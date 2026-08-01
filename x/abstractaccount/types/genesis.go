package types

import (
	"errors"
	"fmt"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewGenesisState(nextAccountID uint64, params *Params) *GenesisState {
	return &GenesisState{
		NextAccountId: nextAccountID,
		Params:        params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(1, DefaultParams())
}

func (gs *GenesisState) Validate() error {
	if gs.Params == nil {
		return errors.New("params cannot be nil")
	}
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	seen := make(map[string]struct{}, len(gs.AccountAddresses))
	for _, entry := range gs.AccountAddresses {
		if entry == nil {
			return errors.New("account address registry entry cannot be nil")
		}
		if _, err := sdk.AccAddressFromBech32(entry.Sender); err != nil {
			return fmt.Errorf("invalid account address registry sender: %w", err)
		}
		if _, err := sdk.AccAddressFromBech32(entry.Address); err != nil {
			return fmt.Errorf("invalid account address registry address: %w", err)
		}
		if err := wasmtypes.ValidateSalt(entry.Salt); err != nil {
			return fmt.Errorf("invalid account address registry salt: %w", err)
		}

		key := entry.Sender + "\x00" + string(entry.Salt)
		if _, ok := seen[key]; ok {
			return errors.New("duplicate account address registry entry")
		}
		seen[key] = struct{}{}
	}

	return nil
}
