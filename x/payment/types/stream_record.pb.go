// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mechain/payment/stream_record.proto

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

// StreamAccountStatus defines the status of a stream account
type StreamAccountStatus int32

const (
	// STREAM_ACCOUNT_STATUS_ACTIVE defines the active status of a stream account.
	STREAM_ACCOUNT_STATUS_ACTIVE StreamAccountStatus = 0
	// STREAM_ACCOUNT_STATUS_FROZEN defines the frozen status of a stream account.
	// A frozen stream account cannot be used as payment address for buckets.
	// It can be unfrozen by depositing more azkme to the stream account.
	STREAM_ACCOUNT_STATUS_FROZEN StreamAccountStatus = 1
)

var StreamAccountStatus_name = map[int32]string{
	0: "STREAM_ACCOUNT_STATUS_ACTIVE",
	1: "STREAM_ACCOUNT_STATUS_FROZEN",
}

var StreamAccountStatus_value = map[string]int32{
	"STREAM_ACCOUNT_STATUS_ACTIVE": 0,
	"STREAM_ACCOUNT_STATUS_FROZEN": 1,
}

func (x StreamAccountStatus) String() string {
	return proto.EnumName(StreamAccountStatus_name, int32(x))
}

func (StreamAccountStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f03268c583f4314b, []int{0}
}

// Stream Payment Record of a stream account
type StreamRecord struct {
	// account address
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// latest update timestamp of the stream record
	CrudTimestamp int64 `protobuf:"varint,2,opt,name=crud_timestamp,json=crudTimestamp,proto3" json:"crud_timestamp,omitempty"`
	// The per-second rate that an account's balance is changing.
	// It is the sum of the account's inbound and outbound flow rates.
	NetflowRate github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=netflow_rate,json=netflowRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"netflow_rate"`
	// The balance of the stream account at the latest CRUD timestamp.
	StaticBalance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=static_balance,json=staticBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"static_balance"`
	// reserved balance of the stream account
	// If the netflow rate is negative, the reserved balance is `netflow_rate * reserve_time`
	BufferBalance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=buffer_balance,json=bufferBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"buffer_balance"`
	// the locked balance of the stream account after it puts a new object and before the object is sealed
	LockBalance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=lock_balance,json=lockBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"lock_balance"`
	// the status of the stream account
	Status StreamAccountStatus `protobuf:"varint,7,opt,name=status,proto3,enum=mechain.payment.StreamAccountStatus" json:"status,omitempty"`
	// the unix timestamp when the stream account will be settled
	SettleTimestamp int64 `protobuf:"varint,8,opt,name=settle_timestamp,json=settleTimestamp,proto3" json:"settle_timestamp,omitempty"`
	// the count of its out flows
	OutFlowCount uint64 `protobuf:"varint,9,opt,name=out_flow_count,json=outFlowCount,proto3" json:"out_flow_count,omitempty"`
	// the frozen netflow rate, which is used when resuming stream account
	FrozenNetflowRate github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=frozen_netflow_rate,json=frozenNetflowRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"frozen_netflow_rate"`
}

func (m *StreamRecord) Reset()         { *m = StreamRecord{} }
func (m *StreamRecord) String() string { return proto.CompactTextString(m) }
func (*StreamRecord) ProtoMessage()    {}
func (*StreamRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_f03268c583f4314b, []int{0}
}
func (m *StreamRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRecord.Merge(m, src)
}
func (m *StreamRecord) XXX_Size() int {
	return m.Size()
}
func (m *StreamRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRecord.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRecord proto.InternalMessageInfo

func (m *StreamRecord) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *StreamRecord) GetCrudTimestamp() int64 {
	if m != nil {
		return m.CrudTimestamp
	}
	return 0
}

func (m *StreamRecord) GetStatus() StreamAccountStatus {
	if m != nil {
		return m.Status
	}
	return STREAM_ACCOUNT_STATUS_ACTIVE
}

func (m *StreamRecord) GetSettleTimestamp() int64 {
	if m != nil {
		return m.SettleTimestamp
	}
	return 0
}

func (m *StreamRecord) GetOutFlowCount() uint64 {
	if m != nil {
		return m.OutFlowCount
	}
	return 0
}

func init() {
	proto.RegisterEnum("mechain.payment.StreamAccountStatus", StreamAccountStatus_name, StreamAccountStatus_value)
	proto.RegisterType((*StreamRecord)(nil), "mechain.payment.StreamRecord")
}

func init() {
	proto.RegisterFile("mechain/payment/stream_record.proto", fileDescriptor_f03268c583f4314b)
}

var fileDescriptor_f03268c583f4314b = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6e, 0xd3, 0x40,
	0x18, 0xb5, 0x69, 0x9b, 0xd2, 0x21, 0x4d, 0x83, 0xdb, 0x85, 0x89, 0x90, 0x1b, 0x41, 0x41, 0xa1,
	0x52, 0x1c, 0x11, 0xb6, 0xdd, 0x38, 0x21, 0x95, 0xb2, 0x20, 0x95, 0x6c, 0x97, 0x45, 0x37, 0xa3,
	0xc9, 0x64, 0x92, 0x46, 0xb5, 0x3d, 0xd1, 0xcc, 0xe7, 0x96, 0x72, 0x02, 0x96, 0xdc, 0x81, 0x2b,
	0xf4, 0x10, 0x5d, 0x56, 0x5d, 0x21, 0x16, 0x15, 0x4a, 0x6e, 0xc0, 0x09, 0x90, 0x67, 0x9c, 0xb4,
	0x20, 0xd8, 0x65, 0xe3, 0x9f, 0xcf, 0xcf, 0xef, 0xcd, 0x9b, 0xf7, 0x06, 0xbd, 0x8c, 0x19, 0x3d,
	0x25, 0xe3, 0xa4, 0x31, 0x21, 0x97, 0x31, 0x4b, 0xa0, 0x21, 0x41, 0x30, 0x12, 0x63, 0xc1, 0x28,
	0x17, 0x03, 0x77, 0x22, 0x38, 0x70, 0x6b, 0x2b, 0x07, 0xb9, 0x39, 0xa8, 0xf2, 0x8c, 0x72, 0x19,
	0x73, 0x89, 0xd5, 0xe7, 0x86, 0x7e, 0xd1, 0xd8, 0xca, 0xce, 0x88, 0x8f, 0xb8, 0x9e, 0x67, 0x4f,
	0x7a, 0xfa, 0xe2, 0xd7, 0x1a, 0x2a, 0x06, 0x8a, 0xd9, 0x57, 0xc4, 0x56, 0x13, 0xad, 0x13, 0x4a,
	0x79, 0x9a, 0x80, 0x6d, 0x56, 0xcd, 0xda, 0x46, 0xcb, 0xbe, 0xbd, 0xaa, 0xef, 0xe4, 0x4c, 0xde,
	0x60, 0x20, 0x98, 0x94, 0x01, 0x88, 0x71, 0x32, 0xf2, 0xe7, 0x40, 0xeb, 0x15, 0x2a, 0x51, 0x91,
	0x0e, 0x30, 0x8c, 0x63, 0x26, 0x81, 0xc4, 0x13, 0xfb, 0x51, 0xd5, 0xac, 0xad, 0xf8, 0x9b, 0xd9,
	0x34, 0x9c, 0x0f, 0x2d, 0x8c, 0x8a, 0x09, 0x83, 0x61, 0xc4, 0x2f, 0xb0, 0x20, 0xc0, 0xec, 0x15,
	0xc5, 0x7f, 0x70, 0x7d, 0xb7, 0x6b, 0xfc, 0xb8, 0xdb, 0x7d, 0x3d, 0x1a, 0xc3, 0x69, 0xda, 0x77,
	0x29, 0x8f, 0xf3, 0x85, 0xe7, 0xb7, 0xba, 0x1c, 0x9c, 0x35, 0xe0, 0x72, 0xc2, 0xa4, 0xdb, 0x4d,
	0xe0, 0xf6, 0xaa, 0x8e, 0xf2, 0xd5, 0x74, 0x13, 0xf0, 0x9f, 0xe4, 0x8c, 0x3e, 0x01, 0x66, 0x51,
	0x54, 0x92, 0x40, 0x60, 0x4c, 0x71, 0x9f, 0x44, 0x24, 0xa1, 0xcc, 0x5e, 0x5d, 0x82, 0xc4, 0xa6,
	0xe6, 0x6c, 0x69, 0xca, 0x4c, 0xa4, 0x9f, 0x0e, 0x87, 0x4c, 0x2c, 0x44, 0xd6, 0x96, 0x21, 0xa2,
	0x39, 0xe7, 0x22, 0x18, 0x15, 0x23, 0x4e, 0xcf, 0x16, 0x12, 0x85, 0x65, 0x6c, 0x55, 0xc6, 0x38,
	0x17, 0x38, 0x40, 0x85, 0xcc, 0x56, 0x2a, 0xed, 0xf5, 0xaa, 0x59, 0x2b, 0x35, 0xf7, 0xdc, 0xbf,
	0xaa, 0xe4, 0xea, 0x56, 0x78, 0x3a, 0xe2, 0x40, 0x61, 0xfd, 0xfc, 0x1f, 0xeb, 0x0d, 0x2a, 0x4b,
	0x06, 0x10, 0xb1, 0x07, 0x91, 0x3f, 0x56, 0x91, 0x6f, 0xe9, 0xf9, 0x7d, 0xe8, 0x7b, 0xa8, 0xc4,
	0x53, 0xc0, 0x2a, 0x75, 0x5d, 0xab, 0x8d, 0xaa, 0x59, 0x5b, 0xf5, 0x8b, 0x3c, 0x85, 0xc3, 0x88,
	0x5f, 0xb4, 0x55, 0x83, 0x22, 0xb4, 0x3d, 0x14, 0xfc, 0x33, 0x4b, 0xf0, 0x1f, 0x0d, 0x41, 0x4b,
	0xb0, 0xfd, 0x54, 0x13, 0xf7, 0xee, 0x7b, 0xb2, 0x8f, 0xd1, 0xf6, 0x3f, 0xdc, 0x59, 0x55, 0xf4,
	0x3c, 0x08, 0xfd, 0x8e, 0xf7, 0x01, 0x7b, 0xed, 0xf6, 0xd1, 0x71, 0x2f, 0xc4, 0x41, 0xe8, 0x85,
	0xc7, 0x01, 0xf6, 0xda, 0x61, 0xf7, 0x63, 0xa7, 0x6c, 0xfc, 0x1f, 0x71, 0xe8, 0x1f, 0x9d, 0x74,
	0x7a, 0x65, 0xb3, 0xb2, 0xfa, 0xe5, 0x9b, 0x63, 0xb4, 0xde, 0x5f, 0x4f, 0x1d, 0xf3, 0x66, 0xea,
	0x98, 0x3f, 0xa7, 0x8e, 0xf9, 0x75, 0xe6, 0x18, 0x37, 0x33, 0xc7, 0xf8, 0x3e, 0x73, 0x8c, 0x93,
	0xfd, 0x07, 0x1e, 0xd8, 0x79, 0x66, 0x41, 0x5f, 0xcf, 0xdf, 0x36, 0x1b, 0x9f, 0x16, 0x67, 0x5d,
	0x79, 0xe9, 0x17, 0xd4, 0x11, 0x7d, 0xf7, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x01, 0xee, 0xb3,
	0x0b, 0x04, 0x00, 0x00,
}

func (m *StreamRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FrozenNetflowRate.Size()
		i -= size
		if _, err := m.FrozenNetflowRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStreamRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.OutFlowCount != 0 {
		i = encodeVarintStreamRecord(dAtA, i, uint64(m.OutFlowCount))
		i--
		dAtA[i] = 0x48
	}
	if m.SettleTimestamp != 0 {
		i = encodeVarintStreamRecord(dAtA, i, uint64(m.SettleTimestamp))
		i--
		dAtA[i] = 0x40
	}
	if m.Status != 0 {
		i = encodeVarintStreamRecord(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	{
		size := m.LockBalance.Size()
		i -= size
		if _, err := m.LockBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStreamRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.BufferBalance.Size()
		i -= size
		if _, err := m.BufferBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStreamRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.StaticBalance.Size()
		i -= size
		if _, err := m.StaticBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStreamRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.NetflowRate.Size()
		i -= size
		if _, err := m.NetflowRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStreamRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.CrudTimestamp != 0 {
		i = encodeVarintStreamRecord(dAtA, i, uint64(m.CrudTimestamp))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintStreamRecord(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStreamRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovStreamRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StreamRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovStreamRecord(uint64(l))
	}
	if m.CrudTimestamp != 0 {
		n += 1 + sovStreamRecord(uint64(m.CrudTimestamp))
	}
	l = m.NetflowRate.Size()
	n += 1 + l + sovStreamRecord(uint64(l))
	l = m.StaticBalance.Size()
	n += 1 + l + sovStreamRecord(uint64(l))
	l = m.BufferBalance.Size()
	n += 1 + l + sovStreamRecord(uint64(l))
	l = m.LockBalance.Size()
	n += 1 + l + sovStreamRecord(uint64(l))
	if m.Status != 0 {
		n += 1 + sovStreamRecord(uint64(m.Status))
	}
	if m.SettleTimestamp != 0 {
		n += 1 + sovStreamRecord(uint64(m.SettleTimestamp))
	}
	if m.OutFlowCount != 0 {
		n += 1 + sovStreamRecord(uint64(m.OutFlowCount))
	}
	l = m.FrozenNetflowRate.Size()
	n += 1 + l + sovStreamRecord(uint64(l))
	return n
}

func sovStreamRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStreamRecord(x uint64) (n int) {
	return sovStreamRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreamRecord
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
			return fmt.Errorf("proto: StreamRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamRecord
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
				return ErrInvalidLengthStreamRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreamRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrudTimestamp", wireType)
			}
			m.CrudTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CrudTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetflowRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamRecord
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
				return ErrInvalidLengthStreamRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreamRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NetflowRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StaticBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamRecord
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
				return ErrInvalidLengthStreamRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreamRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StaticBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamRecord
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
				return ErrInvalidLengthStreamRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreamRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BufferBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamRecord
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
				return ErrInvalidLengthStreamRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreamRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LockBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= StreamAccountStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettleTimestamp", wireType)
			}
			m.SettleTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SettleTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutFlowCount", wireType)
			}
			m.OutFlowCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OutFlowCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrozenNetflowRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamRecord
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
				return ErrInvalidLengthStreamRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreamRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FrozenNetflowRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStreamRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreamRecord
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
func skipStreamRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStreamRecord
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
					return 0, ErrIntOverflowStreamRecord
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
					return 0, ErrIntOverflowStreamRecord
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
				return 0, ErrInvalidLengthStreamRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStreamRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStreamRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStreamRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStreamRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStreamRecord = fmt.Errorf("proto: unexpected end of group")
)
