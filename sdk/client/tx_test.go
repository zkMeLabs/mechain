package client

import (
	"context"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/evmos/evmos/v12/sdk/client/test"
	"github.com/evmos/evmos/v12/sdk/keys"
	"github.com/evmos/evmos/v12/sdk/types"
)

func TestSendTokenSucceedWithSimulatedGas(t *testing.T) {
	km, err := keys.NewPrivateKeyManager(test.TestPrivateKey)
	assert.NoError(t, err)
	gnfdCli, err := NewMechainClient(test.TestRPCAddr, test.TestChainID, WithKeyManager(km))
	assert.NoError(t, err)
	to, err := sdk.AccAddressFromHexUnsafe(test.TestAddr)
	assert.NoError(t, err)
	transfer := banktypes.NewMsgSend(km.GetAddr(), to, sdk.NewCoins(sdk.NewInt64Coin(test.TestTokenName, 12)))
	response, err := gnfdCli.BroadcastTx(context.Background(), []sdk.Msg{transfer}, nil)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), response.TxResponse.Code)
	t.Log(response.TxResponse.String())
}

func TestSendTokenWithTxOptionSucceed(t *testing.T) {
	km, err := keys.NewPrivateKeyManager(test.TestPrivateKey)
	assert.NoError(t, err)
	gnfdCli, err := NewMechainClient(test.TestRPCAddr, test.TestChainID, WithKeyManager(km))
	assert.NoError(t, err)
	to, err := sdk.AccAddressFromHexUnsafe(test.TestAddr)
	assert.NoError(t, err)
	transfer := banktypes.NewMsgSend(km.GetAddr(), to, sdk.NewCoins(sdk.NewInt64Coin(test.TestTokenName, 100)))
	payerAddr, err := sdk.AccAddressFromHexUnsafe(km.GetAddr().String())
	assert.NoError(t, err)
	mode := tx.BroadcastMode_BROADCAST_MODE_SYNC
	feeAmt := sdk.NewCoins(sdk.NewCoin("azkme", sdk.NewInt(int64(10000000000000)))) // gasPrice * gasLimit

	txOpt := &types.TxOption{
		Mode:       &mode,
		NoSimulate: true,
		GasLimit:   2000,
		Memo:       "test",
		FeePayer:   payerAddr,
		FeeAmount:  feeAmt, // 2000 * 5000000000
	}
	response, err := gnfdCli.BroadcastTx(context.Background(), []sdk.Msg{transfer}, txOpt)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), response.TxResponse.Code)
	t.Log(response.TxResponse.String())
}

func TestErrorOutWhenGasInfoNotFullProvided(t *testing.T) {
	km, err := keys.NewPrivateKeyManager(test.TestPrivateKey)
	assert.NoError(t, err)
	gnfdCli, err := NewMechainClient(test.TestRPCAddr, test.TestChainID, WithKeyManager(km))
	assert.NoError(t, err)
	to, err := sdk.AccAddressFromHexUnsafe(test.TestAddr)
	assert.NoError(t, err)
	transfer := banktypes.NewMsgSend(km.GetAddr(), to, sdk.NewCoins(sdk.NewInt64Coin(test.TestTokenName, 100)))
	payerAddr, err := sdk.AccAddressFromHexUnsafe(km.GetAddr().String())
	assert.NoError(t, err)
	mode := tx.BroadcastMode_BROADCAST_MODE_SYNC
	txOpt := &types.TxOption{
		Mode:       &mode,
		NoSimulate: true,
		Memo:       "test",
		FeePayer:   payerAddr,
	}
	_, err = gnfdCli.BroadcastTx(context.Background(), []sdk.Msg{transfer}, txOpt)
	assert.Equal(t, err, types.ErrGasInfoNotProvided)
}

func TestSimulateTx(t *testing.T) {
	km, err := keys.NewPrivateKeyManager(test.TestPrivateKey)
	assert.NoError(t, err)
	gnfdCli, err := NewMechainClient(test.TestRPCAddr, test.TestChainID, WithKeyManager(km))
	assert.NoError(t, err)
	to, err := sdk.AccAddressFromHexUnsafe(test.TestAddr)
	assert.NoError(t, err)
	transfer := banktypes.NewMsgSend(km.GetAddr(), to, sdk.NewCoins(sdk.NewInt64Coin(test.TestTokenName, 100)))
	simulateRes, err := gnfdCli.SimulateTx(context.Background(), []sdk.Msg{transfer}, nil)
	assert.NoError(t, err)
	t.Log(simulateRes.GasInfo.String())
}

func TestSendTokenWithCustomizedNonce(t *testing.T) {
	km, err := keys.NewPrivateKeyManager(test.TestPrivateKey)
	assert.NoError(t, err)
	gnfdCli, err := NewMechainClient(test.TestRPCAddr, test.TestChainID, WithKeyManager(km))
	assert.NoError(t, err)
	to, err := sdk.AccAddressFromHexUnsafe(test.TestAddr)
	assert.NoError(t, err)
	transfer := banktypes.NewMsgSend(km.GetAddr(), to, sdk.NewCoins(sdk.NewInt64Coin(test.TestTokenName, 100)))
	payerAddr, err := sdk.AccAddressFromHexUnsafe(km.GetAddr().String())
	assert.NoError(t, err)
	nonce, err := gnfdCli.GetNonce(context.Background())
	assert.NoError(t, err)
	for i := 0; i < 50; i++ {
		txOpt := &types.TxOption{
			GasLimit: 123456,
			Memo:     "test",
			FeePayer: payerAddr,
			Nonce:    nonce,
		}
		response, err := gnfdCli.BroadcastTx(context.Background(), []sdk.Msg{transfer}, txOpt)
		assert.NoError(t, err)
		nonce++
		assert.Equal(t, uint32(0), response.TxResponse.Code)
		t.Log(response.TxResponse.String())
	}
}

func TestSendTxWithGrpcConn(t *testing.T) {
	km, err := keys.NewPrivateKeyManager(test.TestPrivateKey)
	assert.NoError(t, err)
	gnfdCli, err := NewMechainClient("", test.TestChainID, WithKeyManager(km), WithGrpcConnectionAndDialOption(test.TestGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials())))
	assert.NoError(t, err)
	to, err := sdk.AccAddressFromHexUnsafe(test.TestAddr)
	assert.NoError(t, err)
	transfer := banktypes.NewMsgSend(km.GetAddr(), to, sdk.NewCoins(sdk.NewInt64Coin(test.TestTokenName, 100)))
	payerAddr, err := sdk.AccAddressFromHexUnsafe(km.GetAddr().String())
	assert.NoError(t, err)
	nonce, err := gnfdCli.GetNonce(context.Background())
	assert.NoError(t, err)
	txOpt := &types.TxOption{
		GasLimit: 123456,
		Memo:     "test",
		FeePayer: payerAddr,
		Nonce:    nonce,
	}
	response, err := gnfdCli.BroadcastTx(context.Background(), []sdk.Msg{transfer}, txOpt)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), response.TxResponse.Code)
	t.Log(response.TxResponse.String())
}

func TestSendTokenWithOverrideAccount(t *testing.T) {
	// which is not being used to send tx
	km, err := keys.NewPrivateKeyManager("2a3f0f19fbcb057e053696879207324c24f601ab47db92676cc4958ea9089761")
	assert.NoError(t, err)
	gnfdCli, err := NewMechainClient(test.TestRPCAddr, test.TestChainID, WithKeyManager(km))
	assert.NoError(t, err)

	km2, err := keys.NewPrivateKeyManager(test.TestPrivateKey)
	assert.NoError(t, err)

	assert.NoError(t, err)
	to, err := sdk.AccAddressFromHexUnsafe(test.TestAddr)
	assert.NoError(t, err)
	transfer := banktypes.NewMsgSend(km2.GetAddr(), to, sdk.NewCoins(sdk.NewInt64Coin(test.TestTokenName, 100)))
	payerAddr, err := sdk.AccAddressFromHexUnsafe(km2.GetAddr().String())
	assert.NoError(t, err)
	mode := tx.BroadcastMode_BROADCAST_MODE_SYNC
	feeAmt := sdk.NewCoins(sdk.NewCoin("azkme", sdk.NewInt(int64(10000000000000)))) // gasPrice * gasLimit
	txOpt := &types.TxOption{
		Mode:               &mode,
		NoSimulate:         true,
		GasLimit:           2000,
		Memo:               "test",
		FeePayer:           payerAddr,
		FeeAmount:          feeAmt, // 2000 * 5000000000
		OverrideKeyManager: &km2,
	}
	response, err := gnfdCli.BroadcastTx(context.Background(), []sdk.Msg{transfer}, txOpt)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), response.TxResponse.Code)
	t.Log(response.TxResponse.String())
}

func TestSendTXViaWebsocketClient(t *testing.T) {
	km, err := keys.NewPrivateKeyManager(test.TestPrivateKey)
	assert.NoError(t, err)
	gnfdCli, err := NewMechainClient(test.TestRPCAddr, test.TestChainID, WithKeyManager(km), WithWebSocketClient())
	assert.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	to := sdk.MustAccAddressFromHex(test.TestAddr)
	nonce, _ := gnfdCli.GetNonce(ctx)
	for i := 0; i < 500; i++ {
		assert.NoError(t, err)
		transfer := banktypes.NewMsgSend(km.GetAddr(), to, sdk.NewCoins(sdk.NewInt64Coin(test.TestTokenName, 12)))
		response, err := gnfdCli.BroadcastTx(ctx, []sdk.Msg{transfer}, &types.TxOption{Nonce: nonce})
		assert.NoError(t, err)
		nonce++
		assert.Equal(t, uint32(0), response.TxResponse.Code)
	}
}
