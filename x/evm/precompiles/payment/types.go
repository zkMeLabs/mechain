package payment

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	paymentAddress = common.HexToAddress(types.PaymentAddress)
	paymentABI     = types.MustABIJson(IPaymentMetaData.ABI)
	invalidMethod  = abi.Method{}
)

type (
	PageRequestJSON = PageRequest
)

func GetAddress() common.Address {
	return paymentAddress
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return invalidMethod, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range paymentABI.Methods {
		if bytes.Equal(input[:4], method.ID) {
			return method, nil
		}
	}
	return invalidMethod, fmt.Errorf("method id %s is not exist", string(input[:4]))
}

func GetAbiMethod(name string) abi.Method {
	return paymentABI.Methods[name]
}

func GetAbiEvent(name string) abi.Event {
	return paymentABI.Events[name]
}

type DepositArgs struct {
	// Creator string   `abi:"creator"`
	To     string   `abi:"to"`
	Amount *big.Int `abi:"amount"`
}

// Validate DepositArgs the args
func (args *DepositArgs) Validate() error {
	return nil
}

type DisableRefundArgs struct {
	// Owner string `abi:"owner"`
	Addr string `abi:"addr"`
}

// Validate DisableRefundArgs the args
func (args *DisableRefundArgs) Validate() error {
	return nil
}

type WithdrawArgs struct {
	// Creator string   `abi:"creator"`
	From   string   `abi:"from"`
	Amount *big.Int `abi:"amount"`
}

// Validate WithdrawArgs the args
func (args *WithdrawArgs) Validate() error {
	return nil
}

type PaymentAccountsByOwnerArgs struct {
	Owner string `abi:"owner"`
}

// Validate PaymentAccountsByOwnerArgs the args
func (args *PaymentAccountsByOwnerArgs) Validate() error {
	return nil
}

type PaymentAccountArgs struct {
	Addr string `abi:"addr"`
}

// Validate PaymentAccountArgs the args
func (args *PaymentAccountArgs) Validate() error {
	return nil
}

type UpdateParamsArgs struct {
	Authority string `abi:"authority"`
	Params    Params `abi:"params"`
}

// Validate UpdateParamsArgs the args
func (args *UpdateParamsArgs) Validate() error {
	return nil
}

type ParamsByTimestampArgs struct {
	Timestamp int64 `abi:"timestamp"`
}

// Validate ParamsByTimestampArgs the args
func (args *ParamsByTimestampArgs) Validate() error {
	return nil
}

type OutFlowsArgs struct {
	Account string `abi:"account"`
}

// Validate OutFlowsArgs the args
func (args *OutFlowsArgs) Validate() error {
	return nil
}

type StreamRecordArgs struct {
	Account string `abi:"account"`
}

// Validate StreamRecordArgs the args
func (args *StreamRecordArgs) Validate() error {
	return nil
}

type StreamRecordsArgs struct {
	Pagination PageRequestJSON `abi:"pagination"`
}

// Validate StreamRecordsArgs the args
func (args *StreamRecordsArgs) Validate() error {
	return nil
}

type PaymentAccountCountArgs struct {
	Owner string `abi:"owner"`
}

// Validate PaymentAccountCountArgs the args
func (args *PaymentAccountCountArgs) Validate() error {
	return nil
}

type PaymentAccountCountsArgs struct {
	Pagination PageRequestJSON `abi:"pagination"`
}

// Validate PaymentAccountCountsArgs the args
func (args *PaymentAccountCountsArgs) Validate() error {
	return nil
}

type PaymentAccountsArgs struct {
	Pagination PageRequestJSON `abi:"pagination"`
}

// Validate PaymentAccountsArgs the args
func (args *PaymentAccountsArgs) Validate() error {
	return nil
}

type DynamicBalanceArgs struct {
	Account string `abi:"account"`
}

// Validate DynamicBalanceArgs the args
func (args *DynamicBalanceArgs) Validate() error {
	return nil
}

type AutoSettleRecordsArgs struct {
	Pagination PageRequestJSON `abi:"pagination"`
}

// Validate AutoSettleRecordsArgs the args
func (args *AutoSettleRecordsArgs) Validate() error {
	return nil
}

type DelayedWithdrawalArgs struct {
	Account string `abi:"account"`
}

// Validate DelayedWithdrawalArgs the args
func (args *DelayedWithdrawalArgs) Validate() error {
	return nil
}
