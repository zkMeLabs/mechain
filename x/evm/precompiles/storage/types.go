package storage

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	storageAddress = common.HexToAddress(types.StorageAddress)
	storageABI     = types.MustABIJson(IStorageMetaData.ABI)
	invalidMethod  = abi.Method{}
)

func GetAddress() common.Address {
	return storageAddress
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return invalidMethod, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range storageABI.Methods {
		if bytes.Equal(input[:4], method.ID) {
			return method, nil
		}
	}
	return invalidMethod, fmt.Errorf("method id %s is not exist", string(input[:4]))
}

func GetAbiMethod(name string) abi.Method {
	return storageABI.Methods[name]
}

func GetAbiEvent(name string) abi.Event {
	return storageABI.Events[name]
}

type (
	ApprovalJSON    = Approval
	PageRequestJSON = PageRequest
)

type CreateBucketArgs struct {
	BucketName        string         `abi:"bucketName"`
	Visibility        uint8          `abi:"visibility"`
	PaymentAddress    common.Address `abi:"paymentAddress"`
	PrimarySpAddress  common.Address `abi:"primarySpAddress"`
	PrimarySpApproval ApprovalJSON   `abi:"primarySpApproval"`
	ChargedReadQuota  uint64         `abi:"chargedReadQuota"`
}

// Validate CreateBucketArgs args
func (args *CreateBucketArgs) Validate() error {
	return nil
}

type UpdateBucketInfoArgs struct {
	BucketName       string         `abi:"bucketName"`
	ChargedReadQuota *big.Int       `abi:"chargedReadQuota"`
	PaymentAddress   common.Address `abi:"paymentAddress"`
	Visibility       uint8          `abi:"visibility"`
}

// Validate CreateBucketArgs args
func (args *UpdateBucketInfoArgs) Validate() error {
	if args.BucketName == "" {
		return errors.New("empty bucket name")
	}

	if args.ChargedReadQuota.Int64() != -1 && !args.ChargedReadQuota.IsUint64() {
		return errors.New("charged read quota is invalid")
	}
	return nil
}

type ListBucketsArgs struct {
	Pagination PageRequestJSON `abi:"pagination"`
}

// Validate ListBucketsArgs the args
func (args *ListBucketsArgs) Validate() error {
	return nil
}

type HeadBucketArgs struct {
	BucketName string `abi:"bucketName"`
}

// Validate HeadBucketArgs the args
func (args *HeadBucketArgs) Validate() error {
	return nil
}

type DeleteBucketArgs struct {
	BucketName string `abi:"bucketName"`
}

// Validate DeleteBucketArgs args
func (args *DeleteBucketArgs) Validate() error {
	return nil
}

type DiscontinueBucketArgs struct {
	BucketName string `abi:"bucketName"`
	Reason     string `abi:"reason"`
}

// Validate DiscontinueBucketArgs args
func (args *DiscontinueBucketArgs) Validate() error {
	return nil
}

type CompleteMigrateBucketArgs struct {
	BucketName                 string       `abi:"bucketName"`
	GlobalVirtualGroupFamilyId uint32       `abi:"globalVirtualGroupFamilyId"`
	GvgMappings                []GVGMapping `abi:"gvgMappings"`
}

// Validate CompleteMigrateBucketArgs args
func (args *CompleteMigrateBucketArgs) Validate() error {
	return nil
}

type RejectMigrateBucketArgs struct {
	BucketName string `abi:"bucketName"`
}

// Validate RejectMigrateBucketArgs args
func (args *RejectMigrateBucketArgs) Validate() error {
	return nil
}

type CreateObjectArgs struct {
	BucketName        string       `abi:"bucketName"`
	ObjectName        string       `abi:"objectName"`
	PayloadSize       uint64       `abi:"payloadSize"`
	Visibility        uint8        `abi:"visibility"`
	ContentType       string       `abi:"contentType"`
	PrimarySpApproval ApprovalJSON `abi:"primarySpApproval"`
	ExpectChecksums   []string     `abi:"expectChecksums"`
	RedundancyType    uint8        `abi:"redundancyType"`
}

// Validate CreateObjectArgs args
func (args *CreateObjectArgs) Validate() error {
	return nil
}

type ListObjectsArgs struct {
	Pagination PageRequestJSON `abi:"pagination"`
	BucketName string          `abi:"bucketName"`
}

// Validate ListObjectsArgs the args
func (args *ListObjectsArgs) Validate() error {
	return nil
}

type SealObjectArgs struct {
	SealAddress                 common.Address `abi:"sealAddress"` // primary sp's operator addr or secondary sp's seal addr
	BucketName                  string         `abi:"bucketName"`
	ObjectName                  string         `abi:"objectName"`
	GlobalVirtualGroupID        uint32         `abi:"globalVirtualGroupId"`
	SecondarySpBlsAggSignatures string         `abi:"secondarySpBlsAggSignatures"`
}

// Validate SealObjectArgs args
func (args *SealObjectArgs) Validate() error {
	return nil
}

type SealObjectV2Args struct {
	SealAddress                 common.Address `abi:"sealAddress"` // primary sp's operator addr or secondary sp's seal addr
	BucketName                  string         `abi:"bucketName"`
	ObjectName                  string         `abi:"objectName"`
	GlobalVirtualGroupID        uint32         `abi:"globalVirtualGroupId"`
	SecondarySpBlsAggSignatures string         `abi:"secondarySpBlsAggSignatures"`
	ExpectChecksums             []string       `abi:"expectChecksums"`
}

// Validate SealObjectV2Args args
func (args *SealObjectV2Args) Validate() error {
	return nil
}

type RejectSealObjectArgs struct {
	BucketName string `abi:"bucketName"`
	ObjectName string `abi:"objectName"`
}

// Validate RejectSealObjectArgs args
func (args *RejectSealObjectArgs) Validate() error {
	return nil
}

type DelegateCreateObjectArgs struct {
	Creator         string   `abi:"creator"`
	BucketName      string   `abi:"bucketName"`
	ObjectName      string   `abi:"objectName"`
	PayloadSize     uint64   `abi:"payloadSize"`
	ContentType     string   `abi:"contentType"`
	Visibility      uint8    `abi:"visibility"`
	ExpectChecksums []string `abi:"expectChecksums"`
	RedundancyType  uint8    `abi:"redundancyType"`
}

// Validate DelegateCreateObjectArgs args
func (args *DelegateCreateObjectArgs) Validate() error {
	return nil
}

type DelegateUpdateObjectContentArgs struct {
	Operator        string   `abi:"operator"`
	Updater         string   `abi:"updater"`
	BucketName      string   `abi:"bucketName"`
	ObjectName      string   `abi:"objectName"`
	PayloadSize     uint64   `abi:"payloadSize"`
	ContentType     string   `abi:"contentType"`
	ExpectChecksums []string `abi:"expectChecksums"`
}

// Validate DelegateUpdateObjectContentArgs args
func (args *DelegateUpdateObjectContentArgs) Validate() error {
	return nil
}

type UpdateObjectInfoArgs struct {
	BucketName string `abi:"bucketName"`
	ObjectName string `abi:"objectName"`
	Visibility uint8  `abi:"visibility"`
}

// Validate UpdateObjectInfoArgs args
func (args *UpdateObjectInfoArgs) Validate() error {
	return nil
}

type CreateGroupArgs struct {
	GroupName string `abi:"groupName"`
	Extra     string `abi:"extra"`
}

// Validate CreateGroupArgs args
func (args *CreateGroupArgs) Validate() error {
	return nil
}

type ListGroupsArgs struct {
	Pagination PageRequestJSON `abi:"pagination"`
	GroupOwner common.Address  `abi:"groupOwner"`
}

// Validate ListGroupsArgs the args
func (args *ListGroupsArgs) Validate() error {
	return nil
}

type UpdateGroupArgs struct {
	GroupOwner      common.Address   `abi:"groupOwner"`
	GroupName       string           `abi:"groupName"`
	MembersToAdd    []common.Address `abi:"membersToAdd"`
	ExpirationTime  []int64          `abi:"expirationTime"`
	MembersToDelete []common.Address `abi:"membersToDelete"`
}

// Validate UpdateGroupArgs the args
func (args *UpdateGroupArgs) Validate() error {
	if args.GroupName == "" {
		return errors.New("group name is empty")
	}
	if len(args.MembersToAdd) == 0 && len(args.MembersToDelete) == 0 {
		return errors.New("no update member")
	}
	if args.ExpirationTime != nil && len(args.MembersToAdd) != len(args.ExpirationTime) {
		return errors.New("please provide expirationTime for every new add member")
	}
	return nil
}

type HeadGroupArgs struct {
	GroupOwner common.Address `abi:"groupOwner"`
	GroupName  string         `abi:"groupName"`
}

// Validate HeadGroupArgs the args
func (args *HeadGroupArgs) Validate() error {
	return nil
}

type DeleteGroupArgs struct {
	GroupName string `abi:"groupName"`
}

// Validate DeleteGroupArgs the args
func (args *DeleteGroupArgs) Validate() error {
	return nil
}

type HeadGroupMemberArgs struct {
	Member     common.Address `abi:"member"`
	GroupOwner common.Address `abi:"groupOwner"`
	GroupName  string         `abi:"groupName"`
}

// Validate HeadGroupMemberArgs the args
func (args *HeadGroupMemberArgs) Validate() error {
	return nil
}

type RenewGroupMemberArgs struct {
	GroupOwner     common.Address   `abi:"groupOwner"`
	GroupName      string           `abi:"groupName"`
	Members        []common.Address `abi:"members"`
	ExpirationTime []int64          `abi:"expirationTime"`
}

// Validate RenewGroupMemberArgs the args
func (args *RenewGroupMemberArgs) Validate() error {
	return nil
}

type SetTagForGroupArgs struct {
	GroupName string `abi:"groupName"`
	Tags      []Tag  `abi:"tags"`
}

// Validate SetTagForGroupArgs the args
func (args *SetTagForGroupArgs) Validate() error {
	if args.Tags == nil {
		return errors.New("invalid tags parameter")
	}
	if args.GroupName == "" {
		return errors.New("group name is empty")
	}
	return nil
}

type HeadObjectArgs struct {
	BucketName string `abi:"bucketName"`
	ObjectName string `abi:"objectName"`
}

// Validate HeadObjectArgs the args
func (args *HeadObjectArgs) Validate() error {
	if args.BucketName == "" {
		return errors.New("empty bucket name")
	}
	if args.ObjectName == "" {
		return errors.New("empty object name")
	}
	return nil
}

type HeadObjectByIDArgs struct {
	ObjectID string `abi:"objectId"`
}

// Validate HeadObjectByIdArgs the args
func (args *HeadObjectByIDArgs) Validate() error {
	if args.ObjectID == "" {
		return errors.New("empty object id")
	}
	return nil
}
