// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mechain/payment/genesis.proto

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

// GenesisState defines the payment module's genesis state.
type GenesisState struct {
	Params                  Params                `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	StreamRecordList        []StreamRecord        `protobuf:"bytes,2,rep,name=stream_record_list,json=streamRecordList,proto3" json:"stream_record_list"`
	PaymentAccountCountList []PaymentAccountCount `protobuf:"bytes,3,rep,name=payment_account_count_list,json=paymentAccountCountList,proto3" json:"payment_account_count_list"`
	PaymentAccountList      []PaymentAccount      `protobuf:"bytes,4,rep,name=payment_account_list,json=paymentAccountList,proto3" json:"payment_account_list"`
	AutoSettleRecordList    []AutoSettleRecord    `protobuf:"bytes,5,rep,name=auto_settle_record_list,json=autoSettleRecordList,proto3" json:"auto_settle_record_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d504cf9a68fafcd, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetStreamRecordList() []StreamRecord {
	if m != nil {
		return m.StreamRecordList
	}
	return nil
}

func (m *GenesisState) GetPaymentAccountCountList() []PaymentAccountCount {
	if m != nil {
		return m.PaymentAccountCountList
	}
	return nil
}

func (m *GenesisState) GetPaymentAccountList() []PaymentAccount {
	if m != nil {
		return m.PaymentAccountList
	}
	return nil
}

func (m *GenesisState) GetAutoSettleRecordList() []AutoSettleRecord {
	if m != nil {
		return m.AutoSettleRecordList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "mechain.payment.GenesisState")
}

func init() { proto.RegisterFile("mechain/payment/genesis.proto", fileDescriptor_5d504cf9a68fafcd) }

var fileDescriptor_5d504cf9a68fafcd = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xc7, 0x5b, 0x37, 0x77, 0xc8, 0x04, 0xa5, 0x0c, 0x36, 0x8a, 0xeb, 0xe6, 0x3f, 0x18, 0x0a,
	0x2d, 0x4e, 0x7c, 0x80, 0x4d, 0xc1, 0x8b, 0x07, 0xdd, 0x0e, 0x82, 0x07, 0x4b, 0x56, 0x43, 0x57,
	0x58, 0x9b, 0xd2, 0xfc, 0x3a, 0xdc, 0x5b, 0xf8, 0x3c, 0x3e, 0xc1, 0x8e, 0x3b, 0x7a, 0x12, 0xd9,
	0x5e, 0x44, 0xf6, 0x6b, 0x90, 0x35, 0x1d, 0xec, 0x92, 0x94, 0x7c, 0x3f, 0xfd, 0x7c, 0x93, 0x10,
	0xd2, 0x0c, 0x99, 0x37, 0xa6, 0x41, 0xe4, 0xc4, 0x74, 0x16, 0xb2, 0x08, 0x1c, 0x9f, 0x45, 0x4c,
	0x04, 0xc2, 0x8e, 0x13, 0x0e, 0xdc, 0x38, 0x94, 0xb1, 0x2d, 0x63, 0xb3, 0xe6, 0x73, 0x9f, 0x63,
	0xe6, 0xac, 0xbf, 0x32, 0xcc, 0xec, 0xa8, 0x16, 0x9a, 0x02, 0x77, 0x05, 0x03, 0x98, 0x30, 0x37,
	0x61, 0x1e, 0x4f, 0xde, 0x25, 0x79, 0xac, 0x92, 0x31, 0x4d, 0x68, 0x28, 0xeb, 0xcc, 0x8b, 0x62,
	0x8a, 0xb3, 0x4b, 0x3d, 0x8f, 0xa7, 0x11, 0x48, 0xec, 0x6a, 0x07, 0xe6, 0x6e, 0xc2, 0x67, 0x2a,
	0x2c, 0x20, 0x61, 0x34, 0xcc, 0x6d, 0xeb, 0xf4, 0xab, 0x44, 0x0e, 0x1e, 0xb2, 0x93, 0x0f, 0x81,
	0x02, 0x33, 0x6e, 0x49, 0x25, 0xdb, 0x59, 0x43, 0x6f, 0xeb, 0x9d, 0x6a, 0xb7, 0x6e, 0x2b, 0x37,
	0x61, 0x3f, 0x61, 0xdc, 0x2f, 0xcf, 0x7f, 0x5a, 0xda, 0x40, 0xc2, 0xc6, 0x33, 0x31, 0x72, 0x7a,
	0x77, 0x12, 0x08, 0x68, 0xec, 0xb5, 0x4b, 0x9d, 0x6a, 0xb7, 0x59, 0x50, 0x0c, 0x11, 0x1d, 0x20,
	0x29, 0x45, 0x47, 0x62, 0x63, 0xed, 0x31, 0x10, 0x60, 0xf8, 0xc4, 0xdc, 0x7a, 0xbc, 0x4c, 0x5d,
	0x42, 0xf5, 0xf9, 0x96, 0xdd, 0xe1, 0xdc, 0xcb, 0xfe, 0xb8, 0x5b, 0x0f, 0xb2, 0xa1, 0x1e, 0x17,
	0x23, 0x2c, 0x7a, 0x21, 0x35, 0xb5, 0x08, 0x2b, 0xca, 0x58, 0xd1, 0xda, 0x51, 0x21, 0xed, 0x46,
	0xde, 0x8e, 0xe2, 0x37, 0x52, 0x2f, 0xbe, 0x87, 0xcc, 0xbd, 0x8f, 0xee, 0x93, 0x82, 0xbb, 0x97,
	0x02, 0x1f, 0x22, 0x9e, 0xbb, 0x9d, 0x1a, 0x55, 0xd6, 0xd7, 0xfe, 0xfe, 0xfd, 0x7c, 0x69, 0xe9,
	0x8b, 0xa5, 0xa5, 0xff, 0x2e, 0x2d, 0xfd, 0x73, 0x65, 0x69, 0x8b, 0x95, 0xa5, 0x7d, 0xaf, 0x2c,
	0xed, 0xf5, 0xd2, 0x0f, 0x60, 0x9c, 0x8e, 0x6c, 0x8f, 0x87, 0x0e, 0x9b, 0x86, 0x5c, 0xc8, 0x71,
	0x7a, 0xdd, 0x75, 0x3e, 0xfe, 0x1f, 0x04, 0xcc, 0x62, 0x26, 0x46, 0x15, 0x7c, 0x09, 0x37, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x6a, 0x7b, 0xcd, 0x12, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AutoSettleRecordList) > 0 {
		for iNdEx := len(m.AutoSettleRecordList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AutoSettleRecordList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PaymentAccountList) > 0 {
		for iNdEx := len(m.PaymentAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PaymentAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PaymentAccountCountList) > 0 {
		for iNdEx := len(m.PaymentAccountCountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PaymentAccountCountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.StreamRecordList) > 0 {
		for iNdEx := len(m.StreamRecordList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StreamRecordList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.StreamRecordList) > 0 {
		for _, e := range m.StreamRecordList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PaymentAccountCountList) > 0 {
		for _, e := range m.PaymentAccountCountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PaymentAccountList) > 0 {
		for _, e := range m.PaymentAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AutoSettleRecordList) > 0 {
		for _, e := range m.AutoSettleRecordList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamRecordList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StreamRecordList = append(m.StreamRecordList, StreamRecord{})
			if err := m.StreamRecordList[len(m.StreamRecordList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentAccountCountList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentAccountCountList = append(m.PaymentAccountCountList, PaymentAccountCount{})
			if err := m.PaymentAccountCountList[len(m.PaymentAccountCountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentAccountList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentAccountList = append(m.PaymentAccountList, PaymentAccount{})
			if err := m.PaymentAccountList[len(m.PaymentAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoSettleRecordList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AutoSettleRecordList = append(m.AutoSettleRecordList, AutoSettleRecord{})
			if err := m.AutoSettleRecordList[len(m.AutoSettleRecordList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
