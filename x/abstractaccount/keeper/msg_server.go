package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/burnt-labs/abstract-account/x/abstractaccount/types"
)

type msgServer struct {
	k Keeper
}

func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{k}
}

// ------------------------------- UpdateParams --------------------------------

func (ms msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if req.Sender != ms.k.authority {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("sender is not authority: expect %s, found %s", ms.k.authority, req.Sender)
	}

	currentParams, err := ms.k.GetParams(ctx)
	if err != nil {
		return nil, err
	}
	if currentParams.BootstrapCodeID != 0 && req.Params.BootstrapCodeID != currentParams.BootstrapCodeID {
		return nil, types.ErrImmutableBootstrap.Wrapf(
			"expected %d, found %d",
			currentParams.BootstrapCodeID,
			req.Params.BootstrapCodeID,
		)
	}
	if err := ms.validateRegistrationCodeIDs(ctx, req.Params); err != nil {
		return nil, err
	}

	if err := ms.k.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

func (ms msgServer) validateRegistrationCodeIDs(ctx sdk.Context, params *types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}
	if !params.RegistrationEnabled() {
		return nil
	}

	if ms.k.vk.GetCodeInfo(ctx, params.BootstrapCodeID) == nil {
		return types.ErrCodeIDNotFound.Wrapf("bootstrap code ID %d", params.BootstrapCodeID)
	}
	if params.ImplementationCodeID != params.BootstrapCodeID && ms.k.vk.GetCodeInfo(ctx, params.ImplementationCodeID) == nil {
		return types.ErrCodeIDNotFound.Wrapf("implementation code ID %d", params.ImplementationCodeID)
	}

	return nil
}

// ------------------------------ RegisterAccount ------------------------------

func (ms msgServer) ensureAccountNamespaceAvailable(
	ctx sdk.Context,
	sender sdk.AccAddress,
	salt []byte,
) (sdk.AccAddress, error) {
	if address, found := ms.k.GetAccountAddress(ctx, sender, salt); found {
		return nil, types.ErrAccountAlreadyRegistered.Wrap(address.String())
	}

	predicted, err := ms.k.PredictAccountAddress(ctx, sender, salt)
	if err != nil {
		return nil, err
	}
	if ms.k.IsAbstractAccount(ctx, predicted) {
		return nil, types.ErrAccountAlreadyRegistered.Wrap(predicted.String())
	}

	return predicted, nil
}

func (ms msgServer) instantiateAccount(
	ctx sdk.Context,
	params *types.Params,
	sender, predicted sdk.AccAddress,
	req *types.MsgRegisterAccount,
) (sdk.AccAddress, []byte, error) {
	contractAddr, data, err := ms.k.ck.Instantiate2(
		ctx,
		params.BootstrapCodeID,
		sender,
		sender,
		req.Msg,
		fmt.Sprintf("%s/%d", types.ModuleName, ms.k.GetAndIncrementNextAccountID(ctx)),
		req.Funds,
		req.Salt,
		false,
	)
	if err != nil {
		return nil, nil, err
	}
	if !contractAddr.Equals(predicted) {
		return nil, nil, fmt.Errorf("instantiated account address %s does not match predicted address %s", contractAddr, predicted)
	}

	// The XION account contract's MigrateMsg is an empty object. Keeping this in
	// the module means callers never need implementation-specific migration data.
	if params.ImplementationCodeID != params.BootstrapCodeID {
		if _, err = ms.k.ck.Migrate(ctx, contractAddr, sender, params.ImplementationCodeID, []byte("{}")); err != nil {
			return nil, nil, err
		}
	}

	return contractAddr, data, nil
}

func (ms msgServer) RegisterAccount(goCtx context.Context, req *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params, err := ms.k.GetParams(ctx)
	if err != nil {
		return nil, err
	}

	if !params.RegistrationEnabled() {
		return nil, types.ErrRegistrationDisabled
	}

	senderAddr, err := sdk.AccAddressFromBech32(req.Sender)
	if err != nil {
		return nil, err
	}
	predictedAddr, err := ms.ensureAccountNamespaceAvailable(ctx, senderAddr, req.Salt)
	if err != nil {
		return nil, err
	}

	contractAddr, data, err := ms.instantiateAccount(ctx, params, senderAddr, predictedAddr, req)
	if err != nil {
		return nil, err
	}

	// set the contract's admin to itself
	if err = ms.k.ck.UpdateContractAdmin(ctx, contractAddr, senderAddr, contractAddr); err != nil {
		return nil, err
	}

	// the contract instantiation should have created a BaseAccount
	acc := ms.k.ak.GetAccount(ctx, contractAddr)
	if _, ok := acc.(*authtypes.BaseAccount); !ok {
		return nil, types.ErrNotBaseAccount
	}

	// we overwrite this BaseAccount with our AbstractAccount
	ms.k.ak.SetAccount(ctx, types.NewAbstractAccountFromAccount(acc))
	ms.k.SetAccountAddress(ctx, senderAddr, req.Salt, contractAddr)

	ms.k.Logger(ctx).Info(
		"account registered",
		types.AttributeKeyCreator, req.Sender,
		types.AttributeKeyBootstrapCodeID, params.BootstrapCodeID,
		types.AttributeKeyCodeID, params.ImplementationCodeID,
		types.AttributeKeyContractAddr, contractAddr.String(),
		types.AttributeKeyAccountNumber, acc.GetAccountNumber(),
	)

	if err = ctx.EventManager().EmitTypedEvent(&types.EventAccountRegistered{
		Creator:         req.Sender,
		CodeID:          params.ImplementationCodeID,
		BootstrapCodeID: params.BootstrapCodeID,
		ContractAddr:    contractAddr.String(),
		AccountNumber:   acc.GetAccountNumber(),
	}); err != nil {
		return nil, err
	}

	return &types.MsgRegisterAccountResponse{Address: contractAddr.String(), Data: data}, nil
}

// ------------------------------- MigrateAccount -----------------------------

func (ms msgServer) MigrateAccount(goCtx context.Context, req *types.MsgMigrateAccount) (*types.MsgMigrateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	address, err := sdk.AccAddressFromBech32(req.Sender)
	if err != nil {
		return nil, err
	}
	if !ms.k.IsAbstractAccount(ctx, address) {
		return nil, types.ErrNotAbstractAccount
	}

	params, err := ms.k.GetParams(ctx)
	if err != nil {
		return nil, err
	}
	if !params.RegistrationEnabled() {
		return nil, types.ErrRegistrationDisabled
	}
	if ms.k.vk.GetCodeInfo(ctx, params.ImplementationCodeID) == nil {
		return nil, types.ErrCodeIDNotFound.Wrapf("implementation code ID %d", params.ImplementationCodeID)
	}

	contractInfo := ms.k.vk.GetContractInfo(ctx, address)
	if contractInfo == nil {
		return nil, types.ErrNotAbstractAccount.Wrap("contract info not found")
	}
	if contractInfo.CodeID == params.ImplementationCodeID {
		return &types.MsgMigrateAccountResponse{Migrated: false}, nil
	}

	data, err := ms.k.ck.Migrate(ctx, address, address, params.ImplementationCodeID, []byte("{}"))
	if err != nil {
		return nil, err
	}
	if err = ctx.EventManager().EmitTypedEvent(&types.EventAccountMigrated{
		ContractAddr: address.String(),
		CodeID:       params.ImplementationCodeID,
	}); err != nil {
		return nil, err
	}

	return &types.MsgMigrateAccountResponse{Migrated: true, Data: data}, nil
}
