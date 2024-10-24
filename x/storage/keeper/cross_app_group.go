package keeper

import (
	"encoding/hex"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/evmos/evmos/v12/x/storage/types"
)

var _ sdk.CrossChainApplication = &GroupApp{}

type GroupApp struct {
	storageKeeper types.StorageKeeper
}

func NewGroupApp(keeper types.StorageKeeper) *GroupApp {
	return &GroupApp{
		storageKeeper: keeper,
	}
}

func (app *GroupApp) ExecuteAckPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, payload []byte) sdk.ExecuteResult {
	pack, err := types.DeserializeCrossChainPackage(payload, types.GroupChannelID, sdk.AckCrossChainPackageType)
	if err != nil {
		app.storageKeeper.Logger(ctx).Error("deserialize group cross chain package error", "payload", hex.EncodeToString(payload), "error", err.Error())
		panic("deserialize group cross chain package error")
	}

	var operationType uint8
	var result sdk.ExecuteResult
	switch p := pack.(type) {
	case *types.MirrorGroupAckPackage:
		operationType = types.OperationMirrorGroup
		result = app.handleMirrorGroupAckPackage(ctx, appCtx, p)
	case *types.CreateGroupAckPackage:
		operationType = types.OperationCreateGroup
		result = app.handleCreateGroupAckPackage(ctx, appCtx, p)
	case *types.DeleteGroupAckPackage:
		operationType = types.OperationDeleteGroup
		result = app.handleDeleteGroupAckPackage(ctx, appCtx, p)
	case *types.UpdateGroupMemberAckPackage:
		operationType = types.OperationUpdateGroupMember
		result = app.handleUpdateGroupMemberAckPackage(ctx, appCtx, p)
	default:
		panic("unknown cross chain ack package type")
	}

	if len(result.Payload) != 0 {
		wrapPayload := types.CrossChainPackage{
			OperationType: operationType,
			Package:       result.Payload,
		}
		result.Payload = wrapPayload.MustSerialize()
	}

	return result
}

func (app *GroupApp) ExecuteFailAckPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, payload []byte) sdk.ExecuteResult {
	pack, err := types.DeserializeCrossChainPackage(payload, types.GroupChannelID, sdk.FailAckCrossChainPackageType)
	if err != nil {
		app.storageKeeper.Logger(ctx).Error("deserialize group cross chain package error", "payload", hex.EncodeToString(payload), "error", err.Error())
		panic("deserialize group cross chain package error")
	}

	var operationType uint8
	var result sdk.ExecuteResult
	switch p := pack.(type) {
	case *types.MirrorGroupSynPackage:
		operationType = types.OperationMirrorGroup
		result = app.handleMirrorGroupFailAckPackage(ctx, appCtx, p)
	case *types.CreateGroupSynPackage:
		operationType = types.OperationCreateGroup
		result = app.handleCreateGroupFailAckPackage(ctx, appCtx, p)
	case *types.DeleteGroupSynPackage:
		operationType = types.OperationDeleteGroup
		result = app.handleDeleteGroupFailAckPackage(ctx, appCtx, p)
	case *types.UpdateGroupMemberSynPackage:
		operationType = types.OperationUpdateGroupMember
		result = app.handleUpdateGroupMemberFailAckPackage(ctx, appCtx, p)
	default:
		panic("unknown cross chain ack package type")
	}

	if len(result.Payload) != 0 {
		wrapPayload := types.CrossChainPackage{
			OperationType: operationType,
			Package:       result.Payload,
		}
		result.Payload = wrapPayload.MustSerialize()
	}

	return result
}

func (app *GroupApp) ExecuteSynPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, payload []byte) sdk.ExecuteResult {
	pack, err := types.DeserializeCrossChainPackage(payload, types.GroupChannelID, sdk.SynCrossChainPackageType)
	if err != nil {
		app.storageKeeper.Logger(ctx).Error("deserialize group cross chain package error", "payload", hex.EncodeToString(payload), "error", err.Error())
		panic("deserialize group cross chain package error")
	}

	var operationType uint8
	var result sdk.ExecuteResult
	switch p := pack.(type) {
	case *types.MirrorGroupSynPackage:
		operationType = types.OperationMirrorGroup
		result = app.handleMirrorGroupSynPackage(ctx, appCtx, p)
	case *types.CreateGroupSynPackage:
		operationType = types.OperationCreateGroup
		result = app.handleCreateGroupSynPackage(ctx, appCtx, p)
	case *types.DeleteGroupSynPackage:
		operationType = types.OperationDeleteGroup
		result = app.handleDeleteGroupSynPackage(ctx, appCtx, p)
	case *types.UpdateGroupMemberSynPackage:
		operationType = types.OperationUpdateGroupMember
		result = app.handleUpdateGroupMemberSynPackage(ctx, appCtx, p)
	default:
		return sdk.ExecuteResult{
			Err: types.ErrInvalidCrossChainPackage,
		}
	}

	if len(result.Payload) != 0 {
		wrapPayload := types.CrossChainPackage{
			OperationType: operationType,
			Package:       result.Payload,
		}
		result.Payload = wrapPayload.MustSerialize()
	}

	return result
}

//nolint:unparam
func (app *GroupApp) handleDeleteGroupAckPackage(ctx sdk.Context, _ *sdk.CrossChainAppContext, _ *types.DeleteGroupAckPackage) sdk.ExecuteResult {
	app.storageKeeper.Logger(ctx).Error("received delete group ack package ")

	return sdk.ExecuteResult{}
}

//nolint:unparam
func (app *GroupApp) handleDeleteGroupFailAckPackage(ctx sdk.Context, _ *sdk.CrossChainAppContext, _ *types.DeleteGroupSynPackage) sdk.ExecuteResult {
	app.storageKeeper.Logger(ctx).Error("received delete group fail ack package ")

	return sdk.ExecuteResult{}
}

func (app *GroupApp) handleDeleteGroupSynPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, deleteGroupPackage *types.DeleteGroupSynPackage) sdk.ExecuteResult {
	err := deleteGroupPackage.ValidateBasic()
	if err != nil {
		return sdk.ExecuteResult{
			Payload: types.DeleteGroupAckPackage{
				Status:    types.StatusFail,
				ExtraData: deleteGroupPackage.ExtraData,
			}.MustSerialize(),
			Err: err,
		}
	}

	app.storageKeeper.Logger(ctx).Info("process delete group syn package", "group id", deleteGroupPackage.ID.String())

	groupInfo, found := app.storageKeeper.GetGroupInfoById(ctx, math.NewUintFromBigInt(deleteGroupPackage.ID))
	if !found {
		return sdk.ExecuteResult{
			Payload: types.DeleteGroupAckPackage{
				Status:    types.StatusFail,
				ExtraData: deleteGroupPackage.ExtraData,
			}.MustSerialize(),
			Err: types.ErrNoSuchGroup,
		}
	}

	sourceType, err := app.storageKeeper.GetSourceTypeByChainId(ctx, appCtx.SrcChainId)
	if err != nil {
		return sdk.ExecuteResult{
			Err: err,
		}
	}

	err = app.storageKeeper.DeleteGroup(ctx,
		deleteGroupPackage.Operator,
		groupInfo.GroupName,
		types.DeleteGroupOptions{
			SourceType: sourceType,
		},
	)
	if err != nil {
		return sdk.ExecuteResult{
			Payload: types.DeleteGroupAckPackage{
				Status:    types.StatusFail,
				ExtraData: deleteGroupPackage.ExtraData,
			}.MustSerialize(),
			Err: err,
		}
	}

	return sdk.ExecuteResult{
		Payload: types.DeleteGroupAckPackage{
			Status:    types.StatusSuccess,
			ID:        groupInfo.Id.BigInt(),
			ExtraData: deleteGroupPackage.ExtraData,
		}.MustSerialize(),
	}
}

//nolint:unparam
func (app *GroupApp) handleCreateGroupAckPackage(ctx sdk.Context, _ *sdk.CrossChainAppContext, _ *types.CreateGroupAckPackage) sdk.ExecuteResult {
	app.storageKeeper.Logger(ctx).Error("received create group ack package ")

	return sdk.ExecuteResult{}
}

//nolint:unparam
func (app *GroupApp) handleCreateGroupFailAckPackage(ctx sdk.Context, _ *sdk.CrossChainAppContext, _ *types.CreateGroupSynPackage) sdk.ExecuteResult {
	app.storageKeeper.Logger(ctx).Error("received create group fail ack package ")

	return sdk.ExecuteResult{}
}

func (app *GroupApp) handleCreateGroupSynPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, createGroupPackage *types.CreateGroupSynPackage) sdk.ExecuteResult {
	err := createGroupPackage.ValidateBasic()
	if err != nil {
		return sdk.ExecuteResult{
			Payload: types.CreateGroupAckPackage{
				Status:    types.StatusFail,
				Creator:   createGroupPackage.Creator,
				ExtraData: createGroupPackage.ExtraData,
			}.MustSerialize(),
			Err: err,
		}
	}
	app.storageKeeper.Logger(ctx).Info("process create group syn package", "group name", createGroupPackage.GroupName)

	sourceType, err := app.storageKeeper.GetSourceTypeByChainId(ctx, appCtx.SrcChainId)
	if err != nil {
		return sdk.ExecuteResult{
			Err: err,
		}
	}

	groupID, err := app.storageKeeper.CreateGroup(ctx,
		createGroupPackage.Creator,
		createGroupPackage.GroupName,
		types.CreateGroupOptions{
			SourceType: sourceType,
		},
	)
	if err != nil {
		return sdk.ExecuteResult{
			Payload: types.CreateGroupAckPackage{
				Status:    types.StatusFail,
				Creator:   createGroupPackage.Creator,
				ExtraData: createGroupPackage.ExtraData,
			}.MustSerialize(),
			Err: err,
		}
	}

	return sdk.ExecuteResult{
		Payload: types.CreateGroupAckPackage{
			Status:    types.StatusSuccess,
			ID:        groupID.BigInt(),
			Creator:   createGroupPackage.Creator,
			ExtraData: createGroupPackage.ExtraData,
		}.MustSerialize(),
	}
}

func (app *GroupApp) handleMirrorGroupAckPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, ackPackage *types.MirrorGroupAckPackage) sdk.ExecuteResult {
	groupInfo, found := app.storageKeeper.GetGroupInfoById(ctx, math.NewUintFromBigInt(ackPackage.Id))
	if !found {
		app.storageKeeper.Logger(ctx).Error("group does not exist", "group id", ackPackage.Id.String())
		return sdk.ExecuteResult{
			Err: types.ErrNoSuchGroup,
		}
	}

	sourceType, err := app.storageKeeper.GetSourceTypeByChainId(ctx, appCtx.SrcChainId)
	if err != nil {
		return sdk.ExecuteResult{
			Err: err,
		}
	}

	if ackPackage.Status == types.StatusSuccess {
		groupInfo.SourceType = sourceType

		app.storageKeeper.SetGroupInfo(ctx, groupInfo)
	}

	if err := ctx.EventManager().EmitTypedEvents(&types.EventMirrorGroupResult{
		Status:      uint32(ackPackage.Status),
		GroupName:   groupInfo.GroupName,
		GroupId:     groupInfo.Id,
		DestChainId: uint32(appCtx.SrcChainId),
	}); err != nil {
		return sdk.ExecuteResult{
			Err: err,
		}
	}

	return sdk.ExecuteResult{}
}

func (app *GroupApp) handleMirrorGroupFailAckPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, mirrorGroupPackage *types.MirrorGroupSynPackage) sdk.ExecuteResult {
	groupInfo, found := app.storageKeeper.GetGroupInfoById(ctx, math.NewUintFromBigInt(mirrorGroupPackage.ID))
	if !found {
		app.storageKeeper.Logger(ctx).Error("group does not exist", "group id", mirrorGroupPackage.ID.String())
		return sdk.ExecuteResult{
			Err: types.ErrNoSuchGroup,
		}
	}

	groupInfo.SourceType = types.SOURCE_TYPE_ORIGIN
	app.storageKeeper.SetGroupInfo(ctx, groupInfo)

	if err := ctx.EventManager().EmitTypedEvents(&types.EventMirrorGroupResult{
		Status:      uint32(types.StatusFail),
		GroupName:   groupInfo.GroupName,
		GroupId:     groupInfo.Id,
		DestChainId: uint32(appCtx.SrcChainId),
	}); err != nil {
		return sdk.ExecuteResult{
			Err: err,
		}
	}
	return sdk.ExecuteResult{}
}

//nolint:unparam
func (app *GroupApp) handleMirrorGroupSynPackage(ctx sdk.Context, _ *sdk.CrossChainAppContext, _ *types.MirrorGroupSynPackage) sdk.ExecuteResult {
	app.storageKeeper.Logger(ctx).Error("received mirror group syn ack package ")

	return sdk.ExecuteResult{}
}

func (app *GroupApp) handleUpdateGroupMemberSynPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, updateGroupPackage *types.UpdateGroupMemberSynPackage) sdk.ExecuteResult {
	err := updateGroupPackage.ValidateBasic()
	if err != nil {
		return sdk.ExecuteResult{
			Payload: types.UpdateGroupMemberAckPackage{
				Status:    types.StatusFail,
				Operator:  updateGroupPackage.Operator,
				ExtraData: updateGroupPackage.ExtraData,
			}.MustSerialize(),
			Err: err,
		}
	}

	groupInfo, found := app.storageKeeper.GetGroupInfoById(ctx, math.NewUintFromBigInt(updateGroupPackage.GroupID))
	if !found {
		return sdk.ExecuteResult{
			Payload: types.UpdateGroupMemberAckPackage{
				Status:    types.StatusFail,
				Operator:  updateGroupPackage.Operator,
				ExtraData: updateGroupPackage.ExtraData,
			}.MustSerialize(),
			Err: types.ErrNoSuchGroup,
		}
	}

	sourceType, err := app.storageKeeper.GetSourceTypeByChainId(ctx, appCtx.SrcChainId)
	if err != nil {
		return sdk.ExecuteResult{
			Err: err,
		}
	}

	switch updateGroupPackage.OperationType {
	case types.OperationAddGroupMember, types.OperationDeleteGroupMember:
		err = app.handleAddOrDeleteGroupMemberOperation(ctx, sourceType, groupInfo, updateGroupPackage)
	case types.OperationRenewGroupMember:
		err = app.handleRenewGroupOperation(ctx, sourceType, groupInfo, updateGroupPackage)
	}

	if err != nil {
		return sdk.ExecuteResult{
			Payload: types.UpdateGroupMemberAckPackage{
				Status:    types.StatusFail,
				Operator:  updateGroupPackage.Operator,
				ExtraData: updateGroupPackage.ExtraData,
			}.MustSerialize(),
			Err: err,
		}
	}

	return sdk.ExecuteResult{
		Payload: types.UpdateGroupMemberAckPackage{
			Status:        types.StatusSuccess,
			ID:            groupInfo.Id.BigInt(),
			Operator:      updateGroupPackage.Operator,
			OperationType: updateGroupPackage.OperationType,
			Members:       updateGroupPackage.Members,
			ExtraData:     updateGroupPackage.ExtraData,
		}.MustSerialize(),
	}
}

func (app *GroupApp) handleAddOrDeleteGroupMemberOperation(ctx sdk.Context, sourceType types.SourceType, groupInfo *types.GroupInfo, updateGroupPackage *types.UpdateGroupMemberSynPackage) error {
	options := types.UpdateGroupMemberOptions{
		SourceType: sourceType,
	}
	if updateGroupPackage.OperationType == types.OperationAddGroupMember {
		options.MembersToAdd = updateGroupPackage.GetMembers()
		options.MembersExpirationToAdd = updateGroupPackage.GetMemberExpiration()
	} else {
		options.MembersToDelete = updateGroupPackage.GetMembers()
	}

	return app.storageKeeper.UpdateGroupMember(
		ctx,
		updateGroupPackage.Operator,
		groupInfo,
		options,
	)
}

func (app *GroupApp) handleRenewGroupOperation(ctx sdk.Context, sourceType types.SourceType, groupInfo *types.GroupInfo, updateGroupPackage *types.UpdateGroupMemberSynPackage) error {
	options := types.RenewGroupMemberOptions{
		SourceType:        sourceType,
		Members:           updateGroupPackage.GetMembers(),
		MembersExpiration: updateGroupPackage.GetMemberExpiration(),
	}

	return app.storageKeeper.RenewGroupMember(
		ctx,
		updateGroupPackage.Operator,
		groupInfo,
		options,
	)
}

//nolint:unparam
func (app *GroupApp) handleUpdateGroupMemberAckPackage(ctx sdk.Context, _ *sdk.CrossChainAppContext, _ *types.UpdateGroupMemberAckPackage) sdk.ExecuteResult {
	app.storageKeeper.Logger(ctx).Error("received update group member ack package ")

	return sdk.ExecuteResult{}
}

//nolint:unparam
func (app *GroupApp) handleUpdateGroupMemberFailAckPackage(ctx sdk.Context, _ *sdk.CrossChainAppContext, _ *types.UpdateGroupMemberSynPackage) sdk.ExecuteResult {
	app.storageKeeper.Logger(ctx).Error("received update group member fail ack package ")

	return sdk.ExecuteResult{}
}
