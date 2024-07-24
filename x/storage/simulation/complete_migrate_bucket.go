package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/evmos/evmos/v12/x/storage/keeper"
	"github.com/evmos/evmos/v12/x/storage/types"
)

func SimulateMsgCompleteMigrateBucket(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, _ *baseapp.BaseApp, _ sdk.Context, accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCompleteMigrateBucket{
			Operator: simAccount.Address.String(),
		}

		// TODO: Handling the CompleteMigrateBucket simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CompleteMigrateBucket simulation not implemented"), nil, nil
	}
}
