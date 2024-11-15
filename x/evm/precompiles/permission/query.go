package permission

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
	permissiontypes "github.com/evmos/evmos/v12/x/permission/types"
)

const (
	ParamsMethodName = "params"
)

func (c *Contract) registerQuery() {
	c.registerMethod(ParamsMethodName, 50_000, c.Params, "")
}

// Params queries the parameters of the module.
func (c *Contract) Params(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(ParamsMethodName)
	// parse args
	var args ParamsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &permissiontypes.QueryParamsRequest{}
	res, err := c.permissionKeeper.Params(ctx, msg)
	if err != nil {
		return nil, err
	}
	params := Params{
		MaximumStatementsNum:                  res.Params.MaximumStatementsNum,
		MaximumGroupNum:                       res.Params.MaximumGroupNum,
		MaximumRemoveExpiredPoliciesIteration: res.Params.MaximumRemoveExpiredPoliciesIteration,
	}

	return method.Outputs.Pack(params)
}
