package keeper

import (
	"encoding/hex"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/evmos/evmos/v12/x/storage/types"
)

var _ sdk.CrossChainApplication = &ObjectApp{}

type ObjectApp struct {
	storageKeeper types.StorageKeeper
}

func NewObjectApp(keeper types.StorageKeeper) *ObjectApp {
	return &ObjectApp{
		storageKeeper: keeper,
	}
}

func (app *ObjectApp) ExecuteAckPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, payload []byte) sdk.ExecuteResult {
	pack, err := types.DeserializeCrossChainPackage(payload, types.ObjectChannelID, sdk.AckCrossChainPackageType)
	if err != nil {
		app.storageKeeper.Logger(ctx).Error("deserialize object cross chain package error", "payload", hex.EncodeToString(payload), "error", err.Error())
		panic("deserialize object cross chain package error")
	}

	var operationType uint8
	var result sdk.ExecuteResult

	switch p := pack.(type) {
	case *types.MirrorObjectAckPackage:
		operationType = types.OperationMirrorObject
		result = app.handleMirrorObjectAckPackage(ctx, appCtx, p)
	case *types.DeleteObjectAckPackage:
		operationType = types.OperationDeleteObject
		result = app.handleDeleteObjectAckPackage(ctx, appCtx, p)
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

func (app *ObjectApp) ExecuteFailAckPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, payload []byte) sdk.ExecuteResult {
	pack, err := types.DeserializeCrossChainPackage(payload, types.ObjectChannelID, sdk.FailAckCrossChainPackageType)
	if err != nil {
		app.storageKeeper.Logger(ctx).Error("deserialize object cross chain package error", "payload", hex.EncodeToString(payload), "error", err.Error())
		panic("deserialize object cross chain package error")
	}

	var operationType uint8
	var result sdk.ExecuteResult
	switch p := pack.(type) {
	case *types.MirrorObjectSynPackage:
		operationType = types.OperationMirrorObject
		result = app.handleMirrorObjectFailAckPackage(ctx, appCtx, p)
	case *types.DeleteObjectSynPackage:
		operationType = types.OperationDeleteObject
		result = app.handleDeleteObjectFailAckPackage(ctx, appCtx, p)
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

func (app *ObjectApp) ExecuteSynPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, payload []byte) sdk.ExecuteResult {
	pack, err := types.DeserializeCrossChainPackage(payload, types.ObjectChannelID, sdk.SynCrossChainPackageType)
	if err != nil {
		app.storageKeeper.Logger(ctx).Error("deserialize object cross chain package error", "payload", hex.EncodeToString(payload), "error", err.Error())
		panic("deserialize object cross chain package error")
	}

	var operationType uint8
	var result sdk.ExecuteResult
	switch p := pack.(type) {
	case *types.MirrorObjectSynPackage:
		operationType = types.OperationMirrorObject
		result = app.handleMirrorObjectSynPackage(ctx, appCtx, p)
	case *types.DeleteObjectSynPackage:
		operationType = types.OperationDeleteObject
		result = app.handleDeleteObjectSynPackage(ctx, appCtx, p)
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

func (app *ObjectApp) handleMirrorObjectAckPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, ackPackage *types.MirrorObjectAckPackage) sdk.ExecuteResult {
	app.storageKeeper.Logger(ctx).Error("received mirror object ack package ")

	objectInfo, found := app.storageKeeper.GetObjectInfoById(ctx, math.NewUintFromBigInt(ackPackage.ID))
	if !found {
		app.storageKeeper.Logger(ctx).Error("object does not exist", "object id", ackPackage.ID.String())
		return sdk.ExecuteResult{
			Err: types.ErrNoSuchObject,
		}
	}

	sourceType, err := app.storageKeeper.GetSourceTypeByChainId(ctx, appCtx.SrcChainId)
	if err != nil {
		return sdk.ExecuteResult{
			Err: err,
		}
	}

	// update object
	if ackPackage.Status == types.StatusSuccess {
		objectInfo.SourceType = sourceType

		app.storageKeeper.SetObjectInfo(ctx, objectInfo)
	}

	if err := ctx.EventManager().EmitTypedEvents(&types.EventMirrorObjectResult{
		Status:      uint32(ackPackage.Status),
		BucketName:  objectInfo.BucketName,
		ObjectName:  objectInfo.ObjectName,
		ObjectId:    objectInfo.Id,
		DestChainId: uint32(appCtx.SrcChainId),
	}); err != nil {
		return sdk.ExecuteResult{
			Err: err,
		}
	}

	return sdk.ExecuteResult{}
}

func (app *ObjectApp) handleMirrorObjectFailAckPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, mirrorObjectPackage *types.MirrorObjectSynPackage) sdk.ExecuteResult {
	app.storageKeeper.Logger(ctx).Error("received mirror object fail ack package ")

	objectInfo, found := app.storageKeeper.GetObjectInfoById(ctx, math.NewUintFromBigInt(mirrorObjectPackage.ID))
	if !found {
		app.storageKeeper.Logger(ctx).Error("object does not exist", "object id", mirrorObjectPackage.ID.String())
		return sdk.ExecuteResult{
			Err: types.ErrNoSuchObject,
		}
	}

	objectInfo.SourceType = types.SOURCE_TYPE_ORIGIN
	app.storageKeeper.SetObjectInfo(ctx, objectInfo)

	if err := ctx.EventManager().EmitTypedEvents(&types.EventMirrorObjectResult{
		Status:      uint32(types.StatusFail),
		BucketName:  objectInfo.BucketName,
		ObjectName:  objectInfo.ObjectName,
		ObjectId:    objectInfo.Id,
		DestChainId: uint32(appCtx.SrcChainId),
	}); err != nil {
		return sdk.ExecuteResult{
			Err: err,
		}
	}

	return sdk.ExecuteResult{}
}

func (app *ObjectApp) handleMirrorObjectSynPackage(_ sdk.Context, _ *sdk.CrossChainAppContext, _ *types.MirrorObjectSynPackage) sdk.ExecuteResult {
	return sdk.ExecuteResult{}
}

func (app *ObjectApp) handleDeleteObjectSynPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, deleteObjectPackage *types.DeleteObjectSynPackage) sdk.ExecuteResult {
	err := deleteObjectPackage.ValidateBasic()
	if err != nil {
		return sdk.ExecuteResult{
			Payload: types.DeleteObjectAckPackage{
				Status:    types.StatusFail,
				ExtraData: deleteObjectPackage.ExtraData,
			}.MustSerialize(),
			Err: err,
		}
	}

	app.storageKeeper.Logger(ctx).Info("process delete object syn package", "object id", deleteObjectPackage.ID.String())

	objectInfo, found := app.storageKeeper.GetObjectInfoById(ctx, math.NewUintFromBigInt(deleteObjectPackage.ID))
	if !found {
		return sdk.ExecuteResult{
			Payload: types.DeleteObjectAckPackage{
				Status:    types.StatusFail,
				ExtraData: deleteObjectPackage.ExtraData,
			}.MustSerialize(),
			Err: types.ErrNoSuchObject,
		}
	}

	sourceType, err := app.storageKeeper.GetSourceTypeByChainId(ctx, appCtx.SrcChainId)
	if err != nil {
		return sdk.ExecuteResult{
			Err: err,
		}
	}

	err = app.storageKeeper.DeleteObject(ctx,
		deleteObjectPackage.Operator,
		objectInfo.BucketName,
		objectInfo.ObjectName,
		types.DeleteObjectOptions{
			SourceType: sourceType,
		},
	)
	if err != nil {
		return sdk.ExecuteResult{
			Payload: types.DeleteObjectAckPackage{
				Status:    types.StatusFail,
				ExtraData: deleteObjectPackage.ExtraData,
			}.MustSerialize(),
			Err: err,
		}
	}

	return sdk.ExecuteResult{
		Payload: types.DeleteObjectAckPackage{
			Status:    types.StatusSuccess,
			ID:        objectInfo.Id.BigInt(),
			ExtraData: deleteObjectPackage.ExtraData,
		}.MustSerialize(),
	}
}

//nolint:unparam
func (app *ObjectApp) handleDeleteObjectAckPackage(ctx sdk.Context, _ *sdk.CrossChainAppContext, _ *types.DeleteObjectAckPackage) sdk.ExecuteResult {
	app.storageKeeper.Logger(ctx).Error("received delete object ack package ")

	return sdk.ExecuteResult{}
}

//nolint:unparam
func (app *ObjectApp) handleDeleteObjectFailAckPackage(ctx sdk.Context, _ *sdk.CrossChainAppContext, _ *types.DeleteObjectSynPackage) sdk.ExecuteResult {
	app.storageKeeper.Logger(ctx).Error("received delete object fail ack package ")

	return sdk.ExecuteResult{}
}
