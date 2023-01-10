package keeper_test

import (
	"context"
	"testing"

	keepertest "apollo/testutil/keeper"
	"apollo/x/apollo/keeper"
	"apollo/x/apollo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ApolloKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
