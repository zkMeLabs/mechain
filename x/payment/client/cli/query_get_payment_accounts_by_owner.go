package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	evmostypes "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/evm/precompiles/payment"

	"github.com/evmos/evmos/v12/x/payment/types"
)

var _ = strconv.Itoa(0)

func CmdGetPaymentAccountsByOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-payment-accounts-by-owner [owner]",
		Short: "Query get-payment-accounts-by-owner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqOwner := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			contract, err := payment.NewIPayment(common.HexToAddress(evmostypes.PaymentAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.PaymentAccountsByOwner(&bind.CallOpts{}, reqOwner)
			if err != nil {
				return err
			}
			res := &types.QueryPaymentAccountsByOwnerResponse{
				PaymentAccounts: result,
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
