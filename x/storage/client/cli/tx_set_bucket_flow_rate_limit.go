package cli

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	sdkclient "github.com/evmos/evmos/v12/sdk/client"
	"github.com/evmos/evmos/v12/sdk/keys"
	gnfdSdkTypes "github.com/evmos/evmos/v12/sdk/types"
	types2 "github.com/evmos/evmos/v12/types"
)

var _ = strconv.Itoa(0)

func CmdSetBucketFlowRateLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-bucket-flow-rate-limit [bucket-name] [payment-account] [bucket-owner] [flow-rate-limit] --privatekey xxx",
		Short: "set flow rate limit for a bucket",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argBucketName := args[0]
			argPaymentAcc := args[1]
			paymentAcc, err := sdk.AccAddressFromHexUnsafe(argPaymentAcc)
			if err != nil {
				return err
			}
			argBucketOwner := args[2]
			bucketOwner, err := sdk.AccAddressFromHexUnsafe(argBucketOwner)
			if err != nil {
				return err
			}
			argFlowRateLimit, ok := sdkmath.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("invalid flow-rate-limit: %s", args[3])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			km, err := keys.NewPrivateKeyManager(argPrivateKey)
			gnfdCli, err := sdkclient.NewMechainClient(clientCtx.NodeURI, clientCtx.EvmNodeURI, gnfdSdkTypes.ChainID, sdkclient.WithKeyManager(km))
			if err != nil {
				return err
			}
			nonce, err := gnfdCli.GetNonce(context.Background())
			if err != nil {
				return err
			}
			txOpts, err := sdkclient.CreateTxOpts(context.Background(), clientCtx.EvmClient, argPrivateKey, big.NewInt(gnfdSdkTypes.DefaultChainId), gnfdSdkTypes.DefaultGasLimit, nonce)
			if err != nil {
				// return fmt.Errorf("failed to create tx opts")
				return err
			}

			session, err := sdkclient.CreateStorageSession(clientCtx.EvmClient, *txOpts, types2.StorageAddress)
			if err != nil {
				// return fmt.Errorf("failed to create session")
				return err
			}

			txRsp, err := session.SetBucketFlowRateLimit(
				bucketOwner.String(),
				paymentAcc.String(),
				argBucketName,
				argFlowRateLimit.BigInt(),
			)
			if err != nil {
				// return fmt.Errorf("failed to set bucket flow rate limit")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to set bucket flow rate limit.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to set bucket flow rate limit")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
