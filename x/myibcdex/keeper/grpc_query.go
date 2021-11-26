package keeper

import (
	"github.com/coreators/interchange/x/myibcdex/types"
)

var _ types.QueryServer = Keeper{}
