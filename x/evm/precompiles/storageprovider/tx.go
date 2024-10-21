package storageprovider

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
	spkeeper "github.com/evmos/evmos/v12/x/sp/keeper"
	sptypes "github.com/evmos/evmos/v12/x/sp/types"

	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	UpdateSPPriceMethodName = "updateSPPrice"
)

func (c *Contract) registerTx() {
	c.registerMethod(UpdateSPPriceMethodName, 60_000, c.UpdateSPPrice, "UpdateSPPrice")
}

func (c *Contract) UpdateSPPrice(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("update sp price method readonly")
	}
	method := GetAbiMethod(UpdateSPPriceMethodName)
	var args UpdateSPPriceArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &sptypes.MsgUpdateSpStoragePrice{
		SpAddress:     contract.Caller().String(),
		ReadPrice:     args.ReadPrice,
		FreeReadQuota: args.FreeReadQuota,
		StorePrice:    args.StorePrice,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := spkeeper.NewMsgServerImpl(c.spKeeper)
	if _, err := server.UpdateSpStoragePrice(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[UpdateSPPriceMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}
