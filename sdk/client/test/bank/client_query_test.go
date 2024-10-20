package bank

import (
	"context"
	"testing"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/assert"

	gnfdclient "github.com/evmos/evmos/v12/sdk/client"
	"github.com/evmos/evmos/v12/sdk/client/test"
	"github.com/evmos/evmos/v12/sdk/keys"
)

func TestBankBalance(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := banktypes.QueryBalanceRequest{
		Address: test.TestAddr,
		Denom:   "azkme",
	}
	res, err := client.BankQueryClient.Balance(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.Balance.String())
}

func TestBankAllBalances(t *testing.T) {
	km, err := keys.NewPrivateKeyManager("e3ac46e277677f0f103774019d03bd89c7b4b5ecc554b2650bd5d5127992c20c")
	assert.NoError(t, err)
	t.Log(km)
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID, gnfdclient.WithWebSocketClient())
	assert.NoError(t, err)

	query := banktypes.QueryAllBalancesRequest{
		Address: test.TestAddr,
	}
	res, err := client.BankQueryClient.AllBalances(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.Balances.String())
}

func TestBankDenomMetadata(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := banktypes.QueryDenomMetadataRequest{}
	res, err := client.BankQueryClient.DenomMetadata(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.String())
}

func TestBankDenomOwners(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := banktypes.QueryDenomOwnersRequest{
		Denom: "azkme",
	}
	res, err := client.BankQueryClient.DenomOwners(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.String())
}

func TestBankDenomsMetadata(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := banktypes.QueryDenomsMetadataRequest{}
	res, err := client.BankQueryClient.DenomsMetadata(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.String())
}

func TestBankParams(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := banktypes.QueryParamsRequest{}
	res, err := client.BankQueryClient.Params(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.String())
}

func TestBankSpendableBalance(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := banktypes.QuerySpendableBalancesRequest{
		Address: test.TestAddr,
	}
	res, err := client.BankQueryClient.SpendableBalances(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.GetBalances().String())
}

func TestBankSupplyOf(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := banktypes.QuerySupplyOfRequest{
		Denom: "azkme",
	}
	res, err := client.BankQueryClient.SupplyOf(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.String())
}

func TestBankTotalSupply(t *testing.T) {
	client, err := gnfdclient.NewMechainClient(test.TestRPCAddr, test.TestChainID)
	assert.NoError(t, err)

	query := banktypes.QueryTotalSupplyRequest{}
	res, err := client.BankQueryClient.TotalSupply(context.Background(), &query)
	assert.NoError(t, err)

	t.Log(res.Supply.String())
}
