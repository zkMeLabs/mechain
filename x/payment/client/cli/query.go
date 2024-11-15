package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/spf13/cobra"

	"github.com/evmos/evmos/v12/x/evm/precompiles/payment"
	"github.com/evmos/evmos/v12/x/payment/types"
)

func ToPaymentPageReq(in *query.PageRequest) *payment.PageRequest {
	if in == nil {
		return nil
	}
	return &payment.PageRequest{
		Key:        in.Key,
		Offset:     in.Offset,
		Limit:      in.Limit,
		CountTotal: in.CountTotal,
		Reverse:    in.Reverse,
	}
}

func ToPageResp(p *payment.PageResponse) *query.PageResponse {
	if p == nil {
		return nil
	}
	return &query.PageResponse{
		NextKey: p.NextKey,
		Total:   p.Total,
	}
}

func ToStreamRecord(p *payment.StreamRecord) *types.StreamRecord {
	if p == nil {
		return nil
	}
	s := &types.StreamRecord{
		Account:           p.Account,
		CrudTimestamp:     p.CrudTimestamp,
		NetflowRate:       sdktypes.NewIntFromBigInt(p.NetflowRate),
		StaticBalance:     sdktypes.NewIntFromBigInt(p.StaticBalance),
		BufferBalance:     sdktypes.NewIntFromBigInt(p.BufferBalance),
		LockBalance:       sdktypes.NewIntFromBigInt(p.LockBalance),
		Status:            types.StreamAccountStatus(p.Status),
		SettleTimestamp:   p.SettleTimestamp,
		OutFlowCount:      p.OutFlowCount,
		FrozenNetflowRate: sdktypes.NewIntFromBigInt(p.FrozenNetflowRate),
	}
	return s
}

func ToPaymentAccount(p *payment.PaymentAccount) *types.PaymentAccount {
	if p == nil {
		return nil
	}
	s := &types.PaymentAccount{
		Addr:       p.Addr,
		Owner:      p.Owner,
		Refundable: p.Refundable,
	}
	return s
}

func ToPaymentAccountCount(p *payment.PaymentAccountCount) *types.PaymentAccountCount {
	if p == nil {
		return nil
	}
	s := &types.PaymentAccountCount{
		Owner: p.Owner,
		Count: p.Count,
	}
	return s
}

func ToAutoSettleRecord(p *payment.AutoSettleRecord) *types.AutoSettleRecord {
	if p == nil {
		return nil
	}
	s := &types.AutoSettleRecord{
		Timestamp: p.Timestamp,
		Addr:      p.Addr,
	}
	return s
}

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group payment queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListStreamRecord())
	cmd.AddCommand(CmdShowStreamRecord())
	cmd.AddCommand(CmdListPaymentAccountCount())
	cmd.AddCommand(CmdShowPaymentAccountCount())
	cmd.AddCommand(CmdListPaymentAccount())
	cmd.AddCommand(CmdShowPaymentAccount())
	cmd.AddCommand(CmdDynamicBalance())
	cmd.AddCommand(CmdGetPaymentAccountsByOwner())
	cmd.AddCommand(CmdListAutoSettleRecord())

	return cmd
}
