package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/evmos/evmos/v12/x/bridge/keeper"
	"github.com/evmos/evmos/v12/x/bridge/types"
)

func SimulateMsgTransferOut(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(_ *rand.Rand, _ *baseapp.BaseApp, _ sdk.Context, _ []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTransferOut{}

		// TODO: Handling the TransferOut simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TransferOut simulation not implemented"), nil, nil
	}
}
