// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mechain/permission/params.proto

package types

import (
	fmt "fmt"
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
	// maximum_statements_num defines the maximum number of statements allowed in a policy
	MaximumStatementsNum uint64 `protobuf:"varint,1,opt,name=maximum_statements_num,json=maximumStatementsNum,proto3" json:"maximum_statements_num,omitempty"`
	// maximum_group_num used to set the upper limit on the number of groups to which a resource can grant access permissions.
	// By placing a cap on the number of group permissions, permission control policies can be made more robust and better
	// enforced, thereby reducing the chances of DDos and other security incidents.
	MaximumGroupNum uint64 `protobuf:"varint,2,opt,name=maximum_group_num,json=maximumGroupNum,proto3" json:"maximum_group_num,omitempty"`
	// the maximum iteration number of `RemoveExpiredPolicies` loops in endblocker
	MaximumRemoveExpiredPoliciesIteration uint64 `protobuf:"varint,3,opt,name=maximum_remove_expired_policies_iteration,json=maximumRemoveExpiredPoliciesIteration,proto3" json:"maximum_remove_expired_policies_iteration,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_203a3e460c8207c5, []int{0}
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

func (m *Params) GetMaximumStatementsNum() uint64 {
	if m != nil {
		return m.MaximumStatementsNum
	}
	return 0
}

func (m *Params) GetMaximumGroupNum() uint64 {
	if m != nil {
		return m.MaximumGroupNum
	}
	return 0
}

func (m *Params) GetMaximumRemoveExpiredPoliciesIteration() uint64 {
	if m != nil {
		return m.MaximumRemoveExpiredPoliciesIteration
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "mechain.permission.Params")
}

func init() { proto.RegisterFile("mechain/permission/params.proto", fileDescriptor_203a3e460c8207c5) }

var fileDescriptor_203a3e460c8207c5 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x05, 0xe0, 0x89, 0xca, 0x2c, 0xba, 0x11, 0xcb, 0x20, 0x83, 0x8b, 0x28, 0x82, 0xa0, 0x82,
	0x2d, 0xfe, 0x3c, 0x81, 0x20, 0x83, 0x1b, 0x19, 0xc6, 0x8d, 0xb8, 0x29, 0x99, 0x7a, 0xe9, 0x5c,
	0xf0, 0xf6, 0x86, 0x24, 0x2d, 0xf5, 0x2d, 0x7c, 0x28, 0x17, 0x2e, 0x67, 0xe9, 0x52, 0xda, 0x17,
	0x91, 0x66, 0x52, 0x75, 0x13, 0x42, 0xce, 0x77, 0xb2, 0x38, 0xd1, 0x21, 0x41, 0xbe, 0x52, 0x58,
	0xa6, 0x1a, 0x0c, 0xa1, 0xb5, 0xc8, 0x65, 0xaa, 0x95, 0x51, 0x64, 0x13, 0x6d, 0xd8, 0x71, 0x1c,
	0x07, 0x90, 0xfc, 0x81, 0x83, 0x49, 0xc1, 0x05, 0xfb, 0x38, 0xed, 0x6f, 0x1b, 0x79, 0xfc, 0x21,
	0xa2, 0xf1, 0xdc, 0x57, 0xe3, 0x9b, 0x68, 0x9f, 0x54, 0x83, 0x54, 0x51, 0x66, 0x9d, 0x72, 0x40,
	0x50, 0x3a, 0x9b, 0x95, 0x15, 0x4d, 0xc5, 0x91, 0x38, 0xdd, 0x59, 0x4c, 0x42, 0xfa, 0xf8, 0x1b,
	0x3e, 0x54, 0x14, 0x9f, 0x47, 0x7b, 0x43, 0xab, 0x30, 0x5c, 0x69, 0x5f, 0xd8, 0xf2, 0x85, 0xdd,
	0x10, 0xcc, 0xfa, 0xf7, 0xde, 0x3e, 0x45, 0x67, 0x83, 0x35, 0x40, 0x5c, 0x43, 0x06, 0x8d, 0x46,
	0x03, 0x2f, 0x99, 0xe6, 0x57, 0xcc, 0x11, 0x6c, 0x86, 0x0e, 0x8c, 0x72, 0xc8, 0xe5, 0x74, 0xdb,
	0xff, 0x71, 0x12, 0x0a, 0x0b, 0xef, 0xef, 0x36, 0x7c, 0x1e, 0xf4, 0xfd, 0x80, 0x6f, 0x67, 0x9f,
	0xad, 0x14, 0xeb, 0x56, 0x8a, 0xef, 0x56, 0x8a, 0xf7, 0x4e, 0x8e, 0xd6, 0x9d, 0x1c, 0x7d, 0x75,
	0x72, 0xf4, 0x7c, 0x51, 0xa0, 0x5b, 0x55, 0xcb, 0x24, 0x67, 0x4a, 0xa1, 0x26, 0xb6, 0xe1, 0xac,
	0x2f, 0xaf, 0xd2, 0xe6, 0xff, 0x80, 0xee, 0x4d, 0x83, 0x5d, 0x8e, 0xfd, 0x2c, 0xd7, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x3e, 0x9c, 0x79, 0xa9, 0x63, 0x01, 0x00, 0x00,
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
	if m.MaximumRemoveExpiredPoliciesIteration != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaximumRemoveExpiredPoliciesIteration))
		i--
		dAtA[i] = 0x18
	}
	if m.MaximumGroupNum != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaximumGroupNum))
		i--
		dAtA[i] = 0x10
	}
	if m.MaximumStatementsNum != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaximumStatementsNum))
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
	if m.MaximumStatementsNum != 0 {
		n += 1 + sovParams(uint64(m.MaximumStatementsNum))
	}
	if m.MaximumGroupNum != 0 {
		n += 1 + sovParams(uint64(m.MaximumGroupNum))
	}
	if m.MaximumRemoveExpiredPoliciesIteration != 0 {
		n += 1 + sovParams(uint64(m.MaximumRemoveExpiredPoliciesIteration))
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
				return fmt.Errorf("proto: wrong wireType = %d for field MaximumStatementsNum", wireType)
			}
			m.MaximumStatementsNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaximumStatementsNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaximumGroupNum", wireType)
			}
			m.MaximumGroupNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaximumGroupNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaximumRemoveExpiredPoliciesIteration", wireType)
			}
			m.MaximumRemoveExpiredPoliciesIteration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaximumRemoveExpiredPoliciesIteration |= uint64(b&0x7F) << shift
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
