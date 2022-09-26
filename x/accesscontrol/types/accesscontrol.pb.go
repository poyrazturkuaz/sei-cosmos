// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol/accesscontrol.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type AccessOperation struct {
	AccessType         AccessType   `protobuf:"varint,1,opt,name=access_type,json=accessType,proto3,enum=cosmos.accesscontrol.v1beta1.AccessType" json:"access_type,omitempty"`
	ResourceType       ResourceType `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=cosmos.accesscontrol.v1beta1.ResourceType" json:"resource_type,omitempty"`
	IdentifierTemplate string       `protobuf:"bytes,3,opt,name=identifier_template,json=identifierTemplate,proto3" json:"identifier_template,omitempty"`
}

func (m *AccessOperation) Reset()         { *m = AccessOperation{} }
func (m *AccessOperation) String() string { return proto.CompactTextString(m) }
func (*AccessOperation) ProtoMessage()    {}
func (*AccessOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d636a082612ba091, []int{0}
}
func (m *AccessOperation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessOperation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessOperation.Merge(m, src)
}
func (m *AccessOperation) XXX_Size() int {
	return m.Size()
}
func (m *AccessOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AccessOperation proto.InternalMessageInfo

func (m *AccessOperation) GetAccessType() AccessType {
	if m != nil {
		return m.AccessType
	}
	return AccessType_UNKNOWN
}

func (m *AccessOperation) GetResourceType() ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return ResourceType_ANY
}

func (m *AccessOperation) GetIdentifierTemplate() string {
	if m != nil {
		return m.IdentifierTemplate
	}
	return ""
}

type MessageDependencyMapping struct {
	MessageKey string            `protobuf:"bytes,1,opt,name=message_key,json=messageKey,proto3" json:"message_key,omitempty"`
	AccessOps  []AccessOperation `protobuf:"bytes,2,rep,name=access_ops,json=accessOps,proto3" json:"access_ops"`
}

func (m *MessageDependencyMapping) Reset()         { *m = MessageDependencyMapping{} }
func (m *MessageDependencyMapping) String() string { return proto.CompactTextString(m) }
func (*MessageDependencyMapping) ProtoMessage()    {}
func (*MessageDependencyMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_d636a082612ba091, []int{1}
}
func (m *MessageDependencyMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageDependencyMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageDependencyMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageDependencyMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageDependencyMapping.Merge(m, src)
}
func (m *MessageDependencyMapping) XXX_Size() int {
	return m.Size()
}
func (m *MessageDependencyMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageDependencyMapping.DiscardUnknown(m)
}

var xxx_messageInfo_MessageDependencyMapping proto.InternalMessageInfo

func (m *MessageDependencyMapping) GetMessageKey() string {
	if m != nil {
		return m.MessageKey
	}
	return ""
}

func (m *MessageDependencyMapping) GetAccessOps() []AccessOperation {
	if m != nil {
		return m.AccessOps
	}
	return nil
}

func init() {
	proto.RegisterType((*AccessOperation)(nil), "cosmos.accesscontrol.v1beta1.AccessOperation")
	proto.RegisterType((*MessageDependencyMapping)(nil), "cosmos.accesscontrol.v1beta1.MessageDependencyMapping")
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol/accesscontrol.proto", fileDescriptor_d636a082612ba091)
}

var fileDescriptor_d636a082612ba091 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x6a, 0xf2, 0x40,
	0x10, 0xc7, 0xb3, 0xfa, 0xf1, 0x81, 0xeb, 0xf7, 0xb5, 0xb0, 0xed, 0x21, 0x48, 0x89, 0x22, 0x3d,
	0x84, 0x82, 0x09, 0xda, 0x27, 0xa8, 0xf4, 0x52, 0x5a, 0x11, 0x82, 0xa7, 0x5e, 0xc2, 0xba, 0x4e,
	0xd3, 0xa0, 0xd9, 0x5d, 0x76, 0xd7, 0xd2, 0x3c, 0x45, 0xfb, 0x58, 0x1e, 0x3d, 0xf6, 0x24, 0x45,
	0x5f, 0xa4, 0x98, 0x0d, 0xb6, 0x96, 0x22, 0x3d, 0xed, 0xcc, 0xfc, 0xff, 0xf3, 0x63, 0x67, 0x06,
	0xfb, 0x4c, 0xe8, 0x4c, 0xe8, 0x90, 0x32, 0x06, 0x5a, 0x33, 0xc1, 0x8d, 0x12, 0xb3, 0xfd, 0x2c,
	0x90, 0x4a, 0x18, 0x41, 0xce, 0xac, 0x33, 0xd8, 0xd7, 0x9e, 0xba, 0x63, 0x30, 0xb4, 0xdb, 0x38,
	0x4d, 0x44, 0x22, 0x0a, 0x63, 0xb8, 0x8d, 0x6c, 0x4f, 0xe3, 0xfc, 0x47, 0x3a, 0x13, 0x5c, 0x1b,
	0xca, 0x8d, 0xb6, 0xae, 0xf6, 0x0a, 0xe1, 0xe3, 0xab, 0xc2, 0x31, 0x94, 0xa0, 0xa8, 0x49, 0x05,
	0x27, 0x37, 0xb8, 0x6e, 0x9b, 0x62, 0x93, 0x4b, 0x70, 0x51, 0x0b, 0xf9, 0x47, 0x3d, 0x3f, 0x38,
	0xf4, 0x87, 0xc0, 0x32, 0x46, 0xb9, 0x84, 0x08, 0xd3, 0x5d, 0x4c, 0x86, 0xf8, 0xbf, 0x02, 0x2d,
	0xe6, 0x8a, 0x81, 0x85, 0x55, 0x0a, 0xd8, 0xc5, 0x61, 0x58, 0x54, 0xb6, 0x14, 0xb8, 0x7f, 0xea,
	0x4b, 0x46, 0x42, 0x7c, 0x92, 0x4e, 0x80, 0x9b, 0xf4, 0x21, 0x05, 0x15, 0x1b, 0xc8, 0xe4, 0x8c,
	0x1a, 0x70, 0xab, 0x2d, 0xe4, 0xd7, 0x22, 0xf2, 0x29, 0x8d, 0x4a, 0xa5, 0xfd, 0x82, 0xb0, 0x3b,
	0x00, 0xad, 0x69, 0x02, 0xd7, 0x20, 0x81, 0x4f, 0x80, 0xb3, 0x7c, 0x40, 0xa5, 0x4c, 0x79, 0x42,
	0x9a, 0xb8, 0x9e, 0x59, 0x2d, 0x9e, 0x42, 0x5e, 0x4c, 0x5a, 0x8b, 0x70, 0x59, 0xba, 0x85, 0x9c,
	0x44, 0xb8, 0x9c, 0x26, 0x16, 0x52, 0xbb, 0x95, 0x56, 0xd5, 0xaf, 0xf7, 0x3a, 0xbf, 0xd9, 0xc4,
	0x6e, 0x9b, 0xfd, 0x3f, 0x8b, 0x55, 0xd3, 0x89, 0x6a, 0xb4, 0x2c, 0xeb, 0xfe, 0xdd, 0x62, 0xed,
	0xa1, 0xe5, 0xda, 0x43, 0xef, 0x6b, 0x0f, 0xbd, 0x6e, 0x3c, 0x67, 0xb9, 0xf1, 0x9c, 0xb7, 0x8d,
	0xe7, 0xdc, 0xf7, 0x92, 0xd4, 0x3c, 0xce, 0xc7, 0x01, 0x13, 0x59, 0x58, 0x5e, 0xcf, 0x3e, 0x1d,
	0x3d, 0x99, 0x86, 0xcf, 0xdf, 0x4e, 0xb9, 0xdd, 0xa7, 0x1e, 0xff, 0x2d, 0xee, 0x78, 0xf9, 0x11,
	0x00, 0x00, 0xff, 0xff, 0x6d, 0x7a, 0x65, 0x41, 0x4d, 0x02, 0x00, 0x00,
}

func (m *AccessOperation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessOperation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessOperation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IdentifierTemplate) > 0 {
		i -= len(m.IdentifierTemplate)
		copy(dAtA[i:], m.IdentifierTemplate)
		i = encodeVarintAccesscontrol(dAtA, i, uint64(len(m.IdentifierTemplate)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ResourceType != 0 {
		i = encodeVarintAccesscontrol(dAtA, i, uint64(m.ResourceType))
		i--
		dAtA[i] = 0x10
	}
	if m.AccessType != 0 {
		i = encodeVarintAccesscontrol(dAtA, i, uint64(m.AccessType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MessageDependencyMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageDependencyMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MessageDependencyMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccessOps) > 0 {
		for iNdEx := len(m.AccessOps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccessOps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccesscontrol(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.MessageKey) > 0 {
		i -= len(m.MessageKey)
		copy(dAtA[i:], m.MessageKey)
		i = encodeVarintAccesscontrol(dAtA, i, uint64(len(m.MessageKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccesscontrol(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccesscontrol(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccessOperation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccessType != 0 {
		n += 1 + sovAccesscontrol(uint64(m.AccessType))
	}
	if m.ResourceType != 0 {
		n += 1 + sovAccesscontrol(uint64(m.ResourceType))
	}
	l = len(m.IdentifierTemplate)
	if l > 0 {
		n += 1 + l + sovAccesscontrol(uint64(l))
	}
	return n
}

func (m *MessageDependencyMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MessageKey)
	if l > 0 {
		n += 1 + l + sovAccesscontrol(uint64(l))
	}
	if len(m.AccessOps) > 0 {
		for _, e := range m.AccessOps {
			l = e.Size()
			n += 1 + l + sovAccesscontrol(uint64(l))
		}
	}
	return n
}

func sovAccesscontrol(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccesscontrol(x uint64) (n int) {
	return sovAccesscontrol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessOperation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccesscontrol
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
			return fmt.Errorf("proto: AccessOperation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessOperation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessType", wireType)
			}
			m.AccessType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessType |= AccessType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceType", wireType)
			}
			m.ResourceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResourceType |= ResourceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentifierTemplate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
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
				return ErrInvalidLengthAccesscontrol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccesscontrol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdentifierTemplate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccesscontrol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccesscontrol
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
func (m *MessageDependencyMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccesscontrol
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
			return fmt.Errorf("proto: MessageDependencyMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageDependencyMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
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
				return ErrInvalidLengthAccesscontrol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccesscontrol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessOps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
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
				return ErrInvalidLengthAccesscontrol
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccesscontrol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessOps = append(m.AccessOps, AccessOperation{})
			if err := m.AccessOps[len(m.AccessOps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccesscontrol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccesscontrol
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
func skipAccesscontrol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccesscontrol
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
					return 0, ErrIntOverflowAccesscontrol
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
					return 0, ErrIntOverflowAccesscontrol
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
				return 0, ErrInvalidLengthAccesscontrol
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccesscontrol
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccesscontrol
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccesscontrol        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccesscontrol          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccesscontrol = fmt.Errorf("proto: unexpected end of group")
)