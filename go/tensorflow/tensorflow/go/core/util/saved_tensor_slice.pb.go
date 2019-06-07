// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/util/saved_tensor_slice.proto

package util // import "github.com/netbrain/tf-grpc/go/tensorflow/tensorflow/go/core/util"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import framework "github.com/netbrain/tf-grpc/go/tensorflow/tensorflow/go/core/framework"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Metadata describing the set of slices of the same tensor saved in a
// checkpoint file.
type SavedSliceMeta struct {
	// Name of the tensor.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Shape of the tensor
	Shape *framework.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	// Type of the tensor
	Type framework.DataType `protobuf:"varint,3,opt,name=type,proto3,enum=tensorflow.DataType" json:"type,omitempty"`
	// Explicit list of slices saved in the checkpoint file.
	Slice                []*framework.TensorSliceProto `protobuf:"bytes,4,rep,name=slice,proto3" json:"slice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SavedSliceMeta) Reset()         { *m = SavedSliceMeta{} }
func (m *SavedSliceMeta) String() string { return proto.CompactTextString(m) }
func (*SavedSliceMeta) ProtoMessage()    {}
func (*SavedSliceMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_saved_tensor_slice_292b038af2c5134a, []int{0}
}
func (m *SavedSliceMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedSliceMeta.Unmarshal(m, b)
}
func (m *SavedSliceMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedSliceMeta.Marshal(b, m, deterministic)
}
func (dst *SavedSliceMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedSliceMeta.Merge(dst, src)
}
func (m *SavedSliceMeta) XXX_Size() int {
	return xxx_messageInfo_SavedSliceMeta.Size(m)
}
func (m *SavedSliceMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedSliceMeta.DiscardUnknown(m)
}

var xxx_messageInfo_SavedSliceMeta proto.InternalMessageInfo

func (m *SavedSliceMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SavedSliceMeta) GetShape() *framework.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *SavedSliceMeta) GetType() framework.DataType {
	if m != nil {
		return m.Type
	}
	return framework.DataType_DT_INVALID
}

func (m *SavedSliceMeta) GetSlice() []*framework.TensorSliceProto {
	if m != nil {
		return m.Slice
	}
	return nil
}

// Metadata describing the set of tensor slices saved in a checkpoint file.
// It is always stored at the beginning of each checkpoint file.
type SavedTensorSliceMeta struct {
	// Each SavedSliceMeta describes the slices for one tensor.
	Tensor []*SavedSliceMeta `protobuf:"bytes,1,rep,name=tensor,proto3" json:"tensor,omitempty"`
	// Compatibility version of this checkpoint.  See core/public/version.h
	// for version history.
	Versions             *framework.VersionDef `protobuf:"bytes,2,opt,name=versions,proto3" json:"versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SavedTensorSliceMeta) Reset()         { *m = SavedTensorSliceMeta{} }
func (m *SavedTensorSliceMeta) String() string { return proto.CompactTextString(m) }
func (*SavedTensorSliceMeta) ProtoMessage()    {}
func (*SavedTensorSliceMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_saved_tensor_slice_292b038af2c5134a, []int{1}
}
func (m *SavedTensorSliceMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedTensorSliceMeta.Unmarshal(m, b)
}
func (m *SavedTensorSliceMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedTensorSliceMeta.Marshal(b, m, deterministic)
}
func (dst *SavedTensorSliceMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedTensorSliceMeta.Merge(dst, src)
}
func (m *SavedTensorSliceMeta) XXX_Size() int {
	return xxx_messageInfo_SavedTensorSliceMeta.Size(m)
}
func (m *SavedTensorSliceMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedTensorSliceMeta.DiscardUnknown(m)
}

var xxx_messageInfo_SavedTensorSliceMeta proto.InternalMessageInfo

func (m *SavedTensorSliceMeta) GetTensor() []*SavedSliceMeta {
	if m != nil {
		return m.Tensor
	}
	return nil
}

func (m *SavedTensorSliceMeta) GetVersions() *framework.VersionDef {
	if m != nil {
		return m.Versions
	}
	return nil
}

// Saved tensor slice: it stores the name of the tensors, the slice, and the
// raw data.
type SavedSlice struct {
	// Name of the tensor that this slice belongs to. This must be identical to
	// the name used to encode the key for this record.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Extent of the slice.  Must have one entry for each of the dimension of the
	// tensor that this slice belongs to.
	Slice *framework.TensorSliceProto `protobuf:"bytes,2,opt,name=slice,proto3" json:"slice,omitempty"`
	// The raw data of the slice is stored as a TensorProto. Only raw data are
	// stored (we don't fill in fields such as dtype or tensor_shape).
	Data                 *framework.TensorProto `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SavedSlice) Reset()         { *m = SavedSlice{} }
func (m *SavedSlice) String() string { return proto.CompactTextString(m) }
func (*SavedSlice) ProtoMessage()    {}
func (*SavedSlice) Descriptor() ([]byte, []int) {
	return fileDescriptor_saved_tensor_slice_292b038af2c5134a, []int{2}
}
func (m *SavedSlice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedSlice.Unmarshal(m, b)
}
func (m *SavedSlice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedSlice.Marshal(b, m, deterministic)
}
func (dst *SavedSlice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedSlice.Merge(dst, src)
}
func (m *SavedSlice) XXX_Size() int {
	return xxx_messageInfo_SavedSlice.Size(m)
}
func (m *SavedSlice) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedSlice.DiscardUnknown(m)
}

var xxx_messageInfo_SavedSlice proto.InternalMessageInfo

func (m *SavedSlice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SavedSlice) GetSlice() *framework.TensorSliceProto {
	if m != nil {
		return m.Slice
	}
	return nil
}

func (m *SavedSlice) GetData() *framework.TensorProto {
	if m != nil {
		return m.Data
	}
	return nil
}

// Each record in a v3 checkpoint file is a serialized SavedTensorSlices
// message.
type SavedTensorSlices struct {
	// This is only present at the first item of each checkpoint file and serves
	// as a table of contents, listing all the tensor slices saved in this file.
	Meta *SavedTensorSliceMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// This exists in all but the first item of each checkpoint file.
	Data                 *SavedSlice `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SavedTensorSlices) Reset()         { *m = SavedTensorSlices{} }
func (m *SavedTensorSlices) String() string { return proto.CompactTextString(m) }
func (*SavedTensorSlices) ProtoMessage()    {}
func (*SavedTensorSlices) Descriptor() ([]byte, []int) {
	return fileDescriptor_saved_tensor_slice_292b038af2c5134a, []int{3}
}
func (m *SavedTensorSlices) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedTensorSlices.Unmarshal(m, b)
}
func (m *SavedTensorSlices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedTensorSlices.Marshal(b, m, deterministic)
}
func (dst *SavedTensorSlices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedTensorSlices.Merge(dst, src)
}
func (m *SavedTensorSlices) XXX_Size() int {
	return xxx_messageInfo_SavedTensorSlices.Size(m)
}
func (m *SavedTensorSlices) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedTensorSlices.DiscardUnknown(m)
}

var xxx_messageInfo_SavedTensorSlices proto.InternalMessageInfo

func (m *SavedTensorSlices) GetMeta() *SavedTensorSliceMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *SavedTensorSlices) GetData() *SavedSlice {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SavedSliceMeta)(nil), "tensorflow.SavedSliceMeta")
	proto.RegisterType((*SavedTensorSliceMeta)(nil), "tensorflow.SavedTensorSliceMeta")
	proto.RegisterType((*SavedSlice)(nil), "tensorflow.SavedSlice")
	proto.RegisterType((*SavedTensorSlices)(nil), "tensorflow.SavedTensorSlices")
}

func init() {
	proto.RegisterFile("tensorflow/core/util/saved_tensor_slice.proto", fileDescriptor_saved_tensor_slice_292b038af2c5134a)
}

var fileDescriptor_saved_tensor_slice_292b038af2c5134a = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x4a, 0xfb, 0x30,
	0x14, 0xc6, 0xc9, 0x7f, 0xfd, 0x0f, 0xcd, 0x60, 0x60, 0x1c, 0xb3, 0x0c, 0x2f, 0x4a, 0x41, 0x29,
	0xea, 0x1a, 0xa8, 0xbe, 0x80, 0x63, 0xb7, 0xc2, 0xe8, 0x86, 0x17, 0xde, 0x8c, 0xac, 0x4b, 0xbb,
	0xe2, 0xda, 0x94, 0x34, 0xdd, 0xd8, 0x8d, 0xe0, 0x4b, 0xf9, 0x6c, 0x5e, 0x4a, 0xd2, 0xce, 0xd6,
	0xce, 0xaa, 0x77, 0x07, 0xce, 0xef, 0xfb, 0xce, 0xf9, 0x0e, 0x09, 0x1c, 0x0a, 0x1a, 0xa7, 0x8c,
	0xfb, 0x6b, 0xb6, 0xc5, 0x1e, 0xe3, 0x14, 0x67, 0x22, 0x5c, 0xe3, 0x94, 0x6c, 0xe8, 0x72, 0x9e,
	0xb7, 0xe6, 0xe9, 0x3a, 0xf4, 0xa8, 0x9d, 0x70, 0x26, 0x18, 0x82, 0x25, 0x3e, 0xb8, 0xa9, 0x4b,
	0x7d, 0x4e, 0x22, 0xba, 0x65, 0xfc, 0x19, 0xef, 0x95, 0x2b, 0x92, 0x14, 0xca, 0xbf, 0xd0, 0xe5,
	0x9c, 0xc1, 0xe5, 0x6f, 0x74, 0xc1, 0x5d, 0xfc, 0xc0, 0xed, 0x12, 0x9a, 0x16, 0x98, 0xd5, 0x8c,
	0x6d, 0x28, 0x4f, 0x43, 0x16, 0x17, 0xa4, 0xf9, 0x06, 0x60, 0x77, 0x2a, 0xd3, 0x4f, 0xe5, 0x36,
	0x0f, 0x54, 0x10, 0x84, 0xa0, 0x16, 0x93, 0x88, 0xea, 0xc0, 0x00, 0xd6, 0xb1, 0xab, 0x6a, 0xe4,
	0xc0, 0xff, 0x2a, 0x9c, 0xfe, 0xcf, 0x00, 0x56, 0xc7, 0x39, 0xb7, 0xcb, 0x01, 0xf6, 0x4c, 0x95,
	0x53, 0xd9, 0x9e, 0x48, 0x4f, 0x37, 0x47, 0x91, 0x05, 0x35, 0xb9, 0x93, 0xde, 0x32, 0x80, 0xd5,
	0x75, 0x7a, 0x55, 0xc9, 0x98, 0x08, 0x32, 0xdb, 0x25, 0xd4, 0x55, 0x84, 0x72, 0x97, 0xe3, 0x75,
	0xcd, 0x68, 0x35, 0xb8, 0xcb, 0xf6, 0xde, 0x5d, 0xd6, 0xe6, 0x0b, 0xec, 0xa9, 0xbd, 0x2b, 0x7d,
	0xb5, 0xbd, 0x03, 0xdb, 0xb9, 0x5a, 0x07, 0xca, 0x6c, 0x50, 0x35, 0xfb, 0x9a, 0xd4, 0x2d, 0x48,
	0xe4, 0xc0, 0xa3, 0xfd, 0x59, 0x8a, 0x80, 0xfd, 0xaa, 0xea, 0x31, 0xef, 0x8d, 0xa9, 0xef, 0x7e,
	0x72, 0xe6, 0x2b, 0x80, 0xb0, 0xb4, 0x6b, 0x3c, 0x9a, 0x8a, 0xd5, 0x7c, 0xb4, 0x7a, 0x2c, 0x74,
	0x0d, 0xb5, 0x25, 0x11, 0x44, 0x1d, 0xad, 0xe3, 0x9c, 0x1d, 0x4a, 0x72, 0x5a, 0x41, 0x66, 0x06,
	0x4f, 0xea, 0x37, 0x48, 0xd1, 0x1d, 0xd4, 0x22, 0x2a, 0x88, 0xda, 0xa4, 0xe3, 0x18, 0x07, 0xf1,
	0x6b, 0x07, 0x73, 0x15, 0x8d, 0xae, 0x8a, 0xb9, 0xdf, 0xc4, 0x2f, 0x53, 0xe6, 0x63, 0x47, 0x19,
	0x3c, 0x65, 0x3c, 0xa8, 0x22, 0xf2, 0x13, 0x8d, 0xfa, 0x75, 0x7b, 0xb5, 0x6a, 0x3a, 0x01, 0x4f,
	0xf7, 0x41, 0x28, 0x56, 0xd9, 0xc2, 0xf6, 0x58, 0x84, 0x63, 0x2a, 0x16, 0x9c, 0x84, 0x31, 0x16,
	0xfe, 0x30, 0xe0, 0x89, 0x87, 0x03, 0x86, 0x2b, 0xaf, 0xb5, 0x52, 0x06, 0xac, 0xfc, 0xa1, 0xef,
	0x00, 0x2c, 0xda, 0xea, 0xc5, 0xde, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xbe, 0xa7, 0x9c,
	0xc3, 0x03, 0x00, 0x00,
}
