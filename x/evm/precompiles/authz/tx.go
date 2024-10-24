package authz

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	bridgetypes "github.com/evmos/evmos/v12/x/bridge/types"
	challengetypes "github.com/evmos/evmos/v12/x/challenge/types"
	erc20types "github.com/evmos/evmos/v12/x/erc20/types"
	"github.com/evmos/evmos/v12/x/evm/types"
	feemarkettypes "github.com/evmos/evmos/v12/x/feemarket/types"
	gensptypes "github.com/evmos/evmos/v12/x/gensp/types"
	paymenttypes "github.com/evmos/evmos/v12/x/payment/types"
	permissiontypes "github.com/evmos/evmos/v12/x/permission/types"
	sptypes "github.com/evmos/evmos/v12/x/sp/types"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
	virtualgrouptypes "github.com/evmos/evmos/v12/x/virtualgroup/types"
)

const (
	GrantGas  = 60_000
	RevokeGas = 60_000
	ExecGas   = 60_000

	GrantMethodName  = "grant"
	RevokeMethodName = "revoke"
	ExecMethodName   = "exec"

	GrantEventName  = "Grant"
	RevokeEventName = "Revoke"
	ExecEventName   = "Exec"

	AuthzTypeSend       = "send"
	AuthzTypeGeneric    = "generic"
	AuthzTypeDelegate   = "delegate"
	AuthzTypeUnbond     = "unbond"
	AuthzTypeRedelegate = "redelegate"
	AuthzTypeSpDeposit  = "spDeposit"
)

// Grant implements the MsgServer.Grant method to create a new grant.
func (c *Contract) Grant(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(GrantMethodName)

	var args GrantArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var limit sdk.Coins
	for _, coin := range args.Limit {
		if coin.Amount.Sign() > 0 {
			limit = limit.Add(sdk.Coin{
				Denom:  coin.Denom,
				Amount: sdk.NewIntFromBigInt(coin.Amount),
			})
		}
	}

	// more details see https://github.com/zkMeLabs/mechain-cosmos-sdk/blob/1ad031a3d3a4b73997d72b8012397633b3cdcae2/x/authz/client/cli/tx.go#L56-L202
	// TODO
	var authorization authz.Authorization
	switch args.AuthzType {
	case AuthzTypeSend:
		// Authorization input example
		// allowed:0x00000004e1E16f249E2b71c2dc66545215FE9d84,0x1111102Dd32160B064F2A512CDEf74bFdB6a9F96
		allowed, err := args.SendParams()
		if err != nil {
			return nil, err
		}
		authorization = banktypes.NewSendAuthorization(limit, allowed)
	case AuthzTypeGeneric:
		authorization = authz.NewGenericAuthorization(args.Authorization)
	case AuthzTypeSpDeposit:
		spAddress := sdk.MustAccAddressFromHex(args.Authorization)
		find, amount := limit.Find(sptypes.DefaultDepositDenom)
		if !find || len(limit.Denoms()) > 1 {
			return nil, fmt.Errorf("limit %s is invalid", limit.String())
		}
		authorization = sptypes.NewDepositAuthorization(spAddress, &amount)
	case AuthzTypeDelegate, AuthzTypeUnbond, AuthzTypeRedelegate:
		if limit.Len() != 1 {
			return nil, fmt.Errorf("limit length must be 1, but limit is %s", limit.String())
		}
		// Authorization input example
		// allowed:0x00000004e1E16f249E2b71c2dc66545215FE9d84,0x1111102Dd32160B064F2A512CDEf74bFdB6a9F96
		// or
		// denied:0x00000004e1E16f249E2b71c2dc66545215FE9d84
		allowed, denied, err := args.StakingParams()
		if err != nil {
			return nil, err
		}

		authzType := stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE
		if args.AuthzType == AuthzTypeDelegate {
			authzType = stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE
		} else if args.AuthzType == AuthzTypeUnbond {
			authzType = stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE
		}

		authorization, err = stakingtypes.NewStakeAuthorization(allowed, denied, authzType, &limit[0])
	default:
		return nil, fmt.Errorf("invalid authorization type %s", args.AuthzType)
	}

	var expiration *time.Time = nil
	if args.Expiration > 0 {
		*expiration = time.Unix(args.Expiration, 0)
	}
	msg, err := authz.NewMsgGrant(contract.Caller().Bytes(), args.Grantee.Bytes(), authorization, expiration)
	if err != nil {
		return nil, err
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	_, err = c.authzKeeper.Grant(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add grant log
	if err := c.AddLog(
		evm,
		MustEvent(GrantEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes()), common.BytesToHash(args.Grantee.Bytes())},
		args.AuthzType,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// Revoke implements the MsgServer.Revoke method.
func (c *Contract) Revoke(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(RevokeMethodName)

	var args RevokeArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &authz.MsgRevoke{
		Granter:    sdk.AccAddress(contract.Caller().Bytes()).String(),
		Grantee:    sdk.AccAddress(args.Grantee.Bytes()).String(),
		MsgTypeUrl: args.MsgTypeUrl,
	}
	if err != nil {
		return nil, err
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	_, err = c.authzKeeper.Revoke(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add revoke log
	if err := c.AddLog(
		evm,
		MustEvent(RevokeEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes()), common.BytesToHash(args.Grantee.Bytes())},
		args.MsgTypeUrl,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// Exec implements the MsgServer.Exec method.
func (c *Contract) Exec(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(ExecMethodName)

	var args ExecArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var messages []json.RawMessage
	err = json.Unmarshal([]byte(args.Msgs), &messages)
	if err != nil {
		return nil, err
	}

	interfaceRegistry := codectypes.NewInterfaceRegistry()

	authtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	distrtypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	govv1beta1.RegisterInterfaces(interfaceRegistry)
	govv1.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	ibctransfertypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	proposal.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	types.RegisterInterfaces(interfaceRegistry)

	bridgetypes.RegisterInterfaces(interfaceRegistry)
	challengetypes.RegisterInterfaces(interfaceRegistry)
	erc20types.RegisterInterfaces(interfaceRegistry)
	feemarkettypes.RegisterInterfaces(interfaceRegistry)
	gensptypes.RegisterInterfaces(interfaceRegistry)
	paymenttypes.RegisterInterfaces(interfaceRegistry)
	permissiontypes.RegisterInterfaces(interfaceRegistry)
	sptypes.RegisterInterfaces(interfaceRegistry)
	storagetypes.RegisterInterfaces(interfaceRegistry)
	virtualgrouptypes.RegisterInterfaces(interfaceRegistry)

	ethosCodec := codec.NewProtoCodec(interfaceRegistry)

	msgs := make([]sdk.Msg, len(messages))
	for i, message := range messages {
		var msg sdk.Msg
		err := ethosCodec.UnmarshalInterfaceJSON(message, &msg)
		if err != nil {
			return nil, err
		}

		msgs[i] = msg
	}

	msg := authz.NewMsgExec(sdk.AccAddress(contract.Caller().Bytes()), msgs)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	_, err = c.authzKeeper.Exec(ctx, &msg)
	if err != nil {
		return nil, err
	}

	// add exec log
	if err := c.AddLog(
		evm,
		MustEvent(ExecEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
