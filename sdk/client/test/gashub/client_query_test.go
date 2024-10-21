package gashub

import (
	"context"
	"testing"

	gashubtypes "github.com/cosmos/cosmos-sdk/x/gashub/types"
	"github.com/stretchr/testify/assert"

	gnfdclient "github.com/evmos/evmos/v12/sdk/client"
	"github.com/evmos/evmos/v12/sdk/client/test"
)

func TestGashubParams(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := gashubtypes.QueryParamsRequest{}
	res, err := client.GashubQueryClient.Params(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.String())
}
