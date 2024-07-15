package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
)

// This is an evil token. Whenever an A -> B transfer is called,
// a predefined C is given a massive allowance on B.
var (
	//go:embed compiled_contracts/ERC721NonTransferable.json
	ERC721NonTransferableJSON []byte //nolint: golint

	// ERC721NonTransferableContract is the compiled erc721 contract
	ERC721NonTransferableContract evmtypes.CompiledContract

	// ERC721NonTransferableAddress is the erc721 module address
	ERC721NonTransferableAddress common.Address

	// ObjectControlHubAddress is the object nft control hub address
	ObjectControlHubAddress common.Address
)

func init() {
	ERC721NonTransferableAddress = common.HexToAddress("0x0000000000000000000000000000000000003000")
	ObjectControlHubAddress = common.HexToAddress("0x00000Be6819f41400225702D32d3dd23663Dd690")

	err := json.Unmarshal(ERC721NonTransferableJSON, &ERC721NonTransferableContract)
	if err != nil {
		panic(err)
	}
}
