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

func CmdListPaymentAccountCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-payment-account-count",
		Short: "list all payment-account-count",
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
			result, err := contract.PaymentAccountCounts(&bind.CallOpts{}, *ToPaymentPageReq(pageReq))
			if err != nil {
				return err
			}

			paymentAccountCounts := make([]types.PaymentAccountCount, 0)
			for _, paymentAccountCount := range result.PaymentAccountCounts {
				paymentAccountCounts = append(paymentAccountCounts, *ToPaymentAccountCount(&paymentAccountCount))
			}
			res := &types.QueryPaymentAccountCountsResponse{
				PaymentAccountCounts: paymentAccountCounts,
				Pagination:           ToPageResp(&result.PageResponse),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPaymentAccountCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-payment-account-count [owner]",
		Short: "shows a payment-account-count",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			argOwner := args[0]
			contract, err := payment.NewIPayment(common.HexToAddress(evmostypes.PaymentAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.PaymentAccountCount(&bind.CallOpts{}, argOwner)
			if err != nil {
				return err
			}

			res := &types.QueryPaymentAccountCountResponse{
				PaymentAccountCount: *ToPaymentAccountCount(&result),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
