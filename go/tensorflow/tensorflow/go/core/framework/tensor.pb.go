// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/tensor.proto

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

// Protocol buffer representing a tensor.
type TensorProto struct {
	Dtype DataType `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// Shape of the tensor.  TODO(touts): sort out the 0-rank issues.
	TensorShape *TensorShapeProto `protobuf:"bytes,2,opt,name=tensor_shape,json=tensorShape,proto3" json:"tensor_shape,omitempty"`
	// Version number.
	//
	// In version 0, if the "repeated xxx" representations contain only one
	// element, that element is repeated to fill the shape.  This makes it easy
	// to represent a constant Tensor with a single value.
	VersionNumber int32 `protobuf:"varint,3,opt,name=version_number,json=versionNumber,proto3" json:"version_number,omitempty"`
	// Serialized raw tensor content from either Tensor::AsProtoTensorContent or
	// memcpy in tensorflow::grpc::EncodeTensorToByteBuffer. This representation
	// can be used for all tensor types. The purpose of this representation is to
	// reduce serialization overhead during RPC call by avoiding serialization of
	// many repeated small items.
	TensorContent []byte `protobuf:"bytes,4,opt,name=tensor_content,json=tensorContent,proto3" json:"tensor_content,omitempty"`
	// DT_HALF, DT_BFLOAT16. Note that since protobuf has no int16 type, we'll
	// have some pointless zero padding for each value here.
	HalfVal []int32 `protobuf:"varint,13,rep,packed,name=half_val,json=halfVal,proto3" json:"half_val,omitempty"`
	// DT_FLOAT.
	FloatVal []float32 `protobuf:"fixed32,5,rep,packed,name=float_val,json=floatVal,proto3" json:"float_val,omitempty"`
	// DT_DOUBLE.
	DoubleVal []float64 `protobuf:"fixed64,6,rep,packed,name=double_val,json=doubleVal,proto3" json:"double_val,omitempty"`
	// DT_INT32, DT_INT16, DT_INT8, DT_UINT8.
	IntVal []int32 `protobuf:"varint,7,rep,packed,name=int_val,json=intVal,proto3" json:"int_val,omitempty"`
	// DT_STRING
	StringVal [][]byte `protobuf:"bytes,8,rep,name=string_val,json=stringVal,proto3" json:"string_val,omitempty"`
	// DT_COMPLEX64. scomplex_val(2*i) and scomplex_val(2*i+1) are real
	// and imaginary parts of i-th single precision complex.
	ScomplexVal []float32 `protobuf:"fixed32,9,rep,packed,name=scomplex_val,json=scomplexVal,proto3" json:"scomplex_val,omitempty"`
	// DT_INT64
	Int64Val []int64 `protobuf:"varint,10,rep,packed,name=int64_val,json=int64Val,proto3" json:"int64_val,omitempty"`
	// DT_BOOL
	BoolVal []bool `protobuf:"varint,11,rep,packed,name=bool_val,json=boolVal,proto3" json:"bool_val,omitempty"`
	// DT_COMPLEX128. dcomplex_val(2*i) and dcomplex_val(2*i+1) are real
	// and imaginary parts of i-th double precision complex.
	DcomplexVal []float64 `protobuf:"fixed64,12,rep,packed,name=dcomplex_val,json=dcomplexVal,proto3" json:"dcomplex_val,omitempty"`
	// DT_RESOURCE
	ResourceHandleVal []*ResourceHandleProto `protobuf:"bytes,14,rep,name=resource_handle_val,json=resourceHandleVal,proto3" json:"resource_handle_val,omitempty"`
	// DT_VARIANT
	VariantVal []*VariantTensorDataProto `protobuf:"bytes,15,rep,name=variant_val,json=variantVal,proto3" json:"variant_val,omitempty"`
	// DT_UINT32
	Uint32Val []uint32 `protobuf:"varint,16,rep,packed,name=uint32_val,json=uint32Val,proto3" json:"uint32_val,omitempty"`
	// DT_UINT64
	Uint64Val            []uint64 `protobuf:"varint,17,rep,packed,name=uint64_val,json=uint64Val,proto3" json:"uint64_val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TensorProto) Reset()         { *m = TensorProto{} }
func (m *TensorProto) String() string { return proto.CompactTextString(m) }
func (*TensorProto) ProtoMessage()    {}
func (*TensorProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_efa68180bc31e4fc, []int{0}
}

func (m *TensorProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TensorProto.Unmarshal(m, b)
}
func (m *TensorProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TensorProto.Marshal(b, m, deterministic)
}
func (m *TensorProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TensorProto.Merge(m, src)
}
func (m *TensorProto) XXX_Size() int {
	return xxx_messageInfo_TensorProto.Size(m)
}
func (m *TensorProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TensorProto.DiscardUnknown(m)
}

var xxx_messageInfo_TensorProto proto.InternalMessageInfo

func (m *TensorProto) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *TensorProto) GetTensorShape() *TensorShapeProto {
	if m != nil {
		return m.TensorShape
	}
	return nil
}

func (m *TensorProto) GetVersionNumber() int32 {
	if m != nil {
		return m.VersionNumber
	}
	return 0
}

func (m *TensorProto) GetTensorContent() []byte {
	if m != nil {
		return m.TensorContent
	}
	return nil
}

func (m *TensorProto) GetHalfVal() []int32 {
	if m != nil {
		return m.HalfVal
	}
	return nil
}

func (m *TensorProto) GetFloatVal() []float32 {
	if m != nil {
		return m.FloatVal
	}
	return nil
}

func (m *TensorProto) GetDoubleVal() []float64 {
	if m != nil {
		return m.DoubleVal
	}
	return nil
}

func (m *TensorProto) GetIntVal() []int32 {
	if m != nil {
		return m.IntVal
	}
	return nil
}

func (m *TensorProto) GetStringVal() [][]byte {
	if m != nil {
		return m.StringVal
	}
	return nil
}

func (m *TensorProto) GetScomplexVal() []float32 {
	if m != nil {
		return m.ScomplexVal
	}
	return nil
}

func (m *TensorProto) GetInt64Val() []int64 {
	if m != nil {
		return m.Int64Val
	}
	return nil
}

func (m *TensorProto) GetBoolVal() []bool {
	if m != nil {
		return m.BoolVal
	}
	return nil
}

func (m *TensorProto) GetDcomplexVal() []float64 {
	if m != nil {
		return m.DcomplexVal
	}
	return nil
}

func (m *TensorProto) GetResourceHandleVal() []*ResourceHandleProto {
	if m != nil {
		return m.ResourceHandleVal
	}
	return nil
}

func (m *TensorProto) GetVariantVal() []*VariantTensorDataProto {
	if m != nil {
		return m.VariantVal
	}
	return nil
}

func (m *TensorProto) GetUint32Val() []uint32 {
	if m != nil {
		return m.Uint32Val
	}
	return nil
}

func (m *TensorProto) GetUint64Val() []uint64 {
	if m != nil {
		return m.Uint64Val
	}
	return nil
}

// Protocol buffer representing the serialization format of DT_VARIANT tensors.
type VariantTensorDataProto struct {
	// Name of the type of objects being serialized.
	TypeName string `protobuf:"bytes,1,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// Portions of the object that are not Tensors.
	Metadata []byte `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Tensors contained within objects being serialized.
	Tensors              []*TensorProto `protobuf:"bytes,3,rep,name=tensors,proto3" json:"tensors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *VariantTensorDataProto) Reset()         { *m = VariantTensorDataProto{} }
func (m *VariantTensorDataProto) String() string { return proto.CompactTextString(m) }
func (*VariantTensorDataProto) ProtoMessage()    {}
func (*VariantTensorDataProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_efa68180bc31e4fc, []int{1}
}

func (m *VariantTensorDataProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VariantTensorDataProto.Unmarshal(m, b)
}
func (m *VariantTensorDataProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VariantTensorDataProto.Marshal(b, m, deterministic)
}
func (m *VariantTensorDataProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VariantTensorDataProto.Merge(m, src)
}
func (m *VariantTensorDataProto) XXX_Size() int {
	return xxx_messageInfo_VariantTensorDataProto.Size(m)
}
func (m *VariantTensorDataProto) XXX_DiscardUnknown() {
	xxx_messageInfo_VariantTensorDataProto.DiscardUnknown(m)
}

var xxx_messageInfo_VariantTensorDataProto proto.InternalMessageInfo

func (m *VariantTensorDataProto) GetTypeName() string {
	if m != nil {
		return m.TypeName
	}
	return ""
}

func (m *VariantTensorDataProto) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *VariantTensorDataProto) GetTensors() []*TensorProto {
	if m != nil {
		return m.Tensors
	}
	return nil
}

func init() {
	proto.RegisterType((*TensorProto)(nil), "tensorflow.TensorProto")
	proto.RegisterType((*VariantTensorDataProto)(nil), "tensorflow.VariantTensorDataProto")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/tensor.proto", fileDescriptor_efa68180bc31e4fc)
}

var fileDescriptor_efa68180bc31e4fc = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcf, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0x95, 0x79, 0x6d, 0x13, 0x27, 0xed, 0x77, 0xcb, 0x17, 0x41, 0xd4, 0x31, 0xcd, 0x54,
	0x2a, 0xb2, 0x10, 0x34, 0xa2, 0x43, 0x5c, 0x91, 0x3a, 0x84, 0x38, 0x8d, 0x29, 0x4c, 0x3b, 0x70,
	0xa9, 0x9c, 0xc4, 0x4d, 0x23, 0x12, 0x3b, 0x72, 0x9c, 0x8e, 0xdd, 0xf9, 0x07, 0xf9, 0x6f, 0x38,
	0x22, 0xdb, 0x69, 0x1b, 0xc6, 0xaf, 0x5b, 0xfb, 0xbc, 0x9f, 0xf7, 0x79, 0xfc, 0xe3, 0x75, 0xe0,
	0x53, 0x49, 0x59, 0xcd, 0xc5, 0xaa, 0xe0, 0xb7, 0x61, 0xc2, 0x05, 0x0d, 0x57, 0x82, 0x94, 0xf4,
	0x96, 0x8b, 0xcf, 0xa1, 0xa9, 0xcc, 0x2a, 0xc1, 0x25, 0xf7, 0xe1, 0x9e, 0x1b, 0x87, 0x7f, 0xee,
	0x11, 0xb4, 0xe6, 0x8d, 0x48, 0xe8, 0x72, 0x4d, 0x58, 0x5a, 0x50, 0xd3, 0x3c, 0x7e, 0xfe, 0xaf,
	0x90, 0x65, 0xbd, 0x26, 0xd5, 0x96, 0x9e, 0xfe, 0x85, 0xbe, 0xab, 0x68, 0x6d, 0xb0, 0xc9, 0xb7,
	0x1e, 0x74, 0xaf, 0x35, 0x79, 0xa5, 0x57, 0xf8, 0x0c, 0xf6, 0x52, 0x55, 0x0f, 0x2c, 0x64, 0xe1,
	0xd1, 0xfc, 0xc1, 0x6c, 0x6f, 0x33, 0x7b, 0x4b, 0x24, 0xb9, 0xbe, 0xab, 0x68, 0x64, 0x10, 0xff,
	0x0d, 0xf4, 0xba, 0xc1, 0xc1, 0x01, 0xb2, 0xb0, 0x3b, 0x7f, 0xdc, 0x6d, 0x31, 0xd6, 0x1f, 0x55,
	0x59, 0xfb, 0x47, 0xae, 0xdc, 0x2b, 0xfe, 0x14, 0x8e, 0x36, 0x54, 0xd4, 0x39, 0x67, 0x4b, 0xd6,
	0x94, 0x31, 0x15, 0x01, 0x40, 0x16, 0xee, 0x45, 0xc3, 0x56, 0xbd, 0xd4, 0xa2, 0xc2, 0xda, 0x9c,
	0x84, 0x33, 0x49, 0x99, 0x0c, 0x0e, 0x91, 0x85, 0xbd, 0x68, 0x68, 0xd4, 0x0b, 0x23, 0xfa, 0xa7,
	0xd0, 0x5e, 0x93, 0x62, 0xb5, 0xdc, 0x90, 0x22, 0x18, 0x22, 0x80, 0x7b, 0x8b, 0x83, 0x23, 0x2b,
	0x1a, 0x28, 0xed, 0x86, 0x14, 0xfe, 0x19, 0x74, 0x56, 0x05, 0x27, 0x52, 0xd7, 0x7b, 0x08, 0xe0,
	0x03, 0x5d, 0xb7, 0xb5, 0xa8, 0x80, 0x27, 0x10, 0xa6, 0xbc, 0x89, 0x0b, 0xaa, 0x89, 0x3e, 0x02,
	0xd8, 0xd2, 0x84, 0x63, 0x54, 0x85, 0x9c, 0xc0, 0x41, 0xce, 0x8c, 0xc3, 0x60, 0x97, 0xd0, 0xcf,
	0x99, 0xee, 0x3f, 0x85, 0xb0, 0x96, 0x22, 0x67, 0x99, 0xae, 0xdb, 0x08, 0x60, 0x2f, 0x72, 0x8c,
	0xa2, 0xca, 0x53, 0xe8, 0xd5, 0x09, 0x2f, 0xab, 0x82, 0x7e, 0xd1, 0x80, 0xb3, 0x5b, 0x82, 0xbb,
	0xd5, 0xdb, 0x65, 0xe6, 0x4c, 0xbe, 0x7e, 0xa5, 0x19, 0x88, 0x00, 0x06, 0x66, 0x99, 0x5a, 0x34,
	0x31, 0x76, 0xcc, 0x79, 0xa1, 0xeb, 0x2e, 0x02, 0xd8, 0x36, 0xdb, 0x54, 0x5a, 0x1b, 0x93, 0x76,
	0x63, 0xbc, 0xdd, 0x3e, 0xdc, 0xb4, 0x13, 0xf3, 0x01, 0xfe, 0x7f, 0x6f, 0xca, 0x34, 0x3d, 0x42,
	0x00, 0xbb, 0xf3, 0xb3, 0xee, 0x15, 0x46, 0x2d, 0xf6, 0x5e, 0x53, 0xe6, 0x16, 0x8f, 0xc5, 0x4f,
	0xa2, 0x32, 0xbc, 0x80, 0xee, 0x86, 0x88, 0x9c, 0xb4, 0xc7, 0xf3, 0x9f, 0x36, 0x9a, 0x74, 0x8d,
	0x6e, 0x4c, 0xd9, 0x8c, 0x84, 0x9a, 0x25, 0xe3, 0x05, 0xdb, 0xb6, 0xf6, 0x0a, 0x9a, 0x9c, 0xc9,
	0xf3, 0xb9, 0xf6, 0x38, 0x42, 0x00, 0x0f, 0xcd, 0x15, 0x18, 0xb5, 0x83, 0xb4, 0x07, 0x74, 0x8c,
	0x00, 0x3e, 0xdc, 0x23, 0xfa, 0x84, 0x26, 0x5f, 0x2d, 0xf8, 0xf0, 0xf7, 0x61, 0xfe, 0x09, 0x74,
	0xd4, 0xe8, 0x2e, 0x19, 0x29, 0xcd, 0x88, 0x3b, 0x91, 0xad, 0x84, 0x4b, 0x52, 0x52, 0x7f, 0x0c,
	0xed, 0x92, 0x4a, 0x92, 0x12, 0x49, 0xf4, 0x2c, 0x7b, 0xd1, 0xee, 0xbf, 0xff, 0x12, 0x0e, 0xcc,
	0x56, 0xea, 0x00, 0xe8, 0xad, 0x3d, 0xfa, 0x75, 0xcc, 0xcd, 0x7e, 0xb6, 0xdc, 0xa2, 0x81, 0x01,
	0x17, 0x59, 0x17, 0xdb, 0x3d, 0xc1, 0x85, 0xd7, 0xe9, 0xa8, 0xaf, 0xac, 0x4f, 0xef, 0xb2, 0x5c,
	0xae, 0x9b, 0x78, 0x96, 0xf0, 0x32, 0x64, 0x54, 0xc6, 0x82, 0xe4, 0x2c, 0x94, 0xab, 0x17, 0x99,
	0xa8, 0x92, 0x30, 0xe3, 0xdd, 0x6f, 0x45, 0xe7, 0x67, 0xc6, 0xef, 0x3d, 0xed, 0xef, 0x96, 0x15,
	0xf7, 0xf5, 0xc3, 0x3e, 0xff, 0x11, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x5f, 0xab, 0x57, 0x94, 0x04,
	0x00, 0x00,
}
