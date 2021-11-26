package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/coreators/interchange/testutil/keeper"
	"github.com/coreators/interchange/x/myibcdex/keeper"
	"github.com/coreators/interchange/x/myibcdex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MyibcdexKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
