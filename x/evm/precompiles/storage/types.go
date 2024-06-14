package storage

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	storageAddress = common.HexToAddress(types.StorageAddress)
	storageABI     = types.MustABIJson(IStorageMetaData.ABI)
)

func GetAddress() common.Address {
	return storageAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := storageABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range storageABI.Methods {
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
	event := storageABI.Events[name]
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
	ApprovalJson    = Approval
	PageRequestJson = PageRequest
)

type CreateBucketArgs struct {
	BucketName        string         `abi:"bucketName"`
	Visibility        uint8          `abi:"visibility"`
	PaymentAddress    common.Address `abi:"paymentAddress"`
	PrimarySpAddress  common.Address `abi:"primarySpAddress"`
	PrimarySpApproval ApprovalJson   `abi:"primarySpApproval"`
	ChargedReadQuota  uint64         `abi:"chargedReadQuota"`
}

// Validate CreateBucketArgs args
func (args *CreateBucketArgs) Validate() error {
	return nil
}

type ListBucketsArgs struct {
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate ListBucketsArgs the args
func (args *ListBucketsArgs) Validate() error {
	return nil
}

type CreateObjectArgs struct {
	BucketName        string       `abi:"bucketName"`
	ObjectName        string       `abi:"objectName"`
	PayloadSize       uint64       `abi:"payloadSize"`
	Visibility        uint8        `abi:"visibility"`
	ContentType       string       `abi:"contentType"`
	PrimarySpApproval ApprovalJson `abi:"primarySpApproval"`
	ExpectChecksums   []string     `abi:"expectChecksums"`
	RedundancyType    uint8        `abi:"redundancyType"`
}

// Validate CreateObjectArgs args
func (args *CreateObjectArgs) Validate() error {
	return nil
}

type ListObjectsArgs struct {
	Pagination PageRequestJson `abi:"pagination"`
	BucketName string          `abi:"bucketName"`
}

// Validate ListObjectsArgs the args
func (args *ListObjectsArgs) Validate() error {
	return nil
}

type SealObjectArgs struct {
	SealAddress                 common.Address `abi:"sealAddress"`
	BucketName                  string         `abi:"bucketName"`
	ObjectName                  string         `abi:"objectName"`
	GlobalVirtualGroupId        uint32         `abi:"globalVirtualGroupId"`
	SecondarySpBlsAggSignatures string         `abi:"secondarySpBlsAggSignatures"`
}

// Validate SealObjectArgs args
func (args *SealObjectArgs) Validate() error {
	return nil
}
