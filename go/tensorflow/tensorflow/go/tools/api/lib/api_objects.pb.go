// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/tools/api/lib/api_objects.proto

package lib // import "github.com/netbrain/tf-grpc/go/tensorflow/tensorflow/go/tools/api/lib"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TFAPIMember struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Mtype                *string  `protobuf:"bytes,2,opt,name=mtype" json:"mtype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TFAPIMember) Reset()         { *m = TFAPIMember{} }
func (m *TFAPIMember) String() string { return proto.CompactTextString(m) }
func (*TFAPIMember) ProtoMessage()    {}
func (*TFAPIMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_objects_2bcde4a5ccb32fef, []int{0}
}
func (m *TFAPIMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIMember.Unmarshal(m, b)
}
func (m *TFAPIMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIMember.Marshal(b, m, deterministic)
}
func (dst *TFAPIMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIMember.Merge(dst, src)
}
func (m *TFAPIMember) XXX_Size() int {
	return xxx_messageInfo_TFAPIMember.Size(m)
}
func (m *TFAPIMember) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIMember.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIMember proto.InternalMessageInfo

func (m *TFAPIMember) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TFAPIMember) GetMtype() string {
	if m != nil && m.Mtype != nil {
		return *m.Mtype
	}
	return ""
}

type TFAPIMethod struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Path                 *string  `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Argspec              *string  `protobuf:"bytes,3,opt,name=argspec" json:"argspec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TFAPIMethod) Reset()         { *m = TFAPIMethod{} }
func (m *TFAPIMethod) String() string { return proto.CompactTextString(m) }
func (*TFAPIMethod) ProtoMessage()    {}
func (*TFAPIMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_objects_2bcde4a5ccb32fef, []int{1}
}
func (m *TFAPIMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIMethod.Unmarshal(m, b)
}
func (m *TFAPIMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIMethod.Marshal(b, m, deterministic)
}
func (dst *TFAPIMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIMethod.Merge(dst, src)
}
func (m *TFAPIMethod) XXX_Size() int {
	return xxx_messageInfo_TFAPIMethod.Size(m)
}
func (m *TFAPIMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIMethod.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIMethod proto.InternalMessageInfo

func (m *TFAPIMethod) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TFAPIMethod) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TFAPIMethod) GetArgspec() string {
	if m != nil && m.Argspec != nil {
		return *m.Argspec
	}
	return ""
}

type TFAPIModule struct {
	Member               []*TFAPIMember `protobuf:"bytes,1,rep,name=member" json:"member,omitempty"`
	MemberMethod         []*TFAPIMethod `protobuf:"bytes,2,rep,name=member_method,json=memberMethod" json:"member_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TFAPIModule) Reset()         { *m = TFAPIModule{} }
func (m *TFAPIModule) String() string { return proto.CompactTextString(m) }
func (*TFAPIModule) ProtoMessage()    {}
func (*TFAPIModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_objects_2bcde4a5ccb32fef, []int{2}
}
func (m *TFAPIModule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIModule.Unmarshal(m, b)
}
func (m *TFAPIModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIModule.Marshal(b, m, deterministic)
}
func (dst *TFAPIModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIModule.Merge(dst, src)
}
func (m *TFAPIModule) XXX_Size() int {
	return xxx_messageInfo_TFAPIModule.Size(m)
}
func (m *TFAPIModule) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIModule.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIModule proto.InternalMessageInfo

func (m *TFAPIModule) GetMember() []*TFAPIMember {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *TFAPIModule) GetMemberMethod() []*TFAPIMethod {
	if m != nil {
		return m.MemberMethod
	}
	return nil
}

type TFAPIClass struct {
	IsInstance           []string       `protobuf:"bytes,1,rep,name=is_instance,json=isInstance" json:"is_instance,omitempty"`
	Member               []*TFAPIMember `protobuf:"bytes,2,rep,name=member" json:"member,omitempty"`
	MemberMethod         []*TFAPIMethod `protobuf:"bytes,3,rep,name=member_method,json=memberMethod" json:"member_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TFAPIClass) Reset()         { *m = TFAPIClass{} }
func (m *TFAPIClass) String() string { return proto.CompactTextString(m) }
func (*TFAPIClass) ProtoMessage()    {}
func (*TFAPIClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_objects_2bcde4a5ccb32fef, []int{3}
}
func (m *TFAPIClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIClass.Unmarshal(m, b)
}
func (m *TFAPIClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIClass.Marshal(b, m, deterministic)
}
func (dst *TFAPIClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIClass.Merge(dst, src)
}
func (m *TFAPIClass) XXX_Size() int {
	return xxx_messageInfo_TFAPIClass.Size(m)
}
func (m *TFAPIClass) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIClass.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIClass proto.InternalMessageInfo

func (m *TFAPIClass) GetIsInstance() []string {
	if m != nil {
		return m.IsInstance
	}
	return nil
}

func (m *TFAPIClass) GetMember() []*TFAPIMember {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *TFAPIClass) GetMemberMethod() []*TFAPIMethod {
	if m != nil {
		return m.MemberMethod
	}
	return nil
}

type TFAPIProto struct {
	Descriptor_          *descriptor.DescriptorProto `protobuf:"bytes,1,opt,name=descriptor" json:"descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TFAPIProto) Reset()         { *m = TFAPIProto{} }
func (m *TFAPIProto) String() string { return proto.CompactTextString(m) }
func (*TFAPIProto) ProtoMessage()    {}
func (*TFAPIProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_objects_2bcde4a5ccb32fef, []int{4}
}
func (m *TFAPIProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIProto.Unmarshal(m, b)
}
func (m *TFAPIProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIProto.Marshal(b, m, deterministic)
}
func (dst *TFAPIProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIProto.Merge(dst, src)
}
func (m *TFAPIProto) XXX_Size() int {
	return xxx_messageInfo_TFAPIProto.Size(m)
}
func (m *TFAPIProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIProto.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIProto proto.InternalMessageInfo

func (m *TFAPIProto) GetDescriptor_() *descriptor.DescriptorProto {
	if m != nil {
		return m.Descriptor_
	}
	return nil
}

type TFAPIObject struct {
	Path                 *string      `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	TfModule             *TFAPIModule `protobuf:"bytes,2,opt,name=tf_module,json=tfModule" json:"tf_module,omitempty"`
	TfClass              *TFAPIClass  `protobuf:"bytes,3,opt,name=tf_class,json=tfClass" json:"tf_class,omitempty"`
	TfProto              *TFAPIProto  `protobuf:"bytes,4,opt,name=tf_proto,json=tfProto" json:"tf_proto,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TFAPIObject) Reset()         { *m = TFAPIObject{} }
func (m *TFAPIObject) String() string { return proto.CompactTextString(m) }
func (*TFAPIObject) ProtoMessage()    {}
func (*TFAPIObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_objects_2bcde4a5ccb32fef, []int{5}
}
func (m *TFAPIObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIObject.Unmarshal(m, b)
}
func (m *TFAPIObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIObject.Marshal(b, m, deterministic)
}
func (dst *TFAPIObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIObject.Merge(dst, src)
}
func (m *TFAPIObject) XXX_Size() int {
	return xxx_messageInfo_TFAPIObject.Size(m)
}
func (m *TFAPIObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIObject.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIObject proto.InternalMessageInfo

func (m *TFAPIObject) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TFAPIObject) GetTfModule() *TFAPIModule {
	if m != nil {
		return m.TfModule
	}
	return nil
}

func (m *TFAPIObject) GetTfClass() *TFAPIClass {
	if m != nil {
		return m.TfClass
	}
	return nil
}

func (m *TFAPIObject) GetTfProto() *TFAPIProto {
	if m != nil {
		return m.TfProto
	}
	return nil
}

func init() {
	proto.RegisterType((*TFAPIMember)(nil), "third_party.tensorflow.tools.api.TFAPIMember")
	proto.RegisterType((*TFAPIMethod)(nil), "third_party.tensorflow.tools.api.TFAPIMethod")
	proto.RegisterType((*TFAPIModule)(nil), "third_party.tensorflow.tools.api.TFAPIModule")
	proto.RegisterType((*TFAPIClass)(nil), "third_party.tensorflow.tools.api.TFAPIClass")
	proto.RegisterType((*TFAPIProto)(nil), "third_party.tensorflow.tools.api.TFAPIProto")
	proto.RegisterType((*TFAPIObject)(nil), "third_party.tensorflow.tools.api.TFAPIObject")
}

func init() {
	proto.RegisterFile("tensorflow/tools/api/lib/api_objects.proto", fileDescriptor_api_objects_2bcde4a5ccb32fef)
}

var fileDescriptor_api_objects_2bcde4a5ccb32fef = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcf, 0x8b, 0xd4, 0x30,
	0x14, 0xa6, 0x33, 0xa3, 0xb3, 0x93, 0x2a, 0x48, 0xf0, 0x50, 0xbc, 0x58, 0x7a, 0x5a, 0xc4, 0x4d,
	0x60, 0x2e, 0x82, 0x27, 0x7f, 0xad, 0xcb, 0x0a, 0xb2, 0x4b, 0xf0, 0xe4, 0xa5, 0xa4, 0x9d, 0xa4,
	0x13, 0x69, 0x9b, 0x90, 0xbc, 0x41, 0xf6, 0x3f, 0xf2, 0x4f, 0xf1, 0xbf, 0x52, 0xe6, 0xa5, 0x9d,
	0xe9, 0x41, 0xd0, 0xc1, 0x3d, 0xf5, 0xbd, 0x97, 0x7c, 0x1f, 0xef, 0xeb, 0xf7, 0x85, 0xbc, 0x00,
	0xd5, 0x07, 0xeb, 0x75, 0x6b, 0xbf, 0x73, 0xb0, 0xb6, 0x0d, 0x5c, 0x3a, 0xc3, 0x5b, 0x53, 0xed,
	0xbf, 0xa5, 0xad, 0xbe, 0xa9, 0x1a, 0x02, 0x73, 0xde, 0x82, 0xa5, 0x39, 0x6c, 0x8d, 0xdf, 0x94,
	0x4e, 0x7a, 0xb8, 0x63, 0x47, 0x1c, 0x43, 0x1c, 0x93, 0xce, 0x3c, 0xcb, 0x1b, 0x6b, 0x9b, 0x56,
	0x71, 0xbc, 0x5f, 0xed, 0x34, 0xdf, 0xa8, 0x50, 0x7b, 0xe3, 0xc0, 0xfa, 0xc8, 0x51, 0xbc, 0x22,
	0xe9, 0x97, 0x8f, 0x6f, 0x6f, 0xaf, 0x3f, 0xab, 0xae, 0x52, 0x9e, 0x52, 0xb2, 0xe8, 0x65, 0xa7,
	0xb2, 0x24, 0x4f, 0xce, 0x57, 0x02, 0x6b, 0xfa, 0x94, 0x3c, 0xe8, 0xe0, 0xce, 0xa9, 0x6c, 0x86,
	0xc3, 0xd8, 0x14, 0x37, 0x07, 0x20, 0x6c, 0xed, 0xe6, 0x8f, 0x40, 0x4a, 0x16, 0x4e, 0xc2, 0x76,
	0xc0, 0x61, 0x4d, 0x33, 0xb2, 0x94, 0xbe, 0x09, 0x4e, 0xd5, 0xd9, 0x1c, 0xc7, 0x63, 0x5b, 0xfc,
	0x48, 0x46, 0x46, 0xbb, 0xd9, 0xb5, 0x8a, 0x5e, 0x92, 0x87, 0x1d, 0x2e, 0x95, 0x25, 0xf9, 0xfc,
	0x3c, 0x5d, 0x5f, 0xb0, 0xbf, 0xc9, 0x65, 0x13, 0x25, 0x62, 0x00, 0x53, 0x41, 0x1e, 0xc7, 0xaa,
	0xec, 0x70, 0xd3, 0x6c, 0x76, 0x22, 0xdb, 0x1e, 0x24, 0x1e, 0x45, 0x8e, 0xd8, 0x15, 0x3f, 0x13,
	0x42, 0xf0, 0xf4, 0x7d, 0x2b, 0x43, 0xa0, 0xcf, 0x49, 0x6a, 0x42, 0x69, 0xfa, 0x00, 0xb2, 0xaf,
	0x15, 0xae, 0xbb, 0x12, 0xc4, 0x84, 0xeb, 0x61, 0x32, 0x91, 0x32, 0xbb, 0x57, 0x29, 0xf3, 0xff,
	0x97, 0x22, 0x06, 0x25, 0xb7, 0x98, 0xa8, 0x37, 0x84, 0x1c, 0x13, 0x82, 0x5e, 0xa6, 0xeb, 0x9c,
	0xc5, 0x10, 0xb1, 0x31, 0x44, 0xec, 0xc3, 0xe1, 0x0a, 0xa2, 0xc4, 0x04, 0xf3, 0x7a, 0xf6, 0x24,
	0x29, 0x7e, 0x8d, 0x4e, 0xde, 0x60, 0x5c, 0x0f, 0x39, 0x48, 0x26, 0x39, 0xf8, 0x44, 0x56, 0xa0,
	0xcb, 0x0e, 0xad, 0xc6, 0x80, 0x9c, 0xa0, 0x03, 0x41, 0xe2, 0x0c, 0xf4, 0x90, 0x94, 0x2b, 0x72,
	0x06, 0xba, 0xac, 0xf7, 0x5e, 0x60, 0xa8, 0xd2, 0xf5, 0xcb, 0x7f, 0xa4, 0x42, 0xff, 0xc4, 0x12,
	0x74, 0x34, 0x32, 0x12, 0xa1, 0xce, 0x6c, 0x71, 0x12, 0x51, 0xfc, 0x11, 0x4b, 0xd0, 0x58, 0xbc,
	0xbb, 0xfa, 0x7a, 0xd9, 0x18, 0xd8, 0xee, 0x2a, 0x56, 0xdb, 0x8e, 0xf7, 0x0a, 0x2a, 0x2f, 0x4d,
	0xcf, 0x41, 0x5f, 0x34, 0xde, 0xd5, 0xbc, 0xb1, 0x7c, 0xfa, 0xcc, 0x8f, 0xe5, 0xfe, 0x60, 0xfa,
	0xe8, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xeb, 0x14, 0x53, 0x0f, 0x04, 0x00, 0x00,
}
