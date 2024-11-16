package distribution

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	distributionAddress = common.HexToAddress(types.DistributionAddress)
	distributionABI     = types.MustABIJson(IDistributionMetaData.ABI)
)

type (
	PageRequestJson = PageRequest
)

func GetAddress() common.Address {
	return distributionAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := distributionABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range distributionABI.Methods {
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
	event := distributionABI.Events[name]
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

type SetWithdrawAddressArgs struct {
	WithdrawAddress common.Address `abi:"withdrawAddress"`
}

// Validate validates the args
func (args *SetWithdrawAddressArgs) Validate() error {
	if args.WithdrawAddress == (common.Address{}) {
		return fmt.Errorf("invalid withdraw address: %s", args.WithdrawAddress)
	}
	return nil
}

type DelegationRewardsArgs struct {
	DelegatorAddress common.Address `abi:"delegatorAddress"`
	ValidatorAddress common.Address `abi:"validatorAddress"`
}

// Validate DelegationRewardsArgs the args
func (args *DelegationRewardsArgs) Validate() error {
	if args.DelegatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid delegator address: %s", args.DelegatorAddress)
	}
	if args.ValidatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddress)
	}
	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *DelegationRewardsArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddress.Bytes())
	return valAddr
}

// GetDelegator returns the delegator address, caller must ensure the delegator address is valid
func (args *DelegationRewardsArgs) GetDelegator() sdk.AccAddress {
	accAddr := sdk.AccAddress(args.DelegatorAddress.Bytes())
	return accAddr
}

type ValidatorAddressArgs struct {
	ValidatorAddress common.Address `abi:"validatorAddress"`
}

// Validate WithdrawDelegatorRewardArgs the args
func (args *ValidatorAddressArgs) Validate() error {
	if args.ValidatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddress)
	}
	return nil
}

type DelegatorAddressArgs struct {
	DelegatorAddress common.Address `abi:"delegatorAddress"`
}

// Validate WithdrawDelegatorRewardArgs the args
func (args *DelegatorAddressArgs) Validate() error {
	if args.DelegatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid delegator address: %s", args.DelegatorAddress)
	}
	return nil
}

// GetDelegator returns the delegator address, caller must ensure the delegator address is valid
func (args *DelegatorAddressArgs) GetDelegator() sdk.AccAddress {
	valAddr := sdk.AccAddress(args.DelegatorAddress.Bytes())
	return valAddr
}

type CoinJson = Coin

type FundCommunityPoolArgs struct {
	Amount []CoinJson `abi:"amount"`
}

// Validate FundCommunityPoolrArgs args
func (args *FundCommunityPoolArgs) Validate() error {
	if len(args.Amount) == 0 {
		return fmt.Errorf("no coin send")
	}

	for _, coin := range args.Amount {
		if coin.Amount.Sign() <= 0 {
			return fmt.Errorf("coin amount is %s  less than or equal 0", coin.Amount.String())
		}
	}

	return nil
}

type ValidatorSlashesArgs struct {
	ValidatorAddress common.Address  `abi:"validatorAddress"`
	StartingHeight   uint64          `abi:"startingHeight"`
	EndingHeight     uint64          `abi:"endingHeight"`
	Pagination       PageRequestJson `abi:"pagination"`
}

// Validate WithdrawDelegatorRewardArgs the args
func (args *ValidatorSlashesArgs) Validate() error {
	if args.ValidatorAddress == (common.Address{}) {
		return fmt.Errorf("invalid validator address: %s", args.ValidatorAddress)
	}
	if args.StartingHeight < 0 || args.EndingHeight < 0 {
		return fmt.Errorf("startingHeight %d is less than 0 or endingHeight %d is less than 0", args.StartingHeight, args.EndingHeight)
	}
	if args.StartingHeight > args.EndingHeight {
		return fmt.Errorf("startingHeight %d is greater than endingHeight %d", args.StartingHeight, args.EndingHeight)
	}
	return nil
}

// GetValidator returns the validator address, caller must ensure the validator address is valid
func (args *ValidatorSlashesArgs) GetValidator() sdk.ValAddress {
	valAddr := sdk.ValAddress(args.ValidatorAddress.Bytes())
	return valAddr
}
