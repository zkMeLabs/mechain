package keeper

import (
	// "encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/evmos/evmos/v12/x/storage/types"
)

var _ sdk.CrossChainApplication = &ZkmeSBTApp{}

type ZkmeSBTApp struct {
	storageKeeper types.StorageKeeper
}

func NewZkmeSBTApp(keeper types.StorageKeeper) *ZkmeSBTApp {
	return &ZkmeSBTApp{
		storageKeeper: keeper,
	}
}

func (app *ZkmeSBTApp) ExecuteAckPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, payload []byte) sdk.ExecuteResult {
	// pack, err := types.DeserializeCrossChainPackage(payload, types.ZkmeSBTChannelId, sdk.AckCrossChainPackageType)
	// if err != nil {
	// 	app.storageKeeper.Logger(ctx).Error("deserialize zkmesbt cross chain package error", "payload", hex.EncodeToString(payload), "error", err.Error())
	// 	panic("deserialize zkmesbt cross chain package error")
	// }

	// var operationType uint8
	// var result sdk.ExecuteResult
	// switch p := pack.(type) {
	// case *types.ZkmeSBTAckPackage:
	// 	operationType = types.OperationMirrorGroup
	// 	result = app.handleZkmeSBTAckPackage(ctx, appCtx, p)
	// default:
	// 	panic("unknown cross chain ack package type")
	// }

	// if len(result.Payload) != 0 {
	// 	wrapPayload := types.CrossChainPackage{
	// 		OperationType: operationType,
	// 		Package:       result.Payload,
	// 	}
	// 	result.Payload = wrapPayload.MustSerialize()
	// }

	var result sdk.ExecuteResult
	return result
}

func (app *ZkmeSBTApp) ExecuteFailAckPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, payload []byte) sdk.ExecuteResult {
	// pack, err := types.DeserializeCrossChainPackage(payload, types.ZkmeSBTChannelId, sdk.FailAckCrossChainPackageType)
	// if err != nil {
	// 	app.storageKeeper.Logger(ctx).Error("deserialize zkmesbt cross chain package error", "payload", hex.EncodeToString(payload), "error", err.Error())
	// 	panic("deserialize zkmesbt cross chain package error")
	// }

	// var operationType uint8
	// var result sdk.ExecuteResult
	// switch p := pack.(type) {
	// case *types.ZkmeSBTSynPackage:
	// 	operationType = types.OperationMirrorGroup
	// 	result = app.handleZkmeSBTFailAckPackage(ctx, appCtx, p)
	// default:
	// 	panic("unknown cross chain ack package type")
	// }

	// if len(result.Payload) != 0 {
	// 	wrapPayload := types.CrossChainPackage{
	// 		OperationType: operationType,
	// 		Package:       result.Payload,
	// 	}
	// 	result.Payload = wrapPayload.MustSerialize()
	// }

	var result sdk.ExecuteResult
	return result
}

func (app *ZkmeSBTApp) ExecuteSynPackage(ctx sdk.Context, appCtx *sdk.CrossChainAppContext, payload []byte) sdk.ExecuteResult {
	var result sdk.ExecuteResult
	return result
}
