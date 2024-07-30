package keeper_test

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v12/testutil/sample"
	"github.com/evmos/evmos/v12/x/storage/keeper"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
	"go.uber.org/mock/gomock"
)

func (s *TestSuite) TestSynDeleteBucket() {
	ctrl := gomock.NewController(s.T())
	storageKeeper := storagetypes.NewMockStorageKeeper(ctrl)
	storageKeeper.EXPECT().Logger(gomock.Any()).Return(s.ctx.Logger()).AnyTimes()

	app := keeper.NewBucketApp(storageKeeper)
	deleteSynPackage := storagetypes.DeleteBucketSynPackage{
		Operator:  sample.RandAccAddress(),
		ID:        big.NewInt(10),
		ExtraData: []byte("extra data"),
	}

	serializedSynPackage := deleteSynPackage.MustSerialize()
	serializedSynPackage = append([]byte{storagetypes.OperationDeleteBucket}, serializedSynPackage...)

	storageKeeper.EXPECT().GetSourceTypeByChainId(gomock.Any(), gomock.Any()).Return(storagetypes.SOURCE_TYPE_BSC_CROSS_CHAIN, nil).AnyTimes()

	// case 1: bucket not found
	storageKeeper.EXPECT().GetBucketInfoById(gomock.Any(), gomock.Any()).Return(nil, false)
	res := app.ExecuteSynPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedSynPackage)
	s.Require().ErrorIs(res.Err, storagetypes.ErrNoSuchBucket)

	// case 2: delete bucket error
	storageKeeper.EXPECT().GetBucketInfoById(gomock.Any(), gomock.Any()).Return(&storagetypes.BucketInfo{
		BucketName: "bucket",
	}, true)
	storageKeeper.EXPECT().DeleteBucket(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("delete error"))
	res = app.ExecuteSynPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedSynPackage)
	s.Require().ErrorContains(res.Err, "delete error")

	// case 3: delete bucket success
	storageKeeper.EXPECT().GetBucketInfoById(gomock.Any(), gomock.Any()).Return(&storagetypes.BucketInfo{
		BucketName: "bucket",
		Id:         sdk.NewUint(10),
	}, true)
	storageKeeper.EXPECT().DeleteBucket(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	res = app.ExecuteSynPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedSynPackage)
	s.Require().NoError(res.Err)
}

func (s *TestSuite) TestSynCreateBucket() {
	ctrl := gomock.NewController(s.T())
	storageKeeper := storagetypes.NewMockStorageKeeper(ctrl)
	storageKeeper.EXPECT().Logger(gomock.Any()).Return(s.ctx.Logger()).AnyTimes()

	app := keeper.NewBucketApp(storageKeeper)
	createSynPackage := storagetypes.CreateBucketSynPackage{
		Creator:          sample.RandAccAddress(),
		BucketName:       "bucketName",
		ExtraData:        []byte("extra data"),
		PaymentAddress:   sample.RandAccAddress(),
		PrimarySpAddress: sample.RandAccAddress(),
	}
	serializedSynPackage := createSynPackage.MustSerialize()
	serializedSynPackage = append([]byte{storagetypes.OperationCreateBucket}, serializedSynPackage...)

	storageKeeper.EXPECT().GetSourceTypeByChainId(gomock.Any(), gomock.Any()).Return(storagetypes.SOURCE_TYPE_BSC_CROSS_CHAIN, nil).AnyTimes()

	// case 1: invalid package
	res := app.ExecuteSynPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedSynPackage)
	s.Require().ErrorContains(res.Err, "Invalid type of visibility")

	// case 2: create bucket error
	createSynPackage.Visibility = uint32(storagetypes.VISIBILITY_TYPE_PUBLIC_READ)
	serializedSynPackage = createSynPackage.MustSerialize()
	serializedSynPackage = append([]byte{storagetypes.OperationCreateBucket}, serializedSynPackage...)

	storageKeeper.EXPECT().CreateBucket(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(sdk.NewUint(1), fmt.Errorf("create error"))
	res = app.ExecuteSynPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedSynPackage)
	s.Require().ErrorContains(res.Err, "create error")

	// case 3: create bucket success
	createSynPackage.Visibility = uint32(storagetypes.VISIBILITY_TYPE_PUBLIC_READ)
	serializedSynPackage = createSynPackage.MustSerialize()
	serializedSynPackage = append([]byte{storagetypes.OperationCreateBucket}, serializedSynPackage...)

	storageKeeper.EXPECT().CreateBucket(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(sdk.NewUint(1), nil)
	res = app.ExecuteSynPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedSynPackage)
	s.Require().NoError(res.Err)
}

func (s *TestSuite) TestSynMirrorBucket() {
	ctrl := gomock.NewController(s.T())
	storageKeeper := storagetypes.NewMockStorageKeeper(ctrl)
	storageKeeper.EXPECT().Logger(gomock.Any()).Return(s.ctx.Logger()).AnyTimes()

	app := keeper.NewBucketApp(storageKeeper)
	synPackage := storagetypes.MirrorBucketSynPackage{
		Owner: sample.RandAccAddress(),
		Id:    big.NewInt(10),
	}

	serializedSynPack, err := synPackage.Serialize()
	s.Require().NoError(err)
	serializedSynPack = append([]byte{storagetypes.OperationMirrorBucket}, serializedSynPack...)

	storageKeeper.EXPECT().GetSourceTypeByChainId(gomock.Any(), gomock.Any()).Return(storagetypes.SOURCE_TYPE_BSC_CROSS_CHAIN, nil).AnyTimes()

	// case 1:  normal case
	res := app.ExecuteSynPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedSynPack)
	s.Require().NoError(res.Err)
}

func (s *TestSuite) TestAckMirrorBucket() {
	ctrl := gomock.NewController(s.T())
	storageKeeper := storagetypes.NewMockStorageKeeper(ctrl)
	storageKeeper.EXPECT().Logger(gomock.Any()).Return(s.ctx.Logger()).AnyTimes()

	app := keeper.NewBucketApp(storageKeeper)
	ackPackage := storagetypes.MirrorBucketAckPackage{
		Status: storagetypes.StatusSuccess,
		Id:     big.NewInt(10),
	}

	serializedAckPack, err := ackPackage.Serialize()
	s.Require().NoError(err)
	serializedAckPack = append([]byte{storagetypes.OperationMirrorBucket}, serializedAckPack...)

	storageKeeper.EXPECT().GetSourceTypeByChainId(gomock.Any(), gomock.Any()).Return(storagetypes.SOURCE_TYPE_BSC_CROSS_CHAIN, nil).AnyTimes()

	// case 1: bucket not found
	storageKeeper.EXPECT().GetBucketInfoById(gomock.Any(), gomock.Any()).Return(nil, false)

	res := app.ExecuteAckPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedAckPack)
	s.Require().ErrorIs(res.Err, storagetypes.ErrNoSuchBucket)

	// case 2: success case
	storageKeeper.EXPECT().GetBucketInfoById(gomock.Any(), gomock.Any()).Return(&storagetypes.BucketInfo{}, true)
	storageKeeper.EXPECT().SetBucketInfo(gomock.Any(), gomock.Any()).Return()

	res = app.ExecuteAckPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedAckPack)
	s.Require().NoError(res.Err)
}

func (s *TestSuite) TestAckCreateBucket() {
	ctrl := gomock.NewController(s.T())
	storageKeeper := storagetypes.NewMockStorageKeeper(ctrl)
	storageKeeper.EXPECT().Logger(gomock.Any()).Return(s.ctx.Logger()).AnyTimes()

	app := keeper.NewBucketApp(storageKeeper)
	ackPackage := storagetypes.CreateBucketAckPackage{
		Status:    storagetypes.StatusSuccess,
		ID:        big.NewInt(10),
		Creator:   sample.RandAccAddress(),
		ExtraData: []byte("extra data"),
	}

	serializedAckPack := ackPackage.MustSerialize()
	serializedAckPack = append([]byte{storagetypes.OperationCreateBucket}, serializedAckPack...)

	storageKeeper.EXPECT().GetSourceTypeByChainId(gomock.Any(), gomock.Any()).Return(storagetypes.SOURCE_TYPE_BSC_CROSS_CHAIN, nil).AnyTimes()

	// case 1:  normal case
	res := app.ExecuteAckPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedAckPack)
	s.Require().NoError(res.Err)
}

func (s *TestSuite) TestAckDeleteBucket() {
	ctrl := gomock.NewController(s.T())
	storageKeeper := storagetypes.NewMockStorageKeeper(ctrl)
	storageKeeper.EXPECT().Logger(gomock.Any()).Return(s.ctx.Logger()).AnyTimes()

	app := keeper.NewBucketApp(storageKeeper)
	ackPackage := storagetypes.DeleteBucketAckPackage{
		Status:    storagetypes.StatusSuccess,
		ID:        big.NewInt(10),
		ExtraData: []byte("extra data"),
	}

	serializedAckPack := ackPackage.MustSerialize()
	serializedAckPack = append([]byte{storagetypes.OperationDeleteBucket}, serializedAckPack...)

	storageKeeper.EXPECT().GetSourceTypeByChainId(gomock.Any(), gomock.Any()).Return(storagetypes.SOURCE_TYPE_BSC_CROSS_CHAIN, nil).AnyTimes()

	// case 1:  normal case
	res := app.ExecuteAckPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedAckPack)
	s.Require().NoError(res.Err)
}

func (s *TestSuite) TestFailAckMirrorBucket() {
	ctrl := gomock.NewController(s.T())
	storageKeeper := storagetypes.NewMockStorageKeeper(ctrl)
	storageKeeper.EXPECT().Logger(gomock.Any()).Return(s.ctx.Logger()).AnyTimes()

	app := keeper.NewBucketApp(storageKeeper)
	ackPackage := storagetypes.MirrorBucketSynPackage{
		Id:    big.NewInt(10),
		Owner: sample.RandAccAddress(),
	}

	serializedAckPack, err := ackPackage.Serialize()
	s.Require().NoError(err)
	serializedAckPack = append([]byte{storagetypes.OperationMirrorBucket}, serializedAckPack...)

	storageKeeper.EXPECT().GetSourceTypeByChainId(gomock.Any(), gomock.Any()).Return(storagetypes.SOURCE_TYPE_BSC_CROSS_CHAIN, nil).AnyTimes()

	// case 1:  bucket not found
	storageKeeper.EXPECT().GetBucketInfoById(gomock.Any(), gomock.Any()).Return(&storagetypes.BucketInfo{}, false)

	res := app.ExecuteFailAckPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedAckPack)
	s.Require().ErrorIs(res.Err, storagetypes.ErrNoSuchBucket)

	// case 2: normal case
	storageKeeper.EXPECT().GetBucketInfoById(gomock.Any(), gomock.Any()).Return(&storagetypes.BucketInfo{}, true)
	storageKeeper.EXPECT().SetBucketInfo(gomock.Any(), gomock.Any()).Return()

	res = app.ExecuteFailAckPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedAckPack)
	s.Require().NoError(res.Err)
}

func (s *TestSuite) TestFailAckCreateBucket() {
	ctrl := gomock.NewController(s.T())
	storageKeeper := storagetypes.NewMockStorageKeeper(ctrl)
	storageKeeper.EXPECT().Logger(gomock.Any()).Return(s.ctx.Logger()).AnyTimes()

	app := keeper.NewBucketApp(storageKeeper)
	createSynPackage := storagetypes.CreateBucketSynPackage{
		Creator:          sample.RandAccAddress(),
		BucketName:       "bucketname",
		ExtraData:        []byte("extra data"),
		PaymentAddress:   sample.RandAccAddress(),
		PrimarySpAddress: sample.RandAccAddress(),
	}
	serializedSynPackage := createSynPackage.MustSerialize()
	serializedSynPackage = append([]byte{storagetypes.OperationCreateBucket}, serializedSynPackage...)

	storageKeeper.EXPECT().GetSourceTypeByChainId(gomock.Any(), gomock.Any()).Return(storagetypes.SOURCE_TYPE_BSC_CROSS_CHAIN, nil).AnyTimes()

	// case 1:  normal case
	res := app.ExecuteFailAckPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedSynPackage)
	s.Require().NoError(res.Err)
}

func (s *TestSuite) TestFailAckDeleteBucket() {
	ctrl := gomock.NewController(s.T())
	storageKeeper := storagetypes.NewMockStorageKeeper(ctrl)
	storageKeeper.EXPECT().Logger(gomock.Any()).Return(s.ctx.Logger()).AnyTimes()

	app := keeper.NewBucketApp(storageKeeper)
	deleteSynPackage := storagetypes.DeleteBucketSynPackage{
		Operator:  sample.RandAccAddress(),
		ID:        big.NewInt(10),
		ExtraData: []byte("extra data"),
	}

	serializedSynPackage := deleteSynPackage.MustSerialize()
	serializedSynPackage = append([]byte{storagetypes.OperationDeleteBucket}, serializedSynPackage...)

	storageKeeper.EXPECT().GetSourceTypeByChainId(gomock.Any(), gomock.Any()).Return(storagetypes.SOURCE_TYPE_BSC_CROSS_CHAIN, nil).AnyTimes()

	// case 1:  normal case
	res := app.ExecuteFailAckPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedSynPackage)
	s.Require().NoError(res.Err)
}
