package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/larry0x/abstract-account/x/abstractaccount/types"
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

	if err := ms.k.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

// ------------------------------ RegisterAccount ------------------------------

func (ms msgServer) RegisterAccount(goCtx context.Context, req *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params, err := ms.k.GetParams(ctx)
	if err != nil {
		return nil, err
	}

	if !params.IsAllowed(req.CodeID) {
		return nil, types.ErrNotAllowedCodeID
	}

	senderAddr, err := sdk.AccAddressFromBech32(req.Sender)
	if err != nil {
		return nil, err
	}

	ms.k.Logger(ctx).Info("registering account",
		"code id", req.CodeID,
		"sender", senderAddr,
		"msg", req.Msg,
		"funds", req.Funds,
		"salt", req.Salt)

	contractAddr, data, err := ms.k.ck.Instantiate2(
		ctx,
		req.CodeID,
		senderAddr,
		senderAddr,
		req.Msg,
		fmt.Sprintf("%s/%d", types.ModuleName, ms.k.GetAndIncrementNextAccountID(ctx)),
		req.Funds,
		req.Salt,
		// we set fix_msg to false because there simply isn't any good reason
		// otherwise, given that we already have full control over the address by
		// providing a salt. read more:
		// https://medium.com/cosmwasm/dev-note-3-limitations-of-instantiate2-and-how-to-deal-with-them-a3f946874230
		false,
	)
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

	ms.k.Logger(ctx).Info(
		"account registered",
		types.AttributeKeyCreator, req.Sender,
		types.AttributeKeyCodeID, req.CodeID,
		types.AttributeKeyContractAddr, contractAddr.String(),
	)

	if err = ctx.EventManager().EmitTypedEvent(&types.EventAccountRegistered{
		Creator:      req.Sender,
		CodeID:       req.CodeID,
		ContractAddr: contractAddr.String(),
	}); err != nil {
		return nil, err
	}

	return &types.MsgRegisterAccountResponse{Address: contractAddr.String(), Data: data}, nil
}
