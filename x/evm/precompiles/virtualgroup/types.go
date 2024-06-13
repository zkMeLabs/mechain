package virtualgroup

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	virtualGroupAddress = common.HexToAddress(types.VirtualGroupAddress)
	virtualGroupABI     = types.MustABIJson(IVirtualGroupMetaData.ABI)
)

func GetAddress() common.Address {
	return virtualGroupAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := virtualGroupABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range virtualGroupABI.Methods {
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
	event := virtualGroupABI.Events[name]
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

type (
	CoinJson        = Coin
	PageRequestJson = PageRequest
)

type CreateGlobalVirtualGroupArgs struct {
	FamilyId       uint32   `abi:"familyId"`
	SecondarySpIds []uint32 `abi:"secondarySpIds"`
	Deposit        CoinJson `abi:"deposit"`
}

// Validate CreateGlobalVirtualGroupArgs args
func (args *CreateGlobalVirtualGroupArgs) Validate() error {
	return nil
}

type GlobalVirtualGroupFamiliesArgs struct {
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate GlobalVirtualGroupFamiliesArgs the args
func (args *GlobalVirtualGroupFamiliesArgs) Validate() error {
	return nil
}
