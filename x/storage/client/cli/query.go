package cli

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	cmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	evmostypes "github.com/evmos/evmos/v12/types"
	mechaincommon "github.com/evmos/evmos/v12/types/common"
	"github.com/evmos/evmos/v12/types/resource"
	"github.com/evmos/evmos/v12/x/evm/precompiles/storage"
	pertypes "github.com/evmos/evmos/v12/x/permission/types"
	"github.com/evmos/evmos/v12/x/storage/types"
	vgtypes "github.com/evmos/evmos/v12/x/virtualgroup/types"
)

func ToStoragePageReq(in *query.PageRequest) *storage.PageRequest {
	if in == nil {
		return nil
	}
	return &storage.PageRequest{
		Key:        in.Key,
		Offset:     in.Offset,
		Limit:      in.Limit,
		CountTotal: in.CountTotal,
		Reverse:    in.Reverse,
	}
}

func ToPageResp(p *storage.PageResponse) *query.PageResponse {
	if p == nil {
		return nil
	}
	return &query.PageResponse{
		NextKey: p.NextKey,
		Total:   p.Total,
	}
}

func ToResourceTags(tags []storage.Tag) *types.ResourceTags {
	t := make([]types.ResourceTags_Tag, 0)
	if tags == nil {
		return &types.ResourceTags{Tags: t}
	}
	for _, tag := range tags {
		t = append(t, types.ResourceTags_Tag{
			Key:   tag.Key,
			Value: tag.Value,
		})
	}
	return &types.ResourceTags{Tags: t}
}

func ToChecksums(checksumsStr []string) [][]byte {
	var t [][]byte
	for _, checksum := range checksumsStr {
		tmp, _ := hex.DecodeString(checksum)
		t = append(t, tmp)
	}

	return t
}

func ToBucketInfo(p *storage.BucketInfo) *types.BucketInfo {
	if p == nil {
		return nil
	}
	s := &types.BucketInfo{
		Owner:                      p.Owner.String(),
		BucketName:                 p.BucketName,
		Visibility:                 types.VisibilityType(p.Visibility),
		Id:                         cmath.NewUintFromBigInt(p.Id),
		SourceType:                 types.SourceType(p.SourceType),
		CreateAt:                   p.CreateAt,
		PaymentAddress:             p.PaymentAddress.String(),
		GlobalVirtualGroupFamilyId: p.GlobalVirtualGroupFamilyId,
		ChargedReadQuota:           p.ChargedReadQuota,
		BucketStatus:               types.BucketStatus(p.BucketStatus),
		Tags:                       ToResourceTags(p.Tags),
		SpAsDelegatedAgentDisabled: p.SpAsDelegatedAgentDisabled,
	}
	return s
}

func ToBucketExtraInfo(p *storage.BucketExtraInfo) *types.BucketExtraInfo {
	if p == nil {
		return nil
	}
	s := &types.BucketExtraInfo{
		IsRateLimited:   p.IsRateLimited,
		FlowRateLimit:   cmath.NewIntFromBigInt(p.FlowRateLimit),
		CurrentFlowRate: cmath.NewIntFromBigInt(p.CurrentFlowRate),
	}
	return s
}

func ToObjectInfo(p *storage.ObjectInfo) *types.ObjectInfo {
	if p == nil {
		return nil
	}
	s := &types.ObjectInfo{
		Owner:               p.Owner.String(),
		Creator:             p.Creator.String(),
		BucketName:          p.BucketName,
		ObjectName:          p.ObjectName,
		Id:                  cmath.NewUintFromBigInt(p.Id),
		LocalVirtualGroupId: p.LocalVirtualGroupId,
		PayloadSize:         p.PayloadSize,
		Visibility:          types.VisibilityType(p.Visibility),
		ContentType:         p.ContentType,
		CreateAt:            p.CreateAt,
		ObjectStatus:        types.ObjectStatus(p.ObjectStatus),
		RedundancyType:      types.RedundancyType(p.RedundancyType),
		SourceType:          types.SourceType(p.SourceType),
		Checksums:           ToChecksums(p.Checksums),
		Tags:                ToResourceTags(p.Tags),
		IsUpdating:          p.IsUpdating,
		UpdatedAt:           p.UpdatedAt,
		UpdatedBy:           p.UpdatedBy.String(),
		Version:             p.Version,
	}
	return s
}

func ToGroupInfo(p *storage.GroupInfo) *types.GroupInfo {
	if p == nil {
		return nil
	}
	s := &types.GroupInfo{
		Owner:      p.Owner.String(),
		GroupName:  p.GroupName,
		SourceType: types.SourceType(p.SourceType),
		Id:         cmath.NewUintFromBigInt(p.Id),
		Extra:      p.Extra,
		Tags:       ToResourceTags(p.Tags),
	}
	return s
}

func ToGroupMember(p *storage.GroupMember) *pertypes.GroupMember {
	if p == nil {
		return nil
	}
	expirationTime := time.Unix(p.ExpirationTime, 0)
	s := &pertypes.GroupMember{
		Id:             cmath.NewUintFromBigInt(p.Id),
		GroupId:        cmath.NewUintFromBigInt(p.GroupId),
		Member:         p.Member.String(),
		ExpirationTime: &expirationTime,
	}
	return s
}

func ToGlobalVirtualGroup(p *storage.GlobalVirtualGroup) *vgtypes.GlobalVirtualGroup {
	if p == nil {
		return nil
	}
	totalDeposit, _ := sdk.NewIntFromString(p.TotalDeposit)
	s := &vgtypes.GlobalVirtualGroup{
		Id:                    p.Id,
		FamilyId:              p.FamilyId,
		PrimarySpId:           p.PrimarySpId,
		SecondarySpIds:        p.SecondarySpIds,
		StoredSize:            p.StoredSize,
		VirtualPaymentAddress: p.VirtualPaymentAddress.String(),
		TotalDeposit:          totalDeposit,
	}
	return s
}

func ToStatements(p []storage.Statement) []*pertypes.Statement {
	statements := make([]*pertypes.Statement, 0)
	if p == nil {
		return statements
	}

	for _, statement := range p {
		actions := make([]pertypes.ActionType, 0)
		for _, action := range statement.Actions {
			actions = append(actions, pertypes.ActionType(action))
		}
		expirationTime := time.Unix(statement.ExpirationTime, 0)
		statements = append(statements, &pertypes.Statement{
			Effect:         pertypes.Effect(statement.Effect),
			Actions:        actions,
			Resources:      statement.Resources,
			ExpirationTime: &expirationTime,
			LimitSize:      &mechaincommon.UInt64Value{Value: statement.LimitSize},
		})
	}

	return statements
}

func ToPolicy(p *storage.Policy) *pertypes.Policy {
	if p == nil {
		return nil
	}
	expirationTime := time.Unix(p.ExpirationTime, 0)
	s := &pertypes.Policy{
		Id:             cmath.NewUintFromBigInt(p.Id),
		Principal:      &pertypes.Principal{Type: pertypes.PrincipalType(p.Principal.PrincipalType), Value: p.Principal.Value},
		ResourceType:   resource.ResourceType(p.ResourceType),
		ResourceId:     cmath.NewUintFromBigInt(p.ResourceId),
		Statements:     ToStatements(p.Statements),
		ExpirationTime: &expirationTime,
	}
	return s
}

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group storage queries under a subcommand
	storageQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	storageQueryCmd.AddCommand(
		CmdQueryParams(),
		CmdHeadBucket(),
		CmdHeadObject(),
		CmdListBuckets(),
		CmdListObjects(),
		CmdVerifyPermission(),
		CmdHeadGroup(),
		CmdListGroups(),
		CmdHeadGroupMember(),
		CmdQueryAccountPolicy(),
		CmdQueryGroupPolicy(),
	)

	return storageQueryCmd
}

func CmdHeadBucket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "head-bucket [bucket-name]",
		Short: "Query bucket by bucket name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqBucketName := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			contract, err := storage.NewIStorage(common.HexToAddress(evmostypes.StorageAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.HeadBucket(&bind.CallOpts{}, reqBucketName)
			if err != nil {
				return err
			}

			res := &types.QueryHeadBucketResponse{
				BucketInfo: ToBucketInfo(&result.BucketInfo),
				ExtraInfo:  ToBucketExtraInfo(&result.BucketExtraInfo),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdHeadObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "head-object [bucket-name] [object-name]",
		Short: "Query object by bucket-name and object-name",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqBucketName := args[0]
			reqObjectName := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			contract, err := storage.NewIStorage(common.HexToAddress(evmostypes.StorageAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.HeadObject(&bind.CallOpts{}, reqBucketName, reqObjectName)
			if err != nil {
				return err
			}

			res := &types.QueryHeadObjectResponse{
				ObjectInfo:         ToObjectInfo(&result.ObjectInfo),
				GlobalVirtualGroup: ToGlobalVirtualGroup(&result.GlobalVirtualGroup),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListBuckets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-buckets",
		Short: "Query all list buckets",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			contract, err := storage.NewIStorage(common.HexToAddress(evmostypes.StorageAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.ListBuckets(&bind.CallOpts{}, *ToStoragePageReq(pageReq))
			if err != nil {
				return err
			}

			bucketInfos := make([]*types.BucketInfo, 0)
			for _, bucketInfo := range result.BucketInfos {
				bucketInfos = append(bucketInfos, ToBucketInfo(&bucketInfo))
			}
			res := &types.QueryListBucketsResponse{
				BucketInfos: bucketInfos,
				Pagination:  ToPageResp(&result.PageResponse),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListObjects() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-objects [bucket-name]",
		Short: "Query list objects of the bucket",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqBucketName := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			contract, err := storage.NewIStorage(common.HexToAddress(evmostypes.StorageAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.ListObjects(&bind.CallOpts{}, *ToStoragePageReq(pageReq), reqBucketName)
			if err != nil {
				return err
			}

			objectInfos := make([]*types.ObjectInfo, 0)
			for _, objectInfo := range result.ObjectInfos {
				objectInfos = append(objectInfos, ToObjectInfo(&objectInfo))
			}
			res := &types.QueryListObjectsResponse{
				ObjectInfos: objectInfos,
				Pagination:  ToPageResp(&result.PageResponse),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdVerifyPermission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-permission [operator] [bucket-name] [object-name] [action-type]",
		Short: "Query verify if the operator has permission for the bucket/object's action",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqOperator := args[0]
			reqBucketName := args[1]
			reqObjectName := args[2]
			reqActionType := args[3]

			actionType, err := GetActionType(reqActionType)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			contract, err := storage.NewIStorage(common.HexToAddress(evmostypes.StorageAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.VerifyPermission(&bind.CallOpts{From: common.HexToAddress(reqOperator)}, reqBucketName, reqObjectName, int32(actionType))
			if err != nil {
				return err
			}

			res := &types.QueryVerifyPermissionResponse{
				Effect: pertypes.Effect(result),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdHeadGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "head-group [group-owner] [group-name]",
		Short: "Query the group info",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqGroupOwner := args[0]
			reqGroupName := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			contract, err := storage.NewIStorage(common.HexToAddress(evmostypes.StorageAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.HeadGroup(&bind.CallOpts{}, common.HexToAddress(reqGroupOwner), reqGroupName)
			if err != nil {
				return err
			}

			res := &types.QueryHeadGroupResponse{
				GroupInfo: ToGroupInfo(&result),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListGroups() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-groups [group-owner]",
		Short: "Query list groups of owner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqGroupOwner := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			contract, err := storage.NewIStorage(common.HexToAddress(evmostypes.StorageAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.ListGroups(&bind.CallOpts{}, *ToStoragePageReq(pageReq), common.HexToAddress(reqGroupOwner))
			if err != nil {
				return err
			}

			groupInfos := make([]*types.GroupInfo, 0)
			for _, groupInfo := range result.GroupInfos {
				groupInfos = append(groupInfos, ToGroupInfo(&groupInfo))
			}
			res := &types.QueryListGroupsResponse{
				GroupInfos: groupInfos,
				Pagination: ToPageResp(&result.PageResponse),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdHeadGroupMember() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "head-group-member [group-owner] [group-name] [group-member]",
		Short: "Query the group member info",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqGroupOwner := args[0]
			reqGroupName := args[1]
			reqGroupMember := args[2]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			contract, err := storage.NewIStorage(common.HexToAddress(evmostypes.StorageAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.HeadGroupMember(&bind.CallOpts{}, common.HexToAddress(reqGroupMember), common.HexToAddress(reqGroupOwner), reqGroupName)
			if err != nil {
				return err
			}

			res := &types.QueryHeadGroupMemberResponse{
				GroupMember: ToGroupMember(&result),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryAccountPolicy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-policy [grn] [principle-address]",
		Short: "Query the policy for a account that enforced on the resource",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the polciy that a account has on the resource.

Examples:
 $ %s query %s account-policy grn:o::bucketName/objectName 0x....
	`, version.AppName, types.ModuleName),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			grnStr := args[0]
			var grn evmostypes.GRN
			err = grn.ParseFromString(grnStr, false)
			if err != nil {
				return err
			}
			principalAcc, err := sdk.AccAddressFromHexUnsafe(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			contract, err := storage.NewIStorage(common.HexToAddress(evmostypes.StorageAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.QueryPolicyForAccount(&bind.CallOpts{}, grn.String(), principalAcc.String())
			if err != nil {
				return err
			}

			res := &types.QueryPolicyForAccountResponse{
				Policy: ToPolicy(&result),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryGroupPolicy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group-policy [grn] [principle-group-id]",
		Short: "Query the policy for a group that enforced on the resource",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the policy for a group that enforced on the resource

Examples:
 $ %s query %s group-policy grn:o::bucketName/objectName 1
	`, version.AppName, types.ModuleName),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			grnStr := args[0]
			var grn evmostypes.GRN
			err = grn.ParseFromString(grnStr, false)
			if err != nil {
				return err
			}
			groupID, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("failed to convert group id")
			}
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			contract, err := storage.NewIStorage(common.HexToAddress(evmostypes.StorageAddress), clientCtx.EvmClient)
			if err != nil {
				return err
			}
			result, err := contract.QueryPolicyForGroup(&bind.CallOpts{}, grn.String(), groupID.BigInt())
			if err != nil {
				return err
			}

			res := &types.QueryPolicyForGroupResponse{
				Policy: ToPolicy(&result),
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
