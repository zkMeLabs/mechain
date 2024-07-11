package storage

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	testRPC       = "http://localhost:8545"
	privateKeyHex = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
)

var (
	privateKey, _     = crypto.HexToECDSA(privateKeyHex)
	publicKey         = privateKey.Public()
	publicKeyECDSA, _ = publicKey.(*ecdsa.PublicKey)
	testAddressHex    = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
)

func getTestRPC() string {
	return "http://localhost:8545"
}

func getPrivateKeyHex() string {
	return "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
}
