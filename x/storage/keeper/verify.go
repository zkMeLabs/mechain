package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) VerifyPaymentAccount(_ sdk.Context, paymentAddress string, ownerAcc sdk.AccAddress) (sdk.AccAddress, error) {
	paymentAcc, err := sdk.AccAddressFromHexUnsafe(paymentAddress)
	if err == sdk.ErrEmptyHexAddress {
		return ownerAcc, nil
	} else if err != nil {
		return nil, err
	}

	return paymentAcc, nil
}
