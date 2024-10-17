package authz

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"time"

	"github.com/evmos/evmos/v12/x/evm/types"
	sptypes "github.com/evmos/evmos/v12/x/sp/types"
)

const (
	GrantGas = 60_000

	GrantMethodName = "grant"

	GrantEventName = "Grant"

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
		return nil, fmt.Errorf("the method %s is awaiting implementation", args.AuthzType)
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
		return nil, fmt.Errorf("the method %s is awaiting implementation", args.AuthzType)
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
