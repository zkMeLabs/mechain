package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/evmos/evmos/v12/x/virtualgroup/keeper"
	"github.com/evmos/evmos/v12/x/virtualgroup/types"
)

func SimulateMsgCompleteStorageProviderExit(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, _ *baseapp.BaseApp, _ sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCompleteStorageProviderExit{
			StorageProvider: simAccount.Address.String(),
		}

		// TODO: Handling the CompleteStorageProviderExit simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CompleteStorageProviderExit simulation not implemented"), nil, nil
	}
}
