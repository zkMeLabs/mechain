package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/evmos/evmos/v12/x/challenge/keeper"
	"github.com/evmos/evmos/v12/x/challenge/types"
)

func SimulateMsgSubmit(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSubmit{
			Challenger: simAccount.Address.String(),
		}

		// TODO: Handling the Submit simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Submit simulation not implemented"), nil, nil
	}
}
