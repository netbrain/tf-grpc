// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/saved_model.proto

package protobuf

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

// SavedModel is the high level serialization format for TensorFlow Models.
// See [todo: doc links, similar to session_bundle] for more information.
type SavedModel struct {
	// The schema version of the SavedModel instance. Used for versioning when
	// making future changes to the specification/implementation. Initial value
	// at release will be 1.
	SavedModelSchemaVersion int64 `protobuf:"varint,1,opt,name=saved_model_schema_version,json=savedModelSchemaVersion,proto3" json:"saved_model_schema_version,omitempty"`
	// One or more MetaGraphs.
	MetaGraphs           []*MetaGraphDef `protobuf:"bytes,2,rep,name=meta_graphs,json=metaGraphs,proto3" json:"meta_graphs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SavedModel) Reset()         { *m = SavedModel{} }
func (m *SavedModel) String() string { return proto.CompactTextString(m) }
func (*SavedModel) ProtoMessage()    {}
func (*SavedModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_537826d0bcc2f334, []int{0}
}

func (m *SavedModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedModel.Unmarshal(m, b)
}
func (m *SavedModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedModel.Marshal(b, m, deterministic)
}
func (m *SavedModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedModel.Merge(m, src)
}
func (m *SavedModel) XXX_Size() int {
	return xxx_messageInfo_SavedModel.Size(m)
}
func (m *SavedModel) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedModel.DiscardUnknown(m)
}

var xxx_messageInfo_SavedModel proto.InternalMessageInfo

func (m *SavedModel) GetSavedModelSchemaVersion() int64 {
	if m != nil {
		return m.SavedModelSchemaVersion
	}
	return 0
}

func (m *SavedModel) GetMetaGraphs() []*MetaGraphDef {
	if m != nil {
		return m.MetaGraphs
	}
	return nil
}

func init() {
	proto.RegisterType((*SavedModel)(nil), "tensorflow.SavedModel")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/saved_model.proto", fileDescriptor_537826d0bcc2f334)
}

var fileDescriptor_537826d0bcc2f334 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x59, 0x0b, 0x1e, 0xa6, 0x17, 0xc9, 0xc5, 0xd0, 0x53, 0xf1, 0x54, 0x05, 0xb3, 0xa0,
	0x27, 0xf1, 0x56, 0x14, 0x4f, 0x85, 0xd2, 0x82, 0x07, 0x2f, 0x61, 0x93, 0xce, 0x6e, 0x82, 0xdd,
	0x4c, 0x98, 0xd9, 0xb6, 0xfe, 0x00, 0x7f, 0xb4, 0x47, 0x49, 0x82, 0x66, 0x3d, 0x78, 0xdb, 0xd9,
	0xf7, 0xde, 0xc7, 0xe3, 0xc1, 0x4d, 0xc0, 0x46, 0x88, 0xed, 0x9e, 0x4e, 0xba, 0x24, 0x46, 0xdd,
	0x32, 0x05, 0x2a, 0x0e, 0x56, 0x8b, 0x39, 0xe2, 0x2e, 0xf7, 0xb4, 0xc3, 0x7d, 0xd6, 0x7f, 0x26,
	0x30, 0x7a, 0x67, 0xd7, 0xff, 0xe6, 0x3c, 0x06, 0x93, 0x3b, 0x36, 0x6d, 0x35, 0xc4, 0xae, 0x3e,
	0x15, 0xc0, 0xb6, 0x83, 0xad, 0x3a, 0x56, 0xf2, 0x08, 0xb3, 0x08, 0x9d, 0x4b, 0x59, 0xa1, 0x37,
	0xf9, 0x11, 0x59, 0x6a, 0x6a, 0x52, 0x35, 0x57, 0x8b, 0xc9, 0xe6, 0x52, 0x7e, 0xfd, 0xdb, 0x5e,
	0x7f, 0x1d, 0xe4, 0xe4, 0x01, 0xa6, 0x23, 0x5f, 0xd2, 0xb3, 0xf9, 0x64, 0x31, 0xbd, 0x4b, 0xb3,
	0xb1, 0x4c, 0xb6, 0xc2, 0x60, 0x5e, 0x3a, 0xf5, 0x09, 0xed, 0x06, 0xfc, 0xcf, 0x25, 0xcb, 0x0f,
	0x48, 0x89, 0x5d, 0x6c, 0xb5, 0x6c, 0x3c, 0x9e, 0x88, 0xdf, 0x97, 0x17, 0x63, 0xbf, 0x75, 0xd7,
	0x59, 0xd6, 0xea, 0xed, 0xd9, 0xd5, 0xa1, 0x3a, 0x14, 0x59, 0x49, 0x5e, 0x37, 0x18, 0x0a, 0x36,
	0x75, 0xa3, 0x83, 0xbd, 0x75, 0xdc, 0x96, 0xda, 0x91, 0x8e, 0x06, 0x88, 0x9e, 0x8e, 0xfe, 0xce,
	0xf1, 0xa5, 0x54, 0x71, 0xde, 0x1f, 0xf7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0xe9, 0x0d,
	0x72, 0x6c, 0x01, 0x00, 0x00,
}