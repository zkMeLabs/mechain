package virtualgroup

import (
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	virtualgroupkeeper "github.com/evmos/evmos/v12/x/virtualgroup/keeper"
	virtualgrouptypes "github.com/evmos/evmos/v12/x/virtualgroup/types"

	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	CreateGlobalVirtualGroupGas = 60_000

	CreateGlobalVirtualGroupMethodName = "createGlobalVirtualGroup"

	CreateGlobalVirtualGroupEventName = "CreateGlobalVirtualGroup"
)

// CreateGlobalVirtualGroup defines a method for sp create a global virtual group.
func (c *Contract) CreateGlobalVirtualGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("send method readonly")
	}

	method := MustMethod(CreateGlobalVirtualGroupMethodName)

	var args CreateGlobalVirtualGroupArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &virtualgrouptypes.MsgCreateGlobalVirtualGroup{
		StorageProvider: contract.Caller().String(),
		FamilyId:        args.FamilyID,
		SecondarySpIds:  args.SecondarySpIDs,
		Deposit: sdk.Coin{
			Denom:  args.Deposit.Denom,
			Amount: sdk.NewIntFromBigInt(args.Deposit.Amount),
		},
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := virtualgroupkeeper.NewMsgServerImpl(c.virtualGroupKeeper)
	_, err = server.CreateGlobalVirtualGroup(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(CreateGlobalVirtualGroupEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		big.NewInt(int64(args.FamilyID)),
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
