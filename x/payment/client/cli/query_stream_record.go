package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	evmostypes "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/evm/precompiles/payment"
	"github.com/evmos/evmos/v12/x/payment/types"
)

func CmdListStreamRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-stream-record",
		Short: "list all stream-record",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			contract, err := payment.NewIPayment(common.HexToAddress(evmostypes.PaymentAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.StreamRecords(&bind.CallOpts{}, *ToPaymentPageReq(pageReq))
			if err != nil {
				return err
			}

			streamRecords := make([]types.StreamRecord, 0)
			for _, streamRecord := range result.StreamRecords {
				streamRecords = append(streamRecords, *ToStreamRecord(&streamRecord))
			}
			res := &types.QueryStreamRecordsResponse{
				StreamRecords: streamRecords,
				Pagination:    ToPageResp(&result.PageResponse),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowStreamRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-stream-record [account]",
		Short: "shows a stream-record",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			argAccount := args[0]

			contract, err := payment.NewIPayment(common.HexToAddress(evmostypes.PaymentAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.StreamRecord(&bind.CallOpts{}, argAccount)
			if err != nil {
				return err
			}

			res := &types.QueryGetStreamRecordResponse{
				StreamRecord: *ToStreamRecord(&result),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
