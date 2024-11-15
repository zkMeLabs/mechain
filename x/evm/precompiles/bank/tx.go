package bank

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	SendGas      = 60_000
	MultiSendGas = 80_000

	SendMethodName      = "send"
	MultiSendMethodName = "multiSend"

	SendEventName      = "Send"
	MultiSendEventName = "MultiSend"
)

func (c *Contract) Send(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}
	if evm.Origin != contract.Caller() {
		return nil, errors.New("only allow EOA can call this method")
	}

	method := MustMethod(SendMethodName)

	var args SendArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var amount sdk.Coins
	for _, coin := range args.Amount {
		amount = amount.Add(sdk.Coin{
			Denom:  coin.Denom,
			Amount: sdk.NewIntFromBigInt(coin.Amount),
		})
	}

	msg := &banktypes.MsgSend{
		FromAddress: sdk.AccAddress(contract.Caller().Bytes()).String(),
		ToAddress:   sdk.AccAddress(args.ToAddress.Bytes()).String(),
		Amount:      amount,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := bankkeeper.NewMsgServerImpl(c.bankKeeper, c.paymentKeeper)
	_, err = server.Send(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add send log
	if err := c.AddLog(
		evm,
		MustEvent(SendEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes()), common.BytesToHash(args.ToAddress.Bytes())},
		amount.String(),
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// MultiSend defines a method for sending coins from an account to some other accounts.
func (c *Contract) MultiSend(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}
	if evm.Origin != contract.Caller() {
		return nil, errors.New("only allow EOA can call this method")
	}

	method := MustMethod(MultiSendMethodName)

	var args MultiSendArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var totalCoins sdk.Coins
	var outputs []banktypes.Output
	for _, output := range args.Outputs {
		var coins sdk.Coins
		for _, coin := range output.Amount {
			coins = coins.Add(sdk.Coin{
				Denom:  coin.Denom,
				Amount: sdk.NewIntFromBigInt(coin.Amount),
			})
		}

		outputs = append(outputs, banktypes.Output{
			Address: sdk.AccAddress(output.ToAddress.Bytes()).String(),
			Coins:   coins,
		})

		totalCoins = totalCoins.Add(coins.Sort()...)
	}

	msg := &banktypes.MsgMultiSend{
		Inputs: []banktypes.Input{{
			Address: sdk.AccAddress(contract.Caller().Bytes()).String(),
			Coins:   totalCoins,
		}},
		Outputs: outputs,
	}

	if err = msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := bankkeeper.NewMsgServerImpl(c.bankKeeper, c.paymentKeeper)
	_, err = server.MultiSend(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add multi send log
	if err := c.AddLog(
		evm,
		MustEvent(MultiSendEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		totalCoins.String(),
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
