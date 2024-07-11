package storage

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func Test_object(t *testing.T) {
	assert := assert.New(t)
	conn, err := ethclient.Dial(rpc)
	assert.NoErrorf(err, "failed to connect to the rpc server: %v", err)
	contract, err := NewIStorage(storageAddress, conn)
	assert.NoErrorf(err, "failed to create a new contract: %v", err)
	// create bucket
	{
	}
	// create object
	{
	}
	// list objects
	{
		pr := PageRequest{Limit: 10}
		ret, err := contract.ListObjects(nil, pr, bucketName)
		assert.NoError(err, "failed to list objects")
		assert.Equal(len(ret.ObjectInfos), 1)
		assert.Equal(ret.ObjectInfos[0].BucketName, bucketName)
		assert.Equal(ret.ObjectInfos[0].ObjectName, objectName)
	}
	// head object
	{
		ret, err := contract.HeadObject(nil, bucketName, objectName)
		assert.NoErrorf(err, "failed to head object: %v", err)
		assert.Equal(bucketName, ret.ObjectInfo.BucketName)
		assert.Equal(objectName, ret.ObjectInfo.ObjectName)
	}
}
