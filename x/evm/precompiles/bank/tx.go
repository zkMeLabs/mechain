package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	SendGas = 60_000

	SendMethodName = "send"

	SendEventName = "Send"
)

func (c *Contract) Send(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
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

	server := bankkeeper.NewMsgServerImpl(c.bankKeeper, nil)
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
