package types

import (
	"fmt"
	"math"
	"math/big"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

// storage params default values
const (
	DefaultMaxSegmentSize            uint64 = 16 * 1024 * 1024 // 16M
	DefaultRedundantDataChunkNum     uint32 = 4
	DefaultRedundantParityChunkNum   uint32 = 2
	DefaultMaxPayloadSize            uint64 = 64 * 1024 * 1024 * 1024
	DefaultMaxBucketsPerAccount      uint32 = 100
	DefaultMinChargeSize             uint64 = 1 * 1024 * 1024 // 1M
	DefaultDiscontinueCountingWindow uint64 = 10000
	DefaultDiscontinueObjectMax      uint64 = math.MaxUint64
	DefaultDiscontinueBucketMax      uint64 = math.MaxUint64
	DefaultDiscontinueConfirmPeriod  int64  = 604800 // 7 days (in second)
	DefaultDiscontinueDeletionMax    uint64 = 100
	DefaultStalePolicyCleanupMax     uint64 = 200
	DefaultMinUpdateQuotaInterval    uint64 = 2592000 // 30 days (in second)

	// TODO
	DefaultMaxLocalVirtualGroupNumPerBucket  uint32 = 10
	DefaultBscMirrorBucketRelayerFee                = "1300000000000000" // 0.0013
	DefaultBscMirrorBucketAckRelayerFee             = "250000000000000"  // 0.00025
	DefaultBscMirrorObjectRelayerFee                = "1300000000000000" // 0.0013
	DefaultBscMirrorObjectAckRelayerFee             = "250000000000000"  // 0.00025
	DefaultBscMirrorGroupRelayerFee                 = "1300000000000000" // 0.0013
	DefaultBscMirrorGroupAckRelayerFee              = "250000000000000"  // 0.00025
	DefaultOpMirrorBucketRelayerFee                 = "130000000000000"  // 0.00013
	DefaultOpMirrorBucketAckRelayerFee              = "25000000000000"   // 0.000025
	DefaultOpMirrorObjectRelayerFee                 = "130000000000000"  // 0.00013
	DefaultOpMirrorObjectAckRelayerFee              = "25000000000000"   // 0.000025
	DefaultOpMirrorGroupRelayerFee                  = "130000000000000"  // 0.00013
	DefaultOpMirrorGroupAckRelayerFee               = "25000000000000"   // 0.000025
	DefaultPolygonMirrorBucketRelayerFee            = "130000000000000"  // 0.00013
	DefaultPolygonMirrorBucketAckRelayerFee         = "25000000000000"   // 0.000025
	DefaultPolygonMirrorObjectRelayerFee            = "130000000000000"  // 0.00013
	DefaultPolygonMirrorObjectAckRelayerFee         = "25000000000000"   // 0.000025
	DefaultPolygonMirrorGroupRelayerFee             = "130000000000000"  // 0.00013
	DefaultPolygonMirrorGroupAckRelayerFee          = "25000000000000"   // 0.000025
	DefaultScrollMirrorBucketRelayerFee             = "130000000000000"  // 0.00013
	DefaultScrollMirrorBucketAckRelayerFee          = "25000000000000"   // 0.000025
	DefaultScrollMirrorObjectRelayerFee             = "130000000000000"  // 0.00013
	DefaultScrollMirrorObjectAckRelayerFee          = "25000000000000"   // 0.000025
	DefaultScrollMirrorGroupRelayerFee              = "130000000000000"  // 0.00013
	DefaultScrollMirrorGroupAckRelayerFee           = "25000000000000"   // 0.000025
	DefaultLineaMirrorBucketRelayerFee              = "130000000000000"  // 0.00013
	DefaultLineaMirrorBucketAckRelayerFee           = "25000000000000"   // 0.000025
	DefaultLineaMirrorObjectRelayerFee              = "130000000000000"  // 0.00013
	DefaultLineaMirrorObjectAckRelayerFee           = "25000000000000"   // 0.000025
	DefaultLineaMirrorGroupRelayerFee               = "130000000000000"  // 0.00013
	DefaultLineaMirrorGroupAckRelayerFee            = "25000000000000"   // 0.000025
	DefaultMantleMirrorBucketRelayerFee             = "130000000000000"  // 0.00013
	DefaultMantleMirrorBucketAckRelayerFee          = "25000000000000"   // 0.000025
	DefaultMantleMirrorObjectRelayerFee             = "130000000000000"  // 0.00013
	DefaultMantleMirrorObjectAckRelayerFee          = "25000000000000"   // 0.000025
	DefaultMantleMirrorGroupRelayerFee              = "130000000000000"  // 0.00013
	DefaultMantleMirrorGroupAckRelayerFee           = "25000000000000"   // 0.000025
	DefaultArbitrumMirrorBucketRelayerFee           = "130000000000000"  // 0.00013
	DefaultArbitrumMirrorBucketAckRelayerFee        = "25000000000000"   // 0.000025
	DefaultArbitrumMirrorObjectRelayerFee           = "130000000000000"  // 0.00013
	DefaultArbitrumMirrorObjectAckRelayerFee        = "25000000000000"   // 0.000025
	DefaultArbitrumMirrorGroupRelayerFee            = "130000000000000"  // 0.00013
	DefaultArbitrumMirrorGroupAckRelayerFee         = "25000000000000"   // 0.000025
	DefaultOptimismMirrorBucketRelayerFee           = "130000000000000"  // 0.00013
	DefaultOptimismMirrorBucketAckRelayerFee        = "25000000000000"   // 0.000025
	DefaultOptimismMirrorObjectRelayerFee           = "130000000000000"  // 0.00013
	DefaultOptimismMirrorObjectAckRelayerFee        = "25000000000000"   // 0.000025
	DefaultOptimismMirrorGroupRelayerFee            = "130000000000000"  // 0.00013
	DefaultOptimismMirrorGroupAckRelayerFee         = "25000000000000"   // 0.000025
)

var (
	KeyMaxSegmentSize                    = []byte("MaxSegmentSize")
	KeyRedundantDataChunkNum             = []byte("RedundantDataChunkNum")
	KeyRedundantParityChunkNum           = []byte("RedundantParityChunkNum")
	KeyMaxPayloadSize                    = []byte("MaxPayloadSize")
	KeyMinChargeSize                     = []byte("MinChargeSize")
	KeyMaxBucketsPerAccount              = []byte("MaxBucketsPerAccount")
	KeyDiscontinueCountingWindow         = []byte("DiscontinueCountingWindow")
	KeyDiscontinueObjectMax              = []byte("DiscontinueObjectMax")
	KeyDiscontinueBucketMax              = []byte("DiscontinueBucketMax")
	KeyDiscontinueConfirmPeriod          = []byte("DiscontinueConfirmPeriod")
	KeyDiscontinueDeletionMax            = []byte("DiscontinueDeletionMax")
	KeyStalePolicyCleanupMax             = []byte("StalePolicyCleanupMax")
	KeyMinUpdateQuotaInterval            = []byte("MinUpdateQuotaInterval")
	KeyBscMirrorBucketRelayerFee         = []byte("BscMirrorBucketRelayerFee")
	KeyBscMirrorBucketAckRelayerFee      = []byte("BscMirrorBucketAckRelayerFee")
	KeyBscMirrorObjectRelayerFee         = []byte("BscMirrorObjectRelayerFee")
	KeyBscMirrorObjectAckRelayerFee      = []byte("BscMirrorObjectAckRelayerFee")
	KeyBscMirrorGroupRelayerFee          = []byte("BscMirrorGroupRelayerFee")
	KeyBscMirrorGroupAckRelayerFee       = []byte("BscMirrorGroupAckRelayerFee")
	KeyOpMirrorBucketRelayerFee          = []byte("OpMirrorBucketRelayerFee")
	KeyOpMirrorBucketAckRelayerFee       = []byte("OpMirrorBucketAckRelayerFee")
	KeyOpMirrorObjectRelayerFee          = []byte("OpMirrorObjectRelayerFee")
	KeyOpMirrorObjectAckRelayerFee       = []byte("OpMirrorObjectAckRelayerFee")
	KeyOpMirrorGroupRelayerFee           = []byte("OpMirrorGroupRelayerFee")
	KeyOpMirrorGroupAckRelayerFee        = []byte("OpMirrorGroupAckRelayerFee")
	KeyPolygonMirrorBucketRelayerFee     = []byte("PolygonMirrorBucketRelayerFee")
	KeyPolygonMirrorBucketAckRelayerFee  = []byte("PolygonMirrorBucketAckRelayerFee")
	KeyPolygonMirrorObjectRelayerFee     = []byte("PolygonMirrorObjectRelayerFee")
	KeyPolygonMirrorObjectAckRelayerFee  = []byte("PolygonMirrorObjectAckRelayerFee")
	KeyPolygonMirrorGroupRelayerFee      = []byte("PolygonMirrorGroupRelayerFee")
	KeyPolygonMirrorGroupAckRelayerFee   = []byte("PolygonMirrorGroupAckRelayerFee")
	KeyScrollMirrorBucketRelayerFee      = []byte("ScrollMirrorBucketRelayerFee")
	KeyScrollMirrorBucketAckRelayerFee   = []byte("ScrollMirrorBucketAckRelayerFee")
	KeyScrollMirrorObjectRelayerFee      = []byte("ScrollMirrorObjectRelayerFee")
	KeyScrollMirrorObjectAckRelayerFee   = []byte("ScrollMirrorObjectAckRelayerFee")
	KeyScrollMirrorGroupRelayerFee       = []byte("ScrollMirrorGroupRelayerFee")
	KeyScrollMirrorGroupAckRelayerFee    = []byte("ScrollMirrorGroupAckRelayerFee")
	KeyLineaMirrorBucketRelayerFee       = []byte("LineaMirrorBucketRelayerFee")
	KeyLineaMirrorBucketAckRelayerFee    = []byte("LineaMirrorBucketAckRelayerFee")
	KeyLineaMirrorObjectRelayerFee       = []byte("LineaMirrorObjectRelayerFee")
	KeyLineaMirrorObjectAckRelayerFee    = []byte("LineaMirrorObjectAckRelayerFee")
	KeyLineaMirrorGroupRelayerFee        = []byte("LineaMirrorGroupRelayerFee")
	KeyLineaMirrorGroupAckRelayerFee     = []byte("LineaMirrorGroupAckRelayerFee")
	KeyMantleMirrorBucketRelayerFee      = []byte("MantleMirrorBucketRelayerFee")
	KeyMantleMirrorBucketAckRelayerFee   = []byte("MantleMirrorBucketAckRelayerFee")
	KeyMantleMirrorObjectRelayerFee      = []byte("MantleMirrorObjectRelayerFee")
	KeyMantleMirrorObjectAckRelayerFee   = []byte("MantleMirrorObjectAckRelayerFee")
	KeyMantleMirrorGroupRelayerFee       = []byte("MantleMirrorGroupRelayerFee")
	KeyMantleMirrorGroupAckRelayerFee    = []byte("MantleMirrorGroupAckRelayerFee")
	KeyArbitrumMirrorBucketRelayerFee    = []byte("ArbitrumMirrorBucketRelayerFee")
	KeyArbitrumMirrorBucketAckRelayerFee = []byte("ArbitrumMirrorBucketAckRelayerFee")
	KeyArbitrumMirrorObjectRelayerFee    = []byte("ArbitrumMirrorObjectRelayerFee")
	KeyArbitrumMirrorObjectAckRelayerFee = []byte("ArbitrumMirrorObjectAckRelayerFee")
	KeyArbitrumMirrorGroupRelayerFee     = []byte("ArbitrumMirrorGroupRelayerFee")
	KeyArbitrumMirrorGroupAckRelayerFee  = []byte("ArbitrumMirrorGroupAckRelayerFee")
	KeyOptimismMirrorBucketRelayerFee    = []byte("OptimismMirrorBucketRelayerFee")
	KeyOptimismMirrorBucketAckRelayerFee = []byte("OptimismMirrorBucketAckRelayerFee")
	KeyOptimismMirrorObjectRelayerFee    = []byte("OptimismMirrorObjectRelayerFee")
	KeyOptimismMirrorObjectAckRelayerFee = []byte("OptimismMirrorObjectAckRelayerFee")
	KeyOptimismMirrorGroupRelayerFee     = []byte("OptimismMirrorGroupRelayerFee")
	KeyOptimismMirrorGroupAckRelayerFee  = []byte("OptimismMirrorGroupAckRelayerFee")
	KeyMaxLocalVirtualGroupNumPerBucket  = []byte("MaxLocalVirtualGroupNumPerBucket")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	maxSegmentSize uint64, redundantDataChunkNum uint32,
	redundantParityChunkNum uint32, maxPayloadSize uint64, maxBucketsPerAccount uint32,
	minChargeSize uint64, bscMirrorBucketRelayerFee, bscMirrorBucketAckRelayerFee string,
	bscMirrorObjectRelayerFee, bscMirrorObjectAckRelayerFee string,
	bscMirrorGroupRelayerFee, bscMirrorGroupAckRelayerFee string,
	opMirrorBucketRelayerFee, opMirrorBucketAckRelayerFee string,
	opMirrorObjectRelayerFee, opMirrorObjectAckRelayerFee string,
	opMirrorGroupRelayerFee, opMirrorGroupAckRelayerFee string,
	polygonMirrorBucketRelayerFee, polygonMirrorBucketAckRelayerFee string,
	polygonMirrorObjectRelayerFee, polygonMirrorObjectAckRelayerFee string,
	polygonMirrorGroupRelayerFee, polygonMirrorGroupAckRelayerFee string,
	scrollMirrorBucketRelayerFee, scrollMirrorBucketAckRelayerFee string,
	scrollMirrorObjectRelayerFee, scrollMirrorObjectAckRelayerFee string,
	scrollMirrorGroupRelayerFee, scrollMirrorGroupAckRelayerFee string,
	lineaMirrorBucketRelayerFee, lineaMirrorBucketAckRelayerFee string,
	lineaMirrorObjectRelayerFee, lineaMirrorObjectAckRelayerFee string,
	lineaMirrorGroupRelayerFee, lineaMirrorGroupAckRelayerFee string,
	mantleMirrorBucketRelayerFee, mantleMirrorBucketAckRelayerFee string,
	mantleMirrorObjectRelayerFee, mantleMirrorObjectAckRelayerFee string,
	mantleMirrorGroupRelayerFee, mantleMirrorGroupAckRelayerFee string,
	arbitrumMirrorBucketRelayerFee, arbitrumMirrorBucketAckRelayerFee string,
	arbitrumMirrorObjectRelayerFee, arbitrumMirrorObjectAckRelayerFee string,
	arbitrumMirrorGroupRelayerFee, arbitrumMirrorGroupAckRelayerFee string,
	optimismMirrorBucketRelayerFee, optimismMirrorBucketAckRelayerFee string,
	optimismMirrorObjectRelayerFee, optimismMirrorObjectAckRelayerFee string,
	optimismMirrorGroupRelayerFee, optimismMirrorGroupAckRelayerFee string,
	discontinueCountingWindow, discontinueObjectMax, discontinueBucketMax uint64,
	discontinueConfirmPeriod int64,
	discontinueDeletionMax uint64,
	stalePoliesCleanupMax uint64,
	minUpdateQuotaInterval uint64,
	maxLocalVirtualGroupNumPerBucket uint32,
) Params {
	return Params{
		VersionedParams: VersionedParams{
			MaxSegmentSize:          maxSegmentSize,
			RedundantDataChunkNum:   redundantDataChunkNum,
			RedundantParityChunkNum: redundantParityChunkNum,
			MinChargeSize:           minChargeSize,
		},
		MaxPayloadSize:                    maxPayloadSize,
		MaxBucketsPerAccount:              maxBucketsPerAccount,
		BscMirrorBucketRelayerFee:         bscMirrorBucketRelayerFee,
		BscMirrorBucketAckRelayerFee:      bscMirrorBucketAckRelayerFee,
		BscMirrorObjectRelayerFee:         bscMirrorObjectRelayerFee,
		BscMirrorObjectAckRelayerFee:      bscMirrorObjectAckRelayerFee,
		BscMirrorGroupRelayerFee:          bscMirrorGroupRelayerFee,
		BscMirrorGroupAckRelayerFee:       bscMirrorGroupAckRelayerFee,
		OpMirrorBucketRelayerFee:          opMirrorBucketRelayerFee,
		OpMirrorBucketAckRelayerFee:       opMirrorBucketAckRelayerFee,
		OpMirrorObjectRelayerFee:          opMirrorObjectRelayerFee,
		OpMirrorObjectAckRelayerFee:       opMirrorObjectAckRelayerFee,
		OpMirrorGroupRelayerFee:           opMirrorGroupRelayerFee,
		OpMirrorGroupAckRelayerFee:        opMirrorGroupAckRelayerFee,
		PolygonMirrorBucketRelayerFee:     polygonMirrorBucketRelayerFee,
		PolygonMirrorBucketAckRelayerFee:  polygonMirrorBucketAckRelayerFee,
		PolygonMirrorObjectRelayerFee:     polygonMirrorObjectRelayerFee,
		PolygonMirrorObjectAckRelayerFee:  polygonMirrorObjectAckRelayerFee,
		PolygonMirrorGroupRelayerFee:      polygonMirrorGroupRelayerFee,
		PolygonMirrorGroupAckRelayerFee:   polygonMirrorGroupAckRelayerFee,
		ScrollMirrorBucketRelayerFee:      scrollMirrorBucketRelayerFee,
		ScrollMirrorBucketAckRelayerFee:   scrollMirrorBucketAckRelayerFee,
		ScrollMirrorObjectRelayerFee:      scrollMirrorObjectRelayerFee,
		ScrollMirrorObjectAckRelayerFee:   scrollMirrorObjectAckRelayerFee,
		ScrollMirrorGroupRelayerFee:       scrollMirrorGroupRelayerFee,
		ScrollMirrorGroupAckRelayerFee:    scrollMirrorGroupAckRelayerFee,
		LineaMirrorBucketRelayerFee:       lineaMirrorBucketRelayerFee,
		LineaMirrorBucketAckRelayerFee:    lineaMirrorBucketAckRelayerFee,
		LineaMirrorObjectRelayerFee:       lineaMirrorObjectRelayerFee,
		LineaMirrorObjectAckRelayerFee:    lineaMirrorObjectAckRelayerFee,
		LineaMirrorGroupRelayerFee:        lineaMirrorGroupRelayerFee,
		LineaMirrorGroupAckRelayerFee:     lineaMirrorGroupAckRelayerFee,
		MantleMirrorBucketRelayerFee:      mantleMirrorBucketRelayerFee,
		MantleMirrorBucketAckRelayerFee:   mantleMirrorBucketAckRelayerFee,
		MantleMirrorObjectRelayerFee:      mantleMirrorObjectRelayerFee,
		MantleMirrorObjectAckRelayerFee:   mantleMirrorObjectAckRelayerFee,
		MantleMirrorGroupRelayerFee:       mantleMirrorGroupRelayerFee,
		MantleMirrorGroupAckRelayerFee:    mantleMirrorGroupAckRelayerFee,
		ArbitrumMirrorBucketRelayerFee:    arbitrumMirrorBucketRelayerFee,
		ArbitrumMirrorBucketAckRelayerFee: arbitrumMirrorBucketAckRelayerFee,
		ArbitrumMirrorObjectRelayerFee:    arbitrumMirrorObjectRelayerFee,
		ArbitrumMirrorObjectAckRelayerFee: arbitrumMirrorObjectAckRelayerFee,
		ArbitrumMirrorGroupRelayerFee:     arbitrumMirrorGroupRelayerFee,
		ArbitrumMirrorGroupAckRelayerFee:  arbitrumMirrorGroupAckRelayerFee,
		OptimismMirrorBucketRelayerFee:    optimismMirrorBucketRelayerFee,
		OptimismMirrorBucketAckRelayerFee: optimismMirrorBucketAckRelayerFee,
		OptimismMirrorObjectRelayerFee:    optimismMirrorObjectRelayerFee,
		OptimismMirrorObjectAckRelayerFee: optimismMirrorObjectAckRelayerFee,
		OptimismMirrorGroupRelayerFee:     optimismMirrorGroupRelayerFee,
		OptimismMirrorGroupAckRelayerFee:  optimismMirrorGroupAckRelayerFee,
		DiscontinueCountingWindow:         discontinueCountingWindow,
		DiscontinueObjectMax:              discontinueObjectMax,
		DiscontinueBucketMax:              discontinueBucketMax,
		DiscontinueConfirmPeriod:          discontinueConfirmPeriod,
		DiscontinueDeletionMax:            discontinueDeletionMax,
		StalePolicyCleanupMax:             stalePoliesCleanupMax,
		MinQuotaUpdateInterval:            minUpdateQuotaInterval,
		MaxLocalVirtualGroupNumPerBucket:  maxLocalVirtualGroupNumPerBucket,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMaxSegmentSize, DefaultRedundantDataChunkNum,
		DefaultRedundantParityChunkNum, DefaultMaxPayloadSize, DefaultMaxBucketsPerAccount,
		DefaultMinChargeSize, DefaultBscMirrorBucketRelayerFee, DefaultBscMirrorBucketAckRelayerFee,
		DefaultBscMirrorObjectRelayerFee, DefaultBscMirrorObjectAckRelayerFee,
		DefaultBscMirrorGroupRelayerFee, DefaultBscMirrorGroupAckRelayerFee,
		DefaultOpMirrorBucketRelayerFee, DefaultOpMirrorBucketAckRelayerFee,
		DefaultOpMirrorObjectRelayerFee, DefaultOpMirrorObjectAckRelayerFee,
		DefaultOpMirrorGroupRelayerFee, DefaultOpMirrorGroupAckRelayerFee,
		DefaultPolygonMirrorBucketRelayerFee, DefaultPolygonMirrorBucketAckRelayerFee,
		DefaultPolygonMirrorObjectRelayerFee, DefaultPolygonMirrorObjectAckRelayerFee,
		DefaultPolygonMirrorGroupRelayerFee, DefaultPolygonMirrorGroupAckRelayerFee,
		DefaultScrollMirrorBucketRelayerFee, DefaultScrollMirrorBucketAckRelayerFee,
		DefaultScrollMirrorObjectRelayerFee, DefaultScrollMirrorObjectAckRelayerFee,
		DefaultScrollMirrorGroupRelayerFee, DefaultScrollMirrorGroupAckRelayerFee,
		DefaultLineaMirrorBucketRelayerFee, DefaultLineaMirrorBucketAckRelayerFee,
		DefaultLineaMirrorObjectRelayerFee, DefaultLineaMirrorObjectAckRelayerFee,
		DefaultLineaMirrorGroupRelayerFee, DefaultLineaMirrorGroupAckRelayerFee,
		DefaultMantleMirrorBucketRelayerFee, DefaultMantleMirrorBucketAckRelayerFee,
		DefaultMantleMirrorObjectRelayerFee, DefaultMantleMirrorObjectAckRelayerFee,
		DefaultMantleMirrorGroupRelayerFee, DefaultMantleMirrorGroupAckRelayerFee,
		DefaultArbitrumMirrorBucketRelayerFee, DefaultArbitrumMirrorBucketAckRelayerFee,
		DefaultArbitrumMirrorObjectRelayerFee, DefaultArbitrumMirrorObjectAckRelayerFee,
		DefaultArbitrumMirrorGroupRelayerFee, DefaultArbitrumMirrorGroupAckRelayerFee,
		DefaultOptimismMirrorBucketRelayerFee, DefaultOptimismMirrorBucketAckRelayerFee,
		DefaultOptimismMirrorObjectRelayerFee, DefaultOptimismMirrorObjectAckRelayerFee,
		DefaultOptimismMirrorGroupRelayerFee, DefaultOptimismMirrorGroupAckRelayerFee,
		DefaultDiscontinueCountingWindow, DefaultDiscontinueObjectMax, DefaultDiscontinueBucketMax,
		DefaultDiscontinueConfirmPeriod, DefaultDiscontinueDeletionMax, DefaultStalePolicyCleanupMax,
		DefaultMinUpdateQuotaInterval, DefaultMaxLocalVirtualGroupNumPerBucket,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxSegmentSize, &p.VersionedParams.MaxSegmentSize, validateMaxSegmentSize),
		paramtypes.NewParamSetPair(KeyRedundantDataChunkNum, &p.VersionedParams.RedundantDataChunkNum, validateRedundantDataChunkNum),
		paramtypes.NewParamSetPair(KeyRedundantParityChunkNum, &p.VersionedParams.RedundantParityChunkNum, validateRedundantParityChunkNum),
		paramtypes.NewParamSetPair(KeyMinChargeSize, &p.VersionedParams.MinChargeSize, validateMinChargeSize),

		paramtypes.NewParamSetPair(KeyMaxPayloadSize, &p.MaxPayloadSize, validateMaxPayloadSize),
		paramtypes.NewParamSetPair(KeyMaxBucketsPerAccount, &p.MaxBucketsPerAccount, validateMaxBucketsPerAccount),
		paramtypes.NewParamSetPair(KeyBscMirrorBucketRelayerFee, &p.BscMirrorBucketRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyBscMirrorBucketAckRelayerFee, &p.BscMirrorBucketAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyBscMirrorObjectRelayerFee, &p.BscMirrorObjectRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyBscMirrorObjectAckRelayerFee, &p.BscMirrorObjectAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyBscMirrorGroupRelayerFee, &p.BscMirrorGroupRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyBscMirrorGroupAckRelayerFee, &p.BscMirrorGroupAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOpMirrorBucketRelayerFee, &p.OpMirrorBucketRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOpMirrorBucketAckRelayerFee, &p.OpMirrorBucketAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOpMirrorObjectRelayerFee, &p.OpMirrorObjectRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOpMirrorObjectAckRelayerFee, &p.OpMirrorObjectAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOpMirrorGroupRelayerFee, &p.OpMirrorGroupRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOpMirrorGroupAckRelayerFee, &p.OpMirrorGroupAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyPolygonMirrorBucketRelayerFee, &p.PolygonMirrorBucketRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyPolygonMirrorBucketAckRelayerFee, &p.PolygonMirrorBucketAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyPolygonMirrorObjectRelayerFee, &p.PolygonMirrorObjectRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyPolygonMirrorObjectAckRelayerFee, &p.PolygonMirrorObjectAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyPolygonMirrorGroupRelayerFee, &p.PolygonMirrorGroupRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyPolygonMirrorGroupAckRelayerFee, &p.PolygonMirrorGroupAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyScrollMirrorBucketRelayerFee, &p.ScrollMirrorBucketRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyScrollMirrorBucketAckRelayerFee, &p.ScrollMirrorBucketAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyScrollMirrorObjectRelayerFee, &p.ScrollMirrorObjectRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyScrollMirrorObjectAckRelayerFee, &p.ScrollMirrorObjectAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyScrollMirrorGroupRelayerFee, &p.ScrollMirrorGroupRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyScrollMirrorGroupAckRelayerFee, &p.ScrollMirrorGroupAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyLineaMirrorBucketRelayerFee, &p.LineaMirrorBucketRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyLineaMirrorBucketAckRelayerFee, &p.LineaMirrorBucketAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyLineaMirrorObjectRelayerFee, &p.LineaMirrorObjectRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyLineaMirrorObjectAckRelayerFee, &p.LineaMirrorObjectAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyLineaMirrorGroupRelayerFee, &p.LineaMirrorGroupRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyLineaMirrorGroupAckRelayerFee, &p.LineaMirrorGroupAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyMantleMirrorBucketRelayerFee, &p.MantleMirrorBucketRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyMantleMirrorBucketAckRelayerFee, &p.MantleMirrorBucketAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyMantleMirrorObjectRelayerFee, &p.MantleMirrorObjectRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyMantleMirrorObjectAckRelayerFee, &p.MantleMirrorObjectAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyMantleMirrorGroupRelayerFee, &p.MantleMirrorGroupRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyMantleMirrorGroupAckRelayerFee, &p.MantleMirrorGroupAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyArbitrumMirrorBucketRelayerFee, &p.ArbitrumMirrorBucketRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyArbitrumMirrorBucketAckRelayerFee, &p.ArbitrumMirrorBucketAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyArbitrumMirrorObjectRelayerFee, &p.ArbitrumMirrorObjectRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyArbitrumMirrorObjectAckRelayerFee, &p.ArbitrumMirrorObjectAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyArbitrumMirrorGroupRelayerFee, &p.ArbitrumMirrorGroupRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyArbitrumMirrorGroupAckRelayerFee, &p.ArbitrumMirrorGroupAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOptimismMirrorBucketRelayerFee, &p.OptimismMirrorBucketRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOptimismMirrorBucketAckRelayerFee, &p.OptimismMirrorBucketAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOptimismMirrorObjectRelayerFee, &p.OptimismMirrorObjectRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOptimismMirrorObjectAckRelayerFee, &p.OptimismMirrorObjectAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOptimismMirrorGroupRelayerFee, &p.OptimismMirrorGroupRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyOptimismMirrorGroupAckRelayerFee, &p.OptimismMirrorGroupAckRelayerFee, validateRelayerFee),
		paramtypes.NewParamSetPair(KeyDiscontinueCountingWindow, &p.DiscontinueCountingWindow, validateDiscontinueCountingWindow),
		paramtypes.NewParamSetPair(KeyDiscontinueObjectMax, &p.DiscontinueObjectMax, validateDiscontinueObjectMax),
		paramtypes.NewParamSetPair(KeyDiscontinueBucketMax, &p.DiscontinueBucketMax, validateDiscontinueBucketMax),
		paramtypes.NewParamSetPair(KeyDiscontinueConfirmPeriod, &p.DiscontinueConfirmPeriod, validateDiscontinueConfirmPeriod),
		paramtypes.NewParamSetPair(KeyDiscontinueDeletionMax, &p.DiscontinueDeletionMax, validateDiscontinueDeletionMax),
		paramtypes.NewParamSetPair(KeyStalePolicyCleanupMax, &p.StalePolicyCleanupMax, validateStalePolicyCleanupMax),
		paramtypes.NewParamSetPair(KeyMinUpdateQuotaInterval, &p.MinQuotaUpdateInterval, validateMinUpdateQuotaInterval),
		paramtypes.NewParamSetPair(KeyMaxLocalVirtualGroupNumPerBucket, &p.MaxLocalVirtualGroupNumPerBucket, validateMaxLocalVirtualGroupNumPerBucket),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMaxSegmentSize(p.VersionedParams.MaxSegmentSize); err != nil {
		return err
	}
	if err := validateRedundantDataChunkNum(p.VersionedParams.RedundantDataChunkNum); err != nil {
		return err
	}
	if err := validateRedundantParityChunkNum(p.VersionedParams.RedundantParityChunkNum); err != nil {
		return err
	}
	if err := validateMinChargeSize(p.VersionedParams.MinChargeSize); err != nil {
		return err
	}
	if err := validateMaxPayloadSize(p.MaxPayloadSize); err != nil {
		return err
	}
	if err := validateMaxBucketsPerAccount(p.MaxBucketsPerAccount); err != nil {
		return err
	}
	if err := validateRelayerFee(p.BscMirrorBucketRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.BscMirrorBucketAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.BscMirrorObjectRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.BscMirrorObjectAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.BscMirrorGroupRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.BscMirrorGroupAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OpMirrorBucketRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OpMirrorBucketAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OpMirrorObjectRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OpMirrorObjectAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OpMirrorGroupRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OpMirrorGroupAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.PolygonMirrorBucketRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.PolygonMirrorBucketAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.PolygonMirrorObjectRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.PolygonMirrorObjectAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.PolygonMirrorGroupRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.PolygonMirrorGroupAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ScrollMirrorBucketRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ScrollMirrorBucketAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ScrollMirrorObjectRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ScrollMirrorObjectAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ScrollMirrorGroupRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ScrollMirrorGroupAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.LineaMirrorBucketRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.LineaMirrorBucketAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.LineaMirrorObjectRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.LineaMirrorObjectAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.LineaMirrorGroupRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.LineaMirrorGroupAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.MantleMirrorBucketRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.MantleMirrorBucketAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.MantleMirrorObjectRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.MantleMirrorObjectAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.MantleMirrorGroupRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.MantleMirrorGroupAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ArbitrumMirrorBucketRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ArbitrumMirrorBucketAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ArbitrumMirrorObjectRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ArbitrumMirrorObjectAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ArbitrumMirrorGroupRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.ArbitrumMirrorGroupAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OptimismMirrorBucketRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OptimismMirrorBucketAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OptimismMirrorObjectRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OptimismMirrorObjectAckRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OptimismMirrorGroupRelayerFee); err != nil {
		return err
	}
	if err := validateRelayerFee(p.OptimismMirrorGroupAckRelayerFee); err != nil {
		return err
	}
	if err := validateDiscontinueCountingWindow(p.DiscontinueCountingWindow); err != nil {
		return err
	}
	if err := validateDiscontinueObjectMax(p.DiscontinueObjectMax); err != nil {
		return err
	}
	if err := validateDiscontinueBucketMax(p.DiscontinueBucketMax); err != nil {
		return err
	}
	if err := validateDiscontinueConfirmPeriod(p.DiscontinueConfirmPeriod); err != nil {
		return err
	}
	if err := validateDiscontinueDeletionMax(p.DiscontinueDeletionMax); err != nil {
		return err
	}
	if err := validateStalePolicyCleanupMax(p.StalePolicyCleanupMax); err != nil {
		return err
	}
	if err := validateMinUpdateQuotaInterval(p.MinQuotaUpdateInterval); err != nil {
		return err
	}
	if err := validateMaxLocalVirtualGroupNumPerBucket(p.MaxLocalVirtualGroupNumPerBucket); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// String implements the Stringer interface.
func (p VersionedParams) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (p *Params) GetMaxSegmentSize() uint64 {
	if p != nil {
		return p.VersionedParams.MaxSegmentSize
	}
	return 0
}

func (p *Params) GetRedundantDataChunkNum() uint32 {
	if p != nil {
		return p.VersionedParams.RedundantDataChunkNum
	}
	return 0
}

func (p *Params) GetRedundantParityChunkNum() uint32 {
	if p != nil {
		return p.VersionedParams.RedundantParityChunkNum
	}
	return 0
}

func (p *Params) GetMinChargeSize() uint64 {
	if p != nil {
		return p.VersionedParams.MinChargeSize
	}
	return 0
}

func validateMaxSegmentSize(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max segment size must be positive: %d", v)
	}

	return nil
}

func validateMaxPayloadSize(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max payload size must be positive: %d", v)
	}

	return nil
}

func validateMaxBucketsPerAccount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max buckets per account must be positive: %d", v)
	}

	return nil
}

func validateMinChargeSize(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("min charge size must be positive: %d", v)
	}

	return nil
}

func validateRedundantDataChunkNum(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("redundant data chunk num must be positive: %d", v)
	}

	return nil
}

func validateRedundantParityChunkNum(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("redundant parity size chunk num must be positive: %d", v)
	}

	return nil
}

func validateRelayerFee(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	relayerFee := big.NewInt(0)
	relayerFee, valid := relayerFee.SetString(v, 10)

	if !valid {
		return fmt.Errorf("invalid transfer out relayer fee, %s", v)
	}

	if relayerFee.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf("invalid transfer out relayer fee, %s", v)
	}

	return nil
}

func validateDiscontinueCountingWindow(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("discontinue counting window must be positive: %d", v)
	}

	return nil
}

func validateDiscontinueObjectMax(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateDiscontinueBucketMax(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateDiscontinueConfirmPeriod(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("discontinue confirm period must be positive: %d", v)
	}
	return nil
}

func validateDiscontinueDeletionMax(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("discontinue deletion max must be positive: %d", v)
	}
	return nil
}

func validateStalePolicyCleanupMax(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf(" max stale policy to cleanup must be positive: %d", v)
	}
	return nil
}

func validateMinUpdateQuotaInterval(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateMaxLocalVirtualGroupNumPerBucket(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max LVG per bucket must be positive: %d", v)
	}

	return nil
}
