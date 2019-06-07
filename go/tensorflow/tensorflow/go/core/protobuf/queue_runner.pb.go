// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/queue_runner.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	core "github.com/netbrain/tf-grpc/go/tensorflow/tensorflow/go/core/lib/core"
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

// Protocol buffer representing a QueueRunner.
type QueueRunnerDef struct {
	// Queue name.
	QueueName string `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	// A list of enqueue operations.
	EnqueueOpName []string `protobuf:"bytes,2,rep,name=enqueue_op_name,json=enqueueOpName,proto3" json:"enqueue_op_name,omitempty"`
	// The operation to run to close the queue.
	CloseOpName string `protobuf:"bytes,3,opt,name=close_op_name,json=closeOpName,proto3" json:"close_op_name,omitempty"`
	// The operation to run to cancel the queue.
	CancelOpName string `protobuf:"bytes,4,opt,name=cancel_op_name,json=cancelOpName,proto3" json:"cancel_op_name,omitempty"`
	// A list of exception types considered to signal a safely closed queue
	// if raised during enqueue operations.
	QueueClosedExceptionTypes []core.Code `protobuf:"varint,5,rep,packed,name=queue_closed_exception_types,json=queueClosedExceptionTypes,proto3,enum=tensorflow.error.Code" json:"queue_closed_exception_types,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}    `json:"-"`
	XXX_unrecognized          []byte      `json:"-"`
	XXX_sizecache             int32       `json:"-"`
}

func (m *QueueRunnerDef) Reset()         { *m = QueueRunnerDef{} }
func (m *QueueRunnerDef) String() string { return proto.CompactTextString(m) }
func (*QueueRunnerDef) ProtoMessage()    {}
func (*QueueRunnerDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_7af35200d68d14ae, []int{0}
}

func (m *QueueRunnerDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueRunnerDef.Unmarshal(m, b)
}
func (m *QueueRunnerDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueRunnerDef.Marshal(b, m, deterministic)
}
func (m *QueueRunnerDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueRunnerDef.Merge(m, src)
}
func (m *QueueRunnerDef) XXX_Size() int {
	return xxx_messageInfo_QueueRunnerDef.Size(m)
}
func (m *QueueRunnerDef) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueRunnerDef.DiscardUnknown(m)
}

var xxx_messageInfo_QueueRunnerDef proto.InternalMessageInfo

func (m *QueueRunnerDef) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *QueueRunnerDef) GetEnqueueOpName() []string {
	if m != nil {
		return m.EnqueueOpName
	}
	return nil
}

func (m *QueueRunnerDef) GetCloseOpName() string {
	if m != nil {
		return m.CloseOpName
	}
	return ""
}

func (m *QueueRunnerDef) GetCancelOpName() string {
	if m != nil {
		return m.CancelOpName
	}
	return ""
}

func (m *QueueRunnerDef) GetQueueClosedExceptionTypes() []core.Code {
	if m != nil {
		return m.QueueClosedExceptionTypes
	}
	return nil
}

func init() {
	proto.RegisterType((*QueueRunnerDef)(nil), "tensorflow.QueueRunnerDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/queue_runner.proto", fileDescriptor_7af35200d68d14ae)
}

var fileDescriptor_7af35200d68d14ae = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0x89, 0x53, 0x61, 0xd1, 0x4d, 0xdc, 0x83, 0x4c, 0x51, 0x18, 0x43, 0x64, 0x28, 0x36,
	0xa0, 0xdf, 0x60, 0x73, 0xaf, 0x3a, 0x8b, 0x20, 0xf8, 0x52, 0xda, 0xec, 0x5a, 0x8b, 0x6d, 0x2e,
	0x5e, 0x53, 0xe6, 0xbe, 0xb9, 0xf8, 0x24, 0xbb, 0x4c, 0x57, 0x7d, 0xbb, 0xfc, 0xf9, 0xfd, 0xee,
	0xb8, 0x8b, 0xbc, 0x72, 0x60, 0x2a, 0xa4, 0xb4, 0xc0, 0x85, 0xd2, 0x48, 0xa0, 0x2c, 0xa1, 0xc3,
	0xa4, 0x4e, 0xd5, 0x7b, 0x0d, 0x35, 0x44, 0x54, 0x1b, 0x03, 0x14, 0x70, 0xda, 0x93, 0x1b, 0xf8,
	0xe4, 0xf2, 0xbf, 0x58, 0xe4, 0x89, 0x2f, 0x80, 0x08, 0x29, 0xd2, 0x38, 0x87, 0xca, 0x7b, 0xc3,
	0x2f, 0x21, 0xbb, 0x8f, 0xab, 0x76, 0x21, 0x77, 0xbb, 0x83, 0xb4, 0x77, 0x26, 0xa5, 0x1f, 0x60,
	0xe2, 0x12, 0xfa, 0x62, 0x20, 0x46, 0xed, 0xb0, 0xcd, 0xc9, 0x7d, 0x5c, 0x42, 0xef, 0x42, 0x1e,
	0x80, 0xf1, 0x00, 0x5a, 0xcf, 0x6c, 0x0d, 0x5a, 0xa3, 0x76, 0xd8, 0x59, 0xc7, 0x0f, 0x96, 0xb9,
	0xa1, 0xec, 0xe8, 0x02, 0xab, 0x0d, 0xd5, 0xe2, 0x4e, 0x7b, 0x1c, 0xae, 0x99, 0x73, 0xd9, 0xd5,
	0xb1, 0xd1, 0x50, 0xfc, 0x42, 0xdb, 0x0c, 0xed, 0xfb, 0x74, 0x4d, 0x3d, 0xcb, 0x53, 0x3f, 0x8f,
	0xd5, 0x79, 0x04, 0x1f, 0x1a, 0xac, 0xcb, 0xd1, 0x44, 0x6e, 0x69, 0xa1, 0xea, 0xef, 0x0c, 0x5a,
	0xa3, 0xee, 0xcd, 0x51, 0xb0, 0x59, 0x3b, 0xe0, 0x45, 0x83, 0x09, 0xce, 0x21, 0x3c, 0x66, 0x77,
	0xc2, 0xea, 0xf4, 0xc7, 0x7c, 0x5a, 0x89, 0xe3, 0xa5, 0xec, 0x23, 0x65, 0x4d, 0x2f, 0xa5, 0xb8,
	0x84, 0x05, 0xd2, 0xdb, 0xf8, 0xb0, 0x71, 0x95, 0xd9, 0xea, 0x54, 0xd5, 0x4c, 0xbc, 0x4c, 0xb3,
	0xdc, 0xbd, 0xd6, 0x49, 0xa0, 0xb1, 0x54, 0x06, 0x5c, 0x42, 0x71, 0x6e, 0x94, 0x4b, 0xaf, 0x33,
	0xb2, 0x5a, 0x65, 0xa8, 0x1a, 0x87, 0x6f, 0x94, 0x19, 0xfe, 0xfd, 0xbf, 0x4f, 0x21, 0x92, 0x5d,
	0x7e, 0xdc, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x9a, 0x8a, 0x7f, 0xe5, 0x01, 0x00, 0x00,
}