// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/lite/toco/model_flags.proto

package toco

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

type InputArrayShape struct {
	Dims                 []int32  `protobuf:"varint,2,rep,name=dims" json:"dims,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputArrayShape) Reset()         { *m = InputArrayShape{} }
func (m *InputArrayShape) String() string { return proto.CompactTextString(m) }
func (*InputArrayShape) ProtoMessage()    {}
func (*InputArrayShape) Descriptor() ([]byte, []int) {
	return fileDescriptor_53fa85a52e940e82, []int{0}
}

func (m *InputArrayShape) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputArrayShape.Unmarshal(m, b)
}
func (m *InputArrayShape) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputArrayShape.Marshal(b, m, deterministic)
}
func (m *InputArrayShape) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputArrayShape.Merge(m, src)
}
func (m *InputArrayShape) XXX_Size() int {
	return xxx_messageInfo_InputArrayShape.Size(m)
}
func (m *InputArrayShape) XXX_DiscardUnknown() {
	xxx_messageInfo_InputArrayShape.DiscardUnknown(m)
}

var xxx_messageInfo_InputArrayShape proto.InternalMessageInfo

func (m *InputArrayShape) GetDims() []int32 {
	if m != nil {
		return m.Dims
	}
	return nil
}

// Next ID to USE: 7.
type InputArray struct {
	// Name of the input arrays, i.e. the arrays from which input activations
	// will be read.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Shape of the input.  For many applications the dimensions are {batch,
	// height, width, depth}.  Often the batch is left "unspecified" by providing
	// a value of -1.
	//
	// The last dimension is typically called 'depth' or 'channels'. For example,
	// for an image model taking RGB images as input, this would have the value 3.
	Shape *InputArrayShape `protobuf:"bytes,6,opt,name=shape" json:"shape,omitempty"`
	// mean_value and std_value parameters control the interpretation of raw input
	// activation values (elements of the input array) as real numbers. The
	// mapping is given by:
	//
	//    real_value = (raw_input_value - mean_value) / std_value
	//
	// In particular, the defaults (mean_value=0, std_value=1) yield
	// real_value = raw_input_value. Often, non-default values are used in image
	// models. For example, an image model taking uint8 image channel values as
	// its raw inputs, in [0, 255] range, may use mean_value=128, std_value=128 to
	// map them into the interval [-1, 1).
	//
	// Note: this matches exactly the meaning of mean_value and std_value in
	// (TensorFlow via LegacyFedInput).
	MeanValue *float32 `protobuf:"fixed32,3,opt,name=mean_value,json=meanValue" json:"mean_value,omitempty"`
	StdValue  *float32 `protobuf:"fixed32,4,opt,name=std_value,json=stdValue,def=1" json:"std_value,omitempty"`
	// Data type of the input.
	//
	// In many graphs, the input arrays already have defined data types,
	// e.g. Placeholder nodes in a TensorFlow GraphDef have a dtype attribute.
	// In those cases, it is not needed to specify this data_type flag.
	// The purpose of this flag is only to define the data type of input
	// arrays whose type isn't defined in the input graph file. For example,
	// when specifying an arbitrary (not Placeholder) --input_array into
	// a TensorFlow GraphDef.
	//
	// When this data_type is quantized (e.g. QUANTIZED_UINT8), the
	// corresponding quantization parameters are the mean_value, std_value
	// fields.
	//
	// It is also important to understand the nuance between this data_type
	// flag and the inference_input_type in TocoFlags. The basic difference
	// is that this data_type (like all ModelFlags) describes a property
	// of the input graph, while inference_input_type (like all TocoFlags)
	// describes an aspect of the toco transformation process and thus of
	// the output file. The types of input arrays may be different between
	// the input and output files if quantization or dequantization occurred.
	// Such differences can only occur for real-number data i.e. only
	// between FLOAT and quantized types (e.g. QUANTIZED_UINT8).
	DataType             *IODataType `protobuf:"varint,5,opt,name=data_type,json=dataType,enum=toco.IODataType" json:"data_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *InputArray) Reset()         { *m = InputArray{} }
func (m *InputArray) String() string { return proto.CompactTextString(m) }
func (*InputArray) ProtoMessage()    {}
func (*InputArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_53fa85a52e940e82, []int{1}
}

func (m *InputArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputArray.Unmarshal(m, b)
}
func (m *InputArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputArray.Marshal(b, m, deterministic)
}
func (m *InputArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputArray.Merge(m, src)
}
func (m *InputArray) XXX_Size() int {
	return xxx_messageInfo_InputArray.Size(m)
}
func (m *InputArray) XXX_DiscardUnknown() {
	xxx_messageInfo_InputArray.DiscardUnknown(m)
}

var xxx_messageInfo_InputArray proto.InternalMessageInfo

const Default_InputArray_StdValue float32 = 1

func (m *InputArray) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *InputArray) GetShape() *InputArrayShape {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *InputArray) GetMeanValue() float32 {
	if m != nil && m.MeanValue != nil {
		return *m.MeanValue
	}
	return 0
}

func (m *InputArray) GetStdValue() float32 {
	if m != nil && m.StdValue != nil {
		return *m.StdValue
	}
	return Default_InputArray_StdValue
}

func (m *InputArray) GetDataType() IODataType {
	if m != nil && m.DataType != nil {
		return *m.DataType
	}
	return IODataType_IO_DATA_TYPE_UNKNOWN
}

type RnnState struct {
	StateArray          *string `protobuf:"bytes,1,opt,name=state_array,json=stateArray" json:"state_array,omitempty"`
	BackEdgeSourceArray *string `protobuf:"bytes,2,opt,name=back_edge_source_array,json=backEdgeSourceArray" json:"back_edge_source_array,omitempty"`
	Discardable         *bool   `protobuf:"varint,5,opt,name=discardable" json:"discardable,omitempty"`
	// size allows to specify a 1-D shape for the RNN state array.
	// Will be expanded with 1's to fit the model.
	// TODO(benoitjacob): should allow a generic, explicit shape.
	Size                 *int32   `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RnnState) Reset()         { *m = RnnState{} }
func (m *RnnState) String() string { return proto.CompactTextString(m) }
func (*RnnState) ProtoMessage()    {}
func (*RnnState) Descriptor() ([]byte, []int) {
	return fileDescriptor_53fa85a52e940e82, []int{2}
}

func (m *RnnState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RnnState.Unmarshal(m, b)
}
func (m *RnnState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RnnState.Marshal(b, m, deterministic)
}
func (m *RnnState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RnnState.Merge(m, src)
}
func (m *RnnState) XXX_Size() int {
	return xxx_messageInfo_RnnState.Size(m)
}
func (m *RnnState) XXX_DiscardUnknown() {
	xxx_messageInfo_RnnState.DiscardUnknown(m)
}

var xxx_messageInfo_RnnState proto.InternalMessageInfo

func (m *RnnState) GetStateArray() string {
	if m != nil && m.StateArray != nil {
		return *m.StateArray
	}
	return ""
}

func (m *RnnState) GetBackEdgeSourceArray() string {
	if m != nil && m.BackEdgeSourceArray != nil {
		return *m.BackEdgeSourceArray
	}
	return ""
}

func (m *RnnState) GetDiscardable() bool {
	if m != nil && m.Discardable != nil {
		return *m.Discardable
	}
	return false
}

func (m *RnnState) GetSize() int32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

// An ArraysExtraInfo message stores a collection of additional Information
// about arrays in a model, complementing the information in the model itself.
// It is intentionally a separate message so that it may be serialized and
// passed separately from the model. See --arrays_extra_info_file.
//
// A typical use case is to manually specify MinMax for specific arrays in a
// model that does not already contain such MinMax information.
type ArraysExtraInfo struct {
	Entries              []*ArraysExtraInfo_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ArraysExtraInfo) Reset()         { *m = ArraysExtraInfo{} }
func (m *ArraysExtraInfo) String() string { return proto.CompactTextString(m) }
func (*ArraysExtraInfo) ProtoMessage()    {}
func (*ArraysExtraInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_53fa85a52e940e82, []int{3}
}

func (m *ArraysExtraInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArraysExtraInfo.Unmarshal(m, b)
}
func (m *ArraysExtraInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArraysExtraInfo.Marshal(b, m, deterministic)
}
func (m *ArraysExtraInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArraysExtraInfo.Merge(m, src)
}
func (m *ArraysExtraInfo) XXX_Size() int {
	return xxx_messageInfo_ArraysExtraInfo.Size(m)
}
func (m *ArraysExtraInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ArraysExtraInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ArraysExtraInfo proto.InternalMessageInfo

func (m *ArraysExtraInfo) GetEntries() []*ArraysExtraInfo_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type ArraysExtraInfo_Entry struct {
	// Next ID to use: 8.
	Name                 *string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NameRegexp           *string          `protobuf:"bytes,7,opt,name=name_regexp,json=nameRegexp" json:"name_regexp,omitempty"`
	Min                  *float64         `protobuf:"fixed64,2,opt,name=min" json:"min,omitempty"`
	Max                  *float64         `protobuf:"fixed64,3,opt,name=max" json:"max,omitempty"`
	DataType             *IODataType      `protobuf:"varint,4,opt,name=data_type,json=dataType,enum=toco.IODataType" json:"data_type,omitempty"`
	Shape                *InputArrayShape `protobuf:"bytes,5,opt,name=shape" json:"shape,omitempty"`
	ConstantFloatValue   *float32         `protobuf:"fixed32,6,opt,name=constant_float_value,json=constantFloatValue" json:"constant_float_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ArraysExtraInfo_Entry) Reset()         { *m = ArraysExtraInfo_Entry{} }
func (m *ArraysExtraInfo_Entry) String() string { return proto.CompactTextString(m) }
func (*ArraysExtraInfo_Entry) ProtoMessage()    {}
func (*ArraysExtraInfo_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_53fa85a52e940e82, []int{3, 0}
}

func (m *ArraysExtraInfo_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArraysExtraInfo_Entry.Unmarshal(m, b)
}
func (m *ArraysExtraInfo_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArraysExtraInfo_Entry.Marshal(b, m, deterministic)
}
func (m *ArraysExtraInfo_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArraysExtraInfo_Entry.Merge(m, src)
}
func (m *ArraysExtraInfo_Entry) XXX_Size() int {
	return xxx_messageInfo_ArraysExtraInfo_Entry.Size(m)
}
func (m *ArraysExtraInfo_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_ArraysExtraInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_ArraysExtraInfo_Entry proto.InternalMessageInfo

func (m *ArraysExtraInfo_Entry) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ArraysExtraInfo_Entry) GetNameRegexp() string {
	if m != nil && m.NameRegexp != nil {
		return *m.NameRegexp
	}
	return ""
}

func (m *ArraysExtraInfo_Entry) GetMin() float64 {
	if m != nil && m.Min != nil {
		return *m.Min
	}
	return 0
}

func (m *ArraysExtraInfo_Entry) GetMax() float64 {
	if m != nil && m.Max != nil {
		return *m.Max
	}
	return 0
}

func (m *ArraysExtraInfo_Entry) GetDataType() IODataType {
	if m != nil && m.DataType != nil {
		return *m.DataType
	}
	return IODataType_IO_DATA_TYPE_UNKNOWN
}

func (m *ArraysExtraInfo_Entry) GetShape() *InputArrayShape {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *ArraysExtraInfo_Entry) GetConstantFloatValue() float32 {
	if m != nil && m.ConstantFloatValue != nil {
		return *m.ConstantFloatValue
	}
	return 0
}

// ModelFlags encodes properties of a model that, depending on the file
// format, may or may not be recorded in the model file. The purpose of
// representing these properties in ModelFlags is to allow passing them
// separately from the input model file, for instance as command-line
// parameters, so that we can offer a single uniform interface that can
// handle files from different input formats.
//
// For each of these properties, and each supported file format, we
// detail in comments below whether the property exists in the given file
// format.
//
// Obsolete flags that have been removed:
//   optional int32 input_depth = 3;
//   optional int32 input_width = 4;
//   optional int32 input_height = 5;
//   optional int32 batch = 6 [ default = 1];
//   optional float mean_value = 7;
//   optional float std_value = 8 [default = 1.];
//   optional int32 input_dims = 11 [ default = 4];
//   repeated int32 input_shape = 13;
//
// Next ID to USE: 20.
type ModelFlags struct {
	// Information about the input arrays, i.e. the arrays from which input
	// activations will be read.
	InputArrays []*InputArray `protobuf:"bytes,1,rep,name=input_arrays,json=inputArrays" json:"input_arrays,omitempty"`
	// Name of the output arrays, i.e. the arrays into which output activations
	// will be written.
	OutputArrays []string `protobuf:"bytes,2,rep,name=output_arrays,json=outputArrays" json:"output_arrays,omitempty"`
	// If true, the model accepts an arbitrary batch size. Mutually exclusive with
	// the 'batch' field: at most one of these two fields can be set.
	VariableBatch *bool                    `protobuf:"varint,10,opt,name=variable_batch,json=variableBatch" json:"variable_batch,omitempty"`
	RnnStates     []*RnnState              `protobuf:"bytes,12,rep,name=rnn_states,json=rnnStates" json:"rnn_states,omitempty"`
	ModelChecks   []*ModelFlags_ModelCheck `protobuf:"bytes,14,rep,name=model_checks,json=modelChecks" json:"model_checks,omitempty"`
	// If true, will allow passing inexistent arrays in --input_arrays
	// and --output_arrays. This makes little sense, is only useful to
	// more easily get graph visualizations.
	AllowNonexistentArrays *bool `protobuf:"varint,16,opt,name=allow_nonexistent_arrays,json=allowNonexistentArrays" json:"allow_nonexistent_arrays,omitempty"`
	// If true, will allow passing non-ascii-printable characters in
	// --input_arrays and --output_arrays. By default (if false), only
	// ascii printable characters are allowed, i.e. character codes
	// ranging from 32 to 127. This is disallowed by default so as to
	// catch common copy-and-paste issues where invisible unicode
	// characters are unwittingly added to these strings.
	AllowNonasciiArrays *bool `protobuf:"varint,17,opt,name=allow_nonascii_arrays,json=allowNonasciiArrays" json:"allow_nonascii_arrays,omitempty"`
	// If set, this ArraysExtraInfo allows to pass extra information about arrays
	// not specified in the input model file, such as extra MinMax information.
	ArraysExtraInfo *ArraysExtraInfo `protobuf:"bytes,18,opt,name=arrays_extra_info,json=arraysExtraInfo" json:"arrays_extra_info,omitempty"`
	// When set to false, toco will not change the input ranges and the output
	// ranges of concat operator to the overlap of all input ranges.
	ChangeConcatInputRanges *bool    `protobuf:"varint,19,opt,name=change_concat_input_ranges,json=changeConcatInputRanges,def=1" json:"change_concat_input_ranges,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ModelFlags) Reset()         { *m = ModelFlags{} }
func (m *ModelFlags) String() string { return proto.CompactTextString(m) }
func (*ModelFlags) ProtoMessage()    {}
func (*ModelFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_53fa85a52e940e82, []int{4}
}

func (m *ModelFlags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelFlags.Unmarshal(m, b)
}
func (m *ModelFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelFlags.Marshal(b, m, deterministic)
}
func (m *ModelFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelFlags.Merge(m, src)
}
func (m *ModelFlags) XXX_Size() int {
	return xxx_messageInfo_ModelFlags.Size(m)
}
func (m *ModelFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelFlags.DiscardUnknown(m)
}

var xxx_messageInfo_ModelFlags proto.InternalMessageInfo

const Default_ModelFlags_ChangeConcatInputRanges bool = true

func (m *ModelFlags) GetInputArrays() []*InputArray {
	if m != nil {
		return m.InputArrays
	}
	return nil
}

func (m *ModelFlags) GetOutputArrays() []string {
	if m != nil {
		return m.OutputArrays
	}
	return nil
}

func (m *ModelFlags) GetVariableBatch() bool {
	if m != nil && m.VariableBatch != nil {
		return *m.VariableBatch
	}
	return false
}

func (m *ModelFlags) GetRnnStates() []*RnnState {
	if m != nil {
		return m.RnnStates
	}
	return nil
}

func (m *ModelFlags) GetModelChecks() []*ModelFlags_ModelCheck {
	if m != nil {
		return m.ModelChecks
	}
	return nil
}

func (m *ModelFlags) GetAllowNonexistentArrays() bool {
	if m != nil && m.AllowNonexistentArrays != nil {
		return *m.AllowNonexistentArrays
	}
	return false
}

func (m *ModelFlags) GetAllowNonasciiArrays() bool {
	if m != nil && m.AllowNonasciiArrays != nil {
		return *m.AllowNonasciiArrays
	}
	return false
}

func (m *ModelFlags) GetArraysExtraInfo() *ArraysExtraInfo {
	if m != nil {
		return m.ArraysExtraInfo
	}
	return nil
}

func (m *ModelFlags) GetChangeConcatInputRanges() bool {
	if m != nil && m.ChangeConcatInputRanges != nil {
		return *m.ChangeConcatInputRanges
	}
	return Default_ModelFlags_ChangeConcatInputRanges
}

// Checks applied to the model, typically after toco's comprehensive
// graph transformations.
// Next ID to USE: 4.
type ModelFlags_ModelCheck struct {
	// Use the name of a type of operator to check its counts.
	// Use "Total" for overall operator counts.
	// Use "Arrays" for overall array counts.
	CountType *string `protobuf:"bytes,1,opt,name=count_type,json=countType,def=None" json:"count_type,omitempty"`
	// A count of zero is a meaningful check, so negative used to mean disable.
	CountMin *int32 `protobuf:"varint,2,opt,name=count_min,json=countMin,def=-1" json:"count_min,omitempty"`
	// If count_max < count_min, then count_min is only allowed value.
	CountMax             *int32   `protobuf:"varint,3,opt,name=count_max,json=countMax,def=-1" json:"count_max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelFlags_ModelCheck) Reset()         { *m = ModelFlags_ModelCheck{} }
func (m *ModelFlags_ModelCheck) String() string { return proto.CompactTextString(m) }
func (*ModelFlags_ModelCheck) ProtoMessage()    {}
func (*ModelFlags_ModelCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_53fa85a52e940e82, []int{4, 0}
}

func (m *ModelFlags_ModelCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelFlags_ModelCheck.Unmarshal(m, b)
}
func (m *ModelFlags_ModelCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelFlags_ModelCheck.Marshal(b, m, deterministic)
}
func (m *ModelFlags_ModelCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelFlags_ModelCheck.Merge(m, src)
}
func (m *ModelFlags_ModelCheck) XXX_Size() int {
	return xxx_messageInfo_ModelFlags_ModelCheck.Size(m)
}
func (m *ModelFlags_ModelCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelFlags_ModelCheck.DiscardUnknown(m)
}

var xxx_messageInfo_ModelFlags_ModelCheck proto.InternalMessageInfo

const Default_ModelFlags_ModelCheck_CountType string = "None"
const Default_ModelFlags_ModelCheck_CountMin int32 = -1
const Default_ModelFlags_ModelCheck_CountMax int32 = -1

func (m *ModelFlags_ModelCheck) GetCountType() string {
	if m != nil && m.CountType != nil {
		return *m.CountType
	}
	return Default_ModelFlags_ModelCheck_CountType
}

func (m *ModelFlags_ModelCheck) GetCountMin() int32 {
	if m != nil && m.CountMin != nil {
		return *m.CountMin
	}
	return Default_ModelFlags_ModelCheck_CountMin
}

func (m *ModelFlags_ModelCheck) GetCountMax() int32 {
	if m != nil && m.CountMax != nil {
		return *m.CountMax
	}
	return Default_ModelFlags_ModelCheck_CountMax
}

func init() {
	proto.RegisterType((*InputArrayShape)(nil), "toco.InputArrayShape")
	proto.RegisterType((*InputArray)(nil), "toco.InputArray")
	proto.RegisterType((*RnnState)(nil), "toco.RnnState")
	proto.RegisterType((*ArraysExtraInfo)(nil), "toco.ArraysExtraInfo")
	proto.RegisterType((*ArraysExtraInfo_Entry)(nil), "toco.ArraysExtraInfo.Entry")
	proto.RegisterType((*ModelFlags)(nil), "toco.ModelFlags")
	proto.RegisterType((*ModelFlags_ModelCheck)(nil), "toco.ModelFlags.ModelCheck")
}

func init() {
	proto.RegisterFile("tensorflow/lite/toco/model_flags.proto", fileDescriptor_53fa85a52e940e82)
}

var fileDescriptor_53fa85a52e940e82 = []byte{
	// 750 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0x32, 0x37,
	0x14, 0xd5, 0xf0, 0xf3, 0x7d, 0x70, 0xe1, 0x23, 0xc4, 0x69, 0xd2, 0x11, 0x55, 0x9b, 0x11, 0x51,
	0x2a, 0xa4, 0x0a, 0x68, 0x88, 0x2a, 0x55, 0x2c, 0x2a, 0x91, 0x34, 0x91, 0xb2, 0x48, 0x2b, 0x39,
	0x55, 0x17, 0xdd, 0x58, 0x66, 0xc6, 0x0c, 0x56, 0x06, 0x1b, 0xd9, 0x26, 0x81, 0x3e, 0x48, 0xdf,
	0xa5, 0xea, 0xf3, 0x74, 0xdb, 0x67, 0xa8, 0x6c, 0xcf, 0x00, 0x41, 0x91, 0xda, 0x15, 0x87, 0x73,
	0xcf, 0x9d, 0xeb, 0xe3, 0x7b, 0x66, 0xe0, 0x6b, 0xc3, 0x84, 0x96, 0x6a, 0x96, 0xc9, 0xd7, 0x61,
	0xc6, 0x0d, 0x1b, 0x1a, 0x19, 0xcb, 0xe1, 0x42, 0x26, 0x2c, 0x23, 0xb3, 0x8c, 0xa6, 0x7a, 0xb0,
	0x54, 0xd2, 0x48, 0x54, 0xb1, 0x7c, 0x27, 0x7a, 0x57, 0x6d, 0x36, 0x4b, 0x96, 0xeb, 0xba, 0x97,
	0x70, 0xf4, 0x20, 0x96, 0x2b, 0x33, 0x51, 0x8a, 0x6e, 0x9e, 0xe6, 0x74, 0xc9, 0x10, 0x82, 0x4a,
	0xc2, 0x17, 0x3a, 0x2c, 0x45, 0xe5, 0x5e, 0x15, 0x3b, 0xdc, 0xfd, 0x2b, 0x00, 0xd8, 0xe9, 0xac,
	0x44, 0xd0, 0x05, 0x0b, 0x83, 0x28, 0xe8, 0xd5, 0xb1, 0xc3, 0xe8, 0x1b, 0xa8, 0x6a, 0xdb, 0x1f,
	0x7e, 0x88, 0x82, 0x5e, 0x63, 0x74, 0x3a, 0xb0, 0xb3, 0x06, 0x07, 0x0f, 0xc7, 0x5e, 0x83, 0xbe,
	0x04, 0x58, 0x30, 0x2a, 0xc8, 0x0b, 0xcd, 0x56, 0x2c, 0x2c, 0x47, 0x41, 0xaf, 0x84, 0xeb, 0x96,
	0xf9, 0xd5, 0x12, 0xe8, 0x2b, 0xa8, 0x6b, 0x93, 0xe4, 0xd5, 0x8a, 0xad, 0x8e, 0x83, 0x2b, 0x5c,
	0xd3, 0x26, 0xf1, 0xf5, 0x3e, 0xd4, 0x13, 0x6a, 0x28, 0xb1, 0x4e, 0xc2, 0x6a, 0x14, 0xf4, 0x5a,
	0xa3, 0x76, 0x3e, 0xef, 0xe7, 0x1f, 0xa9, 0xa1, 0xbf, 0x6c, 0x96, 0x0c, 0xd7, 0x92, 0x1c, 0x75,
	0xff, 0x08, 0xa0, 0x86, 0x85, 0x78, 0x32, 0xd4, 0x30, 0x74, 0x0e, 0x0d, 0x6d, 0x01, 0xa1, 0xf6,
	0x54, 0xb9, 0x05, 0x70, 0x94, 0x37, 0x77, 0x0d, 0x67, 0x53, 0x1a, 0x3f, 0x13, 0x96, 0xa4, 0x8c,
	0x68, 0xb9, 0x52, 0x71, 0xa1, 0x2d, 0x39, 0xed, 0x89, 0xad, 0xde, 0x25, 0x29, 0x7b, 0x72, 0x35,
	0xdf, 0x14, 0x41, 0x23, 0xe1, 0x3a, 0xa6, 0x2a, 0xa1, 0xd3, 0xcc, 0x9f, 0xa9, 0x86, 0xf7, 0x29,
	0x7b, 0x67, 0x9a, 0xff, 0xee, 0xcd, 0x56, 0xb1, 0xc3, 0xdd, 0x3f, 0x4b, 0x70, 0xe4, 0xfa, 0xf5,
	0xdd, 0xda, 0x28, 0xfa, 0x20, 0x66, 0x12, 0x7d, 0x07, 0x1f, 0x99, 0x30, 0x8a, 0x33, 0x1d, 0x06,
	0x51, 0xb9, 0xd7, 0x18, 0x7d, 0xe1, 0x9d, 0x1d, 0xe8, 0x06, 0x77, 0xc2, 0xa8, 0x0d, 0x2e, 0xb4,
	0x9d, 0x7f, 0x02, 0xa8, 0x3a, 0xea, 0xdd, 0xe5, 0x9c, 0x43, 0xc3, 0xfe, 0x12, 0xc5, 0x52, 0xb6,
	0x5e, 0x86, 0x1f, 0xbd, 0x69, 0x4b, 0x61, 0xc7, 0xa0, 0x36, 0x94, 0x17, 0x5c, 0x38, 0x87, 0x01,
	0xb6, 0xd0, 0x31, 0x74, 0xed, 0x8e, 0x6b, 0x19, 0xba, 0x7e, 0x7b, 0xeb, 0x95, 0xff, 0xba, 0xf5,
	0x5d, 0x20, 0xaa, 0xff, 0x23, 0x10, 0xdf, 0xc2, 0x67, 0xb1, 0x14, 0xda, 0x50, 0x61, 0xc8, 0x2c,
	0x93, 0xd4, 0xe4, 0xcb, 0xff, 0xe0, 0xa2, 0x81, 0x8a, 0xda, 0xbd, 0x2d, 0xb9, 0x0c, 0x74, 0xff,
	0xae, 0x00, 0x3c, 0xda, 0xdc, 0xdf, 0xdb, 0xd8, 0xa3, 0x6b, 0x68, 0x72, 0xfb, 0x68, 0xbf, 0xaa,
	0xe2, 0xee, 0xda, 0x87, 0x43, 0x71, 0x83, 0x6f, 0xb1, 0x46, 0x17, 0xf0, 0x49, 0xae, 0xcc, 0x5e,
	0x97, 0xcd, 0x7c, 0x1d, 0x37, 0x3d, 0x99, 0x8b, 0x2e, 0xa1, 0xf5, 0x42, 0x15, 0xb7, 0x4b, 0x24,
	0x53, 0x6a, 0xe2, 0x79, 0x08, 0x6e, 0xbb, 0x9f, 0x0a, 0xf6, 0xc6, 0x92, 0xa8, 0x0f, 0xa0, 0x84,
	0x20, 0x2e, 0x48, 0x3a, 0x6c, 0xba, 0xf1, 0x2d, 0x3f, 0xbe, 0xc8, 0x1e, 0xae, 0xab, 0x1c, 0x69,
	0xf4, 0x03, 0x34, 0xfd, 0x5b, 0x1b, 0xcf, 0x59, 0xfc, 0xac, 0xc3, 0xd6, 0xfe, 0xae, 0x77, 0xbe,
	0x3c, 0xbc, 0xb5, 0x1a, 0xdc, 0x58, 0x6c, 0xb1, 0x46, 0xdf, 0x43, 0x48, 0xb3, 0x4c, 0xbe, 0x12,
	0x21, 0x05, 0x5b, 0x73, 0x6d, 0x98, 0xd8, 0xba, 0x68, 0xbb, 0xf3, 0x9d, 0xb9, 0xfa, 0x4f, 0xbb,
	0x72, 0xee, 0x67, 0x04, 0xa7, 0xdb, 0x4e, 0xaa, 0x63, 0xce, 0x8b, 0xb6, 0x63, 0xd7, 0x76, 0x52,
	0xb4, 0xb9, 0x5a, 0xde, 0x33, 0x81, 0x63, 0x2f, 0x22, 0xcc, 0x06, 0x90, 0x70, 0x31, 0x93, 0x21,
	0xda, 0xdf, 0xeb, 0x41, 0x3c, 0xf1, 0x11, 0x3d, 0xc8, 0xf5, 0x04, 0x3a, 0xf1, 0x9c, 0x8a, 0x94,
	0x91, 0x58, 0x8a, 0x98, 0x1a, 0xe2, 0xd7, 0xa5, 0x2c, 0xa5, 0xc3, 0x13, 0x3b, 0x7b, 0x5c, 0x31,
	0x6a, 0xc5, 0xf0, 0xe7, 0x5e, 0x77, 0xeb, 0x64, 0x6e, 0x75, 0xd8, 0x89, 0x3a, 0x3a, 0xdf, 0xb8,
	0xbb, 0x02, 0x74, 0x01, 0x10, 0xcb, 0x95, 0x30, 0x3e, 0x8f, 0x2e, 0xed, 0xe3, 0x8a, 0xb5, 0x8b,
	0xeb, 0x8e, 0x77, 0x21, 0x3c, 0x07, 0xff, 0x87, 0x14, 0xe9, 0xae, 0x8e, 0x4b, 0xfd, 0x2b, 0x5c,
	0x73, 0xe4, 0x23, 0x17, 0x7b, 0x82, 0x3c, 0xec, 0x6f, 0x04, 0x74, 0x7d, 0x73, 0xfb, 0xdb, 0x24,
	0xe5, 0x66, 0xbe, 0x9a, 0x0e, 0x62, 0xb9, 0x18, 0x0a, 0x66, 0xa6, 0x8a, 0x72, 0x31, 0x34, 0xb3,
	0x7e, 0xaa, 0x96, 0xf1, 0x30, 0x95, 0xc3, 0xbd, 0x8f, 0xec, 0x1e, 0x4c, 0xe5, 0xee, 0x93, 0xfb,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0xef, 0x7f, 0x49, 0xb7, 0x05, 0x00, 0x00,
}