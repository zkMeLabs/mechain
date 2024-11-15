package cli

import (
	"strconv"

	cmath "cosmossdk.io/math"
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

func CmdDynamicBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dynamic-balance [account]",
		Short: "Query dynamic-balance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAccount := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			contract, err := payment.NewIPayment(common.HexToAddress(evmostypes.PaymentAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.DynamicBalance(&bind.CallOpts{}, reqAccount)
			if err != nil {
				return err
			}
			res := &types.QueryDynamicBalanceResponse{
				DynamicBalance:   cmath.NewIntFromBigInt(result.DynamicBalance),
				StreamRecord:     *ToStreamRecord(&result.StreamRecord),
				CurrentTimestamp: result.CurrentTimestamp,
				BankBalance:      cmath.NewIntFromBigInt(result.BankBalance),
				AvailableBalance: cmath.NewIntFromBigInt(result.AvailableBalance),
				LockedFee:        cmath.NewIntFromBigInt(result.LockedFee),
				ChangeRate:       cmath.NewIntFromBigInt(result.ChangeRate),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
