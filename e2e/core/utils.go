package core

import (
	cryptoRand "crypto/rand"
	"fmt"
	"math/rand"

	"github.com/0xPolygon/polygon-edge/bls"
	"github.com/cometbft/cometbft/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sigs.k8s.io/yaml"

	"github.com/evmos/evmos/v12/sdk/keys"
	gnfdtypes "github.com/evmos/evmos/v12/types"
)

func GenRandomAddr() sdk.AccAddress {
	// #nosec
	return sdk.AccAddress(crypto.AddressHash([]byte(fmt.Sprintf("%d", rand.Int()))))
}

func GenRandomHexString(len int) string {
	b := make([]byte, len)
	// #nosec
	_, err := cryptoRand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}

func GenRandomKeyManager() keys.KeyManager {
	keyManager, err := keys.NewPrivateKeyManager(GenRandomHexString(32))
	if err != nil {
		panic(err)
	}
	return keyManager
}

func YamlString(data interface{}) string {
	bz, err := yaml.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(bz)
}

// RandInt64 generate random int64 between min and max
func RandInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min) //nolint
}

func BlsSignAndVerify(sp *StorageProvider, signBz [32]byte) ([]byte, error) {
	secondarySig, err := sp.BlsKey.Sign(signBz[:])
	if err != nil {
		return nil, err
	}
	pubKey, err := bls.UnmarshalPublicKey(sp.BlsKey.PubKey().Bytes())
	if err != nil {
		return nil, err
	}
	err = gnfdtypes.VerifyBlsSignature(pubKey, signBz, secondarySig)
	if err != nil {
		return nil, err
	}
	return secondarySig, nil
}

func BlsAggregateAndVerify(secondarySPBlsPubKeys []*bls.PublicKey, signBz [32]byte, secondarySigs [][]byte) ([]byte, error) {
	signatures := make(bls.Signatures, 0, len(secondarySigs))
	for _, sigBts := range secondarySigs {
		sig, err := bls.UnmarshalSignature(sigBts)
		if err != nil {
			return nil, err
		}
		signatures = append(signatures, sig)
	}
	blsSigs, err := signatures.Aggregate().Marshal()
	if err != nil {
		return nil, err
	}
	err = gnfdtypes.VerifyBlsAggSignature(secondarySPBlsPubKeys, signBz, blsSigs)
	if err != nil {
		return nil, err
	}
	return blsSigs, nil
}
