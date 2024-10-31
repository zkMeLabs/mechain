package payment

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
	paymentkeeper "github.com/evmos/evmos/v12/x/payment/keeper"
	paymenttypes "github.com/evmos/evmos/v12/x/payment/types"

	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	CreatePaymentAccountMethodName = "createPaymentAccount"
	DepositMethodName              = "deposit"
	DisableRefundMethodName        = "disableRefund"
	WithdrawMethodName             = "withdraw"
	UpdateParamsMethodName         = "updateParams"
)

func (c *Contract) registerTx() {
	c.registerMethod(CreatePaymentAccountMethodName, 60_000, c.CreatePaymentAccount, "CreatePaymentAccount")
	c.registerMethod(DepositMethodName, 60_000, c.Deposit, "Deposit")
	c.registerMethod(DisableRefundMethodName, 60_000, c.DisableRefund, "DisableRefund")
	c.registerMethod(WithdrawMethodName, 60_000, c.Withdraw, "Withdraw")
	c.registerMethod(UpdateParamsMethodName, 60_000, c.UpdateParams, "UpdateParams")
}

func (c *Contract) CreatePaymentAccount(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("create payment account method readonly")
	}
	method := GetAbiMethod(CreatePaymentAccountMethodName)
	var args CreatePaymentAccountArgs

	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &paymenttypes.MsgCreatePaymentAccount{
		Creator: contract.Caller().String(),
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := paymentkeeper.NewMsgServerImpl(c.paymentKeeper)
	_, err := server.CreatePaymentAccount(ctx, msg)
	if err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[CreatePaymentAccountMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
		},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) Deposit(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("deposit method readonly")
	}
	method := GetAbiMethod(DepositMethodName)
	var args DepositArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &paymenttypes.MsgDeposit{
		Creator: contract.CallerAddress.String(),
		To:      args.To,
		Amount:  sdk.NewIntFromBigInt(args.Amount),
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := paymentkeeper.NewMsgServerImpl(c.paymentKeeper)
	if _, err := server.Deposit(ctx, msg); err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DepositMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
		},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) DisableRefund(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("disable refund method readonly")
	}

	method := GetAbiMethod(DisableRefundMethodName)

	var args DisableRefundArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &paymenttypes.MsgDisableRefund{
		Owner: contract.Caller().String(),
		Addr:  args.Addr,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := paymentkeeper.NewMsgServerImpl(c.paymentKeeper)
	_, err = server.DisableRefund(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DisableRefundMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) Withdraw(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("withdraw method readonly")
	}

	method := GetAbiMethod(WithdrawMethodName)

	var args WithdrawArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &paymenttypes.MsgWithdraw{
		Creator: contract.Caller().String(),
		From:    args.From,
		Amount:  sdk.NewIntFromBigInt(args.Amount),
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := paymentkeeper.NewMsgServerImpl(c.paymentKeeper)
	_, err = server.Withdraw(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[WithdrawMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
		},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) UpdateParams(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("update params method readonly")
	}

	method := GetAbiMethod(UpdateParamsMethodName)

	var args UpdateParamsArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}
	withdrawTimeLockThreshold := sdk.NewIntFromBigInt(args.Params.WithdrawTimeLockThreshold)
	msg := &paymenttypes.MsgUpdateParams{
		Authority: args.Authority,
		Params: paymenttypes.Params{
			VersionedParams: paymenttypes.VersionedParams{
				ReserveTime:      args.Params.VersionedParams.ReserveTime,
				ValidatorTaxRate: sdk.NewDecFromBigInt(args.Params.VersionedParams.ValidatorTaxRate),
			},
			PaymentAccountCountLimit:  args.Params.PaymentAccountCountLimit,
			ForcedSettleTime:          args.Params.ForcedSettleTime,
			MaxAutoSettleFlowCount:    args.Params.MaxAutoSettleFlowCount,
			MaxAutoResumeFlowCount:    args.Params.MaxAutoResumeFlowCount,
			FeeDenom:                  args.Params.FeeDenom,
			WithdrawTimeLockThreshold: &withdrawTimeLockThreshold,
			WithdrawTimeLockDuration:  args.Params.WithdrawTimeLockDuration,
		},
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := paymentkeeper.NewMsgServerImpl(c.paymentKeeper)
	_, err = server.UpdateParams(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
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
