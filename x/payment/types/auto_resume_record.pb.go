// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/payment/auto_resume_record.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// AutoResumeRecord is the record keeps the auto resume information.
// The EndBlocker of payment module will scan the list of AutoResumeRecord
// and resume the stream account one by one.
type AutoResumeRecord struct {
	// timestamp is the unix timestamp to order the records
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// the stream account address
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *AutoResumeRecord) Reset()         { *m = AutoResumeRecord{} }
func (m *AutoResumeRecord) String() string { return proto.CompactTextString(m) }
func (*AutoResumeRecord) ProtoMessage()    {}
func (*AutoResumeRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_462258ebd1137873, []int{0}
}
func (m *AutoResumeRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AutoResumeRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AutoResumeRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AutoResumeRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoResumeRecord.Merge(m, src)
}
func (m *AutoResumeRecord) XXX_Size() int {
	return m.Size()
}
func (m *AutoResumeRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoResumeRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AutoResumeRecord proto.InternalMessageInfo

func (m *AutoResumeRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AutoResumeRecord) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*AutoResumeRecord)(nil), "greenfield.payment.AutoResumeRecord")
}

func init() {
	proto.RegisterFile("greenfield/payment/auto_resume_record.proto", fileDescriptor_462258ebd1137873)
}

var fileDescriptor_462258ebd1137873 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0x2f, 0x4a, 0x4d,
	0xcd, 0x4b, 0xcb, 0x4c, 0xcd, 0x49, 0xd1, 0x2f, 0x48, 0xac, 0xcc, 0x4d, 0xcd, 0x2b, 0xd1, 0x4f,
	0x2c, 0x2d, 0xc9, 0x8f, 0x2f, 0x4a, 0x2d, 0x2e, 0xcd, 0x4d, 0x8d, 0x2f, 0x4a, 0x4d, 0xce, 0x2f,
	0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0x28, 0xd6, 0x83, 0x2a, 0x96, 0x92,
	0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x07, 0xab, 0xd0, 0x87, 0x70, 0x20, 0xca, 0x95, 0xe2,
	0xb8, 0x04, 0x1c, 0x4b, 0x4b, 0xf2, 0x83, 0xc0, 0x26, 0x05, 0x81, 0x0d, 0x12, 0x92, 0xe1, 0xe2,
	0x2c, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e,
	0x42, 0x08, 0x08, 0xe9, 0x70, 0xb1, 0x24, 0xa6, 0xa4, 0x14, 0x49, 0x30, 0x29, 0x30, 0x6a, 0x70,
	0x3a, 0x49, 0x5c, 0xda, 0xa2, 0x2b, 0x02, 0x35, 0xd1, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x38,
	0xb8, 0xa4, 0x28, 0x33, 0x2f, 0x3d, 0x08, 0xac, 0xca, 0xc9, 0xe5, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73,
	0xf5, 0x53, 0xcb, 0x72, 0xf3, 0x8b, 0xa1, 0x64, 0x99, 0xa1, 0x91, 0x7e, 0x05, 0xdc, 0x9f, 0x25,
	0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xc7, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x91,
	0x30, 0x28, 0x36, 0x0a, 0x01, 0x00, 0x00,
}

func (m *AutoResumeRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AutoResumeRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AutoResumeRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintAutoResumeRecord(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintAutoResumeRecord(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAutoResumeRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovAutoResumeRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AutoResumeRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovAutoResumeRecord(uint64(m.Timestamp))
	}
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovAutoResumeRecord(uint64(l))
	}
	return n
}

func sovAutoResumeRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAutoResumeRecord(x uint64) (n int) {
	return sovAutoResumeRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AutoResumeRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAutoResumeRecord
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
			return fmt.Errorf("proto: AutoResumeRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AutoResumeRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAutoResumeRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAutoResumeRecord
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
				return ErrInvalidLengthAutoResumeRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAutoResumeRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAutoResumeRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAutoResumeRecord
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
func skipAutoResumeRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAutoResumeRecord
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
					return 0, ErrIntOverflowAutoResumeRecord
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
					return 0, ErrIntOverflowAutoResumeRecord
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
				return 0, ErrInvalidLengthAutoResumeRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAutoResumeRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAutoResumeRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAutoResumeRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAutoResumeRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAutoResumeRecord = fmt.Errorf("proto: unexpected end of group")
)