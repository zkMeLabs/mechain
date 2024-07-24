package payment

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/evmos/evmos/v12/x/payment/keeper"
	"github.com/evmos/evmos/v12/x/payment/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the streamRecord
	for _, e := range genState.StreamRecordList {
		elem := e
		k.SetStreamRecord(ctx, &elem)
	}
	// Set all the paymentAccountCount
	for _, e := range genState.PaymentAccountCountList {
		elem := e
		k.SetPaymentAccountCount(ctx, &elem)
	}
	// Set all the paymentAccount
	for _, e := range genState.PaymentAccountList {
		elem := e
		k.SetPaymentAccount(ctx, &elem)
	}
	// Set all the autoSettleRecord
	for _, e := range genState.AutoSettleRecordList {
		elem := e
		k.SetAutoSettleRecord(ctx, &elem)
	}
	err := k.SetParams(ctx, genState.Params)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.StreamRecordList = k.GetAllStreamRecord(ctx)
	genesis.PaymentAccountCountList = k.GetAllPaymentAccountCount(ctx)
	genesis.PaymentAccountList = k.GetAllPaymentAccount(ctx)
	genesis.AutoSettleRecordList = k.GetAllAutoSettleRecord(ctx)

	return genesis
}
