// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/servables/tensorflow/saved_model_bundle_source_adapter.proto

package tensorflow_serving // import "github.com/netbrain/tf-grpc/go/serving/tensorflow_serving"

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

// Config proto for SavedModelBundleSourceAdapter.
type SavedModelBundleSourceAdapterConfig struct {
	// A SessionBundleConfig.
	// FOR INTERNAL USE ONLY DURING TRANSITION TO SAVED_MODEL. WILL BE DEPRECATED.
	// TODO(b/32248363): Replace this field with the "real" field(s).
	LegacyConfig         *SessionBundleConfig `protobuf:"bytes,1000,opt,name=legacy_config,json=legacyConfig,proto3" json:"legacy_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SavedModelBundleSourceAdapterConfig) Reset()         { *m = SavedModelBundleSourceAdapterConfig{} }
func (m *SavedModelBundleSourceAdapterConfig) String() string { return proto.CompactTextString(m) }
func (*SavedModelBundleSourceAdapterConfig) ProtoMessage()    {}
func (*SavedModelBundleSourceAdapterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_saved_model_bundle_source_adapter_81d831d8b3665e1d, []int{0}
}
func (m *SavedModelBundleSourceAdapterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedModelBundleSourceAdapterConfig.Unmarshal(m, b)
}
func (m *SavedModelBundleSourceAdapterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedModelBundleSourceAdapterConfig.Marshal(b, m, deterministic)
}
func (dst *SavedModelBundleSourceAdapterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedModelBundleSourceAdapterConfig.Merge(dst, src)
}
func (m *SavedModelBundleSourceAdapterConfig) XXX_Size() int {
	return xxx_messageInfo_SavedModelBundleSourceAdapterConfig.Size(m)
}
func (m *SavedModelBundleSourceAdapterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedModelBundleSourceAdapterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SavedModelBundleSourceAdapterConfig proto.InternalMessageInfo

func (m *SavedModelBundleSourceAdapterConfig) GetLegacyConfig() *SessionBundleConfig {
	if m != nil {
		return m.LegacyConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*SavedModelBundleSourceAdapterConfig)(nil), "tensorflow.serving.SavedModelBundleSourceAdapterConfig")
}

func init() {
	proto.RegisterFile("tensorflow_serving/servables/tensorflow/saved_model_bundle_source_adapter.proto", fileDescriptor_saved_model_bundle_source_adapter_81d831d8b3665e1d)
}

var fileDescriptor_saved_model_bundle_source_adapter_81d831d8b3665e1d = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0x3d, 0x4b, 0xc5, 0x30,
	0x14, 0x86, 0x71, 0x51, 0xa8, 0xba, 0x74, 0x12, 0x27, 0xd1, 0x41, 0x17, 0x13, 0xd0, 0x49, 0x9c,
	0x6c, 0xe7, 0x22, 0xd8, 0xcd, 0x25, 0xe4, 0xe3, 0x34, 0x06, 0xd2, 0x9c, 0x92, 0xa4, 0x95, 0xfb,
	0x8f, 0xef, 0xcf, 0xb8, 0x34, 0x69, 0xe9, 0xd0, 0xe5, 0x4e, 0x21, 0x9c, 0xf3, 0xbe, 0xcf, 0x73,
	0x8a, 0xef, 0x08, 0x2e, 0xa0, 0xef, 0x2c, 0xfe, 0xb3, 0x00, 0x7e, 0x32, 0x4e, 0xd3, 0xf9, 0xe5,
	0xc2, 0x42, 0xa0, 0xdb, 0x90, 0x06, 0x3e, 0x81, 0x62, 0x3d, 0x2a, 0xb0, 0x4c, 0x8c, 0x4e, 0x59,
	0x60, 0x01, 0x47, 0x2f, 0x81, 0x71, 0xc5, 0x87, 0x08, 0x9e, 0x0c, 0x1e, 0x23, 0x96, 0xe5, 0x96,
	0x21, 0x4b, 0xe1, 0x7d, 0x7d, 0x36, 0x04, 0x42, 0x30, 0xe8, 0x56, 0x80, 0x44, 0xd7, 0x19, 0x9d,
	0x8b, 0x1f, 0x63, 0xf1, 0xd4, 0xce, 0x0e, 0xcd, 0xac, 0x50, 0xa5, 0x85, 0x36, 0x09, 0x7c, 0x65,
	0x7e, 0x9d, 0x96, 0xcb, 0xa6, 0xb8, 0xb5, 0xa0, 0xb9, 0x3c, 0x2c, 0xe9, 0xbb, 0xe3, 0xd5, 0xc3,
	0xc5, 0xcb, 0xf5, 0xdb, 0x33, 0xd9, 0x8b, 0x91, 0x36, 0xf3, 0x72, 0x5b, 0x2e, 0xf8, 0xb9, 0xc9,
	0xf1, 0xfc, 0xab, 0x3e, 0x7f, 0x3f, 0xb4, 0x89, 0x7f, 0xa3, 0x20, 0x12, 0x7b, 0xea, 0x20, 0x0a,
	0xcf, 0x8d, 0xa3, 0xb1, 0x7b, 0xd5, 0x7e, 0x90, 0x54, 0x23, 0x5d, 0x0f, 0xda, 0xdf, 0x28, 0x2e,
	0x93, 0xf9, 0xfb, 0x29, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xa4, 0x47, 0x1f, 0x65, 0x01, 0x00, 0x00,
}
