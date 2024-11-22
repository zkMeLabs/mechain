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

type MultiSendArgs struct {
	Outputs []SendArgs `abi:"outputs"`
}

// Validate MultiSend args
func (args *MultiSendArgs) Validate() error {
	if len(args.Outputs) <= 1 {
		return fmt.Errorf("the number of outputs is %v, need to greater than 1", len(args.Outputs))
	}

	for _, output := range args.Outputs {
		for _, deposit := range output.Amount {
			if deposit.Amount.Sign() <= 0 {
				return fmt.Errorf("multiSend %s amount is %s, need to greater than 0", deposit.Denom, deposit.Amount.String())
			}
		}
	}

	return nil
}

type BalanceArgs struct {
	AccountAddress common.Address `abi:"accountAddress"`
	Denom          string         `abi:"denom"`
}

// Validate Balance args
func (args *BalanceArgs) Validate() error {
	if args.Denom == "" {
		return fmt.Errorf("denom is empty")
	}
	return nil
}

type PageRequestJson = PageRequest

type AllBalancesArgs struct {
	AccountAddress common.Address  `abi:"accountAddress"`
	PageRequest    PageRequestJson `abi:"pageRequest"`
}

// Validate AllBalances args
func (args *AllBalancesArgs) Validate() error {
	return nil
}

type SpendableBalancesArgs = AllBalancesArgs

type SpendableBalanceByDenomArgs struct {
	AccountAddress common.Address `abi:"accountAddress"`
	Denom          string         `abi:"denom"`
}

// Validate SpendableBalanceByDenomArgs args
func (args *SpendableBalanceByDenomArgs) Validate() error {
	return nil
}

type TotalSupplyArgs struct {
	PageRequest PageRequestJson `abi:"pageRequest"`
}

// Validate TotalSupplyArgs args
func (args *TotalSupplyArgs) Validate() error {
	return nil
}

type SupplyOfArgs struct {
	Denom string `abi:"denom"`
}

func (args *SupplyOfArgs) Validate() error {
	if args.Denom == "" {
		return fmt.Errorf("denom is empty")
	}
	return nil
}

type DenomMetadataArgs = SupplyOfArgs

type DenomsMetadataArgs struct {
	PageRequest PageRequestJson `abi:"pageRequest"`
}

// Validate DenomsMetadata args
func (args *DenomsMetadataArgs) Validate() error {
	return nil
}

type DenomOwnersArgs struct {
	Denom       string          `abi:"denom"`
	PageRequest PageRequestJson `abi:"pageRequest"`
}

// Validate DenomOwners args
func (args *DenomOwnersArgs) Validate() error {
	if args.Denom == "" {
		return fmt.Errorf("denom is empty")
	}
	return nil
}

type SendEnabledArgs struct {
	Denoms      []string        `abi:"denoms"`
	PageRequest PageRequestJson `abi:"pageRequest"`
}

// Validate SendEnabledArgs args
func (args *SendEnabledArgs) Validate() error {
	return nil
}
