package storage

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	storagekeeper "github.com/evmos/evmos/v12/x/storage/keeper"

	"github.com/evmos/evmos/v12/x/evm/types"
)

type Contract struct {
	ctx           sdk.Context
	storageKeeper storagekeeper.Keeper
}

func NewPrecompiledContract(ctx sdk.Context, storageKeeper storagekeeper.Keeper) *Contract {
	return &Contract{
		ctx:           ctx,
		storageKeeper: storageKeeper,
	}
}

func (c *Contract) Address() common.Address {
	return storageAddress
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	method, err := GetMethodByID(input)
	if err != nil {
		return 0
	}

	switch method.Name {
	case CreateBucketMethodName:
		return CreateBucketGas
	case ListBucketsMethodName:
		return ListBucketsGas
	case HeadBucketMethodName:
		return HeadBucketGas
	case CreateObjectMethodName:
		return CreateObjectGas
	case ListObjectsMethodName:
		return ListObjectsGas
	case SealObjectMethodName:
		return SealObjectGas
	case SealObjectV2MethodName:
		return SealObjectV2Gas
	case UpdateObjectInfoMethodName:
		return UpdateObjectInfoGas
	case CreateGroupMethodName:
		return CreateGroupGas
	case ListGroupsMethodName:
		return ListGroupsGas
	case UpdateGroupMethodName:
		return UpdateGroupGas
	case HeadGroupMethodName:
		return HeadGroupGas
	case DeleteGroupMethodName:
		return DeleteGroupGas
	case HeadGroupMemberMethodName:
		return HeadGroupMemberGas
	case RenewGroupMemberMethodName:
		return RenewGroupMemberGas
	case SetTagForGroupMethodName:
		return SetTagForGroupGas
	default:
		return 0
	}
}

func (c *Contract) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) (ret []byte, err error) {
	if len(contract.Input) < 4 {
		return types.PackRetError("invalid input")
	}

	ctx, commit := c.ctx.CacheContext()
	snapshot := evm.StateDB.Snapshot()

	method, err := GetMethodByID(contract.Input)
	if err == nil {
		switch method.Name {
		case CreateBucketMethodName:
			ret, err = c.CreateBucket(ctx, evm, contract, readonly)
		case ListBucketsMethodName:
			ret, err = c.ListBuckets(ctx, evm, contract, readonly)
		case HeadBucketMethodName:
			ret, err = c.HeadBucket(ctx, evm, contract, readonly)
		case CreateObjectMethodName:
			ret, err = c.CreateObject(ctx, evm, contract, readonly)
		case ListObjectsMethodName:
			ret, err = c.ListObjects(ctx, evm, contract, readonly)
		case SealObjectMethodName:
			ret, err = c.SealObject(ctx, evm, contract, readonly)
		case SealObjectV2MethodName:
			ret, err = c.SealObjectV2(ctx, evm, contract, readonly)
		case UpdateObjectInfoMethodName:
			ret, err = c.UpdateObjectInfo(ctx, evm, contract, readonly)
		case CreateGroupMethodName:
			ret, err = c.CreateGroup(ctx, evm, contract, readonly)
		case ListGroupsMethodName:
			ret, err = c.ListGroups(ctx, evm, contract, readonly)
		case UpdateGroupMethodName:
			ret, err = c.UpdateGroup(ctx, evm, contract, readonly)
		case HeadGroupMethodName:
			ret, err = c.HeadGroup(ctx, evm, contract, readonly)
		case DeleteGroupMethodName:
			ret, err = c.DeleteGroup(ctx, evm, contract, readonly)
		case HeadGroupMemberMethodName:
			ret, err = c.HeadGroupMember(ctx, evm, contract, readonly)
		case RenewGroupMemberMethodName:
			ret, err = c.RenewGroupMember(ctx, evm, contract, readonly)
		case SetTagForGroupMethodName:
			ret, err = c.SetTagForGroup(ctx, evm, contract, readonly)
		}
	}

	if err != nil {
		// revert evm state
		evm.StateDB.RevertToSnapshot(snapshot)
		return types.PackRetError(err.Error())
	}

	// commit and append events
	commit()
	return ret, nil
}

func (c *Contract) AddLog(evm *vm.EVM, event abi.Event, topics []common.Hash, args ...interface{}) error {
	data, newTopic, err := types.PackTopicData(event, topics, args...)
	if err != nil {
		return err
	}
	evm.StateDB.AddLog(&ethtypes.Log{
		Address:     c.Address(),
		Topics:      newTopic,
		Data:        data,
		BlockNumber: evm.Context.BlockNumber.Uint64(),
	})
	return nil
}
