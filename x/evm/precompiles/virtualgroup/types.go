package virtualgroup

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

type (
	ApprovalJSON = Approval
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
	CoinJSON        = Coin
	PageRequestJSON = PageRequest
)

type CreateGlobalVirtualGroupArgs struct {
	FamilyID       uint32   `abi:"familyId"`
	SecondarySpIDs []uint32 `abi:"secondarySpIds"`
	Deposit        CoinJSON `abi:"deposit"`
}

// Validate CreateGlobalVirtualGroupArgs args
func (args *CreateGlobalVirtualGroupArgs) Validate() error {
	return nil
}

type DeleteGlobalVirtualGroupArgs struct {
	// StorageProvider      string `abi:"storageProvider"`
	GlobalVirtualGroupId uint32 `abi:"globalVirtualGroupId"`
}

// Validate DeleteGlobalVirtualGroupArgs args
func (args *DeleteGlobalVirtualGroupArgs) Validate() error {
	return nil
}

type GlobalVirtualGroupFamiliesArgs struct {
	Pagination PageRequestJSON `abi:"pagination"`
}

// Validate GlobalVirtualGroupFamiliesArgs the args
func (args *GlobalVirtualGroupFamiliesArgs) Validate() error {
	return nil
}

type SwapOutArgs struct {
	GvgFamilyId         uint32       `abi:"gvgFamilyId"`
	GvgIds              []uint32     `abi:"gvgIds"`
	SuccessorSpId       uint32       `abi:"successorSpId"`
	SuccessorSpApproval ApprovalJSON `abi:"successorSpApproval"`
}

// Validate SwapOutArgs the args
func (args *SwapOutArgs) Validate() error {
	return nil
}

type CompleteSwapOutArgs struct {
	// StorageProvider            string   `abi:"storageProvider"`
	GvgFamilyId uint32   `abi:"gvgFamilyId"`
	GvgIds      []uint32 `abi:"gvgIds"`
}

// Validate CompleteSwapOutArgs the args
func (args *CompleteSwapOutArgs) Validate() error {
	return nil
}

type SPExitArgs struct {
	// StorageProvider string `abi:"storageProvider"`
}

// Validate SPExitArgs the args
func (args *SPExitArgs) Validate() error {
	return nil
}

type CompleteSPExitArgs struct {
	StorageProvider string `abi:"storageProvider"`
	Operator        string `abi:"operator"`
}

// Validate CompleteSPExitArgs the args
func (args *CompleteSPExitArgs) Validate() error {
	return nil
}

type DepositArgs struct {
	// StorageProvider      string        `abi:"storageProvider"`
	GlobalVirtualGroupId uint32   `abi:"globalVirtualGroupId"`
	Deposit              CoinJSON `abi:"deposit"`
}

// Validate DepositArgs the args
func (args *DepositArgs) Validate() error {
	return nil
}

type ReserveSwapInArgs struct {
	// StorageProvider            string `abi:"storageProvider"`
	TargetSpId           uint32 `abi:"targetSpId"`
	GvgFamilyId          uint32 `abi:"gvgFamilyId"`
	GlobalVirtualGroupId uint32 `abi:"globalVirtualGroupId"`
}

// Validate ReserveSwapInArgs the args
func (args *ReserveSwapInArgs) Validate() error {
	return nil
}

type CompleteSwapInArgs struct {
	// StorageProvider            string `abi:"storageProvider"`
	GvgFamilyId          uint32 `abi:"gvgFamilyId"`
	GlobalVirtualGroupId uint32 `abi:"globalVirtualGroupId"`
}

// Validate CompleteSwapInArgs the args
func (args *CompleteSwapInArgs) Validate() error {
	return nil
}

type CancelSwapInArgs struct {
	// StorageProvider            string `abi:"storageProvider"`
	GvgFamilyId          uint32 `abi:"gvgFamilyId"`
	GlobalVirtualGroupId uint32 `abi:"globalVirtualGroupId"`
}

// Validate CancelSwapInArgs the args
func (args *CancelSwapInArgs) Validate() error {
	return nil
}

type GlobalVirtualGroupFamilyArgs struct {
	FamilyId uint32 `abi:"familyId"`
}

// Validate GlobalVirtualGroupFamily the args
func (args *GlobalVirtualGroupFamilyArgs) Validate() error {
	return nil
}
