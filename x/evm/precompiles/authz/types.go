package authz

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (args *GrantArgs) StakingParams() (allowed []sdk.AccAddress, denied []sdk.AccAddress, err error) {
	err = fmt.Errorf("authorization input example allowed:0x00000004e1E16f249E2b71c2dc66545215FE9d84,0x1111102Dd32160B064F2A512CDEf74bFdB6a9F96 or denied:0x00000004e1E16f249E2b71c2dc66545215FE9d84, but you input is %s", args.Authorization)

	switch args.AuthzType {
	case AuthzTypeDelegate, AuthzTypeUnbond, AuthzTypeRedelegate:
		// Authorization input example
		// allowed:0x00000004e1E16f249E2b71c2dc66545215FE9d84,0x1111102Dd32160B064F2A512CDEf74bFdB6a9F96
		// or
		// denied:0x00000004e1E16f249E2b71c2dc66545215FE9d84

		authorizationArr := strings.Split(args.Authorization, ":")
		if len(authorizationArr) != 2 {
			return nil, nil, err
		}

		authorizationType := authorizationArr[0]
		authorizationData := authorizationArr[1]

		validatorList := strings.Split(authorizationData, ",")
		if len(validatorList) < 1 {
			return nil, nil, err
		}

		var validators []sdk.AccAddress
		for _, validatorStr := range validatorList {
			validators = append(validators, common.HexToAddress(validatorStr).Bytes())
		}

		if authorizationType == "allowed" {
			return validators, nil, nil
		} else if authorizationType == "denied" {
			return nil, validators, nil
		} else {
			return nil, nil, err
		}
	default:
		return nil, nil, fmt.Errorf("auth type %s not need staking params", args.AuthzType)
	}

	return nil, nil, err
}

func (args *GrantArgs) SendParams() (allowed []sdk.AccAddress, err error) {
	err = fmt.Errorf("authorization input example allowed:0x00000004e1E16f249E2b71c2dc66545215FE9d84,0x1111102Dd32160B064F2A512CDEf74bFdB6a9F96 but you input is %s", args.Authorization)

	switch args.AuthzType {
	case AuthzTypeSend:
		// Authorization input example
		// allowed:0x00000004e1E16f249E2b71c2dc66545215FE9d84,0x1111102Dd32160B064F2A512CDEf74bFdB6a9F96

		authorizationArr := strings.Split(args.Authorization, ":")
		if len(authorizationArr) != 2 {
			return nil, err
		}

		authorizationType := authorizationArr[0]
		authorizationData := authorizationArr[1]

		validatorList := strings.Split(authorizationData, ",")
		if len(validatorList) < 1 {
			return nil, err
		}

		for _, validatorStr := range validatorList {
			allowed = append(allowed, common.HexToAddress(validatorStr).Bytes())
		}

		if authorizationType == "allowed" {
			return allowed, nil
		} else {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("auth type %s not need staking params", args.AuthzType)
	}

	return nil, err
}

type RevokeArgs struct {
	Grantee    common.Address `abi:"grantee"`
	MsgTypeUrl string         `abi:"msgTypeUrl"`
}

// Validate revoke args
func (args *RevokeArgs) Validate() error {
	return nil
}

type ExecArgs struct {
	Msgs []string `abi:"msgs"`
}

// Validate exec args
func (args *ExecArgs) Validate() error {
	if len(args.Msgs) == 0 {
		return errors.New("msgs is empty string")
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
