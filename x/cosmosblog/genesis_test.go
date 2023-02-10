package cosmosblog_test

import (
	"testing"

	keepertest "github.com/preet-jain/CosmosBlog/testutil/keeper"
	"github.com/preet-jain/CosmosBlog/testutil/nullify"
	"github.com/preet-jain/CosmosBlog/x/cosmosblog"
	"github.com/preet-jain/CosmosBlog/x/cosmosblog/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosblogKeeper(t)
	cosmosblog.InitGenesis(ctx, *k, genesisState)
	got := cosmosblog.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
