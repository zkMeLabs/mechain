package permission

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
	permissionkeeper "github.com/evmos/evmos/v12/x/permission/keeper"
	permissiontypes "github.com/evmos/evmos/v12/x/permission/types"

	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	UpdateParamsMethodName = "updateParams"
)

func (c *Contract) registerTx() {
	c.registerMethod(UpdateParamsMethodName, 60_000, c.UpdateParams, "UpdateParams")
}

func (c *Contract) UpdateParams(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("update params method readonly")
	}
	method := GetAbiMethod(UpdateParamsMethodName)
	var args UpdateParamsArgs

	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &permissiontypes.MsgUpdateParams{
		Authority: args.Authority,
		Params: permissiontypes.Params{
			MaximumStatementsNum:                  args.Params.MaximumStatementsNum,
			MaximumGroupNum:                       args.Params.MaximumGroupNum,
			MaximumRemoveExpiredPoliciesIteration: args.Params.MaximumRemoveExpiredPoliciesIteration,
		},
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := permissionkeeper.NewMsgServerImpl(c.permissionKeeper)
	_, err := server.UpdateParams(ctx, msg)
	if err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[UpdateParamsMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
		},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
