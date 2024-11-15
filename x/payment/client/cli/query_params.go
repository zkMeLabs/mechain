package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	evmostypes "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/evm/precompiles/payment"
	"github.com/evmos/evmos/v12/x/payment/types"
)

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "shows the parameters of the module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			contract, err := payment.NewIPayment(common.HexToAddress(evmostypes.PaymentAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.Params(&bind.CallOpts{})
			if err != nil {
				return err
			}
			withdrawTimeLockThreshold := sdktypes.NewIntFromBigInt(result.WithdrawTimeLockThreshold)
			res := &types.QueryParamsResponse{
				Params: types.Params{
					VersionedParams: types.VersionedParams{
						ReserveTime:      result.VersionedParams.ReserveTime,
						ValidatorTaxRate: sdktypes.NewDecFromBigInt(result.VersionedParams.ValidatorTaxRate),
					},
					PaymentAccountCountLimit:  result.PaymentAccountCountLimit,
					ForcedSettleTime:          result.ForcedSettleTime,
					MaxAutoSettleFlowCount:    result.MaxAutoSettleFlowCount,
					MaxAutoResumeFlowCount:    result.MaxAutoResumeFlowCount,
					FeeDenom:                  result.FeeDenom,
					WithdrawTimeLockThreshold: &withdrawTimeLockThreshold,
					WithdrawTimeLockDuration:  result.WithdrawTimeLockDuration,
				}}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
