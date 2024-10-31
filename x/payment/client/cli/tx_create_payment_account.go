package cli

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	sdkclient "github.com/evmos/evmos/v12/sdk/client"
	"github.com/evmos/evmos/v12/sdk/keys"
	gnfdSdkTypes "github.com/evmos/evmos/v12/sdk/types"
	types2 "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/payment/types"
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
			gnfdCli, err := sdkclient.NewMechainClient(clientCtx.NodeURI, gnfdSdkTypes.ChainID, sdkclient.WithKeyManager(km))
			if err != nil {
				return err
			}
			msg := types.NewMsgCreatePaymentAccount(
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			evmClient, err := ethclient.Dial(gnfdSdkTypes.EvmUrl)
			if err != nil {
				// return fmt.Errorf("failed to new a evm client")
				return err
			}
			nonce, err := gnfdCli.GetNonce(context.Background())
			if err != nil {
				return err
			}
			txOpts, err := sdkclient.CreateTxOpts(context.Background(), evmClient, argPrivateKey, big.NewInt(gnfdSdkTypes.DefaultChainId), gnfdSdkTypes.DefaultGasLimit, nonce)
			if err != nil {
				// return fmt.Errorf("failed to create tx opts")
				return err
			}

			session, err := sdkclient.CreatePaymentSession(evmClient, *txOpts, types2.PaymentAddress)
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

			_, err = sdkclient.WaitForEvmTx(context.Background(), evmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to create payment account", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
