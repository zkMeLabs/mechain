package slashing

import (
	"github.com/evmos/evmos/v12/x/evm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	UnjailGas = 60_000

	UnjailMethodName = "unjail"

	UnjailEventName = "Unjail"
)

func (c *Contract) Unjail(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(UnjailMethodName)

	msg := &slashingtypes.MsgUnjail{
		ValidatorAddr: sdk.ValAddress(contract.Caller().Bytes()).String(),
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := slashingkeeper.NewMsgServerImpl(c.slashingkeeper)

	_, err := server.Unjail(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add undelegate log
	if err := c.AddLog(
		evm,
		MustEvent(UnjailEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
