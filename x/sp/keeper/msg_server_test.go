package keeper_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/evmos/evmos/v12/sdk/types"
	"github.com/evmos/evmos/v12/testutil/sample"
	"github.com/evmos/evmos/v12/x/sp/keeper"
	sptypes "github.com/evmos/evmos/v12/x/sp/types"
)

func (s *KeeperTestSuite) TestMsgCreateStorageProvider() {
	govAddr := authtypes.NewModuleAddress(gov.ModuleName)
	// 1. create new newStorageProvider and grant

	operatorAddr, _, err := testutil.GenerateCoinKey(hd.Secp256k1, s.cdc)
	s.Require().Nil(err, "error should be nil")
	fundingAddr, _, err := testutil.GenerateCoinKey(hd.Secp256k1, s.cdc)
	s.Require().Nil(err, "error should be nil")
	sealAddr, _, err := testutil.GenerateCoinKey(hd.Secp256k1, s.cdc)
	s.Require().Nil(err, "error should be nil")
	approvalAddr, _, err := testutil.GenerateCoinKey(hd.Secp256k1, s.cdc)
	s.Require().Nil(err, "error should be nil")
	gcAddr, _, err := testutil.GenerateCoinKey(hd.Secp256k1, s.cdc)
	s.Require().Nil(err, "error should be nil")
	maintenanceAddr, _, err := testutil.GenerateCoinKey(hd.Secp256k1, s.cdc)
	s.Require().Nil(err, "error should be nil")

	blsPubKeyHex := sample.RandBlsPubKeyHex()

	s.accountKeeper.EXPECT().GetAccount(gomock.Any(), fundingAddr).Return(authtypes.NewBaseAccountWithAddress(fundingAddr)).AnyTimes()
	s.accountKeeper.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	testCases := []struct {
		Name      string
		ExceptErr bool
		req       types.MsgCreateStorageProvider
	}{
		{
			Name:      "invalid funding address",
			ExceptErr: true,
			req: types.MsgCreateStorageProvider{
				Creator: govAddr.String(),
				Description: sptypes.Description{
					Moniker:  "sp_test",
					Identity: "",
				},
				SpAddress:          operatorAddr.String(),
				FundingAddress:     sample.RandAccAddressHex(),
				SealAddress:        sealAddr.String(),
				ApprovalAddress:    approvalAddr.String(),
				GcAddress:          gcAddr.String(),
				MaintenanceAddress: maintenanceAddr.String(),
				BlsKey:             blsPubKeyHex,
				Deposit: sdk.Coin{
					Denom:  types.Denom,
					Amount: types.NewIntFromInt64WithDecimal(10000, types.DecimalZKME),
				},
			},
		},
		{
			Name:      "invalid endpoint",
			ExceptErr: true,
			req: types.MsgCreateStorageProvider{
				Creator: govAddr.String(),
				Description: sptypes.Description{
					Moniker:  "sp_test",
					Identity: "",
				},
				SpAddress:          operatorAddr.String(),
				FundingAddress:     fundingAddr.String(),
				SealAddress:        sealAddr.String(),
				ApprovalAddress:    approvalAddr.String(),
				GcAddress:          gcAddr.String(),
				MaintenanceAddress: maintenanceAddr.String(),
				BlsKey:             blsPubKeyHex,
				Endpoint:           "sp.io",
				Deposit: sdk.Coin{
					Denom:  types.Denom,
					Amount: types.NewIntFromInt64WithDecimal(10000, types.DecimalZKME),
				},
			},
		},
		{
			Name:      "invalid bls pub key",
			ExceptErr: true,
			req: types.MsgCreateStorageProvider{
				Creator: govAddr.String(),
				Description: sptypes.Description{
					Moniker:  "sp_test",
					Identity: "",
				},
				SpAddress:          operatorAddr.String(),
				FundingAddress:     fundingAddr.String(),
				SealAddress:        sealAddr.String(),
				ApprovalAddress:    approvalAddr.String(),
				GcAddress:          gcAddr.String(),
				MaintenanceAddress: maintenanceAddr.String(),
				BlsKey:             "InValidBlsPubkey",
				Endpoint:           "sp.io",
				Deposit: sdk.Coin{
					Denom:  types.Denom,
					Amount: types.NewIntFromInt64WithDecimal(10000, types.DecimalZKME),
				},
			},
		},
		{
			Name:      "success",
			ExceptErr: true,
			req: types.MsgCreateStorageProvider{
				Creator: govAddr.String(),
				Description: sptypes.Description{
					Moniker:  "MsgServer_sp_test",
					Identity: "",
				},
				SpAddress:          operatorAddr.String(),
				FundingAddress:     fundingAddr.String(),
				SealAddress:        sealAddr.String(),
				ApprovalAddress:    approvalAddr.String(),
				GcAddress:          gcAddr.String(),
				MaintenanceAddress: maintenanceAddr.String(),
				BlsKey:             blsPubKeyHex,
				Deposit: sdk.Coin{
					Denom:  types.Denom,
					Amount: types.NewIntFromInt64WithDecimal(10000, types.DecimalZKME),
				},
			},
		},
	}
	for _, testCase := range testCases {
		s.Suite.T().Run(testCase.Name, func(t *testing.T) {
			req := testCase.req
			_, err := s.msgServer.CreateStorageProvider(s.ctx, &req)
			if testCase.ExceptErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func (s *KeeperTestSuite) TestIsLastDaysOfTheMonth() {
	s.Require().True(!keeper.IsLastDaysOfTheMonth(time.Unix(1693242061, 0), 2)) // 2023-08-28 UTC
	s.Require().True(!keeper.IsLastDaysOfTheMonth(time.Unix(1693328461, 0), 2)) // 2023-08-29 UTC
	s.Require().True(keeper.IsLastDaysOfTheMonth(time.Unix(1693414861, 0), 2))  // 2023-08-30 UTC
	s.Require().True(!keeper.IsLastDaysOfTheMonth(time.Unix(1693587661, 0), 2)) // 2023-09-01 UTC
}
