package keeper_test

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	types2 "github.com/evmos/evmos/v12/sdk/types"
	"github.com/evmos/evmos/v12/testutil/sample"
	"github.com/evmos/evmos/v12/x/sp/types"
)

func (s *KeeperTestSuite) TestSetGetStorageProvider() {
	keeper := s.spKeeper
	ctx := s.ctx
	sp := &types.StorageProvider{Id: 100}
	spAccStr := sample.RandAccAddressHex()
	spAcc := sdk.MustAccAddressFromHex(spAccStr)
	sp.OperatorAddress = spAcc.String()

	keeper.SetStorageProvider(ctx, sp)
	_, found := keeper.GetStorageProvider(ctx, 100)
	if !found {
		fmt.Printf("no such sp: %s", spAcc)
	}
	require.EqualValues(s.T(), found, true)
}

// TestStorageProviderBasics tests GetStorageProviderByOperatorAddr, GetStorageProviderByFundingAddr,
// GetStorageProviderBySealAddr, GetStorageProviderByApprovalAddr, GetStorageProviderByBlsKey
func (s *KeeperTestSuite) TestStorageProviderBasics() {
	k := s.spKeeper
	ctx := s.ctx
	spAccStr := sample.RandAccAddressHex()
	spAcc := sdk.MustAccAddressFromHex(spAccStr)

	fundingAccStr := sample.RandAccAddressHex()
	fundingAcc := sdk.MustAccAddressFromHex(fundingAccStr)

	sealAccStr := sample.RandAccAddressHex()
	sealAcc := sdk.MustAccAddressFromHex(sealAccStr)

	approvalAccStr := sample.RandAccAddressHex()
	approvalAcc := sdk.MustAccAddressFromHex(approvalAccStr)

	blsPubKey := sample.RandBlsPubKey()
	sp := &types.StorageProvider{
		Id:              100,
		OperatorAddress: spAcc.String(),
		FundingAddress:  fundingAcc.String(),
		SealAddress:     sealAcc.String(),
		ApprovalAddress: approvalAcc.String(),
		BlsKey:          blsPubKey,
	}

	k.SetStorageProvider(ctx, sp)
	_, found := k.GetStorageProvider(ctx, 100)
	if !found {
		fmt.Printf("no such sp: %s", spAcc)
	}
	require.EqualValues(s.T(), found, true)

	k.SetStorageProviderByFundingAddr(ctx, sp)
	_, found = k.GetStorageProviderByFundingAddr(ctx, fundingAcc)
	if !found {
		fmt.Printf("no such sp: %s", spAcc)
	}
	require.EqualValues(s.T(), found, true)

	k.SetStorageProviderBySealAddr(ctx, sp)
	_, found = k.GetStorageProviderBySealAddr(ctx, sealAcc)
	if !found {
		fmt.Printf("no such sp: %s", spAcc)
	}
	require.EqualValues(s.T(), found, true)

	k.SetStorageProviderByApprovalAddr(ctx, sp)
	_, found = k.GetStorageProviderByApprovalAddr(ctx, approvalAcc)
	if !found {
		fmt.Printf("no such sp: %s", spAcc)
	}
	require.EqualValues(s.T(), found, true)

	k.SetStorageProviderByBlsKey(ctx, sp)
	_, found = k.GetStorageProviderByBlsKey(ctx, blsPubKey)
	if !found {
		fmt.Printf("no such sp: %s", spAcc)
	}
	require.EqualValues(s.T(), found, true)
}

func (s *KeeperTestSuite) TestSlashBasic() {
	// mock
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	k := s.spKeeper
	ctx := s.ctx
	spAccStr := sample.RandAccAddressHex()
	spAcc := sdk.MustAccAddressFromHex(spAccStr)

	fundingAccStr := sample.RandAccAddressHex()
	fundingAcc := sdk.MustAccAddressFromHex(fundingAccStr)

	sealAccStr := sample.RandAccAddressHex()
	sealAcc := sdk.MustAccAddressFromHex(sealAccStr)

	approvalAccStr := sample.RandAccAddressHex()
	approvalAcc := sdk.MustAccAddressFromHex(approvalAccStr)

	blsPubKey := sample.RandBlsPubKey()

	sp := &types.StorageProvider{
		Id:              100,
		OperatorAddress: spAcc.String(),
		FundingAddress:  fundingAcc.String(),
		SealAddress:     sealAcc.String(),
		ApprovalAddress: approvalAcc.String(),
		BlsKey:          blsPubKey,
		TotalDeposit:    math.NewIntWithDecimal(2010, types2.DecimalZKME),
	}

	k.SetStorageProvider(ctx, sp)
	_, found := k.GetStorageProvider(ctx, 100)
	if !found {
		fmt.Printf("no such sp: %s", spAcc)
	}
	require.EqualValues(s.T(), found, true)

	rewardInfo := types.RewardInfo{
		Address: sample.RandAccAddressHex(),
		Amount:  sdk.NewCoin(types2.Denom, math.NewIntWithDecimal(10, types2.DecimalZKME)),
	}

	err := k.Slash(ctx, sp.Id, []types.RewardInfo{rewardInfo})
	require.NoError(s.T(), err)

	spAfterSlash, found := k.GetStorageProvider(ctx, 100)
	require.True(s.T(), found)
	s.T().Logf("%s", spAfterSlash.TotalDeposit.String())
	require.True(s.T(), spAfterSlash.TotalDeposit.Equal(math.NewIntWithDecimal(2000, types2.DecimalZKME)))
}
