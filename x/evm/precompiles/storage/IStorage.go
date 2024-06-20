// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Approval is an auto generated low-level Go binding around an user-defined struct.
type Approval struct {
	ExpiredHeight              uint64
	GlobalVirtualGroupFamilyId uint32
	Sig                        []byte
}

// BucketExtraInfo is an auto generated low-level Go binding around an user-defined struct.
type BucketExtraInfo struct {
	IsRateLimited   bool
	FlowRateLimit   *big.Int
	CurrentFlowRate *big.Int
}

// BucketInfo is an auto generated low-level Go binding around an user-defined struct.
type BucketInfo struct {
	Owner                      common.Address
	BucketName                 string
	Visibility                 uint8
	Id                         *big.Int
	SourceType                 uint8
	CreateAt                   int64
	PaymentAddress             common.Address
	GlobalVirtualGroupFamilyId uint32
	ChargedReadQuota           uint64
	BucketStatus               uint8
	Tags                       []Tag
	SpAsDelegatedAgentDisabled bool
}

// GroupInfo is an auto generated low-level Go binding around an user-defined struct.
type GroupInfo struct {
	Owner      common.Address
	GroupName  string
	SourceType uint8
	Id         *big.Int
	Extra      string
	Tags       []Tag
}

// ObjectInfo is an auto generated low-level Go binding around an user-defined struct.
type ObjectInfo struct {
	Owner               common.Address
	Creator             common.Address
	BucketName          string
	ObjectName          string
	Id                  *big.Int
	LocalVirtualGroupId uint32
	PayloadSize         uint64
	Visibility          uint8
	ContentType         string
	CreateAt            int64
	ObjectStatus        uint8
	RedundancyType      uint8
	SourceType          uint8
	Checksums           []string
	Tags                []Tag
	IsUpdating          bool
	UpdatedAt           int64
	UpdatedBy           common.Address
	Version             int64
}

// PageRequest is an auto generated low-level Go binding around an user-defined struct.
type PageRequest struct {
	Key        []byte
	Offset     uint64
	Limit      uint64
	CountTotal bool
	Reverse    bool
}

// PageResponse is an auto generated low-level Go binding around an user-defined struct.
type PageResponse struct {
	NextKey []byte
	Total   uint64
}

// Tag is an auto generated low-level Go binding around an user-defined struct.
type Tag struct {
	Key   string
	Value string
}

// IStorageMetaData contains all meta data concerning the IStorage contract.
var IStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"primarySpAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CreateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CreateGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CreateObject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sealAddress\",\"type\":\"address\"}],\"name\":\"SealObject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sealAddress\",\"type\":\"address\"}],\"name\":\"SealObjectV2\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primarySpAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"expiredHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structApproval\",\"name\":\"primarySpApproval\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"chargedReadQuota\",\"type\":\"uint64\"}],\"name\":\"createBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"}],\"name\":\"createGroup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"expiredHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structApproval\",\"name\":\"primarySpApproval\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"expectChecksums\",\"type\":\"string[]\"},{\"internalType\":\"enumRedundancyType\",\"name\":\"redundancyType\",\"type\":\"uint8\"}],\"name\":\"createObject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"headBucket\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"int64\",\"name\":\"createAt\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chargedReadQuota\",\"type\":\"uint64\"},{\"internalType\":\"enumBucketStatus\",\"name\":\"bucketStatus\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"spAsDelegatedAgentDisabled\",\"type\":\"bool\"}],\"internalType\":\"structBucketInfo\",\"name\":\"bucketInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"IsRateLimited\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"FlowRateLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CurrentFlowRate\",\"type\":\"uint256\"}],\"internalType\":\"structBucketExtraInfo\",\"name\":\"bucketExtraInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"listBuckets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"int64\",\"name\":\"createAt\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chargedReadQuota\",\"type\":\"uint64\"},{\"internalType\":\"enumBucketStatus\",\"name\":\"bucketStatus\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"spAsDelegatedAgentDisabled\",\"type\":\"bool\"}],\"internalType\":\"structBucketInfo[]\",\"name\":\"bucketInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"groupOwner\",\"type\":\"string\"}],\"name\":\"listGroups\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"}],\"internalType\":\"structGroupInfo[]\",\"name\":\"groupInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"listObjects\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"localVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"createAt\",\"type\":\"int64\"},{\"internalType\":\"enumObjectStatus\",\"name\":\"objectStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumRedundancyType\",\"name\":\"redundancyType\",\"type\":\"uint8\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"checksums\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"isUpdating\",\"type\":\"bool\"},{\"internalType\":\"int64\",\"name\":\"updatedAt\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"updatedBy\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"version\",\"type\":\"int64\"}],\"internalType\":\"structObjectInfo[]\",\"name\":\"objectInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sealAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"secondarySpBlsAggSignatures\",\"type\":\"string\"}],\"name\":\"sealObject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sealAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"secondarySpBlsAggSignatures\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"expectChecksums\",\"type\":\"string[]\"}],\"name\":\"sealObjectV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use IStorageMetaData.ABI instead.
var IStorageABI = IStorageMetaData.ABI

// IStorage is an auto generated Go binding around an Ethereum contract.
type IStorage struct {
	IStorageCaller     // Read-only binding to the contract
	IStorageTransactor // Write-only binding to the contract
	IStorageFilterer   // Log filterer for contract events
}

// IStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStorageSession struct {
	Contract     *IStorage         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStorageCallerSession struct {
	Contract *IStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStorageTransactorSession struct {
	Contract     *IStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStorageRaw struct {
	Contract *IStorage // Generic contract binding to access the raw methods on
}

// IStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStorageCallerRaw struct {
	Contract *IStorageCaller // Generic read-only contract binding to access the raw methods on
}

// IStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStorageTransactorRaw struct {
	Contract *IStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStorage creates a new instance of IStorage, bound to a specific deployed contract.
func NewIStorage(address common.Address, backend bind.ContractBackend) (*IStorage, error) {
	contract, err := bindIStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStorage{IStorageCaller: IStorageCaller{contract: contract}, IStorageTransactor: IStorageTransactor{contract: contract}, IStorageFilterer: IStorageFilterer{contract: contract}}, nil
}

// NewIStorageCaller creates a new read-only instance of IStorage, bound to a specific deployed contract.
func NewIStorageCaller(address common.Address, caller bind.ContractCaller) (*IStorageCaller, error) {
	contract, err := bindIStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStorageCaller{contract: contract}, nil
}

// NewIStorageTransactor creates a new write-only instance of IStorage, bound to a specific deployed contract.
func NewIStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*IStorageTransactor, error) {
	contract, err := bindIStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStorageTransactor{contract: contract}, nil
}

// NewIStorageFilterer creates a new log filterer instance of IStorage, bound to a specific deployed contract.
func NewIStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*IStorageFilterer, error) {
	contract, err := bindIStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStorageFilterer{contract: contract}, nil
}

// bindIStorage binds a generic wrapper to an already deployed contract.
func bindIStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStorage *IStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStorage.Contract.IStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStorage *IStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStorage.Contract.IStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStorage *IStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStorage.Contract.IStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStorage *IStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStorage *IStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStorage *IStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStorage.Contract.contract.Transact(opts, method, params...)
}

// HeadBucket is a free data retrieval call binding the contract method 0x0d5269af.
//
// Solidity: function headBucket(string bucketName) view returns((address,string,uint8,uint256,uint8,int64,address,uint32,uint64,uint8,(string,string)[],bool) bucketInfo, (bool,uint256,uint256) bucketExtraInfo)
func (_IStorage *IStorageCaller) HeadBucket(opts *bind.CallOpts, bucketName string) (struct {
	BucketInfo      BucketInfo
	BucketExtraInfo BucketExtraInfo
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "headBucket", bucketName)

	outstruct := new(struct {
		BucketInfo      BucketInfo
		BucketExtraInfo BucketExtraInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BucketInfo = *abi.ConvertType(out[0], new(BucketInfo)).(*BucketInfo)
	outstruct.BucketExtraInfo = *abi.ConvertType(out[1], new(BucketExtraInfo)).(*BucketExtraInfo)

	return *outstruct, err

}

// HeadBucket is a free data retrieval call binding the contract method 0x0d5269af.
//
// Solidity: function headBucket(string bucketName) view returns((address,string,uint8,uint256,uint8,int64,address,uint32,uint64,uint8,(string,string)[],bool) bucketInfo, (bool,uint256,uint256) bucketExtraInfo)
func (_IStorage *IStorageSession) HeadBucket(bucketName string) (struct {
	BucketInfo      BucketInfo
	BucketExtraInfo BucketExtraInfo
}, error) {
	return _IStorage.Contract.HeadBucket(&_IStorage.CallOpts, bucketName)
}

// HeadBucket is a free data retrieval call binding the contract method 0x0d5269af.
//
// Solidity: function headBucket(string bucketName) view returns((address,string,uint8,uint256,uint8,int64,address,uint32,uint64,uint8,(string,string)[],bool) bucketInfo, (bool,uint256,uint256) bucketExtraInfo)
func (_IStorage *IStorageCallerSession) HeadBucket(bucketName string) (struct {
	BucketInfo      BucketInfo
	BucketExtraInfo BucketExtraInfo
}, error) {
	return _IStorage.Contract.HeadBucket(&_IStorage.CallOpts, bucketName)
}

// ListBuckets is a free data retrieval call binding the contract method 0x60c8f8d2.
//
// Solidity: function listBuckets((bytes,uint64,uint64,bool,bool) pagination) view returns((address,string,uint8,uint256,uint8,int64,address,uint32,uint64,uint8,(string,string)[],bool)[] bucketInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageCaller) ListBuckets(opts *bind.CallOpts, pagination PageRequest) (struct {
	BucketInfos  []BucketInfo
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "listBuckets", pagination)

	outstruct := new(struct {
		BucketInfos  []BucketInfo
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BucketInfos = *abi.ConvertType(out[0], new([]BucketInfo)).(*[]BucketInfo)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// ListBuckets is a free data retrieval call binding the contract method 0x60c8f8d2.
//
// Solidity: function listBuckets((bytes,uint64,uint64,bool,bool) pagination) view returns((address,string,uint8,uint256,uint8,int64,address,uint32,uint64,uint8,(string,string)[],bool)[] bucketInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageSession) ListBuckets(pagination PageRequest) (struct {
	BucketInfos  []BucketInfo
	PageResponse PageResponse
}, error) {
	return _IStorage.Contract.ListBuckets(&_IStorage.CallOpts, pagination)
}

// ListBuckets is a free data retrieval call binding the contract method 0x60c8f8d2.
//
// Solidity: function listBuckets((bytes,uint64,uint64,bool,bool) pagination) view returns((address,string,uint8,uint256,uint8,int64,address,uint32,uint64,uint8,(string,string)[],bool)[] bucketInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageCallerSession) ListBuckets(pagination PageRequest) (struct {
	BucketInfos  []BucketInfo
	PageResponse PageResponse
}, error) {
	return _IStorage.Contract.ListBuckets(&_IStorage.CallOpts, pagination)
}

// ListGroups is a free data retrieval call binding the contract method 0xb8e92c80.
//
// Solidity: function listGroups((bytes,uint64,uint64,bool,bool) pagination, string groupOwner) view returns((address,string,uint8,uint256,string,(string,string)[])[] groupInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageCaller) ListGroups(opts *bind.CallOpts, pagination PageRequest, groupOwner string) (struct {
	GroupInfos   []GroupInfo
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "listGroups", pagination, groupOwner)

	outstruct := new(struct {
		GroupInfos   []GroupInfo
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GroupInfos = *abi.ConvertType(out[0], new([]GroupInfo)).(*[]GroupInfo)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// ListGroups is a free data retrieval call binding the contract method 0xb8e92c80.
//
// Solidity: function listGroups((bytes,uint64,uint64,bool,bool) pagination, string groupOwner) view returns((address,string,uint8,uint256,string,(string,string)[])[] groupInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageSession) ListGroups(pagination PageRequest, groupOwner string) (struct {
	GroupInfos   []GroupInfo
	PageResponse PageResponse
}, error) {
	return _IStorage.Contract.ListGroups(&_IStorage.CallOpts, pagination, groupOwner)
}

// ListGroups is a free data retrieval call binding the contract method 0xb8e92c80.
//
// Solidity: function listGroups((bytes,uint64,uint64,bool,bool) pagination, string groupOwner) view returns((address,string,uint8,uint256,string,(string,string)[])[] groupInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageCallerSession) ListGroups(pagination PageRequest, groupOwner string) (struct {
	GroupInfos   []GroupInfo
	PageResponse PageResponse
}, error) {
	return _IStorage.Contract.ListGroups(&_IStorage.CallOpts, pagination, groupOwner)
}

// ListObjects is a free data retrieval call binding the contract method 0xccd7ddf7.
//
// Solidity: function listObjects((bytes,uint64,uint64,bool,bool) pagination, string bucketName) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64)[] objectInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageCaller) ListObjects(opts *bind.CallOpts, pagination PageRequest, bucketName string) (struct {
	ObjectInfos  []ObjectInfo
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "listObjects", pagination, bucketName)

	outstruct := new(struct {
		ObjectInfos  []ObjectInfo
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ObjectInfos = *abi.ConvertType(out[0], new([]ObjectInfo)).(*[]ObjectInfo)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// ListObjects is a free data retrieval call binding the contract method 0xccd7ddf7.
//
// Solidity: function listObjects((bytes,uint64,uint64,bool,bool) pagination, string bucketName) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64)[] objectInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageSession) ListObjects(pagination PageRequest, bucketName string) (struct {
	ObjectInfos  []ObjectInfo
	PageResponse PageResponse
}, error) {
	return _IStorage.Contract.ListObjects(&_IStorage.CallOpts, pagination, bucketName)
}

// ListObjects is a free data retrieval call binding the contract method 0xccd7ddf7.
//
// Solidity: function listObjects((bytes,uint64,uint64,bool,bool) pagination, string bucketName) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64)[] objectInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageCallerSession) ListObjects(pagination PageRequest, bucketName string) (struct {
	ObjectInfos  []ObjectInfo
	PageResponse PageResponse
}, error) {
	return _IStorage.Contract.ListObjects(&_IStorage.CallOpts, pagination, bucketName)
}

// CreateBucket is a paid mutator transaction binding the contract method 0xf2fb9df8.
//
// Solidity: function createBucket(string bucketName, uint8 visibility, address paymentAddress, address primarySpAddress, (uint64,uint32,bytes) primarySpApproval, uint64 chargedReadQuota) returns(bool success)
func (_IStorage *IStorageTransactor) CreateBucket(opts *bind.TransactOpts, bucketName string, visibility uint8, paymentAddress common.Address, primarySpAddress common.Address, primarySpApproval Approval, chargedReadQuota uint64) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "createBucket", bucketName, visibility, paymentAddress, primarySpAddress, primarySpApproval, chargedReadQuota)
}

// CreateBucket is a paid mutator transaction binding the contract method 0xf2fb9df8.
//
// Solidity: function createBucket(string bucketName, uint8 visibility, address paymentAddress, address primarySpAddress, (uint64,uint32,bytes) primarySpApproval, uint64 chargedReadQuota) returns(bool success)
func (_IStorage *IStorageSession) CreateBucket(bucketName string, visibility uint8, paymentAddress common.Address, primarySpAddress common.Address, primarySpApproval Approval, chargedReadQuota uint64) (*types.Transaction, error) {
	return _IStorage.Contract.CreateBucket(&_IStorage.TransactOpts, bucketName, visibility, paymentAddress, primarySpAddress, primarySpApproval, chargedReadQuota)
}

// CreateBucket is a paid mutator transaction binding the contract method 0xf2fb9df8.
//
// Solidity: function createBucket(string bucketName, uint8 visibility, address paymentAddress, address primarySpAddress, (uint64,uint32,bytes) primarySpApproval, uint64 chargedReadQuota) returns(bool success)
func (_IStorage *IStorageTransactorSession) CreateBucket(bucketName string, visibility uint8, paymentAddress common.Address, primarySpAddress common.Address, primarySpApproval Approval, chargedReadQuota uint64) (*types.Transaction, error) {
	return _IStorage.Contract.CreateBucket(&_IStorage.TransactOpts, bucketName, visibility, paymentAddress, primarySpAddress, primarySpApproval, chargedReadQuota)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x1e4a97dd.
//
// Solidity: function createGroup(string groupName, string extra) returns(bool success)
func (_IStorage *IStorageTransactor) CreateGroup(opts *bind.TransactOpts, groupName string, extra string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "createGroup", groupName, extra)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x1e4a97dd.
//
// Solidity: function createGroup(string groupName, string extra) returns(bool success)
func (_IStorage *IStorageSession) CreateGroup(groupName string, extra string) (*types.Transaction, error) {
	return _IStorage.Contract.CreateGroup(&_IStorage.TransactOpts, groupName, extra)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x1e4a97dd.
//
// Solidity: function createGroup(string groupName, string extra) returns(bool success)
func (_IStorage *IStorageTransactorSession) CreateGroup(groupName string, extra string) (*types.Transaction, error) {
	return _IStorage.Contract.CreateGroup(&_IStorage.TransactOpts, groupName, extra)
}

// CreateObject is a paid mutator transaction binding the contract method 0x6c29c2dc.
//
// Solidity: function createObject(string bucketName, string objectName, uint64 payloadSize, uint8 visibility, string contentType, (uint64,uint32,bytes) primarySpApproval, string[] expectChecksums, uint8 redundancyType) returns(bool success)
func (_IStorage *IStorageTransactor) CreateObject(opts *bind.TransactOpts, bucketName string, objectName string, payloadSize uint64, visibility uint8, contentType string, primarySpApproval Approval, expectChecksums []string, redundancyType uint8) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "createObject", bucketName, objectName, payloadSize, visibility, contentType, primarySpApproval, expectChecksums, redundancyType)
}

// CreateObject is a paid mutator transaction binding the contract method 0x6c29c2dc.
//
// Solidity: function createObject(string bucketName, string objectName, uint64 payloadSize, uint8 visibility, string contentType, (uint64,uint32,bytes) primarySpApproval, string[] expectChecksums, uint8 redundancyType) returns(bool success)
func (_IStorage *IStorageSession) CreateObject(bucketName string, objectName string, payloadSize uint64, visibility uint8, contentType string, primarySpApproval Approval, expectChecksums []string, redundancyType uint8) (*types.Transaction, error) {
	return _IStorage.Contract.CreateObject(&_IStorage.TransactOpts, bucketName, objectName, payloadSize, visibility, contentType, primarySpApproval, expectChecksums, redundancyType)
}

// CreateObject is a paid mutator transaction binding the contract method 0x6c29c2dc.
//
// Solidity: function createObject(string bucketName, string objectName, uint64 payloadSize, uint8 visibility, string contentType, (uint64,uint32,bytes) primarySpApproval, string[] expectChecksums, uint8 redundancyType) returns(bool success)
func (_IStorage *IStorageTransactorSession) CreateObject(bucketName string, objectName string, payloadSize uint64, visibility uint8, contentType string, primarySpApproval Approval, expectChecksums []string, redundancyType uint8) (*types.Transaction, error) {
	return _IStorage.Contract.CreateObject(&_IStorage.TransactOpts, bucketName, objectName, payloadSize, visibility, contentType, primarySpApproval, expectChecksums, redundancyType)
}

// SealObject is a paid mutator transaction binding the contract method 0xcb95c612.
//
// Solidity: function sealObject(address sealAddress, string bucketName, string objectName, uint32 globalVirtualGroupId, string secondarySpBlsAggSignatures) returns(bool success)
func (_IStorage *IStorageTransactor) SealObject(opts *bind.TransactOpts, sealAddress common.Address, bucketName string, objectName string, globalVirtualGroupId uint32, secondarySpBlsAggSignatures string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "sealObject", sealAddress, bucketName, objectName, globalVirtualGroupId, secondarySpBlsAggSignatures)
}

// SealObject is a paid mutator transaction binding the contract method 0xcb95c612.
//
// Solidity: function sealObject(address sealAddress, string bucketName, string objectName, uint32 globalVirtualGroupId, string secondarySpBlsAggSignatures) returns(bool success)
func (_IStorage *IStorageSession) SealObject(sealAddress common.Address, bucketName string, objectName string, globalVirtualGroupId uint32, secondarySpBlsAggSignatures string) (*types.Transaction, error) {
	return _IStorage.Contract.SealObject(&_IStorage.TransactOpts, sealAddress, bucketName, objectName, globalVirtualGroupId, secondarySpBlsAggSignatures)
}

// SealObject is a paid mutator transaction binding the contract method 0xcb95c612.
//
// Solidity: function sealObject(address sealAddress, string bucketName, string objectName, uint32 globalVirtualGroupId, string secondarySpBlsAggSignatures) returns(bool success)
func (_IStorage *IStorageTransactorSession) SealObject(sealAddress common.Address, bucketName string, objectName string, globalVirtualGroupId uint32, secondarySpBlsAggSignatures string) (*types.Transaction, error) {
	return _IStorage.Contract.SealObject(&_IStorage.TransactOpts, sealAddress, bucketName, objectName, globalVirtualGroupId, secondarySpBlsAggSignatures)
}

// SealObjectV2 is a paid mutator transaction binding the contract method 0x9b54c033.
//
// Solidity: function sealObjectV2(address sealAddress, string bucketName, string objectName, uint32 globalVirtualGroupId, string secondarySpBlsAggSignatures, string[] expectChecksums) returns(bool success)
func (_IStorage *IStorageTransactor) SealObjectV2(opts *bind.TransactOpts, sealAddress common.Address, bucketName string, objectName string, globalVirtualGroupId uint32, secondarySpBlsAggSignatures string, expectChecksums []string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "sealObjectV2", sealAddress, bucketName, objectName, globalVirtualGroupId, secondarySpBlsAggSignatures, expectChecksums)
}

// SealObjectV2 is a paid mutator transaction binding the contract method 0x9b54c033.
//
// Solidity: function sealObjectV2(address sealAddress, string bucketName, string objectName, uint32 globalVirtualGroupId, string secondarySpBlsAggSignatures, string[] expectChecksums) returns(bool success)
func (_IStorage *IStorageSession) SealObjectV2(sealAddress common.Address, bucketName string, objectName string, globalVirtualGroupId uint32, secondarySpBlsAggSignatures string, expectChecksums []string) (*types.Transaction, error) {
	return _IStorage.Contract.SealObjectV2(&_IStorage.TransactOpts, sealAddress, bucketName, objectName, globalVirtualGroupId, secondarySpBlsAggSignatures, expectChecksums)
}

// SealObjectV2 is a paid mutator transaction binding the contract method 0x9b54c033.
//
// Solidity: function sealObjectV2(address sealAddress, string bucketName, string objectName, uint32 globalVirtualGroupId, string secondarySpBlsAggSignatures, string[] expectChecksums) returns(bool success)
func (_IStorage *IStorageTransactorSession) SealObjectV2(sealAddress common.Address, bucketName string, objectName string, globalVirtualGroupId uint32, secondarySpBlsAggSignatures string, expectChecksums []string) (*types.Transaction, error) {
	return _IStorage.Contract.SealObjectV2(&_IStorage.TransactOpts, sealAddress, bucketName, objectName, globalVirtualGroupId, secondarySpBlsAggSignatures, expectChecksums)
}

// IStorageCreateBucketIterator is returned from FilterCreateBucket and is used to iterate over the raw logs and unpacked data for CreateBucket events raised by the IStorage contract.
type IStorageCreateBucketIterator struct {
	Event *IStorageCreateBucket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IStorageCreateBucketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageCreateBucket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IStorageCreateBucket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IStorageCreateBucketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageCreateBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageCreateBucket represents a CreateBucket event raised by the IStorage contract.
type IStorageCreateBucket struct {
	Creator          common.Address
	PaymentAddress   common.Address
	PrimarySpAddress common.Address
	Id               *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCreateBucket is a free log retrieval operation binding the contract event 0x245039a22720b027adaa35bc8837e8b9dd42e795bbdc7c7539e25ad2043c3723.
//
// Solidity: event CreateBucket(address indexed creator, address indexed paymentAddress, address indexed primarySpAddress, uint256 id)
func (_IStorage *IStorageFilterer) FilterCreateBucket(opts *bind.FilterOpts, creator []common.Address, paymentAddress []common.Address, primarySpAddress []common.Address) (*IStorageCreateBucketIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var paymentAddressRule []interface{}
	for _, paymentAddressItem := range paymentAddress {
		paymentAddressRule = append(paymentAddressRule, paymentAddressItem)
	}
	var primarySpAddressRule []interface{}
	for _, primarySpAddressItem := range primarySpAddress {
		primarySpAddressRule = append(primarySpAddressRule, primarySpAddressItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "CreateBucket", creatorRule, paymentAddressRule, primarySpAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStorageCreateBucketIterator{contract: _IStorage.contract, event: "CreateBucket", logs: logs, sub: sub}, nil
}

// WatchCreateBucket is a free log subscription operation binding the contract event 0x245039a22720b027adaa35bc8837e8b9dd42e795bbdc7c7539e25ad2043c3723.
//
// Solidity: event CreateBucket(address indexed creator, address indexed paymentAddress, address indexed primarySpAddress, uint256 id)
func (_IStorage *IStorageFilterer) WatchCreateBucket(opts *bind.WatchOpts, sink chan<- *IStorageCreateBucket, creator []common.Address, paymentAddress []common.Address, primarySpAddress []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var paymentAddressRule []interface{}
	for _, paymentAddressItem := range paymentAddress {
		paymentAddressRule = append(paymentAddressRule, paymentAddressItem)
	}
	var primarySpAddressRule []interface{}
	for _, primarySpAddressItem := range primarySpAddress {
		primarySpAddressRule = append(primarySpAddressRule, primarySpAddressItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "CreateBucket", creatorRule, paymentAddressRule, primarySpAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageCreateBucket)
				if err := _IStorage.contract.UnpackLog(event, "CreateBucket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateBucket is a log parse operation binding the contract event 0x245039a22720b027adaa35bc8837e8b9dd42e795bbdc7c7539e25ad2043c3723.
//
// Solidity: event CreateBucket(address indexed creator, address indexed paymentAddress, address indexed primarySpAddress, uint256 id)
func (_IStorage *IStorageFilterer) ParseCreateBucket(log types.Log) (*IStorageCreateBucket, error) {
	event := new(IStorageCreateBucket)
	if err := _IStorage.contract.UnpackLog(event, "CreateBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageCreateGroupIterator is returned from FilterCreateGroup and is used to iterate over the raw logs and unpacked data for CreateGroup events raised by the IStorage contract.
type IStorageCreateGroupIterator struct {
	Event *IStorageCreateGroup // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IStorageCreateGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageCreateGroup)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IStorageCreateGroup)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IStorageCreateGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageCreateGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageCreateGroup represents a CreateGroup event raised by the IStorage contract.
type IStorageCreateGroup struct {
	Creator common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateGroup is a free log retrieval operation binding the contract event 0x88d8a40d3d79893e13972978642d5fe29930912ee4c0b62a535815945c1d7bd0.
//
// Solidity: event CreateGroup(address indexed creator, uint256 id)
func (_IStorage *IStorageFilterer) FilterCreateGroup(opts *bind.FilterOpts, creator []common.Address) (*IStorageCreateGroupIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "CreateGroup", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageCreateGroupIterator{contract: _IStorage.contract, event: "CreateGroup", logs: logs, sub: sub}, nil
}

// WatchCreateGroup is a free log subscription operation binding the contract event 0x88d8a40d3d79893e13972978642d5fe29930912ee4c0b62a535815945c1d7bd0.
//
// Solidity: event CreateGroup(address indexed creator, uint256 id)
func (_IStorage *IStorageFilterer) WatchCreateGroup(opts *bind.WatchOpts, sink chan<- *IStorageCreateGroup, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "CreateGroup", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageCreateGroup)
				if err := _IStorage.contract.UnpackLog(event, "CreateGroup", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateGroup is a log parse operation binding the contract event 0x88d8a40d3d79893e13972978642d5fe29930912ee4c0b62a535815945c1d7bd0.
//
// Solidity: event CreateGroup(address indexed creator, uint256 id)
func (_IStorage *IStorageFilterer) ParseCreateGroup(log types.Log) (*IStorageCreateGroup, error) {
	event := new(IStorageCreateGroup)
	if err := _IStorage.contract.UnpackLog(event, "CreateGroup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageCreateObjectIterator is returned from FilterCreateObject and is used to iterate over the raw logs and unpacked data for CreateObject events raised by the IStorage contract.
type IStorageCreateObjectIterator struct {
	Event *IStorageCreateObject // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IStorageCreateObjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageCreateObject)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IStorageCreateObject)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IStorageCreateObjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageCreateObjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageCreateObject represents a CreateObject event raised by the IStorage contract.
type IStorageCreateObject struct {
	Creator common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateObject is a free log retrieval operation binding the contract event 0x036486acfc6433f762a6a6a8c2a77904caf492f679f724f1410394f7d5bc2a1d.
//
// Solidity: event CreateObject(address indexed creator, uint256 id)
func (_IStorage *IStorageFilterer) FilterCreateObject(opts *bind.FilterOpts, creator []common.Address) (*IStorageCreateObjectIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "CreateObject", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageCreateObjectIterator{contract: _IStorage.contract, event: "CreateObject", logs: logs, sub: sub}, nil
}

// WatchCreateObject is a free log subscription operation binding the contract event 0x036486acfc6433f762a6a6a8c2a77904caf492f679f724f1410394f7d5bc2a1d.
//
// Solidity: event CreateObject(address indexed creator, uint256 id)
func (_IStorage *IStorageFilterer) WatchCreateObject(opts *bind.WatchOpts, sink chan<- *IStorageCreateObject, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "CreateObject", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageCreateObject)
				if err := _IStorage.contract.UnpackLog(event, "CreateObject", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateObject is a log parse operation binding the contract event 0x036486acfc6433f762a6a6a8c2a77904caf492f679f724f1410394f7d5bc2a1d.
//
// Solidity: event CreateObject(address indexed creator, uint256 id)
func (_IStorage *IStorageFilterer) ParseCreateObject(log types.Log) (*IStorageCreateObject, error) {
	event := new(IStorageCreateObject)
	if err := _IStorage.contract.UnpackLog(event, "CreateObject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageSealObjectIterator is returned from FilterSealObject and is used to iterate over the raw logs and unpacked data for SealObject events raised by the IStorage contract.
type IStorageSealObjectIterator struct {
	Event *IStorageSealObject // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IStorageSealObjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageSealObject)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IStorageSealObject)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IStorageSealObjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageSealObjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageSealObject represents a SealObject event raised by the IStorage contract.
type IStorageSealObject struct {
	Creator     common.Address
	SealAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSealObject is a free log retrieval operation binding the contract event 0xe974c93deae1d2628e44a9e4d6a09748c1ebe32db999037b4ba05b62e2e331b6.
//
// Solidity: event SealObject(address indexed creator, address indexed sealAddress)
func (_IStorage *IStorageFilterer) FilterSealObject(opts *bind.FilterOpts, creator []common.Address, sealAddress []common.Address) (*IStorageSealObjectIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var sealAddressRule []interface{}
	for _, sealAddressItem := range sealAddress {
		sealAddressRule = append(sealAddressRule, sealAddressItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "SealObject", creatorRule, sealAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStorageSealObjectIterator{contract: _IStorage.contract, event: "SealObject", logs: logs, sub: sub}, nil
}

// WatchSealObject is a free log subscription operation binding the contract event 0xe974c93deae1d2628e44a9e4d6a09748c1ebe32db999037b4ba05b62e2e331b6.
//
// Solidity: event SealObject(address indexed creator, address indexed sealAddress)
func (_IStorage *IStorageFilterer) WatchSealObject(opts *bind.WatchOpts, sink chan<- *IStorageSealObject, creator []common.Address, sealAddress []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var sealAddressRule []interface{}
	for _, sealAddressItem := range sealAddress {
		sealAddressRule = append(sealAddressRule, sealAddressItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "SealObject", creatorRule, sealAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageSealObject)
				if err := _IStorage.contract.UnpackLog(event, "SealObject", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSealObject is a log parse operation binding the contract event 0xe974c93deae1d2628e44a9e4d6a09748c1ebe32db999037b4ba05b62e2e331b6.
//
// Solidity: event SealObject(address indexed creator, address indexed sealAddress)
func (_IStorage *IStorageFilterer) ParseSealObject(log types.Log) (*IStorageSealObject, error) {
	event := new(IStorageSealObject)
	if err := _IStorage.contract.UnpackLog(event, "SealObject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageSealObjectV2Iterator is returned from FilterSealObjectV2 and is used to iterate over the raw logs and unpacked data for SealObjectV2 events raised by the IStorage contract.
type IStorageSealObjectV2Iterator struct {
	Event *IStorageSealObjectV2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IStorageSealObjectV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageSealObjectV2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IStorageSealObjectV2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IStorageSealObjectV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageSealObjectV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageSealObjectV2 represents a SealObjectV2 event raised by the IStorage contract.
type IStorageSealObjectV2 struct {
	Creator     common.Address
	SealAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSealObjectV2 is a free log retrieval operation binding the contract event 0x088ef0ff0fdec74a602b6b9b1e399c7125cd4adb2ecbebcd34b16fbbe7961cde.
//
// Solidity: event SealObjectV2(address indexed creator, address indexed sealAddress)
func (_IStorage *IStorageFilterer) FilterSealObjectV2(opts *bind.FilterOpts, creator []common.Address, sealAddress []common.Address) (*IStorageSealObjectV2Iterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var sealAddressRule []interface{}
	for _, sealAddressItem := range sealAddress {
		sealAddressRule = append(sealAddressRule, sealAddressItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "SealObjectV2", creatorRule, sealAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStorageSealObjectV2Iterator{contract: _IStorage.contract, event: "SealObjectV2", logs: logs, sub: sub}, nil
}

// WatchSealObjectV2 is a free log subscription operation binding the contract event 0x088ef0ff0fdec74a602b6b9b1e399c7125cd4adb2ecbebcd34b16fbbe7961cde.
//
// Solidity: event SealObjectV2(address indexed creator, address indexed sealAddress)
func (_IStorage *IStorageFilterer) WatchSealObjectV2(opts *bind.WatchOpts, sink chan<- *IStorageSealObjectV2, creator []common.Address, sealAddress []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var sealAddressRule []interface{}
	for _, sealAddressItem := range sealAddress {
		sealAddressRule = append(sealAddressRule, sealAddressItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "SealObjectV2", creatorRule, sealAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageSealObjectV2)
				if err := _IStorage.contract.UnpackLog(event, "SealObjectV2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSealObjectV2 is a log parse operation binding the contract event 0x088ef0ff0fdec74a602b6b9b1e399c7125cd4adb2ecbebcd34b16fbbe7961cde.
//
// Solidity: event SealObjectV2(address indexed creator, address indexed sealAddress)
func (_IStorage *IStorageFilterer) ParseSealObjectV2(log types.Log) (*IStorageSealObjectV2, error) {
	event := new(IStorageSealObjectV2)
	if err := _IStorage.contract.UnpackLog(event, "SealObjectV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
