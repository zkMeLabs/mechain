package storage

import (
	"bytes"
	"errors"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	gtypes "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/evm/types"
	storagekeeper "github.com/evmos/evmos/v12/x/storage/keeper"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
)

const (
	CreateGroupMethodName      = "createGroup"
	UpdateGroupMethodName      = "updateGroup"
	DeleteGroupMethodName      = "deleteGroup"
	RenewGroupMemberMethodName = "renewGroupMember"
	SetTagForGroupMethodName   = "setTagForGroup"
	HeadGroupMethodName        = "headGroup"
	HeadGroupMemberMethodName  = "headGroupMember"
	ListGroupsMethodName       = "listGroups"
)

func (c *Contract) registerGroupMethod() {
	c.registerMethod(CreateGroupMethodName, 60_000, c.CreateGroup, "CreateGroup")
	c.registerMethod(ListGroupsMethodName, 50_000, c.ListGroups, "")
	c.registerMethod(HeadGroupMethodName, 50_000, c.HeadGroup, "")
	c.registerMethod(HeadGroupMemberMethodName, 50_000, c.HeadGroupMember, "")
	c.registerMethod(UpdateGroupMethodName, 60_000, c.UpdateGroup, "UpdateGroup")
	c.registerMethod(DeleteGroupMethodName, 60_000, c.DeleteGroup, "DeleteGroup")
	c.registerMethod(RenewGroupMemberMethodName, 60_000, c.RenewGroupMember, "RenewGroupMember")
	c.registerMethod(SetTagForGroupMethodName, 60_000, c.SetTagForGroup, "SetTagForGroup")
}

// ListGroups queries the user's total groups.
func (c *Contract) ListGroups(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(ListGroupsMethodName)
	var args ListGroupsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &storagetypes.QueryListGroupsRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
		GroupOwner: args.GroupOwner.String(),
	}
	res, err := c.storageKeeper.ListGroups(ctx, msg)
	if err != nil {
		return nil, err
	}
	var groupInfos []GroupInfo
	for _, groupInfo := range res.GroupInfos {
		groupInfos = append(groupInfos, GroupInfo{
			Owner:      common.HexToAddress(groupInfo.Owner),
			GroupName:  groupInfo.GroupName,
			SourceType: uint8(groupInfo.SourceType),
			Id:         groupInfo.Id.BigInt(),
			Extra:      groupInfo.Extra,
			Tags:       outputTags(groupInfo.Tags),
		})
	}
	return method.Outputs.Pack(groupInfos, outputPageResponse(res.Pagination))
}

// HeadGroup queries the group's info.
func (c *Contract) HeadGroup(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadGroupMethodName)
	var args HeadGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryHeadGroupRequest{
		GroupOwner: args.GroupOwner.String(),
		GroupName:  args.GroupName,
	}
	res, err := c.storageKeeper.HeadGroup(ctx, msg)
	if err != nil {
		return nil, err
	}
	groupInfo := GroupInfo{
		Owner:      common.HexToAddress(res.GroupInfo.Owner),
		GroupName:  res.GroupInfo.GroupName,
		SourceType: uint8(res.GroupInfo.SourceType),
		Id:         res.GroupInfo.Id.BigInt(),
		Extra:      res.GroupInfo.Extra,
		Tags:       outputTags(res.GroupInfo.Tags),
	}
	return method.Outputs.Pack(groupInfo)
}

// HeadGroupMember queries the group member's info.
func (c *Contract) HeadGroupMember(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadGroupMemberMethodName)
	var args HeadGroupMemberArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryHeadGroupMemberRequest{
		Member:     args.Member.String(),
		GroupOwner: args.GroupOwner.String(),
		GroupName:  args.GroupName,
	}
	res, err := c.storageKeeper.HeadGroupMember(ctx, msg)
	if err != nil {
		return nil, err
	}
	expirationTime := res.GroupMember.ExpirationTime.Unix()
	groupMemberInfo := GroupMember{
		Id:             res.GroupMember.Id.BigInt(),
		GroupId:        res.GroupMember.GroupId.BigInt(),
		Member:         common.HexToAddress(res.GroupMember.Member),
		ExpirationTime: expirationTime,
	}
	return method.Outputs.Pack(groupMemberInfo)
}

func (c *Contract) CreateGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("create group method readonly")
	}
	method := GetAbiMethod(CreateGroupMethodName)
	var args CreateGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.MsgCreateGroup{
		Creator:   contract.Caller().String(),
		GroupName: args.GroupName,
		Extra:     args.Extra,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	res, err := server.CreateGroup(ctx, msg)
	if err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[CreateGroupMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		res.GroupId.BigInt(),
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) UpdateGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("update group method readonly")
	}
	method := GetAbiMethod(UpdateGroupMethodName)
	var args UpdateGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	membersToAdd := make([]*storagetypes.MsgGroupMember, 0)
	if args.MembersToAdd != nil {
		for i, members := range args.MembersToAdd {
			var exp time.Time
			if args.ExpirationTime[i] != 0 {
				exp = time.Unix(args.ExpirationTime[i], 0)
			} else {
				exp = storagetypes.MaxTimeStamp
			}
			membersToAdd = append(membersToAdd, &storagetypes.MsgGroupMember{
				Member:         members.String(),
				ExpirationTime: &exp,
			})
		}
	}
	var membersToDelete []string
	if args.MembersToDelete != nil {
		for _, members := range args.MembersToDelete {
			membersToDelete = append(membersToDelete, members.String())
		}
	}
	msg := &storagetypes.MsgUpdateGroupMember{
		Operator:        contract.Caller().String(),
		GroupOwner:      args.GroupOwner.String(),
		GroupName:       args.GroupName,
		MembersToAdd:    membersToAdd,
		MembersToDelete: membersToDelete,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)

	if _, err := server.UpdateGroupMember(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[UpdateGroupMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) DeleteGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("delete group method readonly")
	}
	method := GetAbiMethod(DeleteGroupMethodName)
	var args DeleteGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.MsgDeleteGroup{
		Operator:  contract.Caller().String(),
		GroupName: args.GroupName,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.DeleteGroup(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DeleteGroupMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) RenewGroupMember(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("renew group member method readonly")
	}
	method := GetAbiMethod(RenewGroupMemberMethodName)
	var args RenewGroupMemberArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if args.GroupName == "" {
		return nil, errors.New("group name is empty")
	}
	if len(args.Members) == 0 {
		return nil, errors.New("no renew member")
	}
	if args.ExpirationTime != nil && len(args.Members) != len(args.ExpirationTime) {
		return nil, errors.New("please provide expirationTime for every renew member")
	}
	membersToRenew := make([]*storagetypes.MsgGroupMember, 0)
	if args.Members != nil {
		for i, members := range args.Members {
			var exp time.Time
			if args.ExpirationTime[i] != 0 {
				exp = time.Unix(args.ExpirationTime[i], 0)
			} else {
				exp = storagetypes.MaxTimeStamp
			}
			membersToRenew = append(membersToRenew, &storagetypes.MsgGroupMember{
				Member:         members.String(),
				ExpirationTime: &exp,
			})
		}
	}
	msg := &storagetypes.MsgRenewGroupMember{
		Operator:   contract.Caller().String(),
		GroupOwner: args.GroupOwner.String(),
		GroupName:  args.GroupName,
		Members:    membersToRenew,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.RenewGroupMember(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[RenewGroupMemberMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) SetTagForGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("set tag for group method readonly")
	}
	method := GetAbiMethod(SetTagForGroupMethodName)
	var args SetTagForGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	addr, err := sdk.AccAddressFromHexUnsafe(contract.Caller().String())
	if err != nil {
		return nil, err
	}
	grn := gtypes.NewGroupGRN(addr, args.GroupName)
	var tags storagetypes.ResourceTags
	if args.Tags != nil {
		for _, tag := range args.Tags {
			tags.Tags = append(tags.Tags, storagetypes.ResourceTags_Tag{
				Key:   tag.Key,
				Value: tag.Value,
			})
		}
	}
	msg := &storagetypes.MsgSetTag{
		Operator: contract.Caller().String(),
		Resource: grn.String(),
		Tags:     &tags,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err = server.SetTag(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[SetTagForGroupMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}
