package virtualgroup

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	virtualgroupkeeper "github.com/evmos/evmos/v12/x/virtualgroup/keeper"
	virtualgrouptypes "github.com/evmos/evmos/v12/x/virtualgroup/types"

	mechaincommon "github.com/evmos/evmos/v12/types/common"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	CreateGlobalVirtualGroupGas = 60_000
	DeleteGlobalVirtualGroupGas = 60_000
	SwapOutGas                  = 60_000
	CompleteSwapOutGas          = 60_000
	SPExitGas                   = 60_000
	CompleteSPExitGas           = 60_000
	DepositGas                  = 60_000
	ReserveSwapInGas            = 60_000
	CompleteSwapInGas           = 60_000
	CancelSwapInGas             = 60_000

	CreateGlobalVirtualGroupMethodName = "createGlobalVirtualGroup"
	DeleteGlobalVirtualGroupMethodName = "deleteGlobalVirtualGroup"
	SwapOutMethodName                  = "swapOut"
	CompleteSwapOutMethodName          = "completeSwapOut"
	SPExitMethodName                   = "sPExit"
	CompleteSPExitMethodName           = "completeSPExit"
	DepositMethodName                  = "deposit"
	ReserveSwapInMethodName            = "reserveSwapIn"
	CompleteSwapInMethodName           = "completeSwapIn"
	CancelSwapInMethodName             = "cancelSwapIn"

	CreateGlobalVirtualGroupEventName = "CreateGlobalVirtualGroup"
	DeleteGlobalVirtualGroupEventName = "DeleteGlobalVirtualGroup"
	SwapOutEventName                  = "SwapOut"
	CompleteSwapOutEventName          = "CompleteSwapOut"
	SPExitEventName                   = "SPExit"
	CompleteSPExitEventName           = "CompleteSPExit"
	DepositEventName                  = "Deposit"
	ReserveSwapInEventName            = "ReserveSwapIn"
	CompleteSwapInEventName           = "CompleteSwapIn"
	CancelSwapInEventName             = "CancelSwapIn"
)

// CreateGlobalVirtualGroup defines a method for sp create a global virtual group.
func (c *Contract) CreateGlobalVirtualGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
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

// DeleteGlobalVirtualGroup defines a method for sp delete a global virtual group.
func (c *Contract) DeleteGlobalVirtualGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(DeleteGlobalVirtualGroupMethodName)

	var args DeleteGlobalVirtualGroupArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &virtualgrouptypes.MsgDeleteGlobalVirtualGroup{
		StorageProvider:      contract.Caller().String(),
		GlobalVirtualGroupId: args.GlobalVirtualGroupId,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := virtualgroupkeeper.NewMsgServerImpl(c.virtualGroupKeeper)
	_, err = server.DeleteGlobalVirtualGroup(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(DeleteGlobalVirtualGroupEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		big.NewInt(int64(args.GlobalVirtualGroupId)),
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// SwapOut defines a method for sp to remove itself from all Virtual Groups.
func (c *Contract) SwapOut(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(SwapOutMethodName)

	var args SwapOutArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &virtualgrouptypes.MsgSwapOut{
		StorageProvider:            contract.Caller().String(),
		GlobalVirtualGroupFamilyId: args.GlobalVirtualGroupFamilyId,
		GlobalVirtualGroupIds:      args.GlobalVirtualGroupIds,
		SuccessorSpId:              args.SuccessorSpId,
		SuccessorSpApproval: &mechaincommon.Approval{
			ExpiredHeight:              args.SuccessorSpApproval.ExpiredHeight,
			GlobalVirtualGroupFamilyId: args.SuccessorSpApproval.GlobalVirtualGroupFamilyId,
			Sig:                        args.SuccessorSpApproval.Sig,
		},
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := virtualgroupkeeper.NewMsgServerImpl(c.virtualGroupKeeper)
	_, err = server.SwapOut(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(SwapOutEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		big.NewInt(int64(args.GlobalVirtualGroupFamilyId)),
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// CompleteSwapOut defines a method for sp somplete to remove itself from all Virtual Groups.
func (c *Contract) CompleteSwapOut(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(CompleteSwapOutMethodName)

	var args CompleteSwapOutArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &virtualgrouptypes.MsgCompleteSwapOut{
		StorageProvider:            contract.Caller().String(),
		GlobalVirtualGroupFamilyId: args.GlobalVirtualGroupFamilyId,
		GlobalVirtualGroupIds:      args.GlobalVirtualGroupIds,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := virtualgroupkeeper.NewMsgServerImpl(c.virtualGroupKeeper)
	_, err = server.CompleteSwapOut(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(SwapOutEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		big.NewInt(int64(args.GlobalVirtualGroupFamilyId)),
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// SPExit defines a method for sp to exit.
func (c *Contract) SPExit(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(SPExitMethodName)

	var args SPExitArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &virtualgrouptypes.MsgStorageProviderExit{
		StorageProvider: contract.Caller().String(),
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := virtualgroupkeeper.NewMsgServerImpl(c.virtualGroupKeeper)
	_, err = server.StorageProviderExit(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(SPExitEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// CompleteSPExit defines a method for sp complete to exit.
func (c *Contract) CompleteSPExit(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(CompleteSPExitMethodName)

	var args CompleteSPExitArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &virtualgrouptypes.MsgCompleteStorageProviderExit{
		StorageProvider: args.StorageProvider,
		Operator:        args.Operator,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := virtualgroupkeeper.NewMsgServerImpl(c.virtualGroupKeeper)
	_, err = server.CompleteStorageProviderExit(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(CompleteSPExitEventName),
		[]common.Hash{common.BytesToHash([]byte(args.StorageProvider))},
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// Deposit defines a method to deposit more tokens for the objects stored on it.
func (c *Contract) Deposit(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(DepositMethodName)

	var args DepositArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &virtualgrouptypes.MsgDeposit{
		StorageProvider:      contract.Caller().String(),
		GlobalVirtualGroupId: args.GlobalVirtualGroupId,
		Deposit: sdk.Coin{
			Denom:  args.Deposit.Denom,
			Amount: sdk.NewIntFromBigInt(args.Deposit.Amount),
		},
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := virtualgroupkeeper.NewMsgServerImpl(c.virtualGroupKeeper)
	_, err = server.Deposit(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(DepositEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// ReserveSwapIn defines a method to deposit more tokens for the objects stored on it.
func (c *Contract) ReserveSwapIn(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(ReserveSwapInMethodName)

	var args ReserveSwapInArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &virtualgrouptypes.MsgReserveSwapIn{
		StorageProvider:            contract.Caller().String(),
		TargetSpId:                 args.TargetSpId,
		GlobalVirtualGroupFamilyId: args.GlobalVirtualGroupFamilyId,
		GlobalVirtualGroupId:       args.GlobalVirtualGroupId,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := virtualgroupkeeper.NewMsgServerImpl(c.virtualGroupKeeper)
	_, err = server.ReserveSwapIn(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(ReserveSwapInEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// ReserveSwapIn defines a method to deposit more tokens for the objects stored on it.
func (c *Contract) CompleteSwapIn(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(CompleteSwapInMethodName)

	var args CompleteSwapInArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &virtualgrouptypes.MsgCompleteSwapIn{
		StorageProvider:            contract.Caller().String(),
		GlobalVirtualGroupFamilyId: args.GlobalVirtualGroupFamilyId,
		GlobalVirtualGroupId:       args.GlobalVirtualGroupId,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := virtualgroupkeeper.NewMsgServerImpl(c.virtualGroupKeeper)
	_, err = server.CompleteSwapIn(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(CompleteSwapInEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// CancelSwapIn defines a method to deposit more tokens for the objects stored on it.
func (c *Contract) CancelSwapIn(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(CancelSwapInMethodName)

	var args CancelSwapInArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &virtualgrouptypes.MsgCancelSwapIn{
		StorageProvider:            contract.Caller().String(),
		GlobalVirtualGroupFamilyId: args.GlobalVirtualGroupFamilyId,
		GlobalVirtualGroupId:       args.GlobalVirtualGroupId,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := virtualgroupkeeper.NewMsgServerImpl(c.virtualGroupKeeper)
	_, err = server.CancelSwapIn(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(CancelSwapInEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
