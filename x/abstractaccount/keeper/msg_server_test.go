package keeper_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	"github.com/larry0x/abstract-account/simapp"
	simapptesting "github.com/larry0x/abstract-account/simapp/testing"
	"github.com/larry0x/abstract-account/x/abstractaccount/keeper"
	"github.com/larry0x/abstract-account/x/abstractaccount/testdata"
	"github.com/larry0x/abstract-account/x/abstractaccount/types"
)

type AccountInitMsg struct {
	PubKey []byte `json:"pubkey"`
}

var (
	user               = simapptesting.MakeRandomAddress()
	userInitialBalance = sdk.NewCoins(sdk.NewCoin(simapptesting.DefaultBondDenom, math.NewInt(123456)))
	acctRegisterFunds  = sdk.NewCoins(sdk.NewCoin(simapptesting.DefaultBondDenom, math.NewInt(88888)))
)

// ------------------------------- UpdateParams --------------------------------

func TestUpdateParams(t *testing.T) {
	for _, tc := range []struct {
		desc      string
		sender    string
		newParams *types.Params
		expErr    bool
	}{
		{
			desc:      "sender is not authority",
			sender:    user.String(),
			newParams: types.DefaultParams(),
			expErr:    true,
		},
		{
			desc:      "invalid params",
			sender:    simapp.Authority,
			newParams: &types.Params{MaxGasBefore: 88888, MaxGasAfter: 0},
			expErr:    true,
		},
		{
			desc:      "sender is authority and params are valid",
			sender:    simapp.Authority,
			newParams: &types.Params{MaxGasBefore: 88888, MaxGasAfter: 99999},
			expErr:    false,
		},
	} {
		app := simapptesting.MakeMockApp([]banktypes.Balance{})
		ctx := app.NewContext(false)

		msgServer := keeper.NewMsgServerImpl(app.AbstractAccountKeeper)

		paramsBefore, err1 := app.AbstractAccountKeeper.GetParams(ctx)
		require.NoError(t, err1)

		_, err2 := msgServer.UpdateParams(ctx, &types.MsgUpdateParams{
			Sender: tc.sender,
			Params: tc.newParams,
		})

		paramsAfter, err3 := app.AbstractAccountKeeper.GetParams(ctx)
		require.NoError(t, err3)

		if tc.expErr {
			require.Error(t, err2)
			require.Equal(t, paramsBefore, paramsAfter)
		} else {
			require.NoError(t, err2)
			require.Equal(t, tc.newParams, paramsAfter)
		}
	}
}

// ------------------------------ RegisterAccount ------------------------------

func TestRegisterAccount(t *testing.T) {
	for _, tc := range []struct {
		allowAllCodeIDs bool
		allowedCodeIDs  []uint64
		expOk           bool
	}{
		{
			allowAllCodeIDs: true,
			allowedCodeIDs:  []uint64{},
			expOk:           true,
		},
		{
			allowAllCodeIDs: false,
			allowedCodeIDs:  []uint64{1, 69, 420},
			expOk:           true,
		},
		{
			allowAllCodeIDs: false,
			allowedCodeIDs:  []uint64{69, 420},
			expOk:           false,
		},
	} {
		app := simapptesting.MakeMockApp([]banktypes.Balance{
			{
				Address: user.String(),
				Coins:   userInitialBalance,
			},
		})

		ctx := app.NewContext(false).WithBlockTime(time.Now())

		params, err := types.NewParams(tc.allowAllCodeIDs, tc.allowedCodeIDs, types.DefaultMaxGas, types.DefaultMaxGas)
		require.NoError(t, err)

		k := app.AbstractAccountKeeper
		k.SetParams(ctx, params)

		// store code
		codeID, err := storeCode(ctx, k.ContractKeeper())
		require.NoError(t, err)
		require.Equal(t, uint64(1), codeID)

		// register account
		accAddr, err := registerAccount(ctx, keeper.NewMsgServerImpl(k), codeID)

		if tc.expOk {
			require.NoError(t, err)

			// check the contract info is correct
			contractInfo := app.WasmKeeper.GetContractInfo(ctx, accAddr)
			require.Equal(t, codeID, contractInfo.CodeID)
			require.Equal(t, user.String(), contractInfo.Creator)
			require.Equal(t, accAddr.String(), contractInfo.Admin)
			require.Equal(t, fmt.Sprintf("%s/%d", types.ModuleName, k.GetNextAccountID(ctx)-1), contractInfo.Label)

			// make sure an AbstractAccount has been created
			_, ok := app.AccountKeeper.GetAccount(ctx, accAddr).(*types.AbstractAccount)
			require.True(t, ok)

			// make sure the contract has received the funds
			balance := app.BankKeeper.GetAllBalances(ctx, accAddr)
			require.Equal(t, acctRegisterFunds, balance)
		} else {
			require.Error(t, err)
		}
	}
}

// ---------------------------------- Helpers ----------------------------------

func storeCode(ctx sdk.Context, contractKeeper wasmtypes.ContractOpsKeeper) (uint64, error) {
	codeID, _, err := contractKeeper.Create(ctx, user, testdata.AccountWasm, nil)

	return codeID, err
}

func registerAccount(ctx sdk.Context, msgServer types.MsgServer, codeID uint64) (sdk.AccAddress, error) {
	msgBytes, err := json.Marshal(&AccountInitMsg{
		PubKey: simapptesting.MakeRandomPubKey().Bytes(),
	})
	if err != nil {
		return nil, err
	}

	res, err := msgServer.RegisterAccount(ctx, &types.MsgRegisterAccount{
		Sender: user.String(),
		CodeID: codeID,
		Msg:    msgBytes,
		Funds:  acctRegisterFunds,
		Salt:   []byte("hello"),
	})
	if err != nil {
		return nil, err
	}

	return sdk.AccAddressFromBech32(res.Address)
}
