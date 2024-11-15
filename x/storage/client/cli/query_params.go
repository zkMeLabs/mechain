package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	evmostypes "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/evm/precompiles/storage"
	"github.com/evmos/evmos/v12/x/storage/types"
)

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the parameters of the storage module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			contract, err := storage.NewIStorage(common.HexToAddress(evmostypes.StorageAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.Params(&bind.CallOpts{})
			if err != nil {
				return err
			}
			res := &types.QueryParamsResponse{
				Params: types.Params{
					VersionedParams: types.VersionedParams{
						MaxSegmentSize:          result.VersionedParams.MaxSegmentSize,
						RedundantDataChunkNum:   result.VersionedParams.RedundantDataChunkNum,
						RedundantParityChunkNum: result.VersionedParams.RedundantParityChunkNum,
						MinChargeSize:           result.VersionedParams.MinChargeSize,
					},
					MaxPayloadSize:                   result.MaxPayloadSize,
					BscMirrorBucketRelayerFee:        result.BscMirrorBucketRelayerFee,
					BscMirrorBucketAckRelayerFee:     result.BscMirrorBucketAckRelayerFee,
					BscMirrorObjectRelayerFee:        result.BscMirrorObjectRelayerFee,
					BscMirrorObjectAckRelayerFee:     result.BscMirrorObjectAckRelayerFee,
					BscMirrorGroupRelayerFee:         result.BscMirrorGroupRelayerFee,
					BscMirrorGroupAckRelayerFee:      result.BscMirrorGroupAckRelayerFee,
					MaxBucketsPerAccount:             result.MaxBucketsPerAccount,
					DiscontinueCountingWindow:        result.DiscontinueCountingWindow,
					DiscontinueObjectMax:             result.DiscontinueObjectMax,
					DiscontinueBucketMax:             result.DiscontinueBucketMax,
					DiscontinueConfirmPeriod:         result.DiscontinueConfirmPeriod,
					DiscontinueDeletionMax:           result.DiscontinueDeletionMax,
					StalePolicyCleanupMax:            result.StalePolicyCleanupMax,
					MinQuotaUpdateInterval:           result.MinQuotaUpdateInterval,
					MaxLocalVirtualGroupNumPerBucket: result.MaxLocalVirtualGroupNumPerBucket,
					OpMirrorBucketRelayerFee:         result.OpMirrorBucketRelayerFee,
					OpMirrorBucketAckRelayerFee:      result.OpMirrorBucketAckRelayerFee,
					OpMirrorObjectRelayerFee:         result.OpMirrorObjectRelayerFee,
					OpMirrorObjectAckRelayerFee:      result.OpMirrorObjectAckRelayerFee,
					OpMirrorGroupRelayerFee:          result.OpMirrorGroupRelayerFee,
					OpMirrorGroupAckRelayerFee:       result.OpMirrorGroupAckRelayerFee,
					PolygonMirrorBucketRelayerFee:    result.PolygonMirrorBucketRelayerFee,
					PolygonMirrorBucketAckRelayerFee: result.PolygonMirrorBucketAckRelayerFee,
				}}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
