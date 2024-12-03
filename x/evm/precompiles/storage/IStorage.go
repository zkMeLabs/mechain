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

// BucketMetaData is an auto generated low-level Go binding around an user-defined struct.
type BucketMetaData struct {
	Description string
	ExternalUrl string
	BucketName  string
	Image       string
	Attributes  []Trait
}

// GVGMapping is an auto generated low-level Go binding around an user-defined struct.
type GVGMapping struct {
	SrcGlobalVirtualGroupId uint32
	DstGlobalVirtualGroupId uint32
	SecondarySpBlsSignature []byte
}

// GlobalVirtualGroup is an auto generated low-level Go binding around an user-defined struct.
type GlobalVirtualGroup struct {
	Id                    uint32
	FamilyId              uint32
	PrimarySpId           uint32
	SecondarySpIds        []uint32
	StoredSize            uint64
	VirtualPaymentAddress common.Address
	TotalDeposit          string
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

// GroupMember is an auto generated low-level Go binding around an user-defined struct.
type GroupMember struct {
	Id             *big.Int
	GroupId        *big.Int
	Member         common.Address
	ExpirationTime int64
}

// GroupMetaData is an auto generated low-level Go binding around an user-defined struct.
type GroupMetaData struct {
	Description string
	ExternalUrl string
	GroupName   string
	Image       string
	Attributes  []Trait
}

// InternalBucketInfo is an auto generated low-level Go binding around an user-defined struct.
type InternalBucketInfo struct {
	PriceTime               int64
	TotalChargeSize         uint64
	LocalVirtualGroups      []LocalVirtualGroup
	NextLocalVirtualGroupId uint32
}

// IsPriceChanged is an auto generated low-level Go binding around an user-defined struct.
type IsPriceChanged struct {
	Changed                    bool
	CurrentReadPrice           *big.Int
	CurrentPrimaryStorePrice   *big.Int
	CurrentSecondaryStorePrice *big.Int
	CurrentValidatorTaxRate    *big.Int
	NewReadPrice               *big.Int
	NewPrimaryStorePrice       *big.Int
	NewSecondaryStorePrice     *big.Int
	NewValidatorTaxRate        *big.Int
}

// LocalVirtualGroup is an auto generated low-level Go binding around an user-defined struct.
type LocalVirtualGroup struct {
	Id                   uint32
	GlobalVirtualGroupId uint32
	StoredSize           uint64
	TotalChargeSize      uint64
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

// ObjectMetaData is an auto generated low-level Go binding around an user-defined struct.
type ObjectMetaData struct {
	Description string
	ExternalUrl string
	ObjectName  string
	Image       string
	Attributes  []Trait
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

// Params is an auto generated low-level Go binding around an user-defined struct.
type Params struct {
	VersionedParams                  VersionedParams
	MaxPayloadSize                   uint64
	BscMirrorBucketRelayerFee        string
	BscMirrorBucketAckRelayerFee     string
	BscMirrorObjectRelayerFee        string
	BscMirrorObjectAckRelayerFee     string
	BscMirrorGroupRelayerFee         string
	BscMirrorGroupAckRelayerFee      string
	MaxBucketsPerAccount             uint32
	DiscontinueCountingWindow        uint64
	DiscontinueObjectMax             uint64
	DiscontinueBucketMax             uint64
	DiscontinueConfirmPeriod         int64
	DiscontinueDeletionMax           uint64
	StalePolicyCleanupMax            uint64
	MinQuotaUpdateInterval           uint64
	MaxLocalVirtualGroupNumPerBucket uint32
	OpMirrorBucketRelayerFee         string
	OpMirrorBucketAckRelayerFee      string
	OpMirrorObjectRelayerFee         string
	OpMirrorObjectAckRelayerFee      string
	OpMirrorGroupRelayerFee          string
	OpMirrorGroupAckRelayerFee       string
	PolygonMirrorBucketRelayerFee    string
	PolygonMirrorBucketAckRelayerFee string
}

// Policy is an auto generated low-level Go binding around an user-defined struct.
type Policy struct {
	Id             *big.Int
	Principal      Principal
	ResourceType   int32
	ResourceId     *big.Int
	Statements     []Statement
	ExpirationTime int64
}

// Principal is an auto generated low-level Go binding around an user-defined struct.
type Principal struct {
	PrincipalType int32
	Value         string
}

// ShadowObjectInfo is an auto generated low-level Go binding around an user-defined struct.
type ShadowObjectInfo struct {
	Operator    string
	Id          *big.Int
	ContentType string
	PayloadSize uint64
	Checksums   []string
	UpdatedAt   int64
	Version     int64
}

// Statement is an auto generated low-level Go binding around an user-defined struct.
type Statement struct {
	Effect         int32
	Actions        []int32
	Resources      []string
	ExpirationTime int64
	LimitSize      uint64
}

// Tag is an auto generated low-level Go binding around an user-defined struct.
type Tag struct {
	Key   string
	Value string
}

// Trait is an auto generated low-level Go binding around an user-defined struct.
type Trait struct {
	TraitType string
	Value     string
}

// VersionedParams is an auto generated low-level Go binding around an user-defined struct.
type VersionedParams struct {
	MaxSegmentSize          uint64
	RedundantDataChunkNum   uint32
	RedundantParityChunkNum uint32
	MinChargeSize           uint64
}

// IStorageMetaData contains all meta data concerning the IStorage contract.
var IStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CancelCreateObject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"CancelMigrateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"CompleteMigrateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CopyObject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"primarySpAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CreateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CreateGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CreateObject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"}],\"name\":\"DelegateCreateObject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"}],\"name\":\"DelegateUpdateObjectContent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"DeleteBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"DeleteGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"DeleteObject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"DeletePolicy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"DiscontinueBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"DiscontinueObject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"}],\"name\":\"LeaveGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"MigrateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"MirrorBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"MirrorGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"MirrorObject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"PutPolicy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"RejectMigrateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"}],\"name\":\"RejectSealObject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"RenewGroupMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sealAddress\",\"type\":\"address\"}],\"name\":\"SealObject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sealAddress\",\"type\":\"address\"}],\"name\":\"SealObjectV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"SetBucketFlowRateLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"SetTag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ToggleSPAsDelegatedAgent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketName\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"visibility\",\"type\":\"uint8\"}],\"name\":\"UpdateBucketInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"UpdateGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"UpdateGroupExtra\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"}],\"name\":\"UpdateObjectContent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"UpdateObjectInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"UpdateParams\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"}],\"name\":\"cancelCreateObject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"cancelMigrateBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"gvgFamilyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcGlobalVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dstGlobalVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"secondarySpBlsSignature\",\"type\":\"bytes\"}],\"internalType\":\"structGVGMapping[]\",\"name\":\"gvgMappings\",\"type\":\"tuple[]\"}],\"name\":\"completeMigrateBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"srcBucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dstBucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"srcObjectName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dstObjectName\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"expiredHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structApproval\",\"name\":\"dstPrimarySpApproval\",\"type\":\"tuple\"}],\"name\":\"copyObject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primarySpAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"expiredHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structApproval\",\"name\":\"primarySpApproval\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"chargedReadQuota\",\"type\":\"uint64\"}],\"name\":\"createBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"}],\"name\":\"createGroup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"expiredHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structApproval\",\"name\":\"primarySpApproval\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"expectChecksums\",\"type\":\"string[]\"},{\"internalType\":\"enumRedundancyType\",\"name\":\"redundancyType\",\"type\":\"uint8\"}],\"name\":\"createObject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"creator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"expectChecksums\",\"type\":\"string[]\"},{\"internalType\":\"enumRedundancyType\",\"name\":\"redundancyType\",\"type\":\"uint8\"}],\"name\":\"delegateCreateObject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"updater\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"expectChecksums\",\"type\":\"string[]\"}],\"name\":\"delegateUpdateObjectContent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"deleteBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"}],\"name\":\"deleteGroup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"}],\"name\":\"deleteObject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int32\",\"name\":\"principalType\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structPrincipal\",\"name\":\"principal\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"}],\"name\":\"deletePolicy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"discontinueBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"objectIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"discontinueObject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"headBucket\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"int64\",\"name\":\"createAt\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chargedReadQuota\",\"type\":\"uint64\"},{\"internalType\":\"enumBucketStatus\",\"name\":\"bucketStatus\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"spAsDelegatedAgentDisabled\",\"type\":\"bool\"}],\"internalType\":\"structBucketInfo\",\"name\":\"bucketInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isRateLimited\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"flowRateLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentFlowRate\",\"type\":\"uint256\"}],\"internalType\":\"structBucketExtraInfo\",\"name\":\"bucketExtraInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketId\",\"type\":\"string\"}],\"name\":\"headBucketById\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"int64\",\"name\":\"createAt\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chargedReadQuota\",\"type\":\"uint64\"},{\"internalType\":\"enumBucketStatus\",\"name\":\"bucketStatus\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"spAsDelegatedAgentDisabled\",\"type\":\"bool\"}],\"internalType\":\"structBucketInfo\",\"name\":\"bucketInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isRateLimited\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"flowRateLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentFlowRate\",\"type\":\"uint256\"}],\"internalType\":\"structBucketExtraInfo\",\"name\":\"bucketExtraInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"headBucketExtra\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"priceTime\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"totalChargeSize\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"storedSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalChargeSize\",\"type\":\"uint64\"}],\"internalType\":\"structLocalVirtualGroup[]\",\"name\":\"localVirtualGroups\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"nextLocalVirtualGroupId\",\"type\":\"uint32\"}],\"internalType\":\"structInternalBucketInfo\",\"name\":\"extraInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenId\",\"type\":\"string\"}],\"name\":\"headBucketNFT\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"externalUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"traitType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTrait[]\",\"name\":\"attributes\",\"type\":\"tuple[]\"}],\"internalType\":\"structBucketMetaData\",\"name\":\"bucketMetaData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"groupOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"}],\"name\":\"headGroup\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"}],\"internalType\":\"structGroupInfo\",\"name\":\"groupInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"groupOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"}],\"name\":\"headGroupMember\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"expirationTime\",\"type\":\"int64\"}],\"internalType\":\"structGroupMember\",\"name\":\"groupMember\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenId\",\"type\":\"string\"}],\"name\":\"headGroupNFT\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"externalUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"traitType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTrait[]\",\"name\":\"attributes\",\"type\":\"tuple[]\"}],\"internalType\":\"structGroupMetaData\",\"name\":\"groupMetaData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"}],\"name\":\"headObject\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"localVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"createAt\",\"type\":\"int64\"},{\"internalType\":\"enumObjectStatus\",\"name\":\"objectStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumRedundancyType\",\"name\":\"redundancyType\",\"type\":\"uint8\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"checksums\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"isUpdating\",\"type\":\"bool\"},{\"internalType\":\"int64\",\"name\":\"updatedAt\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"updatedBy\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"version\",\"type\":\"int64\"}],\"internalType\":\"structObjectInfo\",\"name\":\"objectInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"familyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"primarySpId\",\"type\":\"uint32\"},{\"internalType\":\"uint32[]\",\"name\":\"secondarySpIds\",\"type\":\"uint32[]\"},{\"internalType\":\"uint64\",\"name\":\"storedSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"virtualPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"totalDeposit\",\"type\":\"string\"}],\"internalType\":\"structGlobalVirtualGroup\",\"name\":\"globalVirtualGroup\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"objectId\",\"type\":\"string\"}],\"name\":\"headObjectById\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"localVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"createAt\",\"type\":\"int64\"},{\"internalType\":\"enumObjectStatus\",\"name\":\"objectStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumRedundancyType\",\"name\":\"redundancyType\",\"type\":\"uint8\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"checksums\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"isUpdating\",\"type\":\"bool\"},{\"internalType\":\"int64\",\"name\":\"updatedAt\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"updatedBy\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"version\",\"type\":\"int64\"}],\"internalType\":\"structObjectInfo\",\"name\":\"objectInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"familyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"primarySpId\",\"type\":\"uint32\"},{\"internalType\":\"uint32[]\",\"name\":\"secondarySpIds\",\"type\":\"uint32[]\"},{\"internalType\":\"uint64\",\"name\":\"storedSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"virtualPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"totalDeposit\",\"type\":\"string\"}],\"internalType\":\"structGlobalVirtualGroup\",\"name\":\"globalVirtualGroup\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenId\",\"type\":\"string\"}],\"name\":\"headObjectNFT\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"externalUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"traitType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTrait[]\",\"name\":\"attributes\",\"type\":\"tuple[]\"}],\"internalType\":\"structObjectMetaData\",\"name\":\"objectMetaData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"}],\"name\":\"headShadowObject\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"},{\"internalType\":\"string[]\",\"name\":\"checksums\",\"type\":\"string[]\"},{\"internalType\":\"int64\",\"name\":\"updatedAt\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"version\",\"type\":\"int64\"}],\"internalType\":\"structShadowObjectInfo\",\"name\":\"objectInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"groupOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"}],\"name\":\"leaveGroup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"listBuckets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"int64\",\"name\":\"createAt\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chargedReadQuota\",\"type\":\"uint64\"},{\"internalType\":\"enumBucketStatus\",\"name\":\"bucketStatus\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"spAsDelegatedAgentDisabled\",\"type\":\"bool\"}],\"internalType\":\"structBucketInfo[]\",\"name\":\"bucketInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"groupOwner\",\"type\":\"address\"}],\"name\":\"listGroups\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"}],\"internalType\":\"structGroupInfo[]\",\"name\":\"groupInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"listObjects\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"localVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"createAt\",\"type\":\"int64\"},{\"internalType\":\"enumObjectStatus\",\"name\":\"objectStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumRedundancyType\",\"name\":\"redundancyType\",\"type\":\"uint8\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"checksums\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"isUpdating\",\"type\":\"bool\"},{\"internalType\":\"int64\",\"name\":\"updatedAt\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"updatedBy\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"version\",\"type\":\"int64\"}],\"internalType\":\"structObjectInfo[]\",\"name\":\"objectInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"bucketId\",\"type\":\"string\"}],\"name\":\"listObjectsByBucketId\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"localVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"createAt\",\"type\":\"int64\"},{\"internalType\":\"enumObjectStatus\",\"name\":\"objectStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumRedundancyType\",\"name\":\"redundancyType\",\"type\":\"uint8\"},{\"internalType\":\"enumSourceType\",\"name\":\"sourceType\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"checksums\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"isUpdating\",\"type\":\"bool\"},{\"internalType\":\"int64\",\"name\":\"updatedAt\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"updatedBy\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"version\",\"type\":\"int64\"}],\"internalType\":\"structObjectInfo[]\",\"name\":\"objectInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"dstPrimarySpId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"expiredHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structApproval\",\"name\":\"dstPrimarySpApproval\",\"type\":\"tuple\"}],\"name\":\"migrateBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bucketId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"}],\"name\":\"mirrorBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"}],\"name\":\"mirrorGroup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"objectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"}],\"name\":\"mirrorObject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"maxSegmentSize\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"redundantDataChunkNum\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"redundantParityChunkNum\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"minChargeSize\",\"type\":\"uint64\"}],\"internalType\":\"structVersionedParams\",\"name\":\"versionedParams\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"maxPayloadSize\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"bscMirrorBucketRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorBucketAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorObjectRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorObjectAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorGroupRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorGroupAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"maxBucketsPerAccount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"discontinueCountingWindow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"discontinueObjectMax\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"discontinueBucketMax\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"discontinueConfirmPeriod\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"discontinueDeletionMax\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stalePolicyCleanupMax\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minQuotaUpdateInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxLocalVirtualGroupNumPerBucket\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"opMirrorBucketRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorBucketAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorObjectRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorObjectAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorGroupRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorGroupAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"polygonMirrorBucketRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"polygonMirrorBucketAckRelayerFee\",\"type\":\"string\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int32\",\"name\":\"principalType\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structPrincipal\",\"name\":\"principal\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"int32\",\"name\":\"effect\",\"type\":\"int32\"},{\"internalType\":\"int32[]\",\"name\":\"actions\",\"type\":\"int32[]\"},{\"internalType\":\"string[]\",\"name\":\"resources\",\"type\":\"string[]\"},{\"internalType\":\"int64\",\"name\":\"expirationTime\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"limitSize\",\"type\":\"uint64\"}],\"internalType\":\"structStatement[]\",\"name\":\"statements\",\"type\":\"tuple[]\"},{\"internalType\":\"int64\",\"name\":\"expirationTime\",\"type\":\"int64\"}],\"name\":\"putPolicy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"members\",\"type\":\"string[]\"}],\"name\":\"queryGroupMembersExist\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"checkMembers\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"exists\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupOwner\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"groupNames\",\"type\":\"string[]\"}],\"name\":\"queryGroupsExist\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"checkGroupNames\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"exists\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"groupIds\",\"type\":\"string[]\"}],\"name\":\"queryGroupsExistById\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"checkGroupIds\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"exists\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"queryIsPriceChanged\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"changed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"currentReadPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentPrimaryStorePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentSecondaryStorePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentValidatorTaxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newReadPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newPrimaryStorePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newSecondaryStorePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newValidatorTaxRate\",\"type\":\"uint256\"}],\"internalType\":\"structIsPriceChanged\",\"name\":\"isPriceChanged\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"primarySpAddress\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"createAt\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"}],\"name\":\"queryLockFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"timestamp\",\"type\":\"int64\"}],\"name\":\"queryParamsByTimestamp\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"maxSegmentSize\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"redundantDataChunkNum\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"redundantParityChunkNum\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"minChargeSize\",\"type\":\"uint64\"}],\"internalType\":\"structVersionedParams\",\"name\":\"versionedParams\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"maxPayloadSize\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"bscMirrorBucketRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorBucketAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorObjectRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorObjectAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorGroupRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorGroupAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"maxBucketsPerAccount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"discontinueCountingWindow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"discontinueObjectMax\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"discontinueBucketMax\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"discontinueConfirmPeriod\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"discontinueDeletionMax\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stalePolicyCleanupMax\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minQuotaUpdateInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxLocalVirtualGroupNumPerBucket\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"opMirrorBucketRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorBucketAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorObjectRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorObjectAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorGroupRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorGroupAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"polygonMirrorBucketRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"polygonMirrorBucketAckRelayerFee\",\"type\":\"string\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"paymentAccount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bucketOwner\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"queryPaymentAccountBucketFlowRateLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"flowRateLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"policyId\",\"type\":\"string\"}],\"name\":\"queryPolicyById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int32\",\"name\":\"principalType\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structPrincipal\",\"name\":\"principal\",\"type\":\"tuple\"},{\"internalType\":\"int32\",\"name\":\"resourceType\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"resourceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int32\",\"name\":\"effect\",\"type\":\"int32\"},{\"internalType\":\"int32[]\",\"name\":\"actions\",\"type\":\"int32[]\"},{\"internalType\":\"string[]\",\"name\":\"resources\",\"type\":\"string[]\"},{\"internalType\":\"int64\",\"name\":\"expirationTime\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"limitSize\",\"type\":\"uint64\"}],\"internalType\":\"structStatement[]\",\"name\":\"statements\",\"type\":\"tuple[]\"},{\"internalType\":\"int64\",\"name\":\"expirationTime\",\"type\":\"int64\"}],\"internalType\":\"structPolicy\",\"name\":\"policy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"principalAddr\",\"type\":\"string\"}],\"name\":\"queryPolicyForAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int32\",\"name\":\"principalType\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structPrincipal\",\"name\":\"principal\",\"type\":\"tuple\"},{\"internalType\":\"int32\",\"name\":\"resourceType\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"resourceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int32\",\"name\":\"effect\",\"type\":\"int32\"},{\"internalType\":\"int32[]\",\"name\":\"actions\",\"type\":\"int32[]\"},{\"internalType\":\"string[]\",\"name\":\"resources\",\"type\":\"string[]\"},{\"internalType\":\"int64\",\"name\":\"expirationTime\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"limitSize\",\"type\":\"uint64\"}],\"internalType\":\"structStatement[]\",\"name\":\"statements\",\"type\":\"tuple[]\"},{\"internalType\":\"int64\",\"name\":\"expirationTime\",\"type\":\"int64\"}],\"internalType\":\"structPolicy\",\"name\":\"policy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"queryPolicyForGroup\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int32\",\"name\":\"principalType\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structPrincipal\",\"name\":\"principal\",\"type\":\"tuple\"},{\"internalType\":\"int32\",\"name\":\"resourceType\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"resourceId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int32\",\"name\":\"effect\",\"type\":\"int32\"},{\"internalType\":\"int32[]\",\"name\":\"actions\",\"type\":\"int32[]\"},{\"internalType\":\"string[]\",\"name\":\"resources\",\"type\":\"string[]\"},{\"internalType\":\"int64\",\"name\":\"expirationTime\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"limitSize\",\"type\":\"uint64\"}],\"internalType\":\"structStatement[]\",\"name\":\"statements\",\"type\":\"tuple[]\"},{\"internalType\":\"int64\",\"name\":\"expirationTime\",\"type\":\"int64\"}],\"internalType\":\"structPolicy\",\"name\":\"policy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"queryQuotaUpdateTime\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"updateAt\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"rejectMigrateBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"}],\"name\":\"rejectSealObject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"groupOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"members\",\"type\":\"address[]\"},{\"internalType\":\"int64[]\",\"name\":\"expirationTime\",\"type\":\"int64[]\"}],\"name\":\"renewGroupMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sealAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"secondarySpBlsAggSignatures\",\"type\":\"string\"}],\"name\":\"sealObject\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sealAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupId\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"secondarySpBlsAggSignatures\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"expectChecksums\",\"type\":\"string[]\"}],\"name\":\"sealObjectV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bucketOwner\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"paymentAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"flowRateLimit\",\"type\":\"uint256\"}],\"name\":\"setBucketFlowRateLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structTag[]\",\"name\":\"tags\",\"type\":\"tuple[]\"}],\"name\":\"setTag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"toggleSPAsDelegatedAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"chargedReadQuota\",\"type\":\"int128\"}],\"name\":\"updateBucketInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"groupOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"membersToAdd\",\"type\":\"address[]\"},{\"internalType\":\"int64[]\",\"name\":\"expirationTime\",\"type\":\"int64[]\"},{\"internalType\":\"address[]\",\"name\":\"membersToDelete\",\"type\":\"address[]\"}],\"name\":\"updateGroup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"groupOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"}],\"name\":\"updateGroupExtra\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"payloadSize\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"expectChecksums\",\"type\":\"string[]\"}],\"name\":\"updateObjectContent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"enumVisibilityType\",\"name\":\"visibility\",\"type\":\"uint8\"}],\"name\":\"updateObjectInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"authority\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"maxSegmentSize\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"redundantDataChunkNum\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"redundantParityChunkNum\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"minChargeSize\",\"type\":\"uint64\"}],\"internalType\":\"structVersionedParams\",\"name\":\"versionedParams\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"maxPayloadSize\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"bscMirrorBucketRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorBucketAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorObjectRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorObjectAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorGroupRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bscMirrorGroupAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"maxBucketsPerAccount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"discontinueCountingWindow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"discontinueObjectMax\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"discontinueBucketMax\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"discontinueConfirmPeriod\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"discontinueDeletionMax\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stalePolicyCleanupMax\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minQuotaUpdateInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxLocalVirtualGroupNumPerBucket\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"opMirrorBucketRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorBucketAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorObjectRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorObjectAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorGroupRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"opMirrorGroupAckRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"polygonMirrorBucketRelayerFee\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"polygonMirrorBucketAckRelayerFee\",\"type\":\"string\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"updateParams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"objectName\",\"type\":\"string\"},{\"internalType\":\"int32\",\"name\":\"actionType\",\"type\":\"int32\"}],\"name\":\"verifyPermission\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"effect\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// HeadBucketById is a free data retrieval call binding the contract method 0xb5d73569.
//
// Solidity: function headBucketById(string bucketId) view returns((address,string,uint8,uint256,uint8,int64,address,uint32,uint64,uint8,(string,string)[],bool) bucketInfo, (bool,uint256,uint256) bucketExtraInfo)
func (_IStorage *IStorageCaller) HeadBucketById(opts *bind.CallOpts, bucketId string) (struct {
	BucketInfo      BucketInfo
	BucketExtraInfo BucketExtraInfo
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "headBucketById", bucketId)

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

// HeadBucketById is a free data retrieval call binding the contract method 0xb5d73569.
//
// Solidity: function headBucketById(string bucketId) view returns((address,string,uint8,uint256,uint8,int64,address,uint32,uint64,uint8,(string,string)[],bool) bucketInfo, (bool,uint256,uint256) bucketExtraInfo)
func (_IStorage *IStorageSession) HeadBucketById(bucketId string) (struct {
	BucketInfo      BucketInfo
	BucketExtraInfo BucketExtraInfo
}, error) {
	return _IStorage.Contract.HeadBucketById(&_IStorage.CallOpts, bucketId)
}

// HeadBucketById is a free data retrieval call binding the contract method 0xb5d73569.
//
// Solidity: function headBucketById(string bucketId) view returns((address,string,uint8,uint256,uint8,int64,address,uint32,uint64,uint8,(string,string)[],bool) bucketInfo, (bool,uint256,uint256) bucketExtraInfo)
func (_IStorage *IStorageCallerSession) HeadBucketById(bucketId string) (struct {
	BucketInfo      BucketInfo
	BucketExtraInfo BucketExtraInfo
}, error) {
	return _IStorage.Contract.HeadBucketById(&_IStorage.CallOpts, bucketId)
}

// HeadBucketExtra is a free data retrieval call binding the contract method 0xfe2fb5e1.
//
// Solidity: function headBucketExtra(string bucketName) view returns((int64,uint64,(uint32,uint32,uint64,uint64)[],uint32) extraInfo)
func (_IStorage *IStorageCaller) HeadBucketExtra(opts *bind.CallOpts, bucketName string) (InternalBucketInfo, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "headBucketExtra", bucketName)

	if err != nil {
		return *new(InternalBucketInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(InternalBucketInfo)).(*InternalBucketInfo)

	return out0, err

}

// HeadBucketExtra is a free data retrieval call binding the contract method 0xfe2fb5e1.
//
// Solidity: function headBucketExtra(string bucketName) view returns((int64,uint64,(uint32,uint32,uint64,uint64)[],uint32) extraInfo)
func (_IStorage *IStorageSession) HeadBucketExtra(bucketName string) (InternalBucketInfo, error) {
	return _IStorage.Contract.HeadBucketExtra(&_IStorage.CallOpts, bucketName)
}

// HeadBucketExtra is a free data retrieval call binding the contract method 0xfe2fb5e1.
//
// Solidity: function headBucketExtra(string bucketName) view returns((int64,uint64,(uint32,uint32,uint64,uint64)[],uint32) extraInfo)
func (_IStorage *IStorageCallerSession) HeadBucketExtra(bucketName string) (InternalBucketInfo, error) {
	return _IStorage.Contract.HeadBucketExtra(&_IStorage.CallOpts, bucketName)
}

// HeadBucketNFT is a free data retrieval call binding the contract method 0xec4bc21c.
//
// Solidity: function headBucketNFT(string tokenId) view returns((string,string,string,string,(string,string)[]) bucketMetaData)
func (_IStorage *IStorageCaller) HeadBucketNFT(opts *bind.CallOpts, tokenId string) (BucketMetaData, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "headBucketNFT", tokenId)

	if err != nil {
		return *new(BucketMetaData), err
	}

	out0 := *abi.ConvertType(out[0], new(BucketMetaData)).(*BucketMetaData)

	return out0, err

}

// HeadBucketNFT is a free data retrieval call binding the contract method 0xec4bc21c.
//
// Solidity: function headBucketNFT(string tokenId) view returns((string,string,string,string,(string,string)[]) bucketMetaData)
func (_IStorage *IStorageSession) HeadBucketNFT(tokenId string) (BucketMetaData, error) {
	return _IStorage.Contract.HeadBucketNFT(&_IStorage.CallOpts, tokenId)
}

// HeadBucketNFT is a free data retrieval call binding the contract method 0xec4bc21c.
//
// Solidity: function headBucketNFT(string tokenId) view returns((string,string,string,string,(string,string)[]) bucketMetaData)
func (_IStorage *IStorageCallerSession) HeadBucketNFT(tokenId string) (BucketMetaData, error) {
	return _IStorage.Contract.HeadBucketNFT(&_IStorage.CallOpts, tokenId)
}

// HeadGroup is a free data retrieval call binding the contract method 0x88c070b5.
//
// Solidity: function headGroup(address groupOwner, string groupName) view returns((address,string,uint8,uint256,string,(string,string)[]) groupInfo)
func (_IStorage *IStorageCaller) HeadGroup(opts *bind.CallOpts, groupOwner common.Address, groupName string) (GroupInfo, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "headGroup", groupOwner, groupName)

	if err != nil {
		return *new(GroupInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(GroupInfo)).(*GroupInfo)

	return out0, err

}

// HeadGroup is a free data retrieval call binding the contract method 0x88c070b5.
//
// Solidity: function headGroup(address groupOwner, string groupName) view returns((address,string,uint8,uint256,string,(string,string)[]) groupInfo)
func (_IStorage *IStorageSession) HeadGroup(groupOwner common.Address, groupName string) (GroupInfo, error) {
	return _IStorage.Contract.HeadGroup(&_IStorage.CallOpts, groupOwner, groupName)
}

// HeadGroup is a free data retrieval call binding the contract method 0x88c070b5.
//
// Solidity: function headGroup(address groupOwner, string groupName) view returns((address,string,uint8,uint256,string,(string,string)[]) groupInfo)
func (_IStorage *IStorageCallerSession) HeadGroup(groupOwner common.Address, groupName string) (GroupInfo, error) {
	return _IStorage.Contract.HeadGroup(&_IStorage.CallOpts, groupOwner, groupName)
}

// HeadGroupMember is a free data retrieval call binding the contract method 0xd4a00c23.
//
// Solidity: function headGroupMember(address member, address groupOwner, string groupName) view returns((uint256,uint256,address,int64) groupMember)
func (_IStorage *IStorageCaller) HeadGroupMember(opts *bind.CallOpts, member common.Address, groupOwner common.Address, groupName string) (GroupMember, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "headGroupMember", member, groupOwner, groupName)

	if err != nil {
		return *new(GroupMember), err
	}

	out0 := *abi.ConvertType(out[0], new(GroupMember)).(*GroupMember)

	return out0, err

}

// HeadGroupMember is a free data retrieval call binding the contract method 0xd4a00c23.
//
// Solidity: function headGroupMember(address member, address groupOwner, string groupName) view returns((uint256,uint256,address,int64) groupMember)
func (_IStorage *IStorageSession) HeadGroupMember(member common.Address, groupOwner common.Address, groupName string) (GroupMember, error) {
	return _IStorage.Contract.HeadGroupMember(&_IStorage.CallOpts, member, groupOwner, groupName)
}

// HeadGroupMember is a free data retrieval call binding the contract method 0xd4a00c23.
//
// Solidity: function headGroupMember(address member, address groupOwner, string groupName) view returns((uint256,uint256,address,int64) groupMember)
func (_IStorage *IStorageCallerSession) HeadGroupMember(member common.Address, groupOwner common.Address, groupName string) (GroupMember, error) {
	return _IStorage.Contract.HeadGroupMember(&_IStorage.CallOpts, member, groupOwner, groupName)
}

// HeadGroupNFT is a free data retrieval call binding the contract method 0xd4084ff9.
//
// Solidity: function headGroupNFT(string tokenId) view returns((string,string,string,string,(string,string)[]) groupMetaData)
func (_IStorage *IStorageCaller) HeadGroupNFT(opts *bind.CallOpts, tokenId string) (GroupMetaData, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "headGroupNFT", tokenId)

	if err != nil {
		return *new(GroupMetaData), err
	}

	out0 := *abi.ConvertType(out[0], new(GroupMetaData)).(*GroupMetaData)

	return out0, err

}

// HeadGroupNFT is a free data retrieval call binding the contract method 0xd4084ff9.
//
// Solidity: function headGroupNFT(string tokenId) view returns((string,string,string,string,(string,string)[]) groupMetaData)
func (_IStorage *IStorageSession) HeadGroupNFT(tokenId string) (GroupMetaData, error) {
	return _IStorage.Contract.HeadGroupNFT(&_IStorage.CallOpts, tokenId)
}

// HeadGroupNFT is a free data retrieval call binding the contract method 0xd4084ff9.
//
// Solidity: function headGroupNFT(string tokenId) view returns((string,string,string,string,(string,string)[]) groupMetaData)
func (_IStorage *IStorageCallerSession) HeadGroupNFT(tokenId string) (GroupMetaData, error) {
	return _IStorage.Contract.HeadGroupNFT(&_IStorage.CallOpts, tokenId)
}

// HeadObject is a free data retrieval call binding the contract method 0x2523d202.
//
// Solidity: function headObject(string bucketName, string objectName) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64) objectInfo, (uint32,uint32,uint32,uint32[],uint64,address,string) globalVirtualGroup)
func (_IStorage *IStorageCaller) HeadObject(opts *bind.CallOpts, bucketName string, objectName string) (struct {
	ObjectInfo         ObjectInfo
	GlobalVirtualGroup GlobalVirtualGroup
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "headObject", bucketName, objectName)

	outstruct := new(struct {
		ObjectInfo         ObjectInfo
		GlobalVirtualGroup GlobalVirtualGroup
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ObjectInfo = *abi.ConvertType(out[0], new(ObjectInfo)).(*ObjectInfo)
	outstruct.GlobalVirtualGroup = *abi.ConvertType(out[1], new(GlobalVirtualGroup)).(*GlobalVirtualGroup)

	return *outstruct, err

}

// HeadObject is a free data retrieval call binding the contract method 0x2523d202.
//
// Solidity: function headObject(string bucketName, string objectName) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64) objectInfo, (uint32,uint32,uint32,uint32[],uint64,address,string) globalVirtualGroup)
func (_IStorage *IStorageSession) HeadObject(bucketName string, objectName string) (struct {
	ObjectInfo         ObjectInfo
	GlobalVirtualGroup GlobalVirtualGroup
}, error) {
	return _IStorage.Contract.HeadObject(&_IStorage.CallOpts, bucketName, objectName)
}

// HeadObject is a free data retrieval call binding the contract method 0x2523d202.
//
// Solidity: function headObject(string bucketName, string objectName) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64) objectInfo, (uint32,uint32,uint32,uint32[],uint64,address,string) globalVirtualGroup)
func (_IStorage *IStorageCallerSession) HeadObject(bucketName string, objectName string) (struct {
	ObjectInfo         ObjectInfo
	GlobalVirtualGroup GlobalVirtualGroup
}, error) {
	return _IStorage.Contract.HeadObject(&_IStorage.CallOpts, bucketName, objectName)
}

// HeadObjectById is a free data retrieval call binding the contract method 0xfce9b24b.
//
// Solidity: function headObjectById(string objectId) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64) objectInfo, (uint32,uint32,uint32,uint32[],uint64,address,string) globalVirtualGroup)
func (_IStorage *IStorageCaller) HeadObjectById(opts *bind.CallOpts, objectId string) (struct {
	ObjectInfo         ObjectInfo
	GlobalVirtualGroup GlobalVirtualGroup
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "headObjectById", objectId)

	outstruct := new(struct {
		ObjectInfo         ObjectInfo
		GlobalVirtualGroup GlobalVirtualGroup
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ObjectInfo = *abi.ConvertType(out[0], new(ObjectInfo)).(*ObjectInfo)
	outstruct.GlobalVirtualGroup = *abi.ConvertType(out[1], new(GlobalVirtualGroup)).(*GlobalVirtualGroup)

	return *outstruct, err

}

// HeadObjectById is a free data retrieval call binding the contract method 0xfce9b24b.
//
// Solidity: function headObjectById(string objectId) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64) objectInfo, (uint32,uint32,uint32,uint32[],uint64,address,string) globalVirtualGroup)
func (_IStorage *IStorageSession) HeadObjectById(objectId string) (struct {
	ObjectInfo         ObjectInfo
	GlobalVirtualGroup GlobalVirtualGroup
}, error) {
	return _IStorage.Contract.HeadObjectById(&_IStorage.CallOpts, objectId)
}

// HeadObjectById is a free data retrieval call binding the contract method 0xfce9b24b.
//
// Solidity: function headObjectById(string objectId) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64) objectInfo, (uint32,uint32,uint32,uint32[],uint64,address,string) globalVirtualGroup)
func (_IStorage *IStorageCallerSession) HeadObjectById(objectId string) (struct {
	ObjectInfo         ObjectInfo
	GlobalVirtualGroup GlobalVirtualGroup
}, error) {
	return _IStorage.Contract.HeadObjectById(&_IStorage.CallOpts, objectId)
}

// HeadObjectNFT is a free data retrieval call binding the contract method 0x09f2a357.
//
// Solidity: function headObjectNFT(string tokenId) view returns((string,string,string,string,(string,string)[]) objectMetaData)
func (_IStorage *IStorageCaller) HeadObjectNFT(opts *bind.CallOpts, tokenId string) (ObjectMetaData, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "headObjectNFT", tokenId)

	if err != nil {
		return *new(ObjectMetaData), err
	}

	out0 := *abi.ConvertType(out[0], new(ObjectMetaData)).(*ObjectMetaData)

	return out0, err

}

// HeadObjectNFT is a free data retrieval call binding the contract method 0x09f2a357.
//
// Solidity: function headObjectNFT(string tokenId) view returns((string,string,string,string,(string,string)[]) objectMetaData)
func (_IStorage *IStorageSession) HeadObjectNFT(tokenId string) (ObjectMetaData, error) {
	return _IStorage.Contract.HeadObjectNFT(&_IStorage.CallOpts, tokenId)
}

// HeadObjectNFT is a free data retrieval call binding the contract method 0x09f2a357.
//
// Solidity: function headObjectNFT(string tokenId) view returns((string,string,string,string,(string,string)[]) objectMetaData)
func (_IStorage *IStorageCallerSession) HeadObjectNFT(tokenId string) (ObjectMetaData, error) {
	return _IStorage.Contract.HeadObjectNFT(&_IStorage.CallOpts, tokenId)
}

// HeadShadowObject is a free data retrieval call binding the contract method 0xb3b09508.
//
// Solidity: function headShadowObject(string bucketName, string objectName) view returns((string,uint256,string,uint64,string[],int64,int64) objectInfo)
func (_IStorage *IStorageCaller) HeadShadowObject(opts *bind.CallOpts, bucketName string, objectName string) (ShadowObjectInfo, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "headShadowObject", bucketName, objectName)

	if err != nil {
		return *new(ShadowObjectInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ShadowObjectInfo)).(*ShadowObjectInfo)

	return out0, err

}

// HeadShadowObject is a free data retrieval call binding the contract method 0xb3b09508.
//
// Solidity: function headShadowObject(string bucketName, string objectName) view returns((string,uint256,string,uint64,string[],int64,int64) objectInfo)
func (_IStorage *IStorageSession) HeadShadowObject(bucketName string, objectName string) (ShadowObjectInfo, error) {
	return _IStorage.Contract.HeadShadowObject(&_IStorage.CallOpts, bucketName, objectName)
}

// HeadShadowObject is a free data retrieval call binding the contract method 0xb3b09508.
//
// Solidity: function headShadowObject(string bucketName, string objectName) view returns((string,uint256,string,uint64,string[],int64,int64) objectInfo)
func (_IStorage *IStorageCallerSession) HeadShadowObject(bucketName string, objectName string) (ShadowObjectInfo, error) {
	return _IStorage.Contract.HeadShadowObject(&_IStorage.CallOpts, bucketName, objectName)
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

// ListGroups is a free data retrieval call binding the contract method 0x6d294e2b.
//
// Solidity: function listGroups((bytes,uint64,uint64,bool,bool) pagination, address groupOwner) view returns((address,string,uint8,uint256,string,(string,string)[])[] groupInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageCaller) ListGroups(opts *bind.CallOpts, pagination PageRequest, groupOwner common.Address) (struct {
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

// ListGroups is a free data retrieval call binding the contract method 0x6d294e2b.
//
// Solidity: function listGroups((bytes,uint64,uint64,bool,bool) pagination, address groupOwner) view returns((address,string,uint8,uint256,string,(string,string)[])[] groupInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageSession) ListGroups(pagination PageRequest, groupOwner common.Address) (struct {
	GroupInfos   []GroupInfo
	PageResponse PageResponse
}, error) {
	return _IStorage.Contract.ListGroups(&_IStorage.CallOpts, pagination, groupOwner)
}

// ListGroups is a free data retrieval call binding the contract method 0x6d294e2b.
//
// Solidity: function listGroups((bytes,uint64,uint64,bool,bool) pagination, address groupOwner) view returns((address,string,uint8,uint256,string,(string,string)[])[] groupInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageCallerSession) ListGroups(pagination PageRequest, groupOwner common.Address) (struct {
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

// ListObjectsByBucketId is a free data retrieval call binding the contract method 0x640d2a7d.
//
// Solidity: function listObjectsByBucketId((bytes,uint64,uint64,bool,bool) pagination, string bucketId) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64)[] objectInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageCaller) ListObjectsByBucketId(opts *bind.CallOpts, pagination PageRequest, bucketId string) (struct {
	ObjectInfos  []ObjectInfo
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "listObjectsByBucketId", pagination, bucketId)

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

// ListObjectsByBucketId is a free data retrieval call binding the contract method 0x640d2a7d.
//
// Solidity: function listObjectsByBucketId((bytes,uint64,uint64,bool,bool) pagination, string bucketId) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64)[] objectInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageSession) ListObjectsByBucketId(pagination PageRequest, bucketId string) (struct {
	ObjectInfos  []ObjectInfo
	PageResponse PageResponse
}, error) {
	return _IStorage.Contract.ListObjectsByBucketId(&_IStorage.CallOpts, pagination, bucketId)
}

// ListObjectsByBucketId is a free data retrieval call binding the contract method 0x640d2a7d.
//
// Solidity: function listObjectsByBucketId((bytes,uint64,uint64,bool,bool) pagination, string bucketId) view returns((address,address,string,string,uint256,uint32,uint64,uint8,string,int64,uint8,uint8,uint8,string[],(string,string)[],bool,int64,address,int64)[] objectInfos, (bytes,uint64) pageResponse)
func (_IStorage *IStorageCallerSession) ListObjectsByBucketId(pagination PageRequest, bucketId string) (struct {
	ObjectInfos  []ObjectInfo
	PageResponse PageResponse
}, error) {
	return _IStorage.Contract.ListObjectsByBucketId(&_IStorage.CallOpts, pagination, bucketId)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((uint64,uint32,uint32,uint64),uint64,string,string,string,string,string,string,uint32,uint64,uint64,uint64,int64,uint64,uint64,uint64,uint32,string,string,string,string,string,string,string,string) params)
func (_IStorage *IStorageCaller) Params(opts *bind.CallOpts) (Params, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "params")

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((uint64,uint32,uint32,uint64),uint64,string,string,string,string,string,string,uint32,uint64,uint64,uint64,int64,uint64,uint64,uint64,uint32,string,string,string,string,string,string,string,string) params)
func (_IStorage *IStorageSession) Params() (Params, error) {
	return _IStorage.Contract.Params(&_IStorage.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((uint64,uint32,uint32,uint64),uint64,string,string,string,string,string,string,uint32,uint64,uint64,uint64,int64,uint64,uint64,uint64,uint32,string,string,string,string,string,string,string,string) params)
func (_IStorage *IStorageCallerSession) Params() (Params, error) {
	return _IStorage.Contract.Params(&_IStorage.CallOpts)
}

// QueryGroupMembersExist is a free data retrieval call binding the contract method 0x4645d454.
//
// Solidity: function queryGroupMembersExist(string groupId, string[] members) view returns(string[] checkMembers, bool[] exists)
func (_IStorage *IStorageCaller) QueryGroupMembersExist(opts *bind.CallOpts, groupId string, members []string) (struct {
	CheckMembers []string
	Exists       []bool
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "queryGroupMembersExist", groupId, members)

	outstruct := new(struct {
		CheckMembers []string
		Exists       []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CheckMembers = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Exists = *abi.ConvertType(out[1], new([]bool)).(*[]bool)

	return *outstruct, err

}

// QueryGroupMembersExist is a free data retrieval call binding the contract method 0x4645d454.
//
// Solidity: function queryGroupMembersExist(string groupId, string[] members) view returns(string[] checkMembers, bool[] exists)
func (_IStorage *IStorageSession) QueryGroupMembersExist(groupId string, members []string) (struct {
	CheckMembers []string
	Exists       []bool
}, error) {
	return _IStorage.Contract.QueryGroupMembersExist(&_IStorage.CallOpts, groupId, members)
}

// QueryGroupMembersExist is a free data retrieval call binding the contract method 0x4645d454.
//
// Solidity: function queryGroupMembersExist(string groupId, string[] members) view returns(string[] checkMembers, bool[] exists)
func (_IStorage *IStorageCallerSession) QueryGroupMembersExist(groupId string, members []string) (struct {
	CheckMembers []string
	Exists       []bool
}, error) {
	return _IStorage.Contract.QueryGroupMembersExist(&_IStorage.CallOpts, groupId, members)
}

// QueryGroupsExist is a free data retrieval call binding the contract method 0xde1ecb5b.
//
// Solidity: function queryGroupsExist(string groupOwner, string[] groupNames) view returns(string[] checkGroupNames, bool[] exists)
func (_IStorage *IStorageCaller) QueryGroupsExist(opts *bind.CallOpts, groupOwner string, groupNames []string) (struct {
	CheckGroupNames []string
	Exists          []bool
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "queryGroupsExist", groupOwner, groupNames)

	outstruct := new(struct {
		CheckGroupNames []string
		Exists          []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CheckGroupNames = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Exists = *abi.ConvertType(out[1], new([]bool)).(*[]bool)

	return *outstruct, err

}

// QueryGroupsExist is a free data retrieval call binding the contract method 0xde1ecb5b.
//
// Solidity: function queryGroupsExist(string groupOwner, string[] groupNames) view returns(string[] checkGroupNames, bool[] exists)
func (_IStorage *IStorageSession) QueryGroupsExist(groupOwner string, groupNames []string) (struct {
	CheckGroupNames []string
	Exists          []bool
}, error) {
	return _IStorage.Contract.QueryGroupsExist(&_IStorage.CallOpts, groupOwner, groupNames)
}

// QueryGroupsExist is a free data retrieval call binding the contract method 0xde1ecb5b.
//
// Solidity: function queryGroupsExist(string groupOwner, string[] groupNames) view returns(string[] checkGroupNames, bool[] exists)
func (_IStorage *IStorageCallerSession) QueryGroupsExist(groupOwner string, groupNames []string) (struct {
	CheckGroupNames []string
	Exists          []bool
}, error) {
	return _IStorage.Contract.QueryGroupsExist(&_IStorage.CallOpts, groupOwner, groupNames)
}

// QueryGroupsExistById is a free data retrieval call binding the contract method 0x0e888906.
//
// Solidity: function queryGroupsExistById(string[] groupIds) view returns(string[] checkGroupIds, bool[] exists)
func (_IStorage *IStorageCaller) QueryGroupsExistById(opts *bind.CallOpts, groupIds []string) (struct {
	CheckGroupIds []string
	Exists        []bool
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "queryGroupsExistById", groupIds)

	outstruct := new(struct {
		CheckGroupIds []string
		Exists        []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CheckGroupIds = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Exists = *abi.ConvertType(out[1], new([]bool)).(*[]bool)

	return *outstruct, err

}

// QueryGroupsExistById is a free data retrieval call binding the contract method 0x0e888906.
//
// Solidity: function queryGroupsExistById(string[] groupIds) view returns(string[] checkGroupIds, bool[] exists)
func (_IStorage *IStorageSession) QueryGroupsExistById(groupIds []string) (struct {
	CheckGroupIds []string
	Exists        []bool
}, error) {
	return _IStorage.Contract.QueryGroupsExistById(&_IStorage.CallOpts, groupIds)
}

// QueryGroupsExistById is a free data retrieval call binding the contract method 0x0e888906.
//
// Solidity: function queryGroupsExistById(string[] groupIds) view returns(string[] checkGroupIds, bool[] exists)
func (_IStorage *IStorageCallerSession) QueryGroupsExistById(groupIds []string) (struct {
	CheckGroupIds []string
	Exists        []bool
}, error) {
	return _IStorage.Contract.QueryGroupsExistById(&_IStorage.CallOpts, groupIds)
}

// QueryIsPriceChanged is a free data retrieval call binding the contract method 0x192cbbd4.
//
// Solidity: function queryIsPriceChanged(string bucketName) view returns((bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) isPriceChanged)
func (_IStorage *IStorageCaller) QueryIsPriceChanged(opts *bind.CallOpts, bucketName string) (IsPriceChanged, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "queryIsPriceChanged", bucketName)

	if err != nil {
		return *new(IsPriceChanged), err
	}

	out0 := *abi.ConvertType(out[0], new(IsPriceChanged)).(*IsPriceChanged)

	return out0, err

}

// QueryIsPriceChanged is a free data retrieval call binding the contract method 0x192cbbd4.
//
// Solidity: function queryIsPriceChanged(string bucketName) view returns((bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) isPriceChanged)
func (_IStorage *IStorageSession) QueryIsPriceChanged(bucketName string) (IsPriceChanged, error) {
	return _IStorage.Contract.QueryIsPriceChanged(&_IStorage.CallOpts, bucketName)
}

// QueryIsPriceChanged is a free data retrieval call binding the contract method 0x192cbbd4.
//
// Solidity: function queryIsPriceChanged(string bucketName) view returns((bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) isPriceChanged)
func (_IStorage *IStorageCallerSession) QueryIsPriceChanged(bucketName string) (IsPriceChanged, error) {
	return _IStorage.Contract.QueryIsPriceChanged(&_IStorage.CallOpts, bucketName)
}

// QueryLockFee is a free data retrieval call binding the contract method 0xaff2e95d.
//
// Solidity: function queryLockFee(string primarySpAddress, int64 createAt, uint64 payloadSize) view returns(uint256 amount)
func (_IStorage *IStorageCaller) QueryLockFee(opts *bind.CallOpts, primarySpAddress string, createAt int64, payloadSize uint64) (*big.Int, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "queryLockFee", primarySpAddress, createAt, payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryLockFee is a free data retrieval call binding the contract method 0xaff2e95d.
//
// Solidity: function queryLockFee(string primarySpAddress, int64 createAt, uint64 payloadSize) view returns(uint256 amount)
func (_IStorage *IStorageSession) QueryLockFee(primarySpAddress string, createAt int64, payloadSize uint64) (*big.Int, error) {
	return _IStorage.Contract.QueryLockFee(&_IStorage.CallOpts, primarySpAddress, createAt, payloadSize)
}

// QueryLockFee is a free data retrieval call binding the contract method 0xaff2e95d.
//
// Solidity: function queryLockFee(string primarySpAddress, int64 createAt, uint64 payloadSize) view returns(uint256 amount)
func (_IStorage *IStorageCallerSession) QueryLockFee(primarySpAddress string, createAt int64, payloadSize uint64) (*big.Int, error) {
	return _IStorage.Contract.QueryLockFee(&_IStorage.CallOpts, primarySpAddress, createAt, payloadSize)
}

// QueryParamsByTimestamp is a free data retrieval call binding the contract method 0xbe0697fd.
//
// Solidity: function queryParamsByTimestamp(int64 timestamp) view returns(((uint64,uint32,uint32,uint64),uint64,string,string,string,string,string,string,uint32,uint64,uint64,uint64,int64,uint64,uint64,uint64,uint32,string,string,string,string,string,string,string,string) params)
func (_IStorage *IStorageCaller) QueryParamsByTimestamp(opts *bind.CallOpts, timestamp int64) (Params, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "queryParamsByTimestamp", timestamp)

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// QueryParamsByTimestamp is a free data retrieval call binding the contract method 0xbe0697fd.
//
// Solidity: function queryParamsByTimestamp(int64 timestamp) view returns(((uint64,uint32,uint32,uint64),uint64,string,string,string,string,string,string,uint32,uint64,uint64,uint64,int64,uint64,uint64,uint64,uint32,string,string,string,string,string,string,string,string) params)
func (_IStorage *IStorageSession) QueryParamsByTimestamp(timestamp int64) (Params, error) {
	return _IStorage.Contract.QueryParamsByTimestamp(&_IStorage.CallOpts, timestamp)
}

// QueryParamsByTimestamp is a free data retrieval call binding the contract method 0xbe0697fd.
//
// Solidity: function queryParamsByTimestamp(int64 timestamp) view returns(((uint64,uint32,uint32,uint64),uint64,string,string,string,string,string,string,uint32,uint64,uint64,uint64,int64,uint64,uint64,uint64,uint32,string,string,string,string,string,string,string,string) params)
func (_IStorage *IStorageCallerSession) QueryParamsByTimestamp(timestamp int64) (Params, error) {
	return _IStorage.Contract.QueryParamsByTimestamp(&_IStorage.CallOpts, timestamp)
}

// QueryPaymentAccountBucketFlowRateLimit is a free data retrieval call binding the contract method 0x86de7571.
//
// Solidity: function queryPaymentAccountBucketFlowRateLimit(string paymentAccount, string bucketOwner, string bucketName) view returns(bool isSet, uint256 flowRateLimit)
func (_IStorage *IStorageCaller) QueryPaymentAccountBucketFlowRateLimit(opts *bind.CallOpts, paymentAccount string, bucketOwner string, bucketName string) (struct {
	IsSet         bool
	FlowRateLimit *big.Int
}, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "queryPaymentAccountBucketFlowRateLimit", paymentAccount, bucketOwner, bucketName)

	outstruct := new(struct {
		IsSet         bool
		FlowRateLimit *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSet = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.FlowRateLimit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QueryPaymentAccountBucketFlowRateLimit is a free data retrieval call binding the contract method 0x86de7571.
//
// Solidity: function queryPaymentAccountBucketFlowRateLimit(string paymentAccount, string bucketOwner, string bucketName) view returns(bool isSet, uint256 flowRateLimit)
func (_IStorage *IStorageSession) QueryPaymentAccountBucketFlowRateLimit(paymentAccount string, bucketOwner string, bucketName string) (struct {
	IsSet         bool
	FlowRateLimit *big.Int
}, error) {
	return _IStorage.Contract.QueryPaymentAccountBucketFlowRateLimit(&_IStorage.CallOpts, paymentAccount, bucketOwner, bucketName)
}

// QueryPaymentAccountBucketFlowRateLimit is a free data retrieval call binding the contract method 0x86de7571.
//
// Solidity: function queryPaymentAccountBucketFlowRateLimit(string paymentAccount, string bucketOwner, string bucketName) view returns(bool isSet, uint256 flowRateLimit)
func (_IStorage *IStorageCallerSession) QueryPaymentAccountBucketFlowRateLimit(paymentAccount string, bucketOwner string, bucketName string) (struct {
	IsSet         bool
	FlowRateLimit *big.Int
}, error) {
	return _IStorage.Contract.QueryPaymentAccountBucketFlowRateLimit(&_IStorage.CallOpts, paymentAccount, bucketOwner, bucketName)
}

// QueryPolicyById is a free data retrieval call binding the contract method 0x41220cfe.
//
// Solidity: function queryPolicyById(string policyId) view returns((uint256,(int32,string),int32,uint256,(int32,int32[],string[],int64,uint64)[],int64) policy)
func (_IStorage *IStorageCaller) QueryPolicyById(opts *bind.CallOpts, policyId string) (Policy, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "queryPolicyById", policyId)

	if err != nil {
		return *new(Policy), err
	}

	out0 := *abi.ConvertType(out[0], new(Policy)).(*Policy)

	return out0, err

}

// QueryPolicyById is a free data retrieval call binding the contract method 0x41220cfe.
//
// Solidity: function queryPolicyById(string policyId) view returns((uint256,(int32,string),int32,uint256,(int32,int32[],string[],int64,uint64)[],int64) policy)
func (_IStorage *IStorageSession) QueryPolicyById(policyId string) (Policy, error) {
	return _IStorage.Contract.QueryPolicyById(&_IStorage.CallOpts, policyId)
}

// QueryPolicyById is a free data retrieval call binding the contract method 0x41220cfe.
//
// Solidity: function queryPolicyById(string policyId) view returns((uint256,(int32,string),int32,uint256,(int32,int32[],string[],int64,uint64)[],int64) policy)
func (_IStorage *IStorageCallerSession) QueryPolicyById(policyId string) (Policy, error) {
	return _IStorage.Contract.QueryPolicyById(&_IStorage.CallOpts, policyId)
}

// QueryPolicyForAccount is a free data retrieval call binding the contract method 0xe37ea89c.
//
// Solidity: function queryPolicyForAccount(string resource, string principalAddr) view returns((uint256,(int32,string),int32,uint256,(int32,int32[],string[],int64,uint64)[],int64) policy)
func (_IStorage *IStorageCaller) QueryPolicyForAccount(opts *bind.CallOpts, resource string, principalAddr string) (Policy, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "queryPolicyForAccount", resource, principalAddr)

	if err != nil {
		return *new(Policy), err
	}

	out0 := *abi.ConvertType(out[0], new(Policy)).(*Policy)

	return out0, err

}

// QueryPolicyForAccount is a free data retrieval call binding the contract method 0xe37ea89c.
//
// Solidity: function queryPolicyForAccount(string resource, string principalAddr) view returns((uint256,(int32,string),int32,uint256,(int32,int32[],string[],int64,uint64)[],int64) policy)
func (_IStorage *IStorageSession) QueryPolicyForAccount(resource string, principalAddr string) (Policy, error) {
	return _IStorage.Contract.QueryPolicyForAccount(&_IStorage.CallOpts, resource, principalAddr)
}

// QueryPolicyForAccount is a free data retrieval call binding the contract method 0xe37ea89c.
//
// Solidity: function queryPolicyForAccount(string resource, string principalAddr) view returns((uint256,(int32,string),int32,uint256,(int32,int32[],string[],int64,uint64)[],int64) policy)
func (_IStorage *IStorageCallerSession) QueryPolicyForAccount(resource string, principalAddr string) (Policy, error) {
	return _IStorage.Contract.QueryPolicyForAccount(&_IStorage.CallOpts, resource, principalAddr)
}

// QueryPolicyForGroup is a free data retrieval call binding the contract method 0x2e1eaadd.
//
// Solidity: function queryPolicyForGroup(string resource, uint256 groupId) view returns((uint256,(int32,string),int32,uint256,(int32,int32[],string[],int64,uint64)[],int64) policy)
func (_IStorage *IStorageCaller) QueryPolicyForGroup(opts *bind.CallOpts, resource string, groupId *big.Int) (Policy, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "queryPolicyForGroup", resource, groupId)

	if err != nil {
		return *new(Policy), err
	}

	out0 := *abi.ConvertType(out[0], new(Policy)).(*Policy)

	return out0, err

}

// QueryPolicyForGroup is a free data retrieval call binding the contract method 0x2e1eaadd.
//
// Solidity: function queryPolicyForGroup(string resource, uint256 groupId) view returns((uint256,(int32,string),int32,uint256,(int32,int32[],string[],int64,uint64)[],int64) policy)
func (_IStorage *IStorageSession) QueryPolicyForGroup(resource string, groupId *big.Int) (Policy, error) {
	return _IStorage.Contract.QueryPolicyForGroup(&_IStorage.CallOpts, resource, groupId)
}

// QueryPolicyForGroup is a free data retrieval call binding the contract method 0x2e1eaadd.
//
// Solidity: function queryPolicyForGroup(string resource, uint256 groupId) view returns((uint256,(int32,string),int32,uint256,(int32,int32[],string[],int64,uint64)[],int64) policy)
func (_IStorage *IStorageCallerSession) QueryPolicyForGroup(resource string, groupId *big.Int) (Policy, error) {
	return _IStorage.Contract.QueryPolicyForGroup(&_IStorage.CallOpts, resource, groupId)
}

// QueryQuotaUpdateTime is a free data retrieval call binding the contract method 0xed51f0bf.
//
// Solidity: function queryQuotaUpdateTime(string bucketName) view returns(int64 updateAt)
func (_IStorage *IStorageCaller) QueryQuotaUpdateTime(opts *bind.CallOpts, bucketName string) (int64, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "queryQuotaUpdateTime", bucketName)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// QueryQuotaUpdateTime is a free data retrieval call binding the contract method 0xed51f0bf.
//
// Solidity: function queryQuotaUpdateTime(string bucketName) view returns(int64 updateAt)
func (_IStorage *IStorageSession) QueryQuotaUpdateTime(bucketName string) (int64, error) {
	return _IStorage.Contract.QueryQuotaUpdateTime(&_IStorage.CallOpts, bucketName)
}

// QueryQuotaUpdateTime is a free data retrieval call binding the contract method 0xed51f0bf.
//
// Solidity: function queryQuotaUpdateTime(string bucketName) view returns(int64 updateAt)
func (_IStorage *IStorageCallerSession) QueryQuotaUpdateTime(bucketName string) (int64, error) {
	return _IStorage.Contract.QueryQuotaUpdateTime(&_IStorage.CallOpts, bucketName)
}

// VerifyPermission is a free data retrieval call binding the contract method 0xc9640758.
//
// Solidity: function verifyPermission(string bucketName, string objectName, int32 actionType) view returns(int32 effect)
func (_IStorage *IStorageCaller) VerifyPermission(opts *bind.CallOpts, bucketName string, objectName string, actionType int32) (int32, error) {
	var out []interface{}
	err := _IStorage.contract.Call(opts, &out, "verifyPermission", bucketName, objectName, actionType)

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// VerifyPermission is a free data retrieval call binding the contract method 0xc9640758.
//
// Solidity: function verifyPermission(string bucketName, string objectName, int32 actionType) view returns(int32 effect)
func (_IStorage *IStorageSession) VerifyPermission(bucketName string, objectName string, actionType int32) (int32, error) {
	return _IStorage.Contract.VerifyPermission(&_IStorage.CallOpts, bucketName, objectName, actionType)
}

// VerifyPermission is a free data retrieval call binding the contract method 0xc9640758.
//
// Solidity: function verifyPermission(string bucketName, string objectName, int32 actionType) view returns(int32 effect)
func (_IStorage *IStorageCallerSession) VerifyPermission(bucketName string, objectName string, actionType int32) (int32, error) {
	return _IStorage.Contract.VerifyPermission(&_IStorage.CallOpts, bucketName, objectName, actionType)
}

// CancelCreateObject is a paid mutator transaction binding the contract method 0xa0cd9e20.
//
// Solidity: function cancelCreateObject(string bucketName, string objectName) returns(bool success)
func (_IStorage *IStorageTransactor) CancelCreateObject(opts *bind.TransactOpts, bucketName string, objectName string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "cancelCreateObject", bucketName, objectName)
}

// CancelCreateObject is a paid mutator transaction binding the contract method 0xa0cd9e20.
//
// Solidity: function cancelCreateObject(string bucketName, string objectName) returns(bool success)
func (_IStorage *IStorageSession) CancelCreateObject(bucketName string, objectName string) (*types.Transaction, error) {
	return _IStorage.Contract.CancelCreateObject(&_IStorage.TransactOpts, bucketName, objectName)
}

// CancelCreateObject is a paid mutator transaction binding the contract method 0xa0cd9e20.
//
// Solidity: function cancelCreateObject(string bucketName, string objectName) returns(bool success)
func (_IStorage *IStorageTransactorSession) CancelCreateObject(bucketName string, objectName string) (*types.Transaction, error) {
	return _IStorage.Contract.CancelCreateObject(&_IStorage.TransactOpts, bucketName, objectName)
}

// CancelMigrateBucket is a paid mutator transaction binding the contract method 0xb18ee371.
//
// Solidity: function cancelMigrateBucket(string bucketName) returns(bool success)
func (_IStorage *IStorageTransactor) CancelMigrateBucket(opts *bind.TransactOpts, bucketName string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "cancelMigrateBucket", bucketName)
}

// CancelMigrateBucket is a paid mutator transaction binding the contract method 0xb18ee371.
//
// Solidity: function cancelMigrateBucket(string bucketName) returns(bool success)
func (_IStorage *IStorageSession) CancelMigrateBucket(bucketName string) (*types.Transaction, error) {
	return _IStorage.Contract.CancelMigrateBucket(&_IStorage.TransactOpts, bucketName)
}

// CancelMigrateBucket is a paid mutator transaction binding the contract method 0xb18ee371.
//
// Solidity: function cancelMigrateBucket(string bucketName) returns(bool success)
func (_IStorage *IStorageTransactorSession) CancelMigrateBucket(bucketName string) (*types.Transaction, error) {
	return _IStorage.Contract.CancelMigrateBucket(&_IStorage.TransactOpts, bucketName)
}

// CompleteMigrateBucket is a paid mutator transaction binding the contract method 0x98a7da16.
//
// Solidity: function completeMigrateBucket(string bucketName, uint32 gvgFamilyId, (uint32,uint32,bytes)[] gvgMappings) returns(bool success)
func (_IStorage *IStorageTransactor) CompleteMigrateBucket(opts *bind.TransactOpts, bucketName string, gvgFamilyId uint32, gvgMappings []GVGMapping) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "completeMigrateBucket", bucketName, gvgFamilyId, gvgMappings)
}

// CompleteMigrateBucket is a paid mutator transaction binding the contract method 0x98a7da16.
//
// Solidity: function completeMigrateBucket(string bucketName, uint32 gvgFamilyId, (uint32,uint32,bytes)[] gvgMappings) returns(bool success)
func (_IStorage *IStorageSession) CompleteMigrateBucket(bucketName string, gvgFamilyId uint32, gvgMappings []GVGMapping) (*types.Transaction, error) {
	return _IStorage.Contract.CompleteMigrateBucket(&_IStorage.TransactOpts, bucketName, gvgFamilyId, gvgMappings)
}

// CompleteMigrateBucket is a paid mutator transaction binding the contract method 0x98a7da16.
//
// Solidity: function completeMigrateBucket(string bucketName, uint32 gvgFamilyId, (uint32,uint32,bytes)[] gvgMappings) returns(bool success)
func (_IStorage *IStorageTransactorSession) CompleteMigrateBucket(bucketName string, gvgFamilyId uint32, gvgMappings []GVGMapping) (*types.Transaction, error) {
	return _IStorage.Contract.CompleteMigrateBucket(&_IStorage.TransactOpts, bucketName, gvgFamilyId, gvgMappings)
}

// CopyObject is a paid mutator transaction binding the contract method 0x6cbe1c0b.
//
// Solidity: function copyObject(string srcBucketName, string dstBucketName, string srcObjectName, string dstObjectName, (uint64,uint32,bytes) dstPrimarySpApproval) returns(bool success)
func (_IStorage *IStorageTransactor) CopyObject(opts *bind.TransactOpts, srcBucketName string, dstBucketName string, srcObjectName string, dstObjectName string, dstPrimarySpApproval Approval) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "copyObject", srcBucketName, dstBucketName, srcObjectName, dstObjectName, dstPrimarySpApproval)
}

// CopyObject is a paid mutator transaction binding the contract method 0x6cbe1c0b.
//
// Solidity: function copyObject(string srcBucketName, string dstBucketName, string srcObjectName, string dstObjectName, (uint64,uint32,bytes) dstPrimarySpApproval) returns(bool success)
func (_IStorage *IStorageSession) CopyObject(srcBucketName string, dstBucketName string, srcObjectName string, dstObjectName string, dstPrimarySpApproval Approval) (*types.Transaction, error) {
	return _IStorage.Contract.CopyObject(&_IStorage.TransactOpts, srcBucketName, dstBucketName, srcObjectName, dstObjectName, dstPrimarySpApproval)
}

// CopyObject is a paid mutator transaction binding the contract method 0x6cbe1c0b.
//
// Solidity: function copyObject(string srcBucketName, string dstBucketName, string srcObjectName, string dstObjectName, (uint64,uint32,bytes) dstPrimarySpApproval) returns(bool success)
func (_IStorage *IStorageTransactorSession) CopyObject(srcBucketName string, dstBucketName string, srcObjectName string, dstObjectName string, dstPrimarySpApproval Approval) (*types.Transaction, error) {
	return _IStorage.Contract.CopyObject(&_IStorage.TransactOpts, srcBucketName, dstBucketName, srcObjectName, dstObjectName, dstPrimarySpApproval)
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

// DelegateCreateObject is a paid mutator transaction binding the contract method 0xc38708a0.
//
// Solidity: function delegateCreateObject(string creator, string bucketName, string objectName, uint64 payloadSize, string contentType, uint8 visibility, string[] expectChecksums, uint8 redundancyType) returns(bool success)
func (_IStorage *IStorageTransactor) DelegateCreateObject(opts *bind.TransactOpts, creator string, bucketName string, objectName string, payloadSize uint64, contentType string, visibility uint8, expectChecksums []string, redundancyType uint8) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "delegateCreateObject", creator, bucketName, objectName, payloadSize, contentType, visibility, expectChecksums, redundancyType)
}

// DelegateCreateObject is a paid mutator transaction binding the contract method 0xc38708a0.
//
// Solidity: function delegateCreateObject(string creator, string bucketName, string objectName, uint64 payloadSize, string contentType, uint8 visibility, string[] expectChecksums, uint8 redundancyType) returns(bool success)
func (_IStorage *IStorageSession) DelegateCreateObject(creator string, bucketName string, objectName string, payloadSize uint64, contentType string, visibility uint8, expectChecksums []string, redundancyType uint8) (*types.Transaction, error) {
	return _IStorage.Contract.DelegateCreateObject(&_IStorage.TransactOpts, creator, bucketName, objectName, payloadSize, contentType, visibility, expectChecksums, redundancyType)
}

// DelegateCreateObject is a paid mutator transaction binding the contract method 0xc38708a0.
//
// Solidity: function delegateCreateObject(string creator, string bucketName, string objectName, uint64 payloadSize, string contentType, uint8 visibility, string[] expectChecksums, uint8 redundancyType) returns(bool success)
func (_IStorage *IStorageTransactorSession) DelegateCreateObject(creator string, bucketName string, objectName string, payloadSize uint64, contentType string, visibility uint8, expectChecksums []string, redundancyType uint8) (*types.Transaction, error) {
	return _IStorage.Contract.DelegateCreateObject(&_IStorage.TransactOpts, creator, bucketName, objectName, payloadSize, contentType, visibility, expectChecksums, redundancyType)
}

// DelegateUpdateObjectContent is a paid mutator transaction binding the contract method 0xe5271439.
//
// Solidity: function delegateUpdateObjectContent(string updater, string bucketName, string objectName, uint64 payloadSize, string contentType, string[] expectChecksums) returns(bool success)
func (_IStorage *IStorageTransactor) DelegateUpdateObjectContent(opts *bind.TransactOpts, updater string, bucketName string, objectName string, payloadSize uint64, contentType string, expectChecksums []string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "delegateUpdateObjectContent", updater, bucketName, objectName, payloadSize, contentType, expectChecksums)
}

// DelegateUpdateObjectContent is a paid mutator transaction binding the contract method 0xe5271439.
//
// Solidity: function delegateUpdateObjectContent(string updater, string bucketName, string objectName, uint64 payloadSize, string contentType, string[] expectChecksums) returns(bool success)
func (_IStorage *IStorageSession) DelegateUpdateObjectContent(updater string, bucketName string, objectName string, payloadSize uint64, contentType string, expectChecksums []string) (*types.Transaction, error) {
	return _IStorage.Contract.DelegateUpdateObjectContent(&_IStorage.TransactOpts, updater, bucketName, objectName, payloadSize, contentType, expectChecksums)
}

// DelegateUpdateObjectContent is a paid mutator transaction binding the contract method 0xe5271439.
//
// Solidity: function delegateUpdateObjectContent(string updater, string bucketName, string objectName, uint64 payloadSize, string contentType, string[] expectChecksums) returns(bool success)
func (_IStorage *IStorageTransactorSession) DelegateUpdateObjectContent(updater string, bucketName string, objectName string, payloadSize uint64, contentType string, expectChecksums []string) (*types.Transaction, error) {
	return _IStorage.Contract.DelegateUpdateObjectContent(&_IStorage.TransactOpts, updater, bucketName, objectName, payloadSize, contentType, expectChecksums)
}

// DeleteBucket is a paid mutator transaction binding the contract method 0x0feea0b3.
//
// Solidity: function deleteBucket(string bucketName) returns(bool success)
func (_IStorage *IStorageTransactor) DeleteBucket(opts *bind.TransactOpts, bucketName string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "deleteBucket", bucketName)
}

// DeleteBucket is a paid mutator transaction binding the contract method 0x0feea0b3.
//
// Solidity: function deleteBucket(string bucketName) returns(bool success)
func (_IStorage *IStorageSession) DeleteBucket(bucketName string) (*types.Transaction, error) {
	return _IStorage.Contract.DeleteBucket(&_IStorage.TransactOpts, bucketName)
}

// DeleteBucket is a paid mutator transaction binding the contract method 0x0feea0b3.
//
// Solidity: function deleteBucket(string bucketName) returns(bool success)
func (_IStorage *IStorageTransactorSession) DeleteBucket(bucketName string) (*types.Transaction, error) {
	return _IStorage.Contract.DeleteBucket(&_IStorage.TransactOpts, bucketName)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x2e8b92a9.
//
// Solidity: function deleteGroup(string groupName) returns(bool success)
func (_IStorage *IStorageTransactor) DeleteGroup(opts *bind.TransactOpts, groupName string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "deleteGroup", groupName)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x2e8b92a9.
//
// Solidity: function deleteGroup(string groupName) returns(bool success)
func (_IStorage *IStorageSession) DeleteGroup(groupName string) (*types.Transaction, error) {
	return _IStorage.Contract.DeleteGroup(&_IStorage.TransactOpts, groupName)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x2e8b92a9.
//
// Solidity: function deleteGroup(string groupName) returns(bool success)
func (_IStorage *IStorageTransactorSession) DeleteGroup(groupName string) (*types.Transaction, error) {
	return _IStorage.Contract.DeleteGroup(&_IStorage.TransactOpts, groupName)
}

// DeleteObject is a paid mutator transaction binding the contract method 0x229336a3.
//
// Solidity: function deleteObject(string bucketName, string objectName) returns(bool success)
func (_IStorage *IStorageTransactor) DeleteObject(opts *bind.TransactOpts, bucketName string, objectName string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "deleteObject", bucketName, objectName)
}

// DeleteObject is a paid mutator transaction binding the contract method 0x229336a3.
//
// Solidity: function deleteObject(string bucketName, string objectName) returns(bool success)
func (_IStorage *IStorageSession) DeleteObject(bucketName string, objectName string) (*types.Transaction, error) {
	return _IStorage.Contract.DeleteObject(&_IStorage.TransactOpts, bucketName, objectName)
}

// DeleteObject is a paid mutator transaction binding the contract method 0x229336a3.
//
// Solidity: function deleteObject(string bucketName, string objectName) returns(bool success)
func (_IStorage *IStorageTransactorSession) DeleteObject(bucketName string, objectName string) (*types.Transaction, error) {
	return _IStorage.Contract.DeleteObject(&_IStorage.TransactOpts, bucketName, objectName)
}

// DeletePolicy is a paid mutator transaction binding the contract method 0x7ecda60e.
//
// Solidity: function deletePolicy((int32,string) principal, string resource) returns(bool success)
func (_IStorage *IStorageTransactor) DeletePolicy(opts *bind.TransactOpts, principal Principal, resource string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "deletePolicy", principal, resource)
}

// DeletePolicy is a paid mutator transaction binding the contract method 0x7ecda60e.
//
// Solidity: function deletePolicy((int32,string) principal, string resource) returns(bool success)
func (_IStorage *IStorageSession) DeletePolicy(principal Principal, resource string) (*types.Transaction, error) {
	return _IStorage.Contract.DeletePolicy(&_IStorage.TransactOpts, principal, resource)
}

// DeletePolicy is a paid mutator transaction binding the contract method 0x7ecda60e.
//
// Solidity: function deletePolicy((int32,string) principal, string resource) returns(bool success)
func (_IStorage *IStorageTransactorSession) DeletePolicy(principal Principal, resource string) (*types.Transaction, error) {
	return _IStorage.Contract.DeletePolicy(&_IStorage.TransactOpts, principal, resource)
}

// DiscontinueBucket is a paid mutator transaction binding the contract method 0x1c411577.
//
// Solidity: function discontinueBucket(string bucketName, string reason) returns(bool success)
func (_IStorage *IStorageTransactor) DiscontinueBucket(opts *bind.TransactOpts, bucketName string, reason string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "discontinueBucket", bucketName, reason)
}

// DiscontinueBucket is a paid mutator transaction binding the contract method 0x1c411577.
//
// Solidity: function discontinueBucket(string bucketName, string reason) returns(bool success)
func (_IStorage *IStorageSession) DiscontinueBucket(bucketName string, reason string) (*types.Transaction, error) {
	return _IStorage.Contract.DiscontinueBucket(&_IStorage.TransactOpts, bucketName, reason)
}

// DiscontinueBucket is a paid mutator transaction binding the contract method 0x1c411577.
//
// Solidity: function discontinueBucket(string bucketName, string reason) returns(bool success)
func (_IStorage *IStorageTransactorSession) DiscontinueBucket(bucketName string, reason string) (*types.Transaction, error) {
	return _IStorage.Contract.DiscontinueBucket(&_IStorage.TransactOpts, bucketName, reason)
}

// DiscontinueObject is a paid mutator transaction binding the contract method 0xaffede78.
//
// Solidity: function discontinueObject(string bucketName, uint256[] objectIds, string reason) returns(bool success)
func (_IStorage *IStorageTransactor) DiscontinueObject(opts *bind.TransactOpts, bucketName string, objectIds []*big.Int, reason string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "discontinueObject", bucketName, objectIds, reason)
}

// DiscontinueObject is a paid mutator transaction binding the contract method 0xaffede78.
//
// Solidity: function discontinueObject(string bucketName, uint256[] objectIds, string reason) returns(bool success)
func (_IStorage *IStorageSession) DiscontinueObject(bucketName string, objectIds []*big.Int, reason string) (*types.Transaction, error) {
	return _IStorage.Contract.DiscontinueObject(&_IStorage.TransactOpts, bucketName, objectIds, reason)
}

// DiscontinueObject is a paid mutator transaction binding the contract method 0xaffede78.
//
// Solidity: function discontinueObject(string bucketName, uint256[] objectIds, string reason) returns(bool success)
func (_IStorage *IStorageTransactorSession) DiscontinueObject(bucketName string, objectIds []*big.Int, reason string) (*types.Transaction, error) {
	return _IStorage.Contract.DiscontinueObject(&_IStorage.TransactOpts, bucketName, objectIds, reason)
}

// LeaveGroup is a paid mutator transaction binding the contract method 0xf61a36f1.
//
// Solidity: function leaveGroup(address member, address groupOwner, string groupName) returns(bool success)
func (_IStorage *IStorageTransactor) LeaveGroup(opts *bind.TransactOpts, member common.Address, groupOwner common.Address, groupName string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "leaveGroup", member, groupOwner, groupName)
}

// LeaveGroup is a paid mutator transaction binding the contract method 0xf61a36f1.
//
// Solidity: function leaveGroup(address member, address groupOwner, string groupName) returns(bool success)
func (_IStorage *IStorageSession) LeaveGroup(member common.Address, groupOwner common.Address, groupName string) (*types.Transaction, error) {
	return _IStorage.Contract.LeaveGroup(&_IStorage.TransactOpts, member, groupOwner, groupName)
}

// LeaveGroup is a paid mutator transaction binding the contract method 0xf61a36f1.
//
// Solidity: function leaveGroup(address member, address groupOwner, string groupName) returns(bool success)
func (_IStorage *IStorageTransactorSession) LeaveGroup(member common.Address, groupOwner common.Address, groupName string) (*types.Transaction, error) {
	return _IStorage.Contract.LeaveGroup(&_IStorage.TransactOpts, member, groupOwner, groupName)
}

// MigrateBucket is a paid mutator transaction binding the contract method 0x4ac6f8f8.
//
// Solidity: function migrateBucket(string bucketName, uint32 dstPrimarySpId, (uint64,uint32,bytes) dstPrimarySpApproval) returns(bool success)
func (_IStorage *IStorageTransactor) MigrateBucket(opts *bind.TransactOpts, bucketName string, dstPrimarySpId uint32, dstPrimarySpApproval Approval) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "migrateBucket", bucketName, dstPrimarySpId, dstPrimarySpApproval)
}

// MigrateBucket is a paid mutator transaction binding the contract method 0x4ac6f8f8.
//
// Solidity: function migrateBucket(string bucketName, uint32 dstPrimarySpId, (uint64,uint32,bytes) dstPrimarySpApproval) returns(bool success)
func (_IStorage *IStorageSession) MigrateBucket(bucketName string, dstPrimarySpId uint32, dstPrimarySpApproval Approval) (*types.Transaction, error) {
	return _IStorage.Contract.MigrateBucket(&_IStorage.TransactOpts, bucketName, dstPrimarySpId, dstPrimarySpApproval)
}

// MigrateBucket is a paid mutator transaction binding the contract method 0x4ac6f8f8.
//
// Solidity: function migrateBucket(string bucketName, uint32 dstPrimarySpId, (uint64,uint32,bytes) dstPrimarySpApproval) returns(bool success)
func (_IStorage *IStorageTransactorSession) MigrateBucket(bucketName string, dstPrimarySpId uint32, dstPrimarySpApproval Approval) (*types.Transaction, error) {
	return _IStorage.Contract.MigrateBucket(&_IStorage.TransactOpts, bucketName, dstPrimarySpId, dstPrimarySpApproval)
}

// MirrorBucket is a paid mutator transaction binding the contract method 0x5da51665.
//
// Solidity: function mirrorBucket(uint256 bucketId, string bucketName, uint32 destChainId) returns(bool success)
func (_IStorage *IStorageTransactor) MirrorBucket(opts *bind.TransactOpts, bucketId *big.Int, bucketName string, destChainId uint32) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "mirrorBucket", bucketId, bucketName, destChainId)
}

// MirrorBucket is a paid mutator transaction binding the contract method 0x5da51665.
//
// Solidity: function mirrorBucket(uint256 bucketId, string bucketName, uint32 destChainId) returns(bool success)
func (_IStorage *IStorageSession) MirrorBucket(bucketId *big.Int, bucketName string, destChainId uint32) (*types.Transaction, error) {
	return _IStorage.Contract.MirrorBucket(&_IStorage.TransactOpts, bucketId, bucketName, destChainId)
}

// MirrorBucket is a paid mutator transaction binding the contract method 0x5da51665.
//
// Solidity: function mirrorBucket(uint256 bucketId, string bucketName, uint32 destChainId) returns(bool success)
func (_IStorage *IStorageTransactorSession) MirrorBucket(bucketId *big.Int, bucketName string, destChainId uint32) (*types.Transaction, error) {
	return _IStorage.Contract.MirrorBucket(&_IStorage.TransactOpts, bucketId, bucketName, destChainId)
}

// MirrorGroup is a paid mutator transaction binding the contract method 0x7f6f3ace.
//
// Solidity: function mirrorGroup(uint256 groupId, string groupName, uint32 destChainId) returns(bool success)
func (_IStorage *IStorageTransactor) MirrorGroup(opts *bind.TransactOpts, groupId *big.Int, groupName string, destChainId uint32) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "mirrorGroup", groupId, groupName, destChainId)
}

// MirrorGroup is a paid mutator transaction binding the contract method 0x7f6f3ace.
//
// Solidity: function mirrorGroup(uint256 groupId, string groupName, uint32 destChainId) returns(bool success)
func (_IStorage *IStorageSession) MirrorGroup(groupId *big.Int, groupName string, destChainId uint32) (*types.Transaction, error) {
	return _IStorage.Contract.MirrorGroup(&_IStorage.TransactOpts, groupId, groupName, destChainId)
}

// MirrorGroup is a paid mutator transaction binding the contract method 0x7f6f3ace.
//
// Solidity: function mirrorGroup(uint256 groupId, string groupName, uint32 destChainId) returns(bool success)
func (_IStorage *IStorageTransactorSession) MirrorGroup(groupId *big.Int, groupName string, destChainId uint32) (*types.Transaction, error) {
	return _IStorage.Contract.MirrorGroup(&_IStorage.TransactOpts, groupId, groupName, destChainId)
}

// MirrorObject is a paid mutator transaction binding the contract method 0x3e801b0c.
//
// Solidity: function mirrorObject(uint256 objectId, string bucketName, string objectName, uint32 destChainId) returns(bool success)
func (_IStorage *IStorageTransactor) MirrorObject(opts *bind.TransactOpts, objectId *big.Int, bucketName string, objectName string, destChainId uint32) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "mirrorObject", objectId, bucketName, objectName, destChainId)
}

// MirrorObject is a paid mutator transaction binding the contract method 0x3e801b0c.
//
// Solidity: function mirrorObject(uint256 objectId, string bucketName, string objectName, uint32 destChainId) returns(bool success)
func (_IStorage *IStorageSession) MirrorObject(objectId *big.Int, bucketName string, objectName string, destChainId uint32) (*types.Transaction, error) {
	return _IStorage.Contract.MirrorObject(&_IStorage.TransactOpts, objectId, bucketName, objectName, destChainId)
}

// MirrorObject is a paid mutator transaction binding the contract method 0x3e801b0c.
//
// Solidity: function mirrorObject(uint256 objectId, string bucketName, string objectName, uint32 destChainId) returns(bool success)
func (_IStorage *IStorageTransactorSession) MirrorObject(objectId *big.Int, bucketName string, objectName string, destChainId uint32) (*types.Transaction, error) {
	return _IStorage.Contract.MirrorObject(&_IStorage.TransactOpts, objectId, bucketName, objectName, destChainId)
}

// PutPolicy is a paid mutator transaction binding the contract method 0x49052a67.
//
// Solidity: function putPolicy((int32,string) principal, string resource, (int32,int32[],string[],int64,uint64)[] statements, int64 expirationTime) returns(bool success)
func (_IStorage *IStorageTransactor) PutPolicy(opts *bind.TransactOpts, principal Principal, resource string, statements []Statement, expirationTime int64) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "putPolicy", principal, resource, statements, expirationTime)
}

// PutPolicy is a paid mutator transaction binding the contract method 0x49052a67.
//
// Solidity: function putPolicy((int32,string) principal, string resource, (int32,int32[],string[],int64,uint64)[] statements, int64 expirationTime) returns(bool success)
func (_IStorage *IStorageSession) PutPolicy(principal Principal, resource string, statements []Statement, expirationTime int64) (*types.Transaction, error) {
	return _IStorage.Contract.PutPolicy(&_IStorage.TransactOpts, principal, resource, statements, expirationTime)
}

// PutPolicy is a paid mutator transaction binding the contract method 0x49052a67.
//
// Solidity: function putPolicy((int32,string) principal, string resource, (int32,int32[],string[],int64,uint64)[] statements, int64 expirationTime) returns(bool success)
func (_IStorage *IStorageTransactorSession) PutPolicy(principal Principal, resource string, statements []Statement, expirationTime int64) (*types.Transaction, error) {
	return _IStorage.Contract.PutPolicy(&_IStorage.TransactOpts, principal, resource, statements, expirationTime)
}

// RejectMigrateBucket is a paid mutator transaction binding the contract method 0x7c216b07.
//
// Solidity: function rejectMigrateBucket(string bucketName) returns(bool success)
func (_IStorage *IStorageTransactor) RejectMigrateBucket(opts *bind.TransactOpts, bucketName string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "rejectMigrateBucket", bucketName)
}

// RejectMigrateBucket is a paid mutator transaction binding the contract method 0x7c216b07.
//
// Solidity: function rejectMigrateBucket(string bucketName) returns(bool success)
func (_IStorage *IStorageSession) RejectMigrateBucket(bucketName string) (*types.Transaction, error) {
	return _IStorage.Contract.RejectMigrateBucket(&_IStorage.TransactOpts, bucketName)
}

// RejectMigrateBucket is a paid mutator transaction binding the contract method 0x7c216b07.
//
// Solidity: function rejectMigrateBucket(string bucketName) returns(bool success)
func (_IStorage *IStorageTransactorSession) RejectMigrateBucket(bucketName string) (*types.Transaction, error) {
	return _IStorage.Contract.RejectMigrateBucket(&_IStorage.TransactOpts, bucketName)
}

// RejectSealObject is a paid mutator transaction binding the contract method 0xc07f194e.
//
// Solidity: function rejectSealObject(string bucketName, string objectName) returns(bool success)
func (_IStorage *IStorageTransactor) RejectSealObject(opts *bind.TransactOpts, bucketName string, objectName string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "rejectSealObject", bucketName, objectName)
}

// RejectSealObject is a paid mutator transaction binding the contract method 0xc07f194e.
//
// Solidity: function rejectSealObject(string bucketName, string objectName) returns(bool success)
func (_IStorage *IStorageSession) RejectSealObject(bucketName string, objectName string) (*types.Transaction, error) {
	return _IStorage.Contract.RejectSealObject(&_IStorage.TransactOpts, bucketName, objectName)
}

// RejectSealObject is a paid mutator transaction binding the contract method 0xc07f194e.
//
// Solidity: function rejectSealObject(string bucketName, string objectName) returns(bool success)
func (_IStorage *IStorageTransactorSession) RejectSealObject(bucketName string, objectName string) (*types.Transaction, error) {
	return _IStorage.Contract.RejectSealObject(&_IStorage.TransactOpts, bucketName, objectName)
}

// RenewGroupMember is a paid mutator transaction binding the contract method 0x3a248669.
//
// Solidity: function renewGroupMember(address groupOwner, string groupName, address[] members, int64[] expirationTime) returns(bool success)
func (_IStorage *IStorageTransactor) RenewGroupMember(opts *bind.TransactOpts, groupOwner common.Address, groupName string, members []common.Address, expirationTime []int64) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "renewGroupMember", groupOwner, groupName, members, expirationTime)
}

// RenewGroupMember is a paid mutator transaction binding the contract method 0x3a248669.
//
// Solidity: function renewGroupMember(address groupOwner, string groupName, address[] members, int64[] expirationTime) returns(bool success)
func (_IStorage *IStorageSession) RenewGroupMember(groupOwner common.Address, groupName string, members []common.Address, expirationTime []int64) (*types.Transaction, error) {
	return _IStorage.Contract.RenewGroupMember(&_IStorage.TransactOpts, groupOwner, groupName, members, expirationTime)
}

// RenewGroupMember is a paid mutator transaction binding the contract method 0x3a248669.
//
// Solidity: function renewGroupMember(address groupOwner, string groupName, address[] members, int64[] expirationTime) returns(bool success)
func (_IStorage *IStorageTransactorSession) RenewGroupMember(groupOwner common.Address, groupName string, members []common.Address, expirationTime []int64) (*types.Transaction, error) {
	return _IStorage.Contract.RenewGroupMember(&_IStorage.TransactOpts, groupOwner, groupName, members, expirationTime)
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

// SetBucketFlowRateLimit is a paid mutator transaction binding the contract method 0xf9c523d0.
//
// Solidity: function setBucketFlowRateLimit(string bucketName, string bucketOwner, string paymentAddress, uint256 flowRateLimit) returns(bool success)
func (_IStorage *IStorageTransactor) SetBucketFlowRateLimit(opts *bind.TransactOpts, bucketName string, bucketOwner string, paymentAddress string, flowRateLimit *big.Int) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "setBucketFlowRateLimit", bucketName, bucketOwner, paymentAddress, flowRateLimit)
}

// SetBucketFlowRateLimit is a paid mutator transaction binding the contract method 0xf9c523d0.
//
// Solidity: function setBucketFlowRateLimit(string bucketName, string bucketOwner, string paymentAddress, uint256 flowRateLimit) returns(bool success)
func (_IStorage *IStorageSession) SetBucketFlowRateLimit(bucketName string, bucketOwner string, paymentAddress string, flowRateLimit *big.Int) (*types.Transaction, error) {
	return _IStorage.Contract.SetBucketFlowRateLimit(&_IStorage.TransactOpts, bucketName, bucketOwner, paymentAddress, flowRateLimit)
}

// SetBucketFlowRateLimit is a paid mutator transaction binding the contract method 0xf9c523d0.
//
// Solidity: function setBucketFlowRateLimit(string bucketName, string bucketOwner, string paymentAddress, uint256 flowRateLimit) returns(bool success)
func (_IStorage *IStorageTransactorSession) SetBucketFlowRateLimit(bucketName string, bucketOwner string, paymentAddress string, flowRateLimit *big.Int) (*types.Transaction, error) {
	return _IStorage.Contract.SetBucketFlowRateLimit(&_IStorage.TransactOpts, bucketName, bucketOwner, paymentAddress, flowRateLimit)
}

// SetTag is a paid mutator transaction binding the contract method 0x5d0ccd6e.
//
// Solidity: function setTag(string resource, (string,string)[] tags) returns(bool success)
func (_IStorage *IStorageTransactor) SetTag(opts *bind.TransactOpts, resource string, tags []Tag) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "setTag", resource, tags)
}

// SetTag is a paid mutator transaction binding the contract method 0x5d0ccd6e.
//
// Solidity: function setTag(string resource, (string,string)[] tags) returns(bool success)
func (_IStorage *IStorageSession) SetTag(resource string, tags []Tag) (*types.Transaction, error) {
	return _IStorage.Contract.SetTag(&_IStorage.TransactOpts, resource, tags)
}

// SetTag is a paid mutator transaction binding the contract method 0x5d0ccd6e.
//
// Solidity: function setTag(string resource, (string,string)[] tags) returns(bool success)
func (_IStorage *IStorageTransactorSession) SetTag(resource string, tags []Tag) (*types.Transaction, error) {
	return _IStorage.Contract.SetTag(&_IStorage.TransactOpts, resource, tags)
}

// ToggleSPAsDelegatedAgent is a paid mutator transaction binding the contract method 0xcd4bd67a.
//
// Solidity: function toggleSPAsDelegatedAgent(string bucketName) returns(bool success)
func (_IStorage *IStorageTransactor) ToggleSPAsDelegatedAgent(opts *bind.TransactOpts, bucketName string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "toggleSPAsDelegatedAgent", bucketName)
}

// ToggleSPAsDelegatedAgent is a paid mutator transaction binding the contract method 0xcd4bd67a.
//
// Solidity: function toggleSPAsDelegatedAgent(string bucketName) returns(bool success)
func (_IStorage *IStorageSession) ToggleSPAsDelegatedAgent(bucketName string) (*types.Transaction, error) {
	return _IStorage.Contract.ToggleSPAsDelegatedAgent(&_IStorage.TransactOpts, bucketName)
}

// ToggleSPAsDelegatedAgent is a paid mutator transaction binding the contract method 0xcd4bd67a.
//
// Solidity: function toggleSPAsDelegatedAgent(string bucketName) returns(bool success)
func (_IStorage *IStorageTransactorSession) ToggleSPAsDelegatedAgent(bucketName string) (*types.Transaction, error) {
	return _IStorage.Contract.ToggleSPAsDelegatedAgent(&_IStorage.TransactOpts, bucketName)
}

// UpdateBucketInfo is a paid mutator transaction binding the contract method 0x7a5c5d77.
//
// Solidity: function updateBucketInfo(string bucketName, uint8 visibility, address paymentAddress, int128 chargedReadQuota) returns(bool success)
func (_IStorage *IStorageTransactor) UpdateBucketInfo(opts *bind.TransactOpts, bucketName string, visibility uint8, paymentAddress common.Address, chargedReadQuota *big.Int) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "updateBucketInfo", bucketName, visibility, paymentAddress, chargedReadQuota)
}

// UpdateBucketInfo is a paid mutator transaction binding the contract method 0x7a5c5d77.
//
// Solidity: function updateBucketInfo(string bucketName, uint8 visibility, address paymentAddress, int128 chargedReadQuota) returns(bool success)
func (_IStorage *IStorageSession) UpdateBucketInfo(bucketName string, visibility uint8, paymentAddress common.Address, chargedReadQuota *big.Int) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateBucketInfo(&_IStorage.TransactOpts, bucketName, visibility, paymentAddress, chargedReadQuota)
}

// UpdateBucketInfo is a paid mutator transaction binding the contract method 0x7a5c5d77.
//
// Solidity: function updateBucketInfo(string bucketName, uint8 visibility, address paymentAddress, int128 chargedReadQuota) returns(bool success)
func (_IStorage *IStorageTransactorSession) UpdateBucketInfo(bucketName string, visibility uint8, paymentAddress common.Address, chargedReadQuota *big.Int) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateBucketInfo(&_IStorage.TransactOpts, bucketName, visibility, paymentAddress, chargedReadQuota)
}

// UpdateGroup is a paid mutator transaction binding the contract method 0xc966a7cc.
//
// Solidity: function updateGroup(address groupOwner, string groupName, address[] membersToAdd, int64[] expirationTime, address[] membersToDelete) returns(bool success)
func (_IStorage *IStorageTransactor) UpdateGroup(opts *bind.TransactOpts, groupOwner common.Address, groupName string, membersToAdd []common.Address, expirationTime []int64, membersToDelete []common.Address) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "updateGroup", groupOwner, groupName, membersToAdd, expirationTime, membersToDelete)
}

// UpdateGroup is a paid mutator transaction binding the contract method 0xc966a7cc.
//
// Solidity: function updateGroup(address groupOwner, string groupName, address[] membersToAdd, int64[] expirationTime, address[] membersToDelete) returns(bool success)
func (_IStorage *IStorageSession) UpdateGroup(groupOwner common.Address, groupName string, membersToAdd []common.Address, expirationTime []int64, membersToDelete []common.Address) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateGroup(&_IStorage.TransactOpts, groupOwner, groupName, membersToAdd, expirationTime, membersToDelete)
}

// UpdateGroup is a paid mutator transaction binding the contract method 0xc966a7cc.
//
// Solidity: function updateGroup(address groupOwner, string groupName, address[] membersToAdd, int64[] expirationTime, address[] membersToDelete) returns(bool success)
func (_IStorage *IStorageTransactorSession) UpdateGroup(groupOwner common.Address, groupName string, membersToAdd []common.Address, expirationTime []int64, membersToDelete []common.Address) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateGroup(&_IStorage.TransactOpts, groupOwner, groupName, membersToAdd, expirationTime, membersToDelete)
}

// UpdateGroupExtra is a paid mutator transaction binding the contract method 0x4b047704.
//
// Solidity: function updateGroupExtra(address groupOwner, string groupName, string extra) returns(bool success)
func (_IStorage *IStorageTransactor) UpdateGroupExtra(opts *bind.TransactOpts, groupOwner common.Address, groupName string, extra string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "updateGroupExtra", groupOwner, groupName, extra)
}

// UpdateGroupExtra is a paid mutator transaction binding the contract method 0x4b047704.
//
// Solidity: function updateGroupExtra(address groupOwner, string groupName, string extra) returns(bool success)
func (_IStorage *IStorageSession) UpdateGroupExtra(groupOwner common.Address, groupName string, extra string) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateGroupExtra(&_IStorage.TransactOpts, groupOwner, groupName, extra)
}

// UpdateGroupExtra is a paid mutator transaction binding the contract method 0x4b047704.
//
// Solidity: function updateGroupExtra(address groupOwner, string groupName, string extra) returns(bool success)
func (_IStorage *IStorageTransactorSession) UpdateGroupExtra(groupOwner common.Address, groupName string, extra string) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateGroupExtra(&_IStorage.TransactOpts, groupOwner, groupName, extra)
}

// UpdateObjectContent is a paid mutator transaction binding the contract method 0xb303d5ac.
//
// Solidity: function updateObjectContent(string bucketName, string objectName, uint64 payloadSize, string contentType, string[] expectChecksums) returns(bool success)
func (_IStorage *IStorageTransactor) UpdateObjectContent(opts *bind.TransactOpts, bucketName string, objectName string, payloadSize uint64, contentType string, expectChecksums []string) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "updateObjectContent", bucketName, objectName, payloadSize, contentType, expectChecksums)
}

// UpdateObjectContent is a paid mutator transaction binding the contract method 0xb303d5ac.
//
// Solidity: function updateObjectContent(string bucketName, string objectName, uint64 payloadSize, string contentType, string[] expectChecksums) returns(bool success)
func (_IStorage *IStorageSession) UpdateObjectContent(bucketName string, objectName string, payloadSize uint64, contentType string, expectChecksums []string) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateObjectContent(&_IStorage.TransactOpts, bucketName, objectName, payloadSize, contentType, expectChecksums)
}

// UpdateObjectContent is a paid mutator transaction binding the contract method 0xb303d5ac.
//
// Solidity: function updateObjectContent(string bucketName, string objectName, uint64 payloadSize, string contentType, string[] expectChecksums) returns(bool success)
func (_IStorage *IStorageTransactorSession) UpdateObjectContent(bucketName string, objectName string, payloadSize uint64, contentType string, expectChecksums []string) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateObjectContent(&_IStorage.TransactOpts, bucketName, objectName, payloadSize, contentType, expectChecksums)
}

// UpdateObjectInfo is a paid mutator transaction binding the contract method 0xf167e687.
//
// Solidity: function updateObjectInfo(string bucketName, string objectName, uint8 visibility) returns(bool success)
func (_IStorage *IStorageTransactor) UpdateObjectInfo(opts *bind.TransactOpts, bucketName string, objectName string, visibility uint8) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "updateObjectInfo", bucketName, objectName, visibility)
}

// UpdateObjectInfo is a paid mutator transaction binding the contract method 0xf167e687.
//
// Solidity: function updateObjectInfo(string bucketName, string objectName, uint8 visibility) returns(bool success)
func (_IStorage *IStorageSession) UpdateObjectInfo(bucketName string, objectName string, visibility uint8) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateObjectInfo(&_IStorage.TransactOpts, bucketName, objectName, visibility)
}

// UpdateObjectInfo is a paid mutator transaction binding the contract method 0xf167e687.
//
// Solidity: function updateObjectInfo(string bucketName, string objectName, uint8 visibility) returns(bool success)
func (_IStorage *IStorageTransactorSession) UpdateObjectInfo(bucketName string, objectName string, visibility uint8) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateObjectInfo(&_IStorage.TransactOpts, bucketName, objectName, visibility)
}

// UpdateParams is a paid mutator transaction binding the contract method 0x90adfb47.
//
// Solidity: function updateParams(string authority, ((uint64,uint32,uint32,uint64),uint64,string,string,string,string,string,string,uint32,uint64,uint64,uint64,int64,uint64,uint64,uint64,uint32,string,string,string,string,string,string,string,string) params) returns(bool success)
func (_IStorage *IStorageTransactor) UpdateParams(opts *bind.TransactOpts, authority string, params Params) (*types.Transaction, error) {
	return _IStorage.contract.Transact(opts, "updateParams", authority, params)
}

// UpdateParams is a paid mutator transaction binding the contract method 0x90adfb47.
//
// Solidity: function updateParams(string authority, ((uint64,uint32,uint32,uint64),uint64,string,string,string,string,string,string,uint32,uint64,uint64,uint64,int64,uint64,uint64,uint64,uint32,string,string,string,string,string,string,string,string) params) returns(bool success)
func (_IStorage *IStorageSession) UpdateParams(authority string, params Params) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateParams(&_IStorage.TransactOpts, authority, params)
}

// UpdateParams is a paid mutator transaction binding the contract method 0x90adfb47.
//
// Solidity: function updateParams(string authority, ((uint64,uint32,uint32,uint64),uint64,string,string,string,string,string,string,uint32,uint64,uint64,uint64,int64,uint64,uint64,uint64,uint32,string,string,string,string,string,string,string,string) params) returns(bool success)
func (_IStorage *IStorageTransactorSession) UpdateParams(authority string, params Params) (*types.Transaction, error) {
	return _IStorage.Contract.UpdateParams(&_IStorage.TransactOpts, authority, params)
}

// IStorageCancelCreateObjectIterator is returned from FilterCancelCreateObject and is used to iterate over the raw logs and unpacked data for CancelCreateObject events raised by the IStorage contract.
type IStorageCancelCreateObjectIterator struct {
	Event *IStorageCancelCreateObject // Event containing the contract specifics and raw log

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
func (it *IStorageCancelCreateObjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageCancelCreateObject)
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
		it.Event = new(IStorageCancelCreateObject)
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
func (it *IStorageCancelCreateObjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageCancelCreateObjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageCancelCreateObject represents a CancelCreateObject event raised by the IStorage contract.
type IStorageCancelCreateObject struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancelCreateObject is a free log retrieval operation binding the contract event 0x462124f21e377bd51ecd6641e3d6708f9881b3118fc600a7251f336befc81cff.
//
// Solidity: event CancelCreateObject(address indexed creator)
func (_IStorage *IStorageFilterer) FilterCancelCreateObject(opts *bind.FilterOpts, creator []common.Address) (*IStorageCancelCreateObjectIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "CancelCreateObject", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageCancelCreateObjectIterator{contract: _IStorage.contract, event: "CancelCreateObject", logs: logs, sub: sub}, nil
}

// WatchCancelCreateObject is a free log subscription operation binding the contract event 0x462124f21e377bd51ecd6641e3d6708f9881b3118fc600a7251f336befc81cff.
//
// Solidity: event CancelCreateObject(address indexed creator)
func (_IStorage *IStorageFilterer) WatchCancelCreateObject(opts *bind.WatchOpts, sink chan<- *IStorageCancelCreateObject, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "CancelCreateObject", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageCancelCreateObject)
				if err := _IStorage.contract.UnpackLog(event, "CancelCreateObject", log); err != nil {
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

// ParseCancelCreateObject is a log parse operation binding the contract event 0x462124f21e377bd51ecd6641e3d6708f9881b3118fc600a7251f336befc81cff.
//
// Solidity: event CancelCreateObject(address indexed creator)
func (_IStorage *IStorageFilterer) ParseCancelCreateObject(log types.Log) (*IStorageCancelCreateObject, error) {
	event := new(IStorageCancelCreateObject)
	if err := _IStorage.contract.UnpackLog(event, "CancelCreateObject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageCancelMigrateBucketIterator is returned from FilterCancelMigrateBucket and is used to iterate over the raw logs and unpacked data for CancelMigrateBucket events raised by the IStorage contract.
type IStorageCancelMigrateBucketIterator struct {
	Event *IStorageCancelMigrateBucket // Event containing the contract specifics and raw log

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
func (it *IStorageCancelMigrateBucketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageCancelMigrateBucket)
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
		it.Event = new(IStorageCancelMigrateBucket)
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
func (it *IStorageCancelMigrateBucketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageCancelMigrateBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageCancelMigrateBucket represents a CancelMigrateBucket event raised by the IStorage contract.
type IStorageCancelMigrateBucket struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCancelMigrateBucket is a free log retrieval operation binding the contract event 0x6e27f105dcf7c6cebe8da5531a4fad6f386b853381ef2fa5d998efd9ee28006f.
//
// Solidity: event CancelMigrateBucket(address indexed operator)
func (_IStorage *IStorageFilterer) FilterCancelMigrateBucket(opts *bind.FilterOpts, operator []common.Address) (*IStorageCancelMigrateBucketIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "CancelMigrateBucket", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageCancelMigrateBucketIterator{contract: _IStorage.contract, event: "CancelMigrateBucket", logs: logs, sub: sub}, nil
}

// WatchCancelMigrateBucket is a free log subscription operation binding the contract event 0x6e27f105dcf7c6cebe8da5531a4fad6f386b853381ef2fa5d998efd9ee28006f.
//
// Solidity: event CancelMigrateBucket(address indexed operator)
func (_IStorage *IStorageFilterer) WatchCancelMigrateBucket(opts *bind.WatchOpts, sink chan<- *IStorageCancelMigrateBucket, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "CancelMigrateBucket", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageCancelMigrateBucket)
				if err := _IStorage.contract.UnpackLog(event, "CancelMigrateBucket", log); err != nil {
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

// ParseCancelMigrateBucket is a log parse operation binding the contract event 0x6e27f105dcf7c6cebe8da5531a4fad6f386b853381ef2fa5d998efd9ee28006f.
//
// Solidity: event CancelMigrateBucket(address indexed operator)
func (_IStorage *IStorageFilterer) ParseCancelMigrateBucket(log types.Log) (*IStorageCancelMigrateBucket, error) {
	event := new(IStorageCancelMigrateBucket)
	if err := _IStorage.contract.UnpackLog(event, "CancelMigrateBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageCompleteMigrateBucketIterator is returned from FilterCompleteMigrateBucket and is used to iterate over the raw logs and unpacked data for CompleteMigrateBucket events raised by the IStorage contract.
type IStorageCompleteMigrateBucketIterator struct {
	Event *IStorageCompleteMigrateBucket // Event containing the contract specifics and raw log

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
func (it *IStorageCompleteMigrateBucketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageCompleteMigrateBucket)
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
		it.Event = new(IStorageCompleteMigrateBucket)
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
func (it *IStorageCompleteMigrateBucketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageCompleteMigrateBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageCompleteMigrateBucket represents a CompleteMigrateBucket event raised by the IStorage contract.
type IStorageCompleteMigrateBucket struct {
	Creator    common.Address
	BucketName common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCompleteMigrateBucket is a free log retrieval operation binding the contract event 0xeba72ae70b9defdd2439e51a2205710ca135784c5d8194f7861b9cc08ca62ae6.
//
// Solidity: event CompleteMigrateBucket(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) FilterCompleteMigrateBucket(opts *bind.FilterOpts, creator []common.Address, bucketName []string) (*IStorageCompleteMigrateBucketIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var bucketNameRule []interface{}
	for _, bucketNameItem := range bucketName {
		bucketNameRule = append(bucketNameRule, bucketNameItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "CompleteMigrateBucket", creatorRule, bucketNameRule)
	if err != nil {
		return nil, err
	}
	return &IStorageCompleteMigrateBucketIterator{contract: _IStorage.contract, event: "CompleteMigrateBucket", logs: logs, sub: sub}, nil
}

// WatchCompleteMigrateBucket is a free log subscription operation binding the contract event 0xeba72ae70b9defdd2439e51a2205710ca135784c5d8194f7861b9cc08ca62ae6.
//
// Solidity: event CompleteMigrateBucket(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) WatchCompleteMigrateBucket(opts *bind.WatchOpts, sink chan<- *IStorageCompleteMigrateBucket, creator []common.Address, bucketName []string) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var bucketNameRule []interface{}
	for _, bucketNameItem := range bucketName {
		bucketNameRule = append(bucketNameRule, bucketNameItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "CompleteMigrateBucket", creatorRule, bucketNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageCompleteMigrateBucket)
				if err := _IStorage.contract.UnpackLog(event, "CompleteMigrateBucket", log); err != nil {
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

// ParseCompleteMigrateBucket is a log parse operation binding the contract event 0xeba72ae70b9defdd2439e51a2205710ca135784c5d8194f7861b9cc08ca62ae6.
//
// Solidity: event CompleteMigrateBucket(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) ParseCompleteMigrateBucket(log types.Log) (*IStorageCompleteMigrateBucket, error) {
	event := new(IStorageCompleteMigrateBucket)
	if err := _IStorage.contract.UnpackLog(event, "CompleteMigrateBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageCopyObjectIterator is returned from FilterCopyObject and is used to iterate over the raw logs and unpacked data for CopyObject events raised by the IStorage contract.
type IStorageCopyObjectIterator struct {
	Event *IStorageCopyObject // Event containing the contract specifics and raw log

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
func (it *IStorageCopyObjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageCopyObject)
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
		it.Event = new(IStorageCopyObject)
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
func (it *IStorageCopyObjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageCopyObjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageCopyObject represents a CopyObject event raised by the IStorage contract.
type IStorageCopyObject struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCopyObject is a free log retrieval operation binding the contract event 0x5cb0da52ca887b23dbd6cd90a42105194b8eefca5df979a8d4b0418910e2c7c4.
//
// Solidity: event CopyObject(address indexed creator)
func (_IStorage *IStorageFilterer) FilterCopyObject(opts *bind.FilterOpts, creator []common.Address) (*IStorageCopyObjectIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "CopyObject", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageCopyObjectIterator{contract: _IStorage.contract, event: "CopyObject", logs: logs, sub: sub}, nil
}

// WatchCopyObject is a free log subscription operation binding the contract event 0x5cb0da52ca887b23dbd6cd90a42105194b8eefca5df979a8d4b0418910e2c7c4.
//
// Solidity: event CopyObject(address indexed creator)
func (_IStorage *IStorageFilterer) WatchCopyObject(opts *bind.WatchOpts, sink chan<- *IStorageCopyObject, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "CopyObject", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageCopyObject)
				if err := _IStorage.contract.UnpackLog(event, "CopyObject", log); err != nil {
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

// ParseCopyObject is a log parse operation binding the contract event 0x5cb0da52ca887b23dbd6cd90a42105194b8eefca5df979a8d4b0418910e2c7c4.
//
// Solidity: event CopyObject(address indexed creator)
func (_IStorage *IStorageFilterer) ParseCopyObject(log types.Log) (*IStorageCopyObject, error) {
	event := new(IStorageCopyObject)
	if err := _IStorage.contract.UnpackLog(event, "CopyObject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// IStorageDelegateCreateObjectIterator is returned from FilterDelegateCreateObject and is used to iterate over the raw logs and unpacked data for DelegateCreateObject events raised by the IStorage contract.
type IStorageDelegateCreateObjectIterator struct {
	Event *IStorageDelegateCreateObject // Event containing the contract specifics and raw log

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
func (it *IStorageDelegateCreateObjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageDelegateCreateObject)
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
		it.Event = new(IStorageDelegateCreateObject)
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
func (it *IStorageDelegateCreateObjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageDelegateCreateObjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageDelegateCreateObject represents a DelegateCreateObject event raised by the IStorage contract.
type IStorageDelegateCreateObject struct {
	Creator    common.Address
	ObjectName common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDelegateCreateObject is a free log retrieval operation binding the contract event 0x38f646bc35f901b7766b409f017d1258d0bcc952ce4ea9a3876da70721ab02e1.
//
// Solidity: event DelegateCreateObject(address indexed creator, string indexed objectName)
func (_IStorage *IStorageFilterer) FilterDelegateCreateObject(opts *bind.FilterOpts, creator []common.Address, objectName []string) (*IStorageDelegateCreateObjectIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var objectNameRule []interface{}
	for _, objectNameItem := range objectName {
		objectNameRule = append(objectNameRule, objectNameItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "DelegateCreateObject", creatorRule, objectNameRule)
	if err != nil {
		return nil, err
	}
	return &IStorageDelegateCreateObjectIterator{contract: _IStorage.contract, event: "DelegateCreateObject", logs: logs, sub: sub}, nil
}

// WatchDelegateCreateObject is a free log subscription operation binding the contract event 0x38f646bc35f901b7766b409f017d1258d0bcc952ce4ea9a3876da70721ab02e1.
//
// Solidity: event DelegateCreateObject(address indexed creator, string indexed objectName)
func (_IStorage *IStorageFilterer) WatchDelegateCreateObject(opts *bind.WatchOpts, sink chan<- *IStorageDelegateCreateObject, creator []common.Address, objectName []string) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var objectNameRule []interface{}
	for _, objectNameItem := range objectName {
		objectNameRule = append(objectNameRule, objectNameItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "DelegateCreateObject", creatorRule, objectNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageDelegateCreateObject)
				if err := _IStorage.contract.UnpackLog(event, "DelegateCreateObject", log); err != nil {
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

// ParseDelegateCreateObject is a log parse operation binding the contract event 0x38f646bc35f901b7766b409f017d1258d0bcc952ce4ea9a3876da70721ab02e1.
//
// Solidity: event DelegateCreateObject(address indexed creator, string indexed objectName)
func (_IStorage *IStorageFilterer) ParseDelegateCreateObject(log types.Log) (*IStorageDelegateCreateObject, error) {
	event := new(IStorageDelegateCreateObject)
	if err := _IStorage.contract.UnpackLog(event, "DelegateCreateObject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageDelegateUpdateObjectContentIterator is returned from FilterDelegateUpdateObjectContent and is used to iterate over the raw logs and unpacked data for DelegateUpdateObjectContent events raised by the IStorage contract.
type IStorageDelegateUpdateObjectContentIterator struct {
	Event *IStorageDelegateUpdateObjectContent // Event containing the contract specifics and raw log

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
func (it *IStorageDelegateUpdateObjectContentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageDelegateUpdateObjectContent)
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
		it.Event = new(IStorageDelegateUpdateObjectContent)
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
func (it *IStorageDelegateUpdateObjectContentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageDelegateUpdateObjectContentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageDelegateUpdateObjectContent represents a DelegateUpdateObjectContent event raised by the IStorage contract.
type IStorageDelegateUpdateObjectContent struct {
	Operator   common.Address
	ObjectName common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDelegateUpdateObjectContent is a free log retrieval operation binding the contract event 0x979216b16e4ff68a5371da48a48d0084765fec29c135a9a733ff3e56d39cbec8.
//
// Solidity: event DelegateUpdateObjectContent(address indexed operator, string indexed objectName)
func (_IStorage *IStorageFilterer) FilterDelegateUpdateObjectContent(opts *bind.FilterOpts, operator []common.Address, objectName []string) (*IStorageDelegateUpdateObjectContentIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var objectNameRule []interface{}
	for _, objectNameItem := range objectName {
		objectNameRule = append(objectNameRule, objectNameItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "DelegateUpdateObjectContent", operatorRule, objectNameRule)
	if err != nil {
		return nil, err
	}
	return &IStorageDelegateUpdateObjectContentIterator{contract: _IStorage.contract, event: "DelegateUpdateObjectContent", logs: logs, sub: sub}, nil
}

// WatchDelegateUpdateObjectContent is a free log subscription operation binding the contract event 0x979216b16e4ff68a5371da48a48d0084765fec29c135a9a733ff3e56d39cbec8.
//
// Solidity: event DelegateUpdateObjectContent(address indexed operator, string indexed objectName)
func (_IStorage *IStorageFilterer) WatchDelegateUpdateObjectContent(opts *bind.WatchOpts, sink chan<- *IStorageDelegateUpdateObjectContent, operator []common.Address, objectName []string) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var objectNameRule []interface{}
	for _, objectNameItem := range objectName {
		objectNameRule = append(objectNameRule, objectNameItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "DelegateUpdateObjectContent", operatorRule, objectNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageDelegateUpdateObjectContent)
				if err := _IStorage.contract.UnpackLog(event, "DelegateUpdateObjectContent", log); err != nil {
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

// ParseDelegateUpdateObjectContent is a log parse operation binding the contract event 0x979216b16e4ff68a5371da48a48d0084765fec29c135a9a733ff3e56d39cbec8.
//
// Solidity: event DelegateUpdateObjectContent(address indexed operator, string indexed objectName)
func (_IStorage *IStorageFilterer) ParseDelegateUpdateObjectContent(log types.Log) (*IStorageDelegateUpdateObjectContent, error) {
	event := new(IStorageDelegateUpdateObjectContent)
	if err := _IStorage.contract.UnpackLog(event, "DelegateUpdateObjectContent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageDeleteBucketIterator is returned from FilterDeleteBucket and is used to iterate over the raw logs and unpacked data for DeleteBucket events raised by the IStorage contract.
type IStorageDeleteBucketIterator struct {
	Event *IStorageDeleteBucket // Event containing the contract specifics and raw log

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
func (it *IStorageDeleteBucketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageDeleteBucket)
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
		it.Event = new(IStorageDeleteBucket)
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
func (it *IStorageDeleteBucketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageDeleteBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageDeleteBucket represents a DeleteBucket event raised by the IStorage contract.
type IStorageDeleteBucket struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeleteBucket is a free log retrieval operation binding the contract event 0x1c97df6cb6d987c84ba4a4b4bea21619933907f9964e44d89d409d4f105a0588.
//
// Solidity: event DeleteBucket(address indexed creator)
func (_IStorage *IStorageFilterer) FilterDeleteBucket(opts *bind.FilterOpts, creator []common.Address) (*IStorageDeleteBucketIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "DeleteBucket", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageDeleteBucketIterator{contract: _IStorage.contract, event: "DeleteBucket", logs: logs, sub: sub}, nil
}

// WatchDeleteBucket is a free log subscription operation binding the contract event 0x1c97df6cb6d987c84ba4a4b4bea21619933907f9964e44d89d409d4f105a0588.
//
// Solidity: event DeleteBucket(address indexed creator)
func (_IStorage *IStorageFilterer) WatchDeleteBucket(opts *bind.WatchOpts, sink chan<- *IStorageDeleteBucket, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "DeleteBucket", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageDeleteBucket)
				if err := _IStorage.contract.UnpackLog(event, "DeleteBucket", log); err != nil {
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

// ParseDeleteBucket is a log parse operation binding the contract event 0x1c97df6cb6d987c84ba4a4b4bea21619933907f9964e44d89d409d4f105a0588.
//
// Solidity: event DeleteBucket(address indexed creator)
func (_IStorage *IStorageFilterer) ParseDeleteBucket(log types.Log) (*IStorageDeleteBucket, error) {
	event := new(IStorageDeleteBucket)
	if err := _IStorage.contract.UnpackLog(event, "DeleteBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageDeleteGroupIterator is returned from FilterDeleteGroup and is used to iterate over the raw logs and unpacked data for DeleteGroup events raised by the IStorage contract.
type IStorageDeleteGroupIterator struct {
	Event *IStorageDeleteGroup // Event containing the contract specifics and raw log

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
func (it *IStorageDeleteGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageDeleteGroup)
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
		it.Event = new(IStorageDeleteGroup)
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
func (it *IStorageDeleteGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageDeleteGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageDeleteGroup represents a DeleteGroup event raised by the IStorage contract.
type IStorageDeleteGroup struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeleteGroup is a free log retrieval operation binding the contract event 0xc4258be7f08176821b6d80b1eb490b28c3a886aa9d7d9fa1df3b19e9e0c149f3.
//
// Solidity: event DeleteGroup(address indexed creator)
func (_IStorage *IStorageFilterer) FilterDeleteGroup(opts *bind.FilterOpts, creator []common.Address) (*IStorageDeleteGroupIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "DeleteGroup", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageDeleteGroupIterator{contract: _IStorage.contract, event: "DeleteGroup", logs: logs, sub: sub}, nil
}

// WatchDeleteGroup is a free log subscription operation binding the contract event 0xc4258be7f08176821b6d80b1eb490b28c3a886aa9d7d9fa1df3b19e9e0c149f3.
//
// Solidity: event DeleteGroup(address indexed creator)
func (_IStorage *IStorageFilterer) WatchDeleteGroup(opts *bind.WatchOpts, sink chan<- *IStorageDeleteGroup, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "DeleteGroup", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageDeleteGroup)
				if err := _IStorage.contract.UnpackLog(event, "DeleteGroup", log); err != nil {
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

// ParseDeleteGroup is a log parse operation binding the contract event 0xc4258be7f08176821b6d80b1eb490b28c3a886aa9d7d9fa1df3b19e9e0c149f3.
//
// Solidity: event DeleteGroup(address indexed creator)
func (_IStorage *IStorageFilterer) ParseDeleteGroup(log types.Log) (*IStorageDeleteGroup, error) {
	event := new(IStorageDeleteGroup)
	if err := _IStorage.contract.UnpackLog(event, "DeleteGroup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageDeleteObjectIterator is returned from FilterDeleteObject and is used to iterate over the raw logs and unpacked data for DeleteObject events raised by the IStorage contract.
type IStorageDeleteObjectIterator struct {
	Event *IStorageDeleteObject // Event containing the contract specifics and raw log

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
func (it *IStorageDeleteObjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageDeleteObject)
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
		it.Event = new(IStorageDeleteObject)
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
func (it *IStorageDeleteObjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageDeleteObjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageDeleteObject represents a DeleteObject event raised by the IStorage contract.
type IStorageDeleteObject struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeleteObject is a free log retrieval operation binding the contract event 0xc54adfb04b4d6d98c4b607bacc0f1729e9f801ebaad41b343b22b5aa5e34c91a.
//
// Solidity: event DeleteObject(address indexed creator)
func (_IStorage *IStorageFilterer) FilterDeleteObject(opts *bind.FilterOpts, creator []common.Address) (*IStorageDeleteObjectIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "DeleteObject", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageDeleteObjectIterator{contract: _IStorage.contract, event: "DeleteObject", logs: logs, sub: sub}, nil
}

// WatchDeleteObject is a free log subscription operation binding the contract event 0xc54adfb04b4d6d98c4b607bacc0f1729e9f801ebaad41b343b22b5aa5e34c91a.
//
// Solidity: event DeleteObject(address indexed creator)
func (_IStorage *IStorageFilterer) WatchDeleteObject(opts *bind.WatchOpts, sink chan<- *IStorageDeleteObject, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "DeleteObject", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageDeleteObject)
				if err := _IStorage.contract.UnpackLog(event, "DeleteObject", log); err != nil {
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

// ParseDeleteObject is a log parse operation binding the contract event 0xc54adfb04b4d6d98c4b607bacc0f1729e9f801ebaad41b343b22b5aa5e34c91a.
//
// Solidity: event DeleteObject(address indexed creator)
func (_IStorage *IStorageFilterer) ParseDeleteObject(log types.Log) (*IStorageDeleteObject, error) {
	event := new(IStorageDeleteObject)
	if err := _IStorage.contract.UnpackLog(event, "DeleteObject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageDeletePolicyIterator is returned from FilterDeletePolicy and is used to iterate over the raw logs and unpacked data for DeletePolicy events raised by the IStorage contract.
type IStorageDeletePolicyIterator struct {
	Event *IStorageDeletePolicy // Event containing the contract specifics and raw log

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
func (it *IStorageDeletePolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageDeletePolicy)
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
		it.Event = new(IStorageDeletePolicy)
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
func (it *IStorageDeletePolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageDeletePolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageDeletePolicy represents a DeletePolicy event raised by the IStorage contract.
type IStorageDeletePolicy struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeletePolicy is a free log retrieval operation binding the contract event 0xfb87ae16150525ea57a1260792ea7780a6bec011c39950dfc7480c8ed9231293.
//
// Solidity: event DeletePolicy(address indexed creator)
func (_IStorage *IStorageFilterer) FilterDeletePolicy(opts *bind.FilterOpts, creator []common.Address) (*IStorageDeletePolicyIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "DeletePolicy", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageDeletePolicyIterator{contract: _IStorage.contract, event: "DeletePolicy", logs: logs, sub: sub}, nil
}

// WatchDeletePolicy is a free log subscription operation binding the contract event 0xfb87ae16150525ea57a1260792ea7780a6bec011c39950dfc7480c8ed9231293.
//
// Solidity: event DeletePolicy(address indexed creator)
func (_IStorage *IStorageFilterer) WatchDeletePolicy(opts *bind.WatchOpts, sink chan<- *IStorageDeletePolicy, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "DeletePolicy", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageDeletePolicy)
				if err := _IStorage.contract.UnpackLog(event, "DeletePolicy", log); err != nil {
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

// ParseDeletePolicy is a log parse operation binding the contract event 0xfb87ae16150525ea57a1260792ea7780a6bec011c39950dfc7480c8ed9231293.
//
// Solidity: event DeletePolicy(address indexed creator)
func (_IStorage *IStorageFilterer) ParseDeletePolicy(log types.Log) (*IStorageDeletePolicy, error) {
	event := new(IStorageDeletePolicy)
	if err := _IStorage.contract.UnpackLog(event, "DeletePolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageDiscontinueBucketIterator is returned from FilterDiscontinueBucket and is used to iterate over the raw logs and unpacked data for DiscontinueBucket events raised by the IStorage contract.
type IStorageDiscontinueBucketIterator struct {
	Event *IStorageDiscontinueBucket // Event containing the contract specifics and raw log

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
func (it *IStorageDiscontinueBucketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageDiscontinueBucket)
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
		it.Event = new(IStorageDiscontinueBucket)
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
func (it *IStorageDiscontinueBucketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageDiscontinueBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageDiscontinueBucket represents a DiscontinueBucket event raised by the IStorage contract.
type IStorageDiscontinueBucket struct {
	Creator    common.Address
	BucketName common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiscontinueBucket is a free log retrieval operation binding the contract event 0x2ee5ab990bbd8ed278298b41c204ab2551c1164a1df6270297f554f8e30e151e.
//
// Solidity: event DiscontinueBucket(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) FilterDiscontinueBucket(opts *bind.FilterOpts, creator []common.Address, bucketName []string) (*IStorageDiscontinueBucketIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var bucketNameRule []interface{}
	for _, bucketNameItem := range bucketName {
		bucketNameRule = append(bucketNameRule, bucketNameItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "DiscontinueBucket", creatorRule, bucketNameRule)
	if err != nil {
		return nil, err
	}
	return &IStorageDiscontinueBucketIterator{contract: _IStorage.contract, event: "DiscontinueBucket", logs: logs, sub: sub}, nil
}

// WatchDiscontinueBucket is a free log subscription operation binding the contract event 0x2ee5ab990bbd8ed278298b41c204ab2551c1164a1df6270297f554f8e30e151e.
//
// Solidity: event DiscontinueBucket(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) WatchDiscontinueBucket(opts *bind.WatchOpts, sink chan<- *IStorageDiscontinueBucket, creator []common.Address, bucketName []string) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var bucketNameRule []interface{}
	for _, bucketNameItem := range bucketName {
		bucketNameRule = append(bucketNameRule, bucketNameItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "DiscontinueBucket", creatorRule, bucketNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageDiscontinueBucket)
				if err := _IStorage.contract.UnpackLog(event, "DiscontinueBucket", log); err != nil {
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

// ParseDiscontinueBucket is a log parse operation binding the contract event 0x2ee5ab990bbd8ed278298b41c204ab2551c1164a1df6270297f554f8e30e151e.
//
// Solidity: event DiscontinueBucket(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) ParseDiscontinueBucket(log types.Log) (*IStorageDiscontinueBucket, error) {
	event := new(IStorageDiscontinueBucket)
	if err := _IStorage.contract.UnpackLog(event, "DiscontinueBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageDiscontinueObjectIterator is returned from FilterDiscontinueObject and is used to iterate over the raw logs and unpacked data for DiscontinueObject events raised by the IStorage contract.
type IStorageDiscontinueObjectIterator struct {
	Event *IStorageDiscontinueObject // Event containing the contract specifics and raw log

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
func (it *IStorageDiscontinueObjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageDiscontinueObject)
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
		it.Event = new(IStorageDiscontinueObject)
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
func (it *IStorageDiscontinueObjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageDiscontinueObjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageDiscontinueObject represents a DiscontinueObject event raised by the IStorage contract.
type IStorageDiscontinueObject struct {
	Creator    common.Address
	BucketName common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiscontinueObject is a free log retrieval operation binding the contract event 0x925473a0d60cbdc222d820b2e728f182468750730278055981895ba0b3450efb.
//
// Solidity: event DiscontinueObject(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) FilterDiscontinueObject(opts *bind.FilterOpts, creator []common.Address, bucketName []string) (*IStorageDiscontinueObjectIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var bucketNameRule []interface{}
	for _, bucketNameItem := range bucketName {
		bucketNameRule = append(bucketNameRule, bucketNameItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "DiscontinueObject", creatorRule, bucketNameRule)
	if err != nil {
		return nil, err
	}
	return &IStorageDiscontinueObjectIterator{contract: _IStorage.contract, event: "DiscontinueObject", logs: logs, sub: sub}, nil
}

// WatchDiscontinueObject is a free log subscription operation binding the contract event 0x925473a0d60cbdc222d820b2e728f182468750730278055981895ba0b3450efb.
//
// Solidity: event DiscontinueObject(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) WatchDiscontinueObject(opts *bind.WatchOpts, sink chan<- *IStorageDiscontinueObject, creator []common.Address, bucketName []string) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var bucketNameRule []interface{}
	for _, bucketNameItem := range bucketName {
		bucketNameRule = append(bucketNameRule, bucketNameItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "DiscontinueObject", creatorRule, bucketNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageDiscontinueObject)
				if err := _IStorage.contract.UnpackLog(event, "DiscontinueObject", log); err != nil {
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

// ParseDiscontinueObject is a log parse operation binding the contract event 0x925473a0d60cbdc222d820b2e728f182468750730278055981895ba0b3450efb.
//
// Solidity: event DiscontinueObject(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) ParseDiscontinueObject(log types.Log) (*IStorageDiscontinueObject, error) {
	event := new(IStorageDiscontinueObject)
	if err := _IStorage.contract.UnpackLog(event, "DiscontinueObject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageLeaveGroupIterator is returned from FilterLeaveGroup and is used to iterate over the raw logs and unpacked data for LeaveGroup events raised by the IStorage contract.
type IStorageLeaveGroupIterator struct {
	Event *IStorageLeaveGroup // Event containing the contract specifics and raw log

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
func (it *IStorageLeaveGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageLeaveGroup)
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
		it.Event = new(IStorageLeaveGroup)
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
func (it *IStorageLeaveGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageLeaveGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageLeaveGroup represents a LeaveGroup event raised by the IStorage contract.
type IStorageLeaveGroup struct {
	Creator   common.Address
	GroupName common.Hash
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLeaveGroup is a free log retrieval operation binding the contract event 0x38ba0a44ad9e4b285ba2c7157d9703f451381356ea282c7164022bfac4fbbfe8.
//
// Solidity: event LeaveGroup(address indexed creator, string indexed groupName)
func (_IStorage *IStorageFilterer) FilterLeaveGroup(opts *bind.FilterOpts, creator []common.Address, groupName []string) (*IStorageLeaveGroupIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var groupNameRule []interface{}
	for _, groupNameItem := range groupName {
		groupNameRule = append(groupNameRule, groupNameItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "LeaveGroup", creatorRule, groupNameRule)
	if err != nil {
		return nil, err
	}
	return &IStorageLeaveGroupIterator{contract: _IStorage.contract, event: "LeaveGroup", logs: logs, sub: sub}, nil
}

// WatchLeaveGroup is a free log subscription operation binding the contract event 0x38ba0a44ad9e4b285ba2c7157d9703f451381356ea282c7164022bfac4fbbfe8.
//
// Solidity: event LeaveGroup(address indexed creator, string indexed groupName)
func (_IStorage *IStorageFilterer) WatchLeaveGroup(opts *bind.WatchOpts, sink chan<- *IStorageLeaveGroup, creator []common.Address, groupName []string) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var groupNameRule []interface{}
	for _, groupNameItem := range groupName {
		groupNameRule = append(groupNameRule, groupNameItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "LeaveGroup", creatorRule, groupNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageLeaveGroup)
				if err := _IStorage.contract.UnpackLog(event, "LeaveGroup", log); err != nil {
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

// ParseLeaveGroup is a log parse operation binding the contract event 0x38ba0a44ad9e4b285ba2c7157d9703f451381356ea282c7164022bfac4fbbfe8.
//
// Solidity: event LeaveGroup(address indexed creator, string indexed groupName)
func (_IStorage *IStorageFilterer) ParseLeaveGroup(log types.Log) (*IStorageLeaveGroup, error) {
	event := new(IStorageLeaveGroup)
	if err := _IStorage.contract.UnpackLog(event, "LeaveGroup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageMigrateBucketIterator is returned from FilterMigrateBucket and is used to iterate over the raw logs and unpacked data for MigrateBucket events raised by the IStorage contract.
type IStorageMigrateBucketIterator struct {
	Event *IStorageMigrateBucket // Event containing the contract specifics and raw log

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
func (it *IStorageMigrateBucketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageMigrateBucket)
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
		it.Event = new(IStorageMigrateBucket)
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
func (it *IStorageMigrateBucketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageMigrateBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageMigrateBucket represents a MigrateBucket event raised by the IStorage contract.
type IStorageMigrateBucket struct {
	Creator    common.Address
	BucketName common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMigrateBucket is a free log retrieval operation binding the contract event 0xa76fc2fcb8eef2e9b5f1fcd143e4479356bd05664273f947fe08872be011a82f.
//
// Solidity: event MigrateBucket(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) FilterMigrateBucket(opts *bind.FilterOpts, creator []common.Address, bucketName []string) (*IStorageMigrateBucketIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var bucketNameRule []interface{}
	for _, bucketNameItem := range bucketName {
		bucketNameRule = append(bucketNameRule, bucketNameItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "MigrateBucket", creatorRule, bucketNameRule)
	if err != nil {
		return nil, err
	}
	return &IStorageMigrateBucketIterator{contract: _IStorage.contract, event: "MigrateBucket", logs: logs, sub: sub}, nil
}

// WatchMigrateBucket is a free log subscription operation binding the contract event 0xa76fc2fcb8eef2e9b5f1fcd143e4479356bd05664273f947fe08872be011a82f.
//
// Solidity: event MigrateBucket(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) WatchMigrateBucket(opts *bind.WatchOpts, sink chan<- *IStorageMigrateBucket, creator []common.Address, bucketName []string) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var bucketNameRule []interface{}
	for _, bucketNameItem := range bucketName {
		bucketNameRule = append(bucketNameRule, bucketNameItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "MigrateBucket", creatorRule, bucketNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageMigrateBucket)
				if err := _IStorage.contract.UnpackLog(event, "MigrateBucket", log); err != nil {
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

// ParseMigrateBucket is a log parse operation binding the contract event 0xa76fc2fcb8eef2e9b5f1fcd143e4479356bd05664273f947fe08872be011a82f.
//
// Solidity: event MigrateBucket(address indexed creator, string indexed bucketName)
func (_IStorage *IStorageFilterer) ParseMigrateBucket(log types.Log) (*IStorageMigrateBucket, error) {
	event := new(IStorageMigrateBucket)
	if err := _IStorage.contract.UnpackLog(event, "MigrateBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageMirrorBucketIterator is returned from FilterMirrorBucket and is used to iterate over the raw logs and unpacked data for MirrorBucket events raised by the IStorage contract.
type IStorageMirrorBucketIterator struct {
	Event *IStorageMirrorBucket // Event containing the contract specifics and raw log

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
func (it *IStorageMirrorBucketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageMirrorBucket)
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
		it.Event = new(IStorageMirrorBucket)
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
func (it *IStorageMirrorBucketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageMirrorBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageMirrorBucket represents a MirrorBucket event raised by the IStorage contract.
type IStorageMirrorBucket struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMirrorBucket is a free log retrieval operation binding the contract event 0xb6c1620d165c25ff666f4a0df1f79cb7f97d351d1898f80357d7b0346a65c3f7.
//
// Solidity: event MirrorBucket(address indexed creator)
func (_IStorage *IStorageFilterer) FilterMirrorBucket(opts *bind.FilterOpts, creator []common.Address) (*IStorageMirrorBucketIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "MirrorBucket", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageMirrorBucketIterator{contract: _IStorage.contract, event: "MirrorBucket", logs: logs, sub: sub}, nil
}

// WatchMirrorBucket is a free log subscription operation binding the contract event 0xb6c1620d165c25ff666f4a0df1f79cb7f97d351d1898f80357d7b0346a65c3f7.
//
// Solidity: event MirrorBucket(address indexed creator)
func (_IStorage *IStorageFilterer) WatchMirrorBucket(opts *bind.WatchOpts, sink chan<- *IStorageMirrorBucket, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "MirrorBucket", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageMirrorBucket)
				if err := _IStorage.contract.UnpackLog(event, "MirrorBucket", log); err != nil {
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

// ParseMirrorBucket is a log parse operation binding the contract event 0xb6c1620d165c25ff666f4a0df1f79cb7f97d351d1898f80357d7b0346a65c3f7.
//
// Solidity: event MirrorBucket(address indexed creator)
func (_IStorage *IStorageFilterer) ParseMirrorBucket(log types.Log) (*IStorageMirrorBucket, error) {
	event := new(IStorageMirrorBucket)
	if err := _IStorage.contract.UnpackLog(event, "MirrorBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageMirrorGroupIterator is returned from FilterMirrorGroup and is used to iterate over the raw logs and unpacked data for MirrorGroup events raised by the IStorage contract.
type IStorageMirrorGroupIterator struct {
	Event *IStorageMirrorGroup // Event containing the contract specifics and raw log

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
func (it *IStorageMirrorGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageMirrorGroup)
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
		it.Event = new(IStorageMirrorGroup)
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
func (it *IStorageMirrorGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageMirrorGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageMirrorGroup represents a MirrorGroup event raised by the IStorage contract.
type IStorageMirrorGroup struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMirrorGroup is a free log retrieval operation binding the contract event 0x7235e5f429b926cd88892ff27f14a6d305e2f8c381de25c652b6cc6e15f4b7f2.
//
// Solidity: event MirrorGroup(address indexed creator)
func (_IStorage *IStorageFilterer) FilterMirrorGroup(opts *bind.FilterOpts, creator []common.Address) (*IStorageMirrorGroupIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "MirrorGroup", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageMirrorGroupIterator{contract: _IStorage.contract, event: "MirrorGroup", logs: logs, sub: sub}, nil
}

// WatchMirrorGroup is a free log subscription operation binding the contract event 0x7235e5f429b926cd88892ff27f14a6d305e2f8c381de25c652b6cc6e15f4b7f2.
//
// Solidity: event MirrorGroup(address indexed creator)
func (_IStorage *IStorageFilterer) WatchMirrorGroup(opts *bind.WatchOpts, sink chan<- *IStorageMirrorGroup, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "MirrorGroup", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageMirrorGroup)
				if err := _IStorage.contract.UnpackLog(event, "MirrorGroup", log); err != nil {
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

// ParseMirrorGroup is a log parse operation binding the contract event 0x7235e5f429b926cd88892ff27f14a6d305e2f8c381de25c652b6cc6e15f4b7f2.
//
// Solidity: event MirrorGroup(address indexed creator)
func (_IStorage *IStorageFilterer) ParseMirrorGroup(log types.Log) (*IStorageMirrorGroup, error) {
	event := new(IStorageMirrorGroup)
	if err := _IStorage.contract.UnpackLog(event, "MirrorGroup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageMirrorObjectIterator is returned from FilterMirrorObject and is used to iterate over the raw logs and unpacked data for MirrorObject events raised by the IStorage contract.
type IStorageMirrorObjectIterator struct {
	Event *IStorageMirrorObject // Event containing the contract specifics and raw log

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
func (it *IStorageMirrorObjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageMirrorObject)
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
		it.Event = new(IStorageMirrorObject)
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
func (it *IStorageMirrorObjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageMirrorObjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageMirrorObject represents a MirrorObject event raised by the IStorage contract.
type IStorageMirrorObject struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMirrorObject is a free log retrieval operation binding the contract event 0x0ae2d85a98dbcfa0583dfd7067a32b53ec6d22d0461abbc774d7ba29d3d7d598.
//
// Solidity: event MirrorObject(address indexed creator)
func (_IStorage *IStorageFilterer) FilterMirrorObject(opts *bind.FilterOpts, creator []common.Address) (*IStorageMirrorObjectIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "MirrorObject", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageMirrorObjectIterator{contract: _IStorage.contract, event: "MirrorObject", logs: logs, sub: sub}, nil
}

// WatchMirrorObject is a free log subscription operation binding the contract event 0x0ae2d85a98dbcfa0583dfd7067a32b53ec6d22d0461abbc774d7ba29d3d7d598.
//
// Solidity: event MirrorObject(address indexed creator)
func (_IStorage *IStorageFilterer) WatchMirrorObject(opts *bind.WatchOpts, sink chan<- *IStorageMirrorObject, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "MirrorObject", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageMirrorObject)
				if err := _IStorage.contract.UnpackLog(event, "MirrorObject", log); err != nil {
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

// ParseMirrorObject is a log parse operation binding the contract event 0x0ae2d85a98dbcfa0583dfd7067a32b53ec6d22d0461abbc774d7ba29d3d7d598.
//
// Solidity: event MirrorObject(address indexed creator)
func (_IStorage *IStorageFilterer) ParseMirrorObject(log types.Log) (*IStorageMirrorObject, error) {
	event := new(IStorageMirrorObject)
	if err := _IStorage.contract.UnpackLog(event, "MirrorObject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStoragePutPolicyIterator is returned from FilterPutPolicy and is used to iterate over the raw logs and unpacked data for PutPolicy events raised by the IStorage contract.
type IStoragePutPolicyIterator struct {
	Event *IStoragePutPolicy // Event containing the contract specifics and raw log

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
func (it *IStoragePutPolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStoragePutPolicy)
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
		it.Event = new(IStoragePutPolicy)
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
func (it *IStoragePutPolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStoragePutPolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStoragePutPolicy represents a PutPolicy event raised by the IStorage contract.
type IStoragePutPolicy struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPutPolicy is a free log retrieval operation binding the contract event 0xc46fd2adacf06288b8470f722aa77ee9a299109d9df512b314857eaa09858221.
//
// Solidity: event PutPolicy(address indexed creator)
func (_IStorage *IStorageFilterer) FilterPutPolicy(opts *bind.FilterOpts, creator []common.Address) (*IStoragePutPolicyIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "PutPolicy", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStoragePutPolicyIterator{contract: _IStorage.contract, event: "PutPolicy", logs: logs, sub: sub}, nil
}

// WatchPutPolicy is a free log subscription operation binding the contract event 0xc46fd2adacf06288b8470f722aa77ee9a299109d9df512b314857eaa09858221.
//
// Solidity: event PutPolicy(address indexed creator)
func (_IStorage *IStorageFilterer) WatchPutPolicy(opts *bind.WatchOpts, sink chan<- *IStoragePutPolicy, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "PutPolicy", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStoragePutPolicy)
				if err := _IStorage.contract.UnpackLog(event, "PutPolicy", log); err != nil {
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

// ParsePutPolicy is a log parse operation binding the contract event 0xc46fd2adacf06288b8470f722aa77ee9a299109d9df512b314857eaa09858221.
//
// Solidity: event PutPolicy(address indexed creator)
func (_IStorage *IStorageFilterer) ParsePutPolicy(log types.Log) (*IStoragePutPolicy, error) {
	event := new(IStoragePutPolicy)
	if err := _IStorage.contract.UnpackLog(event, "PutPolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageRejectMigrateBucketIterator is returned from FilterRejectMigrateBucket and is used to iterate over the raw logs and unpacked data for RejectMigrateBucket events raised by the IStorage contract.
type IStorageRejectMigrateBucketIterator struct {
	Event *IStorageRejectMigrateBucket // Event containing the contract specifics and raw log

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
func (it *IStorageRejectMigrateBucketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageRejectMigrateBucket)
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
		it.Event = new(IStorageRejectMigrateBucket)
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
func (it *IStorageRejectMigrateBucketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageRejectMigrateBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageRejectMigrateBucket represents a RejectMigrateBucket event raised by the IStorage contract.
type IStorageRejectMigrateBucket struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRejectMigrateBucket is a free log retrieval operation binding the contract event 0x1f664f86e8533d262593c6ae0467ffd6a53c602ed6bda5f365ad07de48cd75c7.
//
// Solidity: event RejectMigrateBucket(address indexed operator)
func (_IStorage *IStorageFilterer) FilterRejectMigrateBucket(opts *bind.FilterOpts, operator []common.Address) (*IStorageRejectMigrateBucketIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "RejectMigrateBucket", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageRejectMigrateBucketIterator{contract: _IStorage.contract, event: "RejectMigrateBucket", logs: logs, sub: sub}, nil
}

// WatchRejectMigrateBucket is a free log subscription operation binding the contract event 0x1f664f86e8533d262593c6ae0467ffd6a53c602ed6bda5f365ad07de48cd75c7.
//
// Solidity: event RejectMigrateBucket(address indexed operator)
func (_IStorage *IStorageFilterer) WatchRejectMigrateBucket(opts *bind.WatchOpts, sink chan<- *IStorageRejectMigrateBucket, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "RejectMigrateBucket", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageRejectMigrateBucket)
				if err := _IStorage.contract.UnpackLog(event, "RejectMigrateBucket", log); err != nil {
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

// ParseRejectMigrateBucket is a log parse operation binding the contract event 0x1f664f86e8533d262593c6ae0467ffd6a53c602ed6bda5f365ad07de48cd75c7.
//
// Solidity: event RejectMigrateBucket(address indexed operator)
func (_IStorage *IStorageFilterer) ParseRejectMigrateBucket(log types.Log) (*IStorageRejectMigrateBucket, error) {
	event := new(IStorageRejectMigrateBucket)
	if err := _IStorage.contract.UnpackLog(event, "RejectMigrateBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageRejectSealObjectIterator is returned from FilterRejectSealObject and is used to iterate over the raw logs and unpacked data for RejectSealObject events raised by the IStorage contract.
type IStorageRejectSealObjectIterator struct {
	Event *IStorageRejectSealObject // Event containing the contract specifics and raw log

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
func (it *IStorageRejectSealObjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageRejectSealObject)
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
		it.Event = new(IStorageRejectSealObject)
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
func (it *IStorageRejectSealObjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageRejectSealObjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageRejectSealObject represents a RejectSealObject event raised by the IStorage contract.
type IStorageRejectSealObject struct {
	Creator    common.Address
	ObjectName common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRejectSealObject is a free log retrieval operation binding the contract event 0x1afee03fb6ca653f9fddeb95cc858ffce282dff6514f1643055d4bc0b84d270f.
//
// Solidity: event RejectSealObject(address indexed creator, string indexed objectName)
func (_IStorage *IStorageFilterer) FilterRejectSealObject(opts *bind.FilterOpts, creator []common.Address, objectName []string) (*IStorageRejectSealObjectIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var objectNameRule []interface{}
	for _, objectNameItem := range objectName {
		objectNameRule = append(objectNameRule, objectNameItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "RejectSealObject", creatorRule, objectNameRule)
	if err != nil {
		return nil, err
	}
	return &IStorageRejectSealObjectIterator{contract: _IStorage.contract, event: "RejectSealObject", logs: logs, sub: sub}, nil
}

// WatchRejectSealObject is a free log subscription operation binding the contract event 0x1afee03fb6ca653f9fddeb95cc858ffce282dff6514f1643055d4bc0b84d270f.
//
// Solidity: event RejectSealObject(address indexed creator, string indexed objectName)
func (_IStorage *IStorageFilterer) WatchRejectSealObject(opts *bind.WatchOpts, sink chan<- *IStorageRejectSealObject, creator []common.Address, objectName []string) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var objectNameRule []interface{}
	for _, objectNameItem := range objectName {
		objectNameRule = append(objectNameRule, objectNameItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "RejectSealObject", creatorRule, objectNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageRejectSealObject)
				if err := _IStorage.contract.UnpackLog(event, "RejectSealObject", log); err != nil {
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

// ParseRejectSealObject is a log parse operation binding the contract event 0x1afee03fb6ca653f9fddeb95cc858ffce282dff6514f1643055d4bc0b84d270f.
//
// Solidity: event RejectSealObject(address indexed creator, string indexed objectName)
func (_IStorage *IStorageFilterer) ParseRejectSealObject(log types.Log) (*IStorageRejectSealObject, error) {
	event := new(IStorageRejectSealObject)
	if err := _IStorage.contract.UnpackLog(event, "RejectSealObject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageRenewGroupMemberIterator is returned from FilterRenewGroupMember and is used to iterate over the raw logs and unpacked data for RenewGroupMember events raised by the IStorage contract.
type IStorageRenewGroupMemberIterator struct {
	Event *IStorageRenewGroupMember // Event containing the contract specifics and raw log

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
func (it *IStorageRenewGroupMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageRenewGroupMember)
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
		it.Event = new(IStorageRenewGroupMember)
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
func (it *IStorageRenewGroupMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageRenewGroupMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageRenewGroupMember represents a RenewGroupMember event raised by the IStorage contract.
type IStorageRenewGroupMember struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRenewGroupMember is a free log retrieval operation binding the contract event 0x25951b953380a3fb2f6ba0aa76d234904bbc3718509f8c5e5f86489694070090.
//
// Solidity: event RenewGroupMember(address indexed creator)
func (_IStorage *IStorageFilterer) FilterRenewGroupMember(opts *bind.FilterOpts, creator []common.Address) (*IStorageRenewGroupMemberIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "RenewGroupMember", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageRenewGroupMemberIterator{contract: _IStorage.contract, event: "RenewGroupMember", logs: logs, sub: sub}, nil
}

// WatchRenewGroupMember is a free log subscription operation binding the contract event 0x25951b953380a3fb2f6ba0aa76d234904bbc3718509f8c5e5f86489694070090.
//
// Solidity: event RenewGroupMember(address indexed creator)
func (_IStorage *IStorageFilterer) WatchRenewGroupMember(opts *bind.WatchOpts, sink chan<- *IStorageRenewGroupMember, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "RenewGroupMember", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageRenewGroupMember)
				if err := _IStorage.contract.UnpackLog(event, "RenewGroupMember", log); err != nil {
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

// ParseRenewGroupMember is a log parse operation binding the contract event 0x25951b953380a3fb2f6ba0aa76d234904bbc3718509f8c5e5f86489694070090.
//
// Solidity: event RenewGroupMember(address indexed creator)
func (_IStorage *IStorageFilterer) ParseRenewGroupMember(log types.Log) (*IStorageRenewGroupMember, error) {
	event := new(IStorageRenewGroupMember)
	if err := _IStorage.contract.UnpackLog(event, "RenewGroupMember", log); err != nil {
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

// IStorageSetBucketFlowRateLimitIterator is returned from FilterSetBucketFlowRateLimit and is used to iterate over the raw logs and unpacked data for SetBucketFlowRateLimit events raised by the IStorage contract.
type IStorageSetBucketFlowRateLimitIterator struct {
	Event *IStorageSetBucketFlowRateLimit // Event containing the contract specifics and raw log

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
func (it *IStorageSetBucketFlowRateLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageSetBucketFlowRateLimit)
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
		it.Event = new(IStorageSetBucketFlowRateLimit)
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
func (it *IStorageSetBucketFlowRateLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageSetBucketFlowRateLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageSetBucketFlowRateLimit represents a SetBucketFlowRateLimit event raised by the IStorage contract.
type IStorageSetBucketFlowRateLimit struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetBucketFlowRateLimit is a free log retrieval operation binding the contract event 0x722bba3e1fbef468e65fcbfb35df6e99943f0c596f95ec0af215a73132367653.
//
// Solidity: event SetBucketFlowRateLimit(address indexed operator)
func (_IStorage *IStorageFilterer) FilterSetBucketFlowRateLimit(opts *bind.FilterOpts, operator []common.Address) (*IStorageSetBucketFlowRateLimitIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "SetBucketFlowRateLimit", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageSetBucketFlowRateLimitIterator{contract: _IStorage.contract, event: "SetBucketFlowRateLimit", logs: logs, sub: sub}, nil
}

// WatchSetBucketFlowRateLimit is a free log subscription operation binding the contract event 0x722bba3e1fbef468e65fcbfb35df6e99943f0c596f95ec0af215a73132367653.
//
// Solidity: event SetBucketFlowRateLimit(address indexed operator)
func (_IStorage *IStorageFilterer) WatchSetBucketFlowRateLimit(opts *bind.WatchOpts, sink chan<- *IStorageSetBucketFlowRateLimit, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "SetBucketFlowRateLimit", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageSetBucketFlowRateLimit)
				if err := _IStorage.contract.UnpackLog(event, "SetBucketFlowRateLimit", log); err != nil {
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

// ParseSetBucketFlowRateLimit is a log parse operation binding the contract event 0x722bba3e1fbef468e65fcbfb35df6e99943f0c596f95ec0af215a73132367653.
//
// Solidity: event SetBucketFlowRateLimit(address indexed operator)
func (_IStorage *IStorageFilterer) ParseSetBucketFlowRateLimit(log types.Log) (*IStorageSetBucketFlowRateLimit, error) {
	event := new(IStorageSetBucketFlowRateLimit)
	if err := _IStorage.contract.UnpackLog(event, "SetBucketFlowRateLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageSetTagIterator is returned from FilterSetTag and is used to iterate over the raw logs and unpacked data for SetTag events raised by the IStorage contract.
type IStorageSetTagIterator struct {
	Event *IStorageSetTag // Event containing the contract specifics and raw log

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
func (it *IStorageSetTagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageSetTag)
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
		it.Event = new(IStorageSetTag)
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
func (it *IStorageSetTagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageSetTagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageSetTag represents a SetTag event raised by the IStorage contract.
type IStorageSetTag struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetTag is a free log retrieval operation binding the contract event 0x270bf20dd1c6d607c4c6c75cfa0f88efd299bb5e9bc897183edd135857175d74.
//
// Solidity: event SetTag(address indexed creator)
func (_IStorage *IStorageFilterer) FilterSetTag(opts *bind.FilterOpts, creator []common.Address) (*IStorageSetTagIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "SetTag", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageSetTagIterator{contract: _IStorage.contract, event: "SetTag", logs: logs, sub: sub}, nil
}

// WatchSetTag is a free log subscription operation binding the contract event 0x270bf20dd1c6d607c4c6c75cfa0f88efd299bb5e9bc897183edd135857175d74.
//
// Solidity: event SetTag(address indexed creator)
func (_IStorage *IStorageFilterer) WatchSetTag(opts *bind.WatchOpts, sink chan<- *IStorageSetTag, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "SetTag", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageSetTag)
				if err := _IStorage.contract.UnpackLog(event, "SetTag", log); err != nil {
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

// ParseSetTag is a log parse operation binding the contract event 0x270bf20dd1c6d607c4c6c75cfa0f88efd299bb5e9bc897183edd135857175d74.
//
// Solidity: event SetTag(address indexed creator)
func (_IStorage *IStorageFilterer) ParseSetTag(log types.Log) (*IStorageSetTag, error) {
	event := new(IStorageSetTag)
	if err := _IStorage.contract.UnpackLog(event, "SetTag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageToggleSPAsDelegatedAgentIterator is returned from FilterToggleSPAsDelegatedAgent and is used to iterate over the raw logs and unpacked data for ToggleSPAsDelegatedAgent events raised by the IStorage contract.
type IStorageToggleSPAsDelegatedAgentIterator struct {
	Event *IStorageToggleSPAsDelegatedAgent // Event containing the contract specifics and raw log

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
func (it *IStorageToggleSPAsDelegatedAgentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageToggleSPAsDelegatedAgent)
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
		it.Event = new(IStorageToggleSPAsDelegatedAgent)
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
func (it *IStorageToggleSPAsDelegatedAgentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageToggleSPAsDelegatedAgentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageToggleSPAsDelegatedAgent represents a ToggleSPAsDelegatedAgent event raised by the IStorage contract.
type IStorageToggleSPAsDelegatedAgent struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterToggleSPAsDelegatedAgent is a free log retrieval operation binding the contract event 0xa3a131308b1f8ab3c5332b1bc3b5be97830c8193888dcfb9f6877da1cd834b4f.
//
// Solidity: event ToggleSPAsDelegatedAgent(address indexed creator)
func (_IStorage *IStorageFilterer) FilterToggleSPAsDelegatedAgent(opts *bind.FilterOpts, creator []common.Address) (*IStorageToggleSPAsDelegatedAgentIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "ToggleSPAsDelegatedAgent", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageToggleSPAsDelegatedAgentIterator{contract: _IStorage.contract, event: "ToggleSPAsDelegatedAgent", logs: logs, sub: sub}, nil
}

// WatchToggleSPAsDelegatedAgent is a free log subscription operation binding the contract event 0xa3a131308b1f8ab3c5332b1bc3b5be97830c8193888dcfb9f6877da1cd834b4f.
//
// Solidity: event ToggleSPAsDelegatedAgent(address indexed creator)
func (_IStorage *IStorageFilterer) WatchToggleSPAsDelegatedAgent(opts *bind.WatchOpts, sink chan<- *IStorageToggleSPAsDelegatedAgent, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "ToggleSPAsDelegatedAgent", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageToggleSPAsDelegatedAgent)
				if err := _IStorage.contract.UnpackLog(event, "ToggleSPAsDelegatedAgent", log); err != nil {
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

// ParseToggleSPAsDelegatedAgent is a log parse operation binding the contract event 0xa3a131308b1f8ab3c5332b1bc3b5be97830c8193888dcfb9f6877da1cd834b4f.
//
// Solidity: event ToggleSPAsDelegatedAgent(address indexed creator)
func (_IStorage *IStorageFilterer) ParseToggleSPAsDelegatedAgent(log types.Log) (*IStorageToggleSPAsDelegatedAgent, error) {
	event := new(IStorageToggleSPAsDelegatedAgent)
	if err := _IStorage.contract.UnpackLog(event, "ToggleSPAsDelegatedAgent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IStorage contract.
type IStorageTransferIterator struct {
	Event *IStorageTransfer // Event containing the contract specifics and raw log

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
func (it *IStorageTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageTransfer)
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
		it.Event = new(IStorageTransfer)
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
func (it *IStorageTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageTransfer represents a Transfer event raised by the IStorage contract.
type IStorageTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IStorage *IStorageFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IStorageTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IStorageTransferIterator{contract: _IStorage.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IStorage *IStorageFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IStorageTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageTransfer)
				if err := _IStorage.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IStorage *IStorageFilterer) ParseTransfer(log types.Log) (*IStorageTransfer, error) {
	event := new(IStorageTransfer)
	if err := _IStorage.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageUpdateBucketInfoIterator is returned from FilterUpdateBucketInfo and is used to iterate over the raw logs and unpacked data for UpdateBucketInfo events raised by the IStorage contract.
type IStorageUpdateBucketInfoIterator struct {
	Event *IStorageUpdateBucketInfo // Event containing the contract specifics and raw log

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
func (it *IStorageUpdateBucketInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageUpdateBucketInfo)
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
		it.Event = new(IStorageUpdateBucketInfo)
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
func (it *IStorageUpdateBucketInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageUpdateBucketInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageUpdateBucketInfo represents a UpdateBucketInfo event raised by the IStorage contract.
type IStorageUpdateBucketInfo struct {
	Operator       common.Address
	BucketName     [32]byte
	PaymentAddress common.Address
	Visibility     uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateBucketInfo is a free log retrieval operation binding the contract event 0xb576fbee7f104c3265342bd1e54aac187a5a2c2abcd2e66b119b2760dcc34af8.
//
// Solidity: event UpdateBucketInfo(address indexed operator, bytes32 indexed bucketName, address indexed paymentAddress, uint8 visibility)
func (_IStorage *IStorageFilterer) FilterUpdateBucketInfo(opts *bind.FilterOpts, operator []common.Address, bucketName [][32]byte, paymentAddress []common.Address) (*IStorageUpdateBucketInfoIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var bucketNameRule []interface{}
	for _, bucketNameItem := range bucketName {
		bucketNameRule = append(bucketNameRule, bucketNameItem)
	}
	var paymentAddressRule []interface{}
	for _, paymentAddressItem := range paymentAddress {
		paymentAddressRule = append(paymentAddressRule, paymentAddressItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "UpdateBucketInfo", operatorRule, bucketNameRule, paymentAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStorageUpdateBucketInfoIterator{contract: _IStorage.contract, event: "UpdateBucketInfo", logs: logs, sub: sub}, nil
}

// WatchUpdateBucketInfo is a free log subscription operation binding the contract event 0xb576fbee7f104c3265342bd1e54aac187a5a2c2abcd2e66b119b2760dcc34af8.
//
// Solidity: event UpdateBucketInfo(address indexed operator, bytes32 indexed bucketName, address indexed paymentAddress, uint8 visibility)
func (_IStorage *IStorageFilterer) WatchUpdateBucketInfo(opts *bind.WatchOpts, sink chan<- *IStorageUpdateBucketInfo, operator []common.Address, bucketName [][32]byte, paymentAddress []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var bucketNameRule []interface{}
	for _, bucketNameItem := range bucketName {
		bucketNameRule = append(bucketNameRule, bucketNameItem)
	}
	var paymentAddressRule []interface{}
	for _, paymentAddressItem := range paymentAddress {
		paymentAddressRule = append(paymentAddressRule, paymentAddressItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "UpdateBucketInfo", operatorRule, bucketNameRule, paymentAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageUpdateBucketInfo)
				if err := _IStorage.contract.UnpackLog(event, "UpdateBucketInfo", log); err != nil {
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

// ParseUpdateBucketInfo is a log parse operation binding the contract event 0xb576fbee7f104c3265342bd1e54aac187a5a2c2abcd2e66b119b2760dcc34af8.
//
// Solidity: event UpdateBucketInfo(address indexed operator, bytes32 indexed bucketName, address indexed paymentAddress, uint8 visibility)
func (_IStorage *IStorageFilterer) ParseUpdateBucketInfo(log types.Log) (*IStorageUpdateBucketInfo, error) {
	event := new(IStorageUpdateBucketInfo)
	if err := _IStorage.contract.UnpackLog(event, "UpdateBucketInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageUpdateGroupIterator is returned from FilterUpdateGroup and is used to iterate over the raw logs and unpacked data for UpdateGroup events raised by the IStorage contract.
type IStorageUpdateGroupIterator struct {
	Event *IStorageUpdateGroup // Event containing the contract specifics and raw log

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
func (it *IStorageUpdateGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageUpdateGroup)
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
		it.Event = new(IStorageUpdateGroup)
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
func (it *IStorageUpdateGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageUpdateGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageUpdateGroup represents a UpdateGroup event raised by the IStorage contract.
type IStorageUpdateGroup struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateGroup is a free log retrieval operation binding the contract event 0xe6783cc561026c566511a9d9b537069bd52e48ef5766e0220f1ab532dc962b66.
//
// Solidity: event UpdateGroup(address indexed creator)
func (_IStorage *IStorageFilterer) FilterUpdateGroup(opts *bind.FilterOpts, creator []common.Address) (*IStorageUpdateGroupIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "UpdateGroup", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageUpdateGroupIterator{contract: _IStorage.contract, event: "UpdateGroup", logs: logs, sub: sub}, nil
}

// WatchUpdateGroup is a free log subscription operation binding the contract event 0xe6783cc561026c566511a9d9b537069bd52e48ef5766e0220f1ab532dc962b66.
//
// Solidity: event UpdateGroup(address indexed creator)
func (_IStorage *IStorageFilterer) WatchUpdateGroup(opts *bind.WatchOpts, sink chan<- *IStorageUpdateGroup, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "UpdateGroup", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageUpdateGroup)
				if err := _IStorage.contract.UnpackLog(event, "UpdateGroup", log); err != nil {
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

// ParseUpdateGroup is a log parse operation binding the contract event 0xe6783cc561026c566511a9d9b537069bd52e48ef5766e0220f1ab532dc962b66.
//
// Solidity: event UpdateGroup(address indexed creator)
func (_IStorage *IStorageFilterer) ParseUpdateGroup(log types.Log) (*IStorageUpdateGroup, error) {
	event := new(IStorageUpdateGroup)
	if err := _IStorage.contract.UnpackLog(event, "UpdateGroup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageUpdateGroupExtraIterator is returned from FilterUpdateGroupExtra and is used to iterate over the raw logs and unpacked data for UpdateGroupExtra events raised by the IStorage contract.
type IStorageUpdateGroupExtraIterator struct {
	Event *IStorageUpdateGroupExtra // Event containing the contract specifics and raw log

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
func (it *IStorageUpdateGroupExtraIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageUpdateGroupExtra)
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
		it.Event = new(IStorageUpdateGroupExtra)
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
func (it *IStorageUpdateGroupExtraIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageUpdateGroupExtraIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageUpdateGroupExtra represents a UpdateGroupExtra event raised by the IStorage contract.
type IStorageUpdateGroupExtra struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateGroupExtra is a free log retrieval operation binding the contract event 0xaaa1b43dc57fea6bc2db630d9eaa4ae9e8b6d575144857ced396f0b28d349083.
//
// Solidity: event UpdateGroupExtra(address indexed creator)
func (_IStorage *IStorageFilterer) FilterUpdateGroupExtra(opts *bind.FilterOpts, creator []common.Address) (*IStorageUpdateGroupExtraIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "UpdateGroupExtra", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageUpdateGroupExtraIterator{contract: _IStorage.contract, event: "UpdateGroupExtra", logs: logs, sub: sub}, nil
}

// WatchUpdateGroupExtra is a free log subscription operation binding the contract event 0xaaa1b43dc57fea6bc2db630d9eaa4ae9e8b6d575144857ced396f0b28d349083.
//
// Solidity: event UpdateGroupExtra(address indexed creator)
func (_IStorage *IStorageFilterer) WatchUpdateGroupExtra(opts *bind.WatchOpts, sink chan<- *IStorageUpdateGroupExtra, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "UpdateGroupExtra", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageUpdateGroupExtra)
				if err := _IStorage.contract.UnpackLog(event, "UpdateGroupExtra", log); err != nil {
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

// ParseUpdateGroupExtra is a log parse operation binding the contract event 0xaaa1b43dc57fea6bc2db630d9eaa4ae9e8b6d575144857ced396f0b28d349083.
//
// Solidity: event UpdateGroupExtra(address indexed creator)
func (_IStorage *IStorageFilterer) ParseUpdateGroupExtra(log types.Log) (*IStorageUpdateGroupExtra, error) {
	event := new(IStorageUpdateGroupExtra)
	if err := _IStorage.contract.UnpackLog(event, "UpdateGroupExtra", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageUpdateObjectContentIterator is returned from FilterUpdateObjectContent and is used to iterate over the raw logs and unpacked data for UpdateObjectContent events raised by the IStorage contract.
type IStorageUpdateObjectContentIterator struct {
	Event *IStorageUpdateObjectContent // Event containing the contract specifics and raw log

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
func (it *IStorageUpdateObjectContentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageUpdateObjectContent)
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
		it.Event = new(IStorageUpdateObjectContent)
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
func (it *IStorageUpdateObjectContentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageUpdateObjectContentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageUpdateObjectContent represents a UpdateObjectContent event raised by the IStorage contract.
type IStorageUpdateObjectContent struct {
	Operator   common.Address
	ObjectName common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateObjectContent is a free log retrieval operation binding the contract event 0xa0ab3d5251651b77737cb63c98341a07873831b11852369afa7942d97e561ce5.
//
// Solidity: event UpdateObjectContent(address indexed operator, string indexed objectName)
func (_IStorage *IStorageFilterer) FilterUpdateObjectContent(opts *bind.FilterOpts, operator []common.Address, objectName []string) (*IStorageUpdateObjectContentIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var objectNameRule []interface{}
	for _, objectNameItem := range objectName {
		objectNameRule = append(objectNameRule, objectNameItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "UpdateObjectContent", operatorRule, objectNameRule)
	if err != nil {
		return nil, err
	}
	return &IStorageUpdateObjectContentIterator{contract: _IStorage.contract, event: "UpdateObjectContent", logs: logs, sub: sub}, nil
}

// WatchUpdateObjectContent is a free log subscription operation binding the contract event 0xa0ab3d5251651b77737cb63c98341a07873831b11852369afa7942d97e561ce5.
//
// Solidity: event UpdateObjectContent(address indexed operator, string indexed objectName)
func (_IStorage *IStorageFilterer) WatchUpdateObjectContent(opts *bind.WatchOpts, sink chan<- *IStorageUpdateObjectContent, operator []common.Address, objectName []string) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var objectNameRule []interface{}
	for _, objectNameItem := range objectName {
		objectNameRule = append(objectNameRule, objectNameItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "UpdateObjectContent", operatorRule, objectNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageUpdateObjectContent)
				if err := _IStorage.contract.UnpackLog(event, "UpdateObjectContent", log); err != nil {
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

// ParseUpdateObjectContent is a log parse operation binding the contract event 0xa0ab3d5251651b77737cb63c98341a07873831b11852369afa7942d97e561ce5.
//
// Solidity: event UpdateObjectContent(address indexed operator, string indexed objectName)
func (_IStorage *IStorageFilterer) ParseUpdateObjectContent(log types.Log) (*IStorageUpdateObjectContent, error) {
	event := new(IStorageUpdateObjectContent)
	if err := _IStorage.contract.UnpackLog(event, "UpdateObjectContent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageUpdateObjectInfoIterator is returned from FilterUpdateObjectInfo and is used to iterate over the raw logs and unpacked data for UpdateObjectInfo events raised by the IStorage contract.
type IStorageUpdateObjectInfoIterator struct {
	Event *IStorageUpdateObjectInfo // Event containing the contract specifics and raw log

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
func (it *IStorageUpdateObjectInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageUpdateObjectInfo)
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
		it.Event = new(IStorageUpdateObjectInfo)
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
func (it *IStorageUpdateObjectInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageUpdateObjectInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageUpdateObjectInfo represents a UpdateObjectInfo event raised by the IStorage contract.
type IStorageUpdateObjectInfo struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateObjectInfo is a free log retrieval operation binding the contract event 0x36b6cd148113bd68cf4f008f817459149021b90b5ff3ed0ecf51b58c61d620cd.
//
// Solidity: event UpdateObjectInfo(address indexed creator)
func (_IStorage *IStorageFilterer) FilterUpdateObjectInfo(opts *bind.FilterOpts, creator []common.Address) (*IStorageUpdateObjectInfoIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "UpdateObjectInfo", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageUpdateObjectInfoIterator{contract: _IStorage.contract, event: "UpdateObjectInfo", logs: logs, sub: sub}, nil
}

// WatchUpdateObjectInfo is a free log subscription operation binding the contract event 0x36b6cd148113bd68cf4f008f817459149021b90b5ff3ed0ecf51b58c61d620cd.
//
// Solidity: event UpdateObjectInfo(address indexed creator)
func (_IStorage *IStorageFilterer) WatchUpdateObjectInfo(opts *bind.WatchOpts, sink chan<- *IStorageUpdateObjectInfo, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "UpdateObjectInfo", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageUpdateObjectInfo)
				if err := _IStorage.contract.UnpackLog(event, "UpdateObjectInfo", log); err != nil {
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

// ParseUpdateObjectInfo is a log parse operation binding the contract event 0x36b6cd148113bd68cf4f008f817459149021b90b5ff3ed0ecf51b58c61d620cd.
//
// Solidity: event UpdateObjectInfo(address indexed creator)
func (_IStorage *IStorageFilterer) ParseUpdateObjectInfo(log types.Log) (*IStorageUpdateObjectInfo, error) {
	event := new(IStorageUpdateObjectInfo)
	if err := _IStorage.contract.UnpackLog(event, "UpdateObjectInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStorageUpdateParamsIterator is returned from FilterUpdateParams and is used to iterate over the raw logs and unpacked data for UpdateParams events raised by the IStorage contract.
type IStorageUpdateParamsIterator struct {
	Event *IStorageUpdateParams // Event containing the contract specifics and raw log

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
func (it *IStorageUpdateParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageUpdateParams)
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
		it.Event = new(IStorageUpdateParams)
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
func (it *IStorageUpdateParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageUpdateParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageUpdateParams represents a UpdateParams event raised by the IStorage contract.
type IStorageUpdateParams struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateParams is a free log retrieval operation binding the contract event 0xdb2e743561971fc97db45300fcd12c1e9e20b20d485aa6f0259c3374f4a4dafd.
//
// Solidity: event UpdateParams(address indexed creator)
func (_IStorage *IStorageFilterer) FilterUpdateParams(opts *bind.FilterOpts, creator []common.Address) (*IStorageUpdateParamsIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.FilterLogs(opts, "UpdateParams", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IStorageUpdateParamsIterator{contract: _IStorage.contract, event: "UpdateParams", logs: logs, sub: sub}, nil
}

// WatchUpdateParams is a free log subscription operation binding the contract event 0xdb2e743561971fc97db45300fcd12c1e9e20b20d485aa6f0259c3374f4a4dafd.
//
// Solidity: event UpdateParams(address indexed creator)
func (_IStorage *IStorageFilterer) WatchUpdateParams(opts *bind.WatchOpts, sink chan<- *IStorageUpdateParams, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IStorage.contract.WatchLogs(opts, "UpdateParams", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageUpdateParams)
				if err := _IStorage.contract.UnpackLog(event, "UpdateParams", log); err != nil {
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

// ParseUpdateParams is a log parse operation binding the contract event 0xdb2e743561971fc97db45300fcd12c1e9e20b20d485aa6f0259c3374f4a4dafd.
//
// Solidity: event UpdateParams(address indexed creator)
func (_IStorage *IStorageFilterer) ParseUpdateParams(log types.Log) (*IStorageUpdateParams, error) {
	event := new(IStorageUpdateParams)
	if err := _IStorage.contract.UnpackLog(event, "UpdateParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
