package authz

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	proposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/utils"
	erc20types "github.com/evmos/evmos/v12/x/erc20/types"
	"github.com/evmos/evmos/v12/x/evm/types"
	feemarkettypes "github.com/evmos/evmos/v12/x/feemarket/types"
	sptypes "github.com/evmos/evmos/v12/x/sp/types"
)

const (
	GrantsGas        = 80_000
	GranterGrantsGas = 50_000
	GranteeGrantsGas = 50_000

	GrantsMethodName        = "grants"
	GranterGrantsMethodName = "granterGrants"
	GranteeGrantsMethodName = "granteeGrants"
)

// Grants returns list of `Authorization`, granted to the grantee by the granter.
func (c *Contract) Grants(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(GrantsMethodName)

	var args GrantsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	msg := &authztypes.QueryGrantsRequest{
		Granter:    args.Granter.String(),
		Grantee:    args.Grantee.String(),
		MsgTypeUrl: args.MsgTypeUrl,
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.authzKeeper.Grants(ctx, msg)
	if err != nil {
		return nil, err
	}

	grants := make([]GrantData, 0, len(res.Grants))
	for _, grant := range res.Grants {
		var expiration int64
		if grant.Expiration != nil {
			expiration = grant.Expiration.Unix()
		}
		grants = append(grants, GrantData{
			Authorization: OutputsAuthorization(grant.Authorization.GetCachedValue().(authztypes.Authorization)),
			Expiration:    expiration,
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(grants, pageResponse)
}

// GranterGrants returns list of `GrantAuthorization`, granted by granter.
func (c *Contract) GranterGrants(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(GranterGrantsMethodName)

	var args GranterGrantsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	msg := &authztypes.QueryGranterGrantsRequest{
		Granter: args.Granter.String(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.authzKeeper.GranterGrants(ctx, msg)
	if err != nil {
		return nil, err
	}

	grants := make([]GrantAuthorization, 0, len(res.Grants))
	for _, grant := range res.Grants {
		var expiration int64
		if grant.Expiration != nil {
			expiration = grant.Expiration.Unix()
		}

		grants = append(grants, GrantAuthorization{
			Granter:       utils.AccAddressMustToHexAddress(grant.Granter),
			Grantee:       utils.AccAddressMustToHexAddress(grant.Grantee),
			Authorization: OutputsAuthorization(grant.Authorization.GetCachedValue().(authztypes.Authorization)),
			Expiration:    expiration,
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(grants, pageResponse)
}

// GranteeGrants returns a list of `GrantAuthorization` by grantee.
func (c *Contract) GranteeGrants(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(GranteeGrantsMethodName)

	var args GranteeGrantsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	msg := &authztypes.QueryGranteeGrantsRequest{
		Grantee: args.Grantee.String(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.authzKeeper.GranteeGrants(ctx, msg)
	if err != nil {
		return nil, err
	}

	grants := make([]GrantAuthorization, 0, len(res.Grants))
	for _, grant := range res.Grants {
		var expiration int64
		if grant.Expiration != nil {
			expiration = grant.Expiration.Unix()
		}

		grants = append(grants, GrantAuthorization{
			Granter:       utils.AccAddressMustToHexAddress(grant.Granter),
			Grantee:       utils.AccAddressMustToHexAddress(grant.Grantee),
			Authorization: OutputsAuthorization(grant.Authorization.GetCachedValue().(authztypes.Authorization)),
			Expiration:    expiration,
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(grants, pageResponse)
}

func OutputsAuthorization(authorization authztypes.Authorization) string {
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	authtypes.RegisterInterfaces(interfaceRegistry)
	authztypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	distrtypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	govv1beta1.RegisterInterfaces(interfaceRegistry)
	govv1.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	types.RegisterInterfaces(interfaceRegistry)
	feemarkettypes.RegisterInterfaces(interfaceRegistry)
	erc20types.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	proposaltypes.RegisterInterfaces(interfaceRegistry)
	sptypes.RegisterInterfaces(interfaceRegistry)

	mechainCodec := codec.NewProtoCodec(interfaceRegistry)

	authorizationBytes, err := mechainCodec.MarshalInterfaceJSON(authorization)
	if err == nil {
		return string(authorizationBytes)
	}

	return authorization.String()
}
