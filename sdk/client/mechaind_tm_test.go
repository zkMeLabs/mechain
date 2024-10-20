package client

import (
	"context"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/assert"

	"github.com/evmos/evmos/v12/sdk/client/test"
	"github.com/evmos/evmos/v12/sdk/keys"
)

func TestTmClient(t *testing.T) {
	km, err := keys.NewPrivateKeyManager(test.TestPrivateKey)
	assert.NoError(t, err)
	gnfdCli, err := NewMechainClient(test.TestRPCAddr, test.TestChainID, WithKeyManager(km), WithWebSocketClient())
	assert.NoError(t, err)
	to, err := sdk.AccAddressFromHexUnsafe(test.TestAddr)
	assert.NoError(t, err)
	transfer := banktypes.NewMsgSend(km.GetAddr(), to, sdk.NewCoins(sdk.NewInt64Coin(test.TestTokenName, 12)))
	response, err := gnfdCli.BroadcastTx(context.Background(), []sdk.Msg{transfer}, nil)
	assert.NoError(t, err)
	time.Sleep(2 * time.Second)
	// get tx
	res, err := gnfdCli.Tx(context.Background(), response.TxResponse.TxHash)
	assert.NoError(t, err)
	t.Log(res.Hash)
	t.Log(res.TxResult.String())

	// get the latest block
	block, err := gnfdCli.GetBlock(context.Background(), nil)
	assert.NoError(t, err)
	t.Log(block)

	h := block.Block.Height

	block, err = gnfdCli.GetBlock(context.Background(), &h)
	assert.NoError(t, err)
	t.Log(block)

	// get block result
	blockResult, err := gnfdCli.GetBlockResults(context.Background(), &h)
	assert.NoError(t, err)
	t.Log(blockResult)

	// get validator
	validators, err := gnfdCli.GetValidators(context.Background(), &h)
	assert.NoError(t, err)
	t.Log(validators.Validators)
}
