package sp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/evmos/evmos/v12/x/sp/keeper"
	"github.com/evmos/evmos/v12/x/sp/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	err := k.SetParams(ctx, genState.Params)
	if err != nil {
		panic(err)
	}

	k.InitGenesis(ctx, genState)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.StorageProviders = k.GetAllStorageProviders(ctx)
	genesis.SpStoragePriceList = k.GetAllSpStoragePrice(ctx)

	return genesis
}
