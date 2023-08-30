// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: abstractaccount/v1/events.proto

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

type EventAccountRegistered struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	CodeID       uint64 `protobuf:"varint,2,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	ContractAddr string `protobuf:"bytes,3,opt,name=contract_addr,json=contractAddr,proto3" json:"contract_addr,omitempty"`
}

func (m *EventAccountRegistered) Reset()         { *m = EventAccountRegistered{} }
func (m *EventAccountRegistered) String() string { return proto.CompactTextString(m) }
func (*EventAccountRegistered) ProtoMessage()    {}
func (*EventAccountRegistered) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bd53b0900cad0fc, []int{0}
}
func (m *EventAccountRegistered) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAccountRegistered) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAccountRegistered.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAccountRegistered) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAccountRegistered.Merge(m, src)
}
func (m *EventAccountRegistered) XXX_Size() int {
	return m.Size()
}
func (m *EventAccountRegistered) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAccountRegistered.DiscardUnknown(m)
}

var xxx_messageInfo_EventAccountRegistered proto.InternalMessageInfo

func (m *EventAccountRegistered) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *EventAccountRegistered) GetCodeID() uint64 {
	if m != nil {
		return m.CodeID
	}
	return 0
}

func (m *EventAccountRegistered) GetContractAddr() string {
	if m != nil {
		return m.ContractAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*EventAccountRegistered)(nil), "abstractaccount.v1.EventAccountRegistered")
}

func init() { proto.RegisterFile("abstractaccount/v1/events.proto", fileDescriptor_6bd53b0900cad0fc) }

var fileDescriptor_6bd53b0900cad0fc = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x4c, 0x2a, 0x2e,
	0x29, 0x4a, 0x4c, 0x2e, 0x49, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2f, 0x33, 0xd4, 0x4f,
	0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0x53, 0xa0,
	0x57, 0x66, 0x28, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x96, 0xd6, 0x07, 0xb1, 0x20, 0x2a, 0x95,
	0x6a, 0xb8, 0xc4, 0x5c, 0x41, 0x3a, 0x1d, 0x21, 0x0a, 0x83, 0x52, 0xd3, 0x33, 0x8b, 0x4b, 0x52,
	0x8b, 0x52, 0x53, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x65, 0x2e, 0xf6, 0xe4, 0xfc, 0x94, 0xd4, 0xf8,
	0xcc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x16, 0x27, 0xae, 0x47, 0xf7, 0xe4, 0xd9, 0x9c, 0xf3,
	0x53, 0x52, 0x3d, 0x5d, 0x82, 0xd8, 0x40, 0x52, 0x9e, 0x29, 0x42, 0xca, 0x5c, 0xbc, 0xc9, 0xf9,
	0x79, 0x60, 0x47, 0xc4, 0x27, 0xa6, 0xa4, 0x14, 0x49, 0x30, 0x83, 0x0d, 0xe1, 0x81, 0x09, 0x3a,
	0xa6, 0xa4, 0x14, 0x39, 0x85, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94,
	0x75, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x4e, 0x62, 0x51, 0x51,
	0xa5, 0x41, 0x85, 0x3e, 0xcc, 0x53, 0xba, 0x30, 0x6f, 0x23, 0x84, 0x60, 0x22, 0x25, 0x95, 0x05,
	0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xbf, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x90, 0x31, 0x33,
	0x1f, 0x28, 0x01, 0x00, 0x00,
}

func (m *EventAccountRegistered) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAccountRegistered) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAccountRegistered) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddr) > 0 {
		i -= len(m.ContractAddr)
		copy(dAtA[i:], m.ContractAddr)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.CodeID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
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
func (m *EventAccountRegistered) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.CodeID != 0 {
		n += 1 + sovEvents(uint64(m.CodeID))
	}
	l = len(m.ContractAddr)
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
func (m *EventAccountRegistered) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventAccountRegistered: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAccountRegistered: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddr", wireType)
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
			m.ContractAddr = string(dAtA[iNdEx:postIndex])
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
