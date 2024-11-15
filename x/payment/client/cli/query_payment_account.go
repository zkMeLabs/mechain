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

func CmdListPaymentAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-payment-account",
		Short: "list all payment-account",
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
			result, err := contract.PaymentAccounts(&bind.CallOpts{}, *ToPaymentPageReq(pageReq))
			if err != nil {
				return err
			}

			paymentAccounts := make([]types.PaymentAccount, 0)
			for _, paymentAccount := range result.PaymentAccounts {
				paymentAccounts = append(paymentAccounts, *ToPaymentAccount(&paymentAccount))
			}
			res := &types.QueryPaymentAccountsResponse{
				PaymentAccounts: paymentAccounts,
				Pagination:      ToPageResp(&result.PageResponse),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPaymentAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-payment-account [addr]",
		Short: "shows a payment-account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			argAddr := args[0]

			contract, err := payment.NewIPayment(common.HexToAddress(evmostypes.PaymentAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.PaymentAccount(&bind.CallOpts{}, argAddr)
			if err != nil {
				return err
			}
			res := &types.QueryPaymentAccountResponse{
				PaymentAccount: *ToPaymentAccount(&result),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
