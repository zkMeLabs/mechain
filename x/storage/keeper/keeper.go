package keeper

import (
	"encoding/binary"
	"fmt"

	"cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	ecommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/evmos/v12/contracts"
	"github.com/evmos/evmos/v12/internal/sequence"
	"github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/types/common"
	gnfderrors "github.com/evmos/evmos/v12/types/errors"
	"github.com/evmos/evmos/v12/types/resource"
	paymenttypes "github.com/evmos/evmos/v12/x/payment/types"
	permtypes "github.com/evmos/evmos/v12/x/permission/types"
	sptypes "github.com/evmos/evmos/v12/x/sp/types"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
	virtualgroupmoduletypes "github.com/evmos/evmos/v12/x/virtualgroup/types"
)

type (
	Keeper struct {
		cdc                codec.BinaryCodec
		storeKey           storetypes.StoreKey
		tStoreKey          storetypes.StoreKey
		spKeeper           storagetypes.SpKeeper
		paymentKeeper      storagetypes.PaymentKeeper
		accountKeeper      storagetypes.AccountKeeper
		permKeeper         storagetypes.PermissionKeeper
		crossChainKeeper   storagetypes.CrossChainKeeper
		virtualGroupKeeper storagetypes.VirtualGroupKeeper
		evmKeeper          storagetypes.EVMKeeper
		// sequence
		bucketSeq sequence.Sequence[sdkmath.Uint]
		objectSeq sequence.Sequence[sdkmath.Uint]
		groupSeq  sequence.Sequence[sdkmath.Uint]

		authority string

		// payment check config
		cfg *paymentCheckConfig
	}
)

type paymentCheckConfig struct {
	Enabled  bool
	Interval uint32
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	tStoreKey storetypes.StoreKey,
	accountKeeper storagetypes.AccountKeeper,
	spKeeper storagetypes.SpKeeper,
	paymentKeeper storagetypes.PaymentKeeper,
	permKeeper storagetypes.PermissionKeeper,
	crossChainKeeper storagetypes.CrossChainKeeper,
	virtualGroupKeeper storagetypes.VirtualGroupKeeper,
	evmKeeper storagetypes.EVMKeeper,
	authority string,
) *Keeper {
	k := Keeper{
		cdc:                cdc,
		storeKey:           storeKey,
		tStoreKey:          tStoreKey,
		accountKeeper:      accountKeeper,
		spKeeper:           spKeeper,
		paymentKeeper:      paymentKeeper,
		permKeeper:         permKeeper,
		crossChainKeeper:   crossChainKeeper,
		virtualGroupKeeper: virtualGroupKeeper,
		evmKeeper:          evmKeeper,
		authority:          authority,
		cfg:                &paymentCheckConfig{Enabled: false, Interval: 0},
	}

	k.bucketSeq = sequence.NewSequence[sdkmath.Uint](storagetypes.BucketSequencePrefix)
	k.objectSeq = sequence.NewSequence[sdkmath.Uint](storagetypes.ObjectSequencePrefix)
	k.groupSeq = sequence.NewSequence[sdkmath.Uint](storagetypes.GroupSequencePrefix)
	return &k
}

func (k Keeper) GetAuthority() string {
	return k.authority
}

func (k Keeper) IsPaymentCheckEnabled() bool {
	return k.cfg.Enabled
}

func (k Keeper) GetPaymentCheckInterval() uint32 {
	return k.cfg.Interval
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", storagetypes.ModuleName))
}

func (k Keeper) CreateBucket(
	ctx sdk.Context, ownerAcc sdk.AccAddress, bucketName string,
	primarySpAcc sdk.AccAddress, opts *storagetypes.CreateBucketOptions,
) (sdkmath.Uint, error) {
	store := ctx.KVStore(k.storeKey)

	// check if the bucket exist
	bucketKey := storagetypes.GetBucketKey(bucketName)
	if store.Has(bucketKey) {
		return sdkmath.ZeroUint(), storagetypes.ErrBucketAlreadyExists
	}

	// check payment account
	paymentAcc, err := k.VerifyPaymentAccount(ctx, opts.PaymentAddress, ownerAcc)
	if err != nil {
		return sdkmath.ZeroUint(), err
	}

	// check sp and its status
	sp, found := k.spKeeper.GetStorageProviderByOperatorAddr(ctx, primarySpAcc)
	if !found {
		return sdkmath.ZeroUint(), errors.Wrap(storagetypes.ErrNoSuchStorageProvider, "the storage provider is not exist")
	}

	// a sp is not in service, neither in maintenance
	if sp.Status != sptypes.STATUS_IN_SERVICE && !k.fromSpMaintenanceAcct(sp, ownerAcc) {
		return sdkmath.ZeroUint(), errors.Wrap(storagetypes.ErrNoSuchStorageProvider, "the storage provider is not in service")
	}

	err = k.VerifySP(ctx, sp, ownerAcc)
	if err != nil {
		return sdkmath.ZeroUint(), err
	}

	gvgFamily, err := k.virtualGroupKeeper.GetAndCheckGVGFamilyAvailableForNewBucket(ctx, opts.PrimarySpApproval.GlobalVirtualGroupFamilyId)
	if err != nil {
		return sdkmath.ZeroUint(), err
	}

	bucketInfo := storagetypes.BucketInfo{
		Owner:                      ownerAcc.String(),
		BucketName:                 bucketName,
		Visibility:                 opts.Visibility,
		CreateAt:                   ctx.BlockTime().Unix(),
		SourceType:                 opts.SourceType,
		BucketStatus:               storagetypes.BUCKET_STATUS_CREATED,
		ChargedReadQuota:           opts.ChargedReadQuota,
		PaymentAddress:             paymentAcc.String(),
		GlobalVirtualGroupFamilyId: gvgFamily.Id,
	}

	internalBucketInfo := storagetypes.InternalBucketInfo{PriceTime: ctx.BlockTime().Unix()}

	// charge by read quota
	if opts.ChargedReadQuota != 0 {
		err = k.ChargeBucketReadFee(ctx, &bucketInfo, &internalBucketInfo)
		if err != nil {
			return sdkmath.ZeroUint(), err
		}
	}

	// Generate bucket Id
	bucketInfo.Id = k.GenNextBucketId(ctx)

	// store the bucket
	bz := k.cdc.MustMarshal(&bucketInfo)
	store.Set(bucketKey, k.bucketSeq.EncodeSequence(bucketInfo.Id))
	store.Set(storagetypes.GetBucketByIDKey(bucketInfo.Id), bz)
	k.SetInternalBucketInfo(ctx, bucketInfo.Id, &internalBucketInfo)

	// emit CreateBucket Event
	if err = ctx.EventManager().EmitTypedEvents(&storagetypes.EventCreateBucket{
		Owner:                      bucketInfo.Owner,
		BucketName:                 bucketInfo.BucketName,
		Visibility:                 bucketInfo.Visibility,
		CreateAt:                   bucketInfo.CreateAt,
		BucketId:                   bucketInfo.Id,
		SourceType:                 bucketInfo.SourceType,
		Status:                     bucketInfo.BucketStatus,
		ChargedReadQuota:           bucketInfo.ChargedReadQuota,
		PaymentAddress:             bucketInfo.PaymentAddress,
		PrimarySpId:                sp.Id,
		GlobalVirtualGroupFamilyId: bucketInfo.GlobalVirtualGroupFamilyId,
	}); err != nil {
		return sdkmath.Uint{}, err
	}

	// Mint bucket nft token and send to receiver
	_, err = k.CallEVM(
		ctx,
		contracts.ERC721NonTransferableContract.ABI,
		contracts.BucketControlHubAddress,
		contracts.BucketERC721TokenAddress,
		true,
		"mint",
		ecommon.HexToAddress(bucketInfo.Owner),
		bucketInfo.Id.BigInt(),
	)
	if err != nil {
		return sdkmath.ZeroUint(), err
	}

	return bucketInfo.Id, nil
}

// StoreBucketInfo will store the bucket info
// It's designed to be used by the test cases to create a bucket.
func (k Keeper) StoreBucketInfo(ctx sdk.Context, bucketInfo *storagetypes.BucketInfo) {
	store := ctx.KVStore(k.storeKey)
	bucketKey := storagetypes.GetBucketKey(bucketInfo.BucketName)
	bz := k.cdc.MustMarshal(bucketInfo)
	store.Set(bucketKey, k.bucketSeq.EncodeSequence(bucketInfo.Id))
	store.Set(storagetypes.GetBucketByIDKey(bucketInfo.Id), bz)
}

func (k Keeper) DeleteBucket(ctx sdk.Context, operator sdk.AccAddress, bucketName string, opts storagetypes.DeleteBucketOptions) error {
	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}
	if bucketInfo.SourceType != opts.SourceType {
		return storagetypes.ErrSourceTypeMismatch
	}

	// check permission
	effect := k.VerifyBucketPermission(ctx, bucketInfo, operator, permtypes.ACTION_DELETE_BUCKET, nil)
	if effect != permtypes.EFFECT_ALLOW {
		return storagetypes.ErrAccessDenied.Wrapf("The operator(%s) has no DeleteBucket permission of the bucket(%s)",
			operator.String(), bucketName)
	}

	// check if the bucket empty
	if k.isNonEmptyBucket(ctx, bucketName) {
		return storagetypes.ErrBucketNotEmpty
	}

	// check lvgs
	internalBucketInfo := k.MustGetInternalBucketInfo(ctx, bucketInfo.Id)
	for _, lvg := range internalBucketInfo.LocalVirtualGroups {
		if lvg.StoredSize != 0 || lvg.TotalChargeSize != 0 {
			return storagetypes.ErrVirtualGroupOperateFailed.Wrapf("non-empty lvg, %s", lvg.String())
		}
	}

	// change the bill
	err := k.UnChargeBucketReadFee(ctx, bucketInfo, internalBucketInfo)
	if err != nil {
		return storagetypes.ErrChargeFailed.Wrapf("cancel charge bucket read fee error: %s", err)
	}

	return k.doDeleteBucket(ctx, operator, bucketInfo)
}

func (k Keeper) doDeleteBucket(ctx sdk.Context, operator sdk.AccAddress, bucketInfo *storagetypes.BucketInfo) error {
	store := ctx.KVStore(k.storeKey)
	store.Delete(storagetypes.GetBucketKey(bucketInfo.BucketName))
	store.Delete(storagetypes.GetBucketByIDKey(bucketInfo.Id))
	store.Delete(storagetypes.GetQuotaKey(bucketInfo.Id))
	store.Delete(storagetypes.GetInternalBucketInfoKey(bucketInfo.Id))
	store.Delete(storagetypes.GetMigrationBucketKey(bucketInfo.Id))

	store.Delete(storagetypes.GetLockedObjectCountKey(bucketInfo.Id))

	err := k.appendResourceIDForGarbageCollection(ctx, resource.RESOURCE_TYPE_BUCKET, bucketInfo.Id)
	if err != nil {
		return err
	}
	if err = ctx.EventManager().EmitTypedEvents(&storagetypes.EventDeleteBucket{
		Operator:                   operator.String(),
		Owner:                      bucketInfo.Owner,
		BucketName:                 bucketInfo.BucketName,
		BucketId:                   bucketInfo.Id,
		GlobalVirtualGroupFamilyId: bucketInfo.GlobalVirtualGroupFamilyId,
	}); err != nil {
		return err
	}

	if bucketInfo.BucketStatus == storagetypes.BUCKET_STATUS_MIGRATING {
		if err = ctx.EventManager().EmitTypedEvents(&storagetypes.EventCancelMigrationBucket{
			Operator:   operator.String(),
			BucketName: bucketInfo.BucketName,
			BucketId:   bucketInfo.Id,
		}); err != nil {
			return err
		}
	}

	// delete bucket flow rate limit status
	k.deleteBucketFlowRateLimitStatus(ctx, bucketInfo.BucketName, bucketInfo.Id)

	if _, err := k.CallEVM(
		ctx,
		contracts.ERC721NonTransferableContract.ABI,
		contracts.BucketControlHubAddress,
		contracts.BucketERC721TokenAddress,
		true,
		"burn",
		bucketInfo.Id.BigInt(),
	); err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetPrimarySPForBucket(ctx sdk.Context, bucketInfo *storagetypes.BucketInfo) (*sptypes.StorageProvider, error) {
	gvgFamily, found := k.virtualGroupKeeper.GetGVGFamily(ctx, bucketInfo.GlobalVirtualGroupFamilyId)
	if !found {
		return nil, virtualgroupmoduletypes.ErrGVGFamilyNotExist.Wrapf("gvg family (%d) not found.", bucketInfo.GlobalVirtualGroupFamilyId)
	}
	sp := k.spKeeper.MustGetStorageProvider(ctx, gvgFamily.PrimarySpId)
	return sp, nil
}

func (k Keeper) MustGetPrimarySPForBucket(ctx sdk.Context, bucketInfo *storagetypes.BucketInfo) *sptypes.StorageProvider {
	sp, err := k.GetPrimarySPForBucket(ctx, bucketInfo)
	if err != nil {
		panic(err)
	}
	return sp
}

// ForceDeleteBucket will delete bucket without permission check, it is used for discontinue request from sps.
// The cap parameter will limit the max objects can be deleted in the call.
// It will also return 1) whether the bucket is deleted, 2) the objects deleted, and 3) error if there is
func (k Keeper) ForceDeleteBucket(ctx sdk.Context, bucketID sdkmath.Uint, cap uint64) (bool, uint64, error) {
	bucketInfo, found := k.GetBucketInfoById(ctx, bucketID)
	if !found { // the bucket is already deleted
		return true, 0, nil
	}

	bucketDeleted := false

	sp := k.MustGetPrimarySPForBucket(ctx, bucketInfo)
	spOperatorAddr := sdk.MustAccAddressFromHex(sp.OperatorAddress)

	store := ctx.KVStore(k.storeKey)
	objectPrefixStore := prefix.NewStore(store, storagetypes.GetObjectKeyOnlyBucketPrefix(bucketInfo.BucketName))
	iter := objectPrefixStore.Iterator(nil, nil)
	defer iter.Close()
	u256Seq := sequence.Sequence[sdkmath.Uint]{}

	deleted := uint64(0) // deleted object count
	var err error
	for ; iter.Valid(); iter.Next() {
		if deleted >= cap {
			return false, deleted, nil // break is also fine here
		}

		bz := store.Get(storagetypes.GetObjectByIDKey(u256Seq.DecodeSequence(iter.Value())))
		if bz == nil {
			panic("should not happen")
		}

		var objectInfo storagetypes.ObjectInfo
		k.cdc.MustUnmarshal(bz, &objectInfo)

		// An object cannot be discontinued if the bucket is already discontinued,
		// which means that after deleting objects when deleting a bucket the objects in it should be in
		// OBJECT_STATUS_CREATED or OBJECT_STATUS_SEALED status.
		// However, when updating the deletion confirm period parameter, it can lead to that
		// the discontinued bucket can be deleted earlier, then its objects could be in
		// OBJECT_STATUS_DISCONTINUED status.

		objectStatus := objectInfo.ObjectStatus
		if objectStatus == storagetypes.OBJECT_STATUS_DISCONTINUED {
			objectStatus, err = k.getAndDeleteDiscontinueObjectStatus(ctx, objectInfo.Id)
			if err != nil {
				return false, deleted, err
			}
		}

		if objectStatus == storagetypes.OBJECT_STATUS_CREATED {
			if err = k.UnlockObjectStoreFee(ctx, bucketInfo, &objectInfo); err != nil {
				ctx.Logger().Error("unlock store fee error", "err", err)
				return false, deleted, err
			}

			k.DecreaseLockedObjectCount(ctx, bucketInfo.Id)

		} else if objectStatus == storagetypes.OBJECT_STATUS_SEALED {
			internalBucketInfo := k.MustGetInternalBucketInfo(ctx, bucketInfo.Id)
			if err = k.UnChargeObjectStoreFee(ctx, bucketInfo, internalBucketInfo, &objectInfo); err != nil {
				ctx.Logger().Error("charge delete object error", "err", err)
				return false, deleted, err
			}
			k.SetInternalBucketInfo(ctx, bucketInfo.Id, internalBucketInfo)

			// if an object is updating, also need to unlock the shadowObject fee
			if objectInfo.IsUpdating {
				shadowObjectInfo := k.MustGetShadowObjectInfo(ctx, bucketInfo.BucketName, objectInfo.ObjectName)
				err = k.UnlockShadowObjectFeeAndDeleteShadowObjectInfo(ctx, bucketInfo, shadowObjectInfo, objectInfo.ObjectName)
				if err != nil {
					return false, deleted, err
				}
				k.DecreaseLockedObjectCount(ctx, bucketInfo.Id)
			}
		}
		if err := k.doDeleteObject(ctx, spOperatorAddr, bucketInfo, &objectInfo); err != nil {
			ctx.Logger().Error("do delete object err", "err", err)
			return false, deleted, err
		}
		deleted++
	}

	if !iter.Valid() {
		internalBucketInfo := k.MustGetInternalBucketInfo(ctx, bucketInfo.Id)
		if err = k.UnChargeBucketReadFee(ctx, bucketInfo, internalBucketInfo); err != nil {
			ctx.Logger().Error("charge delete bucket error", "err", err)
			return false, deleted, err
		}

		if err := k.doDeleteBucket(ctx, spOperatorAddr, bucketInfo); err != nil {
			ctx.Logger().Error("do delete bucket error", "err", err)
			return false, deleted, err
		}
		bucketDeleted = true
	}

	return bucketDeleted, deleted, nil
}

func (k Keeper) UpdateBucketInfo(ctx sdk.Context, operator sdk.AccAddress, bucketName string, opts storagetypes.UpdateBucketOptions) error {
	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}
	// check bucket source
	if bucketInfo.SourceType != opts.SourceType {
		return storagetypes.ErrSourceTypeMismatch
	}

	sp := k.MustGetPrimarySPForBucket(ctx, bucketInfo)
	if sp.Status == sptypes.STATUS_GRACEFUL_EXITING || sp.Status == sptypes.STATUS_FORCED_EXITING {
		return storagetypes.ErrUpdateQuotaFailed.Wrapf("The SP is in %s, bucket can not be updated", sp.Status)
	}

	// check permission
	effect := k.VerifyBucketPermission(ctx, bucketInfo, operator, permtypes.ACTION_UPDATE_BUCKET_INFO, nil)
	if effect != permtypes.EFFECT_ALLOW {
		return storagetypes.ErrAccessDenied.Wrapf("The operator(%s) has no UpdateBucketInfo permission of the bucket(%s)",
			operator.String(), bucketName)
	}

	// handle fields not changed
	if opts.ChargedReadQuota == nil {
		opts.ChargedReadQuota = &bucketInfo.ChargedReadQuota
	} else if *opts.ChargedReadQuota != bucketInfo.ChargedReadQuota {
		blockTime := uint64(ctx.BlockTime().Unix())
		if *opts.ChargedReadQuota < bucketInfo.ChargedReadQuota {
			minInterval := k.GetParams(ctx).MinQuotaUpdateInterval
			lastUpdateTime, found := k.getQuotaUpdateTime(ctx, bucketInfo.Id)
			if !found {
				return storagetypes.ErrUpdateQuotaFailed
			}
			if lastUpdateTime+minInterval > blockTime {
				return storagetypes.ErrUpdateQuotaFailed.Wrapf("The quota can be updated to a smaller value before %d timestamp",
					lastUpdateTime+minInterval)
			}
		}
		// save quota update time
		k.setQuotaUpdateTime(ctx, bucketInfo.Id, blockTime)
	}

	if opts.Visibility != storagetypes.VISIBILITY_TYPE_UNSPECIFIED {
		bucketInfo.Visibility = opts.Visibility
	}

	var paymentAcc sdk.AccAddress
	var err error
	if opts.PaymentAddress != "" {
		ownerAcc := sdk.MustAccAddressFromHex(bucketInfo.Owner)
		paymentAcc, err = k.VerifyPaymentAccount(ctx, opts.PaymentAddress, ownerAcc)
		if err != nil {
			return err
		}

		if !paymentAcc.Equals(sdk.MustAccAddressFromHex(bucketInfo.PaymentAddress)) && k.GetLockedObjectCount(ctx, bucketInfo.Id) != 0 {
			return storagetypes.ErrUpdatePaymentAccountFailed.Wrapf("The bucket %s has unseald objects", bucketInfo.BucketName)
		}
	} else {
		paymentAcc = sdk.MustAccAddressFromHex(bucketInfo.PaymentAddress)
	}

	internalBucketInfo := k.MustGetInternalBucketInfo(ctx, bucketInfo.Id)
	err = k.UpdateBucketInfoAndCharge(ctx, bucketInfo, internalBucketInfo, paymentAcc.String(), *opts.ChargedReadQuota)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(bucketInfo)
	store.Set(storagetypes.GetBucketByIDKey(bucketInfo.Id), bz)
	k.SetInternalBucketInfo(ctx, bucketInfo.Id, internalBucketInfo)

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventUpdateBucketInfo{
		Operator:                   operator.String(),
		BucketName:                 bucketName,
		BucketId:                   bucketInfo.Id,
		ChargedReadQuota:           bucketInfo.ChargedReadQuota,
		PaymentAddress:             bucketInfo.PaymentAddress,
		Visibility:                 bucketInfo.Visibility,
		GlobalVirtualGroupFamilyId: bucketInfo.GlobalVirtualGroupFamilyId,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) DiscontinueBucket(ctx sdk.Context, operator sdk.AccAddress, bucketName, reason string) error {
	sp, found := k.spKeeper.GetStorageProviderByGcAddr(ctx, operator)
	if !found {
		return storagetypes.ErrNoSuchStorageProvider.Wrapf("SP operator address: %s", operator.String())
	}
	if sp.Status != sptypes.STATUS_IN_SERVICE {
		return sptypes.ErrStorageProviderNotInService
	}

	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}
	if bucketInfo.BucketStatus == storagetypes.BUCKET_STATUS_DISCONTINUED {
		return storagetypes.ErrInvalidBucketStatus
	}

	spInState := k.MustGetPrimarySPForBucket(ctx, bucketInfo)

	if sp.Id != spInState.Id {
		return errors.Wrapf(storagetypes.ErrAccessDenied,
			"only primary sp is allowed to do discontinue bucket, expect sp id : %d", spInState.Id)
	}

	count := k.GetDiscontinueBucketCount(ctx, operator)
	max := k.DiscontinueBucketMax(ctx)
	if count+1 > max {
		return storagetypes.ErrNoMoreDiscontinue.Wrapf("no more buckets can be requested in this window")
	}
	previousStatus := bucketInfo.BucketStatus
	bucketInfo.BucketStatus = storagetypes.BUCKET_STATUS_DISCONTINUED

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(bucketInfo)
	store.Set(storagetypes.GetBucketByIDKey(bucketInfo.Id), bz)

	deleteAt := ctx.BlockTime().Unix() + k.DiscontinueConfirmPeriod(ctx)

	k.appendDiscontinueBucketIDs(ctx, deleteAt, []sdkmath.Uint{bucketInfo.Id})
	k.SetDiscontinueBucketCount(ctx, operator, count+1)

	if previousStatus == storagetypes.BUCKET_STATUS_MIGRATING {
		if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventCancelMigrationBucket{
			Operator:   operator.String(),
			BucketName: bucketInfo.BucketName,
			BucketId:   bucketInfo.Id,
		}); err != nil {
			return err
		}
	}

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventDiscontinueBucket{
		BucketId:   bucketInfo.Id,
		BucketName: bucketInfo.BucketName,
		Reason:     reason,
		DeleteAt:   deleteAt,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) SetBucketInfo(ctx sdk.Context, bucketInfo *storagetypes.BucketInfo) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(bucketInfo)
	store.Set(storagetypes.GetBucketByIDKey(bucketInfo.Id), bz)
}

func (k Keeper) GetBucketInfo(ctx sdk.Context, bucketName string) (*storagetypes.BucketInfo, bool) {
	store := ctx.KVStore(k.storeKey)

	bucketKey := storagetypes.GetBucketKey(bucketName)
	bz := store.Get(bucketKey)
	if bz == nil {
		return nil, false
	}

	return k.GetBucketInfoById(ctx, k.bucketSeq.DecodeSequence(bz))
}

func (k Keeper) GetBucketInfoById(ctx sdk.Context, bucketId sdkmath.Uint) (*storagetypes.BucketInfo, bool) { //nolint
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(storagetypes.GetBucketByIDKey(bucketId))
	if bz == nil {
		return nil, false
	}

	var bucketInfo storagetypes.BucketInfo
	k.cdc.MustUnmarshal(bz, &bucketInfo)

	return &bucketInfo, true
}

func (k Keeper) CreateObject(
	ctx sdk.Context, operator sdk.AccAddress, bucketName, objectName string, payloadSize uint64,
	opts storagetypes.CreateObjectOptions,
) (sdkmath.Uint, error) {
	store := ctx.KVStore(k.storeKey)

	// check payload size
	if payloadSize > k.MaxPayloadSize(ctx) {
		return sdkmath.ZeroUint(), storagetypes.ErrTooLargeObject
	}

	// check bucket
	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return sdkmath.ZeroUint(), storagetypes.ErrNoSuchBucket
	}
	err := bucketInfo.CheckBucketStatus()
	if err != nil {
		return sdkmath.ZeroUint(), err
	}

	// primary sp
	sp := k.MustGetPrimarySPForBucket(ctx, bucketInfo)
	creator := operator
	if opts.Delegated {
		creator = opts.Creator
		if bucketInfo.SpAsDelegatedAgentDisabled {
			return sdkmath.ZeroUint(), storagetypes.ErrAccessDenied.Wrap("the SP is not allowed to create object for delegator, disabled by the bucket owner previously")
		}
		if operator.String() != sp.OperatorAddress {
			return sdkmath.ZeroUint(), storagetypes.ErrAccessDenied.Wrap("only the primary SP is allowed to create object for delegator")
		}
	}

	// verify permission
	verifyOpts := &permtypes.VerifyOptions{
		WantedSize: &payloadSize,
	}
	effect := k.VerifyBucketPermission(ctx, bucketInfo, creator, permtypes.ACTION_CREATE_OBJECT, verifyOpts)
	if effect != permtypes.EFFECT_ALLOW {
		return sdkmath.ZeroUint(), storagetypes.ErrAccessDenied.Wrapf("The creator(%s) has no CreateObject permission of the bucket(%s)",
			operator.String(), bucketName)
	}

	objectInfoCreator := creator
	if objectInfoCreator.Equals(sdk.MustAccAddressFromHex(bucketInfo.Owner)) {
		objectInfoCreator = sdk.AccAddress{}
	}

	err = k.VerifySP(ctx, sp, operator)
	if err != nil {
		return sdkmath.ZeroUint(), err
	}

	objectKey := storagetypes.GetObjectKey(bucketName, objectName)
	if store.Has(objectKey) {
		return sdkmath.ZeroUint(), storagetypes.ErrObjectAlreadyExists
	}

	// check payload size, the empty object doesn't need sealed
	var objectStatus storagetypes.ObjectStatus
	if payloadSize == 0 {
		// empty object does not interact with sp
		objectStatus = storagetypes.OBJECT_STATUS_SEALED
	} else {
		objectStatus = storagetypes.OBJECT_STATUS_CREATED
	}

	// construct objectInfo
	objectInfo := storagetypes.ObjectInfo{
		Owner:          bucketInfo.Owner,
		Creator:        objectInfoCreator.String(),
		BucketName:     bucketName,
		ObjectName:     objectName,
		PayloadSize:    payloadSize,
		Visibility:     opts.Visibility,
		ContentType:    opts.ContentType,
		Id:             k.GenNextObjectID(ctx),
		CreateAt:       ctx.BlockTime().Unix(),
		ObjectStatus:   objectStatus,
		RedundancyType: opts.RedundancyType,
		SourceType:     opts.SourceType,
		Checksums:      opts.Checksums,
	}

	if objectInfo.PayloadSize == 0 {
		_, err := k.SealEmptyObjectOnVirtualGroup(ctx, bucketInfo, &objectInfo)
		if err != nil {
			return sdkmath.ZeroUint(), err
		}
	} else {
		// Lock Fee
		err = k.LockObjectStoreFee(ctx, bucketInfo, &objectInfo)
		if err != nil {
			return sdkmath.ZeroUint(), err
		}
	}

	bbz := k.cdc.MustMarshal(bucketInfo)
	store.Set(storagetypes.GetBucketByIDKey(bucketInfo.Id), bbz)

	if objectInfo.ObjectStatus == storagetypes.OBJECT_STATUS_CREATED {
		k.IncreaseLockedObjectCount(ctx, bucketInfo.Id)
	}

	obz := k.cdc.MustMarshal(&objectInfo)
	store.Set(objectKey, k.objectSeq.EncodeSequence(objectInfo.Id))
	store.Set(storagetypes.GetObjectByIDKey(objectInfo.Id), obz)

	if err = ctx.EventManager().EmitTypedEvents(&storagetypes.EventCreateObject{
		Creator:             creator.String(),
		Owner:               objectInfo.Owner,
		BucketName:          bucketInfo.BucketName,
		ObjectName:          objectInfo.ObjectName,
		BucketId:            bucketInfo.Id,
		ObjectId:            objectInfo.Id,
		CreateAt:            objectInfo.CreateAt,
		PayloadSize:         objectInfo.PayloadSize,
		Visibility:          objectInfo.Visibility,
		PrimarySpId:         sp.Id,
		ContentType:         objectInfo.ContentType,
		Status:              objectInfo.ObjectStatus,
		RedundancyType:      objectInfo.RedundancyType,
		SourceType:          objectInfo.SourceType,
		Checksums:           objectInfo.Checksums,
		LocalVirtualGroupId: objectInfo.LocalVirtualGroupId,
	}); err != nil {
		return objectInfo.Id, err
	}
	return objectInfo.Id, nil
}

// StoreObjectInfo stores object related keys to KVStore,
// it's designed to be used in tests
func (k Keeper) StoreObjectInfo(ctx sdk.Context, objectInfo *storagetypes.ObjectInfo) {
	store := ctx.KVStore(k.storeKey)

	objectKey := storagetypes.GetObjectKey(objectInfo.BucketName, objectInfo.ObjectName)

	obz := k.cdc.MustMarshal(objectInfo)
	store.Set(objectKey, k.objectSeq.EncodeSequence(objectInfo.Id))
	store.Set(storagetypes.GetObjectByIDKey(objectInfo.Id), obz)
}

// DeleteObjectInfo deletes object related keys from KVStore,
// it's designed to be used in tests
func (k Keeper) DeleteObjectInfo(ctx sdk.Context, objectInfo *storagetypes.ObjectInfo) {
	store := ctx.KVStore(k.storeKey)

	objectKey := storagetypes.GetObjectKey(objectInfo.BucketName, objectInfo.ObjectName)

	store.Delete(objectKey)
	store.Delete(storagetypes.GetObjectByIDKey(objectInfo.Id))
}

func (k Keeper) SetObjectInfo(ctx sdk.Context, objectInfo *storagetypes.ObjectInfo) {
	store := ctx.KVStore(k.storeKey)

	obz := k.cdc.MustMarshal(objectInfo)
	store.Set(storagetypes.GetObjectByIDKey(objectInfo.Id), obz)
}

func (k Keeper) GetObjectInfoCount(ctx sdk.Context) sdkmath.Uint {
	store := ctx.KVStore(k.storeKey)

	seq := k.objectSeq.CurVal(store)
	return seq
}

func (k Keeper) GetObjectInfo(ctx sdk.Context, bucketName, objectName string) (*storagetypes.ObjectInfo, bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(storagetypes.GetObjectKey(bucketName, objectName))
	if bz == nil {
		return nil, false
	}

	return k.GetObjectInfoById(ctx, k.objectSeq.DecodeSequence(bz))
}

func (k Keeper) GetObjectInfoById(ctx sdk.Context, objectId sdkmath.Uint) (*storagetypes.ObjectInfo, bool) { //nolint
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(storagetypes.GetObjectByIDKey(objectId))
	if bz == nil {
		return nil, false
	}

	var objectInfo storagetypes.ObjectInfo
	k.cdc.MustUnmarshal(bz, &objectInfo)
	return &objectInfo, true
}

func (k Keeper) GetShadowObjectInfo(ctx sdk.Context, bucketName, objectName string) (*storagetypes.ShadowObjectInfo, bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(storagetypes.GetShadowObjectKey(bucketName, objectName))
	if bz == nil {
		return nil, false
	}

	var objectInfo storagetypes.ShadowObjectInfo
	k.cdc.MustUnmarshal(bz, &objectInfo)
	return &objectInfo, true
}

func (k Keeper) MustGetShadowObjectInfo(ctx sdk.Context, bucketName, objectName string) *storagetypes.ShadowObjectInfo {
	shadowObjectInfo, found := k.GetShadowObjectInfo(ctx, bucketName, objectName)
	if !found {
		panic("Shadow Object Info not found")
	}
	return shadowObjectInfo
}

type SealObjectOptions struct {
	GlobalVirtualGroupID     uint32
	SecondarySpBlsSignatures []byte
	Checksums                [][]byte
}

func (k Keeper) SealObject(
	ctx sdk.Context, spSealAcc sdk.AccAddress,
	bucketName, objectName string, opts SealObjectOptions,
) error {
	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}

	sp, found := k.spKeeper.GetStorageProviderBySealAddr(ctx, spSealAcc)
	if !found {
		return errors.Wrapf(storagetypes.ErrNoSuchStorageProvider, "SP seal address: %s", spSealAcc.String())
	}

	spInState := k.MustGetPrimarySPForBucket(ctx, bucketInfo)

	if sp.Id != spInState.Id {
		return errors.Wrapf(storagetypes.ErrAccessDenied, "Only SP's seal address is allowed to SealObject")
	}

	objectInfo, found := k.GetObjectInfo(ctx, bucketName, objectName)
	if !found {
		return storagetypes.ErrNoSuchObject
	}

	if objectInfo.Checksums == nil {
		if opts.Checksums == nil {
			return storagetypes.ErrObjectChecksumsMissing
		}
		objectInfo.Checksums = opts.Checksums
	}
	store := ctx.KVStore(k.storeKey)

	prevPayloadSize := objectInfo.PayloadSize
	prevCheckSums := objectInfo.Checksums

	isUpdate := objectInfo.IsUpdating

	// an object might be set to OBJECT_STATUS_DISCONTINUED
	if isUpdate && objectInfo.ObjectStatus == storagetypes.OBJECT_STATUS_SEALED {
		internalBucketInfo := k.MustGetInternalBucketInfo(ctx, bucketInfo.Id)
		err := k.UnChargeObjectStoreFee(ctx, bucketInfo, internalBucketInfo, objectInfo)
		if err != nil {
			return err
		}
		k.SetInternalBucketInfo(ctx, bucketInfo.Id, internalBucketInfo)
		err = k.DeleteObjectFromVirtualGroup(ctx, bucketInfo, objectInfo)
		if err != nil {
			return err
		}

		shadowObjectInfo := k.MustGetShadowObjectInfo(ctx, bucketName, objectName)
		objectInfo.UpdatedAt = shadowObjectInfo.UpdatedAt // the updated_at in objetInfo will not be visible until the object is sealed.
		objectInfo.Version = shadowObjectInfo.Version
		if shadowObjectInfo.Checksums == nil {
			if opts.Checksums == nil {
				return storagetypes.ErrObjectChecksumsMissing
			}
			shadowObjectInfo.Checksums = opts.Checksums
		}
		objectInfo.Checksums = shadowObjectInfo.Checksums
		objectInfo.PayloadSize = shadowObjectInfo.PayloadSize
		objectInfo.UpdatedBy = shadowObjectInfo.Operator
		objectInfo.ContentType = shadowObjectInfo.ContentType
		objectInfo.IsUpdating = false

		store.Delete(storagetypes.GetShadowObjectKey(bucketInfo.BucketName, objectName))
	} else if objectInfo.ObjectStatus != storagetypes.OBJECT_STATUS_CREATED {
		return storagetypes.ErrObjectAlreadySealed
	}

	gvg, found := k.virtualGroupKeeper.GetGVG(ctx, opts.GlobalVirtualGroupID)
	if !found {
		return virtualgroupmoduletypes.ErrGVGNotExist
	}

	if gvg.FamilyId != bucketInfo.GlobalVirtualGroupFamilyId || gvg.PrimarySpId != spInState.Id {
		return storagetypes.ErrInvalidGlobalVirtualGroup.Wrapf("Global virtual group mismatch, familyID: %d, bucket family ID: %d", gvg.FamilyId, bucketInfo.GlobalVirtualGroupFamilyId)
	}
	expectSecondarySPNum := k.GetExpectSecondarySPNumForECObject(ctx, objectInfo.GetLatestUpdatedTime())
	if int(expectSecondarySPNum) != len(gvg.SecondarySpIds) {
		return storagetypes.ErrInvalidGlobalVirtualGroup.Wrapf("secondary sp num mismatch, expect (%d), but (%d)",
			expectSecondarySPNum, len(gvg.SecondarySpIds))
	}
	// validate seal object bls aggregated sig from secondary sps
	secondarySpsSealObjectBlsSignHash := storagetypes.NewSecondarySpSealObjectSignDoc(ctx.ChainID(), gvg.Id, objectInfo.Id, storagetypes.GenerateHash(objectInfo.Checksums)).GetBlsSignHash()
	err := k.VerifyGVGSecondarySPsBlsSignature(ctx, gvg, secondarySpsSealObjectBlsSignHash, opts.SecondarySpBlsSignatures)
	if err != nil {
		return err
	}

	_, err = k.SealObjectOnVirtualGroup(ctx, bucketInfo, opts.GlobalVirtualGroupID, objectInfo)
	if err != nil {
		return errors.Wrapf(storagetypes.ErrInvalidGlobalVirtualGroup, "err message: %s", err)
	}

	objectInfo.ObjectStatus = storagetypes.OBJECT_STATUS_SEALED

	bbz := k.cdc.MustMarshal(bucketInfo)
	store.Set(storagetypes.GetBucketByIDKey(bucketInfo.Id), bbz)

	k.DecreaseLockedObjectCount(ctx, bucketInfo.Id)

	obz := k.cdc.MustMarshal(objectInfo)
	store.Set(storagetypes.GetObjectByIDKey(objectInfo.Id), obz)

	if isUpdate {
		if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventUpdateObjectContentSuccess{
			Operator:        spSealAcc.String(),
			BucketName:      bucketInfo.BucketName,
			ObjectName:      objectInfo.ObjectName,
			ObjectId:        objectInfo.Id,
			ContentType:     objectInfo.ContentType,
			PrevPayloadSize: prevPayloadSize,
			NewPayloadSize:  objectInfo.PayloadSize,
			PrevChecksums:   prevCheckSums,
			NewChecksums:    objectInfo.Checksums,
			Version:         objectInfo.Version,
			UpdatedAt:       objectInfo.UpdatedAt,
		}); err != nil {
			return err
		}
	}
	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventSealObject{
		Operator:             spSealAcc.String(),
		BucketName:           bucketInfo.BucketName,
		ObjectName:           objectInfo.ObjectName,
		ObjectId:             objectInfo.Id,
		Status:               objectInfo.ObjectStatus,
		GlobalVirtualGroupId: opts.GlobalVirtualGroupID,
		LocalVirtualGroupId:  objectInfo.LocalVirtualGroupId,
		Checksums:            objectInfo.Checksums,
	}); err != nil {
		return err
	}

	// Mint object nft token and send to receiver
	_, err = k.CallEVM(
		ctx,
		contracts.ERC721NonTransferableContract.ABI,
		contracts.ObjectControlHubAddress,
		contracts.ObjectERC721TokenAddress,
		true,
		"mint",
		ecommon.HexToAddress(objectInfo.Owner),
		objectInfo.Id.BigInt(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) CancelCreateObject(
	ctx sdk.Context, operator sdk.AccAddress,
	bucketName, objectName string, opts storagetypes.CancelCreateObjectOptions,
) error {
	store := ctx.KVStore(k.storeKey)
	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}
	objectInfo, found := k.GetObjectInfo(ctx, bucketName, objectName)
	if !found {
		return storagetypes.ErrNoSuchObject
	}

	spInState := k.MustGetPrimarySPForBucket(ctx, bucketInfo)
	if objectInfo.ObjectStatus != storagetypes.OBJECT_STATUS_CREATED {
		return storagetypes.ErrObjectNotCreated.Wrapf("Object status: %s", objectInfo.ObjectStatus.String())
	}

	if objectInfo.SourceType != opts.SourceType {
		return storagetypes.ErrSourceTypeMismatch
	}

	var creator sdk.AccAddress
	if objectInfo.Creator != "" {
		creator = sdk.MustAccAddressFromHex(objectInfo.Creator)
	}

	// check permission, does not include checking the creator
	effect := k.VerifyObjectPermission(ctx, bucketInfo, objectInfo, operator, permtypes.ACTION_DELETE_OBJECT)
	if effect != permtypes.EFFECT_ALLOW && !operator.Equals(creator) {
		return storagetypes.ErrAccessDenied.Wrapf(
			"The operator(%s) has no DeleteObject permission of the bucket(%s), object(%s)",
			operator.String(), bucketName, objectName)
	}

	err := k.UnlockObjectStoreFee(ctx, bucketInfo, objectInfo)
	if err != nil {
		return err
	}

	bbz := k.cdc.MustMarshal(bucketInfo)
	store.Set(storagetypes.GetBucketByIDKey(bucketInfo.Id), bbz)

	if objectInfo.ObjectStatus == storagetypes.OBJECT_STATUS_CREATED {
		k.DecreaseLockedObjectCount(ctx, bucketInfo.Id)
	}

	store.Delete(storagetypes.GetObjectKey(bucketName, objectName))
	store.Delete(storagetypes.GetObjectByIDKey(objectInfo.Id))

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventCancelCreateObject{
		Operator:    operator.String(),
		BucketName:  bucketInfo.BucketName,
		ObjectName:  objectInfo.ObjectName,
		ObjectId:    objectInfo.Id,
		PrimarySpId: spInState.Id,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) DeleteObject(
	ctx sdk.Context, operator sdk.AccAddress, bucketName, objectName string, opts storagetypes.DeleteObjectOptions,
) error {
	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}

	objectInfo, found := k.GetObjectInfo(ctx, bucketName, objectName)
	if !found {
		return storagetypes.ErrNoSuchObject
	}

	if objectInfo.ObjectStatus == storagetypes.OBJECT_STATUS_DISCONTINUED {
		return storagetypes.ErrInvalidObjectStatus.Wrapf("The object %s is discontinued, will be deleted automatically",
			objectInfo.ObjectName)
	}

	if objectInfo.SourceType != opts.SourceType {
		return storagetypes.ErrSourceTypeMismatch
	}

	if objectInfo.ObjectStatus == storagetypes.OBJECT_STATUS_CREATED {
		return k.CancelCreateObject(ctx, operator, bucketName, objectName, storagetypes.CancelCreateObjectOptions{SourceType: storagetypes.SOURCE_TYPE_ORIGIN})
	}
	// check permission
	effect := k.VerifyObjectPermission(ctx, bucketInfo, objectInfo, operator, permtypes.ACTION_DELETE_OBJECT)
	if effect != permtypes.EFFECT_ALLOW {
		return storagetypes.ErrAccessDenied.Wrapf(
			"The operator(%s) has no DeleteObject permission of the bucket(%s), object(%s)",
			operator.String(), bucketName, objectName)
	}
	if objectInfo.IsUpdating {
		shadowObjectInfo := k.MustGetShadowObjectInfo(ctx, bucketInfo.BucketName, objectInfo.ObjectName)
		err := k.UnlockShadowObjectFeeAndDeleteShadowObjectInfo(ctx, bucketInfo, shadowObjectInfo, objectInfo.ObjectName)
		if err != nil {
			return err
		}
		k.DecreaseLockedObjectCount(ctx, bucketInfo.Id)
	}
	internalBucketInfo := k.MustGetInternalBucketInfo(ctx, bucketInfo.Id)

	err := k.UnChargeObjectStoreFee(ctx, bucketInfo, internalBucketInfo, objectInfo)
	if err != nil {
		return err
	}
	k.SetInternalBucketInfo(ctx, bucketInfo.Id, internalBucketInfo)

	err = k.doDeleteObject(ctx, operator, bucketInfo, objectInfo)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) doDeleteObject(ctx sdk.Context, operator sdk.AccAddress, bucketInfo *storagetypes.BucketInfo, objectInfo *storagetypes.ObjectInfo) error {
	store := ctx.KVStore(k.storeKey)

	bbz := k.cdc.MustMarshal(bucketInfo)
	store.Set(storagetypes.GetBucketByIDKey(bucketInfo.Id), bbz)

	store.Delete(storagetypes.GetObjectKey(bucketInfo.BucketName, objectInfo.ObjectName))
	store.Delete(storagetypes.GetObjectByIDKey(objectInfo.Id))

	// when object was not sealed, the lvg id is 0 by default.
	if objectInfo.LocalVirtualGroupId != 0 {
		err := k.DeleteObjectFromVirtualGroup(ctx, bucketInfo, objectInfo)
		if err != nil {
			return err
		}
	}

	err := k.appendResourceIDForGarbageCollection(ctx, resource.RESOURCE_TYPE_OBJECT, objectInfo.Id)
	if err != nil {
		return err
	}

	err = ctx.EventManager().EmitTypedEvents(&storagetypes.EventDeleteObject{
		Operator:            operator.String(),
		BucketName:          bucketInfo.BucketName,
		ObjectName:          objectInfo.ObjectName,
		ObjectId:            objectInfo.Id,
		LocalVirtualGroupId: objectInfo.LocalVirtualGroupId,
	})
	return err
}

// ForceDeleteObject will delete object without permission check, it is used for discontinue request from sps.
func (k Keeper) ForceDeleteObject(ctx sdk.Context, objectID sdkmath.Uint) error {
	objectInfo, found := k.GetObjectInfoById(ctx, objectID)
	if !found { // the object is deleted already
		return nil
	}

	bucketInfo, found := k.GetBucketInfo(ctx, objectInfo.BucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}

	objectStatus, err := k.getAndDeleteDiscontinueObjectStatus(ctx, objectID)
	if err != nil {
		return err
	}

	spInState := k.MustGetPrimarySPForBucket(ctx, bucketInfo)
	if objectStatus == storagetypes.OBJECT_STATUS_CREATED {
		err := k.UnlockObjectStoreFee(ctx, bucketInfo, objectInfo)
		if err != nil {
			ctx.Logger().Error("unlock store fee error", "err", err)
			return err
		}
		k.DecreaseLockedObjectCount(ctx, bucketInfo.Id)
	} else if objectStatus == storagetypes.OBJECT_STATUS_SEALED {
		internalBucketInfo := k.MustGetInternalBucketInfo(ctx, bucketInfo.Id)
		err := k.UnChargeObjectStoreFee(ctx, bucketInfo, internalBucketInfo, objectInfo)
		if err != nil {
			ctx.Logger().Error("charge delete object error", "err", err)
			return err
		}
		k.SetInternalBucketInfo(ctx, bucketInfo.Id, internalBucketInfo)

		// if an object is updating, also need to unlock the shadowObject fee
		if objectInfo.IsUpdating {
			shadowObjectInfo := k.MustGetShadowObjectInfo(ctx, bucketInfo.BucketName, objectInfo.ObjectName)
			err = k.UnlockShadowObjectFeeAndDeleteShadowObjectInfo(ctx, bucketInfo, shadowObjectInfo, objectInfo.ObjectName)
			if err != nil {
				return err
			}
			k.DecreaseLockedObjectCount(ctx, bucketInfo.Id)
		}
	}

	err = k.doDeleteObject(ctx, sdk.MustAccAddressFromHex(spInState.OperatorAddress), bucketInfo, objectInfo)
	if err != nil {
		ctx.Logger().Error("do delete object err", "err", err)
		return err
	}
	return nil
}

func (k Keeper) CopyObject(
	ctx sdk.Context, operator sdk.AccAddress, srcBucketName, srcObjectName, dstBucketName, dstObjectName string,
	opts storagetypes.CopyObjectOptions,
) (sdkmath.Uint, error) {
	store := ctx.KVStore(k.storeKey)

	srcBucketInfo, found := k.GetBucketInfo(ctx, srcBucketName)
	if !found {
		return sdkmath.ZeroUint(), errors.Wrapf(storagetypes.ErrNoSuchBucket, "src bucket name (%s)", srcBucketName)
	}

	dstBucketInfo, found := k.GetBucketInfo(ctx, dstBucketName)
	if !found {
		return sdkmath.ZeroUint(), errors.Wrapf(storagetypes.ErrNoSuchBucket, "dst bucket name (%s)", dstBucketName)
	}

	dstPrimarySP := k.MustGetPrimarySPForBucket(ctx, dstBucketInfo)

	err := dstBucketInfo.CheckBucketStatus()
	if err != nil {
		return sdkmath.ZeroUint(), err
	}

	srcObjectInfo, found := k.GetObjectInfo(ctx, srcBucketName, srcObjectName)
	if !found {
		return sdkmath.ZeroUint(), errors.Wrapf(storagetypes.ErrNoSuchObject, "src object name (%s)", srcObjectName)
	}

	if srcObjectInfo.SourceType != opts.SourceType {
		return sdkmath.ZeroUint(), storagetypes.ErrSourceTypeMismatch
	}

	if srcObjectInfo.IsUpdating {
		return sdkmath.ZeroUint(), storagetypes.ErrAccessDenied.Wrapf("the object is being updated, can not be copied")
	}

	// check permission
	effect := k.VerifyObjectPermission(ctx, srcBucketInfo, srcObjectInfo, operator, permtypes.ACTION_COPY_OBJECT)
	if effect != permtypes.EFFECT_ALLOW {
		return sdkmath.ZeroUint(), storagetypes.ErrAccessDenied.Wrapf("The operator("+
			"%s) has no CopyObject permission of the bucket(%s), object(%s)",
			operator.String(), srcObjectInfo.BucketName, srcObjectInfo.ObjectName)
	}

	err = k.VerifySP(ctx, dstPrimarySP, operator)
	if err != nil {
		return sdkmath.ZeroUint(), err
	}

	// check payload size, the empty object doesn't need sealed
	var objectStatus storagetypes.ObjectStatus
	if srcObjectInfo.PayloadSize == 0 {
		// empty object does not interact with sp
		objectStatus = storagetypes.OBJECT_STATUS_SEALED
	} else {
		objectStatus = storagetypes.OBJECT_STATUS_CREATED
	}

	objectInfo := storagetypes.ObjectInfo{
		Owner:          operator.String(),
		BucketName:     dstBucketInfo.BucketName,
		ObjectName:     dstObjectName,
		PayloadSize:    srcObjectInfo.PayloadSize,
		Visibility:     opts.Visibility,
		ContentType:    srcObjectInfo.ContentType,
		CreateAt:       ctx.BlockTime().Unix(),
		Id:             k.GenNextObjectID(ctx),
		ObjectStatus:   objectStatus,
		RedundancyType: srcObjectInfo.RedundancyType,
		SourceType:     opts.SourceType,
		Checksums:      srcObjectInfo.Checksums,
	}

	if srcObjectInfo.PayloadSize == 0 {
		_, err := k.SealEmptyObjectOnVirtualGroup(ctx, dstBucketInfo, &objectInfo)
		if err != nil {
			return sdkmath.ZeroUint(), err
		}
	} else {
		err = k.LockObjectStoreFee(ctx, dstBucketInfo, &objectInfo)
		if err != nil {
			return sdkmath.ZeroUint(), err
		}
	}

	bbz := k.cdc.MustMarshal(dstBucketInfo)
	store.Set(storagetypes.GetBucketByIDKey(dstBucketInfo.Id), bbz)
	if objectInfo.ObjectStatus == storagetypes.OBJECT_STATUS_CREATED {
		k.IncreaseLockedObjectCount(ctx, dstBucketInfo.Id)
	}

	obz := k.cdc.MustMarshal(&objectInfo)
	store.Set(storagetypes.GetObjectKey(dstBucketName, dstObjectName), k.objectSeq.EncodeSequence(objectInfo.Id))
	store.Set(storagetypes.GetObjectByIDKey(objectInfo.Id), obz)

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventCopyObject{
		Operator:            operator.String(),
		SrcBucketName:       srcObjectInfo.BucketName,
		SrcObjectName:       srcObjectInfo.ObjectName,
		DstBucketName:       objectInfo.BucketName,
		DstObjectName:       objectInfo.ObjectName,
		SrcObjectId:         srcObjectInfo.Id,
		DstObjectId:         objectInfo.Id,
		LocalVirtualGroupId: objectInfo.LocalVirtualGroupId,
	}); err != nil {
		return sdkmath.ZeroUint(), err
	}
	return objectInfo.Id, nil
}

func (k Keeper) RejectSealObject(ctx sdk.Context, operator sdk.AccAddress, bucketName, objectName string) error {
	store := ctx.KVStore(k.storeKey)
	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}
	objectInfo, found := k.GetObjectInfo(ctx, bucketName, objectName)
	if !found {
		return storagetypes.ErrNoSuchObject
	}
	spInState := k.MustGetPrimarySPForBucket(ctx, bucketInfo)
	if objectInfo.ObjectStatus != storagetypes.OBJECT_STATUS_CREATED && !objectInfo.IsUpdating {
		return storagetypes.ErrObjectNotCreated.Wrapf("Object status: %s", objectInfo.ObjectStatus.String())
	}

	sp, found := k.spKeeper.GetStorageProviderBySealAddr(ctx, operator)
	if !found {
		return errors.Wrapf(storagetypes.ErrNoSuchStorageProvider, "SP seal address: %s", operator.String())
	}
	if sp.Status != sptypes.STATUS_IN_SERVICE && sp.Status != sptypes.STATUS_IN_MAINTENANCE {
		return sptypes.ErrStorageProviderNotInService
	}
	if sp.Id != spInState.Id {
		return errors.Wrapf(storagetypes.ErrAccessDenied, "Only allowed primary SP to do reject seal object")
	}
	forUpdate := objectInfo.IsUpdating
	if forUpdate {
		shadowObjectInfo := k.MustGetShadowObjectInfo(ctx, bucketName, objectName)
		err := k.UnlockShadowObjectFeeAndDeleteShadowObjectInfo(ctx, bucketInfo, shadowObjectInfo, objectName)
		if err != nil {
			return err
		}
		// only this field need to revert
		objectInfo.IsUpdating = false
		obz := k.cdc.MustMarshal(objectInfo)
		store.Set(storagetypes.GetObjectByIDKey(objectInfo.Id), obz)
	} else {
		err := k.UnlockObjectStoreFee(ctx, bucketInfo, objectInfo)
		if err != nil {
			return err
		}
		bbz := k.cdc.MustMarshal(bucketInfo)
		store.Set(storagetypes.GetBucketByIDKey(bucketInfo.Id), bbz)
		store.Delete(storagetypes.GetObjectKey(bucketName, objectName))
		store.Delete(storagetypes.GetObjectByIDKey(objectInfo.Id))
	}

	k.DecreaseLockedObjectCount(ctx, bucketInfo.Id)

	return ctx.EventManager().EmitTypedEvents(&storagetypes.EventRejectSealObject{
		Operator:   operator.String(),
		BucketName: bucketInfo.BucketName,
		ObjectName: objectInfo.ObjectName,
		ObjectId:   objectInfo.Id,
		ForUpdate:  forUpdate,
	})
}

func (k Keeper) DiscontinueObject(ctx sdk.Context, operator sdk.AccAddress, bucketName string, objectIDs []sdkmath.Uint, reason string) error {
	sp, found := k.spKeeper.GetStorageProviderByGcAddr(ctx, operator)
	if !found {
		return storagetypes.ErrNoSuchStorageProvider.Wrapf("SP operator address: %s", operator.String())
	}
	if sp.Status != sptypes.STATUS_IN_SERVICE {
		return sptypes.ErrStorageProviderNotInService
	}

	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}
	if bucketInfo.BucketStatus == storagetypes.BUCKET_STATUS_DISCONTINUED {
		return storagetypes.ErrInvalidBucketStatus
	}
	spInState := k.MustGetPrimarySPForBucket(ctx, bucketInfo)

	if sp.Id != spInState.Id {
		swapInInfo, found := k.virtualGroupKeeper.GetSwapInInfo(ctx, bucketInfo.GlobalVirtualGroupFamilyId, virtualgroupmoduletypes.NoSpecifiedGVGId)
		if !found || swapInInfo.TargetSpId != spInState.Id || swapInInfo.SuccessorSpId != sp.Id {
			return errors.Wrapf(storagetypes.ErrAccessDenied, "the sp is not allowed to do discontinue objects")
		}
	}

	count := k.GetDiscontinueObjectCount(ctx, operator)
	max := k.DiscontinueObjectMax(ctx)
	if count+uint64(len(objectIDs)) > max {
		return storagetypes.ErrNoMoreDiscontinue.Wrapf("only %d objects can be requested in this window", max-count)
	}

	store := ctx.KVStore(k.storeKey)
	for _, objectID := range objectIDs {
		object, found := k.GetObjectInfoById(ctx, objectID)
		if !found {
			return storagetypes.ErrInvalidObjectIDs.Wrapf("object not found, id: %s", objectID)
		}
		if object.BucketName != bucketName {
			return storagetypes.ErrInvalidObjectIDs.Wrapf("object %s should in bucket: %s", objectID, bucketName)
		}
		if object.ObjectStatus != storagetypes.OBJECT_STATUS_SEALED && object.ObjectStatus != storagetypes.OBJECT_STATUS_CREATED {
			return storagetypes.ErrInvalidObjectIDs.Wrapf("object %s should in created or sealed status", objectID)
		}

		// remember object status
		k.saveDiscontinueObjectStatus(ctx, object)

		// update object status
		object.ObjectStatus = storagetypes.OBJECT_STATUS_DISCONTINUED
		store.Set(storagetypes.GetObjectByIDKey(object.Id), k.cdc.MustMarshal(object))
	}

	deleteAt := ctx.BlockTime().Unix() + k.DiscontinueConfirmPeriod(ctx)
	k.AppendDiscontinueObjectIds(ctx, deleteAt, objectIDs)
	k.SetDiscontinueObjectCount(ctx, operator, count+uint64(len(objectIDs)))

	events := make([]proto.Message, 0)
	for _, objectID := range objectIDs {
		events = append(events, &storagetypes.EventDiscontinueObject{
			BucketName: bucketName,
			ObjectId:   objectID,
			Reason:     reason,
			DeleteAt:   deleteAt,
		})
	}
	if err := ctx.EventManager().EmitTypedEvents(events...); err != nil {
		return err
	}
	return nil
}

func (k Keeper) UpdateObjectInfo(ctx sdk.Context, operator sdk.AccAddress, bucketName, objectName string, visibility storagetypes.VisibilityType) error {
	store := ctx.KVStore(k.storeKey)

	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}

	objectInfo, found := k.GetObjectInfo(ctx, bucketName, objectName)
	if !found {
		return storagetypes.ErrNoSuchObject
	}

	// check permission
	effect := k.VerifyObjectPermission(ctx, bucketInfo, objectInfo, operator, permtypes.ACTION_UPDATE_OBJECT_INFO)
	if effect != permtypes.EFFECT_ALLOW {
		return storagetypes.ErrAccessDenied.Wrapf("The operator(%s) has no UpdateObjectInfo permission of the bucket(%s), object(%s)",
			operator.String(), bucketName, objectName)
	}

	objectInfo.Visibility = visibility

	obz := k.cdc.MustMarshal(objectInfo)
	store.Set(storagetypes.GetObjectByIDKey(objectInfo.Id), obz)

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventUpdateObjectInfo{
		Operator:   operator.String(),
		BucketName: bucketName,
		ObjectName: objectName,
		Visibility: visibility,
		ObjectId:   objectInfo.Id,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) CreateGroup(
	ctx sdk.Context, owner sdk.AccAddress,
	groupName string, opts storagetypes.CreateGroupOptions,
) (sdkmath.Uint, error) {
	store := ctx.KVStore(k.storeKey)

	groupInfo := storagetypes.GroupInfo{
		Owner:      owner.String(),
		SourceType: opts.SourceType,
		Id:         k.GenNextGroupId(ctx),
		GroupName:  groupName,
		Extra:      opts.Extra,
	}

	// Can not create a group with the same name.
	groupKey := storagetypes.GetGroupKey(owner, groupName)
	if store.Has(groupKey) {
		return sdkmath.ZeroUint(), storagetypes.ErrGroupAlreadyExists
	}

	gbz := k.cdc.MustMarshal(&groupInfo)
	store.Set(groupKey, k.groupSeq.EncodeSequence(groupInfo.Id))
	store.Set(storagetypes.GetGroupByIDKey(groupInfo.Id), gbz)

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventCreateGroup{
		Owner:      groupInfo.Owner,
		GroupName:  groupInfo.GroupName,
		GroupId:    groupInfo.Id,
		SourceType: groupInfo.SourceType,
		Extra:      opts.Extra,
	}); err != nil {
		return sdkmath.ZeroUint(), err
	}

	// Mint group nft token and send to receiver
	_, err := k.CallEVM(
		ctx,
		contracts.ERC721NonTransferableContract.ABI,
		contracts.GroupControlHubAddress,
		contracts.GroupERC721TokenAddress,
		true,
		"mint",
		ecommon.HexToAddress(groupInfo.Owner),
		groupInfo.Id.BigInt(),
	)
	if err != nil {
		return sdkmath.ZeroUint(), err
	}

	return groupInfo.Id, nil
}

func (k Keeper) SetGroupInfo(ctx sdk.Context, groupInfo *storagetypes.GroupInfo) {
	store := ctx.KVStore(k.storeKey)

	gbz := k.cdc.MustMarshal(groupInfo)
	store.Set(storagetypes.GetGroupByIDKey(groupInfo.Id), gbz)
}

func (k Keeper) GetGroupInfo(ctx sdk.Context, ownerAddr sdk.AccAddress,
	groupName string,
) (*storagetypes.GroupInfo, bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(storagetypes.GetGroupKey(ownerAddr, groupName))
	if bz == nil {
		return nil, false
	}

	return k.GetGroupInfoById(ctx, k.groupSeq.DecodeSequence(bz))
}

func (k Keeper) GetGroupInfoById(ctx sdk.Context, groupId sdkmath.Uint) (*storagetypes.GroupInfo, bool) { //nolint
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(storagetypes.GetGroupByIDKey(groupId))
	if bz == nil {
		return nil, false
	}

	var groupInfo storagetypes.GroupInfo
	k.cdc.MustUnmarshal(bz, &groupInfo)
	return &groupInfo, true
}

func (k Keeper) DeleteGroup(ctx sdk.Context, operator sdk.AccAddress, groupName string, opts storagetypes.DeleteGroupOptions) error {
	store := ctx.KVStore(k.storeKey)

	groupInfo, found := k.GetGroupInfo(ctx, operator, groupName)
	if !found {
		return storagetypes.ErrNoSuchGroup
	}
	if groupInfo.SourceType != opts.SourceType {
		return storagetypes.ErrSourceTypeMismatch
	}
	// check permission
	effect := k.VerifyGroupPermission(ctx, groupInfo, operator, permtypes.ACTION_DELETE_GROUP)
	if effect != permtypes.EFFECT_ALLOW {
		return storagetypes.ErrAccessDenied.Wrapf(
			"The operator(%s) has no DeleteGroup permission of the group(%s), owner(%s)",
			operator.String(), groupInfo.GroupName, groupInfo.Owner)
	}
	// Note: Delete group does not require the group is empty. The group member will be deleted by on-chain GC.
	store.Delete(storagetypes.GetGroupKey(operator, groupName))
	store.Delete(storagetypes.GetGroupByIDKey(groupInfo.Id))

	if err := k.appendResourceIDForGarbageCollection(ctx, resource.RESOURCE_TYPE_GROUP, groupInfo.Id); err != nil {
		return err
	}

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventDeleteGroup{
		Owner:     groupInfo.Owner,
		GroupName: groupInfo.GroupName,
		GroupId:   groupInfo.Id,
	}); err != nil {
		return err
	}

	if _, err := k.CallEVM(
		ctx,
		contracts.ERC721NonTransferableContract.ABI,
		contracts.GroupControlHubAddress,
		contracts.GroupERC721TokenAddress,
		true,
		"burn",
		groupInfo.Id.BigInt(),
	); err != nil {
		return err
	}

	return nil
}

func (k Keeper) LeaveGroup(
	ctx sdk.Context, member, owner sdk.AccAddress,
	groupName string, opts storagetypes.LeaveGroupOptions,
) error {
	groupInfo, found := k.GetGroupInfo(ctx, owner, groupName)
	if !found {
		return storagetypes.ErrNoSuchGroup
	}
	if groupInfo.SourceType != opts.SourceType {
		return storagetypes.ErrSourceTypeMismatch
	}

	// Note: Delete group does not require the group is empty. The group member will be deleted by on-chain GC.
	err := k.permKeeper.RemoveGroupMember(ctx, groupInfo.Id, member)
	if err != nil {
		return err
	}

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventLeaveGroup{
		MemberAddress: member.String(),
		Owner:         groupInfo.Owner,
		GroupName:     groupInfo.GroupName,
		GroupId:       groupInfo.Id,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) UpdateGroupMember(ctx sdk.Context, operator sdk.AccAddress, groupInfo *storagetypes.GroupInfo, opts storagetypes.UpdateGroupMemberOptions) error {
	if groupInfo.SourceType != opts.SourceType {
		return storagetypes.ErrSourceTypeMismatch
	}

	// check permission
	effect := k.VerifyGroupPermission(ctx, groupInfo, operator, permtypes.ACTION_UPDATE_GROUP_MEMBER)
	if effect != permtypes.EFFECT_ALLOW {
		return storagetypes.ErrAccessDenied.Wrapf(
			"The operator(%s) has no UpdateGroupMember permission of the group(%s), operator(%s)",
			operator.String(), groupInfo.GroupName, groupInfo.Owner)
	}

	addedMembersDetailEvent := make([]*storagetypes.EventGroupMemberDetail, 0, len(opts.MembersToAdd))
	for i := range opts.MembersToAdd {
		memberAcc, err := sdk.AccAddressFromHexUnsafe(opts.MembersToAdd[i])
		if err != nil {
			return err
		}

		err = k.permKeeper.AddGroupMember(ctx, groupInfo.Id, memberAcc, opts.MembersExpirationToAdd[i])
		if err != nil {
			return err
		}

		addedMembersDetailEvent = append(addedMembersDetailEvent, &storagetypes.EventGroupMemberDetail{
			Member:         opts.MembersToAdd[i],
			ExpirationTime: opts.MembersExpirationToAdd[i],
		})
	}

	for _, member := range opts.MembersToDelete {
		memberAcc, err := sdk.AccAddressFromHexUnsafe(member)
		if err != nil {
			return err
		}
		err = k.permKeeper.RemoveGroupMember(ctx, groupInfo.Id, memberAcc)
		if err != nil {
			return err
		}

	}
	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventUpdateGroupMember{
		Operator:        operator.String(),
		Owner:           groupInfo.Owner,
		GroupName:       groupInfo.GroupName,
		GroupId:         groupInfo.Id,
		MembersToAdd:    addedMembersDetailEvent,
		MembersToDelete: opts.MembersToDelete,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) RenewGroupMember(ctx sdk.Context, operator sdk.AccAddress, groupInfo *storagetypes.GroupInfo, opts storagetypes.RenewGroupMemberOptions) error {
	if groupInfo.SourceType != opts.SourceType {
		return storagetypes.ErrSourceTypeMismatch
	}

	// check permission
	effect := k.VerifyGroupPermission(ctx, groupInfo, operator, permtypes.ACTION_UPDATE_GROUP_MEMBER)
	if effect != permtypes.EFFECT_ALLOW {
		return storagetypes.ErrAccessDenied.Wrapf(
			"The operator(%s) has no UpdateGroupMember permission of the group(%s), operator(%s)",
			operator.String(), groupInfo.GroupName, groupInfo.Owner)
	}

	eventMembersDetail := make([]*storagetypes.EventGroupMemberDetail, 0, len(opts.Members))
	for i := range opts.Members {
		member := opts.Members[i]
		memberAcc, err := sdk.AccAddressFromHexUnsafe(member)
		if err != nil {
			return err
		}

		groupMember, found := k.permKeeper.GetGroupMember(ctx, groupInfo.Id, memberAcc)
		if !found {
			err = k.permKeeper.AddGroupMember(ctx, groupInfo.Id, memberAcc, opts.MembersExpiration[i])
			if err != nil {
				return err
			}
		} else {
			k.permKeeper.UpdateGroupMember(ctx, groupInfo.Id, memberAcc, groupMember.Id, opts.MembersExpiration[i])
		}

		eventMembersDetail = append(eventMembersDetail, &storagetypes.EventGroupMemberDetail{
			Member:         member,
			ExpirationTime: opts.MembersExpiration[i],
		})
	}

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventRenewGroupMember{
		Operator:   operator.String(),
		Owner:      groupInfo.Owner,
		GroupName:  groupInfo.GroupName,
		GroupId:    groupInfo.Id,
		SourceType: groupInfo.SourceType,
		Members:    eventMembersDetail,
	}); err != nil {
		return err
	}

	return nil
}

func (k Keeper) UpdateGroupExtra(ctx sdk.Context, operator sdk.AccAddress, groupInfo *storagetypes.GroupInfo, extra string) error {
	// check permission
	effect := k.VerifyGroupPermission(ctx, groupInfo, operator, permtypes.ACTION_UPDATE_GROUP_EXTRA)
	if effect != permtypes.EFFECT_ALLOW {
		return storagetypes.ErrAccessDenied.Wrapf(
			"The operator(%s) has no UpdateGroupExtra permission of the group(%s), operator(%s)",
			operator.String(), groupInfo.GroupName, groupInfo.Owner)
	}

	if extra != groupInfo.Extra {
		groupInfo.Extra = extra
		obz := k.cdc.MustMarshal(groupInfo)
		ctx.KVStore(k.storeKey).Set(storagetypes.GetGroupByIDKey(groupInfo.Id), obz)
	}

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventUpdateGroupExtra{
		Operator:  operator.String(),
		Owner:     groupInfo.Owner,
		GroupName: groupInfo.GroupName,
		GroupId:   groupInfo.Id,
		Extra:     extra,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) VerifySPAndSignature(_ sdk.Context, sp *sptypes.StorageProvider, sigData, signature []byte, operator sdk.AccAddress) error {
	if sp.Status != sptypes.STATUS_IN_SERVICE && !k.fromSpMaintenanceAcct(sp, operator) {
		return sptypes.ErrStorageProviderNotInService
	}
	approvalAccAddress := sdk.MustAccAddressFromHex(sp.ApprovalAddress)

	err := types.VerifySignature(approvalAccAddress, crypto.Keccak256(sigData), signature)
	if err != nil {
		return errors.Wrapf(storagetypes.ErrInvalidApproval, "verify signature error: %s", err)
	}
	return nil
}

func (k Keeper) VerifySP(_ sdk.Context, sp *sptypes.StorageProvider, operator sdk.AccAddress) error {
	if sp.Status != sptypes.STATUS_IN_SERVICE && !k.fromSpMaintenanceAcct(sp, operator) {
		return sptypes.ErrStorageProviderNotInService
	}
	return nil
}

func (k Keeper) GenNextBucketId(ctx sdk.Context) sdkmath.Uint { //nolint
	store := ctx.KVStore(k.storeKey)

	seq := k.bucketSeq.NextVal(store)
	return seq
}

func (k Keeper) GenNextObjectID(ctx sdk.Context) sdkmath.Uint {
	store := ctx.KVStore(k.storeKey)

	seq := k.objectSeq.NextVal(store)
	return seq
}

func (k Keeper) GenNextGroupId(ctx sdk.Context) sdkmath.Uint { //nolint
	store := ctx.KVStore(k.storeKey)

	seq := k.groupSeq.NextVal(store)
	return seq
}

func (k Keeper) isNonEmptyBucket(ctx sdk.Context, bucketName string) bool {
	store := ctx.KVStore(k.storeKey)
	objectStore := prefix.NewStore(store, storagetypes.GetObjectKeyOnlyBucketPrefix(bucketName))

	iter := objectStore.Iterator(nil, nil)
	return iter.Valid()
}

func (k Keeper) GetDiscontinueObjectCount(ctx sdk.Context, operator sdk.AccAddress) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storagetypes.DiscontinueObjectCountPrefix)
	bz := store.Get(operator.Bytes())

	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetDiscontinueObjectCount(ctx sdk.Context, operator sdk.AccAddress, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storagetypes.DiscontinueObjectCountPrefix)

	countBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(countBytes, count)

	store.Set(operator.Bytes(), countBytes)
}

func (k Keeper) ClearDiscontinueObjectCount(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storagetypes.DiscontinueObjectCountPrefix)

	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}
}

func (k Keeper) AppendDiscontinueObjectIds(ctx sdk.Context, timestamp int64, objectIds []storagetypes.Uint) { //nolint
	store := ctx.KVStore(k.storeKey)
	key := storagetypes.GetDiscontinueObjectIdsKey(timestamp)
	bz := store.Get(key)
	if bz != nil {
		var existedIDs storagetypes.Ids
		k.cdc.MustUnmarshal(bz, &existedIDs)
		objectIds = append(existedIDs.Id, objectIds...)
	}

	store.Set(key, k.cdc.MustMarshal(&storagetypes.Ids{Id: objectIds}))
}

func (k Keeper) DeleteDiscontinueObjectsUntil(ctx sdk.Context, timestamp int64, maxObjectsToDelete uint64) (deleted uint64, err error) {
	store := ctx.KVStore(k.storeKey)
	key := storagetypes.GetDiscontinueObjectIdsKey(timestamp)
	iterator := store.Iterator(storagetypes.DiscontinueObjectIDsPrefix, storetypes.InclusiveEndBytes(key))
	defer iterator.Close()

	deleted = uint64(0)
	for ; iterator.Valid(); iterator.Next() {
		if deleted >= maxObjectsToDelete {
			break
		}
		var ids storagetypes.Ids
		k.cdc.MustUnmarshal(iterator.Value(), &ids)

		left := make([]storagetypes.Uint, 0)
		for _, id := range ids.Id {
			if deleted >= maxObjectsToDelete {
				left = append(left, id)
				continue
			}

			err = k.ForceDeleteObject(ctx, id)
			if err != nil {
				ctx.Logger().Error("delete object error", "err", err, "id", id, "height", ctx.BlockHeight())
				return deleted, err
			}
			deleted++
		}
		if len(left) > 0 {
			store.Set(iterator.Key(), k.cdc.MustMarshal(&storagetypes.Ids{Id: left}))
		} else {
			store.Delete(iterator.Key())
		}
	}

	return deleted, nil
}

func (k Keeper) GetDiscontinueBucketCount(ctx sdk.Context, operator sdk.AccAddress) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storagetypes.DiscontinueBucketCountPrefix)
	bz := store.Get(operator.Bytes())

	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetDiscontinueBucketCount(ctx sdk.Context, operator sdk.AccAddress, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storagetypes.DiscontinueBucketCountPrefix)

	countBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(countBytes, count)

	store.Set(operator.Bytes(), countBytes)
}

func (k Keeper) ClearDiscontinueBucketCount(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storagetypes.DiscontinueBucketCountPrefix)

	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}
}

func (k Keeper) appendDiscontinueBucketIDs(ctx sdk.Context, timestamp int64, bucketIDs []storagetypes.Uint) {
	store := ctx.KVStore(k.storeKey)
	key := storagetypes.GetDiscontinueBucketIDsKey(timestamp)

	bz := store.Get(key)
	if bz != nil {
		var existedIDs storagetypes.Ids
		k.cdc.MustUnmarshal(bz, &existedIDs)
		bucketIDs = append(existedIDs.Id, bucketIDs...)
	}

	store.Set(key, k.cdc.MustMarshal(&storagetypes.Ids{Id: bucketIDs}))
}

func (k Keeper) DeleteDiscontinueBucketsUntil(ctx sdk.Context, timestamp int64, maxToDelete uint64) (uint64, error) {
	store := ctx.KVStore(k.storeKey)
	key := storagetypes.GetDiscontinueBucketIDsKey(timestamp)
	iterator := store.Iterator(storagetypes.DiscontinueBucketIDsPrefix, storetypes.InclusiveEndBytes(key))
	defer iterator.Close()

	deleted := uint64(0)
	for ; iterator.Valid(); iterator.Next() {
		if deleted >= maxToDelete {
			break
		}
		var ids storagetypes.Ids
		k.cdc.MustUnmarshal(iterator.Value(), &ids)

		left := make([]storagetypes.Uint, 0)
		for _, id := range ids.Id {
			if deleted >= maxToDelete {
				left = append(left, id)
				continue
			}

			bucketDeleted, objectDeleted, err := k.ForceDeleteBucket(ctx, id, maxToDelete-deleted)
			if err != nil {
				ctx.Logger().Error("force delete bucket error", "err", err, "id", id, "height", ctx.BlockHeight())
				return deleted, err
			}
			deleted += objectDeleted

			if !bucketDeleted {
				left = append(left, id)
			} else {
				deleted++
			}
		}
		if len(left) > 0 {
			store.Set(iterator.Key(), k.cdc.MustMarshal(&storagetypes.Ids{Id: left}))
		} else {
			store.Delete(iterator.Key())
		}
	}

	return deleted, nil
}

func (k Keeper) saveDiscontinueObjectStatus(ctx sdk.Context, object *storagetypes.ObjectInfo) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 4)
	binary.BigEndian.PutUint32(bz, uint32(object.ObjectStatus))
	store.Set(storagetypes.GetDiscontinueObjectStatusKey(object.Id), bz)
}

func (k Keeper) getAndDeleteDiscontinueObjectStatus(ctx sdk.Context, objectID storagetypes.Uint) (storagetypes.ObjectStatus, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(storagetypes.GetDiscontinueObjectStatusKey(objectID))
	if bz == nil {
		ctx.Logger().Error("fail to get discontinued object status", "object id", objectID)
		return storagetypes.OBJECT_STATUS_DISCONTINUED, errors.Wrapf(storagetypes.ErrInvalidObjectStatus, "object id: %s", objectID)
	}
	status := int32(binary.BigEndian.Uint32(bz))
	store.Delete(storagetypes.GetDiscontinueObjectStatusKey(objectID)) // remove it at the same time
	return storagetypes.ObjectStatus(status), nil
}

func (k Keeper) appendResourceIDForGarbageCollection(ctx sdk.Context, resourceType resource.ResourceType, resourceID sdkmath.Uint) error {
	if !k.permKeeper.ExistAccountPolicyForResource(ctx, resourceType, resourceID) &&
		!k.permKeeper.ExistGroupPolicyForResource(ctx, resourceType, resourceID) {

		if resourceType != resource.RESOURCE_TYPE_GROUP ||
			(resourceType == resource.RESOURCE_TYPE_GROUP && !k.permKeeper.ExistGroupMemberForGroup(ctx, resourceID)) {
			return nil
		}
	}
	tStore := ctx.TransientStore(k.tStoreKey)
	var deleteInfo storagetypes.DeleteInfo
	if !tStore.Has(storagetypes.CurrentBlockDeleteStalePoliciesKey) {
		deleteInfo = storagetypes.DeleteInfo{
			BucketIds: &storagetypes.Ids{},
			ObjectIds: &storagetypes.Ids{},
			GroupIds:  &storagetypes.Ids{},
		}
	} else {
		bz := tStore.Get(storagetypes.CurrentBlockDeleteStalePoliciesKey)
		k.cdc.MustUnmarshal(bz, &deleteInfo)
	}
	switch resourceType {
	case resource.RESOURCE_TYPE_BUCKET:
		bucketIDs := deleteInfo.BucketIds.Id
		bucketIDs = append(bucketIDs, resourceID)
		deleteInfo.BucketIds = &storagetypes.Ids{Id: bucketIDs}
	case resource.RESOURCE_TYPE_OBJECT:
		objectIDs := deleteInfo.ObjectIds.Id
		objectIDs = append(objectIDs, resourceID)
		deleteInfo.ObjectIds = &storagetypes.Ids{Id: objectIDs}
	case resource.RESOURCE_TYPE_GROUP:
		groupIDs := deleteInfo.GroupIds.Id
		groupIDs = append(groupIDs, resourceID)
		deleteInfo.GroupIds = &storagetypes.Ids{Id: groupIDs}
	default:
		return storagetypes.ErrInvalidResource
	}
	tStore.Set(storagetypes.CurrentBlockDeleteStalePoliciesKey, k.cdc.MustMarshal(&deleteInfo))
	return nil
}

func (k Keeper) PersistDeleteInfo(ctx sdk.Context) {
	tStore := ctx.TransientStore(k.tStoreKey)
	if !tStore.Has(storagetypes.CurrentBlockDeleteStalePoliciesKey) {
		return
	}
	bz := tStore.Get(storagetypes.CurrentBlockDeleteStalePoliciesKey)
	deleteInfo := &storagetypes.DeleteInfo{}
	k.cdc.MustUnmarshal(bz, deleteInfo)

	// persist current block stale permission info to store if exists
	if !deleteInfo.IsEmpty() {
		store := ctx.KVStore(k.storeKey)
		store.Set(storagetypes.GetDeleteStalePoliciesKey(ctx.BlockHeight()), bz)
		_ = ctx.EventManager().EmitTypedEvents(&storagetypes.EventStalePolicyCleanup{
			BlockNum:   ctx.BlockHeight(),
			DeleteInfo: deleteInfo,
		})
	}
}

func (k Keeper) GarbageCollectResourcesStalePolicy(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	deleteStalePoliciesPrefixStore := prefix.NewStore(store, storagetypes.DeleteStalePoliciesPrefix)

	iterator := deleteStalePoliciesPrefixStore.Iterator(nil, nil)
	defer iterator.Close()

	maxCleanup := k.StalePolicyCleanupMax(ctx)

	var deletedTotal uint64
	var done bool

	for ; iterator.Valid(); iterator.Next() {
		deleteInfo := &storagetypes.DeleteInfo{}
		k.cdc.MustUnmarshal(iterator.Value(), deleteInfo)
		deletedTotal, done = k.garbageCollectionForResource(ctx, deleteStalePoliciesPrefixStore, iterator, deleteInfo, resource.RESOURCE_TYPE_OBJECT, deleteInfo.ObjectIds, maxCleanup, deletedTotal)
		if !done {
			return
		}
		deleteInfo.ObjectIds = nil
		deletedTotal, done = k.garbageCollectionForResource(ctx, deleteStalePoliciesPrefixStore, iterator, deleteInfo, resource.RESOURCE_TYPE_BUCKET, deleteInfo.BucketIds, maxCleanup, deletedTotal)
		if !done {
			return
		}
		deleteInfo.BucketIds = nil
		deletedTotal, done = k.garbageCollectionForResource(ctx, deleteStalePoliciesPrefixStore, iterator, deleteInfo, resource.RESOURCE_TYPE_GROUP, deleteInfo.GroupIds, maxCleanup, deletedTotal)
		if !done {
			return
		}
		deleteInfo.GroupIds = nil
		// the specified block height(iterator-key)'s stale resource permission metadata is purged
		if deleteInfo.IsEmpty() {
			deleteStalePoliciesPrefixStore.Delete(iterator.Key())
		}
	}
}

func (k Keeper) garbageCollectionForResource(ctx sdk.Context, deleteStalePoliciesPrefixStore prefix.Store, iterator storetypes.Iterator,
	deleteInfo *storagetypes.DeleteInfo, resourceType resource.ResourceType, resourceIDs *storagetypes.Ids, maxCleanup, deletedTotal uint64,
) (uint64, bool) {
	var done bool
	if resourceIDs != nil && len(resourceIDs.Id) > 0 {
		ids := resourceIDs.Id
		temp := ids
		for idx, id := range ids {
			deletedTotal, done = k.permKeeper.ForceDeleteAccountPolicyForResource(ctx, maxCleanup, deletedTotal, resourceType, id)
			if !done {
				resourceIDs.Id = temp
				deleteStalePoliciesPrefixStore.Set(iterator.Key(), k.cdc.MustMarshal(deleteInfo))
				return deletedTotal, false
			}
			if resourceType == resource.RESOURCE_TYPE_GROUP {
				deletedTotal, done = k.permKeeper.ForceDeleteGroupMembers(ctx, maxCleanup, deletedTotal, id)
				if !done {
					deleteInfo.GroupIds.Id = temp
					deleteStalePoliciesPrefixStore.Set(iterator.Key(), k.cdc.MustMarshal(deleteInfo))
					return deletedTotal, false
				}
				// no need to deal with group policy when resource type is group
				continue
			}
			deletedTotal, done = k.permKeeper.ForceDeleteGroupPolicyForResource(ctx, maxCleanup, deletedTotal, resourceType, id)
			if !done {
				resourceIDs.Id = temp
				deleteStalePoliciesPrefixStore.Set(iterator.Key(), k.cdc.MustMarshal(deleteInfo))
				return deletedTotal, false
			}
			//  remove current resource id from list of ids to be deleted
			temp = ids[idx+1:]
		}
	}
	return deletedTotal, true
}

func (k Keeper) MigrateBucket(ctx sdk.Context, operator sdk.AccAddress, bucketName string, dstPrimarySPID uint32, dstPrimarySPApproval *common.Approval, approvalBytes []byte) error {
	store := ctx.KVStore(k.storeKey)

	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}

	if !operator.Equals(sdk.MustAccAddressFromHex(bucketInfo.Owner)) {
		return storagetypes.ErrAccessDenied.Wrap("Only bucket owner can migrate bucket.")
	}

	if bucketInfo.BucketStatus == storagetypes.BUCKET_STATUS_MIGRATING {
		return storagetypes.ErrInvalidBucketStatus.Wrapf("The bucket already been migrating")
	}

	if bucketInfo.BucketStatus == storagetypes.BUCKET_STATUS_DISCONTINUED {
		return storagetypes.ErrInvalidBucketStatus.Wrapf("The discontinued bucket cannot be migrated")
	}

	srcSP := k.MustGetPrimarySPForBucket(ctx, bucketInfo)

	dstSP, found := k.spKeeper.GetStorageProvider(ctx, dstPrimarySPID)
	if !found {
		return sptypes.ErrStorageProviderNotFound.Wrapf("dst sp not found")
	}

	if srcSP.Id == dstSP.Id {
		return storagetypes.ErrMigrationBucketFailed.Wrapf("The dest sp must not be the origin sp.")
	}

	if !srcSP.IsInService() || !dstSP.IsInService() {
		return sptypes.ErrStorageProviderNotInService.Wrapf(
			"origin SP status: %s, dst SP status: %s", srcSP.Status.String(), dstSP.Status.String())
	}

	streamRecord, found := k.paymentKeeper.GetStreamRecord(ctx, sdk.MustAccAddressFromHex(bucketInfo.PaymentAddress))
	if !found || streamRecord.Status == paymenttypes.STREAM_ACCOUNT_STATUS_FROZEN {
		return paymenttypes.ErrInvalidStreamAccountStatus.Wrap("stream account is frozen")
	}

	// check approval
	if dstPrimarySPApproval.ExpiredHeight < (uint64)(ctx.BlockHeight()) {
		return storagetypes.ErrInvalidApproval.Wrap("dst primary sp approval timeout")
	}
	err := k.VerifySPAndSignature(ctx, dstSP, approvalBytes, dstPrimarySPApproval.Sig, operator)
	if err != nil {
		return err
	}

	isRateLimited := k.IsBucketRateLimited(ctx, bucketInfo.BucketName)
	if isRateLimited {
		return fmt.Errorf("bucket is rate limited: %s", bucketInfo.BucketName)
	}

	key := storagetypes.GetMigrationBucketKey(bucketInfo.Id)
	if store.Has(key) {
		panic("migration bucket key is existed.")
	}

	migrationBucketInfo := &storagetypes.MigrationBucketInfo{
		SrcSpId:                       srcSP.Id,
		DstSpId:                       dstSP.Id,
		SrcGlobalVirtualGroupFamilyId: bucketInfo.GlobalVirtualGroupFamilyId,
		BucketId:                      bucketInfo.Id,
	}

	bz := k.cdc.MustMarshal(migrationBucketInfo)
	store.Set(key, bz)

	bucketInfo.BucketStatus = storagetypes.BUCKET_STATUS_MIGRATING
	k.SetBucketInfo(ctx, bucketInfo)

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventMigrationBucket{
		Operator:       operator.String(),
		BucketName:     bucketName,
		BucketId:       bucketInfo.Id,
		DstPrimarySpId: dstSP.Id,
		Status:         bucketInfo.BucketStatus,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) CompleteMigrateBucket(ctx sdk.Context, operator sdk.AccAddress, bucketName string, gvgFamilyID uint32, gvgMappings []*storagetypes.GVGMapping) error {
	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}

	dstSP, found := k.spKeeper.GetStorageProviderByOperatorAddr(ctx, operator)
	if !found {
		return sptypes.ErrStorageProviderNotFound.Wrapf("dst SP not found.")
	}

	if bucketInfo.BucketStatus != storagetypes.BUCKET_STATUS_MIGRATING {
		return storagetypes.ErrInvalidBucketStatus.Wrapf("The bucket is not been migrating")
	}

	migrationBucketInfo, found := k.GetMigrationBucketInfo(ctx, bucketInfo.Id)
	if !found {
		return storagetypes.ErrMigrationBucketFailed.Wrapf("migration bucket info not found.")
	}

	if dstSP.Id != migrationBucketInfo.DstSpId {
		return storagetypes.ErrMigrationBucketFailed.Wrapf("dst sp info not match")
	}

	_, found = k.virtualGroupKeeper.GetGVGFamily(ctx, gvgFamilyID)
	if !found {
		return virtualgroupmoduletypes.ErrGVGFamilyNotExist
	}

	srcGvgFamily, found := k.virtualGroupKeeper.GetGVGFamily(ctx, bucketInfo.GlobalVirtualGroupFamilyId)
	if !found {
		return virtualgroupmoduletypes.ErrGVGFamilyNotExist
	}

	streamRecord, found := k.paymentKeeper.GetStreamRecord(ctx, sdk.MustAccAddressFromHex(bucketInfo.PaymentAddress))
	if !found || streamRecord.Status == paymenttypes.STREAM_ACCOUNT_STATUS_FROZEN {
		return paymenttypes.ErrInvalidStreamAccountStatus.Wrap("stream account is frozen")
	}

	sp := k.MustGetPrimarySPForBucket(ctx, bucketInfo)

	err := k.virtualGroupKeeper.SettleAndDistributeGVGFamily(ctx, sp, srcGvgFamily)
	if err != nil {
		return virtualgroupmoduletypes.ErrSettleFailed.Wrapf("settle gvg family failed, err: %s", err)
	}

	internalBucketInfo := k.MustGetInternalBucketInfo(ctx, bucketInfo.Id)
	err = k.UnChargeBucketReadStoreFee(ctx, bucketInfo, internalBucketInfo)
	if err != nil {
		return storagetypes.ErrMigrationBucketFailed.Wrapf("cancel charge bucket failed, err: %s", err)
	}

	bucketInfo.GlobalVirtualGroupFamilyId = gvgFamilyID

	// check secondary sp signature
	err = k.verifyGVGSignatures(ctx, bucketInfo.Id, dstSP, gvgMappings)
	if err != nil {
		return storagetypes.ErrMigrationBucketFailed.Wrapf("err: %s", err)
	}

	// rebinding gvg and lvg
	err = k.RebindingVirtualGroup(ctx, bucketInfo, internalBucketInfo, gvgMappings)
	if err != nil {
		return storagetypes.ErrMigrationBucketFailed.Wrapf("err: %s", err)
	}

	bucketInfo.BucketStatus = storagetypes.BUCKET_STATUS_CREATED
	k.SetBucketInfo(ctx, bucketInfo)
	k.DeleteMigrationBucketInfo(ctx, bucketInfo.Id)

	if err = ctx.EventManager().EmitTypedEvents(&storagetypes.EventCompleteMigrationBucket{
		Operator:                   operator.String(),
		BucketName:                 bucketName,
		BucketId:                   bucketInfo.Id,
		GlobalVirtualGroupFamilyId: gvgFamilyID,
		SrcPrimarySpId:             srcGvgFamily.PrimarySpId,
		Status:                     bucketInfo.BucketStatus,
	}); err != nil {
		return err
	}

	if err = ctx.EventManager().EmitTypedEvents(&storagetypes.EventUpdateBucketInfo{
		Operator:                   operator.String(),
		BucketName:                 bucketName,
		BucketId:                   bucketInfo.Id,
		ChargedReadQuota:           bucketInfo.ChargedReadQuota,
		PaymentAddress:             bucketInfo.PaymentAddress,
		Visibility:                 bucketInfo.Visibility,
		GlobalVirtualGroupFamilyId: bucketInfo.GlobalVirtualGroupFamilyId,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) CancelBucketMigration(ctx sdk.Context, operator sdk.AccAddress, bucketName string) error {
	store := ctx.KVStore(k.storeKey)
	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}

	if !operator.Equals(sdk.MustAccAddressFromHex(bucketInfo.Owner)) {
		return storagetypes.ErrAccessDenied.Wrap("Only bucket owner can cancel migrate bucket.")
	}

	if bucketInfo.BucketStatus != storagetypes.BUCKET_STATUS_MIGRATING {
		return storagetypes.ErrInvalidBucketStatus.Wrapf("The bucket is not been migrating")
	}

	key := storagetypes.GetMigrationBucketKey(bucketInfo.Id)
	if !store.Has(key) {
		return storagetypes.ErrMigrationBucketFailed.Wrapf("cancel migrate bucket failed due to the migrate bucket info not found.")
	}

	bucketInfo.BucketStatus = storagetypes.BUCKET_STATUS_CREATED
	k.SetBucketInfo(ctx, bucketInfo)
	store.Delete(key)

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventCancelMigrationBucket{
		Operator:   operator.String(),
		BucketName: bucketName,
		BucketId:   bucketInfo.Id,
		Status:     bucketInfo.BucketStatus,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) RejectBucketMigration(ctx sdk.Context, operator sdk.AccAddress, bucketName string) error {
	store := ctx.KVStore(k.storeKey)
	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}
	if bucketInfo.BucketStatus != storagetypes.BUCKET_STATUS_MIGRATING {
		return storagetypes.ErrInvalidBucketStatus.Wrapf("The bucket is not in migrating status")
	}

	migrationBucketInfo, found := k.GetMigrationBucketInfo(ctx, bucketInfo.Id)
	if !found {
		return storagetypes.ErrMigrationBucketFailed.Wrapf("reject bucket migration failed due to the migrate bucket info not found.")
	}

	sp := k.spKeeper.MustGetStorageProvider(ctx, migrationBucketInfo.DstSpId)
	if !sdk.MustAccAddressFromHex(sp.OperatorAddress).Equals(operator) {
		return storagetypes.ErrAccessDenied.Wrap("Only the dest SP can reject the bucket migration.")
	}

	bucketInfo.BucketStatus = storagetypes.BUCKET_STATUS_CREATED
	k.SetBucketInfo(ctx, bucketInfo)
	store.Delete(storagetypes.GetMigrationBucketKey(bucketInfo.Id))

	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventRejectMigrateBucket{
		Operator:   operator.String(),
		BucketName: bucketName,
		BucketId:   bucketInfo.Id,
		Status:     bucketInfo.BucketStatus,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) GetMigrationBucketInfo(ctx sdk.Context, bucketID sdkmath.Uint) (*storagetypes.MigrationBucketInfo, bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(storagetypes.GetMigrationBucketKey(bucketID))
	if bz == nil {
		return nil, false
	}

	var migrationBucketInfo storagetypes.MigrationBucketInfo
	k.cdc.MustUnmarshal(bz, &migrationBucketInfo)
	return &migrationBucketInfo, true
}

func (k Keeper) DeleteMigrationBucketInfo(ctx sdk.Context, bucketID sdkmath.Uint) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(storagetypes.GetMigrationBucketKey(bucketID))
}

func (k Keeper) setQuotaUpdateTime(ctx sdk.Context, bucketID storagetypes.Uint, timestamp uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, timestamp)
	store.Set(storagetypes.GetQuotaKey(bucketID), bz)
}

func (k Keeper) getQuotaUpdateTime(ctx sdk.Context, bucketID storagetypes.Uint) (uint64, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(storagetypes.GetQuotaKey(bucketID))
	if bz != nil {
		return binary.BigEndian.Uint64(bz), true
	}
	bucketInfo, found := k.GetBucketInfoById(ctx, bucketID)
	if !found {
		return 0, false
	}
	return uint64(bucketInfo.CreateAt), true
}

func (k Keeper) MustGetInternalBucketInfo(ctx sdk.Context, bucketID sdkmath.Uint) *storagetypes.InternalBucketInfo {
	internalBucketInfo, found := k.GetInternalBucketInfo(ctx, bucketID)
	if !found {
		panic("Internal bucket Info not found")
	}
	return internalBucketInfo
}

func (k Keeper) GetInternalBucketInfo(ctx sdk.Context, bucketID sdkmath.Uint) (*storagetypes.InternalBucketInfo, bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(storagetypes.GetInternalBucketInfoKey(bucketID))
	if bz == nil {
		return nil, false
	}

	var internalBucketInfo storagetypes.InternalBucketInfo
	k.cdc.MustUnmarshal(bz, &internalBucketInfo)
	return &internalBucketInfo, true
}

func (k Keeper) SetInternalBucketInfo(ctx sdk.Context, bucketID sdkmath.Uint, internalBucketInfo *storagetypes.InternalBucketInfo) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(internalBucketInfo)
	store.Set(storagetypes.GetInternalBucketInfoKey(bucketID), bz)
}

func (k Keeper) fromSpMaintenanceAcct(sp *sptypes.StorageProvider, operatorAddr sdk.AccAddress) bool {
	return sp.Status == sptypes.STATUS_IN_MAINTENANCE && operatorAddr.Equals(sdk.MustAccAddressFromHex(sp.MaintenanceAddress))
}

func (k Keeper) hasGroup(ctx sdk.Context, groupID sdkmath.Uint) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(storagetypes.GetGroupByIDKey(groupID))
}

func (k Keeper) GetSourceTypeByChainId(ctx sdk.Context, chainId sdk.ChainID) (storagetypes.SourceType, error) { //nolint
	if chainId == 0 {
		return 0, storagetypes.ErrChainNotSupported
	}

	switch chainId {
	case k.crossChainKeeper.GetDestBscChainID():
		return storagetypes.SOURCE_TYPE_BSC_CROSS_CHAIN, nil
	case k.crossChainKeeper.GetDestOpChainID():
		return storagetypes.SOURCE_TYPE_OP_CROSS_CHAIN, nil
	case k.crossChainKeeper.GetDestPolygonChainID():
		return storagetypes.SOURCE_TYPE_POLYGON_CROSS_CHAIN, nil
	case k.crossChainKeeper.GetDestScrollChainID():
		return storagetypes.SOURCE_TYPE_SCROLL_CROSS_CHAIN, nil
	case k.crossChainKeeper.GetDestLineaChainID():
		return storagetypes.SOURCE_TYPE_LINEA_CROSS_CHAIN, nil
	case k.crossChainKeeper.GetDestMantleChainID():
		return storagetypes.SOURCE_TYPE_MANTLE_CROSS_CHAIN, nil
	case k.crossChainKeeper.GetDestArbitrumChainID():
		return storagetypes.SOURCE_TYPE_ARBITRUM_CROSS_CHAIN, nil
	case k.crossChainKeeper.GetDestOptimismChainID():
		return storagetypes.SOURCE_TYPE_OPTIMISM_CROSS_CHAIN, nil
	default:
		return 0, storagetypes.ErrChainNotSupported
	}
}

func (k Keeper) SetTag(ctx sdk.Context, operator sdk.AccAddress, grn types.GRN, tags *storagetypes.ResourceTags) error {
	store := ctx.KVStore(k.storeKey)

	var id sdkmath.Uint
	switch grn.ResourceType() {
	case resource.RESOURCE_TYPE_BUCKET:
		bucketName, grnErr := grn.GetBucketName()
		if grnErr != nil {
			return grnErr
		}
		bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
		if !found {
			return storagetypes.ErrNoSuchBucket.Wrapf("bucketName: %s", bucketName)
		}
		// check permission
		effect := k.VerifyBucketPermission(ctx, bucketInfo, operator, permtypes.ACTION_UPDATE_BUCKET_INFO, nil)
		if effect != permtypes.EFFECT_ALLOW {
			return storagetypes.ErrAccessDenied.Wrapf("The operator(%s) has no updateBucketInfo permission of the bucket(%s)",
				operator.String(), bucketName)
		}

		bucketInfo.Tags = tags
		bz := k.cdc.MustMarshal(bucketInfo)
		store.Set(storagetypes.GetBucketByIDKey(bucketInfo.Id), bz)
		id = bucketInfo.Id
	case resource.RESOURCE_TYPE_OBJECT:
		bucketName, objectName, grnErr := grn.GetBucketAndObjectName()
		if grnErr != nil {
			return grnErr
		}
		objectInfo, found := k.GetObjectInfo(ctx, bucketName, objectName)
		if !found {
			return storagetypes.ErrNoSuchObject.Wrapf("BucketName: %s, objectName: %s", bucketName, objectName)
		}
		bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
		if !found {
			return storagetypes.ErrNoSuchBucket.Wrapf("bucketName: %s", bucketName)
		}
		effect := k.VerifyObjectPermission(ctx, bucketInfo, objectInfo, operator, permtypes.ACTION_UPDATE_OBJECT_INFO)
		if effect != permtypes.EFFECT_ALLOW {
			return storagetypes.ErrAccessDenied.Wrapf(
				"The operator(%s) has no updateObjectInfo permission of the bucket(%s), object(%s)",
				operator.String(), bucketName, objectName)
		}

		objectInfo.Tags = tags
		obz := k.cdc.MustMarshal(objectInfo)
		store.Set(storagetypes.GetObjectByIDKey(objectInfo.Id), obz)
		id = objectInfo.Id
	case resource.RESOURCE_TYPE_GROUP:
		groupOwner, groupName, grnErr := grn.GetGroupOwnerAndAccount()
		if grnErr != nil {
			return grnErr
		}
		groupInfo, found := k.GetGroupInfo(ctx, groupOwner, groupName)
		if !found {
			return storagetypes.ErrNoSuchBucket.Wrapf("groupOwner: %s, groupName: %s", groupOwner.String(), groupName)
		}
		effect := k.VerifyGroupPermission(ctx, groupInfo, operator, permtypes.ACTION_UPDATE_GROUP_INFO)
		if effect != permtypes.EFFECT_ALLOW {
			return storagetypes.ErrAccessDenied.Wrapf(
				"The operator(%s) has no updateGroupInfo permission of the group(%s), owner(%s)",
				operator.String(), groupInfo.GroupName, groupInfo.Owner)
		}

		groupInfo.Tags = tags
		gbz := k.cdc.MustMarshal(groupInfo)
		store.Set(storagetypes.GetGroupByIDKey(groupInfo.Id), gbz)
		id = groupInfo.Id
	default:
		return gnfderrors.ErrInvalidGRN.Wrap("Unknown resource type in mechain resource name")
	}

	// emit Event
	if err := ctx.EventManager().EmitTypedEvents(&storagetypes.EventSetTag{
		Resource: grn.String(),
		Tags:     tags,
		Id:       id,
	}); err != nil {
		return err
	}

	return nil
}

func (k Keeper) UpdateObjectContent(
	ctx sdk.Context, operator sdk.AccAddress, bucketName, objectName string, payloadSize uint64,
	opts storagetypes.UpdateObjectOptions,
) error {
	store := ctx.KVStore(k.storeKey)

	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}
	err := bucketInfo.CheckBucketStatus()
	if err != nil {
		return err
	}
	objectInfo, found := k.GetObjectInfo(ctx, bucketName, objectName)
	if !found {
		return storagetypes.ErrNoSuchObject
	}
	// check object status
	if objectInfo.ObjectStatus != storagetypes.OBJECT_STATUS_SEALED {
		return storagetypes.ErrUpdateObjectNotAllowed.Wrapf("The object is not sealed yet")
	}
	if objectInfo.IsUpdating {
		return storagetypes.ErrObjectIsUpdating.Wrapf("The object is already being updated")
	}
	// check permission
	var updater sdk.AccAddress
	if opts.Delegated {
		updater = opts.Updater
	} else {
		updater = operator
	}
	effect := k.VerifyObjectPermission(ctx, bucketInfo, objectInfo, updater, permtypes.ACTION_UPDATE_OBJECT_CONTENT)
	if effect != permtypes.EFFECT_ALLOW {
		return storagetypes.ErrAccessDenied.Wrapf(
			"The updater(%s) has no updateObjectContent permission of the bucket(%s), object(%s)",
			updater.String(), bucketName, objectName)
	}

	// check payload size
	if payloadSize > k.MaxPayloadSize(ctx) {
		return storagetypes.ErrTooLargeObject
	}

	// primary sp
	sp := k.MustGetPrimarySPForBucket(ctx, bucketInfo)
	// a sp is not in service, neither in maintenance
	if sp.Status != sptypes.STATUS_IN_SERVICE && !k.fromSpMaintenanceAcct(sp, operator) {
		return errors.Wrap(storagetypes.ErrNoSuchStorageProvider, "the storage provider is not in service")
	}
	if opts.Delegated {
		if bucketInfo.SpAsDelegatedAgentDisabled {
			return storagetypes.ErrAccessDenied.Wrap("the SP is not allowed to create object for delegator, disabled by the bucket owner previously")
		}
		if operator.String() != sp.OperatorAddress {
			return storagetypes.ErrAccessDenied.Wrap("only the primary SP is allowed to create object for delegator")
		}
	}
	nextVersion := objectInfo.Version + 1

	if payloadSize == 0 {
		internalBucketInfo := k.MustGetInternalBucketInfo(ctx, bucketInfo.Id)
		err := k.UnChargeObjectStoreFee(ctx, bucketInfo, k.MustGetInternalBucketInfo(ctx, bucketInfo.Id), objectInfo)
		if err != nil {
			return err
		}
		k.SetInternalBucketInfo(ctx, bucketInfo.Id, internalBucketInfo)
		err = k.DeleteObjectFromVirtualGroup(ctx, bucketInfo, objectInfo)
		if err != nil {
			return err
		}
		objectInfo.UpdatedAt = ctx.BlockTime().Unix()
		objectInfo.Version = nextVersion
		objectInfo.PayloadSize = 0
		objectInfo.Checksums = opts.Checksums
		objectInfo.UpdatedBy = updater.String()
		objectInfo.ContentType = opts.ContentType

		_, err = k.SealEmptyObjectOnVirtualGroup(ctx, bucketInfo, objectInfo)
		if err != nil {
			return err
		}
	} else {
		objectInfo.IsUpdating = true
		shadowObjectInfo := &storagetypes.ShadowObjectInfo{
			Operator:    updater.String(),
			Id:          objectInfo.Id,
			PayloadSize: payloadSize,
			Checksums:   opts.Checksums,
			ContentType: opts.ContentType,
			UpdatedAt:   ctx.BlockTime().Unix(),
			Version:     nextVersion,
		}
		store.Set(storagetypes.GetShadowObjectKey(bucketName, objectName), k.cdc.MustMarshal(shadowObjectInfo))
		err = k.LockShadowObjectStoreFee(ctx, bucketInfo, shadowObjectInfo, objectName)
		if err != nil {
			return err
		}
		k.IncreaseLockedObjectCount(ctx, bucketInfo.Id)
	}

	obz := k.cdc.MustMarshal(objectInfo)
	store.Set(storagetypes.GetObjectKey(bucketName, objectName), k.objectSeq.EncodeSequence(objectInfo.Id))
	store.Set(storagetypes.GetObjectByIDKey(objectInfo.Id), obz)

	if err = ctx.EventManager().EmitTypedEvents(&storagetypes.EventUpdateObjectContent{
		Operator:    updater.String(),
		ObjectId:    objectInfo.Id,
		PayloadSize: payloadSize,
		Checksums:   opts.Checksums,
		Version:     nextVersion,
	}); err != nil {
		return err
	}
	return nil
}

func (k Keeper) UnlockShadowObjectFeeAndDeleteShadowObjectInfo(ctx sdk.Context, bucketInfo *storagetypes.BucketInfo, shadowObjectInfo *storagetypes.ShadowObjectInfo, objectName string) (err error) {
	err = k.UnlockShadowObjectStoreFee(ctx, bucketInfo, shadowObjectInfo)
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	store.Delete(storagetypes.GetShadowObjectKey(bucketInfo.BucketName, objectName))
	return
}

func (k Keeper) CancelUpdateObjectContent(
	ctx sdk.Context, operator sdk.AccAddress,
	bucketName, objectName string,
) error {
	store := ctx.KVStore(k.storeKey)
	bucketInfo, found := k.GetBucketInfo(ctx, bucketName)
	if !found {
		return storagetypes.ErrNoSuchBucket
	}
	objectInfo, found := k.GetObjectInfo(ctx, bucketName, objectName)
	if !found {
		return storagetypes.ErrNoSuchObject
	}
	if !objectInfo.IsUpdating {
		return storagetypes.ErrObjectIsNotUpdating
	}
	shadowObjectInfo := k.MustGetShadowObjectInfo(ctx, bucketName, objectName)
	err := k.UnlockShadowObjectFeeAndDeleteShadowObjectInfo(ctx, bucketInfo, shadowObjectInfo, objectName)
	if err != nil {
		return err
	}
	updater := sdk.MustAccAddressFromHex(shadowObjectInfo.Operator)
	owner := sdk.MustAccAddressFromHex(objectInfo.Owner)
	if !operator.Equals(owner) && !operator.Equals(updater) {
		return errors.Wrapf(storagetypes.ErrAccessDenied, "Only allowed owner/updater to do cancel update object")
	}

	objectInfo.IsUpdating = false
	obz := k.cdc.MustMarshal(objectInfo)
	store.Set(storagetypes.GetObjectByIDKey(objectInfo.Id), obz)
	k.DecreaseLockedObjectCount(ctx, bucketInfo.Id)

	return ctx.EventManager().EmitTypedEvents(&storagetypes.EventCancelUpdateObjectContent{
		Operator:   operator.String(),
		BucketName: bucketInfo.BucketName,
		ObjectName: objectInfo.ObjectName,
		ObjectId:   objectInfo.Id,
	})
}

func (k Keeper) GetLockedObjectCount(ctx sdk.Context, bucketID sdkmath.Uint) uint64 {
	store := ctx.KVStore(k.storeKey)

	key := storagetypes.GetLockedObjectCountKey(bucketID)
	bz := store.Get(key)
	current := uint64(0)
	if bz != nil {
		current = binary.BigEndian.Uint64(bz)
	}
	return current
}

func (k Keeper) IncreaseLockedObjectCount(ctx sdk.Context, bucketID sdkmath.Uint) {
	store := ctx.KVStore(k.storeKey)

	key := storagetypes.GetLockedObjectCountKey(bucketID)
	bz := store.Get(key)
	before := uint64(0)
	if bz != nil {
		before = binary.BigEndian.Uint64(bz)
	}
	after := before + 1

	bz = make([]byte, 8)
	binary.BigEndian.PutUint64(bz, after)
	store.Set(key, bz)
}

func (k Keeper) DecreaseLockedObjectCount(ctx sdk.Context, bucketID sdkmath.Uint) {
	store := ctx.KVStore(k.storeKey)

	key := storagetypes.GetLockedObjectCountKey(bucketID)
	bz := store.Get(key)
	before := uint64(0)
	if bz != nil {
		before = binary.BigEndian.Uint64(bz)
	}

	after := before
	if before > 0 {
		after = before - 1
	}
	// this feature is not introduced from the genesis, which means that some buckets do not have such
	// indicators earlier, they could have created objects even the indicator is zero. For such buckets,
	// after they delete or seal all the created objects, the indicator will become correct finally.

	bz = make([]byte, 8)
	binary.BigEndian.PutUint64(bz, after)
	store.Set(key, bz)
}
