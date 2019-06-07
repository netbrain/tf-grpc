// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/tools/proto_text/test.proto

package proto_text

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

type ForeignEnum int32

const (
	ForeignEnum_FOREIGN_ZERO ForeignEnum = 0
	ForeignEnum_FOREIGN_FOO  ForeignEnum = 4
	ForeignEnum_FOREIGN_BAR  ForeignEnum = 5
	ForeignEnum_FOREIGN_BAZ  ForeignEnum = 6
)

var ForeignEnum_name = map[int32]string{
	0: "FOREIGN_ZERO",
	4: "FOREIGN_FOO",
	5: "FOREIGN_BAR",
	6: "FOREIGN_BAZ",
}

var ForeignEnum_value = map[string]int32{
	"FOREIGN_ZERO": 0,
	"FOREIGN_FOO":  4,
	"FOREIGN_BAR":  5,
	"FOREIGN_BAZ":  6,
}

func (x ForeignEnum) String() string {
	return proto.EnumName(ForeignEnum_name, int32(x))
}

func (ForeignEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_54bcf749d548b8e3, []int{0}
}

type TestAllTypes_NestedEnum int32

const (
	TestAllTypes_ZERO TestAllTypes_NestedEnum = 0
	TestAllTypes_FOO  TestAllTypes_NestedEnum = 1
	TestAllTypes_BAR  TestAllTypes_NestedEnum = 2
	TestAllTypes_BAZ  TestAllTypes_NestedEnum = 3
	TestAllTypes_NEG  TestAllTypes_NestedEnum = -1
)

var TestAllTypes_NestedEnum_name = map[int32]string{
	0:  "ZERO",
	1:  "FOO",
	2:  "BAR",
	3:  "BAZ",
	-1: "NEG",
}

var TestAllTypes_NestedEnum_value = map[string]int32{
	"ZERO": 0,
	"FOO":  1,
	"BAR":  2,
	"BAZ":  3,
	"NEG":  -1,
}

func (x TestAllTypes_NestedEnum) String() string {
	return proto.EnumName(TestAllTypes_NestedEnum_name, int32(x))
}

func (TestAllTypes_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_54bcf749d548b8e3, []int{0, 0}
}

type TestAllTypes struct {
	// Singular
	OptionalInt32          int32                       `protobuf:"varint,1000,opt,name=optional_int32,json=optionalInt32,proto3" json:"optional_int32,omitempty"`
	OptionalInt64          int64                       `protobuf:"varint,2,opt,name=optional_int64,json=optionalInt64,proto3" json:"optional_int64,omitempty"`
	OptionalUint32         uint32                      `protobuf:"varint,3,opt,name=optional_uint32,json=optionalUint32,proto3" json:"optional_uint32,omitempty"`
	OptionalUint64         uint64                      `protobuf:"varint,999,opt,name=optional_uint64,json=optionalUint64,proto3" json:"optional_uint64,omitempty"`
	OptionalSint32         int32                       `protobuf:"zigzag32,5,opt,name=optional_sint32,json=optionalSint32,proto3" json:"optional_sint32,omitempty"`
	OptionalSint64         int64                       `protobuf:"zigzag64,6,opt,name=optional_sint64,json=optionalSint64,proto3" json:"optional_sint64,omitempty"`
	OptionalFixed32        uint32                      `protobuf:"fixed32,7,opt,name=optional_fixed32,json=optionalFixed32,proto3" json:"optional_fixed32,omitempty"`
	OptionalFixed64        uint64                      `protobuf:"fixed64,8,opt,name=optional_fixed64,json=optionalFixed64,proto3" json:"optional_fixed64,omitempty"`
	OptionalSfixed32       int32                       `protobuf:"fixed32,9,opt,name=optional_sfixed32,json=optionalSfixed32,proto3" json:"optional_sfixed32,omitempty"`
	OptionalSfixed64       int64                       `protobuf:"fixed64,10,opt,name=optional_sfixed64,json=optionalSfixed64,proto3" json:"optional_sfixed64,omitempty"`
	OptionalFloat          float32                     `protobuf:"fixed32,11,opt,name=optional_float,json=optionalFloat,proto3" json:"optional_float,omitempty"`
	OptionalDouble         float64                     `protobuf:"fixed64,12,opt,name=optional_double,json=optionalDouble,proto3" json:"optional_double,omitempty"`
	OptionalBool           bool                        `protobuf:"varint,13,opt,name=optional_bool,json=optionalBool,proto3" json:"optional_bool,omitempty"`
	OptionalString         string                      `protobuf:"bytes,14,opt,name=optional_string,json=optionalString,proto3" json:"optional_string,omitempty"`
	OptionalBytes          []byte                      `protobuf:"bytes,15,opt,name=optional_bytes,json=optionalBytes,proto3" json:"optional_bytes,omitempty"`
	OptionalNestedMessage  *TestAllTypes_NestedMessage `protobuf:"bytes,18,opt,name=optional_nested_message,json=optionalNestedMessage,proto3" json:"optional_nested_message,omitempty"`
	OptionalForeignMessage *ForeignMessage             `protobuf:"bytes,19,opt,name=optional_foreign_message,json=optionalForeignMessage,proto3" json:"optional_foreign_message,omitempty"`
	OptionalNestedEnum     TestAllTypes_NestedEnum     `protobuf:"varint,21,opt,name=optional_nested_enum,json=optionalNestedEnum,proto3,enum=tensorflow.test.TestAllTypes_NestedEnum" json:"optional_nested_enum,omitempty"`
	OptionalForeignEnum    ForeignEnum                 `protobuf:"varint,22,opt,name=optional_foreign_enum,json=optionalForeignEnum,proto3,enum=tensorflow.test.ForeignEnum" json:"optional_foreign_enum,omitempty"`
	OptionalCord           string                      `protobuf:"bytes,25,opt,name=optional_cord,json=optionalCord,proto3" json:"optional_cord,omitempty"`
	// Repeated
	RepeatedInt32         []int32                       `protobuf:"varint,31,rep,packed,name=repeated_int32,json=repeatedInt32,proto3" json:"repeated_int32,omitempty"`
	RepeatedInt64         []int64                       `protobuf:"varint,32,rep,packed,name=repeated_int64,json=repeatedInt64,proto3" json:"repeated_int64,omitempty"`
	RepeatedUint32        []uint32                      `protobuf:"varint,33,rep,packed,name=repeated_uint32,json=repeatedUint32,proto3" json:"repeated_uint32,omitempty"`
	RepeatedUint64        []uint64                      `protobuf:"varint,34,rep,packed,name=repeated_uint64,json=repeatedUint64,proto3" json:"repeated_uint64,omitempty"`
	RepeatedSint32        []int32                       `protobuf:"zigzag32,35,rep,packed,name=repeated_sint32,json=repeatedSint32,proto3" json:"repeated_sint32,omitempty"`
	RepeatedSint64        []int64                       `protobuf:"zigzag64,36,rep,packed,name=repeated_sint64,json=repeatedSint64,proto3" json:"repeated_sint64,omitempty"`
	RepeatedFixed32       []uint32                      `protobuf:"fixed32,37,rep,packed,name=repeated_fixed32,json=repeatedFixed32,proto3" json:"repeated_fixed32,omitempty"`
	RepeatedFixed64       []uint64                      `protobuf:"fixed64,38,rep,packed,name=repeated_fixed64,json=repeatedFixed64,proto3" json:"repeated_fixed64,omitempty"`
	RepeatedSfixed32      []int32                       `protobuf:"fixed32,39,rep,packed,name=repeated_sfixed32,json=repeatedSfixed32,proto3" json:"repeated_sfixed32,omitempty"`
	RepeatedSfixed64      []int64                       `protobuf:"fixed64,40,rep,packed,name=repeated_sfixed64,json=repeatedSfixed64,proto3" json:"repeated_sfixed64,omitempty"`
	RepeatedFloat         []float32                     `protobuf:"fixed32,41,rep,packed,name=repeated_float,json=repeatedFloat,proto3" json:"repeated_float,omitempty"`
	RepeatedDouble        []float64                     `protobuf:"fixed64,42,rep,packed,name=repeated_double,json=repeatedDouble,proto3" json:"repeated_double,omitempty"`
	RepeatedBool          []bool                        `protobuf:"varint,43,rep,packed,name=repeated_bool,json=repeatedBool,proto3" json:"repeated_bool,omitempty"`
	RepeatedString        []string                      `protobuf:"bytes,44,rep,name=repeated_string,json=repeatedString,proto3" json:"repeated_string,omitempty"`
	RepeatedBytes         [][]byte                      `protobuf:"bytes,45,rep,name=repeated_bytes,json=repeatedBytes,proto3" json:"repeated_bytes,omitempty"`
	RepeatedNestedMessage []*TestAllTypes_NestedMessage `protobuf:"bytes,48,rep,name=repeated_nested_message,json=repeatedNestedMessage,proto3" json:"repeated_nested_message,omitempty"`
	RepeatedNestedEnum    []TestAllTypes_NestedEnum     `protobuf:"varint,51,rep,packed,name=repeated_nested_enum,json=repeatedNestedEnum,proto3,enum=tensorflow.test.TestAllTypes_NestedEnum" json:"repeated_nested_enum,omitempty"`
	RepeatedCord          []string                      `protobuf:"bytes,55,rep,name=repeated_cord,json=repeatedCord,proto3" json:"repeated_cord,omitempty"`
	// Types that are valid to be assigned to OneofField:
	//	*TestAllTypes_OneofUint32
	//	*TestAllTypes_OneofNestedMessage
	//	*TestAllTypes_OneofString
	//	*TestAllTypes_OneofBytes
	//	*TestAllTypes_OneofEnum
	OneofField                isTestAllTypes_OneofField              `protobuf_oneof:"oneof_field"`
	MapStringToMessage        map[string]*TestAllTypes_NestedMessage `protobuf:"bytes,58,rep,name=map_string_to_message,json=mapStringToMessage,proto3" json:"map_string_to_message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapInt32ToMessage         map[int32]*TestAllTypes_NestedMessage  `protobuf:"bytes,59,rep,name=map_int32_to_message,json=mapInt32ToMessage,proto3" json:"map_int32_to_message,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapInt64ToMessage         map[int64]*TestAllTypes_NestedMessage  `protobuf:"bytes,60,rep,name=map_int64_to_message,json=mapInt64ToMessage,proto3" json:"map_int64_to_message,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapBoolToMessage          map[bool]*TestAllTypes_NestedMessage   `protobuf:"bytes,61,rep,name=map_bool_to_message,json=mapBoolToMessage,proto3" json:"map_bool_to_message,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapStringToInt64          map[string]int64                       `protobuf:"bytes,62,rep,name=map_string_to_int64,json=mapStringToInt64,proto3" json:"map_string_to_int64,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapInt64ToString          map[int64]string                       `protobuf:"bytes,63,rep,name=map_int64_to_string,json=mapInt64ToString,proto3" json:"map_int64_to_string,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AnotherMapStringToMessage map[string]*TestAllTypes_NestedMessage `protobuf:"bytes,65,rep,name=another_map_string_to_message,json=anotherMapStringToMessage,proto3" json:"another_map_string_to_message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PackedRepeatedInt64       []int64                                `protobuf:"varint,64,rep,packed,name=packed_repeated_int64,json=packedRepeatedInt64,proto3" json:"packed_repeated_int64,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                               `json:"-"`
	XXX_unrecognized          []byte                                 `json:"-"`
	XXX_sizecache             int32                                  `json:"-"`
}

func (m *TestAllTypes) Reset()         { *m = TestAllTypes{} }
func (m *TestAllTypes) String() string { return proto.CompactTextString(m) }
func (*TestAllTypes) ProtoMessage()    {}
func (*TestAllTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_54bcf749d548b8e3, []int{0}
}

func (m *TestAllTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestAllTypes.Unmarshal(m, b)
}
func (m *TestAllTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestAllTypes.Marshal(b, m, deterministic)
}
func (m *TestAllTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAllTypes.Merge(m, src)
}
func (m *TestAllTypes) XXX_Size() int {
	return xxx_messageInfo_TestAllTypes.Size(m)
}
func (m *TestAllTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAllTypes.DiscardUnknown(m)
}

var xxx_messageInfo_TestAllTypes proto.InternalMessageInfo

func (m *TestAllTypes) GetOptionalInt32() int32 {
	if m != nil {
		return m.OptionalInt32
	}
	return 0
}

func (m *TestAllTypes) GetOptionalInt64() int64 {
	if m != nil {
		return m.OptionalInt64
	}
	return 0
}

func (m *TestAllTypes) GetOptionalUint32() uint32 {
	if m != nil {
		return m.OptionalUint32
	}
	return 0
}

func (m *TestAllTypes) GetOptionalUint64() uint64 {
	if m != nil {
		return m.OptionalUint64
	}
	return 0
}

func (m *TestAllTypes) GetOptionalSint32() int32 {
	if m != nil {
		return m.OptionalSint32
	}
	return 0
}

func (m *TestAllTypes) GetOptionalSint64() int64 {
	if m != nil {
		return m.OptionalSint64
	}
	return 0
}

func (m *TestAllTypes) GetOptionalFixed32() uint32 {
	if m != nil {
		return m.OptionalFixed32
	}
	return 0
}

func (m *TestAllTypes) GetOptionalFixed64() uint64 {
	if m != nil {
		return m.OptionalFixed64
	}
	return 0
}

func (m *TestAllTypes) GetOptionalSfixed32() int32 {
	if m != nil {
		return m.OptionalSfixed32
	}
	return 0
}

func (m *TestAllTypes) GetOptionalSfixed64() int64 {
	if m != nil {
		return m.OptionalSfixed64
	}
	return 0
}

func (m *TestAllTypes) GetOptionalFloat() float32 {
	if m != nil {
		return m.OptionalFloat
	}
	return 0
}

func (m *TestAllTypes) GetOptionalDouble() float64 {
	if m != nil {
		return m.OptionalDouble
	}
	return 0
}

func (m *TestAllTypes) GetOptionalBool() bool {
	if m != nil {
		return m.OptionalBool
	}
	return false
}

func (m *TestAllTypes) GetOptionalString() string {
	if m != nil {
		return m.OptionalString
	}
	return ""
}

func (m *TestAllTypes) GetOptionalBytes() []byte {
	if m != nil {
		return m.OptionalBytes
	}
	return nil
}

func (m *TestAllTypes) GetOptionalNestedMessage() *TestAllTypes_NestedMessage {
	if m != nil {
		return m.OptionalNestedMessage
	}
	return nil
}

func (m *TestAllTypes) GetOptionalForeignMessage() *ForeignMessage {
	if m != nil {
		return m.OptionalForeignMessage
	}
	return nil
}

func (m *TestAllTypes) GetOptionalNestedEnum() TestAllTypes_NestedEnum {
	if m != nil {
		return m.OptionalNestedEnum
	}
	return TestAllTypes_ZERO
}

func (m *TestAllTypes) GetOptionalForeignEnum() ForeignEnum {
	if m != nil {
		return m.OptionalForeignEnum
	}
	return ForeignEnum_FOREIGN_ZERO
}

func (m *TestAllTypes) GetOptionalCord() string {
	if m != nil {
		return m.OptionalCord
	}
	return ""
}

func (m *TestAllTypes) GetRepeatedInt32() []int32 {
	if m != nil {
		return m.RepeatedInt32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedInt64() []int64 {
	if m != nil {
		return m.RepeatedInt64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedUint32() []uint32 {
	if m != nil {
		return m.RepeatedUint32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedUint64() []uint64 {
	if m != nil {
		return m.RepeatedUint64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSint32() []int32 {
	if m != nil {
		return m.RepeatedSint32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSint64() []int64 {
	if m != nil {
		return m.RepeatedSint64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFixed32() []uint32 {
	if m != nil {
		return m.RepeatedFixed32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFixed64() []uint64 {
	if m != nil {
		return m.RepeatedFixed64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSfixed32() []int32 {
	if m != nil {
		return m.RepeatedSfixed32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSfixed64() []int64 {
	if m != nil {
		return m.RepeatedSfixed64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFloat() []float32 {
	if m != nil {
		return m.RepeatedFloat
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedDouble() []float64 {
	if m != nil {
		return m.RepeatedDouble
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedBool() []bool {
	if m != nil {
		return m.RepeatedBool
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedString() []string {
	if m != nil {
		return m.RepeatedString
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedBytes() [][]byte {
	if m != nil {
		return m.RepeatedBytes
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedNestedMessage() []*TestAllTypes_NestedMessage {
	if m != nil {
		return m.RepeatedNestedMessage
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedNestedEnum() []TestAllTypes_NestedEnum {
	if m != nil {
		return m.RepeatedNestedEnum
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedCord() []string {
	if m != nil {
		return m.RepeatedCord
	}
	return nil
}

type isTestAllTypes_OneofField interface {
	isTestAllTypes_OneofField()
}

type TestAllTypes_OneofUint32 struct {
	OneofUint32 uint32 `protobuf:"varint,111,opt,name=oneof_uint32,json=oneofUint32,proto3,oneof"`
}

type TestAllTypes_OneofNestedMessage struct {
	OneofNestedMessage *TestAllTypes_NestedMessage `protobuf:"bytes,112,opt,name=oneof_nested_message,json=oneofNestedMessage,proto3,oneof"`
}

type TestAllTypes_OneofString struct {
	OneofString string `protobuf:"bytes,113,opt,name=oneof_string,json=oneofString,proto3,oneof"`
}

type TestAllTypes_OneofBytes struct {
	OneofBytes []byte `protobuf:"bytes,114,opt,name=oneof_bytes,json=oneofBytes,proto3,oneof"`
}

type TestAllTypes_OneofEnum struct {
	OneofEnum TestAllTypes_NestedEnum `protobuf:"varint,100,opt,name=oneof_enum,json=oneofEnum,proto3,enum=tensorflow.test.TestAllTypes_NestedEnum,oneof"`
}

func (*TestAllTypes_OneofUint32) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofNestedMessage) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofString) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofBytes) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofEnum) isTestAllTypes_OneofField() {}

func (m *TestAllTypes) GetOneofField() isTestAllTypes_OneofField {
	if m != nil {
		return m.OneofField
	}
	return nil
}

func (m *TestAllTypes) GetOneofUint32() uint32 {
	if x, ok := m.GetOneofField().(*TestAllTypes_OneofUint32); ok {
		return x.OneofUint32
	}
	return 0
}

func (m *TestAllTypes) GetOneofNestedMessage() *TestAllTypes_NestedMessage {
	if x, ok := m.GetOneofField().(*TestAllTypes_OneofNestedMessage); ok {
		return x.OneofNestedMessage
	}
	return nil
}

func (m *TestAllTypes) GetOneofString() string {
	if x, ok := m.GetOneofField().(*TestAllTypes_OneofString); ok {
		return x.OneofString
	}
	return ""
}

func (m *TestAllTypes) GetOneofBytes() []byte {
	if x, ok := m.GetOneofField().(*TestAllTypes_OneofBytes); ok {
		return x.OneofBytes
	}
	return nil
}

func (m *TestAllTypes) GetOneofEnum() TestAllTypes_NestedEnum {
	if x, ok := m.GetOneofField().(*TestAllTypes_OneofEnum); ok {
		return x.OneofEnum
	}
	return TestAllTypes_ZERO
}

func (m *TestAllTypes) GetMapStringToMessage() map[string]*TestAllTypes_NestedMessage {
	if m != nil {
		return m.MapStringToMessage
	}
	return nil
}

func (m *TestAllTypes) GetMapInt32ToMessage() map[int32]*TestAllTypes_NestedMessage {
	if m != nil {
		return m.MapInt32ToMessage
	}
	return nil
}

func (m *TestAllTypes) GetMapInt64ToMessage() map[int64]*TestAllTypes_NestedMessage {
	if m != nil {
		return m.MapInt64ToMessage
	}
	return nil
}

func (m *TestAllTypes) GetMapBoolToMessage() map[bool]*TestAllTypes_NestedMessage {
	if m != nil {
		return m.MapBoolToMessage
	}
	return nil
}

func (m *TestAllTypes) GetMapStringToInt64() map[string]int64 {
	if m != nil {
		return m.MapStringToInt64
	}
	return nil
}

func (m *TestAllTypes) GetMapInt64ToString() map[int64]string {
	if m != nil {
		return m.MapInt64ToString
	}
	return nil
}

func (m *TestAllTypes) GetAnotherMapStringToMessage() map[string]*TestAllTypes_NestedMessage {
	if m != nil {
		return m.AnotherMapStringToMessage
	}
	return nil
}

func (m *TestAllTypes) GetPackedRepeatedInt64() []int64 {
	if m != nil {
		return m.PackedRepeatedInt64
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestAllTypes) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestAllTypes_OneofUint32)(nil),
		(*TestAllTypes_OneofNestedMessage)(nil),
		(*TestAllTypes_OneofString)(nil),
		(*TestAllTypes_OneofBytes)(nil),
		(*TestAllTypes_OneofEnum)(nil),
	}
}

type TestAllTypes_NestedMessage struct {
	OptionalInt32        int32                                           `protobuf:"varint,1,opt,name=optional_int32,json=optionalInt32,proto3" json:"optional_int32,omitempty"`
	RepeatedInt32        []int32                                         `protobuf:"varint,2,rep,packed,name=repeated_int32,json=repeatedInt32,proto3" json:"repeated_int32,omitempty"`
	Msg                  *TestAllTypes_NestedMessage_DoubleNestedMessage `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	OptionalInt64        int64                                           `protobuf:"varint,4,opt,name=optional_int64,json=optionalInt64,proto3" json:"optional_int64,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *TestAllTypes_NestedMessage) Reset()         { *m = TestAllTypes_NestedMessage{} }
func (m *TestAllTypes_NestedMessage) String() string { return proto.CompactTextString(m) }
func (*TestAllTypes_NestedMessage) ProtoMessage()    {}
func (*TestAllTypes_NestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_54bcf749d548b8e3, []int{0, 0}
}

func (m *TestAllTypes_NestedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestAllTypes_NestedMessage.Unmarshal(m, b)
}
func (m *TestAllTypes_NestedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestAllTypes_NestedMessage.Marshal(b, m, deterministic)
}
func (m *TestAllTypes_NestedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAllTypes_NestedMessage.Merge(m, src)
}
func (m *TestAllTypes_NestedMessage) XXX_Size() int {
	return xxx_messageInfo_TestAllTypes_NestedMessage.Size(m)
}
func (m *TestAllTypes_NestedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAllTypes_NestedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestAllTypes_NestedMessage proto.InternalMessageInfo

func (m *TestAllTypes_NestedMessage) GetOptionalInt32() int32 {
	if m != nil {
		return m.OptionalInt32
	}
	return 0
}

func (m *TestAllTypes_NestedMessage) GetRepeatedInt32() []int32 {
	if m != nil {
		return m.RepeatedInt32
	}
	return nil
}

func (m *TestAllTypes_NestedMessage) GetMsg() *TestAllTypes_NestedMessage_DoubleNestedMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *TestAllTypes_NestedMessage) GetOptionalInt64() int64 {
	if m != nil {
		return m.OptionalInt64
	}
	return 0
}

type TestAllTypes_NestedMessage_DoubleNestedMessage struct {
	OptionalString       string   `protobuf:"bytes,1,opt,name=optional_string,json=optionalString,proto3" json:"optional_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestAllTypes_NestedMessage_DoubleNestedMessage) Reset() {
	*m = TestAllTypes_NestedMessage_DoubleNestedMessage{}
}
func (m *TestAllTypes_NestedMessage_DoubleNestedMessage) String() string {
	return proto.CompactTextString(m)
}
func (*TestAllTypes_NestedMessage_DoubleNestedMessage) ProtoMessage() {}
func (*TestAllTypes_NestedMessage_DoubleNestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_54bcf749d548b8e3, []int{0, 0, 0}
}

func (m *TestAllTypes_NestedMessage_DoubleNestedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestAllTypes_NestedMessage_DoubleNestedMessage.Unmarshal(m, b)
}
func (m *TestAllTypes_NestedMessage_DoubleNestedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestAllTypes_NestedMessage_DoubleNestedMessage.Marshal(b, m, deterministic)
}
func (m *TestAllTypes_NestedMessage_DoubleNestedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAllTypes_NestedMessage_DoubleNestedMessage.Merge(m, src)
}
func (m *TestAllTypes_NestedMessage_DoubleNestedMessage) XXX_Size() int {
	return xxx_messageInfo_TestAllTypes_NestedMessage_DoubleNestedMessage.Size(m)
}
func (m *TestAllTypes_NestedMessage_DoubleNestedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAllTypes_NestedMessage_DoubleNestedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestAllTypes_NestedMessage_DoubleNestedMessage proto.InternalMessageInfo

func (m *TestAllTypes_NestedMessage_DoubleNestedMessage) GetOptionalString() string {
	if m != nil {
		return m.OptionalString
	}
	return ""
}

// A recursive message.
type NestedTestAllTypes struct {
	Child                *NestedTestAllTypes `protobuf:"bytes,1,opt,name=child,proto3" json:"child,omitempty"`
	Payload              *TestAllTypes       `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	MapStringToInt64     map[string]int64    `protobuf:"bytes,3,rep,name=map_string_to_int64,json=mapStringToInt64,proto3" json:"map_string_to_int64,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NestedTestAllTypes) Reset()         { *m = NestedTestAllTypes{} }
func (m *NestedTestAllTypes) String() string { return proto.CompactTextString(m) }
func (*NestedTestAllTypes) ProtoMessage()    {}
func (*NestedTestAllTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_54bcf749d548b8e3, []int{1}
}

func (m *NestedTestAllTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NestedTestAllTypes.Unmarshal(m, b)
}
func (m *NestedTestAllTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NestedTestAllTypes.Marshal(b, m, deterministic)
}
func (m *NestedTestAllTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NestedTestAllTypes.Merge(m, src)
}
func (m *NestedTestAllTypes) XXX_Size() int {
	return xxx_messageInfo_NestedTestAllTypes.Size(m)
}
func (m *NestedTestAllTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_NestedTestAllTypes.DiscardUnknown(m)
}

var xxx_messageInfo_NestedTestAllTypes proto.InternalMessageInfo

func (m *NestedTestAllTypes) GetChild() *NestedTestAllTypes {
	if m != nil {
		return m.Child
	}
	return nil
}

func (m *NestedTestAllTypes) GetPayload() *TestAllTypes {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *NestedTestAllTypes) GetMapStringToInt64() map[string]int64 {
	if m != nil {
		return m.MapStringToInt64
	}
	return nil
}

type ForeignMessage struct {
	C                    int32    `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForeignMessage) Reset()         { *m = ForeignMessage{} }
func (m *ForeignMessage) String() string { return proto.CompactTextString(m) }
func (*ForeignMessage) ProtoMessage()    {}
func (*ForeignMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_54bcf749d548b8e3, []int{2}
}

func (m *ForeignMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForeignMessage.Unmarshal(m, b)
}
func (m *ForeignMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForeignMessage.Marshal(b, m, deterministic)
}
func (m *ForeignMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForeignMessage.Merge(m, src)
}
func (m *ForeignMessage) XXX_Size() int {
	return xxx_messageInfo_ForeignMessage.Size(m)
}
func (m *ForeignMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ForeignMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ForeignMessage proto.InternalMessageInfo

func (m *ForeignMessage) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

type TestEmptyMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestEmptyMessage) Reset()         { *m = TestEmptyMessage{} }
func (m *TestEmptyMessage) String() string { return proto.CompactTextString(m) }
func (*TestEmptyMessage) ProtoMessage()    {}
func (*TestEmptyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_54bcf749d548b8e3, []int{3}
}

func (m *TestEmptyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestEmptyMessage.Unmarshal(m, b)
}
func (m *TestEmptyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestEmptyMessage.Marshal(b, m, deterministic)
}
func (m *TestEmptyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestEmptyMessage.Merge(m, src)
}
func (m *TestEmptyMessage) XXX_Size() int {
	return xxx_messageInfo_TestEmptyMessage.Size(m)
}
func (m *TestEmptyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestEmptyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestEmptyMessage proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("tensorflow.test.ForeignEnum", ForeignEnum_name, ForeignEnum_value)
	proto.RegisterEnum("tensorflow.test.TestAllTypes_NestedEnum", TestAllTypes_NestedEnum_name, TestAllTypes_NestedEnum_value)
	proto.RegisterType((*TestAllTypes)(nil), "tensorflow.test.TestAllTypes")
	proto.RegisterMapType((map[string]*TestAllTypes_NestedMessage)(nil), "tensorflow.test.TestAllTypes.AnotherMapStringToMessageEntry")
	proto.RegisterMapType((map[bool]*TestAllTypes_NestedMessage)(nil), "tensorflow.test.TestAllTypes.MapBoolToMessageEntry")
	proto.RegisterMapType((map[int32]*TestAllTypes_NestedMessage)(nil), "tensorflow.test.TestAllTypes.MapInt32ToMessageEntry")
	proto.RegisterMapType((map[int64]*TestAllTypes_NestedMessage)(nil), "tensorflow.test.TestAllTypes.MapInt64ToMessageEntry")
	proto.RegisterMapType((map[int64]string)(nil), "tensorflow.test.TestAllTypes.MapInt64ToStringEntry")
	proto.RegisterMapType((map[string]int64)(nil), "tensorflow.test.TestAllTypes.MapStringToInt64Entry")
	proto.RegisterMapType((map[string]*TestAllTypes_NestedMessage)(nil), "tensorflow.test.TestAllTypes.MapStringToMessageEntry")
	proto.RegisterType((*TestAllTypes_NestedMessage)(nil), "tensorflow.test.TestAllTypes.NestedMessage")
	proto.RegisterType((*TestAllTypes_NestedMessage_DoubleNestedMessage)(nil), "tensorflow.test.TestAllTypes.NestedMessage.DoubleNestedMessage")
	proto.RegisterType((*NestedTestAllTypes)(nil), "tensorflow.test.NestedTestAllTypes")
	proto.RegisterMapType((map[string]int64)(nil), "tensorflow.test.NestedTestAllTypes.MapStringToInt64Entry")
	proto.RegisterType((*ForeignMessage)(nil), "tensorflow.test.ForeignMessage")
	proto.RegisterType((*TestEmptyMessage)(nil), "tensorflow.test.TestEmptyMessage")
}

func init() {
	proto.RegisterFile("tensorflow/tools/proto_text/test.proto", fileDescriptor_54bcf749d548b8e3)
}

var fileDescriptor_54bcf749d548b8e3 = []byte{
	// 1361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x58, 0xef, 0x76, 0xd3, 0x36,
	0x1c, 0xad, 0xa3, 0xfe, 0x55, 0xd2, 0xd6, 0xa8, 0x2d, 0x98, 0x9e, 0x01, 0xa2, 0x1d, 0x60, 0x60,
	0x24, 0x3b, 0xad, 0x67, 0x06, 0x63, 0xb0, 0x96, 0xb5, 0x94, 0x9d, 0x03, 0xdd, 0x0c, 0x7c, 0x58,
	0xbf, 0xe4, 0x38, 0x89, 0x9a, 0xe6, 0xe0, 0x58, 0xc1, 0x71, 0x37, 0xf2, 0x65, 0xaf, 0xb1, 0x47,
	0xd9, 0x43, 0xec, 0x41, 0xb6, 0xb7, 0xd8, 0x8e, 0x24, 0x4b, 0x96, 0x1c, 0xb7, 0xb4, 0xe7, 0x94,
	0x7e, 0x4a, 0x6e, 0xae, 0xee, 0xfd, 0x49, 0x3f, 0x5d, 0xc9, 0x2e, 0xbc, 0x9d, 0x92, 0x78, 0x48,
	0x93, 0xc3, 0x88, 0xfe, 0xde, 0x48, 0x29, 0x8d, 0x86, 0x8d, 0x41, 0x42, 0x53, 0xda, 0x4c, 0xc9,
	0xc7, 0xb4, 0x91, 0x92, 0x61, 0x5a, 0xe7, 0xdf, 0xd1, 0x62, 0xce, 0xab, 0x33, 0x78, 0xed, 0xef,
	0x1b, 0xb0, 0xf6, 0x96, 0x0c, 0xd3, 0xad, 0x28, 0x7a, 0x3b, 0x1a, 0x90, 0x21, 0xba, 0x0d, 0x17,
	0xe8, 0x20, 0xed, 0xd1, 0x38, 0x8c, 0x9a, 0xbd, 0x38, 0xdd, 0xdc, 0x70, 0xfe, 0x9d, 0xc1, 0x96,
	0x3b, 0x15, 0xcc, 0x4b, 0xf8, 0x25, 0x43, 0xd1, 0x2d, 0x93, 0xe7, 0x7b, 0x4e, 0x05, 0x5b, 0x2e,
	0x30, 0x68, 0xbe, 0x87, 0xee, 0xc0, 0x45, 0x45, 0x3b, 0x16, 0x7a, 0x00, 0x5b, 0xee, 0x7c, 0xa0,
	0x46, 0xbf, 0xe3, 0x28, 0x72, 0x0b, 0x44, 0xdf, 0x73, 0xfe, 0x61, 0xc6, 0x93, 0x26, 0xb3, 0x20,
	0x39, 0x14, 0x92, 0x53, 0xd8, 0x72, 0x2f, 0xe5, 0xc4, 0x37, 0x42, 0xb2, 0x48, 0xf4, 0x3d, 0x67,
	0x1a, 0x5b, 0x2e, 0x32, 0x89, 0xbe, 0x87, 0xee, 0x42, 0x5b, 0x11, 0x0f, 0x7b, 0x1f, 0x49, 0x67,
	0x73, 0xc3, 0x61, 0xde, 0x33, 0x81, 0x12, 0xd8, 0x15, 0xf0, 0x38, 0xd5, 0xf7, 0x9c, 0x59, 0x6c,
	0xb9, 0xd3, 0x05, 0xaa, 0xef, 0xa1, 0xfb, 0xf0, 0x52, 0x6e, 0x2f, 0x65, 0xe7, 0xb0, 0xe5, 0x2e,
	0x06, 0x4a, 0xe3, 0x4d, 0x86, 0x97, 0x90, 0x7d, 0xcf, 0x81, 0xd8, 0x72, 0xed, 0x22, 0xd9, 0xf7,
	0x8c, 0xb5, 0x3f, 0x8c, 0x68, 0x98, 0x3a, 0x55, 0x6c, 0xb9, 0x95, 0x7c, 0xed, 0x77, 0x19, 0x68,
	0xcc, 0xbf, 0x43, 0x8f, 0x5b, 0x11, 0x71, 0x6a, 0xd8, 0x72, 0xad, 0x7c, 0xfe, 0x3f, 0x72, 0x14,
	0xad, 0x43, 0x35, 0xb2, 0xd9, 0xa2, 0x34, 0x72, 0xe6, 0xb1, 0xe5, 0xce, 0x06, 0x35, 0x09, 0x6e,
	0x53, 0x1a, 0x99, 0xab, 0x99, 0x26, 0xbd, 0xb8, 0xeb, 0x2c, 0x60, 0xcb, 0x9d, 0xd3, 0x56, 0x93,
	0xa3, 0x46, 0x75, 0xad, 0x51, 0x4a, 0x86, 0xce, 0x22, 0xb6, 0xdc, 0x5a, 0x5e, 0xdd, 0x36, 0x03,
	0x51, 0x1b, 0x5e, 0x51, 0xb4, 0x98, 0x0c, 0x53, 0xd2, 0x69, 0xf6, 0xc9, 0x70, 0x18, 0x76, 0x89,
	0x83, 0xb0, 0xe5, 0x56, 0x37, 0xee, 0xd7, 0x0b, 0x9b, 0xb5, 0xae, 0x6f, 0xd4, 0xfa, 0x6b, 0x3e,
	0xe6, 0x95, 0x18, 0x12, 0xac, 0x48, 0x2d, 0x03, 0x46, 0xbf, 0x42, 0x27, 0x5f, 0x29, 0x9a, 0x90,
	0x5e, 0x37, 0x56, 0x2e, 0x4b, 0xdc, 0xe5, 0xc6, 0x98, 0xcb, 0xae, 0xe0, 0x49, 0xe5, 0xcb, 0x6a,
	0x51, 0x0d, 0x1c, 0x1d, 0xc0, 0xe5, 0x62, 0xfd, 0x24, 0x3e, 0xee, 0x3b, 0x2b, 0xd8, 0x72, 0x17,
	0x36, 0xdc, 0xb3, 0x14, 0xbf, 0x13, 0x1f, 0xf7, 0x03, 0x64, 0x56, 0xce, 0x30, 0xf4, 0x33, 0x5c,
	0x19, 0x2b, 0x9b, 0x8b, 0x5f, 0xe6, 0xe2, 0x5f, 0x9c, 0x54, 0x33, 0x17, 0x5c, 0x2a, 0x14, 0xcc,
	0x15, 0xf5, 0x16, 0xb7, 0x69, 0xd2, 0x71, 0xae, 0xf2, 0xde, 0xa9, 0x16, 0x3f, 0xa7, 0x49, 0x87,
	0x75, 0x2e, 0x21, 0x03, 0x12, 0xb2, 0xb9, 0x88, 0x60, 0xdd, 0xc0, 0x80, 0x45, 0x5f, 0xa2, 0x2a,
	0xfa, 0x3a, 0xcd, 0xf7, 0x1c, 0x8c, 0x01, 0x8b, 0xbe, 0x46, 0x13, 0x39, 0x55, 0xb4, 0x2c, 0xfa,
	0x37, 0x31, 0x60, 0xd1, 0x97, 0xf0, 0x3b, 0x95, 0x53, 0x83, 0xe8, 0x7b, 0xce, 0x1a, 0x06, 0x2c,
	0xf9, 0x3a, 0xb1, 0xa0, 0x98, 0x25, 0x7f, 0x1d, 0x03, 0x96, 0x7c, 0x09, 0xbf, 0x19, 0x57, 0xcc,
	0x92, 0xff, 0x25, 0x06, 0x2c, 0xf9, 0x3a, 0x51, 0x24, 0x5f, 0x11, 0x65, 0x44, 0x6f, 0x61, 0xc0,
	0x92, 0x2f, 0x71, 0x2d, 0xf9, 0x26, 0xd5, 0xf7, 0x9c, 0xdb, 0x18, 0xb0, 0xe4, 0x1b, 0x54, 0x91,
	0xfc, 0xdc, 0x5e, 0xca, 0xde, 0xc1, 0x80, 0x25, 0x5f, 0x15, 0xa0, 0x25, 0xbf, 0x40, 0xf6, 0x3d,
	0xc7, 0xc5, 0x80, 0x25, 0xdf, 0x24, 0x8b, 0xe4, 0xe7, 0x45, 0xf0, 0xe4, 0xdf, 0xc5, 0x80, 0x25,
	0x5f, 0x95, 0x20, 0x93, 0xaf, 0x68, 0x59, 0xf2, 0xef, 0x61, 0xc0, 0x92, 0x2f, 0xe1, 0x3c, 0xf9,
	0x8a, 0xc8, 0x93, 0x7f, 0x1f, 0x03, 0x96, 0x7c, 0x09, 0xca, 0xe4, 0xe7, 0x15, 0x8a, 0xe4, 0x7f,
	0x85, 0x01, 0x4b, 0xbe, 0xaa, 0x4f, 0x25, 0x3f, 0x57, 0xe3, 0xc9, 0x7f, 0x80, 0x01, 0x4b, 0xbe,
	0x92, 0x93, 0xc9, 0x57, 0xb4, 0x42, 0xf2, 0xbf, 0xc6, 0xe0, 0xdc, 0xc9, 0x97, 0x5a, 0x66, 0xf2,
	0x0f, 0xe0, 0x72, 0xd1, 0x84, 0x27, 0x68, 0x13, 0x83, 0xf3, 0xc5, 0xd3, 0x94, 0x97, 0x61, 0x52,
	0xda, 0x3c, 0x4c, 0x0f, 0xf9, 0x72, 0xa8, 0x55, 0xe3, 0x61, 0x5a, 0x87, 0x35, 0x1a, 0x13, 0x7a,
	0x28, 0xf7, 0x3e, 0x65, 0xd7, 0xde, 0xde, 0x44, 0x50, 0xe5, 0x68, 0xb6, 0xf5, 0x9b, 0x70, 0x59,
	0x90, 0x0a, 0xeb, 0x30, 0x38, 0xf7, 0x09, 0xb8, 0x37, 0x11, 0x20, 0x2e, 0x65, 0x2e, 0x83, 0xaa,
	0x22, 0x6b, 0xdc, 0x07, 0x16, 0x7b, 0x55, 0x45, 0xd6, 0xb7, 0x9b, 0x50, 0x7c, 0xcd, 0x9a, 0x96,
	0xb0, 0xe3, 0x7a, 0x6f, 0x22, 0x80, 0x1c, 0x14, 0x3d, 0x7b, 0x09, 0xc5, 0x37, 0xb1, 0x88, 0x9d,
	0xf3, 0x9d, 0x71, 0x7b, 0x13, 0xc1, 0x1c, 0x1f, 0xcd, 0x57, 0xef, 0x08, 0xae, 0xf4, 0xc3, 0x41,
	0x56, 0x50, 0x33, 0xa5, 0x6a, 0xd2, 0x8f, 0x79, 0xf3, 0xbf, 0x39, 0x5d, 0xf5, 0x55, 0x38, 0x10,
	0x55, 0xbf, 0xa5, 0xd9, 0x1c, 0x77, 0xe2, 0x34, 0x19, 0x05, 0xa8, 0x3f, 0xf6, 0x03, 0x22, 0x70,
	0x99, 0x39, 0xf1, 0xa5, 0xd6, 0x8d, 0xbe, 0xe3, 0x46, 0xde, 0x27, 0x8d, 0xf8, 0x71, 0x57, 0xf0,
	0xb9, 0xd4, 0x2f, 0xe2, 0x9a, 0x8d, 0xef, 0xe9, 0x36, 0x4f, 0xce, 0x6e, 0xe3, 0x7b, 0xe5, 0x36,
	0x1a, 0x8e, 0x5a, 0x70, 0x89, 0xd9, 0xb0, 0x98, 0xea, 0x2e, 0xdf, 0x73, 0x97, 0xcd, 0x4f, 0xba,
	0xb0, 0x28, 0x17, 0x4c, 0xec, 0x7e, 0x01, 0x96, 0x1e, 0x79, 0x6f, 0xc4, 0xe1, 0xf9, 0xf4, 0x8c,
	0x1e, 0xb2, 0x01, 0xbc, 0xf2, 0xdc, 0xc3, 0x80, 0xa5, 0x87, 0x5a, 0xae, 0x6c, 0x67, 0x3e, 0x3b,
	0xa3, 0x47, 0xb6, 0x2a, 0x42, 0x33, 0xf7, 0x30, 0x60, 0xf4, 0x07, 0xbc, 0x16, 0xc6, 0x34, 0x3d,
	0x22, 0x49, 0xb3, 0x7c, 0xaf, 0x6d, 0x71, 0xb7, 0x27, 0xa7, 0xbb, 0x6d, 0x09, 0x89, 0x93, 0xb6,
	0xdc, 0xd5, 0xf0, 0xa4, 0xdf, 0x91, 0x0f, 0x57, 0x06, 0x61, 0xfb, 0x3d, 0xe9, 0x34, 0x0b, 0x37,
	0xe5, 0x0f, 0xec, 0xa6, 0xdc, 0xae, 0xd8, 0x56, 0xb0, 0x24, 0x08, 0x81, 0x7e, 0x67, 0xae, 0xfe,
	0x59, 0x81, 0xf3, 0x66, 0x80, 0x6f, 0x8d, 0x3d, 0x8f, 0x5b, 0x27, 0x3c, 0x8e, 0x17, 0xae, 0xee,
	0x4a, 0xd9, 0xd5, 0xfd, 0x0b, 0x04, 0xfd, 0x61, 0x97, 0x3f, 0x82, 0x57, 0x37, 0x9e, 0x9d, 0xe3,
	0x78, 0xa9, 0x8b, 0x0b, 0xc3, 0x3c, 0x7a, 0x99, 0x56, 0xc9, 0x8b, 0xc0, 0x64, 0xc9, 0x8b, 0xc0,
	0xea, 0x53, 0xb8, 0x54, 0x22, 0x51, 0xf6, 0x54, 0x69, 0x95, 0x3d, 0x55, 0xae, 0x26, 0xf0, 0xca,
	0x09, 0x7d, 0x40, 0x36, 0x04, 0xef, 0xc9, 0x28, 0x1b, 0xc7, 0x3e, 0xa2, 0x2d, 0x38, 0xf5, 0x5b,
	0x18, 0x1d, 0x13, 0xfe, 0x4e, 0x72, 0xce, 0xfb, 0x44, 0x8c, 0x7c, 0x5c, 0xf9, 0xd6, 0x5a, 0xfd,
	0x00, 0x2f, 0x97, 0x9f, 0x02, 0xba, 0xe5, 0xd4, 0xe7, 0xb1, 0x2c, 0x9e, 0x08, 0xba, 0x25, 0xb8,
	0x50, 0xcb, 0x01, 0x5c, 0x29, 0x3d, 0x1e, 0x74, 0xc7, 0xd9, 0x0b, 0x75, 0x7c, 0xce, 0x1d, 0xc7,
	0x0f, 0x8b, 0x92, 0x4e, 0x2e, 0xeb, 0x8e, 0x60, 0x5c, 0x64, 0xfc, 0x34, 0x28, 0x59, 0x28, 0x43,
	0x64, 0x4e, 0x17, 0x19, 0xc1, 0xeb, 0xa7, 0x87, 0xfc, 0xb3, 0x6d, 0xae, 0xb5, 0xe7, 0x10, 0x6a,
	0x8f, 0x14, 0xb3, 0x70, 0xf2, 0x60, 0x27, 0xd8, 0xb7, 0x27, 0xd0, 0x0c, 0x04, 0xbb, 0xfb, 0xfb,
	0xb6, 0xc5, 0x3e, 0x6c, 0x6f, 0x05, 0x76, 0x45, 0x7c, 0x38, 0xb0, 0x01, 0xab, 0xe5, 0xf5, 0xce,
	0x0b, 0xfb, 0x3f, 0xf9, 0x67, 0x6d, 0xcf, 0xcb, 0x9b, 0xfb, 0xb0, 0x47, 0xa2, 0xce, 0xda, 0x5f,
	0x15, 0x88, 0x84, 0xa8, 0xf1, 0x4e, 0xff, 0x08, 0x4e, 0xb5, 0x8f, 0x7a, 0x51, 0x87, 0xcf, 0xa2,
	0xba, 0xb1, 0x3e, 0x56, 0xf1, 0xf8, 0x98, 0x40, 0x8c, 0x40, 0x0f, 0xe1, 0xcc, 0x20, 0x1c, 0x45,
	0x34, 0xec, 0x64, 0xd3, 0xbd, 0x76, 0xea, 0x74, 0x03, 0xc9, 0x46, 0x47, 0xe5, 0x37, 0x09, 0xe0,
	0xe7, 0xee, 0xa3, 0x33, 0x54, 0x70, 0xd6, 0xfb, 0xe4, 0x42, 0x76, 0xd3, 0xda, 0x75, 0xb8, 0x50,
	0x78, 0xbf, 0xab, 0x41, 0xab, 0x9d, 0x05, 0xdc, 0x6a, 0xaf, 0x21, 0x68, 0xb3, 0xe2, 0x76, 0xfa,
	0x83, 0x74, 0x94, 0x31, 0xee, 0x05, 0xb0, 0xaa, 0xbf, 0x62, 0xd9, 0xb0, 0xb6, 0xbb, 0x1f, 0xec,
	0xbc, 0x7c, 0xf1, 0xba, 0x99, 0xb5, 0x72, 0x11, 0x56, 0x25, 0xc2, 0x5a, 0x3a, 0xa9, 0x03, 0xac,
	0xb5, 0x53, 0x26, 0x70, 0x60, 0x4f, 0x6f, 0xff, 0x74, 0xb0, 0xd7, 0xed, 0xa5, 0x47, 0xc7, 0xad,
	0x7a, 0x9b, 0xf6, 0x1b, 0x31, 0x49, 0x5b, 0x49, 0xd8, 0x8b, 0x1b, 0xe9, 0xe1, 0x83, 0x6e, 0x32,
	0x68, 0x37, 0xba, 0xb4, 0xa1, 0xff, 0xa7, 0x27, 0xff, 0xc8, 0x7e, 0x28, 0xfc, 0xdf, 0xa7, 0x35,
	0xcd, 0x3f, 0x6f, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x53, 0x72, 0x31, 0x4c, 0x1d, 0x12, 0x00,
	0x00,
}