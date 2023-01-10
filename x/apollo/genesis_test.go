package apollo_test

import (
	"testing"

	keepertest "apollo/testutil/keeper"
	"apollo/testutil/nullify"
	"apollo/x/apollo"
	"apollo/x/apollo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ApolloKeeper(t)
	apollo.InitGenesis(ctx, *k, genesisState)
	got := apollo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
