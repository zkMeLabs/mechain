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

func CmdListAutoSettleRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-auto-settle-record",
		Short: "list all auto-settle-record",
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
			result, err := contract.AutoSettleRecords(&bind.CallOpts{}, *ToPaymentPageReq(pageReq))
			if err != nil {
				return err
			}

			autoSettleRecords := make([]types.AutoSettleRecord, 0)
			for _, autoSettleRecord := range result.AutoSettleRecords {
				autoSettleRecords = append(autoSettleRecords, *ToAutoSettleRecord(&autoSettleRecord))
			}
			res := &types.QueryAutoSettleRecordsResponse{
				AutoSettleRecords: autoSettleRecords,
				Pagination:        ToPageResp(&result.PageResponse),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
