package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/coreators/interchange/x/myibcdex/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		channelKeeper types.ChannelKeeper
		portKeeper    types.PortKeeper
		scopedKeeper  types.ScopedKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper types.ScopedKeeper,

) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
