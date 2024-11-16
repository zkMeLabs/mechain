package distribution

import (
	"errors"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	SetWithdrawAddressGas          = 60_000
	WithdrawDelegatorRewardGas     = 60_000
	WithdrawValidatorCommissionGas = 60_000
	FundCommunityPoolGas           = 60_000

	SetWithdrawAddressMethodName          = "setWithdrawAddress"
	WithdrawDelegatorRewardMethodName     = "withdrawDelegatorReward"
	WithdrawValidatorCommissionMethodName = "withdrawValidatorCommission"
	FundCommunityPoolMethodName           = "fundCommunityPool"

	SetWithdrawAddressEventName          = "SetWithdrawAddress"
	WithdrawDelegatorRewardEventName     = "WithdrawDelegatorReward"
	WithdrawValidatorCommissionEventName = "WithdrawValidatorCommission"
	FundCommunityPoolEventName           = "FundCommunityPool"
)

func (c *Contract) SetWithdrawAddress(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, errors.New("only allow EOA can call this method")
	}

	method := MustMethod(SetWithdrawAddressMethodName)

	var args SetWithdrawAddressArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &distributiontypes.MsgSetWithdrawAddress{
		DelegatorAddress: sdk.AccAddress(contract.Caller().Bytes()).String(),
		WithdrawAddress:  sdk.AccAddress(args.WithdrawAddress.Bytes()).String(),
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := distributionkeeper.NewMsgServerImpl(c.distributionKeeper)
	_, err = server.SetWithdrawAddress(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(SetWithdrawAddressEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes()), common.BytesToHash(args.WithdrawAddress.Bytes())},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) WithdrawDelegatorReward(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, errors.New("only allow EOA can call this method")
	}

	method := MustMethod(WithdrawDelegatorRewardMethodName)

	var args ValidatorAddressArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &distributiontypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: sdk.AccAddress(contract.Caller().Bytes()).String(),
		ValidatorAddress: args.ValidatorAddress.String(),
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := distributionkeeper.NewMsgServerImpl(c.distributionKeeper)
	res, err := server.WithdrawDelegatorReward(ctx, msg)
	if err != nil {
		return nil, err
	}

	if evm.Origin != contract.Caller() {
		// ensure that the funds of the contract account in the EVM are consistent with the funds recorded in the bank module account.
		evm.StateDB.AddBalance(contract.Caller(), res.Amount[0].Amount.BigInt())
	}

	if err := c.AddLog(
		evm,
		MustEvent(WithdrawDelegatorRewardEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes()), common.BytesToHash(args.ValidatorAddress.Bytes())},
		res.Amount.String(),
	); err != nil {
		return nil, err
	}

	var rewards []Coin
	for _, amount := range res.Amount {
		rewards = append(rewards, Coin{
			Denom:  amount.Denom,
			Amount: amount.Amount.BigInt(),
		})
	}

	return method.Outputs.Pack(rewards)
}

func (c *Contract) WithdrawValidatorCommission(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, errors.New("only allow EOA can call this method")
	}

	method := MustMethod(WithdrawValidatorCommissionMethodName)

	msg := &distributiontypes.MsgWithdrawValidatorCommission{
		ValidatorAddress: sdk.ValAddress(contract.Caller().Bytes()).String(),
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := distributionkeeper.NewMsgServerImpl(c.distributionKeeper)
	res, err := server.WithdrawValidatorCommission(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(WithdrawValidatorCommissionEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		res.Amount.String(),
	); err != nil {
		return nil, err
	}

	var rewards []Coin
	for _, amount := range res.Amount {
		rewards = append(rewards, Coin{
			Denom:  amount.Denom,
			Amount: amount.Amount.BigInt(),
		})
	}

	return method.Outputs.Pack(rewards)
}

func (c *Contract) FundCommunityPool(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, errors.New("only allow EOA can call this method")
	}

	method := MustMethod(FundCommunityPoolMethodName)

	var args FundCommunityPoolArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var amount []sdk.Coin
	for _, coin := range args.Amount {
		amount = append(amount, sdk.Coin{
			Denom:  coin.Denom,
			Amount: math.NewIntFromBigInt(coin.Amount),
		})
	}
	msg := &distributiontypes.MsgFundCommunityPool{
		Depositor: sdk.AccAddress(contract.Caller().Bytes()).String(),
		Amount:    amount,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := distributionkeeper.NewMsgServerImpl(c.distributionKeeper)
	_, err = server.FundCommunityPool(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(FundCommunityPoolEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		msg.Amount.String(),
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
