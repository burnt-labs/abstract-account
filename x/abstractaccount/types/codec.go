package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "abstract-account/MsgUpdateParams")

	cdc.RegisterConcrete(&NilPubKey{}, "abstract-account/NilPubKey", nil)
	cdc.RegisterConcrete(&Params{}, "abstract-account/Params", nil)

	// try to register the nilpubkey, for some reason it doesn't happen otherwise
	cryptocodec.RegisterCrypto(cdc)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.AccountI)(nil), &AbstractAccount{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &NilPubKey{})

	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgRegisterAccount{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUpdateParams{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
