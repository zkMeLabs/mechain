package cli

import (
	"math"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	flag "github.com/spf13/pflag"

	gnfderrors "github.com/evmos/evmos/v12/types/errors"
	permissiontypes "github.com/evmos/evmos/v12/x/permission/types"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
)

const (
	FlagVisibility           = "visibility"
	FlagPaymentAccount       = "payment-account"
	FlagPrimarySP            = "primary-sp"
	FlagExpectChecksums      = "expect-checksums"
	FlagRedundancyType       = "redundancy-type"
	FlagApproveSignature     = "approve-signature"
	FlagApproveTimeoutHeight = "approve-timeout-height"
	FlagChargedReadQuota     = "charged-read-quota"
	FlagBucketID             = "bucket-id"
	FlagBucketName           = "bucket-name"
	FlagDestChainID          = "dest-chain-id"
	FlagObjectID             = "object-id"
	FlagObjectName           = "object-name"
	FlagGroupID              = "group-id"
	FlagGroupName            = "group-name"
	FlagExtra                = "extra"
	FlagTags                 = "tags"
	FlagPrivateKey           = "privatekey"
	FlagGVGFamilyID          = "gvgfamily-id"
)

func GetVisibilityType(str string) (storagetypes.VisibilityType, error) {
	v, ok := storagetypes.VisibilityType_value[str]
	if !ok {
		return storagetypes.VISIBILITY_TYPE_PRIVATE, gnfderrors.ErrInvalidVisibilityType
	}
	visibility := storagetypes.VisibilityType(v)

	return visibility, nil
}

func GetActionType(str string) (permissiontypes.ActionType, error) {
	v, ok := permissiontypes.ActionType_value[str]
	if !ok {
		return permissiontypes.ACTION_UNSPECIFIED, gnfderrors.ErrInvalidActionType
	}
	actionType := permissiontypes.ActionType(v)

	return actionType, nil
}

func GetPrincipalType(str string) (permissiontypes.PrincipalType, error) {
	v, ok := permissiontypes.PrincipalType_value[str]
	if !ok {
		return permissiontypes.PRINCIPAL_TYPE_UNSPECIFIED, gnfderrors.ErrInvalidPrincipalType
	}
	principalType := permissiontypes.PrincipalType(v)

	return principalType, nil
}

func GetPrincipal(str string) (permissiontypes.Principal, error) {
	principalType := permissiontypes.PRINCIPAL_TYPE_GNFD_ACCOUNT
	principalValue := str
	_, err := sdk.AccAddressFromHexUnsafe(str)
	if err != nil {
		principalType = permissiontypes.PRINCIPAL_TYPE_GNFD_GROUP
	}

	return permissiontypes.Principal{
		Type:  principalType,
		Value: principalValue,
	}, nil
}

// GetPrimarySPField returns a from account address, account name and keyring type, given either an address or key name.
func GetPrimarySPField(kr keyring.Keyring, primarySP string) (sdk.AccAddress, string, keyring.KeyType, error) {
	if primarySP == "" {
		return nil, "", 0, nil
	}

	addr, err := sdk.AccAddressFromHexUnsafe(primarySP)

	var k *keyring.Record
	if err == nil {
		k, err = kr.KeyByAddress(addr)
		if err != nil {
			return nil, "", 0, err
		}
	} else {
		k, err = kr.Key(primarySP)
		if err != nil {
			return nil, "", 0, err
		}
	}

	addr, err = k.GetAddress()
	if err != nil {
		return nil, "", 0, err
	}

	return addr, k.Name, k.GetType(), nil
}

// GetPaymentAccountField returns a from account address, account name and keyring type, given either an address or key name.
func GetPaymentAccountField(kr keyring.Keyring, paymentAcc string) (sdk.AccAddress, string, keyring.KeyType, error) {
	if paymentAcc == "" {
		return nil, "", 0, nil
	}

	addr, err := sdk.AccAddressFromHexUnsafe(paymentAcc)

	var k *keyring.Record
	if err == nil {
		k, err = kr.KeyByAddress(addr)
		if err != nil {
			return nil, "", 0, err
		}
	} else {
		k, err = kr.Key(paymentAcc)
		if err != nil {
			return nil, "", 0, err
		}
	}

	addr, err = k.GetAddress()
	if err != nil {
		return nil, "", 0, err
	}

	return addr, k.Name, k.GetType(), nil
}

// FlagSetVisibility Returns the flagSet for set visibility related operations.
func FlagSetVisibility() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagVisibility, "VISIBILITY_TYPE_PRIVATE", "If private, only owner and grantee can access it. Otherwise,"+
		"every one has permission to access it. Select visibility's type (VISIBILITY_TYPE_PRIVATE|VISIBILITY_TYPE_PUBLIC_READ|VISIBILITY_TYPE_INHERIT)")
	return fs
}

func FlagSetApproval() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagApproveSignature, "", "The approval signature of primarySp")
	fs.Uint64(FlagApproveTimeoutHeight, math.MaxUint, "The approval timeout height of primarySp")
	return fs
}

func GetTags(str string) *storagetypes.ResourceTags {
	var tags storagetypes.ResourceTags
	if str == "" || str == "{}" {
		return nil
	}

	tagsStr := str
	if tagsStr[0] == '{' {
		tagsStr = tagsStr[1:]
	}
	if tagsStr[len(tagsStr)-1] == '}' {
		tagsStr = tagsStr[:len(tagsStr)-1]
	}

	for _, tagStr := range strings.Split(tagsStr, ",") {
		kv := strings.Split(tagStr, "=")
		if len(kv) != 2 {
			continue
		}
		tags.Tags = append(tags.Tags, storagetypes.ResourceTags_Tag{Key: kv[0], Value: kv[1]})
	}

	return &tags
}
