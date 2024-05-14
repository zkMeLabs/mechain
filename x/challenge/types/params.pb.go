// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/challenge/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	// Challenges which will be emitted in each block, including user submitted or randomly triggered.
	ChallengeCountPerBlock uint64 `protobuf:"varint,1,opt,name=challenge_count_per_block,json=challengeCountPerBlock,proto3" json:"challenge_count_per_block,omitempty" yaml:"challenge_count_per_block"`
	// Challenges will be expired after the period, including user submitted or randomly triggered.
	ChallengeKeepAlivePeriod uint64 `protobuf:"varint,2,opt,name=challenge_keep_alive_period,json=challengeKeepAlivePeriod,proto3" json:"challenge_keep_alive_period,omitempty" yaml:"challenge_keep_alive_period"`
	// The count of blocks to stand for the period in which the same storage and object info cannot be slashed again.
	SlashCoolingOffPeriod uint64 `protobuf:"varint,3,opt,name=slash_cooling_off_period,json=slashCoolingOffPeriod,proto3" json:"slash_cooling_off_period,omitempty" yaml:"slash_cooling_off_period"`
	// The slash coin amount will be calculated from the size of object info, and adjusted by this rate.
	SlashAmountSizeRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=slash_amount_size_rate,json=slashAmountSizeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_amount_size_rate" yaml:"slash_amount_size_rate"`
	// The minimal slash amount.
	SlashAmountMin github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=slash_amount_min,json=slashAmountMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"slash_amount_min"`
	// The maximum slash amount.
	SlashAmountMax github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=slash_amount_max,json=slashAmountMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"slash_amount_max"`
	// The ratio of slash amount to reward all current validators.
	RewardValidatorRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=reward_validator_ratio,json=rewardValidatorRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_validator_ratio" yaml:"reward_validator_ratio"`
	// The ratio of reward amount to reward attestation submitter.
	RewardSubmitterRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=reward_submitter_ratio,json=rewardSubmitterRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_submitter_ratio" yaml:"reward_challenger_ratio"`
	// The reward amount to submitter will be adjusted by the threshold.
	RewardSubmitterThreshold github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=reward_submitter_threshold,json=rewardSubmitterThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reward_submitter_threshold"`
	// Heartbeat interval, based on challenge id, defines the frequency of heartbeat attestation.
	HeartbeatInterval uint64 `protobuf:"varint,10,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty" yaml:"heartbeat_interval"`
	// The time duration for each submitter to submit attestations in turn.
	AttestationInturnInterval uint64 `protobuf:"varint,11,opt,name=attestation_inturn_interval,json=attestationInturnInterval,proto3" json:"attestation_inturn_interval,omitempty" yaml:"attestation_inturn_interval"`
	// The number of kept attested challenge ids, which can be queried by clients.
	AttestationKeptCount uint64 `protobuf:"varint,12,opt,name=attestation_kept_count,json=attestationKeptCount,proto3" json:"attestation_kept_count,omitempty" yaml:"attestation_kept_count"`
	// The max slash amount for a sp in a counting window.
	SpSlashMaxAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,13,opt,name=sp_slash_max_amount,json=spSlashMaxAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sp_slash_max_amount"`
	// The number of blocks to count how much a sp had been slashed.
	SpSlashCountingWindow uint64 `protobuf:"varint,14,opt,name=sp_slash_counting_window,json=spSlashCountingWindow,proto3" json:"sp_slash_counting_window,omitempty" yaml:"sp_slash_counting_window"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2396367ee53edf57, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetChallengeCountPerBlock() uint64 {
	if m != nil {
		return m.ChallengeCountPerBlock
	}
	return 0
}

func (m *Params) GetChallengeKeepAlivePeriod() uint64 {
	if m != nil {
		return m.ChallengeKeepAlivePeriod
	}
	return 0
}

func (m *Params) GetSlashCoolingOffPeriod() uint64 {
	if m != nil {
		return m.SlashCoolingOffPeriod
	}
	return 0
}

func (m *Params) GetHeartbeatInterval() uint64 {
	if m != nil {
		return m.HeartbeatInterval
	}
	return 0
}

func (m *Params) GetAttestationInturnInterval() uint64 {
	if m != nil {
		return m.AttestationInturnInterval
	}
	return 0
}

func (m *Params) GetAttestationKeptCount() uint64 {
	if m != nil {
		return m.AttestationKeptCount
	}
	return 0
}

func (m *Params) GetSpSlashCountingWindow() uint64 {
	if m != nil {
		return m.SpSlashCountingWindow
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "greenfield.challenge.Params")
}

func init() { proto.RegisterFile("greenfield/challenge/params.proto", fileDescriptor_2396367ee53edf57) }

var fileDescriptor_2396367ee53edf57 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0x8e, 0xef, 0xe5, 0x72, 0x2f, 0xc3, 0x2d, 0xa2, 0x26, 0x8d, 0x26, 0x20, 0x62, 0x70, 0x2b,
	0xc4, 0xa2, 0x24, 0x6a, 0xbb, 0x43, 0xdd, 0x10, 0xaa, 0x4a, 0x11, 0x45, 0x44, 0xa6, 0x2a, 0x52,
	0x55, 0xc9, 0x9a, 0xd8, 0xc7, 0xc9, 0x28, 0xb6, 0xc7, 0xb2, 0x27, 0x21, 0xb0, 0x6e, 0xbb, 0xee,
	0xb2, 0xcb, 0x3e, 0x44, 0x1f, 0x82, 0x25, 0xea, 0xaa, 0xea, 0xc2, 0xaa, 0xe0, 0x0d, 0xf2, 0x04,
	0x95, 0x67, 0x1c, 0xc7, 0x24, 0x50, 0x09, 0x95, 0x4d, 0x7e, 0xce, 0xf7, 0x9d, 0xf3, 0x7d, 0xc7,
	0x33, 0xe7, 0x18, 0xad, 0xb7, 0x43, 0x00, 0xdf, 0xa1, 0xe0, 0xda, 0x35, 0xab, 0x43, 0x5c, 0x17,
	0xfc, 0x36, 0xd4, 0x02, 0x12, 0x12, 0x2f, 0xaa, 0x06, 0x21, 0xe3, 0x4c, 0x2d, 0x8e, 0x29, 0xd5,
	0x8c, 0xb2, 0x5c, 0xb6, 0x58, 0xe4, 0xb1, 0xc8, 0x14, 0x9c, 0x9a, 0xfc, 0x23, 0x13, 0x96, 0x8b,
	0x6d, 0xd6, 0x66, 0x32, 0x9e, 0xfc, 0x92, 0x51, 0x7d, 0x38, 0x8f, 0x66, 0x9b, 0xa2, 0xae, 0x6a,
	0xa2, 0x72, 0x56, 0xc8, 0xb4, 0x58, 0xcf, 0xe7, 0x66, 0x00, 0xa1, 0xd9, 0x72, 0x99, 0xd5, 0xc5,
	0xca, 0x9a, 0xb2, 0x39, 0x53, 0x7f, 0x34, 0x8c, 0xb5, 0xb5, 0x13, 0xe2, 0xb9, 0xdb, 0xfa, 0x8d,
	0x54, 0xdd, 0x28, 0x65, 0xd8, 0x6e, 0x02, 0x35, 0x21, 0xac, 0x27, 0x80, 0x0a, 0x68, 0x65, 0x9c,
	0xd5, 0x05, 0x08, 0x4c, 0xe2, 0xd2, 0x3e, 0x24, 0xa9, 0x94, 0xd9, 0xf8, 0x2f, 0x21, 0xb1, 0x31,
	0x8c, 0x35, 0x7d, 0x52, 0x62, 0x8a, 0xac, 0x1b, 0x38, 0x43, 0xf7, 0x00, 0x82, 0x9d, 0x04, 0x6b,
	0x0a, 0x48, 0x7d, 0x87, 0x70, 0xe4, 0x92, 0xa8, 0x63, 0x5a, 0x8c, 0xb9, 0xd4, 0x6f, 0x9b, 0xcc,
	0x71, 0x46, 0x1a, 0x7f, 0x0b, 0x8d, 0x87, 0xc3, 0x58, 0xd3, 0xa4, 0xc6, 0x4d, 0x4c, 0xdd, 0x78,
	0x20, 0xa0, 0x5d, 0x89, 0x1c, 0x38, 0x4e, 0x5a, 0xfd, 0xbd, 0x82, 0x4a, 0x32, 0x89, 0x78, 0xa2,
	0xf1, 0x88, 0x9e, 0x82, 0x19, 0x12, 0x0e, 0x78, 0x66, 0x4d, 0xd9, 0x9c, 0xab, 0x1f, 0x9c, 0xc5,
	0x5a, 0xe1, 0x47, 0xac, 0x6d, 0xb4, 0x29, 0xef, 0xf4, 0x5a, 0x55, 0x8b, 0x79, 0xe9, 0x41, 0xa4,
	0x5f, 0x5b, 0x91, 0xdd, 0xad, 0xf1, 0x93, 0x00, 0xa2, 0xea, 0x0b, 0xb0, 0x86, 0xb1, 0xb6, 0x9a,
	0xb7, 0x32, 0x59, 0x55, 0x37, 0x96, 0x04, 0xb0, 0x23, 0xe2, 0x87, 0xf4, 0x14, 0x0c, 0xc2, 0x41,
	0x75, 0xd0, 0xe2, 0x15, 0xbe, 0x47, 0x7d, 0xfc, 0x8f, 0xd0, 0x7f, 0x7e, 0x0b, 0xfd, 0x86, 0xcf,
	0xbf, 0x7d, 0xdd, 0x42, 0xe9, 0x3d, 0x69, 0xf8, 0xdc, 0x58, 0xc8, 0x89, 0xed, 0x53, 0x7f, 0x5a,
	0x87, 0x0c, 0xf0, 0xec, 0x5d, 0xeb, 0x90, 0x81, 0xfa, 0x41, 0x41, 0xa5, 0x10, 0x8e, 0x49, 0x68,
	0x9b, 0x7d, 0xe2, 0x52, 0x9b, 0x70, 0x16, 0x26, 0xfd, 0x53, 0x86, 0xff, 0xfd, 0xb3, 0xc7, 0x7a,
	0x7d, 0x55, 0xdd, 0x28, 0x4a, 0xe0, 0xcd, 0x28, 0x6e, 0x24, 0x61, 0xf5, 0xe3, 0xd8, 0x47, 0xd4,
	0x6b, 0x79, 0x94, 0x73, 0x18, 0xf9, 0xf8, 0x4f, 0xf8, 0x68, 0xde, 0xda, 0x47, 0xe5, 0x8a, 0x8f,
	0xec, 0xda, 0x4e, 0x1a, 0x39, 0x1c, 0xc9, 0x49, 0x23, 0xa7, 0x68, 0x79, 0xca, 0x07, 0xef, 0x84,
	0x10, 0x75, 0x98, 0x6b, 0xe3, 0xb9, 0x3b, 0x38, 0x02, 0x3c, 0xa1, 0xfb, 0x7a, 0x54, 0x5d, 0x7d,
	0x85, 0xd4, 0x0e, 0x90, 0x90, 0xb7, 0x80, 0x70, 0x93, 0xfa, 0x1c, 0xc2, 0x3e, 0x71, 0x31, 0x12,
	0xb3, 0xb3, 0x3a, 0x8c, 0xb5, 0xb2, 0xec, 0x68, 0x9a, 0xa3, 0x1b, 0xf7, 0xb3, 0x60, 0x23, 0x8d,
	0xa9, 0x0e, 0x5a, 0x21, 0x9c, 0x43, 0xc4, 0x93, 0xbe, 0xfc, 0x84, 0xdb, 0x0b, 0xfd, 0x71, 0xd9,
	0xf9, 0xc9, 0xb1, 0xff, 0x0d, 0x59, 0x37, 0xca, 0x39, 0xb4, 0x21, 0xc0, 0x4c, 0xe7, 0x08, 0x95,
	0xf2, 0xa9, 0x5d, 0x08, 0xb8, 0xdc, 0x4d, 0xf8, 0x7f, 0x21, 0xb1, 0x3e, 0xbe, 0x13, 0xd7, 0xf3,
	0x74, 0xa3, 0x98, 0x03, 0xf6, 0x20, 0xe0, 0x62, 0x7f, 0xa9, 0x5d, 0xb4, 0x14, 0x05, 0xa6, 0x1c,
	0x03, 0x8f, 0x0c, 0xd2, 0x51, 0xc0, 0xf7, 0xee, 0xe0, 0x0c, 0x16, 0xa3, 0xe0, 0x30, 0xa9, 0xbb,
	0x4f, 0x06, 0x72, 0x16, 0xc4, 0xf6, 0x1a, 0x89, 0x09, 0x57, 0xc9, 0x5e, 0x3a, 0xa6, 0xbe, 0xcd,
	0x8e, 0xf1, 0xc2, 0xd4, 0xf6, 0xba, 0x81, 0x99, 0x6c, 0x2f, 0x59, 0x78, 0x37, 0x05, 0x8e, 0x44,
	0x7c, 0x7b, 0xe6, 0xf3, 0x17, 0xad, 0x50, 0x7f, 0x79, 0x76, 0x51, 0x51, 0xce, 0x2f, 0x2a, 0xca,
	0xcf, 0x8b, 0x8a, 0xf2, 0xe9, 0xb2, 0x52, 0x38, 0xbf, 0xac, 0x14, 0xbe, 0x5f, 0x56, 0x0a, 0x6f,
	0x1f, 0xe7, 0xba, 0x80, 0x7e, 0xd2, 0x84, 0xfc, 0xec, 0x3f, 0x79, 0x5a, 0x1b, 0xe4, 0x5e, 0x45,
	0xa2, 0x9f, 0xd6, 0xac, 0x78, 0x87, 0x3c, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x55, 0x8d, 0x0a,
	0xeb, 0xaf, 0x06, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SpSlashCountingWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SpSlashCountingWindow))
		i--
		dAtA[i] = 0x70
	}
	{
		size := m.SpSlashMaxAmount.Size()
		i -= size
		if _, err := m.SpSlashMaxAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	if m.AttestationKeptCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AttestationKeptCount))
		i--
		dAtA[i] = 0x60
	}
	if m.AttestationInturnInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AttestationInturnInterval))
		i--
		dAtA[i] = 0x58
	}
	if m.HeartbeatInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.HeartbeatInterval))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.RewardSubmitterThreshold.Size()
		i -= size
		if _, err := m.RewardSubmitterThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.RewardSubmitterRatio.Size()
		i -= size
		if _, err := m.RewardSubmitterRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.RewardValidatorRatio.Size()
		i -= size
		if _, err := m.RewardValidatorRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.SlashAmountMax.Size()
		i -= size
		if _, err := m.SlashAmountMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.SlashAmountMin.Size()
		i -= size
		if _, err := m.SlashAmountMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.SlashAmountSizeRate.Size()
		i -= size
		if _, err := m.SlashAmountSizeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.SlashCoolingOffPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SlashCoolingOffPeriod))
		i--
		dAtA[i] = 0x18
	}
	if m.ChallengeKeepAlivePeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ChallengeKeepAlivePeriod))
		i--
		dAtA[i] = 0x10
	}
	if m.ChallengeCountPerBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ChallengeCountPerBlock))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChallengeCountPerBlock != 0 {
		n += 1 + sovParams(uint64(m.ChallengeCountPerBlock))
	}
	if m.ChallengeKeepAlivePeriod != 0 {
		n += 1 + sovParams(uint64(m.ChallengeKeepAlivePeriod))
	}
	if m.SlashCoolingOffPeriod != 0 {
		n += 1 + sovParams(uint64(m.SlashCoolingOffPeriod))
	}
	l = m.SlashAmountSizeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashAmountMin.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashAmountMax.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardValidatorRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardSubmitterRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardSubmitterThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.HeartbeatInterval != 0 {
		n += 1 + sovParams(uint64(m.HeartbeatInterval))
	}
	if m.AttestationInturnInterval != 0 {
		n += 1 + sovParams(uint64(m.AttestationInturnInterval))
	}
	if m.AttestationKeptCount != 0 {
		n += 1 + sovParams(uint64(m.AttestationKeptCount))
	}
	l = m.SpSlashMaxAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.SpSlashCountingWindow != 0 {
		n += 1 + sovParams(uint64(m.SpSlashCountingWindow))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeCountPerBlock", wireType)
			}
			m.ChallengeCountPerBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengeCountPerBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeKeepAlivePeriod", wireType)
			}
			m.ChallengeKeepAlivePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengeKeepAlivePeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashCoolingOffPeriod", wireType)
			}
			m.SlashCoolingOffPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SlashCoolingOffPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountSizeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashAmountSizeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashAmountMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashAmountMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardValidatorRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardValidatorRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardSubmitterRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardSubmitterRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardSubmitterThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardSubmitterThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartbeatInterval", wireType)
			}
			m.HeartbeatInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeartbeatInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestationInturnInterval", wireType)
			}
			m.AttestationInturnInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AttestationInturnInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestationKeptCount", wireType)
			}
			m.AttestationKeptCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AttestationKeptCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpSlashMaxAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SpSlashMaxAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpSlashCountingWindow", wireType)
			}
			m.SpSlashCountingWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpSlashCountingWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)