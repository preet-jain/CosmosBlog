package keeper_test

import (
	"testing"

	testkeeper "github.com/preet-jain/CosmosBlog/testutil/keeper"
	"github.com/preet-jain/CosmosBlog/x/cosmosblog/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmosblogKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
