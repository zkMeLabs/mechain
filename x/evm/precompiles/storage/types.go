package storage

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	storageAddress = common.HexToAddress(types.StorageAddress)
	storageABI     = types.MustABIJson(IStorageMetaData.ABI)
	invalidMethod  = abi.Method{}
)

const (
	noBalanceErr = "key not found"

	BucketResourcePrefix = "grn:b::"
	ObjectResourcePrefix = "grn:o::"
	GroupResourcePrefix  = "grn:g:"

	ObjectResourceType = 1
	BucketResourceType = 2
	GroupResourceType  = 3
)

type ResourceType int
type NewStatementOptions struct {
	StatementExpireTime *time.Time
	LimitSize           uint64
}

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

type HeadBucketExtraArgs struct {
	BucketName string `abi:"bucketName"`
}

// Validate HeadBucketExtraArgs the args
func (args *HeadBucketExtraArgs) Validate() error {
	return nil
}

type HeadBucketByIdArgs struct {
	BucketId string `abi:"bucketId"`
}

// Validate HeadBucketByIdArgs the args
func (args *HeadBucketByIdArgs) Validate() error {
	return nil
}

type HeadBucketNFTArgs struct {
	TokenId string `abi:"tokenId"`
}

// Validate HeadBucketNFTArgs the args
func (args *HeadBucketNFTArgs) Validate() error {
	return nil
}

type HeadObjectNFTArgs struct {
	TokenId string `abi:"tokenId"`
}

// Validate HeadObjectNFTArgs the args
func (args *HeadObjectNFTArgs) Validate() error {
	return nil
}

type HeadGroupNFTArgs struct {
	TokenId string `abi:"tokenId"`
}

// Validate HeadGroupNFTArgs the args
func (args *HeadGroupNFTArgs) Validate() error {
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

type MigrateBucketArgs struct {
	// Operator             string       `abi:"operator"`
	BucketName           string       `abi:"bucketName"`
	DstPrimarySpId       uint32       `abi:"dstPrimarySpId"`
	DstPrimarySpApproval ApprovalJSON `abi:"dstPrimarySpApproval"`
}

// Validate MigrateBucketArgs args
func (args *MigrateBucketArgs) Validate() error {
	return nil
}

type CompleteMigrateBucketArgs struct {
	BucketName  string       `abi:"bucketName"`
	GvgFamilyId uint32       `abi:"gvgFamilyId"`
	GvgMappings []GVGMapping `abi:"gvgMappings"`
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

type CancelMigrateBucketArgs struct {
	BucketName string `abi:"bucketName"`
}

// Validate CancelMigrateBucketArgs args
func (args *CancelMigrateBucketArgs) Validate() error {
	return nil
}

type SetBucketFlowRateLimitArgs struct {
	BucketName     string   `abi:"bucketName"`
	BucketOwner    string   `abi:"bucketOwner"`
	PaymentAddress string   `abi:"paymentAddress"`
	FlowRateLimit  *big.Int `abi:"flowRateLimit"`
}

// Validate SetBucketFlowRateLimitArgs args
func (args *SetBucketFlowRateLimitArgs) Validate() error {
	return nil
}

type MirrorBucketArgs struct {
	BucketId    *big.Int `abi:"bucketId"`
	BucketName  string   `abi:"bucketName"`
	DestChainId uint32   `abi:"destChainId"`
}

// Validate MirrorBucketArgs args
func (args *MirrorBucketArgs) Validate() error {
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

type CopyObjectArgs struct {
	// Operator             string       `abi:"operator"`
	SrcBucketName        string       `abi:"srcBucketName"`
	DstBucketName        string       `abi:"dstBucketName"`
	SrcObjectName        string       `abi:"srcObjectName"`
	DstObjectName        string       `abi:"dstObjectName"`
	DstPrimarySpApproval ApprovalJSON `abi:"dstPrimarySpApproval"`
}

// Validate CopyObjectArgs args
func (args *CopyObjectArgs) Validate() error {
	return nil
}

type DeleteObjectArgs struct {
	// Operator   string `abi:"operator"`
	BucketName string `abi:"bucketName"`
	ObjectName string `abi:"objectName"`
}

// Validate DeleteObjectArgs args
func (args *DeleteObjectArgs) Validate() error {
	return nil
}

type CancelCreateObjectArgs struct {
	// Operator   string `abi:"operator"`
	BucketName string `abi:"bucketName"`
	ObjectName string `abi:"objectName"`
}

// Validate CancelCreateObjectArgs args
func (args *CancelCreateObjectArgs) Validate() error {
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

type ListObjectsByBucketIdArgs struct {
	Pagination PageRequestJSON `abi:"pagination"`
	BucketId   string          `abi:"bucketId"`
}

// Validate ListObjectsByBucketIdArgs the args
func (args *ListObjectsByBucketIdArgs) Validate() error {
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
	// Operator        string   `abi:"operator"`
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

type UpdateObjectContentArgs struct {
	BucketName      string   `abi:"bucketName"`
	ObjectName      string   `abi:"objectName"`
	PayloadSize     uint64   `abi:"payloadSize"`
	ContentType     string   `abi:"contentType"`
	ExpectChecksums []string `abi:"expectChecksums"`
}

// Validate UpdateObjectInfoArgs args
func (args *UpdateObjectContentArgs) Validate() error {
	return nil
}

type DiscontinueObjectArgs struct {
	// Operator   string     `abi:"operator"`
	BucketName string     `abi:"bucketName"`
	ObjectIds  []*big.Int `abi:"objectIds"`
	Reason     string     `abi:"reason"`
}

// Validate DiscontinueObjectArgs args
func (args *DiscontinueObjectArgs) Validate() error {
	return nil
}

type MirrorObjectArgs struct {
	ObjectId    *big.Int `abi:"objectId"`
	BucketName  string   `abi:"bucketName"`
	ObjectName  string   `abi:"objectName"`
	DestChainId uint32   `abi:"destChainId"`
}

// Validate MirrorObjectArgs args
func (args *MirrorObjectArgs) Validate() error {
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

type UpdateGroupExtraArgs struct {
	// Operator string `abi:"operator"`
	GroupOwner common.Address `abi:"groupOwner"`
	GroupName  string         `abi:"groupName"`
	Extra      string         `abi:"extra"`
}

// Validate UpdateGroupExtraArgs the args
func (args *UpdateGroupExtraArgs) Validate() error {
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

type LeaveGroupArgs struct {
	Member     common.Address `abi:"member"`
	GroupOwner common.Address `abi:"groupOwner"`
	GroupName  string         `abi:"groupName"`
}

// Validate LeaveGroupArgs the args
func (args *LeaveGroupArgs) Validate() error {
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

type MirrorGroupArgs struct {
	GroupId     *big.Int `abi:"groupId"`
	GroupName   string   `abi:"groupName"`
	DestChainId uint32   `abi:"destChainId"`
}

// Validate MirrorGroupArgs args
func (args *MirrorGroupArgs) Validate() error {
	return nil
}

type ToggleSPAsDelegatedAgentArgs struct {
	BucketName string `abi:"bucketName"`
}

// Validate ToggleSPAsDelegatedAgentArgs args
func (args *ToggleSPAsDelegatedAgentArgs) Validate() error {
	return nil
}

type UpdateParamsArgs struct {
	Authority string `abi:"authority"`
	Params    Params `abi:"params"`
}

// Validate UpdateParamsArgs args
func (args *UpdateParamsArgs) Validate() error {
	return nil
}

type SetTagArgs struct {
	Resource string `abi:"resource"`
	Tags     []Tag  `abi:"tags"`
}

// Validate SetTagArgs the args
func (args *SetTagArgs) Validate() error {
	if args.Tags == nil {
		return errors.New("invalid tags parameter")
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

type HeadShadowObjectArgs struct {
	BucketName string `abi:"bucketName"`
	ObjectName string `abi:"objectName"`
}

// Validate HeadShadowObjectArgs the args
func (args *HeadShadowObjectArgs) Validate() error {
	if args.BucketName == "" {
		return errors.New("empty bucket name")
	}
	if args.ObjectName == "" {
		return errors.New("empty object name")
	}
	return nil
}

type PutPolicyArgs struct {
	// Operator       string      `abi:"operator"`
	Principal      Principal   `abi:"principal"`
	Resource       string      `abi:"resource"`
	Statements     []Statement `abi:"statements"`
	ExpirationTime int64       `abi:"expirationTime"`
}

// Validate PutPolicyArgs the args
func (args *PutPolicyArgs) Validate() error {
	return nil
}

type DeletePolicyArgs struct {
	// Operator  string    `abi:"operator"`
	Principal Principal `abi:"principal"`
	Resource  string    `abi:"resource"`
}

// Validate DeletePolicyArgs the args
func (args *DeletePolicyArgs) Validate() error {
	return nil
}

type QueryPolicyForGroupArgs struct {
	Resource string   `abi:"resource"`
	GroupId  *big.Int `abi:"groupId"`
}

// Validate QueryPolicyForGroupArgs the args
func (args *QueryPolicyForGroupArgs) Validate() error {
	return nil
}

type QueryPolicyForAccountArgs struct {
	Resource      string `abi:"resource"`
	PrincipalAddr string `abi:"principalAddr"`
}

// Validate QueryPolicyForAccountArgs the args
func (args *QueryPolicyForAccountArgs) Validate() error {
	return nil
}

type QueryPolicyByIdArgs struct {
	PolicyId string `abi:"policyId"`
}

// Validate QueryPolicyByIdArgs the args
func (args *QueryPolicyByIdArgs) Validate() error {
	return nil
}

type QueryLockFeeArgs struct {
	PrimarySpAddress string `abi:"primarySpAddress"`
	CreateAt         int64  `abi:"createAt"`
	PayloadSize      uint64 `abi:"payloadSize"`
}

// Validate QueryLockFeeArgs the args
func (args *QueryLockFeeArgs) Validate() error {
	return nil
}

type QueryIsPriceChangedArgs struct {
	BucketName string `abi:"bucketName"`
}

// Validate QueryIsPriceChangedArgs the args
func (args *QueryIsPriceChangedArgs) Validate() error {
	return nil
}

type QueryQuotaUpdateTimeArgs struct {
	BucketName string `abi:"bucketName"`
}

// Validate QueryQuotaUpdateTimeArgs the args
func (args *QueryQuotaUpdateTimeArgs) Validate() error {
	return nil
}

type QueryGroupMembersExistArgs struct {
	GroupId string   `abi:"groupId"`
	Members []string `abi:"members"`
}

// Validate QueryGroupMembersExistArgs the args
func (args *QueryGroupMembersExistArgs) Validate() error {
	return nil
}

type QueryGroupsExistArgs struct {
	GroupOwner string   `abi:"groupOwner"`
	GroupNames []string `abi:"groupNames"`
}

// Validate QueryGroupsExistArgs the args
func (args *QueryGroupsExistArgs) Validate() error {
	return nil
}

type QueryGroupsExistByIdArgs struct {
	GroupIds []string `abi:"groupIds"`
}

// Validate QueryGroupsExistByIdArgs the args
func (args *QueryGroupsExistByIdArgs) Validate() error {
	return nil
}

type QueryPaymentAccountBucketFlowRateLimitArgs struct {
	PaymentAccount string `abi:"paymentAccount"`
	BucketOwner    string `abi:"bucketOwner"`
	BucketName     string `abi:"bucketName"`
}

// Validate QueryPaymentAccountBucketFlowRateLimitArgs the args
func (args *QueryPaymentAccountBucketFlowRateLimitArgs) Validate() error {
	return nil
}

type QueryParamsByTimestampArgs struct {
	Timestamp int64 `abi:"timestamp"`
}

// Validate QueryParamsByTimestampArgs the args
func (args *QueryParamsByTimestampArgs) Validate() error {
	return nil
}

type VerifyPermissionArgs struct {
	BucketName string `abi:"bucketName"`
	ObjectName string `abi:"objectName"`
	ActionType int32  `abi:"actionType"`
}

// Validate VerifyPermissionArgs the args
func (args *VerifyPermissionArgs) Validate() error {
	return nil
}
