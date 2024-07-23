package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/evmos/evmos/v12/x/virtualgroup/keeper"
	"github.com/evmos/evmos/v12/x/virtualgroup/types"
)

func SimulateMsgCompleteSwapOut(
	_ types.AccountKeeper,
	bk types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, _ *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCompleteSwapOut{
			StorageProvider: simAccount.Address.String(),
		}

		// TODO: Handling the CompleteSwapOut simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CompleteSwapOut simulation not implemented"), nil, nil
	}
}
