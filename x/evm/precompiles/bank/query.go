package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	TotalSupplyGas = 50_000

	TotalSupplyMethodName = "totalSupply"
)

// TotalSupply queries the total supply of all coins.
func (c *Contract) TotalSupply(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(TotalSupplyMethodName)

	msg := &banktypes.QueryTotalSupplyRequest{}

	res, err := c.bankKeeper.TotalSupply(ctx, msg)
	if err != nil {
		return nil, err
	}

	balances := make([]Coin, 0, len(res.Supply))
	for _, balance := range res.Supply {
		balances = append(balances, Coin{
			Denom:  balance.Denom,
			Amount: balance.Amount.BigInt(),
		})
	}

	return method.Outputs.Pack(balances)
}
