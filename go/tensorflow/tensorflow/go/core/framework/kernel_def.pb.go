// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/kernel_def.proto

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

type KernelDef struct {
	// Must match the name of an Op.
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// Type of device this kernel runs on.
	DeviceType string                      `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	Constraint []*KernelDef_AttrConstraint `protobuf:"bytes,3,rep,name=constraint,proto3" json:"constraint,omitempty"`
	// Names of the Op's input_/output_args that reside in host memory
	// instead of device memory.
	HostMemoryArg []string `protobuf:"bytes,4,rep,name=host_memory_arg,json=hostMemoryArg,proto3" json:"host_memory_arg,omitempty"`
	// This allows experimental kernels to be registered for an op that
	// won't be used unless the user specifies a "_kernel" attr with
	// value matching this.
	Label string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	// Prioritization of kernel amongst different devices. By default we assume
	// priority is 0. The higher the priority the better. By default (i.e. if
	// this is not set), we prefer GPU kernels over CPU.
	Priority             int32    `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KernelDef) Reset()         { *m = KernelDef{} }
func (m *KernelDef) String() string { return proto.CompactTextString(m) }
func (*KernelDef) ProtoMessage()    {}
func (*KernelDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_kernel_def_a3c41f0c94cd5ffa, []int{0}
}
func (m *KernelDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelDef.Unmarshal(m, b)
}
func (m *KernelDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelDef.Marshal(b, m, deterministic)
}
func (dst *KernelDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelDef.Merge(dst, src)
}
func (m *KernelDef) XXX_Size() int {
	return xxx_messageInfo_KernelDef.Size(m)
}
func (m *KernelDef) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelDef.DiscardUnknown(m)
}

var xxx_messageInfo_KernelDef proto.InternalMessageInfo

func (m *KernelDef) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *KernelDef) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *KernelDef) GetConstraint() []*KernelDef_AttrConstraint {
	if m != nil {
		return m.Constraint
	}
	return nil
}

func (m *KernelDef) GetHostMemoryArg() []string {
	if m != nil {
		return m.HostMemoryArg
	}
	return nil
}

func (m *KernelDef) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *KernelDef) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

type KernelDef_AttrConstraint struct {
	// Name of an attr from the Op.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of values that this kernel supports for this attr.
	// Like OpDef.AttrDef.allowed_values, except for kernels instead of Ops.
	AllowedValues        *AttrValue `protobuf:"bytes,2,opt,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *KernelDef_AttrConstraint) Reset()         { *m = KernelDef_AttrConstraint{} }
func (m *KernelDef_AttrConstraint) String() string { return proto.CompactTextString(m) }
func (*KernelDef_AttrConstraint) ProtoMessage()    {}
func (*KernelDef_AttrConstraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_kernel_def_a3c41f0c94cd5ffa, []int{0, 0}
}
func (m *KernelDef_AttrConstraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelDef_AttrConstraint.Unmarshal(m, b)
}
func (m *KernelDef_AttrConstraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelDef_AttrConstraint.Marshal(b, m, deterministic)
}
func (dst *KernelDef_AttrConstraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelDef_AttrConstraint.Merge(dst, src)
}
func (m *KernelDef_AttrConstraint) XXX_Size() int {
	return xxx_messageInfo_KernelDef_AttrConstraint.Size(m)
}
func (m *KernelDef_AttrConstraint) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelDef_AttrConstraint.DiscardUnknown(m)
}

var xxx_messageInfo_KernelDef_AttrConstraint proto.InternalMessageInfo

func (m *KernelDef_AttrConstraint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KernelDef_AttrConstraint) GetAllowedValues() *AttrValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

// A collection of KernelDefs
type KernelList struct {
	Kernel               []*KernelDef `protobuf:"bytes,1,rep,name=kernel,proto3" json:"kernel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *KernelList) Reset()         { *m = KernelList{} }
func (m *KernelList) String() string { return proto.CompactTextString(m) }
func (*KernelList) ProtoMessage()    {}
func (*KernelList) Descriptor() ([]byte, []int) {
	return fileDescriptor_kernel_def_a3c41f0c94cd5ffa, []int{1}
}
func (m *KernelList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelList.Unmarshal(m, b)
}
func (m *KernelList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelList.Marshal(b, m, deterministic)
}
func (dst *KernelList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelList.Merge(dst, src)
}
func (m *KernelList) XXX_Size() int {
	return xxx_messageInfo_KernelList.Size(m)
}
func (m *KernelList) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelList.DiscardUnknown(m)
}

var xxx_messageInfo_KernelList proto.InternalMessageInfo

func (m *KernelList) GetKernel() []*KernelDef {
	if m != nil {
		return m.Kernel
	}
	return nil
}

func init() {
	proto.RegisterType((*KernelDef)(nil), "tensorflow.KernelDef")
	proto.RegisterType((*KernelDef_AttrConstraint)(nil), "tensorflow.KernelDef.AttrConstraint")
	proto.RegisterType((*KernelList)(nil), "tensorflow.KernelList")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/kernel_def.proto", fileDescriptor_kernel_def_a3c41f0c94cd5ffa)
}

var fileDescriptor_kernel_def_a3c41f0c94cd5ffa = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x41, 0xab, 0xd3, 0x40,
	0x10, 0x26, 0xe9, 0x6b, 0xb1, 0x53, 0x5e, 0x1f, 0x2c, 0x3e, 0x58, 0x7a, 0x31, 0x3c, 0x44, 0x82,
	0xd0, 0x04, 0xea, 0x51, 0x2f, 0xad, 0xc5, 0x8b, 0x0a, 0x25, 0x88, 0x07, 0x2f, 0x61, 0x93, 0x4e,
	0xd2, 0xd0, 0x24, 0x13, 0x26, 0xdb, 0xd6, 0xfc, 0x3b, 0x7f, 0x96, 0x47, 0xc9, 0xa6, 0xa4, 0x51,
	0xe4, 0xdd, 0x76, 0x66, 0xbe, 0x99, 0xd9, 0xef, 0xfb, 0x06, 0xde, 0x6a, 0x2c, 0x6b, 0xe2, 0x24,
	0xa7, 0x8b, 0x1f, 0x13, 0xa3, 0x9f, 0xb0, 0x2a, 0xf0, 0x42, 0x7c, 0xf4, 0x8f, 0xc8, 0x25, 0xe6,
	0xe1, 0x1e, 0x13, 0xaf, 0x62, 0xd2, 0x24, 0xe0, 0x86, 0x5d, 0x3c, 0xd3, 0xa7, 0xb4, 0xe6, 0xf0,
	0xac, 0xf2, 0x13, 0x76, 0x7d, 0x4f, 0xbf, 0x6c, 0x98, 0x7e, 0x36, 0xc3, 0xb6, 0x98, 0x88, 0x39,
	0xd8, 0x54, 0x49, 0xcb, 0xb1, 0xdc, 0x69, 0x60, 0x53, 0x25, 0x5e, 0xc1, 0x6c, 0x8f, 0xe7, 0x2c,
	0xc6, 0x50, 0x37, 0x15, 0x4a, 0xdb, 0x14, 0xa0, 0x4b, 0x7d, 0x6b, 0x2a, 0x14, 0x5b, 0x80, 0x98,
	0xca, 0x5a, 0xb3, 0xca, 0x4a, 0x2d, 0x47, 0xce, 0xc8, 0x9d, 0xad, 0x5e, 0x7b, 0xb7, 0xfd, 0x5e,
	0x3f, 0xdb, 0x5b, 0x6b, 0xcd, 0x1f, 0x7b, 0x6c, 0x30, 0xe8, 0x13, 0x6f, 0xe0, 0xe1, 0x40, 0xb5,
	0x0e, 0x0b, 0x2c, 0x88, 0x9b, 0x50, 0x71, 0x2a, 0xef, 0x9c, 0x91, 0x3b, 0x0d, 0xee, 0xdb, 0xf4,
	0x57, 0x93, 0x5d, 0x73, 0x2a, 0x5e, 0xc2, 0x38, 0x57, 0x11, 0xe6, 0x72, 0x6c, 0x3e, 0xd2, 0x05,
	0x62, 0x01, 0x2f, 0x2a, 0xce, 0x88, 0x33, 0xdd, 0xc8, 0x89, 0x63, 0xb9, 0xe3, 0xa0, 0x8f, 0x17,
	0x11, 0xcc, 0xff, 0xde, 0x2b, 0x04, 0xdc, 0x95, 0xaa, 0xc0, 0x2b, 0x49, 0xf3, 0x16, 0x1f, 0x60,
	0xae, 0xf2, 0x9c, 0x2e, 0xb8, 0xef, 0xb4, 0xa9, 0x0d, 0xd3, 0xd9, 0xea, 0x71, 0xc8, 0xa4, 0x9d,
	0xf3, 0xbd, 0xad, 0x06, 0xf7, 0x57, 0xb0, 0x89, 0xea, 0xa7, 0xf7, 0x00, 0x1d, 0xcb, 0x2f, 0x59,
	0xad, 0xc5, 0x12, 0x26, 0x9d, 0x39, 0xd2, 0x32, 0x6a, 0x3c, 0xfe, 0x57, 0x8d, 0xe0, 0x0a, 0xda,
	0xfc, 0x04, 0x49, 0x9c, 0x0e, 0x31, 0xbd, 0x59, 0x9b, 0x87, 0x1e, 0xbe, 0x6b, 0xbd, 0xaa, 0x77,
	0xd6, 0x8f, 0x4f, 0x69, 0xa6, 0x0f, 0xa7, 0xc8, 0x8b, 0xa9, 0xf0, 0x4b, 0xd4, 0x51, 0xcb, 0xc9,
	0xd7, 0xc9, 0x32, 0xe5, 0x2a, 0xf6, 0x53, 0xf2, 0x07, 0xce, 0x0f, 0x9e, 0x29, 0xfd, 0x73, 0x07,
	0xbf, 0x2d, 0x2b, 0x9a, 0x98, 0x03, 0x78, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0xab, 0x5f, 0x29,
	0x0a, 0x66, 0x02, 0x00, 0x00,
}
