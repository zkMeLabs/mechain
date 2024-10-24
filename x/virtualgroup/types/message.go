package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/gogoproto/proto"

	gnfderrors "github.com/evmos/evmos/v12/types/errors"
)

const (
	TypeMsgCreateGlobalVirtualGroup = "create_global_virtual_group"
	TypeMsgDeleteGlobalVirtualGroup = "delete_global_virtual_group"
	TypeMsgDeposit                  = "deposit"
	TypeMsgWithdraw                 = "withdraw"
	TypeMsgSwapOut                  = "swap_out"
	TypeMsgUpdateParams             = "update_params"
	TypeMsgSettle                   = "settle"
	TypeMsgReserveSwapIn            = "reserve_swap_in"
	TypeMsgCancelSwapIn             = "cancel_swap_in"
	TypeMsgCompleteSwapIn           = "complete_swap_in"
)

var (
	_ sdk.Msg = &MsgCreateGlobalVirtualGroup{}
	_ sdk.Msg = &MsgDeleteGlobalVirtualGroup{}
	_ sdk.Msg = &MsgDeposit{}
	_ sdk.Msg = &MsgWithdraw{}
	_ sdk.Msg = &MsgSwapOut{}
	_ sdk.Msg = &MsgUpdateParams{}
	_ sdk.Msg = &MsgSettle{}
	_ sdk.Msg = &MsgReserveSwapIn{}
	_ sdk.Msg = &MsgCancelSwapIn{}
	_ sdk.Msg = &MsgCompleteSwapIn{}
)

func NewMsgCreateGlobalVirtualGroup(primarySpAddress sdk.AccAddress, globalVirtualFamilyID uint32, secondarySpIDs []uint32, deposit sdk.Coin) *MsgCreateGlobalVirtualGroup {
	return &MsgCreateGlobalVirtualGroup{
		StorageProvider: primarySpAddress.String(),
		FamilyId:        globalVirtualFamilyID,
		SecondarySpIds:  secondarySpIDs,
		Deposit:         deposit,
	}
}

func (msg *MsgCreateGlobalVirtualGroup) Route() string {
	return RouterKey
}

func (msg *MsgCreateGlobalVirtualGroup) Type() string {
	return TypeMsgCreateGlobalVirtualGroup
}

// GetSignBytes implements the LegacyMsg interface.
func (msg *MsgCreateGlobalVirtualGroup) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgCreateGlobalVirtualGroup) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgCreateGlobalVirtualGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid storage provider address (%s)", err)
	}

	if !msg.Deposit.IsValid() || !msg.Deposit.Amount.IsPositive() {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "invalid deposit amount")
	}

	return nil
}

func NewMsgDeleteGlobalVirtualGroup(primarySpAddress sdk.AccAddress, globalVirtualGroupID uint32) *MsgDeleteGlobalVirtualGroup {
	return &MsgDeleteGlobalVirtualGroup{
		StorageProvider:      primarySpAddress.String(),
		GlobalVirtualGroupId: globalVirtualGroupID,
	}
}

func (msg *MsgDeleteGlobalVirtualGroup) Route() string {
	return RouterKey
}

func (msg *MsgDeleteGlobalVirtualGroup) Type() string {
	return TypeMsgDeleteGlobalVirtualGroup
}

// GetSignBytes implements the LegacyMsg interface.
func (msg *MsgDeleteGlobalVirtualGroup) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgDeleteGlobalVirtualGroup) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgDeleteGlobalVirtualGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid storage provider address (%s)", err)
	}

	return nil
}

func NewMsgDeposit(fundingAddress sdk.AccAddress, globalVirtualGroupID uint32, deposit sdk.Coin) *MsgDeposit {
	return &MsgDeposit{
		StorageProvider:      fundingAddress.String(),
		GlobalVirtualGroupId: globalVirtualGroupID,
		Deposit:              deposit,
	}
}

func (msg *MsgDeposit) Route() string {
	return RouterKey
}

func (msg *MsgDeposit) Type() string {
	return TypeMsgDeposit
}

// GetSignBytes implements the LegacyMsg interface.
func (msg *MsgDeposit) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgDeposit) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid storage provider address (%s)", err)
	}

	if !msg.Deposit.IsValid() || !msg.Deposit.Amount.IsPositive() {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "invalid deposit amount")
	}

	return nil
}

func NewMsgWithdraw(fundingAddress sdk.AccAddress, globalVirtualGroupID uint32, withdraw sdk.Coin) *MsgWithdraw {
	return &MsgWithdraw{
		StorageProvider:      fundingAddress.String(),
		GlobalVirtualGroupId: globalVirtualGroupID,
		Withdraw:             withdraw,
	}
}

func (msg *MsgWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgWithdraw) Type() string {
	return TypeMsgWithdraw
}

// GetSignBytes implements the LegacyMsg interface.
func (msg *MsgWithdraw) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgWithdraw) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid storage provider address (%s)", err)
	}

	if !msg.Withdraw.IsValid() || !msg.Withdraw.Amount.IsPositive() {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "invalid or non-positive withdraw amount")
	}
	return nil
}

func NewMsgSwapOut(operatorAddress sdk.AccAddress, globalVirtualGroupFamilyID uint32, globalVirtualGroupIDs []uint32, successorSPID uint32) *MsgSwapOut {
	return &MsgSwapOut{
		StorageProvider:            operatorAddress.String(),
		GlobalVirtualGroupFamilyId: globalVirtualGroupFamilyID,
		GlobalVirtualGroupIds:      globalVirtualGroupIDs,
		SuccessorSpId:              successorSPID,
	}
}

func (msg *MsgSwapOut) Route() string {
	return RouterKey
}

func (msg *MsgSwapOut) Type() string {
	return TypeMsgSwapOut
}

// GetSignBytes implements the LegacyMsg interface.
func (msg *MsgSwapOut) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgSwapOut) GetApprovalBytes() []byte {
	fakeMsg := proto.Clone(msg).(*MsgSwapOut)
	fakeMsg.SuccessorSpApproval.Sig = nil
	return fakeMsg.GetSignBytes()
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgSwapOut) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgSwapOut) ValidateBasic() error {
	_, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid storage provider address (%s)", err)
	}

	if msg.GlobalVirtualGroupFamilyId == NoSpecifiedFamilyID {
		if len(msg.GlobalVirtualGroupIds) == 0 {
			return gnfderrors.ErrInvalidMessage.Wrap("The gvgs are not allowed to be empty when familyID is not specified.")
		}
	} else {
		if len(msg.GlobalVirtualGroupIds) > 0 {
			return gnfderrors.ErrInvalidMessage.Wrap("The gvgs are not allowed to be non-empty when familyID is specified.")
		}
	}

	if msg.SuccessorSpId == 0 {
		return gnfderrors.ErrInvalidMessage.Wrap("The successor sp id is not specified.")
	}

	if msg.SuccessorSpApproval == nil {
		return gnfderrors.ErrInvalidMessage.Wrap("The successor sp approval is not specified.")
	}

	return nil
}

// GetSignBytes implements the LegacyMsg interface.
func (msg *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromHexUnsafe(msg.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromHexUnsafe(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	if err := msg.Params.Validate(); err != nil {
		return err
	}

	return nil
}

func NewMsgSettle(submitter sdk.AccAddress, globalVirtualGroupFamilyID uint32, globalVirtualGroupIDs []uint32) *MsgSettle {
	return &MsgSettle{
		StorageProvider:            submitter.String(),
		GlobalVirtualGroupFamilyId: globalVirtualGroupFamilyID,
		GlobalVirtualGroupIds:      globalVirtualGroupIDs,
	}
}

func (msg *MsgSettle) Route() string {
	return RouterKey
}

func (msg *MsgSettle) Type() string {
	return TypeMsgSettle
}

// GetSignBytes implements the LegacyMsg interface.
func (msg *MsgSettle) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgSettle) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgSettle) ValidateBasic() error {
	_, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid submitter address (%s)", err)
	}

	if msg.GlobalVirtualGroupFamilyId == NoSpecifiedFamilyID {
		if len(msg.GlobalVirtualGroupIds) == 0 || len(msg.GlobalVirtualGroupIds) > 10 {
			return ErrInvalidGVGCount
		}
	}

	return nil
}

func NewMsgReserveSwapIn(storageProvider sdk.AccAddress, targetSPId, globalVirtualGroupFamilyID, globalVirtualGroupID uint32) *MsgReserveSwapIn {
	return &MsgReserveSwapIn{
		StorageProvider:            storageProvider.String(),
		TargetSpId:                 targetSPId,
		GlobalVirtualGroupFamilyId: globalVirtualGroupFamilyID,
		GlobalVirtualGroupId:       globalVirtualGroupID,
	}
}

func (msg *MsgReserveSwapIn) Route() string {
	return RouterKey
}

func (msg *MsgReserveSwapIn) Type() string {
	return TypeMsgReserveSwapIn
}

func (msg *MsgReserveSwapIn) ValidateBasic() error {
	_, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid creator address (%s)", err)
	}
	if msg.GlobalVirtualGroupFamilyId == NoSpecifiedFamilyID {
		if msg.GlobalVirtualGroupId == NoSpecifiedGVGId {
			return gnfderrors.ErrInvalidMessage.Wrap("The gvg id need to be specified when familyID is not specified.")
		}
	} else {
		if msg.GlobalVirtualGroupId != NoSpecifiedGVGId {
			return gnfderrors.ErrInvalidMessage.Wrap("The gvg id need to be empty(0) when familyID is specified.")
		}
	}
	if msg.TargetSpId == 0 {
		return gnfderrors.ErrInvalidMessage.Wrap("The target sp id is not specified.")
	}
	return nil
}

func (msg *MsgReserveSwapIn) GetSigners() []sdk.AccAddress {
	operator, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{operator}
}

func (msg *MsgReserveSwapIn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func NewMsgCancelSwapIn(storageProvider sdk.AccAddress, globalVirtualGroupFamilyID, globalVirtualGroupID uint32) *MsgCancelSwapIn {
	return &MsgCancelSwapIn{
		StorageProvider:            storageProvider.String(),
		GlobalVirtualGroupFamilyId: globalVirtualGroupFamilyID,
		GlobalVirtualGroupId:       globalVirtualGroupID,
	}
}

func (msg *MsgCancelSwapIn) Route() string {
	return RouterKey
}

func (msg *MsgCancelSwapIn) Type() string {
	return TypeMsgCancelSwapIn
}

func (msg *MsgCancelSwapIn) GetSigners() []sdk.AccAddress {
	operator, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{operator}
}

func (msg *MsgCancelSwapIn) ValidateBasic() error {
	_, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid creator address (%s)", err)
	}
	if msg.GlobalVirtualGroupFamilyId == NoSpecifiedFamilyID {
		if msg.GlobalVirtualGroupId == NoSpecifiedGVGId {
			return gnfderrors.ErrInvalidMessage.Wrap("The gvg id need to be specified when familyID is not specified.")
		}
	} else {
		if msg.GlobalVirtualGroupId != NoSpecifiedGVGId {
			return gnfderrors.ErrInvalidMessage.Wrap("The gvg id need to be empty(0) when familyID is specified.")
		}
	}
	return nil
}

func NewMsgCompleteSwapIn(storageProvider sdk.AccAddress, globalVirtualGroupFamilyID, globalVirtualGroupID uint32) *MsgCompleteSwapIn {
	return &MsgCompleteSwapIn{
		StorageProvider:            storageProvider.String(),
		GlobalVirtualGroupFamilyId: globalVirtualGroupFamilyID,
		GlobalVirtualGroupId:       globalVirtualGroupID,
	}
}

func (msg *MsgCompleteSwapIn) Route() string {
	return RouterKey
}

func (msg *MsgCompleteSwapIn) Type() string {
	return TypeMsgCompleteSwapIn
}

func (msg *MsgCompleteSwapIn) GetSigners() []sdk.AccAddress {
	operator, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{operator}
}

func (msg *MsgCompleteSwapIn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCompleteSwapIn) ValidateBasic() error {
	_, err := sdk.AccAddressFromHexUnsafe(msg.StorageProvider)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid creator address (%s)", err)
	}
	if msg.GlobalVirtualGroupFamilyId == NoSpecifiedFamilyID {
		if msg.GlobalVirtualGroupId == NoSpecifiedGVGId {
			return gnfderrors.ErrInvalidMessage.Wrap("The gvg id need to be specified when familyID is not specified.")
		}
	} else {
		if msg.GlobalVirtualGroupId != NoSpecifiedGVGId {
			return gnfderrors.ErrInvalidMessage.Wrap("The gvg id need to be empty(0) when familyID is specified.")
		}
	}
	return nil
}
