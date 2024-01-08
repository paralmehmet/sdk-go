// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/types/v1beta1/tx_response.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
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

// base header ak message type, we can cast the bytes into corresponding message
// response type
type TxResponseGenericMessage struct {
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TxResponseGenericMessage) Reset()         { *m = TxResponseGenericMessage{} }
func (m *TxResponseGenericMessage) String() string { return proto.CompactTextString(m) }
func (*TxResponseGenericMessage) ProtoMessage()    {}
func (*TxResponseGenericMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9122428102320deb, []int{0}
}
func (m *TxResponseGenericMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxResponseGenericMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxResponseGenericMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxResponseGenericMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResponseGenericMessage.Merge(m, src)
}
func (m *TxResponseGenericMessage) XXX_Size() int {
	return m.Size()
}
func (m *TxResponseGenericMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResponseGenericMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TxResponseGenericMessage proto.InternalMessageInfo

func (m *TxResponseGenericMessage) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *TxResponseGenericMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// improvised message to unpack length prefixed messages in tx response data
type TxResponseData struct {
	Messages []*TxResponseGenericMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (m *TxResponseData) Reset()         { *m = TxResponseData{} }
func (m *TxResponseData) String() string { return proto.CompactTextString(m) }
func (*TxResponseData) ProtoMessage()    {}
func (*TxResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9122428102320deb, []int{1}
}
func (m *TxResponseData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxResponseData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResponseData.Merge(m, src)
}
func (m *TxResponseData) XXX_Size() int {
	return m.Size()
}
func (m *TxResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_TxResponseData proto.InternalMessageInfo

func (m *TxResponseData) GetMessages() []*TxResponseGenericMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*TxResponseGenericMessage)(nil), "injective.types.v1beta1.TxResponseGenericMessage")
	proto.RegisterType((*TxResponseData)(nil), "injective.types.v1beta1.TxResponseData")
}

func init() {
	proto.RegisterFile("injective/types/v1beta1/tx_response.proto", fileDescriptor_9122428102320deb)
}

var fileDescriptor_9122428102320deb = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcc, 0xcc, 0xcb, 0x4a,
	0x4d, 0x2e, 0xc9, 0x2c, 0x4b, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0x2f, 0xa9, 0x88, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0x2b, 0xd5, 0x03, 0x2b, 0xd5, 0x83, 0x2a, 0x55,
	0x72, 0xe3, 0x92, 0x08, 0xa9, 0x08, 0x82, 0x2a, 0x76, 0x4f, 0xcd, 0x4b, 0x2d, 0xca, 0x4c, 0xf6,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe3, 0x62, 0xcb, 0x48, 0x4d, 0x4c, 0x49, 0x2d,
	0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b,
	0x12, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c, 0xa5, 0x78, 0x2e, 0x3e, 0x84, 0x39,
	0x2e, 0x89, 0x25, 0x89, 0x42, 0xbe, 0x5c, 0x1c, 0xb9, 0x10, 0x83, 0x8a, 0x25, 0x18, 0x15, 0x98,
	0x35, 0xb8, 0x8d, 0x0c, 0xf5, 0x70, 0xb8, 0x42, 0x0f, 0x97, 0x13, 0x82, 0xe0, 0x46, 0x38, 0x85,
	0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb,
	0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6d, 0x7a, 0x66, 0x49, 0x46,
	0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x27, 0xcc, 0x02, 0x9f, 0xc4, 0xa4, 0x62, 0x7d, 0xb8,
	0x75, 0xba, 0xc9, 0xf9, 0x45, 0xa9, 0xc8, 0xdc, 0x8c, 0xc4, 0xcc, 0x3c, 0x48, 0xa0, 0x25, 0xb1,
	0x81, 0x43, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x97, 0x64, 0x5e, 0x96, 0x4e, 0x01, 0x00,
	0x00,
}

func (m *TxResponseGenericMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxResponseGenericMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxResponseGenericMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTxResponse(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Header) > 0 {
		i -= len(m.Header)
		copy(dAtA[i:], m.Header)
		i = encodeVarintTxResponse(dAtA, i, uint64(len(m.Header)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxResponseData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxResponseData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxResponseData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTxResponse(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTxResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovTxResponse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxResponseGenericMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Header)
	if l > 0 {
		n += 1 + l + sovTxResponse(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTxResponse(uint64(l))
	}
	return n
}

func (m *TxResponseData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovTxResponse(uint64(l))
		}
	}
	return n
}

func sovTxResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTxResponse(x uint64) (n int) {
	return sovTxResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxResponseGenericMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxResponse
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
			return fmt.Errorf("proto: TxResponseGenericMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxResponseGenericMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxResponse
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
				return ErrInvalidLengthTxResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Header = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxResponse
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxResponse
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
func (m *TxResponseData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxResponse
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
			return fmt.Errorf("proto: TxResponseData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxResponseData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxResponse
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
				return ErrInvalidLengthTxResponse
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &TxResponseGenericMessage{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxResponse
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
func skipTxResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTxResponse
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
					return 0, ErrIntOverflowTxResponse
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
					return 0, ErrIntOverflowTxResponse
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
				return 0, ErrInvalidLengthTxResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTxResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTxResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTxResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTxResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTxResponse = fmt.Errorf("proto: unexpected end of group")
)
