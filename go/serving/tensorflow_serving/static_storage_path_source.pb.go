// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/sources/storage_path/static_storage_path_source.proto

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

// Config proto for StaticStoragePathSource.
type StaticStoragePathSourceConfig struct {
	// The single servable name, version number and path to supply statically.
	ServableName         string   `protobuf:"bytes,1,opt,name=servable_name,json=servableName,proto3" json:"servable_name,omitempty"`
	VersionNum           int64    `protobuf:"varint,2,opt,name=version_num,json=versionNum,proto3" json:"version_num,omitempty"`
	VersionPath          string   `protobuf:"bytes,3,opt,name=version_path,json=versionPath,proto3" json:"version_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticStoragePathSourceConfig) Reset()         { *m = StaticStoragePathSourceConfig{} }
func (m *StaticStoragePathSourceConfig) String() string { return proto.CompactTextString(m) }
func (*StaticStoragePathSourceConfig) ProtoMessage()    {}
func (*StaticStoragePathSourceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_static_storage_path_source_9e714142e18146c1, []int{0}
}
func (m *StaticStoragePathSourceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticStoragePathSourceConfig.Unmarshal(m, b)
}
func (m *StaticStoragePathSourceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticStoragePathSourceConfig.Marshal(b, m, deterministic)
}
func (dst *StaticStoragePathSourceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticStoragePathSourceConfig.Merge(dst, src)
}
func (m *StaticStoragePathSourceConfig) XXX_Size() int {
	return xxx_messageInfo_StaticStoragePathSourceConfig.Size(m)
}
func (m *StaticStoragePathSourceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticStoragePathSourceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StaticStoragePathSourceConfig proto.InternalMessageInfo

func (m *StaticStoragePathSourceConfig) GetServableName() string {
	if m != nil {
		return m.ServableName
	}
	return ""
}

func (m *StaticStoragePathSourceConfig) GetVersionNum() int64 {
	if m != nil {
		return m.VersionNum
	}
	return 0
}

func (m *StaticStoragePathSourceConfig) GetVersionPath() string {
	if m != nil {
		return m.VersionPath
	}
	return ""
}

func init() {
	proto.RegisterType((*StaticStoragePathSourceConfig)(nil), "tensorflow.serving.StaticStoragePathSourceConfig")
}

func init() {
	proto.RegisterFile("tensorflow_serving/sources/storage_path/static_storage_path_source.proto", fileDescriptor_static_storage_path_source_9e714142e18146c1)
}

var fileDescriptor_static_storage_path_source_9e714142e18146c1 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x2a, 0x21, 0x61, 0xca, 0xe2, 0x29, 0x0b, 0xa2, 0xc0, 0xd2, 0x85, 0x78, 0x60,
	0x42, 0x6c, 0xb0, 0x30, 0x55, 0xa8, 0xd9, 0x58, 0x2c, 0x27, 0xba, 0x38, 0x96, 0xea, 0xbb, 0xe8,
	0x7c, 0x29, 0x7f, 0x81, 0x9f, 0x8d, 0x92, 0xc6, 0x02, 0xa9, 0xeb, 0xd3, 0xfb, 0x3e, 0xfb, 0x9d,
	0xfa, 0x10, 0xc0, 0x44, 0xdc, 0x1d, 0xe8, 0xdb, 0x26, 0xe0, 0x63, 0x40, 0x6f, 0x12, 0x8d, 0xdc,
	0x42, 0x32, 0x49, 0x88, 0x9d, 0x07, 0x3b, 0x38, 0xe9, 0x4d, 0x12, 0x27, 0xa1, 0xb5, 0xff, 0x33,
	0x7b, 0x2a, 0x56, 0x03, 0x93, 0x90, 0xd6, 0x7f, 0xa6, 0x6a, 0x31, 0x3d, 0xfc, 0x14, 0xea, 0xb6,
	0x9e, 0xc1, 0xfa, 0xc4, 0x7d, 0x3a, 0xe9, 0xeb, 0x99, 0x7a, 0x27, 0xec, 0x82, 0xd7, 0x8f, 0xea,
	0x66, 0x2a, 0xbb, 0xe6, 0x00, 0x16, 0x5d, 0x84, 0xb2, 0xd8, 0x14, 0xdb, 0xab, 0xfd, 0x3a, 0x87,
	0x3b, 0x17, 0x41, 0xdf, 0xa9, 0xeb, 0x23, 0x70, 0x0a, 0x84, 0x16, 0xc7, 0x58, 0x5e, 0x6c, 0x8a,
	0xed, 0x6a, 0xaf, 0x96, 0x68, 0x37, 0x46, 0x7d, 0xaf, 0xd6, 0xb9, 0x30, 0x7d, 0xac, 0x5c, 0xcd,
	0x92, 0x0c, 0x4d, 0x8f, 0xbe, 0xbd, 0x7e, 0xbd, 0xf8, 0x20, 0xfd, 0xd8, 0x54, 0x2d, 0x45, 0x83,
	0x20, 0x0d, 0xbb, 0x80, 0x46, 0xba, 0x27, 0xcf, 0x43, 0x6b, 0x3c, 0x99, 0x3c, 0xff, 0xfc, 0x22,
	0xcd, 0xe5, 0x3c, 0xf1, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x39, 0x6d, 0x91, 0x2e, 0x01,
	0x00, 0x00,
}
