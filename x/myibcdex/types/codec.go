package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSendCreatePair{}, "myibcdex/SendCreatePair", nil)
	cdc.RegisterConcrete(&MsgSendSellOrder{}, "myibcdex/SendSellOrder", nil)
	cdc.RegisterConcrete(&MsgSendBuyOrder{}, "myibcdex/SendBuyOrder", nil)
	cdc.RegisterConcrete(&MsgCancelSellOrder{}, "myibcdex/CancelSellOrder", nil)
	cdc.RegisterConcrete(&MsgCancelBuyOrder{}, "myibcdex/CancelBuyOrder", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendCreatePair{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendSellOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendBuyOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelSellOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelBuyOrder{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
