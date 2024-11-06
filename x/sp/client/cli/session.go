package cli

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	spp "github.com/evmos/evmos/v12/x/evm/precompiles/storageprovider"
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

func CreateSpSession(client *ethclient.Client, txOpts bind.TransactOpts, contractAddress string) (*spp.IStorageProviderSession, error) {
	spContract, err := spp.NewIStorageProvider(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}
	session := &spp.IStorageProviderSession{
		Contract: spContract,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: txOpts,
	}
	return session, nil
}
