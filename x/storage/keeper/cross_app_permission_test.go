package keeper_test

import (
	"math/big"
	"math/rand"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.uber.org/mock/gomock"

	"github.com/evmos/evmos/v12/testutil/sample"
	"github.com/evmos/evmos/v12/x/permission/types"
	"github.com/evmos/evmos/v12/x/storage/keeper"
	storageTypes "github.com/evmos/evmos/v12/x/storage/types"
)

func (s *TestSuite) TestSynCreatePolicy() {
	ctrl := gomock.NewController(s.T())
	storageKeeper := storageTypes.NewMockStorageKeeper(ctrl)
	permissionKeeper := storageTypes.NewMockPermissionKeeper(ctrl)

	resourceIds := []math.Uint{math.NewUint(rand.Uint64()), math.NewUint(rand.Uint64()), math.NewUint(rand.Uint64())} //nolint
	// policy without expiry
	policy := types.Policy{
		Principal: &types.Principal{
			Type:  types.PRINCIPAL_TYPE_GNFD_ACCOUNT,
			Value: sample.RandAccAddressHex(),
		},
		ResourceType:   1,
		ResourceId:     resourceIds[0],
		Statements:     nil,
		ExpirationTime: nil,
	}

	app := keeper.NewPermissionApp(storageKeeper, permissionKeeper)
	data, err := policy.Marshal()
	s.NoError(err)

	synPackage := storageTypes.CreatePolicySynPackage{
		Operator:  sample.RandAccAddress(),
		Data:      data,
		ExtraData: []byte("extra data"),
	}
	serializedSynPackage := synPackage.MustSerialize()
	serializedSynPackage = append([]byte{storageTypes.OperationCreatePolicy}, serializedSynPackage...)

	// case 1: bucket not found
	storageKeeper.EXPECT().GetBucketInfoById(gomock.Any(), gomock.Any()).Return(nil, false)
	res := app.ExecuteSynPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedSynPackage)
	s.Require().ErrorIs(res.Err, storageTypes.ErrNoSuchBucket)
}

func (s *TestSuite) TestSynDeletePolicy() {
	ctrl := gomock.NewController(s.T())
	storageKeeper := storageTypes.NewMockStorageKeeper(ctrl)
	permissionKeeper := storageTypes.NewMockPermissionKeeper(ctrl)

	app := keeper.NewPermissionApp(storageKeeper, permissionKeeper)
	synPackage := storageTypes.DeleteBucketSynPackage{
		Operator:  sample.RandAccAddress(),
		Id:        big.NewInt(10),
		ExtraData: []byte("extra data"),
	}

	serializedSynPackage := synPackage.MustSerialize()
	serializedSynPackage = append([]byte{storageTypes.OperationDeletePolicy}, serializedSynPackage...)

	// case 1: No such Policy
	permissionKeeper.EXPECT().GetPolicyByID(gomock.Any(), gomock.Any()).Return(&types.Policy{}, false)
	res := app.ExecuteSynPackage(s.ctx, &sdk.CrossChainAppContext{}, serializedSynPackage)
	s.Require().ErrorIs(res.Err, storageTypes.ErrNoSuchPolicy)
	s.Require().NotEmpty(res.Payload)
}
