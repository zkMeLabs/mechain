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

	// ObjectERC721TokenAddress is the object erc721 module address
	ObjectERC721TokenAddress common.Address

	// BucketERC721TokenAddress is the bucket erc721 module address
	BucketERC721TokenAddress common.Address

	// GroupERC721TokenAddress is the group erc721 module address
	GroupERC721TokenAddress common.Address

	// ObjectControlHubAddress is the object nft control hub address
	ObjectControlHubAddress common.Address

	// BucketControlHubAddress is the bucket nft control hub address
	BucketControlHubAddress common.Address

	// GroupControlHubAddress is the group nft control hub address
	GroupControlHubAddress common.Address
)

func init() {
	ObjectERC721TokenAddress = common.HexToAddress("0x0000000000000000000000000000000000003000")
	BucketERC721TokenAddress = common.HexToAddress("0x0000000000000000000000000000000000003001")
	GroupERC721TokenAddress = common.HexToAddress("0x0000000000000000000000000000000000003002")
	ObjectControlHubAddress = common.HexToAddress("0x000000000000000000000000000000000000dead")
	GroupControlHubAddress = common.HexToAddress("0x000000000000000000000000000000000000dead")
	BucketControlHubAddress = common.HexToAddress("0x000000000000000000000000000000000000dead")

	err := json.Unmarshal(ERC721NonTransferableJSON, &ERC721NonTransferableContract)
	if err != nil {
		panic(err)
	}
}
