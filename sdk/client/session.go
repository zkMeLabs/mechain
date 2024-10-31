package client

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	gnfdSdkTypes "github.com/evmos/evmos/v12/sdk/types"
	"github.com/evmos/evmos/v12/x/evm/precompiles/payment"
	"github.com/evmos/evmos/v12/x/evm/precompiles/storage"
)

func CreateTxOpts(ctx context.Context, client *ethclient.Client, hexPrivateKey string, chain *big.Int, gasLimit uint64, nonce uint64) (*bind.TransactOpts, error) {
	// create private key
	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		return nil, err
	}

	// Build transact tx opts with private key
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chain)
	if err != nil {
		return nil, err
	}

	// set gas limit and gas price
	txOpts.GasLimit = gasLimit
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	txOpts.GasPrice = gasPrice

	txOpts.Nonce = big.NewInt(int64(nonce))

	return txOpts, nil
}

func CreateStorageSession(client *ethclient.Client, txOpts bind.TransactOpts, contractAddress string) (*storage.IStorageSession, error) {
	storageContract, err := storage.NewIStorage(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}
	session := &storage.IStorageSession{
		Contract: storageContract,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: txOpts,
	}
	return session, nil
}

func CreatePaymentSession(client *ethclient.Client, txOpts bind.TransactOpts, contractAddress string) (*payment.IPaymentSession, error) {
	paymentContract, err := payment.NewIPayment(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}
	session := &payment.IPaymentSession{
		Contract: paymentContract,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: txOpts,
	}
	return session, nil
}

// GetLatestBlockHeight - Get the height of the latest block from the chain.
//
// - ctx: Context variables for the current API call.
//
// - ret1: The block height.
//
// - ret2: Return error when the request failed, otherwise return nil.
func GetLatestBlockHeight(ctx context.Context, chainClient *MechainClient) (int64, error) {
	resp, err := chainClient.GetStatus(ctx)
	if err != nil {
		return 0, nil
	}
	return resp.SyncInfo.LatestBlockHeight, nil
}

// WaitForNextBlock - Wait until the next block is committed since current block.
//
// - ctx: Context variables for the current API call.
//
// - ret: Return error when the request failed, otherwise return nil.
func WaitForNextBlock(ctx context.Context, chainClient *MechainClient) error {
	res, err := chainClient.GetBlock(ctx, nil)
	if err != nil {
		return err
	}
	height := res.Block.Header.Height + 1
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		latestBlockHeight, err := GetLatestBlockHeight(ctx, chainClient)
		if err != nil {
			return err
		}
		if latestBlockHeight >= height {
			return nil
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf(err.Error(), "timeout exceeded waiting for block")
		case <-ticker.C:
		}
	}
}

// WaitForEvmTx - Wait for a evm transaction to be confirmed onchian, if transaction not found in current block, wait for the next block. API ends when a transaction is found or context is canceled.
//
// - ctx: Context variables for the current API call.
//
// - hash: The hex representation of transaction hash.
//
// - ret1: The transaction result details.
//
// - ret2: Return error when the request failed, otherwise return nil.
func WaitForEvmTx(ctx context.Context, evmClient *ethclient.Client, gnfdCli *MechainClient, hash common.Hash) (*ethtypes.Receipt, error) {
	for {
		txReceipt, err := evmClient.TransactionReceipt(ctx, hash)
		if err != nil {
			// Tx not found, wait for next block and try again
			if strings.Contains(err.Error(), "not found") {

				err := WaitForNextBlock(ctx, gnfdCli)
				if err != nil {
					return nil, fmt.Errorf("waiting for next block")
				}
				continue
			}
			return nil, fmt.Errorf("fetching tx '%s'", hash)
		}
		// `nil` could mean the transaction is in the mempool, invalidated, or was not sent in the first place.
		if txReceipt == nil {
			err := WaitForNextBlock(ctx, gnfdCli)
			if err != nil {
				return nil, fmt.Errorf("waiting for next block")
			}
			continue
		}
		if txReceipt.Status != gnfdSdkTypes.ReceiptStatusSuccessful {
			return nil, fmt.Errorf("txn: %s has failed with response code: %+v\n", hash.String(), txReceipt)
		}
		return txReceipt, nil
	}
}
