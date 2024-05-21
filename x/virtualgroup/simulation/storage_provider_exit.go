package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/evmos/evmos/v12/x/virtualgroup/keeper"
	"github.com/evmos/evmos/v12/x/virtualgroup/types"
)

func SimulateMsgStorageProviderExit(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgStorageProviderExit{
			StorageProvider: simAccount.Address.String(),
		}

		// TODO: Handling the StorageProviderExit simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "StorageProviderExit simulation not implemented"), nil, nil
	}
}
