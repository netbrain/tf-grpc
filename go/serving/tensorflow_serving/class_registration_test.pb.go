// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/util/class_registration_test.proto

package tensorflow_serving

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Config1 struct {
	StringField          string   `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config1) Reset()         { *m = Config1{} }
func (m *Config1) String() string { return proto.CompactTextString(m) }
func (*Config1) ProtoMessage()    {}
func (*Config1) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a1690d17e90f318, []int{0}
}

func (m *Config1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config1.Unmarshal(m, b)
}
func (m *Config1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config1.Marshal(b, m, deterministic)
}
func (m *Config1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config1.Merge(m, src)
}
func (m *Config1) XXX_Size() int {
	return xxx_messageInfo_Config1.Size(m)
}
func (m *Config1) XXX_DiscardUnknown() {
	xxx_messageInfo_Config1.DiscardUnknown(m)
}

var xxx_messageInfo_Config1 proto.InternalMessageInfo

func (m *Config1) GetStringField() string {
	if m != nil {
		return m.StringField
	}
	return ""
}

type Config2 struct {
	StringField          string   `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config2) Reset()         { *m = Config2{} }
func (m *Config2) String() string { return proto.CompactTextString(m) }
func (*Config2) ProtoMessage()    {}
func (*Config2) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a1690d17e90f318, []int{1}
}

func (m *Config2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config2.Unmarshal(m, b)
}
func (m *Config2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config2.Marshal(b, m, deterministic)
}
func (m *Config2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config2.Merge(m, src)
}
func (m *Config2) XXX_Size() int {
	return xxx_messageInfo_Config2.Size(m)
}
func (m *Config2) XXX_DiscardUnknown() {
	xxx_messageInfo_Config2.DiscardUnknown(m)
}

var xxx_messageInfo_Config2 proto.InternalMessageInfo

func (m *Config2) GetStringField() string {
	if m != nil {
		return m.StringField
	}
	return ""
}

type MessageWithAny struct {
	AnyField             *any.Any `protobuf:"bytes,1,opt,name=any_field,json=anyField,proto3" json:"any_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageWithAny) Reset()         { *m = MessageWithAny{} }
func (m *MessageWithAny) String() string { return proto.CompactTextString(m) }
func (*MessageWithAny) ProtoMessage()    {}
func (*MessageWithAny) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a1690d17e90f318, []int{2}
}

func (m *MessageWithAny) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithAny.Unmarshal(m, b)
}
func (m *MessageWithAny) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithAny.Marshal(b, m, deterministic)
}
func (m *MessageWithAny) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithAny.Merge(m, src)
}
func (m *MessageWithAny) XXX_Size() int {
	return xxx_messageInfo_MessageWithAny.Size(m)
}
func (m *MessageWithAny) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithAny.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithAny proto.InternalMessageInfo

func (m *MessageWithAny) GetAnyField() *any.Any {
	if m != nil {
		return m.AnyField
	}
	return nil
}

func init() {
	proto.RegisterType((*Config1)(nil), "tensorflow.serving.Config1")
	proto.RegisterType((*Config2)(nil), "tensorflow.serving.Config2")
	proto.RegisterType((*MessageWithAny)(nil), "tensorflow.serving.MessageWithAny")
}

func init() {
	proto.RegisterFile("tensorflow_serving/util/class_registration_test.proto", fileDescriptor_2a1690d17e90f318)
}

var fileDescriptor_2a1690d17e90f318 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0x03, 0x41,
	0x10, 0x85, 0x49, 0xa3, 0x66, 0x23, 0x16, 0x87, 0x85, 0x5a, 0xa9, 0x95, 0x85, 0xee, 0x92, 0x88,
	0x85, 0x58, 0xc5, 0x80, 0x9d, 0x4d, 0x1a, 0xc1, 0xe6, 0xd8, 0x3b, 0x77, 0x27, 0x03, 0xeb, 0x4c,
	0xd8, 0x99, 0x53, 0xf6, 0xdf, 0x8b, 0x77, 0x86, 0x13, 0x6c, 0x6c, 0x1f, 0xdf, 0xf7, 0x78, 0x3c,
	0x73, 0xa7, 0x81, 0x84, 0x73, 0x4c, 0xfc, 0x59, 0x4b, 0xc8, 0x1f, 0x48, 0xe0, 0x3a, 0xc5, 0xe4,
	0xda, 0xe4, 0x45, 0xea, 0x1c, 0x00, 0x45, 0xb3, 0x57, 0x64, 0xaa, 0x35, 0x88, 0xda, 0x6d, 0x66,
	0xe5, 0xaa, 0x1a, 0x35, 0xfb, 0xa3, 0x9d, 0x9d, 0x02, 0x33, 0xa4, 0xe0, 0x7a, 0xa2, 0xe9, 0xa2,
	0xf3, 0x54, 0x06, 0xfc, 0xf2, 0xda, 0xec, 0xaf, 0x98, 0x22, 0xc2, 0xbc, 0xba, 0x30, 0x87, 0xa2,
	0x19, 0x09, 0xea, 0x88, 0x21, 0xbd, 0x9d, 0x4c, 0xce, 0x27, 0x57, 0xd3, 0xf5, 0x6c, 0xc8, 0x9e,
	0xbe, 0xa3, 0x91, 0x5e, 0xfc, 0x87, 0x5e, 0x99, 0xa3, 0xe7, 0x20, 0xe2, 0x21, 0xbc, 0xa0, 0x6e,
	0x96, 0x54, 0xaa, 0xb9, 0x99, 0x7a, 0x2a, 0xbf, 0x8c, 0xd9, 0xe2, 0xd8, 0x0e, 0xe3, 0xec, 0x6e,
	0x9c, 0x5d, 0x52, 0x59, 0x1f, 0x78, 0x2a, 0x7d, 0xc9, 0xe3, 0xc3, 0xeb, 0x3d, 0xa0, 0x6e, 0xba,
	0xc6, 0xb6, 0xfc, 0xee, 0x28, 0x68, 0x93, 0x3d, 0x92, 0xd3, 0x78, 0x03, 0x79, 0xdb, 0x3a, 0x60,
	0xb7, 0x3b, 0xe7, 0xef, 0x5f, 0xcd, 0x5e, 0x5f, 0x7a, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0xff,
	0x2e, 0xf2, 0xc8, 0x4c, 0x01, 0x00, 0x00,
}