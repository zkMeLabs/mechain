package cli

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	sdkclient "github.com/evmos/evmos/v12/sdk/client"
	"github.com/evmos/evmos/v12/sdk/keys"
	gnfdSdkTypes "github.com/evmos/evmos/v12/sdk/types"
	types2 "github.com/evmos/evmos/v12/types"
)

var _ = strconv.Itoa(0)

func CmdCreatePaymentAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-payment-account --privatekey xxx",
		Short: "Broadcast message create-payment-account",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
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

			session, err := sdkclient.CreatePaymentSession(clientCtx.EvmClient, *txOpts, types2.PaymentAddress)
			if err != nil {
				// return fmt.Errorf("failed to create session")
				return err
			}

			txRsp, err := session.CreatePaymentAccount(
			// clientCtx.GetFromAddress().String(),
			)
			if err != nil {
				// return fmt.Errorf("failed to create payment account")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to create payment account", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to create payment account")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
