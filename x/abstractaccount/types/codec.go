package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcodec "github.com/cosmos/cosmos-sdk/x/gov/codec"
	groupcodec "github.com/cosmos/cosmos-sdk/x/group/codec"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*authtypes.AccountI)(nil), &AbstractAccount{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &NilPubKey{})

	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgRegisterAccount{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUpdateParams{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers the account interfaces and concrete types on the
// provided LegacyAmino codec. These types are used for Amino JSON serialization
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&NilPubKey{}, "abstract-account/NilPubKey", nil)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/feegrant module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/feegrant and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz  and gov Amino codec so that this can later be
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
	RegisterLegacyAminoCodec(groupcodec.Amino)
	RegisterLegacyAminoCodec(authtypes.ModuleCdc.LegacyAmino)
	RegisterLegacyAminoCodec(legacy.Cdc)
}
