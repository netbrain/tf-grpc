// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/tensor_description.proto

package framework

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TensorDescription struct {
	// Data type of tensor elements
	Dtype DataType `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// Shape of the tensor.
	Shape *TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	// Information about the size and allocator used for the data
	AllocationDescription *AllocationDescription `protobuf:"bytes,4,opt,name=allocation_description,json=allocationDescription,proto3" json:"allocation_description,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *TensorDescription) Reset()         { *m = TensorDescription{} }
func (m *TensorDescription) String() string { return proto.CompactTextString(m) }
func (*TensorDescription) ProtoMessage()    {}
func (*TensorDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa203ffb9e427669, []int{0}
}

func (m *TensorDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TensorDescription.Unmarshal(m, b)
}
func (m *TensorDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TensorDescription.Marshal(b, m, deterministic)
}
func (m *TensorDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TensorDescription.Merge(m, src)
}
func (m *TensorDescription) XXX_Size() int {
	return xxx_messageInfo_TensorDescription.Size(m)
}
func (m *TensorDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_TensorDescription.DiscardUnknown(m)
}

var xxx_messageInfo_TensorDescription proto.InternalMessageInfo

func (m *TensorDescription) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *TensorDescription) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *TensorDescription) GetAllocationDescription() *AllocationDescription {
	if m != nil {
		return m.AllocationDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*TensorDescription)(nil), "tensorflow.TensorDescription")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/tensor_description.proto", fileDescriptor_aa203ffb9e427669)
}

var fileDescriptor_aa203ffb9e427669 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0x89, 0x38, 0x0f, 0x11, 0x04, 0x8b, 0x2f, 0x65, 0x78, 0x98, 0x82, 0x30, 0x44, 0x1b,
	0xa8, 0xe0, 0xdd, 0x31, 0x3c, 0x8f, 0xba, 0x83, 0x78, 0x91, 0x34, 0x4b, 0xb3, 0x62, 0xd7, 0x27,
	0x3c, 0x89, 0x8c, 0xdd, 0xfc, 0x82, 0x7e, 0x1f, 0x8f, 0x92, 0x44, 0xd6, 0xe0, 0xa6, 0xde, 0x4a,
	0xff, 0x2f, 0xcf, 0x2f, 0x7f, 0x9a, 0x5b, 0xd9, 0x1a, 0xc0, 0xaa, 0x81, 0x25, 0x13, 0x80, 0x92,
	0x55, 0xc8, 0x17, 0x72, 0x09, 0xf8, 0xca, 0x82, 0xf2, 0x32, 0x93, 0x46, 0x60, 0xad, 0x6d, 0x0d,
	0x6d, 0xa6, 0x11, 0x2c, 0x24, 0xb4, 0xcb, 0xf4, 0x2f, 0xff, 0xc8, 0xaf, 0xb4, 0x34, 0x21, 0xd2,
	0xbf, 0xfe, 0xf7, 0x8c, 0x99, 0x73, 0x2d, 0xbf, 0xdd, 0x77, 0xbf, 0xbb, 0x79, 0xd3, 0x80, 0xe0,
	0x0e, 0x66, 0x13, 0xec, 0xe2, 0x83, 0xd0, 0xc3, 0xa9, 0x8f, 0x8e, 0x3b, 0x2d, 0xb9, 0xa2, 0xbd,
	0x99, 0x63, 0x49, 0xc9, 0x80, 0x0c, 0x0f, 0xf2, 0xa3, 0xac, 0x6b, 0xcf, 0xc6, 0xdc, 0xf2, 0xe9,
	0x4a, 0xcb, 0x22, 0x58, 0x92, 0x9c, 0xf6, 0x3c, 0x48, 0xba, 0x33, 0x20, 0xc3, 0xfd, 0xfc, 0x2c,
	0xf6, 0x86, 0xe6, 0x47, 0x27, 0x4f, 0xdc, 0xb9, 0x22, 0x58, 0x93, 0x27, 0x7a, 0xb2, 0x9d, 0x2a,
	0xdd, 0xf5, 0x25, 0xe7, 0x71, 0xc9, 0xfd, 0xda, 0x19, 0x21, 0x16, 0xc7, 0x7c, 0xdb, 0xef, 0xd1,
	0x3b, 0xa1, 0x29, 0xa0, 0x8a, 0xf3, 0xeb, 0x25, 0x46, 0xa7, 0x1b, 0x2f, 0xf5, 0x54, 0x66, 0x42,
	0x9e, 0x1f, 0x54, 0x6d, 0xe7, 0x6f, 0x65, 0x26, 0x60, 0xc1, 0x5a, 0x69, 0x4b, 0xe4, 0x75, 0xcb,
	0x6c, 0x75, 0xa3, 0x50, 0x0b, 0xa6, 0x80, 0x45, 0xf3, 0x46, 0x9f, 0x0a, 0x7e, 0x8c, 0xfd, 0x49,
	0x48, 0xb9, 0xe7, 0x97, 0xbd, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x9f, 0xba, 0x93, 0x28,
	0x02, 0x00, 0x00,
}