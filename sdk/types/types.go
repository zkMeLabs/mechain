package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/evmos/evmos/v12/sdk/keys"
)

const (
	Denom = "azkme"

	// DecimalZKME defines number of azkme decimal places
	DecimalZKME = 18

	// DecimalGwei defines number of gweiZKME decimal places
	DecimalGwei = 9
)

type TxOption struct {
	Mode               *tx.BroadcastMode
	NoSimulate         bool
	GasLimit           uint64
	FeeAmount          sdk.Coins
	Nonce              uint64
	FeePayer           sdk.AccAddress
	FeeGranter         sdk.AccAddress
	Tip                *tx.Tip
	Memo               string
	OverrideKeyManager *keys.KeyManager
}

func NewIntFromInt64WithDecimal(amount int64, decimal int64) sdkmath.Int {
	return sdk.NewInt(amount).Mul(sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(decimal), nil)))
}
