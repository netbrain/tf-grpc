// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/versions.proto

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

// Version information for a piece of serialized data
//
// There are different types of versions for each type of data
// (GraphDef, etc.), but they all have the same common shape
// described here.
//
// Each consumer has "consumer" and "min_producer" versions (specified
// elsewhere).  A consumer is allowed to consume this data if
//
//   producer >= min_producer
//   consumer >= min_consumer
//   consumer not in bad_consumers
//
type VersionDef struct {
	// The version of the code that produced this data.
	Producer int32 `protobuf:"varint,1,opt,name=producer,proto3" json:"producer,omitempty"`
	// Any consumer below this version is not allowed to consume this data.
	MinConsumer int32 `protobuf:"varint,2,opt,name=min_consumer,json=minConsumer,proto3" json:"min_consumer,omitempty"`
	// Specific consumer versions which are disallowed (e.g. due to bugs).
	BadConsumers         []int32  `protobuf:"varint,3,rep,packed,name=bad_consumers,json=badConsumers,proto3" json:"bad_consumers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionDef) Reset()         { *m = VersionDef{} }
func (m *VersionDef) String() string { return proto.CompactTextString(m) }
func (*VersionDef) ProtoMessage()    {}
func (*VersionDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_versions_5c77a333a5601135, []int{0}
}
func (m *VersionDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionDef.Unmarshal(m, b)
}
func (m *VersionDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionDef.Marshal(b, m, deterministic)
}
func (dst *VersionDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionDef.Merge(dst, src)
}
func (m *VersionDef) XXX_Size() int {
	return xxx_messageInfo_VersionDef.Size(m)
}
func (m *VersionDef) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionDef.DiscardUnknown(m)
}

var xxx_messageInfo_VersionDef proto.InternalMessageInfo

func (m *VersionDef) GetProducer() int32 {
	if m != nil {
		return m.Producer
	}
	return 0
}

func (m *VersionDef) GetMinConsumer() int32 {
	if m != nil {
		return m.MinConsumer
	}
	return 0
}

func (m *VersionDef) GetBadConsumers() []int32 {
	if m != nil {
		return m.BadConsumers
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionDef)(nil), "tensorflow.VersionDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/versions.proto", fileDescriptor_versions_5c77a333a5601135)
}

var fileDescriptor_versions_5c77a333a5601135 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x65, 0xaa, 0x22, 0x74, 0x14, 0x86, 0x4c, 0x11, 0x53, 0x81, 0x25, 0x0b, 0xf1, 0xc0,
	0x1b, 0x14, 0xc4, 0x5c, 0x75, 0x60, 0x60, 0x41, 0xb1, 0x73, 0x36, 0x16, 0xd8, 0x67, 0xdd, 0x39,
	0xe4, 0xd5, 0x19, 0x11, 0xa5, 0xa4, 0x11, 0xdb, 0xdd, 0xfd, 0xdf, 0xf0, 0xfd, 0x07, 0x4d, 0xc1,
	0x24, 0xc4, 0xee, 0x83, 0x46, 0x6d, 0x89, 0x51, 0x3b, 0xee, 0x22, 0x8e, 0xc4, 0xef, 0xfa, 0x13,
	0x59, 0x02, 0x25, 0x69, 0x33, 0x53, 0xa1, 0x0a, 0x8e, 0xe4, 0x4d, 0x06, 0x78, 0xfe, 0x4d, 0x1f,
	0xd1, 0x55, 0x57, 0x70, 0x96, 0x99, 0xfa, 0xc1, 0x22, 0xd7, 0x6a, 0xad, 0x9a, 0xe5, 0x6e, 0xda,
	0xab, 0x6b, 0x58, 0xc5, 0x90, 0x5e, 0x2d, 0x25, 0x19, 0x22, 0x72, 0x7d, 0xb2, 0xcf, 0xcf, 0x63,
	0x48, 0x0f, 0x87, 0x53, 0x75, 0x0b, 0x17, 0xa6, 0xeb, 0x27, 0x44, 0xea, 0xc5, 0x7a, 0xd1, 0x2c,
	0x77, 0x2b, 0xd3, 0xf5, 0x7f, 0x8c, 0x6c, 0x46, 0xa8, 0x89, 0x7d, 0x7b, 0x74, 0x68, 0x27, 0xd1,
	0xcd, 0xe5, 0xc1, 0x45, 0xb6, 0x3f, 0xa2, 0xb2, 0x55, 0x2f, 0x4f, 0x3e, 0x94, 0xb7, 0xc1, 0xb4,
	0x96, 0xa2, 0x4e, 0x58, 0x0c, 0x77, 0x21, 0xe9, 0xe2, 0xee, 0x3c, 0x67, 0xab, 0x3d, 0xe9, 0x59,
	0xe9, 0xd9, 0xe8, 0xe9, 0xdf, 0x0b, 0xbe, 0x94, 0x32, 0xa7, 0xfb, 0xf6, 0xf7, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xa3, 0x78, 0xae, 0x00, 0x29, 0x01, 0x00, 0x00,
}
