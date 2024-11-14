package storage

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
	storagekeeper "github.com/evmos/evmos/v12/x/storage/keeper"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"

	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	DeleteObjectMethodName = "deleteObject"
)

func (c *Contract) DeleteObject(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("delete bucket method readonly")
	}

	method := GetAbiMethod(DeleteObjectMethodName)

	var args DeleteObjectArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgDeleteObject{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
		ObjectName: args.ObjectName,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.DeleteObject(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DeleteObjectMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

type DeleteObjectArgs struct {
	BucketName string `abi:"bucketName"`
	ObjectName string `abi:"objectName"`
}

// Validate DeleteObjectArgs args
func (args *DeleteObjectArgs) Validate() error {
	if args.BucketName == "" {
		return errors.New("bucket name is empty")
	}
	if args.ObjectName == "" {
		return errors.New("object name is empty")
	}
	return nil
}
