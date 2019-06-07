// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/allocation_description.proto

package framework

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

type AllocationDescription struct {
	// Total number of bytes requested
	RequestedBytes int64 `protobuf:"varint,1,opt,name=requested_bytes,json=requestedBytes,proto3" json:"requested_bytes,omitempty"`
	// Total number of bytes allocated if known
	AllocatedBytes int64 `protobuf:"varint,2,opt,name=allocated_bytes,json=allocatedBytes,proto3" json:"allocated_bytes,omitempty"`
	// Name of the allocator used
	AllocatorName string `protobuf:"bytes,3,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// Identifier of the allocated buffer if known
	AllocationId int64 `protobuf:"varint,4,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Set if this tensor only has one remaining reference
	HasSingleReference bool `protobuf:"varint,5,opt,name=has_single_reference,json=hasSingleReference,proto3" json:"has_single_reference,omitempty"`
	// Address of the allocation.
	Ptr                  uint64   `protobuf:"varint,6,opt,name=ptr,proto3" json:"ptr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocationDescription) Reset()         { *m = AllocationDescription{} }
func (m *AllocationDescription) String() string { return proto.CompactTextString(m) }
func (*AllocationDescription) ProtoMessage()    {}
func (*AllocationDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_1254702e9f0c7d2f, []int{0}
}

func (m *AllocationDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationDescription.Unmarshal(m, b)
}
func (m *AllocationDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationDescription.Marshal(b, m, deterministic)
}
func (m *AllocationDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationDescription.Merge(m, src)
}
func (m *AllocationDescription) XXX_Size() int {
	return xxx_messageInfo_AllocationDescription.Size(m)
}
func (m *AllocationDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationDescription.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationDescription proto.InternalMessageInfo

func (m *AllocationDescription) GetRequestedBytes() int64 {
	if m != nil {
		return m.RequestedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatedBytes() int64 {
	if m != nil {
		return m.AllocatedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *AllocationDescription) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *AllocationDescription) GetHasSingleReference() bool {
	if m != nil {
		return m.HasSingleReference
	}
	return false
}

func (m *AllocationDescription) GetPtr() uint64 {
	if m != nil {
		return m.Ptr
	}
	return 0
}

func init() {
	proto.RegisterType((*AllocationDescription)(nil), "tensorflow.AllocationDescription")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/allocation_description.proto", fileDescriptor_1254702e9f0c7d2f)
}

var fileDescriptor_1254702e9f0c7d2f = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x89, 0x9b, 0x43, 0x83, 0x53, 0x29, 0x0a, 0x05, 0x2f, 0x45, 0x11, 0x7b, 0xb1, 0x11,
	0x04, 0xef, 0x0e, 0x11, 0xbc, 0xc8, 0xa8, 0x37, 0x2f, 0x25, 0x6d, 0x5f, 0xb3, 0x60, 0x9b, 0x57,
	0x5f, 0x32, 0x86, 0x77, 0xff, 0x68, 0xbd, 0x49, 0x37, 0x97, 0x16, 0xf1, 0xf6, 0xf8, 0xbc, 0x4f,
	0x7e, 0xf0, 0xfd, 0xf2, 0x3b, 0x07, 0xc6, 0x22, 0x55, 0x35, 0xae, 0x44, 0x81, 0x04, 0xa2, 0x22,
	0xd9, 0xc0, 0x0a, 0xe9, 0x4d, 0xc8, 0xba, 0xc6, 0x42, 0x3a, 0x8d, 0x26, 0x2b, 0xc1, 0x16, 0xa4,
	0xdb, 0x6e, 0x4e, 0x5a, 0x42, 0x87, 0x01, 0xef, 0xcf, 0x9d, 0x7f, 0x33, 0x7e, 0x7a, 0xef, 0xe5,
	0x87, 0xde, 0x0d, 0xae, 0xf8, 0x11, 0xc1, 0xfb, 0x12, 0xac, 0x83, 0x32, 0xcb, 0x3f, 0x1c, 0xd8,
	0x90, 0x45, 0x2c, 0x1e, 0xa5, 0x87, 0x1e, 0xcf, 0x3a, 0xda, 0x89, 0xbf, 0xcf, 0x79, 0x71, 0x67,
	0x23, 0x7a, 0xbc, 0x11, 0x2f, 0xf9, 0x96, 0x20, 0x65, 0x46, 0x36, 0x10, 0x8e, 0x22, 0x16, 0xef,
	0xa7, 0x53, 0x4f, 0x9f, 0x65, 0x03, 0xc1, 0x05, 0x9f, 0x0e, 0xbe, 0xaf, 0xcb, 0x70, 0xbc, 0xbe,
	0xed, 0xa0, 0x87, 0x4f, 0x65, 0x70, 0xc3, 0x4f, 0x16, 0xd2, 0x66, 0x56, 0x1b, 0x55, 0x43, 0x46,
	0x50, 0x01, 0x81, 0x29, 0x20, 0xdc, 0x8d, 0x58, 0xbc, 0x97, 0x06, 0x0b, 0x69, 0x5f, 0xd6, 0xab,
	0x74, 0xbb, 0x09, 0x8e, 0xf9, 0xa8, 0x75, 0x14, 0x4e, 0x22, 0x16, 0x8f, 0xd3, 0x6e, 0x9c, 0x7d,
	0x32, 0x1e, 0x22, 0xa9, 0xa4, 0x8f, 0x23, 0xf1, 0x09, 0xce, 0xce, 0xfe, 0x4d, 0x65, 0xde, 0x05,
	0x68, 0xe7, 0xec, 0xf5, 0x51, 0x69, 0xb7, 0x58, 0xe6, 0x49, 0x81, 0x8d, 0x30, 0xe0, 0x72, 0x92,
	0xda, 0x08, 0x57, 0x5d, 0x2b, 0x6a, 0x0b, 0xa1, 0x50, 0x0c, 0xaa, 0x19, 0x8c, 0x0a, 0xff, 0x14,
	0xf5, 0xc5, 0x58, 0x3e, 0x59, 0xb7, 0x72, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xcf, 0xa3,
	0x33, 0xcf, 0x01, 0x00, 0x00,
}
