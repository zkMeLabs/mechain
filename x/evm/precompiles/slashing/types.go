package slashing

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/evmos/evmos/v12/types"
)

var (
	slashingAddress = common.HexToAddress(types.SlashingAddress)
	slashingABI     = types.MustABIJson(ISlashingMetaData.ABI)
)

func GetAddress() common.Address {
	return slashingAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := slashingABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range slashingABI.Methods {
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
	event := slashingABI.Events[name]
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

type PageRequestJson = PageRequest

type SigningInfoArgs struct {
	ConsAddress common.Address `abi:"consAddress"`
}

// Validate SigningInfo args
func (args *SigningInfoArgs) Validate() error {
	if args.ConsAddress == (common.Address{}) {
		return fmt.Errorf("invalid consensus address: %s", args.ConsAddress)
	}
	return nil
}

// GetConsAddress returns the consensus address, caller must ensure the consensus address is valid
func (args *SigningInfoArgs) GetConsAddress() sdk.ConsAddress {
	consAddress := sdk.ConsAddress(args.ConsAddress.Bytes())
	return consAddress
}

type SigningInfosArgs struct {
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate SigningInfos args
func (args *SigningInfosArgs) Validate() error {
	return nil
}
