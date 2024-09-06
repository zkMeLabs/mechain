package sample

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/0xPolygon/polygon-edge/bls"
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cometbft/cometbft/votepool"
	"github.com/cosmos/cosmos-sdk/crypto/keys/eth/ethsecp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func RandAccAddress() sdk.AccAddress {
	pk, err := ethsecp256k1.GenPrivKey()
	if err != nil {
		panic(err)
	}
	return sdk.AccAddress(pk.PubKey().Address())
}

func RandAccAddressHex() string {
	pk, err := ethsecp256k1.GenPrivKey()
	if err != nil {
		panic(err)
	}
	return sdk.AccAddress(pk.PubKey().Address()).String()
}

func RandSignBytes() (addr sdk.AccAddress, signBytes, sig []byte) {
	signBytes = RandStr(256)
	privKey, _ := ethsecp256k1.GenPrivKey()

	sig, _ = privKey.Sign(crypto.Keccak256(signBytes))
	pk := privKey.PubKey()
	addr = sdk.AccAddress(pk.Address())
	return addr, signBytes, sig
}

func Checksum() []byte {
	return crypto.Keccak256(RandStr(256))
}

func RandStr(length int) []byte {
	randBytes := make([]byte, length/2)
	// #nosec
	_, _ = rand.Read(randBytes)
	return randBytes
}

func RandBlsPubKey() []byte {
	blsPrivKey, _ := bls.GenerateBlsKey()
	return blsPrivKey.PublicKey().Marshal()
}

func RandBlsPubKeyHex() string {
	blsPrivKey, _ := bls.GenerateBlsKey()
	return hex.EncodeToString(blsPrivKey.PublicKey().Marshal())
}

func RandBlsPubKeyAndBlsProofBz() ([]byte, []byte) {
	blsPriv, _ := bls.GenerateBlsKey()
	blsPubKeyBz := blsPriv.PublicKey().Marshal()
	blsProof, _ := blsPriv.Sign(tmhash.Sum(blsPubKeyBz), votepool.DST)
	blsProofBz, _ := blsProof.Marshal()
	return blsPubKeyBz, blsProofBz
}

func RandBlsPubKeyAndBlsProof() (string, string) {
	blsPubKey, proof := RandBlsPubKeyAndBlsProofBz()
	return hex.EncodeToString(blsPubKey), hex.EncodeToString(proof)
}
