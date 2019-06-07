// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/reader_base.proto

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

// For serializing and restoring the state of ReaderBase, see
// reader_base.h for details.
type ReaderBaseState struct {
	WorkStarted          int64    `protobuf:"varint,1,opt,name=work_started,json=workStarted,proto3" json:"work_started,omitempty"`
	WorkFinished         int64    `protobuf:"varint,2,opt,name=work_finished,json=workFinished,proto3" json:"work_finished,omitempty"`
	NumRecordsProduced   int64    `protobuf:"varint,3,opt,name=num_records_produced,json=numRecordsProduced,proto3" json:"num_records_produced,omitempty"`
	CurrentWork          []byte   `protobuf:"bytes,4,opt,name=current_work,json=currentWork,proto3" json:"current_work,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReaderBaseState) Reset()         { *m = ReaderBaseState{} }
func (m *ReaderBaseState) String() string { return proto.CompactTextString(m) }
func (*ReaderBaseState) ProtoMessage()    {}
func (*ReaderBaseState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8282e7620a01b6, []int{0}
}

func (m *ReaderBaseState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReaderBaseState.Unmarshal(m, b)
}
func (m *ReaderBaseState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReaderBaseState.Marshal(b, m, deterministic)
}
func (m *ReaderBaseState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReaderBaseState.Merge(m, src)
}
func (m *ReaderBaseState) XXX_Size() int {
	return xxx_messageInfo_ReaderBaseState.Size(m)
}
func (m *ReaderBaseState) XXX_DiscardUnknown() {
	xxx_messageInfo_ReaderBaseState.DiscardUnknown(m)
}

var xxx_messageInfo_ReaderBaseState proto.InternalMessageInfo

func (m *ReaderBaseState) GetWorkStarted() int64 {
	if m != nil {
		return m.WorkStarted
	}
	return 0
}

func (m *ReaderBaseState) GetWorkFinished() int64 {
	if m != nil {
		return m.WorkFinished
	}
	return 0
}

func (m *ReaderBaseState) GetNumRecordsProduced() int64 {
	if m != nil {
		return m.NumRecordsProduced
	}
	return 0
}

func (m *ReaderBaseState) GetCurrentWork() []byte {
	if m != nil {
		return m.CurrentWork
	}
	return nil
}

func init() {
	proto.RegisterType((*ReaderBaseState)(nil), "tensorflow.ReaderBaseState")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/reader_base.proto", fileDescriptor_9d8282e7620a01b6)
}

var fileDescriptor_9d8282e7620a01b6 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x59, 0x2b, 0x1e, 0xb6, 0x15, 0x25, 0x78, 0xc8, 0xb1, 0xea, 0xa5, 0x20, 0x66, 0x05,
	0xdf, 0x20, 0x87, 0x9e, 0x43, 0x7a, 0x10, 0xbc, 0x84, 0xcd, 0xee, 0x24, 0x0d, 0x35, 0x3b, 0x61,
	0x76, 0x42, 0xf1, 0xa5, 0x7c, 0x3e, 0x8f, 0xb2, 0xdb, 0x60, 0x4a, 0x6f, 0xcb, 0x3f, 0xdf, 0xec,
	0xf0, 0xfd, 0xf2, 0x85, 0xc1, 0x79, 0xa4, 0xe6, 0x0b, 0x8f, 0xca, 0x20, 0x81, 0x6a, 0x48, 0xf7,
	0x70, 0x44, 0x3a, 0x28, 0x02, 0x6d, 0x81, 0xaa, 0x5a, 0x7b, 0xc8, 0x06, 0x42, 0xc6, 0x44, 0xce,
	0xf0, 0xd3, 0x8f, 0x90, 0x77, 0x65, 0x24, 0x72, 0xed, 0x61, 0xc7, 0x9a, 0x21, 0x79, 0x94, 0xab,
	0xb0, 0x59, 0x79, 0xd6, 0xc4, 0x60, 0x53, 0xb1, 0x16, 0x9b, 0x45, 0xb9, 0x0c, 0xd9, 0xee, 0x14,
	0x25, 0xcf, 0xf2, 0x36, 0x22, 0x4d, 0xe7, 0x3a, 0xbf, 0x07, 0x9b, 0x5e, 0x45, 0x26, 0xee, 0x6d,
	0xa7, 0x2c, 0x79, 0x93, 0x0f, 0x6e, 0xec, 0x2b, 0x02, 0x83, 0x64, 0x7d, 0x35, 0x10, 0xda, 0xd1,
	0x80, 0x4d, 0x17, 0x91, 0x4d, 0xdc, 0xd8, 0x97, 0xa7, 0x51, 0x31, 0x4d, 0xc2, 0x65, 0x33, 0x12,
	0x81, 0xe3, 0x2a, 0xfc, 0x94, 0x5e, 0xaf, 0xc5, 0x66, 0x55, 0x2e, 0xa7, 0xec, 0x03, 0xe9, 0x90,
	0x7f, 0xcb, 0x14, 0xa9, 0xcd, 0x66, 0x85, 0xec, 0x5f, 0x35, 0xbf, 0x9f, 0x4d, 0x8a, 0x60, 0xea,
	0x0b, 0xf1, 0xb9, 0x6d, 0x3b, 0xde, 0x8f, 0x75, 0x66, 0xb0, 0x57, 0x0e, 0xb8, 0x26, 0xdd, 0x39,
	0xc5, 0xcd, 0x6b, 0x4b, 0x83, 0x51, 0x2d, 0xaa, 0xb3, 0xe2, 0xce, 0x9e, 0x2d, 0x5e, 0xd4, 0xf8,
	0x2b, 0x44, 0x7d, 0x13, 0xeb, 0x7b, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x17, 0x4c, 0x6b, 0x61,
	0x6d, 0x01, 0x00, 0x00,
}