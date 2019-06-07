// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/tensor_shape.proto

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

// Dimensions of a tensor.
type TensorShapeProto struct {
	// Dimensions of the tensor, such as {"input", 30}, {"output", 40}
	// for a 30 x 40 2D tensor.  If an entry has size -1, this
	// corresponds to a dimension of unknown size. The names are
	// optional.
	//
	// The order of entries in "dim" matters: It indicates the layout of the
	// values in the tensor in-memory representation.
	//
	// The first entry in "dim" is the outermost dimension used to layout the
	// values, the last entry is the innermost dimension.  This matches the
	// in-memory layout of RowMajor Eigen tensors.
	//
	// If "dim.size()" > 0, "unknown_rank" must be false.
	Dim []*TensorShapeProto_Dim `protobuf:"bytes,2,rep,name=dim,proto3" json:"dim,omitempty"`
	// If true, the number of dimensions in the shape is unknown.
	//
	// If true, "dim.size()" must be 0.
	UnknownRank          bool     `protobuf:"varint,3,opt,name=unknown_rank,json=unknownRank,proto3" json:"unknown_rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TensorShapeProto) Reset()         { *m = TensorShapeProto{} }
func (m *TensorShapeProto) String() string { return proto.CompactTextString(m) }
func (*TensorShapeProto) ProtoMessage()    {}
func (*TensorShapeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd43873e75c1f7ac, []int{0}
}

func (m *TensorShapeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TensorShapeProto.Unmarshal(m, b)
}
func (m *TensorShapeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TensorShapeProto.Marshal(b, m, deterministic)
}
func (m *TensorShapeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TensorShapeProto.Merge(m, src)
}
func (m *TensorShapeProto) XXX_Size() int {
	return xxx_messageInfo_TensorShapeProto.Size(m)
}
func (m *TensorShapeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TensorShapeProto.DiscardUnknown(m)
}

var xxx_messageInfo_TensorShapeProto proto.InternalMessageInfo

func (m *TensorShapeProto) GetDim() []*TensorShapeProto_Dim {
	if m != nil {
		return m.Dim
	}
	return nil
}

func (m *TensorShapeProto) GetUnknownRank() bool {
	if m != nil {
		return m.UnknownRank
	}
	return false
}

// One dimension of the tensor.
type TensorShapeProto_Dim struct {
	// Size of the tensor in that dimension.
	// This value must be >= -1, but values of -1 are reserved for "unknown"
	// shapes (values of -1 mean "unknown" dimension).  Certain wrappers
	// that work with TensorShapeProto may fail at runtime when deserializing
	// a TensorShapeProto containing a dim value of -1.
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Optional name of the tensor dimension.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TensorShapeProto_Dim) Reset()         { *m = TensorShapeProto_Dim{} }
func (m *TensorShapeProto_Dim) String() string { return proto.CompactTextString(m) }
func (*TensorShapeProto_Dim) ProtoMessage()    {}
func (*TensorShapeProto_Dim) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd43873e75c1f7ac, []int{0, 0}
}

func (m *TensorShapeProto_Dim) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TensorShapeProto_Dim.Unmarshal(m, b)
}
func (m *TensorShapeProto_Dim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TensorShapeProto_Dim.Marshal(b, m, deterministic)
}
func (m *TensorShapeProto_Dim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TensorShapeProto_Dim.Merge(m, src)
}
func (m *TensorShapeProto_Dim) XXX_Size() int {
	return xxx_messageInfo_TensorShapeProto_Dim.Size(m)
}
func (m *TensorShapeProto_Dim) XXX_DiscardUnknown() {
	xxx_messageInfo_TensorShapeProto_Dim.DiscardUnknown(m)
}

var xxx_messageInfo_TensorShapeProto_Dim proto.InternalMessageInfo

func (m *TensorShapeProto_Dim) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *TensorShapeProto_Dim) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*TensorShapeProto)(nil), "tensorflow.TensorShapeProto")
	proto.RegisterType((*TensorShapeProto_Dim)(nil), "tensorflow.TensorShapeProto.Dim")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/tensor_shape.proto", fileDescriptor_cd43873e75c1f7ac)
}

var fileDescriptor_cd43873e75c1f7ac = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0x1a, 0x21, 0x70, 0x19, 0x20, 0x93, 0xc5, 0x64, 0x98, 0x32, 0x50, 0x5b, 0x2a,
	0x6f, 0x50, 0x55, 0xcc, 0x95, 0x61, 0x62, 0xa9, 0x9c, 0xe0, 0xb8, 0x56, 0xb0, 0x2f, 0xba, 0xb8,
	0x8a, 0xd4, 0x27, 0xe1, 0x51, 0x19, 0x51, 0x0c, 0xa2, 0x51, 0xb6, 0xdf, 0xbf, 0x3f, 0xe9, 0xee,
	0x3e, 0xf6, 0x94, 0x6c, 0xec, 0x01, 0x9b, 0x4f, 0x18, 0x54, 0x0d, 0x68, 0x55, 0x83, 0x26, 0xd8,
	0x01, 0xb0, 0x55, 0xbf, 0x3f, 0xfb, 0xfe, 0x60, 0x3a, 0x2b, 0x3b, 0x84, 0x04, 0x05, 0x3b, 0xd3,
	0x8f, 0x5f, 0x84, 0xdd, 0xbe, 0xe5, 0xe7, 0xeb, 0x48, 0xec, 0x32, 0xb0, 0x66, 0xf4, 0xc3, 0x07,
	0xbe, 0x10, 0xb4, 0x5c, 0xae, 0x85, 0x3c, 0xe3, 0x72, 0x8e, 0xca, 0xad, 0x0f, 0x7a, 0x84, 0x8b,
	0x07, 0x76, 0x73, 0x8c, 0x6d, 0x84, 0x21, 0xee, 0xd1, 0xc4, 0x96, 0x53, 0x41, 0xca, 0x2b, 0xbd,
	0xfc, 0xeb, 0xb4, 0x89, 0xed, 0xfd, 0x8a, 0xd1, 0xad, 0x0f, 0x45, 0xc1, 0x2e, 0x7a, 0x7f, 0xb2,
	0x9c, 0x08, 0x52, 0x52, 0x9d, 0xf3, 0xd8, 0x45, 0x13, 0x2c, 0x5f, 0x08, 0x52, 0x5e, 0xeb, 0x9c,
	0x37, 0x27, 0xc6, 0x01, 0xdd, 0x74, 0xfa, 0xff, 0x55, 0x9b, 0xbb, 0xf9, 0x22, 0xfd, 0x8e, 0xbc,
	0xbf, 0x38, 0x9f, 0x0e, 0xc7, 0x4a, 0xd6, 0x10, 0x54, 0xb4, 0xa9, 0x42, 0xe3, 0xa3, 0x4a, 0xcd,
	0xca, 0x61, 0x57, 0x2b, 0x07, 0x6a, 0x22, 0x69, 0x12, 0x1d, 0xcc, 0x94, 0x7d, 0x13, 0x52, 0x5d,
	0x66, 0x53, 0xcf, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xfc, 0x8e, 0x03, 0x59, 0x01, 0x00,
	0x00,
}