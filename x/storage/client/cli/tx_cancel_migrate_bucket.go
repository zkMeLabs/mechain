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
	"github.com/evmos/evmos/v12/x/storage/types"
)

var _ = strconv.Itoa(0)

func CmdCancelMigrateBucket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-migrate-bucket [bucket-name] --privatekey xxx",
		Short: "cancel a existing bucket migration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argBucketName := args[0]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			km, err := keys.NewPrivateKeyManager(argPrivateKey)
			gnfdCli, err := sdkclient.NewMechainClient(clientCtx.NodeURI, gnfdSdkTypes.ChainID, sdkclient.WithKeyManager(km))
			if err != nil {
				return err
			}
			msg := types.NewMsgCancelMigrateBucket(
				clientCtx.GetFromAddress(),
				argBucketName,
			)
			if err = msg.ValidateBasic(); err != nil {
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

			session, err := sdkclient.CreateStorageSession(evmClient, *txOpts, types2.StorageAddress)
			if err != nil {
				// return fmt.Errorf("failed to create session")
				return err
			}

			txRsp, err := session.CancelMigrateBucket(argBucketName)
			if err != nil {
				// return fmt.Errorf("failed to cancel migrate bucket")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), evmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to cancel migrate bucket", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
