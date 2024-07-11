package storage

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
	"github.com/stretchr/testify/suite"
)

type bucketSuite struct {
	suite.Suite
	cli *ethclient.Client
	c   *IStorage
}

func TestBucketSuite(t *testing.T) {
	suite.Run(t, new(bucketSuite))
}

func (s *bucketSuite) SetupTest() {
	var err error
	s.cli, err = ethclient.Dial(testRPC)
	s.NoErrorf(err, "failed to connect to the rpc server: %v", err)
	s.c, err = NewIStorage(storageAddress, s.cli)
	s.NoErrorf(err, "failed to create a new contract: %v", err)
}

func (s *bucketSuite) TearDownTest() {
	s.cli.Close()
}

func (s *bucketSuite) TestE2E() {
	ctx := context.Background()
	chainID, err := s.cli.ChainID(ctx)
	s.NoError(err, "failed to get chain id")
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	s.NoError(err, "failed to create a new transactor")
	opts.GasLimit = 1000000
	// create bucket
	bucketName := fmt.Sprintf("bucket-%s", randStr(5))
	approval := Approval{
		ExpiredHeight:              0,
		GlobalVirtualGroupFamilyId: 1,
		Sig:                        []byte("0x00"),
	}
	chargedReadQuota := uint64(100000000000000)
	createTx, err := s.c.CreateBucket(opts, bucketName, 0, common.Address{}, common.Address{}, approval, chargedReadQuota)
	s.NoError(err, "failed to create bucket")
	_, err = WaitReceipt(ctx, s.cli, createTx)
	s.NoError(err, "can not get create receipt")

	// list buckets
	listRet, err := s.c.ListBuckets(nil, PageRequest{Limit: 10})
	s.NoError(err, "failed to list buckets")
	s.Equal(len(listRet.BucketInfos), 1)
	s.Equal(listRet.BucketInfos[0].BucketName, bucketName)
	s.Equal(listRet.BucketInfos[0].PaymentAddress.Hex(), testAddressHex)

	// update bucket info
	s.T().Log(listRet.BucketInfos[0].Visibility)
	newVisibility := storagetypes.VISIBILITY_TYPE_PUBLIC_READ
	if listRet.BucketInfos[0].Visibility == uint8(storagetypes.VISIBILITY_TYPE_PUBLIC_READ) {
		newVisibility = storagetypes.VISIBILITY_TYPE_PRIVATE
	}

	ChargedReadQuota := big.NewInt(-1)
	updateTx, err := s.c.UpdateBucketInfo(opts, bucketName, uint8(newVisibility), paymentAddress, ChargedReadQuota)
	s.NoError(err, "failed to update bucket info")
	_, err = WaitReceipt(ctx, s.cli, updateTx)
	s.NoError(err, "can not get update receipt")
	// head bucket
	headRet, err := s.c.HeadBucket(nil, bucketName)
	s.NoErrorf(err, "failed to head bucket: %v", err)
	s.Equal(uint8(newVisibility), headRet.BucketInfo.Visibility)
}

var waitReceiptInterval, _ = time.ParseDuration("10s")

func WaitReceipt(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	ticker := time.NewTicker(waitReceiptInterval)
	ctxWithTimeout, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
		ticker.Stop()
	}()

	for {
		select {
		case <-ctxWithTimeout.Done():
			return nil, ctxWithTimeout.Err()
		case <-ticker.C:
			receipt, err := client.TransactionReceipt(ctxWithTimeout, tx.Hash())
			if err != nil {
				continue
			}

			if receipt.Status != types.ReceiptStatusSuccessful {
				return nil, fmt.Errorf("transaction reverted, hash: %s", tx.Hash())
			}

			return receipt, nil
		}
	}
}
