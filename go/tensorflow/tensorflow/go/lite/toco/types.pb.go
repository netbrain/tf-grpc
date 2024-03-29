// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/lite/toco/types.proto

package toco // import "github.com/netbrain/tf-grpc/go/tensorflow/tensorflow/go/lite/toco"

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

// IODataType describes the numeric data types of input and output arrays
// of a model.
type IODataType int32

const (
	IODataType_IO_DATA_TYPE_UNKNOWN IODataType = 0
	// Float32, not quantized
	IODataType_FLOAT IODataType = 1
	// Uint8, quantized
	IODataType_QUANTIZED_UINT8 IODataType = 2
	// Int32, not quantized
	IODataType_INT32 IODataType = 3
	// Int64, not quantized
	IODataType_INT64 IODataType = 4
	// String, not quantized
	IODataType_STRING IODataType = 5
	// Int16, quantized
	IODataType_QUANTIZED_INT16 IODataType = 6
	// Boolean
	IODataType_BOOL IODataType = 7
	// Complex64, not quantized
	IODataType_COMPLEX64 IODataType = 8
	// Int8, quantized based on QuantizationParameters in schema.
	IODataType_INT8 IODataType = 9
)

var IODataType_name = map[int32]string{
	0: "IO_DATA_TYPE_UNKNOWN",
	1: "FLOAT",
	2: "QUANTIZED_UINT8",
	3: "INT32",
	4: "INT64",
	5: "STRING",
	6: "QUANTIZED_INT16",
	7: "BOOL",
	8: "COMPLEX64",
	9: "INT8",
}
var IODataType_value = map[string]int32{
	"IO_DATA_TYPE_UNKNOWN": 0,
	"FLOAT":                1,
	"QUANTIZED_UINT8":      2,
	"INT32":                3,
	"INT64":                4,
	"STRING":               5,
	"QUANTIZED_INT16":      6,
	"BOOL":                 7,
	"COMPLEX64":            8,
	"INT8":                 9,
}

func (x IODataType) Enum() *IODataType {
	p := new(IODataType)
	*p = x
	return p
}
func (x IODataType) String() string {
	return proto.EnumName(IODataType_name, int32(x))
}
func (x *IODataType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IODataType_value, data, "IODataType")
	if err != nil {
		return err
	}
	*x = IODataType(value)
	return nil
}
func (IODataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_types_a590f8975207ee0e, []int{0}
}

func init() {
	proto.RegisterEnum("toco.IODataType", IODataType_name, IODataType_value)
}

func init() {
	proto.RegisterFile("tensorflow/lite/toco/types.proto", fileDescriptor_types_a590f8975207ee0e)
}

var fileDescriptor_types_a590f8975207ee0e = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x46, 0xad, 0xa6, 0xb5, 0x19, 0x10, 0x97, 0xd5, 0x83, 0x47, 0xcf, 0x82, 0x5d, 0xd4, 0x12,
	0xbc, 0xa6, 0x4d, 0x94, 0xc5, 0xb8, 0x5b, 0x75, 0x82, 0xda, 0x4b, 0x48, 0xc3, 0x36, 0x06, 0x6a,
	0x26, 0xa4, 0x23, 0xd2, 0x7f, 0xe3, 0x4f, 0x95, 0x88, 0x60, 0xbd, 0x3d, 0xe6, 0x31, 0x1f, 0x3c,
	0x38, 0x65, 0x57, 0xaf, 0xa9, 0x5d, 0xae, 0xe8, 0x53, 0xad, 0x2a, 0x76, 0x8a, 0xa9, 0x20, 0xc5,
	0x9b, 0xc6, 0xad, 0x47, 0x4d, 0x4b, 0x4c, 0xd2, 0xeb, 0x2e, 0x67, 0x5f, 0x3d, 0x00, 0x6d, 0xa3,
	0x9c, 0x73, 0xdc, 0x34, 0x4e, 0x9e, 0xc0, 0xb1, 0xb6, 0x59, 0x14, 0x62, 0x98, 0xe1, 0xeb, 0x2c,
	0xce, 0x52, 0x73, 0x67, 0xec, 0xb3, 0x11, 0x3b, 0xd2, 0x87, 0xfe, 0x4d, 0x62, 0x43, 0x14, 0x3d,
	0x79, 0x04, 0x87, 0x0f, 0x69, 0x68, 0x50, 0xcf, 0xe3, 0x28, 0x4b, 0xb5, 0xc1, 0x6b, 0xb1, 0xdb,
	0x79, 0x6d, 0xf0, 0xea, 0x52, 0xec, 0xfd, 0x62, 0x30, 0x16, 0x9e, 0x04, 0x18, 0x3c, 0xe1, 0xa3,
	0x36, 0xb7, 0xa2, 0xff, 0xff, 0x4d, 0x1b, 0xbc, 0x08, 0xc4, 0x40, 0x0e, 0xc1, 0x9b, 0x58, 0x9b,
	0x88, 0x7d, 0x79, 0x00, 0xfe, 0xd4, 0xde, 0xcf, 0x92, 0xf8, 0x25, 0x18, 0x8b, 0x61, 0x27, 0x7e,
	0x96, 0xfd, 0xc9, 0x74, 0x1e, 0x96, 0x15, 0xbf, 0x7d, 0x2c, 0x46, 0x05, 0xbd, 0xab, 0xda, 0xf1,
	0xa2, 0xcd, 0xab, 0x5a, 0xf1, 0xf2, 0xbc, 0x6c, 0x9b, 0x42, 0x95, 0xa4, 0xb6, 0x5a, 0xb7, 0xb0,
	0xa4, 0xbf, 0xf2, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0x28, 0x0d, 0xd7, 0x10, 0x01, 0x00,
	0x00,
}
