package hd

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/evmos/evmos/v12/types"
)

func BenchmarkEthSecp256k1Algo_Derive(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		deriveFn := hd.EthSecp256k1.Derive()
		if _, err := deriveFn(mnemonic, keyring.DefaultBIP39Passphrase, types.BIP44HDPath); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEthSecp256k1Algo_Generate(b *testing.B) {
	bz, err := hd.EthSecp256k1.Derive()(mnemonic, keyring.DefaultBIP39Passphrase, types.BIP44HDPath)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		(&hd.EthSecp256k1).Generate()(bz)
	}
}
