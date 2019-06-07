// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/stream_executor/logging.proto

package stream_executor

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

type CudnnVersion struct {
	Major                int32    `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor                int32    `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch                int32    `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CudnnVersion) Reset()         { *m = CudnnVersion{} }
func (m *CudnnVersion) String() string { return proto.CompactTextString(m) }
func (*CudnnVersion) ProtoMessage()    {}
func (*CudnnVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f78c0ee808dbfb0, []int{0}
}

func (m *CudnnVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CudnnVersion.Unmarshal(m, b)
}
func (m *CudnnVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CudnnVersion.Marshal(b, m, deterministic)
}
func (m *CudnnVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CudnnVersion.Merge(m, src)
}
func (m *CudnnVersion) XXX_Size() int {
	return xxx_messageInfo_CudnnVersion.Size(m)
}
func (m *CudnnVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_CudnnVersion.DiscardUnknown(m)
}

var xxx_messageInfo_CudnnVersion proto.InternalMessageInfo

func (m *CudnnVersion) GetMajor() int32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *CudnnVersion) GetMinor() int32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *CudnnVersion) GetPatch() int32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

type ComputeCapability struct {
	Major                int32    `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor                int32    `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeCapability) Reset()         { *m = ComputeCapability{} }
func (m *ComputeCapability) String() string { return proto.CompactTextString(m) }
func (*ComputeCapability) ProtoMessage()    {}
func (*ComputeCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f78c0ee808dbfb0, []int{1}
}

func (m *ComputeCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeCapability.Unmarshal(m, b)
}
func (m *ComputeCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeCapability.Marshal(b, m, deterministic)
}
func (m *ComputeCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeCapability.Merge(m, src)
}
func (m *ComputeCapability) XXX_Size() int {
	return xxx_messageInfo_ComputeCapability.Size(m)
}
func (m *ComputeCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeCapability.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeCapability proto.InternalMessageInfo

func (m *ComputeCapability) GetMajor() int32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *ComputeCapability) GetMinor() int32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

type CudaInfo struct {
	CudnnVersion         *CudnnVersion      `protobuf:"bytes,1,opt,name=cudnn_version,json=cudnnVersion,proto3" json:"cudnn_version,omitempty"`
	ComputeCapability    *ComputeCapability `protobuf:"bytes,2,opt,name=compute_capability,json=computeCapability,proto3" json:"compute_capability,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CudaInfo) Reset()         { *m = CudaInfo{} }
func (m *CudaInfo) String() string { return proto.CompactTextString(m) }
func (*CudaInfo) ProtoMessage()    {}
func (*CudaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f78c0ee808dbfb0, []int{2}
}

func (m *CudaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CudaInfo.Unmarshal(m, b)
}
func (m *CudaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CudaInfo.Marshal(b, m, deterministic)
}
func (m *CudaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CudaInfo.Merge(m, src)
}
func (m *CudaInfo) XXX_Size() int {
	return xxx_messageInfo_CudaInfo.Size(m)
}
func (m *CudaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CudaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CudaInfo proto.InternalMessageInfo

func (m *CudaInfo) GetCudnnVersion() *CudnnVersion {
	if m != nil {
		return m.CudnnVersion
	}
	return nil
}

func (m *CudaInfo) GetComputeCapability() *ComputeCapability {
	if m != nil {
		return m.ComputeCapability
	}
	return nil
}

func init() {
	proto.RegisterType((*CudnnVersion)(nil), "stream_executor.CudnnVersion")
	proto.RegisterType((*ComputeCapability)(nil), "stream_executor.ComputeCapability")
	proto.RegisterType((*CudaInfo)(nil), "stream_executor.CudaInfo")
}

func init() {
	proto.RegisterFile("tensorflow/stream_executor/logging.proto", fileDescriptor_9f78c0ee808dbfb0)
}

var fileDescriptor_9f78c0ee808dbfb0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x14, 0xc4, 0xa9, 0xb2, 0x22, 0xd9, 0x15, 0xd9, 0xe2, 0xa1, 0x17, 0x41, 0x7a, 0xda, 0x8b, 0x0d,
	0xac, 0x1f, 0x40, 0xd8, 0x1e, 0x64, 0x6f, 0xda, 0x83, 0x07, 0x2f, 0x25, 0xcd, 0xa6, 0xd9, 0x48,
	0x9b, 0x17, 0x5e, 0x5f, 0xfc, 0xf3, 0x71, 0xfc, 0xa6, 0xd2, 0x54, 0xa4, 0xc4, 0x93, 0xb7, 0xcc,
	0x04, 0xe6, 0xfd, 0x86, 0x61, 0x1b, 0x52, 0x76, 0x00, 0x6c, 0x3b, 0x78, 0xe7, 0x03, 0xa1, 0x12,
	0x7d, 0xad, 0x3e, 0x94, 0xf4, 0x04, 0xc8, 0x3b, 0xd0, 0xda, 0x58, 0x5d, 0x38, 0x04, 0x82, 0xf4,
	0x32, 0xfa, 0xce, 0x1f, 0xd9, 0xaa, 0xf4, 0x07, 0x6b, 0x9f, 0x15, 0x0e, 0x06, 0x6c, 0x7a, 0xc5,
	0x16, 0xbd, 0x78, 0x05, 0xcc, 0x92, 0x9b, 0x64, 0xb3, 0xa8, 0x26, 0x11, 0x5c, 0x63, 0x01, 0xb3,
	0x93, 0x1f, 0x77, 0x14, 0xa3, 0xeb, 0x04, 0xc9, 0x63, 0x76, 0x3a, 0xb9, 0x41, 0xe4, 0xf7, 0x6c,
	0x5d, 0x42, 0xef, 0x3c, 0xa9, 0x52, 0x38, 0xd1, 0x98, 0xce, 0xd0, 0xe7, 0x7f, 0x62, 0xf3, 0xaf,
	0x84, 0x9d, 0x97, 0xfe, 0x20, 0xf6, 0xb6, 0x85, 0x74, 0xc7, 0x2e, 0xe4, 0xc8, 0x57, 0xbf, 0x4d,
	0x80, 0x21, 0x60, 0xb9, 0xbd, 0x2e, 0xa2, 0x22, 0xc5, 0xbc, 0x45, 0xb5, 0x92, 0xf3, 0x4e, 0x4f,
	0x2c, 0x95, 0x13, 0x51, 0x2d, 0x7f, 0x91, 0xc2, 0xcd, 0xe5, 0x36, 0xff, 0x1b, 0x14, 0xc3, 0x57,
	0x6b, 0x19, 0x5b, 0xbb, 0xfd, 0xcb, 0x83, 0x36, 0x74, 0xf4, 0x4d, 0x21, 0xa1, 0xe7, 0x56, 0x51,
	0x83, 0xc2, 0x58, 0x4e, 0xed, 0xad, 0x46, 0x27, 0xb9, 0x06, 0x3e, 0x9b, 0x64, 0xf6, 0xd4, 0x10,
	0x0f, 0xd4, 0x9c, 0x85, 0x65, 0xee, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xaa, 0xbd, 0x3b,
	0xc5, 0x01, 0x00, 0x00,
}
