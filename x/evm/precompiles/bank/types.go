package bank

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	bankAddress = common.HexToAddress(types.BankAddress)
	bankABI     = types.MustABIJson(IBankMetaData.ABI)
)

func GetAddress() common.Address {
	return bankAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := bankABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range bankABI.Methods {
		if bytes.Equal(input[:4], method.ID) {
			return method, nil
		}
	}
	return abi.Method{}, fmt.Errorf("method id %s is not exist", string(input[:4]))
}

func MustMethod(name string) abi.Method {
	method, err := GetMethod(name)
	if err != nil {
		panic(err)
	}
	return method
}

func GetEvent(name string) (abi.Event, error) {
	event := bankABI.Events[name]
	if event.ID == (common.Hash{}) {
		return abi.Event{}, fmt.Errorf("event %s is not exist", name)
	}
	return event, nil
}

func MustEvent(name string) abi.Event {
	event, err := GetEvent(name)
	if err != nil {
		panic(err)
	}
	return event
}

type CoinJson = Coin

type SendArgs struct {
	ToAddress common.Address `abi:"toAddress"`
	Amount    []CoinJson     `abi:"amount"`
}

// Validate Send args
func (args *SendArgs) Validate() error {
	for _, deposit := range args.Amount {
		if deposit.Amount.Sign() <= 0 {
			return fmt.Errorf("send %s amount is %s, need to greater than 0", deposit.Denom, deposit.Amount.String())
		}
	}
	return nil
}
