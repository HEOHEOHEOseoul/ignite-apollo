package keeper

import (
	"context"

	"apollo/x/apollo/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Apollo(goCtx context.Context, req *types.QueryApolloRequest) (*types.QueryApolloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryApolloResponse{Text: "Hello, Ignite CLI"}, nil
}
