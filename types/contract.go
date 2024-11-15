package types

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	EmptyEvmAddress     = "0x0000000000000000000000000000000000000000"
	BankAddress         = "0x0000000000000000000000000000000000001000"
	AuthAddress         = "0x0000000000000000000000000000000000001001"
	GovAddress          = "0x0000000000000000000000000000000000001002"
	StakingAddress      = "0x0000000000000000000000000000000000001003"
	DistributionAddress = "0x0000000000000000000000000000000000001004"
	SlashingAddress     = "0x0000000000000000000000000000000000001005"
	EvidenceAddress     = "0x0000000000000000000000000000000000001006"
	EpochsAddress       = "0x0000000000000000000000000000000000001007"
	AuthzAddress        = "0x0000000000000000000000000000000000001008"
	FeemarketAddress    = "0x0000000000000000000000000000000000001009"
	PaymentAddress      = "0x000000000000000000000000000000000000100a"
	PermissionAddress   = "0x000000000000000000000000000000000000100b"
	VirtualGroupAddress = "0x0000000000000000000000000000000000002000"
	StorageAddress      = "0x0000000000000000000000000000000000002001"
	SpAddress           = "0x0000000000000000000000000000000000002002"
)

type Contract struct {
	Address common.Address
	ABI     abi.ABI
	Bin     []byte
	Code    []byte
}

func (c Contract) CodeHash() common.Hash {
	return crypto.Keccak256Hash(c.Code)
}

func MustABIJson(str string) abi.ABI {
	j, err := abi.JSON(strings.NewReader(str))
	if err != nil {
		panic(err)
	}
	return j
}
