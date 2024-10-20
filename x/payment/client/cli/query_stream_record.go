package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

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

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryStreamRecordsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.StreamRecords(context.Background(), params)
			if err != nil {
				return err
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

			queryClient := types.NewQueryClient(clientCtx)

			argAccount := args[0]

			params := &types.QueryGetStreamRecordRequest{
				Account: argAccount,
			}

			res, err := queryClient.StreamRecord(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
