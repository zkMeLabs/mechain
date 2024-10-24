package storageprovider

import (
	"bytes"
	"fmt"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	spAddress     = common.HexToAddress(types.SpAddress)
	spABI         = types.MustABIJson(IStorageProviderMetaData.ABI)
	invalidMethod = abi.Method{}
)

type (
	PageRequestJson = PageRequest
)

func GetAddress() common.Address {
	return spAddress
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return invalidMethod, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range spABI.Methods {
		if bytes.Equal(input[:4], method.ID) {
			return method, nil
		}
	}
	return invalidMethod, fmt.Errorf("method id %s is not exist", string(input[:4]))
}

func GetAbiMethod(name string) abi.Method {
	return spABI.Methods[name]
}

func GetAbiEvent(name string) abi.Event {
	return spABI.Events[name]
}

type UpdateSPPriceArgs struct {
	ReadPrice     github_com_cosmos_cosmos_sdk_types.Dec `abi:"readPrice"`
	FreeReadQuota uint64                                 `abi:"freeReadQuota"`
	StorePrice    github_com_cosmos_cosmos_sdk_types.Dec `abi:"storePrice"`
}

// Validate UpdateSPPriceArgs the args
func (args *UpdateSPPriceArgs) Validate() error {
	return nil
}

type StorageProviderArgs struct {
	Id uint32 `abi:"id"`
}

// Validate StorageProvider args
func (args *StorageProviderArgs) Validate() error {
	return nil
}

type StorageProvidersArgs struct {
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate StorageProviders the args
func (args *StorageProvidersArgs) Validate() error {
	return nil
}
