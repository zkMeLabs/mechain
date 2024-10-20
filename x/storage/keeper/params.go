package keeper

import (
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/evmos/evmos/v12/x/storage/types"
)

func (k Keeper) MaxBucketsPerAccount(ctx sdk.Context) (res uint32) {
	params := k.GetParams(ctx)
	return params.MaxBucketsPerAccount
}

func (k Keeper) GetExpectSecondarySPNumForECObject(ctx sdk.Context, createTime int64) (res uint32) {
	versionParams, err := k.GetVersionedParamsWithTS(ctx, createTime)
	if err != nil {
		panic(fmt.Sprintf("get expect secondary sp num error, msg: %s", err))
	}
	return versionParams.RedundantParityChunkNum + versionParams.RedundantDataChunkNum
}

func (k Keeper) MaxPayloadSize(ctx sdk.Context) (res uint64) {
	params := k.GetParams(ctx)
	return params.MaxPayloadSize
}

func (k Keeper) MirrorBucketRelayerFee(ctx sdk.Context, destChainID sdk.ChainID) *big.Int {
	params := k.GetParams(ctx)

	var relayerFeeParam string
	switch {
	case k.crossChainKeeper.GetDestBscChainID() == destChainID:
		relayerFeeParam = params.BscMirrorBucketRelayerFee
	case k.crossChainKeeper.GetDestOpChainID() == destChainID:
		relayerFeeParam = params.OpMirrorBucketRelayerFee
	case k.crossChainKeeper.GetDestPolygonChainID() == destChainID:
		relayerFeeParam = params.PolygonMirrorBucketRelayerFee
	case k.crossChainKeeper.GetDestScrollChainID() == destChainID:
		relayerFeeParam = params.ScrollMirrorBucketRelayerFee
	case k.crossChainKeeper.GetDestLineaChainID() == destChainID:
		relayerFeeParam = params.LineaMirrorBucketRelayerFee
	case k.crossChainKeeper.GetDestMantleChainID() == destChainID:
		relayerFeeParam = params.MantleMirrorBucketRelayerFee
	case k.crossChainKeeper.GetDestArbitrumChainID() == destChainID:
		relayerFeeParam = params.ArbitrumMirrorBucketRelayerFee
	case k.crossChainKeeper.GetDestOptimismChainID() == destChainID:
		relayerFeeParam = params.OptimismMirrorBucketRelayerFee
	default:
		panic(fmt.Sprintf("chain id(%d) is not supported", destChainID))
	}

	relayerFee, valid := big.NewInt(0).SetString(relayerFeeParam, 10)
	if !valid {
		panic(fmt.Sprintf("invalid relayer fee: %s", relayerFeeParam))
	}

	return relayerFee
}

func (k Keeper) MirrorBucketAckRelayerFee(ctx sdk.Context, destChainID sdk.ChainID) *big.Int {
	params := k.GetParams(ctx)

	var relayerFeeParam string
	switch {
	case k.crossChainKeeper.GetDestBscChainID() == destChainID:
		relayerFeeParam = params.BscMirrorBucketAckRelayerFee
	case k.crossChainKeeper.GetDestOpChainID() == destChainID:
		relayerFeeParam = params.OpMirrorBucketAckRelayerFee
	case k.crossChainKeeper.GetDestPolygonChainID() == destChainID:
		relayerFeeParam = params.PolygonMirrorBucketAckRelayerFee
	case k.crossChainKeeper.GetDestScrollChainID() == destChainID:
		relayerFeeParam = params.ScrollMirrorBucketAckRelayerFee
	case k.crossChainKeeper.GetDestLineaChainID() == destChainID:
		relayerFeeParam = params.LineaMirrorBucketAckRelayerFee
	case k.crossChainKeeper.GetDestMantleChainID() == destChainID:
		relayerFeeParam = params.MantleMirrorBucketAckRelayerFee
	case k.crossChainKeeper.GetDestArbitrumChainID() == destChainID:
		relayerFeeParam = params.ArbitrumMirrorBucketAckRelayerFee
	case k.crossChainKeeper.GetDestOptimismChainID() == destChainID:
		relayerFeeParam = params.OptimismMirrorBucketAckRelayerFee
	default:
		panic(fmt.Sprintf("chain id(%d) is not supported", destChainID))
	}

	relayerFee, valid := big.NewInt(0).SetString(relayerFeeParam, 10)
	if !valid {
		panic(fmt.Sprintf("invalid relayer fee: %s", relayerFeeParam))
	}

	return relayerFee
}

func (k Keeper) MirrorObjectRelayerFee(ctx sdk.Context, destChainID sdk.ChainID) *big.Int {
	params := k.GetParams(ctx)

	var relayerFeeParam string
	switch {
	case k.crossChainKeeper.GetDestBscChainID() == destChainID:
		relayerFeeParam = params.BscMirrorObjectRelayerFee
	case k.crossChainKeeper.GetDestOpChainID() == destChainID:
		relayerFeeParam = params.OpMirrorObjectRelayerFee
	case k.crossChainKeeper.GetDestPolygonChainID() == destChainID:
		relayerFeeParam = params.PolygonMirrorObjectRelayerFee
	case k.crossChainKeeper.GetDestScrollChainID() == destChainID:
		relayerFeeParam = params.ScrollMirrorObjectRelayerFee
	case k.crossChainKeeper.GetDestLineaChainID() == destChainID:
		relayerFeeParam = params.LineaMirrorObjectRelayerFee
	case k.crossChainKeeper.GetDestMantleChainID() == destChainID:
		relayerFeeParam = params.MantleMirrorObjectRelayerFee
	case k.crossChainKeeper.GetDestArbitrumChainID() == destChainID:
		relayerFeeParam = params.ArbitrumMirrorObjectRelayerFee
	case k.crossChainKeeper.GetDestOptimismChainID() == destChainID:
		relayerFeeParam = params.OptimismMirrorObjectRelayerFee
	default:
		panic(fmt.Sprintf("chain id(%d) is not supported", destChainID))
	}

	relayerFee, valid := big.NewInt(0).SetString(relayerFeeParam, 10)
	if !valid {
		panic(fmt.Sprintf("invalid relayer fee: %s", relayerFeeParam))
	}

	return relayerFee
}

func (k Keeper) MirrorObjectAckRelayerFee(ctx sdk.Context, destChainID sdk.ChainID) *big.Int {
	params := k.GetParams(ctx)

	var relayerFeeParam string
	switch {
	case k.crossChainKeeper.GetDestBscChainID() == destChainID:
		relayerFeeParam = params.BscMirrorObjectAckRelayerFee
	case k.crossChainKeeper.GetDestOpChainID() == destChainID:
		relayerFeeParam = params.OpMirrorObjectAckRelayerFee
	case k.crossChainKeeper.GetDestPolygonChainID() == destChainID:
		relayerFeeParam = params.PolygonMirrorObjectAckRelayerFee
	case k.crossChainKeeper.GetDestScrollChainID() == destChainID:
		relayerFeeParam = params.ScrollMirrorObjectAckRelayerFee
	case k.crossChainKeeper.GetDestLineaChainID() == destChainID:
		relayerFeeParam = params.LineaMirrorObjectAckRelayerFee
	case k.crossChainKeeper.GetDestMantleChainID() == destChainID:
		relayerFeeParam = params.MantleMirrorObjectAckRelayerFee
	case k.crossChainKeeper.GetDestArbitrumChainID() == destChainID:
		relayerFeeParam = params.ArbitrumMirrorObjectAckRelayerFee
	case k.crossChainKeeper.GetDestOptimismChainID() == destChainID:
		relayerFeeParam = params.OptimismMirrorObjectAckRelayerFee
	default:
		panic(fmt.Sprintf("chain id(%d) is not supported", destChainID))

	}
	relayerFee, valid := big.NewInt(0).SetString(relayerFeeParam, 10)
	if !valid {
		panic(fmt.Sprintf("invalid relayer fee: %s", relayerFeeParam))
	}

	return relayerFee
}

func (k Keeper) MirrorGroupRelayerFee(ctx sdk.Context, destChainID sdk.ChainID) *big.Int {
	params := k.GetParams(ctx)
	var relayerFeeParam string
	switch {
	case k.crossChainKeeper.GetDestBscChainID() == destChainID:
		relayerFeeParam = params.BscMirrorGroupRelayerFee
	case k.crossChainKeeper.GetDestOpChainID() == destChainID:
		relayerFeeParam = params.OpMirrorGroupRelayerFee
	case k.crossChainKeeper.GetDestPolygonChainID() == destChainID:
		relayerFeeParam = params.PolygonMirrorGroupRelayerFee
	case k.crossChainKeeper.GetDestScrollChainID() == destChainID:
		relayerFeeParam = params.ScrollMirrorGroupRelayerFee
	case k.crossChainKeeper.GetDestLineaChainID() == destChainID:
		relayerFeeParam = params.LineaMirrorGroupRelayerFee
	case k.crossChainKeeper.GetDestMantleChainID() == destChainID:
		relayerFeeParam = params.MantleMirrorGroupRelayerFee
	case k.crossChainKeeper.GetDestArbitrumChainID() == destChainID:
		relayerFeeParam = params.ArbitrumMirrorGroupRelayerFee
	case k.crossChainKeeper.GetDestOptimismChainID() == destChainID:
		relayerFeeParam = params.OptimismMirrorGroupRelayerFee
	default:
		panic(fmt.Sprintf("chain id(%d) is not supported", destChainID))

	}

	relayerFee, valid := big.NewInt(0).SetString(relayerFeeParam, 10)
	if !valid {
		panic(fmt.Sprintf("invalid relayer fee: %s", relayerFeeParam))
	}

	return relayerFee
}

func (k Keeper) MirrorGroupAckRelayerFee(ctx sdk.Context, destChainID sdk.ChainID) *big.Int {
	params := k.GetParams(ctx)

	var relayerFeeParam string
	switch {
	case k.crossChainKeeper.GetDestBscChainID() == destChainID:
		relayerFeeParam = params.BscMirrorGroupAckRelayerFee
	case k.crossChainKeeper.GetDestOpChainID() == destChainID:
		relayerFeeParam = params.OpMirrorGroupAckRelayerFee
	case k.crossChainKeeper.GetDestPolygonChainID() == destChainID:
		relayerFeeParam = params.PolygonMirrorGroupAckRelayerFee
	case k.crossChainKeeper.GetDestScrollChainID() == destChainID:
		relayerFeeParam = params.ScrollMirrorGroupAckRelayerFee
	case k.crossChainKeeper.GetDestLineaChainID() == destChainID:
		relayerFeeParam = params.LineaMirrorGroupAckRelayerFee
	case k.crossChainKeeper.GetDestMantleChainID() == destChainID:
		relayerFeeParam = params.MantleMirrorGroupAckRelayerFee
	case k.crossChainKeeper.GetDestArbitrumChainID() == destChainID:
		relayerFeeParam = params.ArbitrumMirrorGroupAckRelayerFee
	case k.crossChainKeeper.GetDestOptimismChainID() == destChainID:
		relayerFeeParam = params.OptimismMirrorGroupAckRelayerFee
	default:
		panic(fmt.Sprintf("chain id(%d) is not supported", destChainID))
	}

	relayerFee, valid := big.NewInt(0).SetString(relayerFeeParam, 10)
	if !valid {
		panic(fmt.Sprintf("invalid relayer fee: %s", relayerFeeParam))
	}

	return relayerFee
}

func (k Keeper) DiscontinueCountingWindow(ctx sdk.Context) (res uint64) {
	params := k.GetParams(ctx)
	return params.DiscontinueCountingWindow
}

func (k Keeper) DiscontinueObjectMax(ctx sdk.Context) (res uint64) {
	params := k.GetParams(ctx)
	return params.DiscontinueObjectMax
}

func (k Keeper) DiscontinueBucketMax(ctx sdk.Context) (res uint64) {
	params := k.GetParams(ctx)
	return params.DiscontinueBucketMax
}

func (k Keeper) DiscontinueConfirmPeriod(ctx sdk.Context) (res int64) {
	params := k.GetParams(ctx)
	return params.DiscontinueConfirmPeriod
}

func (k Keeper) DiscontinueDeletionMax(ctx sdk.Context) (res uint64) {
	params := k.GetParams(ctx)
	return params.DiscontinueDeletionMax
}

func (k Keeper) MaxSegmentSize(ctx sdk.Context, timestamp int64) (res uint64, err error) {
	params, err := k.GetVersionedParamsWithTS(ctx, timestamp)
	if err != nil {
		return 0, err
	}
	return params.MaxSegmentSize, err
}

func (k Keeper) RedundantDataChunkNum(ctx sdk.Context) (res uint32) {
	p := k.GetParams(ctx)
	params := p.GetVersionedParams()
	return params.RedundantDataChunkNum
}

func (k Keeper) RedundantParityChunkNum(ctx sdk.Context) (res uint32) {
	p := k.GetParams(ctx)
	params := p.GetVersionedParams()
	return params.RedundantParityChunkNum
}

func (k Keeper) MinChargeSize(ctx sdk.Context) (res uint64) {
	p := k.GetParams(ctx)
	params := p.GetVersionedParams()
	return params.MinChargeSize
}

func (k Keeper) StalePolicyCleanupMax(ctx sdk.Context) (res uint64) {
	params := k.GetParams(ctx)
	return params.StalePolicyCleanupMax
}

// GetParams returns the current storage module parameters.
func (k Keeper) GetParams(ctx sdk.Context) (p types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return p
	}

	k.cdc.MustUnmarshal(bz, &p)
	return p
}

// SetParams sets the params of storage module
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, bz)

	// store another kv with timestamp
	err := k.SetVersionedParamsWithTS(ctx, params.VersionedParams)
	if err != nil {
		return err
	}

	return nil
}

// SetVersionedParamsWithTS set a specific params in the store from its index
func (k Keeper) SetVersionedParamsWithTS(ctx sdk.Context, verParams types.VersionedParams) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VersionedParamsKeyPrefix)
	key := types.GetParamsKeyWithTimestamp(ctx.BlockTime().Unix())

	b := k.cdc.MustMarshal(&verParams)
	store.Set(key, b)

	return nil
}

// GetVersionedParamsWithTS find the latest params before and equal than the specific timestamp
func (k Keeper) GetVersionedParamsWithTS(ctx sdk.Context, ts int64) (verParams types.VersionedParams, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VersionedParamsKeyPrefix)

	// params are updated in the endblock, so we do not need to make the ts to be included
	// for example, if the params is updated in 100 timestamp: the txs that are executed in 100 timestamp
	// will use the old parameter, after 100 timestamp, when we passing 100 to query, we should still get
	// the old parameter.
	startKey := types.GetParamsKeyWithTimestamp(ts)
	iterator := store.ReverseIterator(nil, startKey)
	defer iterator.Close()
	if !iterator.Valid() {
		return verParams, fmt.Errorf("no versioned params found, ts:%d", uint64(ts))
	}

	k.cdc.MustUnmarshal(iterator.Value(), &verParams)

	return verParams, nil
}

func (k Keeper) MaxLocalVirtualGroupNumPerBucket(ctx sdk.Context) (res uint32) {
	params := k.GetParams(ctx)
	return params.MaxLocalVirtualGroupNumPerBucket
}
