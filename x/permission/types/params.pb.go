// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/permission/params.proto

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
	return fileDescriptor_819487f28ea0fa75, []int{0}
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
	proto.RegisterType((*Params)(nil), "greenfield.permission.Params")
}

func init() {
	proto.RegisterFile("greenfield/permission/params.proto", fileDescriptor_819487f28ea0fa75)
}

var fileDescriptor_819487f28ea0fa75 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x05, 0xe0, 0x8e, 0x4a, 0x17, 0xd9, 0x88, 0xa1, 0x4a, 0x71, 0x31, 0x48, 0x41, 0x50, 0xc1,
	0x04, 0xd4, 0x27, 0x10, 0x44, 0x04, 0x91, 0x52, 0x37, 0xe2, 0x26, 0x4c, 0xd2, 0x6b, 0x7a, 0xa1,
	0x33, 0x77, 0x98, 0x1f, 0x89, 0x6f, 0xe1, 0x43, 0xb9, 0x70, 0xd9, 0xa5, 0x4b, 0x49, 0x5e, 0x44,
	0x32, 0x4d, 0x6c, 0x77, 0xc3, 0x9c, 0xef, 0xdc, 0xc5, 0x89, 0x26, 0xa5, 0x01, 0x50, 0x6f, 0x08,
	0xcb, 0x79, 0xaa, 0xc1, 0x48, 0xb4, 0x16, 0x49, 0xa5, 0x5a, 0x18, 0x21, 0x6d, 0xa2, 0x0d, 0x39,
	0x8a, 0x0f, 0x37, 0x26, 0xd9, 0x98, 0xe3, 0x51, 0x49, 0x25, 0x05, 0x91, 0xb6, 0xaf, 0x35, 0x9e,
	0x7c, 0xb1, 0x68, 0x38, 0x0d, 0xed, 0xf8, 0x26, 0x3a, 0x92, 0xa2, 0x42, 0xe9, 0x65, 0x66, 0x9d,
	0x70, 0x20, 0x41, 0x39, 0x9b, 0x29, 0x2f, 0xc7, 0xec, 0x84, 0x9d, 0xed, 0xcd, 0x46, 0x5d, 0xfa,
	0xfc, 0x1f, 0x3e, 0x79, 0x19, 0x5f, 0x44, 0x07, 0x7d, 0xab, 0x34, 0xe4, 0x75, 0x28, 0xec, 0x84,
	0xc2, 0x7e, 0x17, 0xdc, 0xb7, 0xff, 0xad, 0x7d, 0x89, 0xce, 0x7b, 0x6b, 0x40, 0xd2, 0x3b, 0x64,
	0x50, 0x69, 0x34, 0x30, 0xcf, 0x34, 0x2d, 0xb1, 0x40, 0xb0, 0x19, 0x3a, 0x30, 0xc2, 0x21, 0xa9,
	0xf1, 0x6e, 0xb8, 0x71, 0xda, 0x15, 0x66, 0xc1, 0xdf, 0xad, 0xf9, 0xb4, 0xd3, 0x0f, 0x3d, 0xbe,
	0x7d, 0xfc, 0xae, 0x39, 0x5b, 0xd5, 0x9c, 0xfd, 0xd6, 0x9c, 0x7d, 0x36, 0x7c, 0xb0, 0x6a, 0xf8,
	0xe0, 0xa7, 0xe1, 0x83, 0xd7, 0xab, 0x12, 0xdd, 0xc2, 0xe7, 0x49, 0x41, 0x32, 0xcd, 0x55, 0x7e,
	0x59, 0x2c, 0x04, 0xaa, 0x74, 0x6b, 0xc6, 0x6a, 0x7b, 0x48, 0xf7, 0xa1, 0xc1, 0xe6, 0xc3, 0xb0,
	0xcd, 0xf5, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x7d, 0xf8, 0x85, 0x6e, 0x01, 0x00, 0x00,
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
