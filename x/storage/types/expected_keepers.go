package types

import (
	"context"
	"math/big"
	time "time"

	"cosmossdk.io/math"
	sdkmath "cosmossdk.io/math"
	"github.com/cometbft/cometbft/libs/log"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/types/resource"
	"github.com/evmos/evmos/v12/x/evm/statedb"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
	paymenttypes "github.com/evmos/evmos/v12/x/payment/types"
	permtypes "github.com/evmos/evmos/v12/x/permission/types"
	sptypes "github.com/evmos/evmos/v12/x/sp/types"
	vgtypes "github.com/evmos/evmos/v12/x/virtualgroup/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetSequence(sdktypes.Context, sdktypes.AccAddress) (uint64, error)
	GetAccount(ctx sdktypes.Context, addr sdktypes.AccAddress) authtypes.AccountI
	GetModuleAddress(name string) sdktypes.AccAddress
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdktypes.Context, addr sdktypes.AccAddress) sdktypes.Coins
	GetBalance(ctx sdktypes.Context, addr sdktypes.AccAddress, denom string) sdktypes.Coin
	GetAllBalances(ctx sdktypes.Context, addr sdktypes.AccAddress) sdktypes.Coins
	SendCoinsFromModuleToAccount(ctx sdktypes.Context, senderModule string, recipientAddr sdktypes.AccAddress, amt sdktypes.Coins) error
	// Methods imported from bank should be defined here
}

type SpKeeper interface {
	GetStorageProvider(ctx sdktypes.Context, id uint32) (*sptypes.StorageProvider, bool)
	MustGetStorageProvider(ctx sdktypes.Context, id uint32) *sptypes.StorageProvider
	GetStorageProviderByOperatorAddr(ctx sdktypes.Context, addr sdktypes.AccAddress) (sp *sptypes.StorageProvider, found bool)
	GetStorageProviderBySealAddr(ctx sdktypes.Context, sealAddr sdktypes.AccAddress) (sp *sptypes.StorageProvider, found bool)
	GetStorageProviderByGcAddr(ctx sdktypes.Context, gcAddr sdktypes.AccAddress) (sp *sptypes.StorageProvider, found bool)
	GetGlobalSpStorePriceByTime(ctx sdktypes.Context, time int64) (val sptypes.GlobalSpStorePrice, err error)
}

type PaymentKeeper interface {
	GetVersionedParamsWithTs(ctx sdktypes.Context, time int64) (paymenttypes.VersionedParams, error)
	IsPaymentAccountOwner(ctx sdktypes.Context, addr, owner sdktypes.AccAddress) bool
	ApplyUserFlowsList(ctx sdktypes.Context, userFlows []paymenttypes.UserFlows) (err error)
	UpdateStreamRecordByAddr(ctx sdktypes.Context, change *paymenttypes.StreamRecordChange) (ret *paymenttypes.StreamRecord, err error)
	GetStreamRecord(ctx sdktypes.Context, account sdktypes.AccAddress) (ret *paymenttypes.StreamRecord, found bool)
	MergeOutFlows(flows []paymenttypes.OutFlow) []paymenttypes.OutFlow
	GetAllStreamRecord(ctx sdktypes.Context) (list []paymenttypes.StreamRecord)
	GetOutFlows(ctx sdktypes.Context, addr sdktypes.AccAddress) []paymenttypes.OutFlow
}

type PermissionKeeper interface {
	PutPolicy(ctx sdktypes.Context, policy *permtypes.Policy) (math.Uint, error)
	DeletePolicy(ctx sdktypes.Context, principal *permtypes.Principal, resourceType resource.ResourceType,
		resourceID math.Uint) (math.Uint, error)
	AddGroupMember(ctx sdktypes.Context, groupID math.Uint, member sdktypes.AccAddress, expiration *time.Time) error
	UpdateGroupMember(ctx sdktypes.Context, groupID math.Uint, member sdktypes.AccAddress, memberID math.Uint, expiration *time.Time)
	MustGetPolicyByID(ctx sdktypes.Context, policyID math.Uint) *permtypes.Policy
	GetPolicyGroupForResource(ctx sdktypes.Context, resourceID math.Uint, resourceType resource.ResourceType) (*permtypes.PolicyGroup, bool)
	RemoveGroupMember(ctx sdktypes.Context, groupID math.Uint, member sdktypes.AccAddress) error
	GetPolicyByID(ctx sdktypes.Context, policyID math.Uint) (*permtypes.Policy, bool)
	GetPolicyForAccount(ctx sdktypes.Context, resourceID math.Uint, resourceType resource.ResourceType, addr sdktypes.AccAddress) (policy *permtypes.Policy, isFound bool)
	GetPolicyForGroup(ctx sdktypes.Context, resourceID math.Uint, resourceType resource.ResourceType,
		groupID math.Uint) (policy *permtypes.Policy, isFound bool)
	GetGroupMember(ctx sdktypes.Context, groupID math.Uint, member sdktypes.AccAddress) (*permtypes.GroupMember, bool)
	GetGroupMemberByID(ctx sdktypes.Context, groupMemberID math.Uint) (*permtypes.GroupMember, bool)
	ForceDeleteAccountPolicyForResource(ctx sdktypes.Context, maxDelete, deletedCount uint64, resourceType resource.ResourceType, resourceID math.Uint) (uint64, bool)
	ForceDeleteGroupPolicyForResource(ctx sdktypes.Context, maxDelete, deletedCount uint64, resourceType resource.ResourceType, resourceID math.Uint) (uint64, bool)
	ForceDeleteGroupMembers(ctx sdktypes.Context, maxDelete, deletedTotal uint64, groupId math.Uint) (uint64, bool)
	ExistAccountPolicyForResource(ctx sdktypes.Context, resourceType resource.ResourceType, resourceID math.Uint) bool
	ExistGroupPolicyForResource(ctx sdktypes.Context, resourceType resource.ResourceType, resourceID math.Uint) bool
	ExistGroupMemberForGroup(ctx sdktypes.Context, groupId math.Uint) bool
}

type CrossChainKeeper interface {
	GetDestBscChainID() sdktypes.ChainID
	GetDestOpChainID() sdktypes.ChainID

	CreateRawIBCPackageWithFee(ctx sdktypes.Context, chainID sdktypes.ChainID, channelID sdktypes.ChannelID, packageType sdktypes.CrossChainPackageType,
		packageLoad []byte, relayerFee *big.Int, ackRelayerFee *big.Int,
	) (uint64, error)

	IsDestChainSupported(chainID sdktypes.ChainID) bool

	RegisterChannel(name string, id sdktypes.ChannelID, app sdktypes.CrossChainApplication) error
}

type VirtualGroupKeeper interface {
	SetGVGAndEmitUpdateEvent(ctx sdktypes.Context, gvg *vgtypes.GlobalVirtualGroup) error
	GetGVGFamily(ctx sdktypes.Context, familyID uint32) (*vgtypes.GlobalVirtualGroupFamily, bool)
	GetGVG(ctx sdktypes.Context, gvgID uint32) (*vgtypes.GlobalVirtualGroup, bool)
	SettleAndDistributeGVGFamily(ctx sdktypes.Context, sp *sptypes.StorageProvider, family *vgtypes.GlobalVirtualGroupFamily) error
	SettleAndDistributeGVG(ctx sdktypes.Context, gvg *vgtypes.GlobalVirtualGroup) error
	GetAndCheckGVGFamilyAvailableForNewBucket(ctx sdktypes.Context, familyID uint32) (*vgtypes.GlobalVirtualGroupFamily, error)
	GetGlobalVirtualGroupIfAvailable(ctx sdktypes.Context, gvgID uint32, expectedStoreSize uint64) (*vgtypes.GlobalVirtualGroup, error)
	GetSwapInInfo(ctx sdktypes.Context, familyID, gvgID uint32) (*vgtypes.SwapInInfo, bool)
}

// StorageKeeper used by the cross-chain applications
type StorageKeeper interface {
	Logger(ctx sdktypes.Context) log.Logger
	GetBucketInfoById(ctx sdktypes.Context, bucketId sdkmath.Uint) (*BucketInfo, bool)
	SetBucketInfo(ctx sdktypes.Context, bucketInfo *BucketInfo)
	CreateBucket(
		ctx sdktypes.Context, ownerAcc sdktypes.AccAddress, bucketName string,
		primarySpAcc sdktypes.AccAddress, opts *CreateBucketOptions) (sdkmath.Uint, error)
	DeleteBucket(ctx sdktypes.Context, operator sdktypes.AccAddress, bucketName string, opts DeleteBucketOptions) error
	GetGroupInfoById(ctx sdktypes.Context, groupId sdkmath.Uint) (*GroupInfo, bool)
	GetGroupInfo(ctx sdktypes.Context, ownerAddr sdktypes.AccAddress, groupName string) (*GroupInfo, bool)
	DeleteGroup(ctx sdktypes.Context, operator sdktypes.AccAddress, groupName string, opts DeleteGroupOptions) error
	CreateGroup(
		ctx sdktypes.Context, owner sdktypes.AccAddress,
		groupName string, opts CreateGroupOptions) (sdkmath.Uint, error)
	SetGroupInfo(ctx sdktypes.Context, groupInfo *GroupInfo)
	UpdateGroupMember(ctx sdktypes.Context, operator sdktypes.AccAddress, groupInfo *GroupInfo, opts UpdateGroupMemberOptions) error
	RenewGroupMember(ctx sdktypes.Context, operator sdktypes.AccAddress, groupInfo *GroupInfo, opts RenewGroupMemberOptions) error
	GetObjectInfoById(ctx sdktypes.Context, objectId sdkmath.Uint) (*ObjectInfo, bool)
	SetObjectInfo(ctx sdktypes.Context, objectInfo *ObjectInfo)
	DeleteObject(
		ctx sdktypes.Context, operator sdktypes.AccAddress, bucketName, objectName string, opts DeleteObjectOptions) error
	GetSourceTypeByChainId(ctx sdktypes.Context, chainId sdktypes.ChainID) (SourceType, error)

	NormalizePrincipal(ctx sdktypes.Context, principal *permtypes.Principal)
	ValidatePrincipal(ctx sdktypes.Context, resOwner sdktypes.AccAddress, principal *permtypes.Principal) error
}

type PaymentMsgServer interface {
	CreatePaymentAccount(context.Context, *paymenttypes.MsgCreatePaymentAccount) (*paymenttypes.MsgCreatePaymentAccountResponse, error)
	Deposit(context.Context, *paymenttypes.MsgDeposit) (*paymenttypes.MsgDepositResponse, error)
	Withdraw(context.Context, *paymenttypes.MsgWithdraw) (*paymenttypes.MsgWithdrawResponse, error)
	DisableRefund(context.Context, *paymenttypes.MsgDisableRefund) (*paymenttypes.MsgDisableRefundResponse, error)
}

type StorageMsgServer interface {
	UpdateBucketInfo(context.Context, *MsgUpdateBucketInfo) (*MsgUpdateBucketInfoResponse, error)
	ToggleSPAsDelegatedAgent(context.Context, *MsgToggleSPAsDelegatedAgent) (*MsgToggleSPAsDelegatedAgentResponse, error)
	CopyObject(context.Context, *MsgCopyObject) (*MsgCopyObjectResponse, error)
	UpdateObjectInfo(context.Context, *MsgUpdateObjectInfo) (*MsgUpdateObjectInfoResponse, error)
	UpdateGroupExtra(context.Context, *MsgUpdateGroupExtra) (*MsgUpdateGroupExtraResponse, error)
	MigrateBucket(context.Context, *MsgMigrateBucket) (*MsgMigrateBucketResponse, error)
	CancelMigrateBucket(context.Context, *MsgCancelMigrateBucket) (*MsgCancelMigrateBucketResponse, error)
	SetTag(context.Context, *MsgSetTag) (*MsgSetTagResponse, error)
	SetBucketFlowRateLimit(context.Context, *MsgSetBucketFlowRateLimit) (*MsgSetBucketFlowRateLimitResponse, error)
}

// EVMKeeper defines the expected EVM keeper interface used on erc20
type EVMKeeper interface {
	GetParams(ctx sdktypes.Context) evmtypes.Params
	GetAccountWithoutBalance(ctx sdktypes.Context, addr common.Address) *statedb.Account
	EstimateGas(c context.Context, req *evmtypes.EthCallRequest) (*evmtypes.EstimateGasResponse, error)
	ApplyMessage(ctx sdktypes.Context, msg core.Message, tracer vm.EVMLogger, commit bool) (*evmtypes.MsgEthereumTxResponse, error)
}
