package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/preet-jain/CosmosBlog/testutil/keeper"
	"github.com/preet-jain/CosmosBlog/x/cosmosblog/keeper"
	"github.com/preet-jain/CosmosBlog/x/cosmosblog/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmosblogKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
