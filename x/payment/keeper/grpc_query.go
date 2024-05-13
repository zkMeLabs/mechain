package keeper

import (
	"github.com/evmos/evmos/v12/x/payment/types"
)

var _ types.QueryServer = Keeper{}
