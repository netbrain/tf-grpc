// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/tensor_bundle.proto

package protobuf // import "github.com/netbrain/tf-grpc/go/tensorflow/tensorflow/go/core/protobuf"

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

// An enum indicating the endianness of the platform that produced this
// bundle.  A bundle can only be read by a platform with matching endianness.
// Defaults to LITTLE, as most modern platforms are little-endian.
//
// Affects the binary tensor data bytes only, not the metadata in protobufs.
type BundleHeaderProto_Endianness int32

const (
	BundleHeaderProto_LITTLE BundleHeaderProto_Endianness = 0
	BundleHeaderProto_BIG    BundleHeaderProto_Endianness = 1
)

var BundleHeaderProto_Endianness_name = map[int32]string{
	0: "LITTLE",
	1: "BIG",
}
var BundleHeaderProto_Endianness_value = map[string]int32{
	"LITTLE": 0,
	"BIG":    1,
}

func (x BundleHeaderProto_Endianness) String() string {
	return proto.EnumName(BundleHeaderProto_Endianness_name, int32(x))
}
func (BundleHeaderProto_Endianness) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tensor_bundle_d2e7fe6edd69b13b, []int{0, 0}
}

// Special header that is associated with a bundle.
//
// TODO(zongheng,zhifengc): maybe in the future, we can add information about
// which binary produced this checkpoint, timestamp, etc. Sometime, these can be
// valuable debugging information. And if needed, these can be used as defensive
// information ensuring reader (binary version) of the checkpoint and the writer
// (binary version) must match within certain range, etc.
type BundleHeaderProto struct {
	// Number of data files in the bundle.
	NumShards  int32                        `protobuf:"varint,1,opt,name=num_shards,json=numShards,proto3" json:"num_shards,omitempty"`
	Endianness BundleHeaderProto_Endianness `protobuf:"varint,2,opt,name=endianness,proto3,enum=tensorflow.BundleHeaderProto_Endianness" json:"endianness,omitempty"`
	// Versioning of the tensor bundle format.
	Version              *framework.VersionDef `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BundleHeaderProto) Reset()         { *m = BundleHeaderProto{} }
func (m *BundleHeaderProto) String() string { return proto.CompactTextString(m) }
func (*BundleHeaderProto) ProtoMessage()    {}
func (*BundleHeaderProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_tensor_bundle_d2e7fe6edd69b13b, []int{0}
}
func (m *BundleHeaderProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleHeaderProto.Unmarshal(m, b)
}
func (m *BundleHeaderProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleHeaderProto.Marshal(b, m, deterministic)
}
func (dst *BundleHeaderProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleHeaderProto.Merge(dst, src)
}
func (m *BundleHeaderProto) XXX_Size() int {
	return xxx_messageInfo_BundleHeaderProto.Size(m)
}
func (m *BundleHeaderProto) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleHeaderProto.DiscardUnknown(m)
}

var xxx_messageInfo_BundleHeaderProto proto.InternalMessageInfo

func (m *BundleHeaderProto) GetNumShards() int32 {
	if m != nil {
		return m.NumShards
	}
	return 0
}

func (m *BundleHeaderProto) GetEndianness() BundleHeaderProto_Endianness {
	if m != nil {
		return m.Endianness
	}
	return BundleHeaderProto_LITTLE
}

func (m *BundleHeaderProto) GetVersion() *framework.VersionDef {
	if m != nil {
		return m.Version
	}
	return nil
}

// Describes the metadata related to a checkpointed tensor.
type BundleEntryProto struct {
	// The tensor dtype and shape.
	Dtype framework.DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape *framework.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	// The binary content of the tensor lies in:
	//   File "shard_id": bytes [offset, offset + size).
	ShardId int32 `protobuf:"varint,3,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	Offset  int64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Size    int64 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	// The CRC32C checksum of the tensor bytes.
	Crc32C uint32 `protobuf:"fixed32,6,opt,name=crc32c,proto3" json:"crc32c,omitempty"`
	// Iff present, this entry represents a partitioned tensor.  The previous
	// fields are interpreted as follows:
	//
	//   "dtype", "shape": describe the full tensor.
	//   "shard_id", "offset", "size", "crc32c": all IGNORED.
	//      These information for each slice can be looked up in their own
	//      BundleEntryProto, keyed by each "slice_name".
	Slices               []*framework.TensorSliceProto `protobuf:"bytes,7,rep,name=slices,proto3" json:"slices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *BundleEntryProto) Reset()         { *m = BundleEntryProto{} }
func (m *BundleEntryProto) String() string { return proto.CompactTextString(m) }
func (*BundleEntryProto) ProtoMessage()    {}
func (*BundleEntryProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_tensor_bundle_d2e7fe6edd69b13b, []int{1}
}
func (m *BundleEntryProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleEntryProto.Unmarshal(m, b)
}
func (m *BundleEntryProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleEntryProto.Marshal(b, m, deterministic)
}
func (dst *BundleEntryProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleEntryProto.Merge(dst, src)
}
func (m *BundleEntryProto) XXX_Size() int {
	return xxx_messageInfo_BundleEntryProto.Size(m)
}
func (m *BundleEntryProto) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleEntryProto.DiscardUnknown(m)
}

var xxx_messageInfo_BundleEntryProto proto.InternalMessageInfo

func (m *BundleEntryProto) GetDtype() framework.DataType {
	if m != nil {
		return m.Dtype
	}
	return framework.DataType_DT_INVALID
}

func (m *BundleEntryProto) GetShape() *framework.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *BundleEntryProto) GetShardId() int32 {
	if m != nil {
		return m.ShardId
	}
	return 0
}

func (m *BundleEntryProto) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *BundleEntryProto) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BundleEntryProto) GetCrc32C() uint32 {
	if m != nil {
		return m.Crc32C
	}
	return 0
}

func (m *BundleEntryProto) GetSlices() []*framework.TensorSliceProto {
	if m != nil {
		return m.Slices
	}
	return nil
}

func init() {
	proto.RegisterType((*BundleHeaderProto)(nil), "tensorflow.BundleHeaderProto")
	proto.RegisterType((*BundleEntryProto)(nil), "tensorflow.BundleEntryProto")
	proto.RegisterEnum("tensorflow.BundleHeaderProto_Endianness", BundleHeaderProto_Endianness_name, BundleHeaderProto_Endianness_value)
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/tensor_bundle.proto", fileDescriptor_tensor_bundle_d2e7fe6edd69b13b)
}

var fileDescriptor_tensor_bundle_d2e7fe6edd69b13b = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0x9d, 0x6e, 0x93, 0xe8, 0x29, 0x94, 0x75, 0x94, 0x12, 0x45, 0x21, 0x2e, 0x08, 0x41,
	0x34, 0x91, 0xd4, 0x27, 0x58, 0xba, 0xd8, 0x85, 0x5e, 0x94, 0xe9, 0xe2, 0x85, 0x37, 0x25, 0x7f,
	0x26, 0xe9, 0xe0, 0xee, 0x4c, 0x98, 0x99, 0x58, 0xd6, 0x17, 0xf0, 0xf9, 0x7c, 0x1b, 0x2f, 0x25,
	0x67, 0xb2, 0x6e, 0xa4, 0x54, 0x7a, 0x37, 0xe7, 0x9c, 0xdf, 0x77, 0xe6, 0x7c, 0x27, 0x13, 0x78,
	0x6f, 0xb9, 0x34, 0x4a, 0xd7, 0x6b, 0x75, 0x9b, 0x96, 0x4a, 0xf3, 0xb4, 0xd5, 0xca, 0xaa, 0xa2,
	0xab, 0x53, 0x57, 0xb8, 0x2e, 0x3a, 0x59, 0xad, 0x79, 0x82, 0x69, 0x0a, 0x7b, 0xfa, 0xe5, 0x1d,
	0x65, 0xad, 0xf3, 0x0d, 0xbf, 0x55, 0xfa, 0xdb, 0x4e, 0x6a, 0x6e, 0xf2, 0x76, 0x50, 0x3e, 0x84,
	0x5e, 0x8b, 0x72, 0x47, 0xbf, 0xfd, 0x0f, 0xbd, 0x6d, 0xb9, 0x19, 0xb0, 0xf8, 0x7e, 0xec, 0x3b,
	0xd7, 0x46, 0x28, 0x39, 0x90, 0xb3, 0x5f, 0x04, 0x9e, 0xce, 0xd1, 0xc9, 0x39, 0xcf, 0x2b, 0xae,
	0x2f, 0xd1, 0xce, 0x6b, 0x00, 0xd9, 0x6d, 0xfa, 0x39, 0x75, 0x65, 0x42, 0x12, 0x91, 0xd8, 0x63,
	0x4f, 0x64, 0xb7, 0xb9, 0xc2, 0x04, 0x3d, 0x07, 0xe0, 0xb2, 0x12, 0xb9, 0x94, 0xdc, 0x98, 0xf0,
	0x20, 0x22, 0xf1, 0x71, 0x16, 0x27, 0xfb, 0x3b, 0x93, 0x3b, 0x1d, 0x93, 0xc5, 0x5f, 0x9e, 0x8d,
	0xb4, 0xf4, 0x23, 0x04, 0xc3, 0x40, 0xe1, 0x24, 0x22, 0xf1, 0x51, 0x76, 0x32, 0x6e, 0xf3, 0xc5,
	0x95, 0xce, 0x78, 0xcd, 0x76, 0xd8, 0xec, 0x0d, 0xc0, 0xbe, 0x17, 0x05, 0xf0, 0x2f, 0x96, 0xab,
	0xd5, 0xc5, 0x62, 0xfa, 0x88, 0x06, 0x30, 0x99, 0x2f, 0x3f, 0x4f, 0xc9, 0xec, 0xe7, 0x01, 0x4c,
	0xdd, 0x04, 0x0b, 0x69, 0xf5, 0xd6, 0x59, 0x7a, 0x07, 0x5e, 0xd5, 0xaf, 0x08, 0xdd, 0x1c, 0x67,
	0xcf, 0xc7, 0xf7, 0x9c, 0xe5, 0x36, 0x5f, 0x6d, 0x5b, 0xce, 0x1c, 0x42, 0x33, 0xf0, 0xf0, 0x13,
	0xa1, 0xb5, 0xa3, 0xec, 0xd5, 0x98, 0x5d, 0xe1, 0xf1, 0xaa, 0x2f, 0x63, 0x63, 0xe6, 0x50, 0xfa,
	0x02, 0x1e, 0xe3, 0xba, 0xae, 0x45, 0x85, 0x56, 0x3c, 0x16, 0x60, 0xbc, 0xac, 0xe8, 0x09, 0xf8,
	0xaa, 0xae, 0x0d, 0xb7, 0xe1, 0x61, 0x44, 0xe2, 0x09, 0x1b, 0x22, 0x4a, 0xe1, 0xd0, 0x88, 0x1f,
	0x3c, 0xf4, 0x30, 0x8b, 0xe7, 0x9e, 0x2d, 0x75, 0x79, 0x9a, 0x95, 0xa1, 0x1f, 0x91, 0x38, 0x60,
	0x43, 0x44, 0x3f, 0x81, 0x8f, 0xef, 0xc0, 0x84, 0x41, 0x34, 0xb9, 0x67, 0xa6, 0xbe, 0xee, 0x66,
	0x1a, 0xd8, 0x79, 0x07, 0xcf, 0x94, 0x6e, 0xc6, 0x68, 0x67, 0xc5, 0x7a, 0x4e, 0x9d, 0xc0, 0xed,
	0x08, 0x15, 0xe6, 0x92, 0x7c, 0x5d, 0x34, 0xc2, 0xde, 0x74, 0x45, 0x52, 0xaa, 0x4d, 0x2a, 0xb9,
	0x2d, 0x74, 0x2e, 0x64, 0x6a, 0xeb, 0x0f, 0x8d, 0x6e, 0xcb, 0xb4, 0x51, 0xe9, 0xe8, 0x4d, 0x8d,
	0x8e, 0x8d, 0xfa, 0xf7, 0xf7, 0xf8, 0x4d, 0x48, 0xe1, 0x63, 0x70, 0xfa, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x9a, 0x04, 0x00, 0x27, 0x44, 0x03, 0x00, 0x00,
}
