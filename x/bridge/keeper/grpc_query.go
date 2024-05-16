package keeper

import (
	"github.com/evmos/evmos/v12/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
