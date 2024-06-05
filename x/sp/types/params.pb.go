// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/sp/params.proto

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
	// deposit_denom defines the staking coin denomination.
	DepositDenom string `protobuf:"bytes,1,opt,name=deposit_denom,json=depositDenom,proto3" json:"deposit_denom,omitempty"`
	// min_deposit defines the minimum deposit amount for storage providers.
	MinDeposit github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=min_deposit,json=minDeposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_deposit"`
	// the ratio of the store price of the secondary sp to the primary sp, the default value is 80%
	SecondarySpStorePriceRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=secondary_sp_store_price_ratio,json=secondarySpStorePriceRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"secondary_sp_store_price_ratio"`
	// previous blocks that be traced back to for maintenance_records
	NumOfHistoricalBlocksForMaintenanceRecords int64 `protobuf:"varint,4,opt,name=num_of_historical_blocks_for_maintenance_records,json=numOfHistoricalBlocksForMaintenanceRecords,proto3" json:"num_of_historical_blocks_for_maintenance_records,omitempty" yaml:"num_of_historical_blocks_for_maintenance_records"`
	// the max duration that a SP can be in_maintenance within num_of_historical_blocks_for_maintenance_records
	MaintenanceDurationQuota int64 `protobuf:"varint,5,opt,name=maintenance_duration_quota,json=maintenanceDurationQuota,proto3" json:"maintenance_duration_quota,omitempty" yaml:"maintenance_duration_quota"`
	// the number of blocks to be wait for sp to be in maintenance mode again if already requested
	NumOfLockupBlocksForMaintenance int64 `protobuf:"varint,6,opt,name=num_of_lockup_blocks_for_maintenance,json=numOfLockupBlocksForMaintenance,proto3" json:"num_of_lockup_blocks_for_maintenance,omitempty" yaml:"num_of_lockup_blocks_for_maintenance"`
	// the time interval to update global storage price, if it is not set then the price will be updated at the first block of each natural month
	UpdateGlobalPriceInterval uint64 `protobuf:"varint,7,opt,name=update_global_price_interval,json=updateGlobalPriceInterval,proto3" json:"update_global_price_interval,omitempty" yaml:"update_global_price_interval"`
	// the days counting backwards from end of a month in which a sp cannot update its price
	UpdatePriceDisallowedDays uint32 `protobuf:"varint,8,opt,name=update_price_disallowed_days,json=updatePriceDisallowedDays,proto3" json:"update_price_disallowed_days,omitempty" yaml:"update_price_disallowed_days"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5353d8e6e407d7e, []int{0}
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

func (m *Params) GetDepositDenom() string {
	if m != nil {
		return m.DepositDenom
	}
	return ""
}

func (m *Params) GetNumOfHistoricalBlocksForMaintenanceRecords() int64 {
	if m != nil {
		return m.NumOfHistoricalBlocksForMaintenanceRecords
	}
	return 0
}

func (m *Params) GetMaintenanceDurationQuota() int64 {
	if m != nil {
		return m.MaintenanceDurationQuota
	}
	return 0
}

func (m *Params) GetNumOfLockupBlocksForMaintenance() int64 {
	if m != nil {
		return m.NumOfLockupBlocksForMaintenance
	}
	return 0
}

func (m *Params) GetUpdateGlobalPriceInterval() uint64 {
	if m != nil {
		return m.UpdateGlobalPriceInterval
	}
	return 0
}

func (m *Params) GetUpdatePriceDisallowedDays() uint32 {
	if m != nil {
		return m.UpdatePriceDisallowedDays
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "greenfield.sp.Params")
}

func init() { proto.RegisterFile("greenfield/sp/params.proto", fileDescriptor_a5353d8e6e407d7e) }

var fileDescriptor_a5353d8e6e407d7e = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x63, 0x5a, 0x42, 0x31, 0x64, 0xb1, 0x18, 0xdc, 0x08, 0xd9, 0xc1, 0xa5, 0x10, 0x81,
	0x1a, 0xf3, 0x63, 0x40, 0x2a, 0x48, 0x48, 0x91, 0x05, 0x44, 0x02, 0x51, 0xdc, 0x0d, 0x09, 0x9d,
	0x2e, 0xf6, 0x25, 0x39, 0xd5, 0xbe, 0x3b, 0xee, 0xce, 0x81, 0x2c, 0x88, 0x3f, 0x81, 0x91, 0xb1,
	0xfc, 0x0f, 0xfc, 0x11, 0x1d, 0x2b, 0x26, 0xc4, 0x60, 0xa1, 0x64, 0x61, 0xce, 0xc0, 0x8c, 0xee,
	0xce, 0x4a, 0x83, 0x94, 0x56, 0xea, 0xe2, 0x1f, 0xef, 0xfb, 0x7d, 0xef, 0xf3, 0x9e, 0xfd, 0x6c,
	0xbb, 0x39, 0xe4, 0x08, 0x91, 0x01, 0x46, 0x59, 0x1a, 0x0a, 0x16, 0x32, 0xc8, 0x61, 0x2e, 0x3a,
	0x8c, 0x53, 0x49, 0x9d, 0xc6, 0x89, 0xd6, 0x11, 0xac, 0xb9, 0x99, 0x50, 0x91, 0x53, 0x01, 0xb4,
	0x18, 0x9a, 0x1b, 0xe3, 0x6c, 0x5e, 0x1b, 0xd2, 0x21, 0x35, 0x71, 0x75, 0x65, 0xa2, 0xc1, 0xdf,
	0xba, 0x5d, 0xdf, 0xd3, 0x05, 0x9d, 0x2d, 0xbb, 0x91, 0x22, 0x46, 0x05, 0x96, 0x20, 0x45, 0x84,
	0xe6, 0xae, 0xd5, 0xb2, 0xda, 0x97, 0xe3, 0xab, 0x55, 0x30, 0x52, 0x31, 0xe7, 0x9d, 0x7d, 0x25,
	0xc7, 0x04, 0x54, 0x31, 0xf7, 0x82, 0xb2, 0x74, 0x9f, 0x1c, 0x95, 0x7e, 0xed, 0x57, 0xe9, 0xdf,
	0x1a, 0x62, 0x39, 0x2a, 0xfa, 0x9d, 0x84, 0xe6, 0x15, 0xbb, 0x3a, 0xed, 0x88, 0xf4, 0x20, 0x94,
	0x13, 0x86, 0x44, 0xa7, 0x47, 0xe4, 0x8f, 0xef, 0x3b, 0x76, 0xd5, 0x5a, 0x8f, 0xc8, 0xd8, 0xce,
	0x31, 0x89, 0x4c, 0x3d, 0xe7, 0xb3, 0x65, 0x7b, 0x02, 0x25, 0x94, 0xa4, 0x90, 0x4f, 0x80, 0x60,
	0x40, 0x48, 0xca, 0x11, 0x60, 0x1c, 0x27, 0x08, 0x70, 0x28, 0x31, 0x75, 0xd7, 0xce, 0x8d, 0x8c,
	0x50, 0xb2, 0x84, 0x8c, 0x50, 0x12, 0x37, 0x17, 0x8c, 0x7d, 0xb6, 0xaf, 0x08, 0x7b, 0x0a, 0x10,
	0xab, 0xfa, 0xce, 0x37, 0xcb, 0xbe, 0x47, 0x8a, 0x1c, 0xd0, 0x01, 0x18, 0x61, 0x85, 0xc7, 0x09,
	0xcc, 0x40, 0x3f, 0xa3, 0xc9, 0x81, 0x00, 0x03, 0xca, 0x41, 0x0e, 0x31, 0x91, 0x88, 0x40, 0xa2,
	0x5a, 0x42, 0x09, 0xe5, 0xa9, 0x70, 0xd7, 0x5b, 0x56, 0x7b, 0xad, 0xfb, 0x78, 0x5e, 0xfa, 0x8f,
	0x26, 0x30, 0xcf, 0x76, 0x83, 0xf3, 0x56, 0x08, 0xe2, 0x3b, 0xa4, 0xc8, 0x5f, 0x0f, 0x5e, 0x2c,
	0x12, 0xba, 0xda, 0xff, 0x8c, 0xf2, 0x57, 0x27, 0xee, 0xd8, 0x98, 0x9d, 0xc4, 0x6e, 0x2e, 0xd7,
	0x48, 0x0b, 0xfd, 0x68, 0x08, 0x78, 0x5f, 0x50, 0x09, 0xdd, 0x8b, 0xba, 0x99, 0xed, 0x79, 0xe9,
	0xdf, 0x30, 0xcd, 0x9c, 0xee, 0x0d, 0x62, 0x77, 0x49, 0x8c, 0x2a, 0xed, 0x8d, 0x92, 0x9c, 0x4f,
	0xf6, 0xcd, 0x6a, 0x0a, 0xd5, 0x49, 0xc1, 0x4e, 0x99, 0xc0, 0xad, 0x6b, 0x5c, 0x38, 0x2f, 0xfd,
	0xbb, 0xff, 0xcd, 0x7e, 0x66, 0x56, 0x10, 0xfb, 0x7a, 0xde, 0x97, 0xda, 0xb4, 0x6a, 0x56, 0x67,
	0x64, 0x5f, 0x2f, 0x58, 0x0a, 0x25, 0x02, 0xc3, 0x8c, 0xf6, 0x61, 0x56, 0x6d, 0x81, 0x32, 0xf0,
	0x31, 0xcc, 0xdc, 0x4b, 0x2d, 0xab, 0xbd, 0xde, 0xbd, 0x3d, 0x2f, 0xfd, 0x2d, 0xc3, 0x3d, 0xcb,
	0x1d, 0xc4, 0x9b, 0x46, 0x7e, 0xae, 0x55, 0xfd, 0xbe, 0x7b, 0x95, 0xb6, 0x44, 0x32, 0x49, 0x29,
	0x16, 0x30, 0xcb, 0xe8, 0x07, 0x94, 0x82, 0x14, 0x4e, 0x84, 0xbb, 0xd1, 0xb2, 0xda, 0x8d, 0x15,
	0xa4, 0x95, 0xee, 0x05, 0x49, 0x33, 0xa2, 0x85, 0x18, 0xc1, 0x89, 0xd8, 0xdd, 0xf8, 0x7a, 0xe8,
	0xd7, 0xfe, 0x1c, 0xfa, 0x56, 0xf7, 0xe9, 0xd1, 0xd4, 0xb3, 0x8e, 0xa7, 0x9e, 0xf5, 0x7b, 0xea,
	0x59, 0x5f, 0x66, 0x5e, 0xed, 0x78, 0xe6, 0xd5, 0x7e, 0xce, 0xbc, 0xda, 0xdb, 0xed, 0xa5, 0x95,
	0x46, 0x63, 0xb5, 0xd1, 0xe6, 0x38, 0xbe, 0xff, 0x20, 0xfc, 0xa8, 0x7e, 0x00, 0x7a, 0xab, 0xfb,
	0x75, 0xfd, 0x01, 0x3f, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x22, 0xf2, 0x5e, 0xaa, 0x1e, 0x04,
	0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.DepositDenom != that1.DepositDenom {
		return false
	}
	if !this.MinDeposit.Equal(that1.MinDeposit) {
		return false
	}
	if !this.SecondarySpStorePriceRatio.Equal(that1.SecondarySpStorePriceRatio) {
		return false
	}
	if this.NumOfHistoricalBlocksForMaintenanceRecords != that1.NumOfHistoricalBlocksForMaintenanceRecords {
		return false
	}
	if this.MaintenanceDurationQuota != that1.MaintenanceDurationQuota {
		return false
	}
	if this.NumOfLockupBlocksForMaintenance != that1.NumOfLockupBlocksForMaintenance {
		return false
	}
	if this.UpdateGlobalPriceInterval != that1.UpdateGlobalPriceInterval {
		return false
	}
	if this.UpdatePriceDisallowedDays != that1.UpdatePriceDisallowedDays {
		return false
	}
	return true
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
	if m.UpdatePriceDisallowedDays != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UpdatePriceDisallowedDays))
		i--
		dAtA[i] = 0x40
	}
	if m.UpdateGlobalPriceInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UpdateGlobalPriceInterval))
		i--
		dAtA[i] = 0x38
	}
	if m.NumOfLockupBlocksForMaintenance != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.NumOfLockupBlocksForMaintenance))
		i--
		dAtA[i] = 0x30
	}
	if m.MaintenanceDurationQuota != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaintenanceDurationQuota))
		i--
		dAtA[i] = 0x28
	}
	if m.NumOfHistoricalBlocksForMaintenanceRecords != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.NumOfHistoricalBlocksForMaintenanceRecords))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.SecondarySpStorePriceRatio.Size()
		i -= size
		if _, err := m.SecondarySpStorePriceRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinDeposit.Size()
		i -= size
		if _, err := m.MinDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.DepositDenom) > 0 {
		i -= len(m.DepositDenom)
		copy(dAtA[i:], m.DepositDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DepositDenom)))
		i--
		dAtA[i] = 0xa
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
	l = len(m.DepositDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MinDeposit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SecondarySpStorePriceRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.NumOfHistoricalBlocksForMaintenanceRecords != 0 {
		n += 1 + sovParams(uint64(m.NumOfHistoricalBlocksForMaintenanceRecords))
	}
	if m.MaintenanceDurationQuota != 0 {
		n += 1 + sovParams(uint64(m.MaintenanceDurationQuota))
	}
	if m.NumOfLockupBlocksForMaintenance != 0 {
		n += 1 + sovParams(uint64(m.NumOfLockupBlocksForMaintenance))
	}
	if m.UpdateGlobalPriceInterval != 0 {
		n += 1 + sovParams(uint64(m.UpdateGlobalPriceInterval))
	}
	if m.UpdatePriceDisallowedDays != 0 {
		n += 1 + sovParams(uint64(m.UpdatePriceDisallowedDays))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositDenom", wireType)
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
			m.DepositDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeposit", wireType)
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
			if err := m.MinDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecondarySpStorePriceRatio", wireType)
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
			if err := m.SecondarySpStorePriceRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumOfHistoricalBlocksForMaintenanceRecords", wireType)
			}
			m.NumOfHistoricalBlocksForMaintenanceRecords = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumOfHistoricalBlocksForMaintenanceRecords |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaintenanceDurationQuota", wireType)
			}
			m.MaintenanceDurationQuota = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaintenanceDurationQuota |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumOfLockupBlocksForMaintenance", wireType)
			}
			m.NumOfLockupBlocksForMaintenance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumOfLockupBlocksForMaintenance |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateGlobalPriceInterval", wireType)
			}
			m.UpdateGlobalPriceInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateGlobalPriceInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatePriceDisallowedDays", wireType)
			}
			m.UpdatePriceDisallowedDays = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatePriceDisallowedDays |= uint32(b&0x7F) << shift
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
