package staking

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/evmos/evmos/v12/types"
)

var (
	stakingAddress = common.HexToAddress(types.StakingAddress)
	stakingABI     = types.MustABIJson(IStakingMetaData.ABI)
)

func GetAddress() common.Address {
	return stakingAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := stakingABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range stakingABI.Methods {
		if bytes.Equal(input[:4], method.ID) {
			return method, nil
		}
	}
	return abi.Method{}, fmt.Errorf("method id %s is not exist", string(input[:4]))
}

func MustMethod(name string) abi.Method {
	method, err := GetMethod(name)
	if err != nil {
		panic(err)
	}
	return method
}

func GetEvent(name string) (abi.Event, error) {
	event := stakingABI.Events[name]
	if event.ID == (common.Hash{}) {
		return abi.Event{}, fmt.Errorf("event %s is not exist", name)
	}
	return event, nil
}

func MustEvent(name string) abi.Event {
	event, err := GetEvent(name)
	if err != nil {
		panic(err)
	}
	return event
}

type (
	DescriptionJson     = Description
	CommissionRatesJson = CommissionRates
	PageRequestJson     = PageRequest
)

type EditValidatorArgs struct {
	Description       DescriptionJson `abi:"description"`
	CommissionRate    *big.Int        `abi:"commissionRate"`
	MinSelfDelegation *big.Int        `abi:"minSelfDelegation"`
	RelayerAddress    common.Address  `abi:"relayerAddress"`
	ChallengerAddress common.Address  `abi:"challengerAddress"`
	BlsKey            string          `abi:"blsKey"`
	BlsProof          string          `abi:"blsProof"`
}

// Validate validates the args
func (args *EditValidatorArgs) Validate() error {
	return nil
}

// GetCommissionRate returns the dec commission rate
func (args *EditValidatorArgs) GetCommissionRate() *sdk.Dec {
	var commissionRate *sdk.Dec
	// if is less than 0, represents the user's unwillingness to modify this value
	if args.CommissionRate.Cmp(big.NewInt(-1)) > 0 {
		tmp := sdk.NewDecFromBigIntWithPrec(args.CommissionRate, sdk.Precision)
		commissionRate = &tmp
	}

	return commissionRate
}

// GetMinSelfDelegation returns the sdk.Int minSelfDelegation
func (args *EditValidatorArgs) GetMinSelfDelegation() *sdk.Int {
	var minSelfDelegation *sdk.Int
	// if is less than 0, represents the user's unwillingness to modify this value
	if args.MinSelfDelegation.Cmp(big.NewInt(-1)) > 0 {
		tmp := sdk.NewIntFromBigInt(args.MinSelfDelegation)
		minSelfDelegation = &tmp
	}

	return minSelfDelegation
}

// GetRelayerAddress returns the relayer address
func (args *EditValidatorArgs) GetRelayerAddress() string {
	if args.RelayerAddress == (common.Address{}) {
		return ""
	}

	return args.RelayerAddress.String()
}

// GetChallengerAddress returns the challenger address
func (args *EditValidatorArgs) GetChallengerAddress() string {
	if args.ChallengerAddress == (common.Address{}) {
		return ""
	}

	return args.ChallengerAddress.String()
}

type DelegateArgs struct {
	ValidatorAddress common.Address `abi:"validatorAddress"`
	Amount           *big.Int       `abi:"amount"`
}

// Validate validates the args
func (args *DelegateArgs) Validate() error {
	if args.ValidatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddress)
	}
	if args.Amount == nil || args.Amount.Sign() <= 0 {
		return errors.New("invalid amount")
	}
	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *DelegateArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddress.Bytes())
	return valAddr
}

type DelegationArgs struct {
	DelegatorAddr common.Address `abi:"delegatorAddr"`
	ValidatorAddr common.Address `abi:"validatorAddr"`
}

// Validate validates the args
func (args *DelegationArgs) Validate() error {
	if args.DelegatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid delegator address: %s", args.DelegatorAddr)
	}
	if args.ValidatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddr)
	}

	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *DelegationArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddr.Bytes())
	return valAddr
}

// GetDelegator returns the Delegator address, caller must ensure the delegator address is valid
func (args *DelegationArgs) GetDelegator() sdk.AccAddress {
	accAddr := sdk.AccAddress(args.DelegatorAddr.Bytes())
	return accAddr
}

type UnbondingDelegationArgs struct {
	DelegatorAddr common.Address `abi:"delegatorAddr"`
	ValidatorAddr common.Address `abi:"validatorAddr"`
}

// Validate validates the args
func (args *UnbondingDelegationArgs) Validate() error {
	if args.DelegatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid delegator address: %s", args.DelegatorAddr)
	}
	if args.ValidatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddr)
	}

	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *UnbondingDelegationArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddr.Bytes())
	return valAddr
}

// GetDelegator returns the Delegator address, caller must ensure the delegator address is valid
func (args *UnbondingDelegationArgs) GetDelegator() sdk.AccAddress {
	accAddr := sdk.AccAddress(args.DelegatorAddr.Bytes())
	return accAddr
}

type UndelegateArgs struct {
	ValidatorAddress common.Address `abi:"validatorAddress"`
	Amount           *big.Int       `abi:"amount"`
}

// Validate validates the args
func (args *UndelegateArgs) Validate() error {
	if args.ValidatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddress)
	}
	if args.Amount == nil || args.Amount.Sign() <= 0 {
		return errors.New("invalid amount")
	}
	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *UndelegateArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddress.Bytes())
	return valAddr
}

type RedelegateArgs struct {
	ValidatorSrcAddress common.Address `abi:"validatorSrcAddress"`
	ValidatorDstAddress common.Address `abi:"validatorDstAddress"`
	Amount              *big.Int       `abi:"amount"`
}

// Validate validates the args
func (args *RedelegateArgs) Validate() error {
	if args.ValidatorSrcAddress == (common.Address{}) {
		return fmt.Errorf("invalid src validator address: %s", args.ValidatorSrcAddress)
	}
	if args.ValidatorDstAddress == (common.Address{}) {
		return fmt.Errorf("invalid dst validator address: %s", args.ValidatorDstAddress)
	}
	if args.Amount == nil || args.Amount.Sign() <= 0 {
		return errors.New("invalid amount")
	}
	return nil
}

// GetSrcValidator returns the validator src address, caller must ensure the validator address is valid
func (args *RedelegateArgs) GetSrcValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorSrcAddress.Bytes())
	return valAddr
}

// GetDstValidator returns the validator dest address, caller must ensure the validator address is valid
func (args *RedelegateArgs) GetDstValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorDstAddress.Bytes())
	return valAddr
}

type CancelUnbondingDelegationArgs struct {
	ValidatorAddress common.Address `abi:"validatorAddress"`
	Amount           *big.Int       `abi:"amount"`
	CreationHeight   *big.Int       `abi:"creationHeight"`
}

// Validate validates the args
func (args *CancelUnbondingDelegationArgs) Validate() error {
	if args.ValidatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddress)
	}

	if args.Amount == nil || args.Amount.Sign() <= 0 {
		return errors.New("invalid amount")
	}

	if args.CreationHeight == nil || args.CreationHeight.Sign() <= 0 || !args.CreationHeight.IsInt64() {
		return errors.New("invalid creation height")
	}

	return nil
}

// GetValidator returns the validator address
func (args *CancelUnbondingDelegationArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddress.Bytes())
	return valAddr
}

// GetCreationHeight returns the creation height
func (args *CancelUnbondingDelegationArgs) GetCreationHeight() int64 {
	return args.CreationHeight.Int64()
}

type ValidatorsArgs struct {
	Status     uint8           `abi:"status"`
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate validates the args
func (args *ValidatorsArgs) Validate() error {
	if args.Status > uint8(stakingtypes.Bonded) {
		return fmt.Errorf("invalid status: %d", args.Status)
	}

	return nil
}

// GetStatus returns the validator status string
func (args *ValidatorsArgs) GetStatus() string {
	switch args.Status {
	case 0:
		return ""
	case 1:
		return stakingtypes.Unbonded.String()
	case 2:
		return stakingtypes.Unbonding.String()
	case 3:
		return stakingtypes.Bonded.String()
	default:
		return ""
	}
	return ""
}

type ValidatorArgs struct {
	ValidatorAddr common.Address `abi:"validatorAddr"`
}

// Validate validates the args
func (args *ValidatorArgs) Validate() error {
	if args.ValidatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddr)
	}

	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *ValidatorArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddr.Bytes())
	return valAddr
}

type ValidatorDelegationsArgs struct {
	ValidatorAddr common.Address  `abi:"validatorAddr"`
	Pagination    PageRequestJson `abi:"pagination"`
}

// Validate validates the args
func (args *ValidatorDelegationsArgs) Validate() error {
	if args.ValidatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddr)
	}

	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *ValidatorDelegationsArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddr.Bytes())
	return valAddr
}

type ValidatorUnbondingDelegationsArgs struct {
	ValidatorAddr common.Address  `abi:"validatorAddr"`
	Pagination    PageRequestJson `abi:"pagination"`
}

// Validate validates the args
func (args *ValidatorUnbondingDelegationsArgs) Validate() error {
	if args.ValidatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddr)
	}

	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *ValidatorUnbondingDelegationsArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddr.Bytes())
	return valAddr
}

type DelegatorDelegationsArgs struct {
	DelegatorAddr common.Address  `abi:"delegatorAddr"`
	Pagination    PageRequestJson `abi:"pagination"`
}

// Validate validates the args
func (args *DelegatorDelegationsArgs) Validate() error {
	if args.DelegatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid delegator address: %s", args.DelegatorAddr)
	}

	return nil
}

// GetDelegator returns the delegator address, caller must ensure the delegator address is valid
func (args *DelegatorDelegationsArgs) GetDelegator() sdk.AccAddress {
	valAddr := sdk.AccAddress(args.DelegatorAddr.Bytes())
	return valAddr
}

type DelegatorUnbondingDelegationsArgs struct {
	DelegatorAddr common.Address  `abi:"delegatorAddr"`
	Pagination    PageRequestJson `abi:"pagination"`
}

// Validate validates the args
func (args *DelegatorUnbondingDelegationsArgs) Validate() error {
	if args.DelegatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid delegator address: %s", args.DelegatorAddr)
	}

	return nil
}

// GetDelegator returns the delegator address, caller must ensure the delegator address is valid
func (args *DelegatorUnbondingDelegationsArgs) GetDelegator() sdk.AccAddress {
	valAddr := sdk.AccAddress(args.DelegatorAddr.Bytes())
	return valAddr
}

type Redelegations struct {
	DelegatorAddr    common.Address  `abi:"delegatorAddr"`
	SrcValidatorAddr common.Address  `abi:"srcValidatorAddr"`
	DstValidatorAddr common.Address  `abi:"dstValidatorAddr"`
	Pagination       PageRequestJson `abi:"pagination"`
}

// Validate validates the args
func (args *Redelegations) Validate() error {
	return nil
}

// GetDelegator returns the delegator address, caller must ensure the delegator address is valid
func (args *Redelegations) GetDelegator() sdk.AccAddress {
	delAddr := sdk.AccAddress(args.DelegatorAddr.Bytes())
	return delAddr
}

// GetSrcValidator returns the src validator address, caller must ensure the validator address is valid
func (args *Redelegations) GetSrcValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.SrcValidatorAddr.Bytes())
	return valAddr
}

// GetDstValidator returns the dst validator address, caller must ensure the validator address is valid
func (args *Redelegations) GetDstValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.DstValidatorAddr.Bytes())
	return valAddr
}

type DelegatorValidators struct {
	DelegatorAddr common.Address  `abi:"delegatorAddr"`
	Pagination    PageRequestJson `abi:"pagination"`
}

// Validate validates the args
func (args *DelegatorValidators) Validate() error {
	if args.DelegatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid delegator address: %s", args.DelegatorAddr)
	}

	return nil
}

// GetDelegator returns the delegator address, caller must ensure the delegator address is valid
func (args *DelegatorValidators) GetDelegator() sdk.AccAddress {
	delAddr := sdk.AccAddress(args.DelegatorAddr.Bytes())
	return delAddr
}

type DelegatorValidator struct {
	DelegatorAddr common.Address `abi:"delegatorAddr"`
	ValidatorAddr common.Address `abi:"validatorAddr"`
}

// Validate validates the args
func (args *DelegatorValidator) Validate() error {
	if args.DelegatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid delegator address: %s", args.DelegatorAddr)
	}

	if args.ValidatorAddr == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddr)
	}

	return nil
}

// GetDelegator returns the delegator address, caller must ensure the delegator address is valid
func (args *DelegatorValidator) GetDelegator() sdk.AccAddress {
	delAddr := sdk.AccAddress(args.DelegatorAddr.Bytes())
	return delAddr
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *DelegatorValidator) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddr.Bytes())
	return valAddr
}

type HistoricalInfoRequest struct {
	Height int64 `abi:"height"`
}

// Validate validates the args
func (args *HistoricalInfoRequest) Validate() error {
	if args.Height < 0 {
		return fmt.Errorf("invalid height: %v", args.Height)
	}

	return nil
}

// GetHeight returns the block height, caller must ensure the block height is valid
func (args *HistoricalInfoRequest) GetHeight() int64 {
	return args.Height
}
