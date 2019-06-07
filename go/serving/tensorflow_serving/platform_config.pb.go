// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/config/platform_config.proto

package tensorflow_serving // import "github.com/netbrain/tf-grpc/go/serving/tensorflow_serving"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Configuration for a servable platform e.g. tensorflow or other ML systems.
type PlatformConfig struct {
	// The config proto for a SourceAdapter in the StoragePathSourceAdapter
	// registry.
	SourceAdapterConfig  *any.Any `protobuf:"bytes,1,opt,name=source_adapter_config,json=sourceAdapterConfig,proto3" json:"source_adapter_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformConfig) Reset()         { *m = PlatformConfig{} }
func (m *PlatformConfig) String() string { return proto.CompactTextString(m) }
func (*PlatformConfig) ProtoMessage()    {}
func (*PlatformConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_config_9b1a15e6c4d97e98, []int{0}
}
func (m *PlatformConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformConfig.Unmarshal(m, b)
}
func (m *PlatformConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformConfig.Marshal(b, m, deterministic)
}
func (dst *PlatformConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformConfig.Merge(dst, src)
}
func (m *PlatformConfig) XXX_Size() int {
	return xxx_messageInfo_PlatformConfig.Size(m)
}
func (m *PlatformConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformConfig proto.InternalMessageInfo

func (m *PlatformConfig) GetSourceAdapterConfig() *any.Any {
	if m != nil {
		return m.SourceAdapterConfig
	}
	return nil
}

type PlatformConfigMap struct {
	// A map from a platform name to a platform config. The platform name is used
	// in ModelConfig.model_platform.
	PlatformConfigs      map[string]*PlatformConfig `protobuf:"bytes,1,rep,name=platform_configs,json=platformConfigs,proto3" json:"platform_configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PlatformConfigMap) Reset()         { *m = PlatformConfigMap{} }
func (m *PlatformConfigMap) String() string { return proto.CompactTextString(m) }
func (*PlatformConfigMap) ProtoMessage()    {}
func (*PlatformConfigMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_config_9b1a15e6c4d97e98, []int{1}
}
func (m *PlatformConfigMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformConfigMap.Unmarshal(m, b)
}
func (m *PlatformConfigMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformConfigMap.Marshal(b, m, deterministic)
}
func (dst *PlatformConfigMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformConfigMap.Merge(dst, src)
}
func (m *PlatformConfigMap) XXX_Size() int {
	return xxx_messageInfo_PlatformConfigMap.Size(m)
}
func (m *PlatformConfigMap) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformConfigMap.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformConfigMap proto.InternalMessageInfo

func (m *PlatformConfigMap) GetPlatformConfigs() map[string]*PlatformConfig {
	if m != nil {
		return m.PlatformConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*PlatformConfig)(nil), "tensorflow.serving.PlatformConfig")
	proto.RegisterType((*PlatformConfigMap)(nil), "tensorflow.serving.PlatformConfigMap")
	proto.RegisterMapType((map[string]*PlatformConfig)(nil), "tensorflow.serving.PlatformConfigMap.PlatformConfigsEntry")
}

func init() {
	proto.RegisterFile("tensorflow_serving/config/platform_config.proto", fileDescriptor_platform_config_9b1a15e6c4d97e98)
}

var fileDescriptor_platform_config_9b1a15e6c4d97e98 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0x86, 0x82, 0x19, 0xe8, 0x8c, 0x13, 0xe6, 0x4e, 0xa3, 0xa7, 0x5d, 0x4c, 0x60,
	0x5e, 0xe6, 0x0e, 0xc2, 0x14, 0xc1, 0x8b, 0x20, 0x3d, 0xee, 0x52, 0xd2, 0x9a, 0xc4, 0x62, 0x97,
	0x17, 0x92, 0x74, 0xd2, 0xaf, 0xeb, 0xa7, 0xf0, 0x28, 0x26, 0x1d, 0xb2, 0x55, 0xf0, 0xd6, 0x3e,
	0xfe, 0xbf, 0x1f, 0xff, 0x97, 0x87, 0x99, 0x17, 0xda, 0x81, 0x95, 0x15, 0x7c, 0x64, 0x4e, 0xd8,
	0x6d, 0xa9, 0x15, 0x2b, 0x40, 0xcb, 0x52, 0x31, 0x53, 0x71, 0x2f, 0xc1, 0x6e, 0xb2, 0xf8, 0x4f,
	0x8d, 0x05, 0x0f, 0x84, 0xfc, 0x02, 0xb4, 0x05, 0x26, 0x57, 0x0a, 0x40, 0x55, 0x82, 0x85, 0x44,
	0x5e, 0x4b, 0xc6, 0x75, 0x13, 0xe3, 0xc9, 0x1a, 0x9f, 0xbe, 0xb4, 0x9e, 0x87, 0xa0, 0x21, 0x4f,
	0xf8, 0xd2, 0x41, 0x6d, 0x0b, 0x91, 0xf1, 0x57, 0x6e, 0xbc, 0xb0, 0xad, 0x7f, 0x8c, 0xa6, 0x68,
	0x36, 0x98, 0x8f, 0x68, 0x94, 0xd1, 0x9d, 0x8c, 0xae, 0x74, 0x93, 0x5e, 0x44, 0x64, 0x15, 0x89,
	0x68, 0x4a, 0x3e, 0x11, 0x3e, 0xdf, 0x97, 0x3f, 0x73, 0x43, 0x04, 0x1e, 0x1e, 0x34, 0x77, 0x63,
	0x34, 0xed, 0xcf, 0x06, 0xf3, 0x25, 0xed, 0x76, 0xa7, 0x1d, 0xc1, 0xc1, 0xc4, 0x3d, 0x6a, 0x6f,
	0x9b, 0xf4, 0xcc, 0xec, 0x4f, 0x27, 0x12, 0x8f, 0xfe, 0x0a, 0x92, 0x21, 0xee, 0xbf, 0x8b, 0x26,
	0x2c, 0x73, 0x92, 0xfe, 0x7c, 0x92, 0x05, 0x3e, 0xda, 0xf2, 0xaa, 0x16, 0xe3, 0x5e, 0x58, 0x30,
	0xf9, 0xbf, 0x45, 0x1a, 0x81, 0x65, 0x6f, 0x81, 0xee, 0xef, 0xd6, 0xb7, 0xaa, 0xf4, 0x6f, 0x75,
	0x4e, 0x0b, 0xd8, 0x30, 0x2d, 0x7c, 0x6e, 0x79, 0xa9, 0x99, 0x97, 0xd7, 0xca, 0x9a, 0x82, 0x29,
	0x60, 0xbb, 0xb3, 0x75, 0x2f, 0xf9, 0x85, 0x50, 0x7e, 0x1c, 0xde, 0xf1, 0xe6, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x16, 0x57, 0xb1, 0xd0, 0xe9, 0x01, 0x00, 0x00,
}
