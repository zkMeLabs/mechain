package bank

import (
	"context"
	"testing"

	crosschaintypes "github.com/cosmos/cosmos-sdk/x/crosschain/types"
	"github.com/stretchr/testify/assert"

	gnfdclient "github.com/evmos/evmos/v12/sdk/client"
	"github.com/evmos/evmos/v12/sdk/client/test"
)

func TestCrosschainParams(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := crosschaintypes.QueryParamsRequest{}
	res, err := client.CrosschainQueryClient.Params(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res)
}

func TestCrosschainPackageRequest(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := crosschaintypes.QueryCrossChainPackageRequest{}
	res, err := client.CrosschainQueryClient.CrossChainPackage(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.String())
}

func TestCrosschainReceiveSequence(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := crosschaintypes.QueryReceiveSequenceRequest{}
	res, err := client.CrosschainQueryClient.ReceiveSequence(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.String())
}

func TestCrosschainSendSequence(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := crosschaintypes.QuerySendSequenceRequest{}
	res, err := client.CrosschainQueryClient.SendSequence(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.String())
}
