package keeper

import (
	"apollo/x/apollo/types"
)

var _ types.QueryServer = Keeper{}
