// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/cockroachdb/cockroach/pkg/errors/errorspb/errors.proto

package errorspb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EncodedError struct {
	// An error is either...
	//
	// Types that are valid to be assigned to Error:
	//	*EncodedError_Leaf
	//	*EncodedError_Wrapper
	Error                isEncodedError_Error `protobuf_oneof:"error"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EncodedError) Reset()         { *m = EncodedError{} }
func (m *EncodedError) String() string { return proto.CompactTextString(m) }
func (*EncodedError) ProtoMessage()    {}
func (*EncodedError) Descriptor() ([]byte, []int) {
	return fileDescriptor_errors_5e4865ddee77a111, []int{0}
}
func (m *EncodedError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncodedError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *EncodedError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodedError.Merge(dst, src)
}
func (m *EncodedError) XXX_Size() int {
	return m.Size()
}
func (m *EncodedError) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodedError.DiscardUnknown(m)
}

var xxx_messageInfo_EncodedError proto.InternalMessageInfo

type isEncodedError_Error interface {
	isEncodedError_Error()
	MarshalTo([]byte) (int, error)
	Size() int
}

type EncodedError_Leaf struct {
	Leaf *EncodedErrorLeaf `protobuf:"bytes,1,opt,name=leaf,proto3,oneof"`
}
type EncodedError_Wrapper struct {
	Wrapper *EncodedWrapper `protobuf:"bytes,2,opt,name=wrapper,proto3,oneof"`
}

func (*EncodedError_Leaf) isEncodedError_Error()    {}
func (*EncodedError_Wrapper) isEncodedError_Error() {}

func (m *EncodedError) GetError() isEncodedError_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *EncodedError) GetLeaf() *EncodedErrorLeaf {
	if x, ok := m.GetError().(*EncodedError_Leaf); ok {
		return x.Leaf
	}
	return nil
}

func (m *EncodedError) GetWrapper() *EncodedWrapper {
	if x, ok := m.GetError().(*EncodedError_Wrapper); ok {
		return x.Wrapper
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EncodedError) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EncodedError_OneofMarshaler, _EncodedError_OneofUnmarshaler, _EncodedError_OneofSizer, []interface{}{
		(*EncodedError_Leaf)(nil),
		(*EncodedError_Wrapper)(nil),
	}
}

func _EncodedError_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EncodedError)
	// error
	switch x := m.Error.(type) {
	case *EncodedError_Leaf:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Leaf); err != nil {
			return err
		}
	case *EncodedError_Wrapper:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Wrapper); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EncodedError.Error has unexpected type %T", x)
	}
	return nil
}

func _EncodedError_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EncodedError)
	switch tag {
	case 1: // error.leaf
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EncodedErrorLeaf)
		err := b.DecodeMessage(msg)
		m.Error = &EncodedError_Leaf{msg}
		return true, err
	case 2: // error.wrapper
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EncodedWrapper)
		err := b.DecodeMessage(msg)
		m.Error = &EncodedError_Wrapper{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EncodedError_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EncodedError)
	// error
	switch x := m.Error.(type) {
	case *EncodedError_Leaf:
		s := proto.Size(x.Leaf)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EncodedError_Wrapper:
		s := proto.Size(x.Wrapper)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A leaf error has...
type EncodedErrorLeaf struct {
	// always a message, that can be printed to human users and may
	// contain PII. This contains the value of the leaf error's
	// Error(), or using a registered encoder.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// always a fully qualified go type name, which will
	// be used to look up a decoding function.
	TypeName string `protobuf:"bytes,2,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// optionally, a reportable payload, which is as descriptive as
	// possible but may not contain PII.
	//
	// This is extracted automatically using a registered encoder, if
	// any, or the SafeDetailer interface.
	ReportablePayload []string `protobuf:"bytes,3,rep,name=reportable_payload,json=reportablePayload,proto3" json:"reportable_payload,omitempty"`
	// optionally, an arbitrary payload.
	//
	// This is extracted automatically using a registered encoder, if
	// any.
	FullDetails          *types.Any `protobuf:"bytes,4,opt,name=full_details,json=fullDetails,proto3" json:"full_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EncodedErrorLeaf) Reset()         { *m = EncodedErrorLeaf{} }
func (m *EncodedErrorLeaf) String() string { return proto.CompactTextString(m) }
func (*EncodedErrorLeaf) ProtoMessage()    {}
func (*EncodedErrorLeaf) Descriptor() ([]byte, []int) {
	return fileDescriptor_errors_5e4865ddee77a111, []int{1}
}
func (m *EncodedErrorLeaf) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncodedErrorLeaf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *EncodedErrorLeaf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodedErrorLeaf.Merge(dst, src)
}
func (m *EncodedErrorLeaf) XXX_Size() int {
	return m.Size()
}
func (m *EncodedErrorLeaf) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodedErrorLeaf.DiscardUnknown(m)
}

var xxx_messageInfo_EncodedErrorLeaf proto.InternalMessageInfo

// An error wrapper has...
type EncodedWrapper struct {
	// always a cause, which is another error.
	// This is populated using Cause() or Unwrap().
	Cause EncodedError `protobuf:"bytes,1,opt,name=cause,proto3" json:"cause"`
	// always a message prefix (which may be empty), which
	// will be printed before the cause's own message when
	// constructing a full message. This may contain PII.
	//
	// This is extracted automatically:
	//
	// - for wrappers that have a registered encoder,
	// - otherwise, when the wrapper's Error() has its cause's Error() as suffix.
	MessagePrefix string `protobuf:"bytes,2,opt,name=message_prefix,json=messagePrefix,proto3" json:"message_prefix,omitempty"`
	// always a fully qualified go type name, which will
	// be used to look up a decoding function.
	TypeName string `protobuf:"bytes,3,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// optionally, a reportable payload, which is as descriptive as
	// possible but may not contain PII.
	//
	// This is extracted automatically for wrappers that have a
	// registered encoder, or that implement the SafeDetailer()
	// interface.
	ReportablePayload []string `protobuf:"bytes,4,rep,name=reportable_payload,json=reportablePayload,proto3" json:"reportable_payload,omitempty"`
	// optionally, an arbitrary payload.
	//
	// This is extracted automatically for wrappers that have
	// a registered encoder.
	FullDetails          *types.Any `protobuf:"bytes,5,opt,name=full_details,json=fullDetails,proto3" json:"full_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EncodedWrapper) Reset()         { *m = EncodedWrapper{} }
func (m *EncodedWrapper) String() string { return proto.CompactTextString(m) }
func (*EncodedWrapper) ProtoMessage()    {}
func (*EncodedWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_errors_5e4865ddee77a111, []int{2}
}
func (m *EncodedWrapper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncodedWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *EncodedWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodedWrapper.Merge(dst, src)
}
func (m *EncodedWrapper) XXX_Size() int {
	return m.Size()
}
func (m *EncodedWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodedWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_EncodedWrapper proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EncodedError)(nil), "cockroach.errorspb.EncodedError")
	proto.RegisterType((*EncodedErrorLeaf)(nil), "cockroach.errorspb.EncodedErrorLeaf")
	proto.RegisterType((*EncodedWrapper)(nil), "cockroach.errorspb.EncodedWrapper")
}
func (m *EncodedError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncodedError) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		nn1, err := m.Error.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *EncodedError_Leaf) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Leaf != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintErrors(dAtA, i, uint64(m.Leaf.Size()))
		n2, err := m.Leaf.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *EncodedError_Wrapper) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Wrapper != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintErrors(dAtA, i, uint64(m.Wrapper.Size()))
		n3, err := m.Wrapper.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *EncodedErrorLeaf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncodedErrorLeaf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if len(m.TypeName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintErrors(dAtA, i, uint64(len(m.TypeName)))
		i += copy(dAtA[i:], m.TypeName)
	}
	if len(m.ReportablePayload) > 0 {
		for _, s := range m.ReportablePayload {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.FullDetails != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintErrors(dAtA, i, uint64(m.FullDetails.Size()))
		n4, err := m.FullDetails.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *EncodedWrapper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncodedWrapper) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintErrors(dAtA, i, uint64(m.Cause.Size()))
	n5, err := m.Cause.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if len(m.MessagePrefix) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintErrors(dAtA, i, uint64(len(m.MessagePrefix)))
		i += copy(dAtA[i:], m.MessagePrefix)
	}
	if len(m.TypeName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintErrors(dAtA, i, uint64(len(m.TypeName)))
		i += copy(dAtA[i:], m.TypeName)
	}
	if len(m.ReportablePayload) > 0 {
		for _, s := range m.ReportablePayload {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.FullDetails != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintErrors(dAtA, i, uint64(m.FullDetails.Size()))
		n6, err := m.FullDetails.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func encodeVarintErrors(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EncodedError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Error != nil {
		n += m.Error.Size()
	}
	return n
}

func (m *EncodedError_Leaf) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Leaf != nil {
		l = m.Leaf.Size()
		n += 1 + l + sovErrors(uint64(l))
	}
	return n
}
func (m *EncodedError_Wrapper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Wrapper != nil {
		l = m.Wrapper.Size()
		n += 1 + l + sovErrors(uint64(l))
	}
	return n
}
func (m *EncodedErrorLeaf) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.TypeName)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if len(m.ReportablePayload) > 0 {
		for _, s := range m.ReportablePayload {
			l = len(s)
			n += 1 + l + sovErrors(uint64(l))
		}
	}
	if m.FullDetails != nil {
		l = m.FullDetails.Size()
		n += 1 + l + sovErrors(uint64(l))
	}
	return n
}

func (m *EncodedWrapper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Cause.Size()
	n += 1 + l + sovErrors(uint64(l))
	l = len(m.MessagePrefix)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.TypeName)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if len(m.ReportablePayload) > 0 {
		for _, s := range m.ReportablePayload {
			l = len(s)
			n += 1 + l + sovErrors(uint64(l))
		}
	}
	if m.FullDetails != nil {
		l = m.FullDetails.Size()
		n += 1 + l + sovErrors(uint64(l))
	}
	return n
}

func sovErrors(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozErrors(x uint64) (n int) {
	return sovErrors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EncodedError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EncodedError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncodedError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leaf", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &EncodedErrorLeaf{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Error = &EncodedError_Leaf{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wrapper", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &EncodedWrapper{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Error = &EncodedError_Wrapper{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
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
func (m *EncodedErrorLeaf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EncodedErrorLeaf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncodedErrorLeaf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TypeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportablePayload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReportablePayload = append(m.ReportablePayload, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FullDetails == nil {
				m.FullDetails = &types.Any{}
			}
			if err := m.FullDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
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
func (m *EncodedWrapper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EncodedWrapper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncodedWrapper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cause", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cause.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessagePrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessagePrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TypeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportablePayload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReportablePayload = append(m.ReportablePayload, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FullDetails == nil {
				m.FullDetails = &types.Any{}
			}
			if err := m.FullDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
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
func skipErrors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrors
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthErrors
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowErrors
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipErrors(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthErrors = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrors   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/cockroachdb/cockroach/pkg/errors/errorspb/errors.proto", fileDescriptor_errors_5e4865ddee77a111)
}

var fileDescriptor_errors_5e4865ddee77a111 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x33, 0x36, 0xb5, 0xcd, 0xb4, 0x16, 0x1d, 0x7a, 0x88, 0x55, 0x62, 0x09, 0x0a, 0xbd,
	0x98, 0x80, 0x1e, 0x04, 0x11, 0xc1, 0x62, 0xa1, 0x07, 0x91, 0x92, 0x8b, 0xe0, 0x25, 0x4c, 0x92,
	0x97, 0x58, 0x4c, 0x33, 0xc3, 0x24, 0x45, 0xf3, 0x2d, 0x04, 0x3f, 0xc9, 0x7e, 0x8b, 0x1e, 0xf7,
	0xb8, 0xa7, 0x65, 0x37, 0xfb, 0x2d, 0xf6, 0xb4, 0x64, 0x26, 0x61, 0xb7, 0xbb, 0xb0, 0xdb, 0x53,
	0x5e, 0xde, 0xfb, 0xfd, 0xf3, 0xf2, 0x7b, 0xf8, 0x25, 0x08, 0xc1, 0x44, 0xee, 0xaa, 0x07, 0x0f,
	0x9a, 0xc2, 0xe1, 0x82, 0x15, 0x8c, 0x90, 0x90, 0x85, 0xbf, 0x05, 0xa3, 0xe1, 0x2f, 0xa7, 0x05,
	0x26, 0xcf, 0x13, 0xc6, 0x92, 0x14, 0x5c, 0x49, 0x04, 0xdb, 0xd8, 0xa5, 0x59, 0xa9, 0xf0, 0xc9,
	0x38, 0x61, 0x09, 0x93, 0xa5, 0x5b, 0x57, 0xaa, 0x6b, 0xff, 0x47, 0x78, 0xb8, 0xc8, 0x42, 0x16,
	0x41, 0xb4, 0xa8, 0x3f, 0x42, 0x3e, 0x62, 0x3d, 0x05, 0x1a, 0x9b, 0x68, 0x8a, 0x66, 0x83, 0x77,
	0xaf, 0x9d, 0xbb, 0x4b, 0x9c, 0x9b, 0xfc, 0x37, 0xa0, 0xf1, 0x52, 0xf3, 0x64, 0x86, 0x7c, 0xc6,
	0xbd, 0x3f, 0x82, 0x72, 0x0e, 0xc2, 0x7c, 0x24, 0xe3, 0xf6, 0x3d, 0xf1, 0x1f, 0x8a, 0x5c, 0x6a,
	0x5e, 0x1b, 0x9a, 0xf7, 0x70, 0x57, 0x52, 0xf6, 0x11, 0xc2, 0x4f, 0x6f, 0x6f, 0x21, 0x26, 0xee,
	0x6d, 0x20, 0xcf, 0x69, 0x02, 0xf2, 0xe7, 0x0c, 0xaf, 0x7d, 0x25, 0x2f, 0xb0, 0x51, 0x94, 0x1c,
	0xfc, 0x8c, 0x6e, 0x40, 0x6e, 0x36, 0xbc, 0x7e, 0xdd, 0xf8, 0x4e, 0x37, 0x40, 0xde, 0x62, 0x22,
	0x80, 0x33, 0x51, 0xd0, 0x20, 0x05, 0x9f, 0xd3, 0x32, 0x65, 0x34, 0x32, 0x3b, 0xd3, 0xce, 0xcc,
	0xf0, 0x9e, 0x5d, 0x4f, 0x56, 0x6a, 0x40, 0x3e, 0xe0, 0x61, 0xbc, 0x4d, 0x53, 0x3f, 0x82, 0x82,
	0xae, 0xd3, 0xdc, 0xd4, 0xa5, 0xc8, 0xd8, 0x51, 0x87, 0x75, 0xda, 0xc3, 0x3a, 0x5f, 0xb2, 0xd2,
	0x1b, 0xd4, 0xe4, 0x57, 0x05, 0xda, 0x97, 0x08, 0x8f, 0xf6, 0xd5, 0xc8, 0x27, 0xdc, 0x0d, 0xe9,
	0x36, 0x87, 0xe6, 0x98, 0xd3, 0x87, 0x8e, 0x39, 0xd7, 0x77, 0xa7, 0xaf, 0x34, 0x4f, 0x85, 0xc8,
	0x1b, 0x3c, 0x6a, 0x04, 0x7d, 0x2e, 0x20, 0x5e, 0xff, 0x6d, 0xd4, 0x9e, 0x34, 0xdd, 0x95, 0x6c,
	0xee, 0xcb, 0x77, 0x0e, 0x92, 0xd7, 0x0f, 0x95, 0xef, 0x1e, 0x28, 0x3f, 0xb7, 0x77, 0xe7, 0x96,
	0xb6, 0xab, 0x2c, 0x74, 0x5c, 0x59, 0xe8, 0xa4, 0xb2, 0xd0, 0x59, 0x65, 0xa1, 0x7f, 0x17, 0x96,
	0xf6, 0xb3, 0xdf, 0x9a, 0x06, 0x8f, 0x65, 0xfc, 0xfd, 0x55, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43,
	0x37, 0xc9, 0xd8, 0xd6, 0x02, 0x00, 0x00,
}
