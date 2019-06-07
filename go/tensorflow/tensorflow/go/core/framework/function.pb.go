// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/function.proto

package framework // import "github.com/netbrain/tf-grpc/go/tensorflow/tensorflow/go/core/framework"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A library is a set of named functions.
type FunctionDefLibrary struct {
	Function             []*FunctionDef `protobuf:"bytes,1,rep,name=function,proto3" json:"function,omitempty"`
	Gradient             []*GradientDef `protobuf:"bytes,2,rep,name=gradient,proto3" json:"gradient,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FunctionDefLibrary) Reset()         { *m = FunctionDefLibrary{} }
func (m *FunctionDefLibrary) String() string { return proto.CompactTextString(m) }
func (*FunctionDefLibrary) ProtoMessage()    {}
func (*FunctionDefLibrary) Descriptor() ([]byte, []int) {
	return fileDescriptor_function_1f66fc09c7489ee5, []int{0}
}
func (m *FunctionDefLibrary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionDefLibrary.Unmarshal(m, b)
}
func (m *FunctionDefLibrary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionDefLibrary.Marshal(b, m, deterministic)
}
func (dst *FunctionDefLibrary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionDefLibrary.Merge(dst, src)
}
func (m *FunctionDefLibrary) XXX_Size() int {
	return xxx_messageInfo_FunctionDefLibrary.Size(m)
}
func (m *FunctionDefLibrary) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionDefLibrary.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionDefLibrary proto.InternalMessageInfo

func (m *FunctionDefLibrary) GetFunction() []*FunctionDef {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *FunctionDefLibrary) GetGradient() []*GradientDef {
	if m != nil {
		return m.Gradient
	}
	return nil
}

// A function can be instantiated when the runtime can bind every attr
// with a value. When a GraphDef has a call to a function, it must
// have binding for every attr defined in the signature.
//
// TODO(zhifengc):
//   * device spec, etc.
type FunctionDef struct {
	// The definition of the function's name, arguments, return values,
	// attrs etc.
	Signature *OpDef `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Attributes specific to this function definition.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// By convention, "op" in node_def is resolved by consulting with a
	// user-defined library first. If not resolved, "func" is assumed to
	// be a builtin op.
	NodeDef []*NodeDef `protobuf:"bytes,3,rep,name=node_def,json=nodeDef,proto3" json:"node_def,omitempty"`
	// A mapping from the output arg names from `signature` to the
	// outputs from `node_def` that should be returned by the function.
	Ret                  map[string]string `protobuf:"bytes,4,rep,name=ret,proto3" json:"ret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FunctionDef) Reset()         { *m = FunctionDef{} }
func (m *FunctionDef) String() string { return proto.CompactTextString(m) }
func (*FunctionDef) ProtoMessage()    {}
func (*FunctionDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_function_1f66fc09c7489ee5, []int{1}
}
func (m *FunctionDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionDef.Unmarshal(m, b)
}
func (m *FunctionDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionDef.Marshal(b, m, deterministic)
}
func (dst *FunctionDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionDef.Merge(dst, src)
}
func (m *FunctionDef) XXX_Size() int {
	return xxx_messageInfo_FunctionDef.Size(m)
}
func (m *FunctionDef) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionDef.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionDef proto.InternalMessageInfo

func (m *FunctionDef) GetSignature() *OpDef {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FunctionDef) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *FunctionDef) GetNodeDef() []*NodeDef {
	if m != nil {
		return m.NodeDef
	}
	return nil
}

func (m *FunctionDef) GetRet() map[string]string {
	if m != nil {
		return m.Ret
	}
	return nil
}

// GradientDef defines the gradient function of a function defined in
// a function library.
//
// A gradient function g (specified by gradient_func) for a function f
// (specified by function_name) must follow the following:
//
// The function 'f' must be a numerical function which takes N inputs
// and produces M outputs. Its gradient function 'g', which is a
// function taking N + M inputs and produces N outputs.
//
// I.e. if we have
//    (y1, y2, ..., y_M) = f(x1, x2, ..., x_N),
// then, g is
//    (dL/dx1, dL/dx2, ..., dL/dx_N) = g(x1, x2, ..., x_N,
//                                      dL/dy1, dL/dy2, ..., dL/dy_M),
// where L is a scalar-value function of (x1, x2, ..., xN) (e.g., the
// loss function). dL/dx_i is the partial derivative of L with respect
// to x_i.
type GradientDef struct {
	FunctionName         string   `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	GradientFunc         string   `protobuf:"bytes,2,opt,name=gradient_func,json=gradientFunc,proto3" json:"gradient_func,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GradientDef) Reset()         { *m = GradientDef{} }
func (m *GradientDef) String() string { return proto.CompactTextString(m) }
func (*GradientDef) ProtoMessage()    {}
func (*GradientDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_function_1f66fc09c7489ee5, []int{2}
}
func (m *GradientDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientDef.Unmarshal(m, b)
}
func (m *GradientDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientDef.Marshal(b, m, deterministic)
}
func (dst *GradientDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientDef.Merge(dst, src)
}
func (m *GradientDef) XXX_Size() int {
	return xxx_messageInfo_GradientDef.Size(m)
}
func (m *GradientDef) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientDef.DiscardUnknown(m)
}

var xxx_messageInfo_GradientDef proto.InternalMessageInfo

func (m *GradientDef) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *GradientDef) GetGradientFunc() string {
	if m != nil {
		return m.GradientFunc
	}
	return ""
}

func init() {
	proto.RegisterType((*FunctionDefLibrary)(nil), "tensorflow.FunctionDefLibrary")
	proto.RegisterType((*FunctionDef)(nil), "tensorflow.FunctionDef")
	proto.RegisterMapType((map[string]*AttrValue)(nil), "tensorflow.FunctionDef.AttrEntry")
	proto.RegisterMapType((map[string]string)(nil), "tensorflow.FunctionDef.RetEntry")
	proto.RegisterType((*GradientDef)(nil), "tensorflow.GradientDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/function.proto", fileDescriptor_function_1f66fc09c7489ee5)
}

var fileDescriptor_function_1f66fc09c7489ee5 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8b, 0x13, 0x31,
	0x14, 0xc6, 0x99, 0x99, 0xae, 0xb6, 0xaf, 0xab, 0xac, 0x51, 0x31, 0xf4, 0x54, 0x2b, 0x48, 0x51,
	0x9c, 0x40, 0x17, 0x45, 0xbc, 0xb9, 0xe8, 0x0a, 0x22, 0x75, 0x99, 0x83, 0x82, 0x97, 0x92, 0xb6,
	0x6f, 0xc6, 0x61, 0xdb, 0xa4, 0xbc, 0xa6, 0x96, 0x5e, 0xfc, 0x77, 0xfd, 0x17, 0x3c, 0x4a, 0x32,
	0x93, 0x4e, 0x50, 0x67, 0x6f, 0x21, 0xf9, 0x7d, 0xdf, 0xf7, 0xf2, 0xde, 0x83, 0xb1, 0x41, 0xb5,
	0xd5, 0x94, 0xaf, 0xf4, 0x5e, 0x2c, 0x34, 0xa1, 0xc8, 0x49, 0xae, 0x71, 0xaf, 0xe9, 0x5a, 0xe4,
	0x3b, 0xb5, 0x30, 0xa5, 0x56, 0xe9, 0x86, 0xb4, 0xd1, 0x0c, 0x1a, 0x72, 0xf0, 0xac, 0x5d, 0x25,
	0x8d, 0xa1, 0xd9, 0x0f, 0xb9, 0xda, 0x61, 0xa5, 0x1b, 0xdc, 0x90, 0xa0, 0xf4, 0x12, 0x67, 0x4b,
	0xcc, 0x6b, 0xf2, 0x69, 0x3b, 0xa9, 0x37, 0x0d, 0x37, 0xfa, 0x09, 0xec, 0xb2, 0xae, 0xed, 0x1d,
	0xe6, 0x9f, 0xca, 0x39, 0x49, 0x3a, 0xb0, 0x73, 0xe8, 0xfa, 0x8a, 0x79, 0x34, 0x4c, 0xc6, 0xfd,
	0xc9, 0xa3, 0xb4, 0x31, 0x4c, 0x03, 0x45, 0x76, 0x04, 0xad, 0xa8, 0x20, 0xb9, 0x2c, 0x51, 0x19,
	0x1e, 0xff, 0x2b, 0xfa, 0x50, 0xbf, 0x39, 0x91, 0x07, 0x47, 0xbf, 0x62, 0xe8, 0x07, 0x76, 0x4c,
	0x40, 0x6f, 0x5b, 0x16, 0x4a, 0x9a, 0x1d, 0x21, 0x8f, 0x86, 0xd1, 0xb8, 0x3f, 0xb9, 0x17, 0xba,
	0x7c, 0xde, 0x58, 0x7d, 0xc3, 0xb0, 0x97, 0xd0, 0xb1, 0x6d, 0xe2, 0x27, 0x2e, 0xf1, 0x71, 0x4b,
	0x99, 0xe9, 0x5b, 0x63, 0xe8, 0xbd, 0x32, 0x74, 0xc8, 0x1c, 0xce, 0x52, 0xe8, 0xfa, 0x8e, 0xf1,
	0xc4, 0x49, 0xef, 0x87, 0xd2, 0xa9, 0x5e, 0xa2, 0x0d, 0xba, 0xad, 0xaa, 0x03, 0x9b, 0x40, 0x42,
	0x68, 0x78, 0xc7, 0xa1, 0xc3, 0xb6, 0x94, 0x0c, 0x4d, 0x15, 0x62, 0xe1, 0xc1, 0x14, 0x7a, 0xc7,
	0x58, 0x76, 0x06, 0xc9, 0x35, 0x1e, 0xdc, 0x97, 0x7a, 0x99, 0x3d, 0xb2, 0xe7, 0x70, 0xe2, 0x66,
	0xcb, 0x63, 0xf7, 0xcd, 0x87, 0xa1, 0xa9, 0xd5, 0x7d, 0xb1, 0x8f, 0x59, 0xc5, 0xbc, 0x89, 0x5f,
	0x47, 0x83, 0x57, 0xd0, 0xf5, 0x01, 0xff, 0xb1, 0x7b, 0x10, 0xda, 0xf5, 0x02, 0xdd, 0xc7, 0x4e,
	0x37, 0x3e, 0x4b, 0x46, 0x5f, 0xa1, 0x1f, 0x8c, 0x80, 0x3d, 0x81, 0x3b, 0x7e, 0x72, 0x33, 0x25,
	0xd7, 0x58, 0x5b, 0x9d, 0xfa, 0xcb, 0xa9, 0x5c, 0xa3, 0x85, 0xfc, 0xa4, 0x66, 0xf6, 0xa1, 0xf6,
	0x3e, 0xf5, 0x97, 0xf6, 0xef, 0x17, 0x7b, 0xe0, 0x9a, 0x8a, 0xb0, 0xfa, 0xe3, 0xae, 0x5d, 0xdc,
	0xf5, 0xdd, 0xb9, 0xb2, 0xdb, 0xb6, 0xbd, 0x8a, 0xbe, 0x5d, 0x16, 0xa5, 0xf9, 0xbe, 0x9b, 0xa7,
	0x0b, 0xbd, 0x16, 0x0a, 0xcd, 0x9c, 0x64, 0xa9, 0x84, 0xc9, 0x5f, 0x14, 0xb4, 0x59, 0x88, 0x42,
	0x8b, 0x60, 0x6f, 0x83, 0x63, 0xa1, 0xff, 0xda, 0xe2, 0xdf, 0x51, 0x34, 0xbf, 0xe5, 0x56, 0xf8,
	0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0xc3, 0x1c, 0xd0, 0x78, 0x03, 0x00, 0x00,
}
