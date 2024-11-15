package permission

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	permissionAddress = common.HexToAddress(types.PermissionAddress)
	permissionABI     = types.MustABIJson(IPermissionMetaData.ABI)
	invalidMethod     = abi.Method{}
)

func GetAddress() common.Address {
	return permissionAddress
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return invalidMethod, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range permissionABI.Methods {
		if bytes.Equal(input[:4], method.ID) {
			return method, nil
		}
	}
	return invalidMethod, fmt.Errorf("method id %s is not exist", string(input[:4]))
}

func GetAbiMethod(name string) abi.Method {
	return permissionABI.Methods[name]
}

func GetAbiEvent(name string) abi.Event {
	return permissionABI.Events[name]
}

type UpdateParamsArgs struct {
	// Operator string `abi:"operator"`
	Authority string `abi:"authority"`
	Params    Params `abi:"params"`
}

// Validate UpdateParamsArgs the args
func (args *UpdateParamsArgs) Validate() error {
	return nil
}

type ParamsArgs struct {
	// Operator string `abi:"operator"`
}

// Validate ParamsArgs the args
func (args *ParamsArgs) Validate() error {
	return nil
}
