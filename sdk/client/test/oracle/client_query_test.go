package bank

import (
	"context"
	"testing"

	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	"github.com/stretchr/testify/assert"

	gnfdclient "github.com/evmos/evmos/v12/sdk/client"
	"github.com/evmos/evmos/v12/sdk/client/test"
)

func TestOracleParams(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := oracletypes.QueryParamsRequest{}
	res, err := client.OracleQueryClient.Params(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.GetParams())
}
