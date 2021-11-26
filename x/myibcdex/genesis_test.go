package myibcdex_test

import (
	"testing"

	keepertest "github.com/coreators/interchange/testutil/keeper"
	"github.com/coreators/interchange/x/myibcdex"
	"github.com/coreators/interchange/x/myibcdex/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MyibcdexKeeper(t)
	myibcdex.InitGenesis(ctx, *k, genesisState)
	got := myibcdex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Equal(t, genesisState.PortId, got.PortId)
	// this line is used by starport scaffolding # genesis/test/assert
}
