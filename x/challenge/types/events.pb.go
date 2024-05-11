// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/challenge/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EventStartChallenge to indicate a challenge has bee created.
type EventStartChallenge struct {
	// The id of challenge, which is generated by blockchain.
	ChallengeId uint64 `protobuf:"varint,1,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	// The id of object info to be challenged.
	ObjectId Uint `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3,customtype=Uint" json:"object_id"`
	// The segment/piece index of the object info.
	SegmentIndex uint32 `protobuf:"varint,3,opt,name=segment_index,json=segmentIndex,proto3" json:"segment_index,omitempty"`
	// The storage provider to be challenged.
	SpId uint32 `protobuf:"varint,4,opt,name=sp_id,json=spId,proto3" json:"sp_id,omitempty"`
	// The storage provider to be challenged.
	SpOperatorAddress string `protobuf:"bytes,5,opt,name=sp_operator_address,json=spOperatorAddress,proto3" json:"sp_operator_address,omitempty"`
	// The redundancy index, which comes from the index of storage providers.
	RedundancyIndex int32 `protobuf:"varint,6,opt,name=redundancy_index,json=redundancyIndex,proto3" json:"redundancy_index,omitempty"`
	// The challenger who submits the challenge.
	ChallengerAddress string `protobuf:"bytes,7,opt,name=challenger_address,json=challengerAddress,proto3" json:"challenger_address,omitempty"`
	// The challenge will be expired after this height
	ExpiredHeight uint64 `protobuf:"varint,8,opt,name=expired_height,json=expiredHeight,proto3" json:"expired_height,omitempty"`
}

func (m *EventStartChallenge) Reset()         { *m = EventStartChallenge{} }
func (m *EventStartChallenge) String() string { return proto.CompactTextString(m) }
func (*EventStartChallenge) ProtoMessage()    {}
func (*EventStartChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9eaa4bfadaa20f8, []int{0}
}
func (m *EventStartChallenge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventStartChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventStartChallenge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventStartChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStartChallenge.Merge(m, src)
}
func (m *EventStartChallenge) XXX_Size() int {
	return m.Size()
}
func (m *EventStartChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStartChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_EventStartChallenge proto.InternalMessageInfo

func (m *EventStartChallenge) GetChallengeId() uint64 {
	if m != nil {
		return m.ChallengeId
	}
	return 0
}

func (m *EventStartChallenge) GetSegmentIndex() uint32 {
	if m != nil {
		return m.SegmentIndex
	}
	return 0
}

func (m *EventStartChallenge) GetSpId() uint32 {
	if m != nil {
		return m.SpId
	}
	return 0
}

func (m *EventStartChallenge) GetSpOperatorAddress() string {
	if m != nil {
		return m.SpOperatorAddress
	}
	return ""
}

func (m *EventStartChallenge) GetRedundancyIndex() int32 {
	if m != nil {
		return m.RedundancyIndex
	}
	return 0
}

func (m *EventStartChallenge) GetChallengerAddress() string {
	if m != nil {
		return m.ChallengerAddress
	}
	return ""
}

func (m *EventStartChallenge) GetExpiredHeight() uint64 {
	if m != nil {
		return m.ExpiredHeight
	}
	return 0
}

// EventAttestChallenge to indicate a challenge has been attested.
type EventAttestChallenge struct {
	// The id of challenge.
	ChallengeId uint64 `protobuf:"varint,1,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	// The result of challenge.
	Result VoteResult `protobuf:"varint,2,opt,name=result,proto3,enum=greenfield.challenge.VoteResult" json:"result,omitempty"`
	// The slashed storage provider address.
	SpId uint32 `protobuf:"varint,3,opt,name=sp_id,json=spId,proto3" json:"sp_id,omitempty"`
	// The slashed amount from the storage provider.
	SlashAmount string `protobuf:"bytes,4,opt,name=slash_amount,json=slashAmount,proto3" json:"slash_amount,omitempty"`
	// The address of challenger.
	ChallengerAddress string `protobuf:"bytes,5,opt,name=challenger_address,json=challengerAddress,proto3" json:"challenger_address,omitempty"`
	// The reward amount to the challenger.
	ChallengerRewardAmount string `protobuf:"bytes,6,opt,name=challenger_reward_amount,json=challengerRewardAmount,proto3" json:"challenger_reward_amount,omitempty"`
	// The submitter of the challenge attestation.
	SubmitterAddress string `protobuf:"bytes,7,opt,name=submitter_address,json=submitterAddress,proto3" json:"submitter_address,omitempty"`
	// The reward amount to the submitter.
	SubmitterRewardAmount string `protobuf:"bytes,8,opt,name=submitter_reward_amount,json=submitterRewardAmount,proto3" json:"submitter_reward_amount,omitempty"`
	// The reward amount to all current validators.
	ValidatorRewardAmount string `protobuf:"bytes,10,opt,name=validator_reward_amount,json=validatorRewardAmount,proto3" json:"validator_reward_amount,omitempty"`
}

func (m *EventAttestChallenge) Reset()         { *m = EventAttestChallenge{} }
func (m *EventAttestChallenge) String() string { return proto.CompactTextString(m) }
func (*EventAttestChallenge) ProtoMessage()    {}
func (*EventAttestChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9eaa4bfadaa20f8, []int{1}
}
func (m *EventAttestChallenge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAttestChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAttestChallenge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAttestChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAttestChallenge.Merge(m, src)
}
func (m *EventAttestChallenge) XXX_Size() int {
	return m.Size()
}
func (m *EventAttestChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAttestChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_EventAttestChallenge proto.InternalMessageInfo

func (m *EventAttestChallenge) GetChallengeId() uint64 {
	if m != nil {
		return m.ChallengeId
	}
	return 0
}

func (m *EventAttestChallenge) GetResult() VoteResult {
	if m != nil {
		return m.Result
	}
	return CHALLENGE_FAILED
}

func (m *EventAttestChallenge) GetSpId() uint32 {
	if m != nil {
		return m.SpId
	}
	return 0
}

func (m *EventAttestChallenge) GetSlashAmount() string {
	if m != nil {
		return m.SlashAmount
	}
	return ""
}

func (m *EventAttestChallenge) GetChallengerAddress() string {
	if m != nil {
		return m.ChallengerAddress
	}
	return ""
}

func (m *EventAttestChallenge) GetChallengerRewardAmount() string {
	if m != nil {
		return m.ChallengerRewardAmount
	}
	return ""
}

func (m *EventAttestChallenge) GetSubmitterAddress() string {
	if m != nil {
		return m.SubmitterAddress
	}
	return ""
}

func (m *EventAttestChallenge) GetSubmitterRewardAmount() string {
	if m != nil {
		return m.SubmitterRewardAmount
	}
	return ""
}

func (m *EventAttestChallenge) GetValidatorRewardAmount() string {
	if m != nil {
		return m.ValidatorRewardAmount
	}
	return ""
}

func init() {
	proto.RegisterType((*EventStartChallenge)(nil), "greenfield.challenge.EventStartChallenge")
	proto.RegisterType((*EventAttestChallenge)(nil), "greenfield.challenge.EventAttestChallenge")
}

func init() { proto.RegisterFile("greenfield/challenge/events.proto", fileDescriptor_e9eaa4bfadaa20f8) }

var fileDescriptor_e9eaa4bfadaa20f8 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x8e, 0xff, 0x5c, 0xfe, 0x64, 0x92, 0x94, 0xd6, 0x09, 0xd4, 0x14, 0xc9, 0x75, 0x8a, 0x90,
	0x82, 0x44, 0x13, 0x51, 0xa4, 0xaa, 0xdb, 0x14, 0x15, 0x9a, 0x15, 0x92, 0x2b, 0x58, 0xb0, 0xb1,
	0x9c, 0xcc, 0xc1, 0x31, 0xb2, 0x67, 0xac, 0x99, 0x71, 0x48, 0xdf, 0x02, 0xde, 0xa5, 0x0f, 0xd1,
	0x65, 0xd5, 0x15, 0x62, 0x51, 0xa1, 0x44, 0xbc, 0x07, 0xf2, 0x78, 0x62, 0x27, 0xa8, 0x08, 0x65,
	0x13, 0x65, 0xbe, 0xcb, 0xf9, 0x8e, 0xcf, 0x99, 0x41, 0x1d, 0x8f, 0x01, 0x90, 0x4f, 0x3e, 0x04,
	0xb8, 0x3f, 0x9e, 0xb8, 0x41, 0x00, 0xc4, 0x83, 0x3e, 0x4c, 0x81, 0x08, 0xde, 0x8b, 0x18, 0x15,
	0x54, 0x6f, 0xe7, 0x92, 0x5e, 0x26, 0xd9, 0x7b, 0x3c, 0xa6, 0x3c, 0xa4, 0xdc, 0x91, 0x9a, 0x7e,
	0x7a, 0x48, 0x0d, 0x7b, 0x6d, 0x8f, 0x7a, 0x34, 0xc5, 0x93, 0x7f, 0x0a, 0xb5, 0xee, 0x4d, 0x12,
	0x97, 0x11, 0x28, 0xdf, 0xc1, 0xb7, 0x22, 0x6a, 0x9d, 0x25, 0xc9, 0x17, 0xc2, 0x65, 0xe2, 0xf5,
	0x52, 0xa3, 0x77, 0x50, 0x23, 0x33, 0x38, 0x3e, 0x36, 0x34, 0x4b, 0xeb, 0x96, 0xec, 0x7a, 0x86,
	0x0d, 0xb1, 0x7e, 0x82, 0x6a, 0x74, 0xf4, 0x19, 0xc6, 0x22, 0xe1, 0xff, 0xb3, 0xb4, 0x6e, 0xed,
	0xf4, 0xc9, 0xf5, 0xdd, 0x7e, 0xe1, 0xc7, 0xdd, 0x7e, 0xe9, 0xbd, 0x4f, 0xc4, 0xed, 0xd5, 0x61,
	0x5d, 0xf5, 0x98, 0x1c, 0xed, 0x6a, 0xaa, 0x1e, 0x62, 0xfd, 0x29, 0x6a, 0x72, 0xf0, 0x42, 0x20,
	0xc2, 0xf1, 0x09, 0x86, 0x99, 0x51, 0xb4, 0xb4, 0x6e, 0xd3, 0x6e, 0x28, 0x70, 0x98, 0x60, 0x7a,
	0x0b, 0x95, 0x79, 0x94, 0x94, 0x2e, 0x49, 0xb2, 0xc4, 0xa3, 0x21, 0xd6, 0xcf, 0x51, 0x8b, 0x47,
	0x0e, 0x8d, 0x80, 0xb9, 0x82, 0x32, 0xc7, 0xc5, 0x98, 0x01, 0xe7, 0x46, 0x59, 0xa6, 0x1b, 0xb7,
	0x57, 0x87, 0x6d, 0x95, 0x38, 0x48, 0x99, 0x0b, 0xc1, 0x7c, 0xe2, 0xd9, 0x3b, 0x3c, 0x7a, 0xa7,
	0x3c, 0x8a, 0xd0, 0x9f, 0xa3, 0x6d, 0x06, 0x38, 0x26, 0xd8, 0x25, 0xe3, 0x4b, 0xd5, 0x46, 0xc5,
	0xd2, 0xba, 0x65, 0xfb, 0x41, 0x8e, 0xa7, 0x9d, 0xbc, 0x45, 0x7a, 0xf6, 0xdd, 0x79, 0xe6, 0xff,
	0xff, 0xca, 0xcc, 0x3d, 0xcb, 0xcc, 0x67, 0x68, 0x0b, 0x66, 0x91, 0xcf, 0x00, 0x3b, 0x13, 0xf0,
	0xbd, 0x89, 0x30, 0xaa, 0x72, 0xac, 0x4d, 0x85, 0x9e, 0x4b, 0xf0, 0xe0, 0x57, 0x11, 0xb5, 0xe5,
	0x4e, 0x06, 0x42, 0x00, 0xdf, 0x74, 0x29, 0x15, 0x06, 0x3c, 0x0e, 0x84, 0xdc, 0xc8, 0xd6, 0x91,
	0xd5, 0xbb, 0xef, 0x26, 0xf5, 0x3e, 0x50, 0x01, 0xb6, 0xd4, 0xd9, 0x4a, 0x9f, 0xcf, 0xbb, 0xb8,
	0x32, 0xef, 0x0e, 0x6a, 0xf0, 0xc0, 0xe5, 0x13, 0xc7, 0x0d, 0x69, 0x4c, 0x84, 0xdc, 0x45, 0xcd,
	0xae, 0x4b, 0x6c, 0x20, 0xa1, 0xbf, 0x4c, 0xa7, 0xbc, 0xf9, 0x74, 0x4e, 0x90, 0xb1, 0x52, 0x88,
	0xc1, 0x17, 0x97, 0xe1, 0x65, 0x6e, 0x45, 0xe6, 0x3e, 0xca, 0x79, 0x5b, 0xd2, 0xaa, 0x85, 0x33,
	0xb4, 0xc3, 0xe3, 0x51, 0xe8, 0x0b, 0xb1, 0xc1, 0x7e, 0xb6, 0x33, 0xcb, 0xb2, 0x81, 0x63, 0xb4,
	0x9b, 0x97, 0x59, 0xcf, 0xaf, 0xca, 0xfc, 0x87, 0x19, 0xbd, 0x16, 0x7f, 0x8c, 0x76, 0xa7, 0x6e,
	0xe0, 0x63, 0x79, 0x25, 0xd7, 0x7d, 0x28, 0xf5, 0x65, 0xf4, 0xaa, 0xef, 0xf4, 0xcd, 0xf5, 0xdc,
	0xd4, 0x6e, 0xe6, 0xa6, 0xf6, 0x73, 0x6e, 0x6a, 0x5f, 0x17, 0x66, 0xe1, 0x66, 0x61, 0x16, 0xbe,
	0x2f, 0xcc, 0xc2, 0xc7, 0x17, 0x9e, 0x2f, 0x26, 0xf1, 0xa8, 0x37, 0xa6, 0x61, 0x1f, 0xa6, 0x21,
	0xe5, 0xea, 0x77, 0xfa, 0xf2, 0xa8, 0x3f, 0xfb, 0xf3, 0x25, 0x8f, 0x2a, 0xf2, 0x29, 0xbf, 0xfa,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xb7, 0xe3, 0xb9, 0x58, 0x04, 0x00, 0x00,
}

func (m *EventStartChallenge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventStartChallenge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventStartChallenge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpiredHeight != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ExpiredHeight))
		i--
		dAtA[i] = 0x40
	}
	if len(m.ChallengerAddress) > 0 {
		i -= len(m.ChallengerAddress)
		copy(dAtA[i:], m.ChallengerAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ChallengerAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if m.RedundancyIndex != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.RedundancyIndex))
		i--
		dAtA[i] = 0x30
	}
	if len(m.SpOperatorAddress) > 0 {
		i -= len(m.SpOperatorAddress)
		copy(dAtA[i:], m.SpOperatorAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SpOperatorAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SpId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SpId))
		i--
		dAtA[i] = 0x20
	}
	if m.SegmentIndex != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SegmentIndex))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.ObjectId.Size()
		i -= size
		if _, err := m.ObjectId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.ChallengeId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ChallengeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventAttestChallenge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAttestChallenge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAttestChallenge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorRewardAmount) > 0 {
		i -= len(m.ValidatorRewardAmount)
		copy(dAtA[i:], m.ValidatorRewardAmount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ValidatorRewardAmount)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.SubmitterRewardAmount) > 0 {
		i -= len(m.SubmitterRewardAmount)
		copy(dAtA[i:], m.SubmitterRewardAmount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SubmitterRewardAmount)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.SubmitterAddress) > 0 {
		i -= len(m.SubmitterAddress)
		copy(dAtA[i:], m.SubmitterAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SubmitterAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ChallengerRewardAmount) > 0 {
		i -= len(m.ChallengerRewardAmount)
		copy(dAtA[i:], m.ChallengerRewardAmount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ChallengerRewardAmount)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ChallengerAddress) > 0 {
		i -= len(m.ChallengerAddress)
		copy(dAtA[i:], m.ChallengerAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ChallengerAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SlashAmount) > 0 {
		i -= len(m.SlashAmount)
		copy(dAtA[i:], m.SlashAmount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SlashAmount)))
		i--
		dAtA[i] = 0x22
	}
	if m.SpId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SpId))
		i--
		dAtA[i] = 0x18
	}
	if m.Result != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x10
	}
	if m.ChallengeId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ChallengeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventStartChallenge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChallengeId != 0 {
		n += 1 + sovEvents(uint64(m.ChallengeId))
	}
	l = m.ObjectId.Size()
	n += 1 + l + sovEvents(uint64(l))
	if m.SegmentIndex != 0 {
		n += 1 + sovEvents(uint64(m.SegmentIndex))
	}
	if m.SpId != 0 {
		n += 1 + sovEvents(uint64(m.SpId))
	}
	l = len(m.SpOperatorAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.RedundancyIndex != 0 {
		n += 1 + sovEvents(uint64(m.RedundancyIndex))
	}
	l = len(m.ChallengerAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ExpiredHeight != 0 {
		n += 1 + sovEvents(uint64(m.ExpiredHeight))
	}
	return n
}

func (m *EventAttestChallenge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChallengeId != 0 {
		n += 1 + sovEvents(uint64(m.ChallengeId))
	}
	if m.Result != 0 {
		n += 1 + sovEvents(uint64(m.Result))
	}
	if m.SpId != 0 {
		n += 1 + sovEvents(uint64(m.SpId))
	}
	l = len(m.SlashAmount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ChallengerAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ChallengerRewardAmount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.SubmitterAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.SubmitterRewardAmount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ValidatorRewardAmount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventStartChallenge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventStartChallenge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventStartChallenge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeId", wireType)
			}
			m.ChallengeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SegmentIndex", wireType)
			}
			m.SegmentIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SegmentIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpId", wireType)
			}
			m.SpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpOperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpOperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedundancyIndex", wireType)
			}
			m.RedundancyIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RedundancyIndex |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChallengerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredHeight", wireType)
			}
			m.ExpiredHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiredHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventAttestChallenge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventAttestChallenge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAttestChallenge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeId", wireType)
			}
			m.ChallengeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= VoteResult(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpId", wireType)
			}
			m.SpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlashAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChallengerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengerRewardAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChallengerRewardAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubmitterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitterRewardAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubmitterRewardAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorRewardAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorRewardAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)