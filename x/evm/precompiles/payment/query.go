package payment

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
	paymenttypes "github.com/evmos/evmos/v12/x/payment/types"
)

const (
	PaymentAccountsByOwnerMethodName = "paymentAccountsByOwner"
	PaymentAccountMethodName         = "paymentAccount"
	ParamsMethodName                 = "params"
	ParamsByTimestampMethodName      = "paramsByTimestamp"
	OutFlowsMethodName               = "outFlows"
	StreamRecordMethodName           = "streamRecord"
	StreamRecordsMethodName          = "streamRecords"
	PaymentAccountCountMethodName    = "paymentAccountCount"
	PaymentAccountCountsMethodName   = "paymentAccountCounts"
	PaymentAccountsMethodName        = "paymentAccounts"
	DynamicBalanceMethodName         = "dynamicBalance"
	AutoSettleRecordsMethodName      = "autoSettleRecords"
	DelayedWithdrawalMethodName      = "delayedWithdrawal"
)

func (c *Contract) registerQuery() {
	c.registerMethod(PaymentAccountsByOwnerMethodName, 50_000, c.PaymentAccountsByOwner, "")
	c.registerMethod(PaymentAccountMethodName, 50_000, c.PaymentAccount, "")
	c.registerMethod(ParamsMethodName, 50_000, c.Params, "")
	c.registerMethod(ParamsByTimestampMethodName, 50_000, c.ParamsByTimestamp, "")
	c.registerMethod(OutFlowsMethodName, 50_000, c.OutFlows, "")
	c.registerMethod(StreamRecordMethodName, 50_000, c.StreamRecord, "")
	c.registerMethod(StreamRecordsMethodName, 50_000, c.StreamRecords, "")
	c.registerMethod(PaymentAccountCountMethodName, 50_000, c.PaymentAccountCount, "")
	c.registerMethod(PaymentAccountCountsMethodName, 50_000, c.PaymentAccountCounts, "")
	c.registerMethod(PaymentAccountsMethodName, 50_000, c.PaymentAccounts, "")
	c.registerMethod(DynamicBalanceMethodName, 50_000, c.DynamicBalance, "")
	c.registerMethod(AutoSettleRecordsMethodName, 50_000, c.AutoSettleRecords, "")
	c.registerMethod(DelayedWithdrawalMethodName, 50_000, c.DelayedWithdrawal, "")
}

// PaymentAccountsByOwner queries all payment accounts by a owner.
func (c *Contract) PaymentAccountsByOwner(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(PaymentAccountsByOwnerMethodName)
	// parse args
	var args PaymentAccountsByOwnerArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryPaymentAccountsByOwnerRequest{
		Owner: args.Owner,
	}
	res, err := c.paymentKeeper.PaymentAccountsByOwner(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.PaymentAccounts)
}

// PaymentAccount queries a payment account by payment account address.
func (c *Contract) PaymentAccount(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(PaymentAccountMethodName)
	var args PaymentAccountArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryPaymentAccountRequest{
		Addr: args.Addr,
	}
	res, err := c.paymentKeeper.PaymentAccount(ctx, msg)
	if err != nil {
		return nil, err
	}
	paymentAccount := PaymentAccount{
		Addr:       res.PaymentAccount.Addr,
		Owner:      res.PaymentAccount.Owner,
		Refundable: res.PaymentAccount.Refundable,
	}

	return method.Outputs.Pack(paymentAccount)
}

// Parameters queries the parameters of the module.
func (c *Contract) Params(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(ParamsMethodName)

	msg := &paymenttypes.QueryParamsRequest{}
	res, err := c.paymentKeeper.Params(ctx, msg)
	if err != nil {
		return nil, err
	}
	params := Params{
		VersionedParams: VersionedParams{
			ReserveTime:      res.Params.VersionedParams.ReserveTime,
			ValidatorTaxRate: res.Params.VersionedParams.ValidatorTaxRate.BigInt(),
		},
		PaymentAccountCountLimit:  res.Params.PaymentAccountCountLimit,
		ForcedSettleTime:          res.Params.ForcedSettleTime,
		MaxAutoSettleFlowCount:    res.Params.MaxAutoSettleFlowCount,
		MaxAutoResumeFlowCount:    res.Params.MaxAutoResumeFlowCount,
		FeeDenom:                  res.Params.FeeDenom,
		WithdrawTimeLockThreshold: res.Params.WithdrawTimeLockThreshold.BigInt(),
		WithdrawTimeLockDuration:  res.Params.WithdrawTimeLockDuration,
	}

	return method.Outputs.Pack(params)
}

// ParamsByTimestamp queries the parameter of the module by timestamp.
func (c *Contract) ParamsByTimestamp(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(ParamsByTimestampMethodName)
	var args ParamsByTimestampArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryParamsByTimestampRequest{
		Timestamp: args.Timestamp,
	}
	res, err := c.paymentKeeper.ParamsByTimestamp(ctx, msg)
	if err != nil {
		return nil, err
	}
	params := Params{
		VersionedParams: VersionedParams{
			ReserveTime:      res.Params.VersionedParams.ReserveTime,
			ValidatorTaxRate: res.Params.VersionedParams.ValidatorTaxRate.BigInt(),
		},
		PaymentAccountCountLimit:  res.Params.PaymentAccountCountLimit,
		ForcedSettleTime:          res.Params.ForcedSettleTime,
		MaxAutoSettleFlowCount:    res.Params.MaxAutoSettleFlowCount,
		MaxAutoResumeFlowCount:    res.Params.MaxAutoResumeFlowCount,
		FeeDenom:                  res.Params.FeeDenom,
		WithdrawTimeLockThreshold: res.Params.WithdrawTimeLockThreshold.BigInt(),
		WithdrawTimeLockDuration:  res.Params.WithdrawTimeLockDuration,
	}

	return method.Outputs.Pack(params)
}

// Queries out flows by account.
func (c *Contract) OutFlows(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(OutFlowsMethodName)
	var args OutFlowsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryOutFlowsRequest{
		Account: args.Account,
	}
	res, err := c.paymentKeeper.OutFlows(ctx, msg)
	if err != nil {
		return nil, err
	}
	outFlows := make([]OutFlow, 0)
	for _, outFlow := range res.OutFlows {
		outFlows = append(outFlows, OutFlow{ToAddress: outFlow.ToAddress, Rate: outFlow.Rate.BigInt(), Status: int32(outFlow.Status)})
	}

	return method.Outputs.Pack(outFlows)
}

// Queries a stream record by account.
func (c *Contract) StreamRecord(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(StreamRecordMethodName)
	var args StreamRecordArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryGetStreamRecordRequest{
		Account: args.Account,
	}
	res, err := c.paymentKeeper.StreamRecord(ctx, msg)
	if err != nil {
		return nil, err
	}
	streamRecord := StreamRecord{
		Account:           res.StreamRecord.Account,
		CrudTimestamp:     res.StreamRecord.CrudTimestamp,
		NetflowRate:       res.StreamRecord.NetflowRate.BigInt(),
		StaticBalance:     res.StreamRecord.StaticBalance.BigInt(),
		BufferBalance:     res.StreamRecord.BufferBalance.BigInt(),
		LockBalance:       res.StreamRecord.LockBalance.BigInt(),
		Status:            int32(res.StreamRecord.Status),
		SettleTimestamp:   res.StreamRecord.SettleTimestamp,
		OutFlowCount:      res.StreamRecord.OutFlowCount,
		FrozenNetflowRate: res.StreamRecord.FrozenNetflowRate.BigInt(),
	}

	return method.Outputs.Pack(streamRecord)
}

// Queries all stream records.
func (c *Contract) StreamRecords(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(StreamRecordsMethodName)
	var args StreamRecordsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryStreamRecordsRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}
	res, err := c.paymentKeeper.StreamRecords(ctx, msg)
	if err != nil {
		return nil, err
	}
	streamRecords := make([]StreamRecord, 0, len(res.StreamRecords))
	for _, streamRecord := range res.StreamRecords {
		streamRecords = append(streamRecords, StreamRecord{
			Account:           streamRecord.Account,
			CrudTimestamp:     streamRecord.CrudTimestamp,
			NetflowRate:       streamRecord.NetflowRate.BigInt(),
			StaticBalance:     streamRecord.StaticBalance.BigInt(),
			BufferBalance:     streamRecord.BufferBalance.BigInt(),
			LockBalance:       streamRecord.LockBalance.BigInt(),
			Status:            int32(streamRecord.Status),
			SettleTimestamp:   streamRecord.SettleTimestamp,
			OutFlowCount:      streamRecord.OutFlowCount,
			FrozenNetflowRate: streamRecord.FrozenNetflowRate.BigInt(),
		})
	}
	return method.Outputs.Pack(streamRecords, outputPageResponse(res.Pagination))
}

// Queries the count of payment account by owner.
func (c *Contract) PaymentAccountCount(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(PaymentAccountCountMethodName)
	var args PaymentAccountCountArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryPaymentAccountCountRequest{
		Owner: args.Owner,
	}
	res, err := c.paymentKeeper.PaymentAccountCount(ctx, msg)
	if err != nil {
		return nil, err
	}
	paymentAccountCount := PaymentAccountCount{
		Owner: res.PaymentAccountCount.Owner,
		Count: res.PaymentAccountCount.Count,
	}

	return method.Outputs.Pack(paymentAccountCount)
}

// Queries all counts of payment account for all owners.
func (c *Contract) PaymentAccountCounts(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(PaymentAccountCountsMethodName)
	var args PaymentAccountCountsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryPaymentAccountCountsRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}
	res, err := c.paymentKeeper.PaymentAccountCounts(ctx, msg)
	if err != nil {
		return nil, err
	}
	paymentAccountCounts := make([]PaymentAccountCount, 0, len(res.PaymentAccountCounts))
	for _, paymentAccountCount := range res.PaymentAccountCounts {
		paymentAccountCounts = append(paymentAccountCounts, PaymentAccountCount{
			Owner: paymentAccountCount.Owner,
			Count: paymentAccountCount.Count,
		})
	}
	return method.Outputs.Pack(paymentAccountCounts, outputPageResponse(res.Pagination))
}

// Queries all payment accounts.
func (c *Contract) PaymentAccounts(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(PaymentAccountsMethodName)
	var args PaymentAccountsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryPaymentAccountsRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}
	res, err := c.paymentKeeper.PaymentAccounts(ctx, msg)
	if err != nil {
		return nil, err
	}
	paymentAccounts := make([]PaymentAccount, 0, len(res.PaymentAccounts))
	for _, paymentAccount := range res.PaymentAccounts {
		paymentAccounts = append(paymentAccounts, PaymentAccount{
			Addr:       paymentAccount.Addr,
			Owner:      paymentAccount.Owner,
			Refundable: paymentAccount.Refundable,
		})
	}
	return method.Outputs.Pack(paymentAccounts, outputPageResponse(res.Pagination))
}

// Queries dynamic balance of a payment account.
func (c *Contract) DynamicBalance(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(DynamicBalanceMethodName)
	var args DynamicBalanceArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryDynamicBalanceRequest{
		Account: args.Account,
	}
	res, err := c.paymentKeeper.DynamicBalance(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.DynamicBalance.BigInt())
}

// Queries all auto settle records.
func (c *Contract) AutoSettleRecords(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(AutoSettleRecordsMethodName)
	var args AutoSettleRecordsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryAutoSettleRecordsRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}
	res, err := c.paymentKeeper.AutoSettleRecords(ctx, msg)
	if err != nil {
		return nil, err
	}
	autoSettleRecords := make([]AutoSettleRecord, 0, len(res.AutoSettleRecords))
	for _, autoSettleRecord := range res.AutoSettleRecords {
		autoSettleRecords = append(autoSettleRecords, AutoSettleRecord{
			Timestamp: autoSettleRecord.Timestamp,
			Addr:      autoSettleRecord.Addr,
		})
	}
	return method.Outputs.Pack(autoSettleRecords, outputPageResponse(res.Pagination))
}

// Queries delayed withdrawal of a account.
func (c *Contract) DelayedWithdrawal(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(DelayedWithdrawalMethodName)
	var args DelayedWithdrawalArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &paymenttypes.QueryDelayedWithdrawalRequest{
		Account: args.Account,
	}
	res, err := c.paymentKeeper.DelayedWithdrawal(ctx, msg)
	if err != nil {
		return nil, err
	}
	delayedWithdrawal := DelayedWithdrawalRecord{
		Addr:            res.DelayedWithdrawal.Addr,
		Amount:          res.DelayedWithdrawal.Amount.BigInt(),
		From:            res.DelayedWithdrawal.From,
		UnlockTimestamp: res.DelayedWithdrawal.UnlockTimestamp,
	}

	return method.Outputs.Pack(delayedWithdrawal)
}

func outputPageResponse(p *query.PageResponse) *PageResponse {
	return &PageResponse{
		NextKey: p.NextKey,
		Total:   p.Total,
	}
}
