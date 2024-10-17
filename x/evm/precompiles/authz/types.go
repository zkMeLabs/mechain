package authz

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	authzAddress = common.HexToAddress(types.AuthzAddress)
	authzABI     = types.MustABIJson(IAuthzMetaData.ABI)
)

func GetAddress() common.Address {
	return authzAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := authzABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range authzABI.Methods {
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
	event := authzABI.Events[name]
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

type GrantArgs struct {
	Grantee       common.Address `abi:"grantee"`
	AuthzType     string         `abi:"authzType"`
	Authorization string         `abi:"authorization"`
	Limit         []CoinJson     `abi:"limit"`
	Expiration    int64          `abi:"expiration"`
}

// Validate grant args
func (args *GrantArgs) Validate() error {
	for _, coin := range args.Limit {
		if coin.Amount.Sign() <= 0 {
			return fmt.Errorf("limit %s amount is %s, need to greater than 0", coin.Denom, coin.Amount.String())
		}
	}

	if args.Expiration < 0 {
		return fmt.Errorf("expiration is %d, need to greater than or equal 0", args.Expiration)
	}
	return nil
}

type GrantsArgs struct {
	Granter    common.Address  `abi:"granter"`
	Grantee    common.Address  `abi:"grantee"`
	MsgTypeUrl string          `abi:"msgTypeUrl"`
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate check grants args
func (args *GrantsArgs) Validate() error {
	return nil
}

type GranterGrantsArgs struct {
	Granter    common.Address  `abi:"granter"`
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate check granter grants args
func (args *GranterGrantsArgs) Validate() error {
	return nil
}

type GranteeGrantsArgs struct {
	Grantee    common.Address  `abi:"grantee"`
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate check grantee grants args
func (args *GranteeGrantsArgs) Validate() error {
	return nil
}
