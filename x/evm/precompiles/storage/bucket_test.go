package storage

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func Test_bucket(t *testing.T) {
	assert := assert.New(t)
	conn, err := ethclient.Dial(rpc)
	assert.NoErrorf(err, "failed to connect to the rpc server: %v", err)
	contract, err := NewIStorage(storageAddress, conn)
	assert.NoErrorf(err, "failed to create a new contract: %v", err)
	// create bucket
	{
	}
	// list buckets
	{
		ret, err := contract.ListBuckets(nil, PageRequest{Limit: 10})
		assert.NoError(err, "failed to list buckets")
		assert.Equal(len(ret.BucketInfos), 1)
		assert.Equal(ret.BucketInfos[0].BucketName, bucketName)
	}
	// get bucket
}
