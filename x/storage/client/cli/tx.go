package cli

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	cmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	sdkclient "github.com/evmos/evmos/v12/sdk/client"
	"github.com/evmos/evmos/v12/sdk/keys"
	gnfdSdkTypes "github.com/evmos/evmos/v12/sdk/types"
	types2 "github.com/evmos/evmos/v12/types"
	gnfderrors "github.com/evmos/evmos/v12/types/errors"
	"github.com/evmos/evmos/v12/x/evm/precompiles/storage"
	"github.com/evmos/evmos/v12/x/storage/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdCreateBucket(),
		CmdDeleteBucket(),
		CmdUpdateBucketInfo(),
		CmdMirrorBucket(),
		CmdDiscontinueBucket(),
		CmdMigrateBucket(),
		CmdCancelMigrateBucket(),
		CmdSetBucketFlowRateLimit(),
	)

	cmd.AddCommand(
		CmdCreateObject(),
		CmdDeleteObject(),
		CmdCancelCreateObject(),
		CmdCopyObject(),
		CmdMirrorObject(),
		CmdDiscontinueObject(),
		CmdUpdateObjectInfo(),
	)

	cmd.AddCommand(
		CmdCreateGroup(),
		CmdDeleteGroup(),
		CmdUpdateGroupMember(),
		CmdUpdateGroupExtra(),
		CmdRenewGroupMember(),
		CmdLeaveGroup(),
		CmdMirrorGroup(),
	)

	cmd.AddCommand(
		CmdPutPolicy(),
		CmdDeletePolicy(),
	)

	cmd.AddCommand(
		CmdSetTag(),
	)

	return cmd
}

// CmdCreateBucket returns a CLI command handler for creating a MsgCreateBucket transaction.
func CmdCreateBucket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-bucket [bucket-name] --privatekey xxx --gvgfamily-id x",
		Short: "create a new bucket which associate to a primary sp",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argGlobalVirtualGroupFamilyId, _ := cmd.Flags().GetUint32(FlagGVGFamilyID)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argBucketName := args[0]

			visibility, err := cmd.Flags().GetString(FlagVisibility)
			if err != nil {
				return err
			}
			visibilityType, err := GetVisibilityType(visibility)
			if err != nil {
				return err
			}

			chargedReadQuota, err := cmd.Flags().GetUint64(FlagChargedReadQuota)
			if err != nil {
				return err
			}

			payment, _ := cmd.Flags().GetString(FlagPaymentAccount)
			paymentAcc, _, _, err := GetPaymentAccountField(clientCtx.Keyring, payment)
			if err != nil {
				return err
			}

			primarySP, _ := cmd.Flags().GetString(FlagPrimarySP)
			primarySPAcc, _, _, err := GetPrimarySPField(clientCtx.Keyring, primarySP)
			if err != nil {
				return err
			}

			approveSignature, _ := cmd.Flags().GetString(FlagApproveSignature)
			approveTimeoutHeight, _ := cmd.Flags().GetUint64(FlagApproveTimeoutHeight)

			tagsStr, _ := cmd.Flags().GetString(FlagTags)
			tags := GetTags(tagsStr)

			approveSignatureBytes, err := hex.DecodeString(approveSignature)
			if err != nil {
				return err
			}
			km, err := keys.NewPrivateKeyManager(argPrivateKey)
			gnfdCli, err := sdkclient.NewMechainClient(clientCtx.NodeURI, clientCtx.EvmNodeURI, gnfdSdkTypes.ChainID, sdkclient.WithKeyManager(km))
			if err != nil {
				return err
			}
			msgCreateBucket := types.NewMsgCreateBucket(
				km.GetAddr(), // clientCtx.GetFromAddress(),
				argBucketName,
				visibilityType,
				primarySPAcc,
				paymentAcc,
				approveTimeoutHeight,
				approveSignatureBytes,
				chargedReadQuota,
			)
			if err := msgCreateBucket.ValidateBasic(); err != nil {
				return err
			}

			if tags != nil {
				grn := types2.NewBucketGRN(argBucketName)
				msgSetTag := types.NewMsgSetTag(clientCtx.GetFromAddress(), grn.String(), tags)
				if err := msgSetTag.ValidateBasic(); err != nil {
					return err
				}
				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCreateBucket, msgSetTag)
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

			txRsp, err := session.CreateBucket(
				argBucketName,
				uint8(visibilityType),
				ethcmn.Address(paymentAcc),
				ethcmn.Address(primarySPAcc),
				storage.Approval{
					ExpiredHeight:              approveTimeoutHeight,
					GlobalVirtualGroupFamilyId: argGlobalVirtualGroupFamilyId,
					Sig:                        approveSignatureBytes,
				},
				chargedReadQuota,
			)
			if err != nil {
				// return fmt.Errorf("failed to create bucket")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to create bucket.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().AddFlagSet(FlagSetVisibility())
	cmd.Flags().AddFlagSet(FlagSetApproval())
	cmd.Flags().Uint64(FlagChargedReadQuota, 0, "The charged read quota of bucket.")
	cmd.Flags().String(FlagPaymentAccount, "", "The address of the account used to pay for the read fee. The default is the sender account.")
	cmd.Flags().String(FlagPrimarySP, "", "The operator account address of primarySp")
	cmd.Flags().String(FlagTags, "", "The tags of the resource. It should be like: `key1=value1,key2=value2`")
	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to create bucket")
	cmd.Flags().Uint32(FlagGVGFamilyID, 1, "The GlobalVirtualGroupFamilyId of bucket.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteBucket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-bucket [bucket-name] --privatekey xxx",
		Short: "delete an existing bucket",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argBucketName := args[0]

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

			txRsp, err := session.DeleteBucket(argBucketName)
			if err != nil {
				// return fmt.Errorf("failed to delete bucket")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to delete bucket.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to delete bucket")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateBucketInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-bucket-info [bucket-name] [charged-read-quota] --privatekey xxx",
		Short: "Update the meta of bucket, E.g ChargedReadQuota, PaymentAccount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argBucketName := args[0]
			argChargedReadQuota, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			visibility, err := cmd.Flags().GetString(FlagVisibility)
			if err != nil {
				return err
			}
			visibilityType, err := GetVisibilityType(visibility)
			if err != nil {
				return err
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

			txRsp, err := session.UpdateBucketInfo(
				argBucketName,
				uint8(visibilityType),
				ethcmn.Address{},
				new(big.Int).SetUint64(argChargedReadQuota),
			)
			if err != nil {
				// return fmt.Errorf("failed to update bucket info")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("update bucket info.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to update bucket info")
	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().AddFlagSet(FlagSetVisibility())

	return cmd
}

func CmdCancelCreateObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-create-object [bucket-name] [object-name] --privatekey xxx",
		Short: "Broadcast message cancel_create_object",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argBucketName := args[0]
			argObjectName := args[1]

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

			txRsp, err := session.CancelCreateObject(
				argBucketName,
				argObjectName,
			)
			if err != nil {
				// return fmt.Errorf("failed to cancel create object")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to cancel create object.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to cancel create object")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdCreateObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-object [bucket-name] [object-name] [payload-size] [content-type] --privatekey xxx",
		Short: "Create a new object in the bucket, checksums split by ','",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argBucketName := args[0]
			argObjectName := args[1]
			argPayloadSize := args[2]
			argContentType := args[3]

			payloadSize, err := strconv.ParseUint(argPayloadSize, 10, 64)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			visibility, err := cmd.Flags().GetString(FlagVisibility)
			if err != nil {
				return err
			}
			visibilityType, err := GetVisibilityType(visibility)
			if err != nil {
				return err
			}

			checksums, _ := cmd.Flags().GetString(FlagExpectChecksums)
			redundancyTypeFlag, _ := cmd.Flags().GetString(FlagRedundancyType)
			approveSignature, _ := cmd.Flags().GetString(FlagApproveSignature)
			approveTimeoutHeight, _ := cmd.Flags().GetUint64(FlagApproveTimeoutHeight)
			tagsStr, _ := cmd.Flags().GetString(FlagTags)
			tags := GetTags(tagsStr)

			approveSignatureBytes, err := hex.DecodeString(approveSignature)
			if err != nil {
				return err
			}

			checksumsStr := strings.Split(checksums, ",")
			if checksumsStr == nil {
				return gnfderrors.ErrInvalidChecksum
			}
			var expectChecksums [][]byte
			for _, checksum := range checksumsStr {
				tmp, err := hex.DecodeString(checksum)
				if err != nil {
					return err
				}
				expectChecksums = append(expectChecksums, tmp)
			}

			var redundancyType types.RedundancyType
			switch redundancyTypeFlag {
			case "EC":
				redundancyType = types.REDUNDANCY_EC_TYPE
			case "Replica":
				redundancyType = types.REDUNDANCY_REPLICA_TYPE
			default:
				return types.ErrInvalidRedundancyType
			}

			km, err := keys.NewPrivateKeyManager(argPrivateKey)
			gnfdCli, err := sdkclient.NewMechainClient(clientCtx.NodeURI, clientCtx.EvmNodeURI, gnfdSdkTypes.ChainID, sdkclient.WithKeyManager(km))
			if err != nil {
				return err
			}
			msgCreateObject := types.NewMsgCreateObject(
				km.GetAddr(), // clientCtx.GetFromAddress(),
				argBucketName,
				argObjectName,
				payloadSize,
				visibilityType,
				expectChecksums,
				argContentType,
				redundancyType,
				approveTimeoutHeight,
				approveSignatureBytes,
			)
			if err := msgCreateObject.ValidateBasic(); err != nil {
				return err
			}

			if tags != nil {
				grn := types2.NewObjectGRN(argBucketName, argObjectName)
				msgSetTag := types.NewMsgSetTag(clientCtx.GetFromAddress(), grn.String(), tags)
				if err := msgSetTag.ValidateBasic(); err != nil {
					return err
				}
				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCreateObject, msgSetTag)
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

			txRsp, err := session.CreateObject(
				argBucketName,
				argObjectName,
				payloadSize,
				uint8(visibilityType),
				argContentType,
				storage.Approval{ExpiredHeight: approveTimeoutHeight, Sig: approveSignatureBytes},
				checksumsStr,
				uint8(redundancyType),
			)
			if err != nil {
				// return fmt.Errorf("failed to create object")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to create object.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().AddFlagSet(FlagSetVisibility())
	cmd.Flags().AddFlagSet(FlagSetApproval())
	cmd.Flags().String(FlagPrimarySP, "", "The operator account address of primarySp")
	cmd.Flags().String(FlagExpectChecksums, "", "The checksums that calculate by redundancy algorithm")
	cmd.Flags().String(FlagRedundancyType, "", "The redundancy type, EC or Replica ")
	cmd.Flags().String(FlagTags, "", "The tags of the resource. It should be like: `key1=value1,key2=value2`")
	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to create object")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdCopyObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy-object [src-bucket-name] [dst-bucket-name] [src-object-name] [dst-object-name] --privatekey xxx",
		Short: "Copy an existing object in a bucket to another bucket",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argSrcBucketName := args[0]
			argDstBucketName := args[1]
			argSrcObjectName := args[2]
			argDstObjectName := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			approveSignature, _ := cmd.Flags().GetString(FlagApproveSignature)
			approveTimeoutHeight, _ := cmd.Flags().GetUint64(FlagApproveTimeoutHeight)

			approveSignatureBytes, err := hex.DecodeString(approveSignature)
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

			txRsp, err := session.CopyObject(
				argSrcBucketName,
				argDstBucketName,
				argSrcObjectName,
				argDstObjectName,
				storage.Approval{ExpiredHeight: approveTimeoutHeight, Sig: approveSignatureBytes},
			)
			if err != nil {
				// return fmt.Errorf("failed to copy object")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to copy object.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to copy object")
	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().AddFlagSet(FlagSetApproval())

	return cmd
}

func CmdDeleteObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-object [bucket-name] [object-name] --privatekey xxx",
		Short: "Delete an existing object",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argBucketName := args[0]
			argObjectName := args[1]

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

			txRsp, err := session.DeleteObject(
				argBucketName,
				argObjectName,
			)
			if err != nil {
				// return fmt.Errorf("failed to delete object")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to delete object.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to delete object")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateObjectInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-object-info [bucket-name] [object-name] [flags] --privatekey xxx",
		Short: "Update the meta of object, Currently only support: Visibility",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argBucketName := args[0]
			argObjectName := args[1]

			visibility, err := cmd.Flags().GetString(FlagVisibility)
			if err != nil {
				return err
			}
			visibilityType, err := GetVisibilityType(visibility)
			if err != nil {
				return err
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

			txRsp, err := session.UpdateObjectInfo(
				argBucketName,
				argObjectName,
				uint8(visibilityType),
			)
			if err != nil {
				// return fmt.Errorf("failed to update object info")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to update object info.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to update object info")
	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().AddFlagSet(FlagSetVisibility())

	return cmd
}

func CmdDiscontinueObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "discontinue-object [bucket-name] [object-ids] [reason] --privatekey xxx",
		Short: "Discontinue to store objects",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argBucketName := args[0]
			argObjectIDs := args[1]
			argObjectReason := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			objectIDs := make([]cmath.Uint, 0)
			objectIds := make([]*big.Int, 0)
			splitIDs := strings.Split(argObjectIDs, ",")
			for _, split := range splitIDs {
				id, ok := big.NewInt(0).SetString(split, 10)
				if !ok {
					return fmt.Errorf("invalid object id: %s", id)
				}
				if id.Cmp(big.NewInt(0)) < 0 {
					return fmt.Errorf("object id should not be negative")
				}

				objectIDs = append(objectIDs, cmath.NewUintFromBigInt(id))
				objectIds = append(objectIds, id)
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

			txRsp, err := session.DiscontinueObject(
				argBucketName,
				objectIds,
				argObjectReason,
			)
			if err != nil {
				// return fmt.Errorf("failed to discontinue object")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to discontinue object.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to discontinue object")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdCreateGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-group [group-name] --privatekey xxx",
		Short: "Create a new group without group members",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argGroupName := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			extra, _ := cmd.Flags().GetString(FlagExtra)
			tagsStr, _ := cmd.Flags().GetString(FlagTags)
			tags := GetTags(tagsStr)

			km, err := keys.NewPrivateKeyManager(argPrivateKey)
			gnfdCli, err := sdkclient.NewMechainClient(clientCtx.NodeURI, clientCtx.EvmNodeURI, gnfdSdkTypes.ChainID, sdkclient.WithKeyManager(km))
			if err != nil {
				return err
			}
			msgCreateGroup := types.NewMsgCreateGroup(
				km.GetAddr(), // clientCtx.GetFromAddress(),
				argGroupName,
				extra,
			)
			if err := msgCreateGroup.ValidateBasic(); err != nil {
				return err
			}

			if tags != nil {
				grn := types2.NewGroupGRN(clientCtx.GetFromAddress(), argGroupName)
				msgSetTag := types.NewMsgSetTag(clientCtx.GetFromAddress(), grn.String(), tags)
				if err := msgSetTag.ValidateBasic(); err != nil {
					return err
				}
				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCreateGroup, msgSetTag)
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

			txRsp, err := session.CreateGroup(
				argGroupName,
				extra,
			)
			if err != nil {
				// return fmt.Errorf("failed to create group")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to create group.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagExtra, "", "extra info for the group")
	cmd.Flags().String(FlagTags, "", "The tags of the resource. It should be like: `key1=value1,key2=value2`")
	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to create group")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-group [group-name] --privatekey xxx",
		Short: "Delete an existing group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argGroupName := args[0]

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

			txRsp, err := session.DeleteGroup(argGroupName)
			if err != nil {
				// return fmt.Errorf("failed to delete group")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to delete group.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to delete group")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdLeaveGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "leave-group [group-owner] [group-name] --privatekey xxx",
		Short: "Leave the group you're a member of",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argGroupOwner := args[0]
			argGroupName := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			groupOwner, err := sdk.AccAddressFromHexUnsafe(argGroupOwner)
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

			txRsp, err := session.LeaveGroup(
				ethcmn.Address(km.GetAddr()),
				ethcmn.Address(groupOwner),
				argGroupName,
			)
			if err != nil {
				// return fmt.Errorf("failed to leave group")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to leave group.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to leave group")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateGroupMember() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-group-member [group-name] [member-to-add] [member-expiration-to-add] [member-to-delete] --privatekey xxx",
		Short: "Update the member of the group you own, split member addresses and expiration(UNIX timestamp) by ,",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argGroupName := args[0]
			argMemberToAdd := args[1]
			argMemberExpirationToAdd := args[2]
			argMemberToDelete := args[3]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			membersToAdd := strings.Split(argMemberToAdd, ",")
			memberExpirationStr := strings.Split(argMemberExpirationToAdd, ",")
			if len(memberExpirationStr) != len(membersToAdd) {
				return errors.New("[member-to-add] and [member-expiration-to-add] should have the same length")
			}

			msgGroupMemberToAdd := make([]*types.MsgGroupMember, 0, len(argMemberToAdd))
			groupMembersToAdd := make([]ethcmn.Address, 0, len(argMemberToAdd))
			var expirationTime []int64
			if len(membersToAdd) > 0 {
				for i := range membersToAdd {
					if len(membersToAdd[i]) > 0 {
						memberToAdd, err := sdk.AccAddressFromHexUnsafe(membersToAdd[i])
						if err != nil {
							return err
						}
						groupMembersToAdd[i] = ethcmn.Address(memberToAdd)
						member := types.MsgGroupMember{
							Member: membersToAdd[i],
						}
						if len(memberExpirationStr[i]) > 0 {
							unix, err := strconv.ParseInt(memberExpirationStr[i], 10, 64)
							if err != nil {
								return err
							}
							expirationTime[i] = unix
							expiration := time.Unix(unix, 0)
							member.ExpirationTime = &expiration
						}

						msgGroupMemberToAdd = append(msgGroupMemberToAdd, &member)
					}
				}
			}

			var memberAddrsToDelete []sdk.AccAddress
			var groupMembersToDelete []ethcmn.Address
			if len(argMemberToDelete) != 0 {
				membersToDelete := strings.Split(argMemberToDelete, ",")
				for _, member := range membersToDelete {
					if len(member) > 0 {
						memberAddr, err := sdk.AccAddressFromHexUnsafe(member)
						if err != nil {
							return err
						}
						groupMembersToDelete = append(groupMembersToDelete, ethcmn.Address(memberAddr))
						memberAddrsToDelete = append(memberAddrsToDelete, memberAddr)
					}
				}
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

			txRsp, err := session.UpdateGroup(
				ethcmn.Address(km.GetAddr()),
				argGroupName,
				groupMembersToAdd,
				expirationTime,
				groupMembersToDelete,
			)
			if err != nil {
				// return fmt.Errorf("failed to update group member")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to update group member.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to update group member")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRenewGroupMember() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew-group-member [group-name] [member] [member-expiration] --privatekey xxx",
		Short: "renew the member of the group you own, split member-addresses and member-expiration(UNIX timestamp) by ,",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argGroupName := args[0]
			argMember := args[1]
			argMemberExpiration := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			memberExpirationStr := strings.Split(argMemberExpiration, ",")
			members := strings.Split(argMember, ",")

			if len(memberExpirationStr) != len(members) {
				return errors.New("member and member-expiration should have the same length")
			}

			msgGroupMember := make([]*types.MsgGroupMember, 0, len(argMember))
			var groupMembers []ethcmn.Address
			var expirationTime []int64
			for i := range members {
				if len(members[i]) > 0 {
					memberToRenew, err := sdk.AccAddressFromHexUnsafe(members[i])
					if err != nil {
						return err
					}
					groupMembers[i] = ethcmn.Address(memberToRenew)
					member := types.MsgGroupMember{
						Member: members[i],
					}
					if len(memberExpirationStr[i]) > 0 {
						unix, err := strconv.ParseInt(memberExpirationStr[i], 10, 64)
						if err != nil {
							return err
						}
						expirationTime[i] = unix
						expiration := time.Unix(unix, 0)
						member.ExpirationTime = &expiration
					}

					msgGroupMember = append(msgGroupMember, &member)
				}
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

			txRsp, err := session.RenewGroupMember(
				ethcmn.Address(km.GetAddr()),
				argGroupName,
				groupMembers,
				expirationTime,
			)
			if err != nil {
				// return fmt.Errorf("failed to renew group member")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to renew group member.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to renew group member")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateGroupExtra() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-group-extra [group-name] [extra] --privatekey xxx",
		Short: "Update the extra info of the group",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argGroupName := args[0]
			argExtra := args[1]

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

			txRsp, err := session.UpdateGroupExtra(
				ethcmn.Address(km.GetAddr()),
				argGroupName,
				argExtra,
			)
			if err != nil {
				// return fmt.Errorf("failed to update group extra")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to update group extra.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to update group extra")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdPutPolicy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put-policy [principle-value] [resource] --privatekey xxx",
		Short: "put a policy to bucket/object/group which can grant permission to others",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argPrincipalValue := args[0]
			argResource := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			principal, err := GetPrincipal(argPrincipalValue)
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

			txRsp, err := session.PutPolicy(
				// ethcmn.Address(km.GetAddr()),
				storage.Principal{PrincipalType: int32(principal.Type), Value: principal.Value},
				argResource,
				[]storage.Statement{},
				0,
			)
			if err != nil {
				// return fmt.Errorf("failed to put policy")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to put policy.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to put policy")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeletePolicy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-policy [principle-value] [resource] --privatekey xxx",
		Short: "Delete policy with specify principle",
		Args:  cobra.ExactArgs(2),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Delete the policy, the principle-value can be account or group id.

Example:
$ %s tx storage delete-policy 0xffffffffffffffffffffff
$ %s tx delete-policy 3
`,
				version.AppName, version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argPrincipalValue := args[0]
			argResource := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			principal, err := GetPrincipal(argPrincipalValue)
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

			txRsp, err := session.DeletePolicy(
				// ethcmn.Address(km.GetAddr()),
				storage.Principal{PrincipalType: int32(principal.Type), Value: principal.Value},
				argResource,
			)
			if err != nil {
				// return fmt.Errorf("failed to delete policy")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to delete policy.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to delete policy")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdMirrorBucket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mirror-bucket --privatekey xxx",
		Short: "Mirror an existing bucket to the destination chain",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argBucketID, _ := cmd.Flags().GetString(FlagBucketID)
			argBucketName, _ := cmd.Flags().GetString(FlagBucketName)
			argDestChainID, _ := cmd.Flags().GetString(FlagDestChainID)

			bucketID := big.NewInt(0)
			switch {
			case argBucketID == "" && argBucketName == "":
				return fmt.Errorf("bucket id or bucket name should be provided")
			case argBucketID != "" && argBucketName != "":
				return fmt.Errorf("bucket id and bucket name should not be provided together")
			case argBucketID != "":
				ok := false
				bucketID, ok = big.NewInt(0).SetString(argBucketID, 10)
				if !ok {
					return fmt.Errorf("invalid bucket id: %s", argBucketID)
				}
				if bucketID.Cmp(big.NewInt(0)) <= 0 {
					return fmt.Errorf("bucket id should be positive")
				}
			}

			if argDestChainID == "" {
				return fmt.Errorf("destination chain id should be provided")
			}
			destChainID, err := strconv.ParseUint(argDestChainID, 10, 16)
			if err != nil {
				return err
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

			txRsp, err := session.MirrorBucket(
				// ethcmn.Address(km.GetAddr()),
				bucketID,
				argBucketName,
				uint32(destChainID),
			)
			if err != nil {
				// return fmt.Errorf("failed to mirror bucket")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to mirror bucket.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagBucketID, "", "Id of the bucket to mirror")
	cmd.Flags().String(FlagBucketName, "", "Name of the bucket to mirror")
	cmd.Flags().String(FlagDestChainID, "", "the destination chain id")
	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to mirror bucket")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDiscontinueBucket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "discontinue-bucket [bucket-name] [reason] --privatekey xxx",
		Short: "Discontinue to store bucket",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argBucketName := args[0]
			argObjectReason := args[1]

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

			txRsp, err := session.DiscontinueBucket(
				argBucketName,
				argObjectReason,
			)
			if err != nil {
				// return fmt.Errorf("failed to discontinue bucket")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to discontinue bucket.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to discontinue bucket")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdMigrateBucket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate-bucket [bucket-name] [dest-primary-sp-id] --privatekey xxx",
		Short: "migrate a bucket to another primary storage provider by user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			bucketName := args[0]
			destPrimarySpID, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			approveSignature, _ := cmd.Flags().GetString(FlagApproveSignature)
			approveTimeoutHeight, _ := cmd.Flags().GetUint64(FlagApproveTimeoutHeight)
			approveSignatureBytes, err := hex.DecodeString(approveSignature)
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

			txRsp, err := session.MigrateBucket(
				// ethcmn.Address(km.GetAddr()),
				bucketName,
				uint32(destPrimarySpID),
				storage.Approval{ExpiredHeight: approveTimeoutHeight, Sig: approveSignatureBytes},
			)
			if err != nil {
				// return fmt.Errorf("failed to migrate bucket")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to migrate bucket.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to migrate bucket")
	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().AddFlagSet(FlagSetApproval())
	return cmd
}

func CmdMirrorObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mirror-object --privatekey xxx",
		Short: "Mirror the object to the destination chain",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argObjectID, _ := cmd.Flags().GetString(FlagObjectID)
			argBucketName, _ := cmd.Flags().GetString(FlagBucketName)
			argObjectName, _ := cmd.Flags().GetString(FlagObjectName)
			argDestChainID, _ := cmd.Flags().GetString(FlagDestChainID)

			objectID := big.NewInt(0)
			switch {
			case argObjectID == "" && argObjectName == "":
				return fmt.Errorf("object id or object name should be provided")
			case argObjectID != "" && argObjectName != "":
				return fmt.Errorf("object id and object name should not be provided together")
			case argObjectID != "":
				ok := false
				objectID, ok = big.NewInt(0).SetString(argObjectID, 10)
				if !ok {
					return fmt.Errorf("invalid object id: %s", argObjectID)
				}
				if objectID.Cmp(big.NewInt(0)) <= 0 {
					return fmt.Errorf("object id should be positive")
				}
			case argObjectName != "" && argBucketName == "":
				return fmt.Errorf("object name and bucket name should not be provided together")
			}

			if argDestChainID == "" {
				return fmt.Errorf("destination chain id should be provided")
			}
			destChainID, err := strconv.ParseUint(argDestChainID, 10, 16)
			if err != nil {
				return err
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

			txRsp, err := session.MirrorObject(
				// ethcmn.Address(km.GetAddr()),
				objectID,
				argBucketName,
				argObjectName,
				uint32(destChainID),
			)
			if err != nil {
				// return fmt.Errorf("failed to mirror object")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to mirror object.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagObjectID, "", "Id of the object to mirror")
	cmd.Flags().String(FlagObjectName, "", "Name of the object to mirror")
	cmd.Flags().String(FlagBucketName, "", "Name of the bucket that the object belongs to")
	cmd.Flags().String(FlagDestChainID, "", "the destination chain id")
	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to mirror object")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdMirrorGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mirror-group --privatekey xxx",
		Short: "Mirror an existing group to the destination chain",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argGroupID, _ := cmd.Flags().GetString(FlagGroupID)
			argGroupName, _ := cmd.Flags().GetString(FlagGroupName)
			argDestChainID, _ := cmd.Flags().GetString(FlagDestChainID)

			groupID := big.NewInt(0)
			switch {
			case argGroupID == "" && argGroupName == "":
				return fmt.Errorf("group id or group name should be provided")
			case argGroupID != "" && argGroupName != "":
				return fmt.Errorf("group id and group name should not be provided together")
			case argGroupID != "":
				ok := false
				groupID, ok = big.NewInt(0).SetString(argGroupID, 10)
				if !ok {
					return fmt.Errorf("invalid groupd id: %s", argGroupID)
				}
				if groupID.Cmp(big.NewInt(0)) <= 0 {
					return fmt.Errorf("groupd id should be positive")
				}
			}

			if argDestChainID == "" {
				return fmt.Errorf("destination chain id should be provided")
			}
			destChainID, err := strconv.ParseUint(argDestChainID, 10, 16)
			if err != nil {
				return err
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

			txRsp, err := session.MirrorGroup(
				// ethcmn.Address(km.GetAddr()),
				groupID,
				argGroupName,
				uint32(destChainID),
			)
			if err != nil {
				// return fmt.Errorf("failed to mirror group")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to mirror group.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagGroupID, "", "Id of the group to mirror")
	cmd.Flags().String(FlagGroupName, "", "Name of the group to mirror")
	cmd.Flags().String(FlagDestChainID, "", "the destination chain id")
	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to mirror group")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSetTag() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-tag [grn] --privatekey xxx",
		Short: "set a bucket/object/group's tag.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrivateKey, _ := cmd.Flags().GetString(FlagPrivateKey)
			argResource := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tagsStr, _ := cmd.Flags().GetString(FlagTags)
			resourceTags := GetTags(tagsStr)

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

			tags := make([]storage.Tag, 0)
			for _, tag := range resourceTags.Tags {
				tags = append(tags, storage.Tag{
					Key:   tag.Key,
					Value: tag.Value,
				})
			}
			txRsp, err := session.SetTag(
				// ethcmn.Address(km.GetAddr()),
				argResource,
				tags,
			)
			if err != nil {
				// return fmt.Errorf("failed to mirror group")
				return err
			}

			_, err = sdkclient.WaitForEvmTx(context.Background(), clientCtx.EvmClient, gnfdCli, txRsp.Hash())
			if err != nil {
				return fmt.Errorf("failed to set tags.%v", err.Error())
			}
			return clientCtx.PrintObjectLegacy(txRsp.Hash().String())
		},
	}

	cmd.Flags().String(FlagTags, "", "The tags of the resource. It should be like: `key1=value1,key2=value2`")
	cmd.Flags().String(FlagPrivateKey, "", "The privatekey of account to set tags")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
