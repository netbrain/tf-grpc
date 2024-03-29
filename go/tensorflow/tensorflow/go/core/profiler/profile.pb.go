// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/profiler/profile.proto

package profiler // import "github.com/netbrain/tf-grpc/go/tensorflow/tensorflow/go/core/profiler"

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

type Profile struct {
	SampleType           []*ValueType `protobuf:"bytes,1,rep,name=sample_type,json=sampleType,proto3" json:"sample_type,omitempty"`
	Sample               []*Sample    `protobuf:"bytes,2,rep,name=sample,proto3" json:"sample,omitempty"`
	Mapping              []*Mapping   `protobuf:"bytes,3,rep,name=mapping,proto3" json:"mapping,omitempty"`
	Location             []*Location  `protobuf:"bytes,4,rep,name=location,proto3" json:"location,omitempty"`
	Function             []*Function  `protobuf:"bytes,5,rep,name=function,proto3" json:"function,omitempty"`
	StringTable          []string     `protobuf:"bytes,6,rep,name=string_table,json=stringTable,proto3" json:"string_table,omitempty"`
	DropFrames           int64        `protobuf:"varint,7,opt,name=drop_frames,json=dropFrames,proto3" json:"drop_frames,omitempty"`
	KeepFrames           int64        `protobuf:"varint,8,opt,name=keep_frames,json=keepFrames,proto3" json:"keep_frames,omitempty"`
	TimeNanos            int64        `protobuf:"varint,9,opt,name=time_nanos,json=timeNanos,proto3" json:"time_nanos,omitempty"`
	DurationNanos        int64        `protobuf:"varint,10,opt,name=duration_nanos,json=durationNanos,proto3" json:"duration_nanos,omitempty"`
	PeriodType           *ValueType   `protobuf:"bytes,11,opt,name=period_type,json=periodType,proto3" json:"period_type,omitempty"`
	Period               int64        `protobuf:"varint,12,opt,name=period,proto3" json:"period,omitempty"`
	Comment              []int64      `protobuf:"varint,13,rep,packed,name=comment,proto3" json:"comment,omitempty"`
	DefaultSampleType    int64        `protobuf:"varint,14,opt,name=default_sample_type,json=defaultSampleType,proto3" json:"default_sample_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_7d4530c80c3943ef, []int{0}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (dst *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(dst, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetSampleType() []*ValueType {
	if m != nil {
		return m.SampleType
	}
	return nil
}

func (m *Profile) GetSample() []*Sample {
	if m != nil {
		return m.Sample
	}
	return nil
}

func (m *Profile) GetMapping() []*Mapping {
	if m != nil {
		return m.Mapping
	}
	return nil
}

func (m *Profile) GetLocation() []*Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Profile) GetFunction() []*Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *Profile) GetStringTable() []string {
	if m != nil {
		return m.StringTable
	}
	return nil
}

func (m *Profile) GetDropFrames() int64 {
	if m != nil {
		return m.DropFrames
	}
	return 0
}

func (m *Profile) GetKeepFrames() int64 {
	if m != nil {
		return m.KeepFrames
	}
	return 0
}

func (m *Profile) GetTimeNanos() int64 {
	if m != nil {
		return m.TimeNanos
	}
	return 0
}

func (m *Profile) GetDurationNanos() int64 {
	if m != nil {
		return m.DurationNanos
	}
	return 0
}

func (m *Profile) GetPeriodType() *ValueType {
	if m != nil {
		return m.PeriodType
	}
	return nil
}

func (m *Profile) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *Profile) GetComment() []int64 {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *Profile) GetDefaultSampleType() int64 {
	if m != nil {
		return m.DefaultSampleType
	}
	return 0
}

type ValueType struct {
	Type                 int64    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Unit                 int64    `protobuf:"varint,2,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueType) Reset()         { *m = ValueType{} }
func (m *ValueType) String() string { return proto.CompactTextString(m) }
func (*ValueType) ProtoMessage()    {}
func (*ValueType) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_7d4530c80c3943ef, []int{1}
}
func (m *ValueType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueType.Unmarshal(m, b)
}
func (m *ValueType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueType.Marshal(b, m, deterministic)
}
func (dst *ValueType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueType.Merge(dst, src)
}
func (m *ValueType) XXX_Size() int {
	return xxx_messageInfo_ValueType.Size(m)
}
func (m *ValueType) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueType.DiscardUnknown(m)
}

var xxx_messageInfo_ValueType proto.InternalMessageInfo

func (m *ValueType) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ValueType) GetUnit() int64 {
	if m != nil {
		return m.Unit
	}
	return 0
}

type Sample struct {
	LocationId           []uint64 `protobuf:"varint,1,rep,packed,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Value                []int64  `protobuf:"varint,2,rep,packed,name=value,proto3" json:"value,omitempty"`
	Label                []*Label `protobuf:"bytes,3,rep,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_7d4530c80c3943ef, []int{2}
}
func (m *Sample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sample.Unmarshal(m, b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
}
func (dst *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(dst, src)
}
func (m *Sample) XXX_Size() int {
	return xxx_messageInfo_Sample.Size(m)
}
func (m *Sample) XXX_DiscardUnknown() {
	xxx_messageInfo_Sample.DiscardUnknown(m)
}

var xxx_messageInfo_Sample proto.InternalMessageInfo

func (m *Sample) GetLocationId() []uint64 {
	if m != nil {
		return m.LocationId
	}
	return nil
}

func (m *Sample) GetValue() []int64 {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Sample) GetLabel() []*Label {
	if m != nil {
		return m.Label
	}
	return nil
}

type Label struct {
	Key                  int64    `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Str                  int64    `protobuf:"varint,2,opt,name=str,proto3" json:"str,omitempty"`
	Num                  int64    `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_7d4530c80c3943ef, []int{3}
}
func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (dst *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(dst, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Label) GetStr() int64 {
	if m != nil {
		return m.Str
	}
	return 0
}

func (m *Label) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type Mapping struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MemoryStart          uint64   `protobuf:"varint,2,opt,name=memory_start,json=memoryStart,proto3" json:"memory_start,omitempty"`
	MemoryLimit          uint64   `protobuf:"varint,3,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	FileOffset           uint64   `protobuf:"varint,4,opt,name=file_offset,json=fileOffset,proto3" json:"file_offset,omitempty"`
	Filename             int64    `protobuf:"varint,5,opt,name=filename,proto3" json:"filename,omitempty"`
	BuildId              int64    `protobuf:"varint,6,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	HasFunctions         bool     `protobuf:"varint,7,opt,name=has_functions,json=hasFunctions,proto3" json:"has_functions,omitempty"`
	HasFilenames         bool     `protobuf:"varint,8,opt,name=has_filenames,json=hasFilenames,proto3" json:"has_filenames,omitempty"`
	HasLineNumbers       bool     `protobuf:"varint,9,opt,name=has_line_numbers,json=hasLineNumbers,proto3" json:"has_line_numbers,omitempty"`
	HasInlineFrames      bool     `protobuf:"varint,10,opt,name=has_inline_frames,json=hasInlineFrames,proto3" json:"has_inline_frames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mapping) Reset()         { *m = Mapping{} }
func (m *Mapping) String() string { return proto.CompactTextString(m) }
func (*Mapping) ProtoMessage()    {}
func (*Mapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_7d4530c80c3943ef, []int{4}
}
func (m *Mapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mapping.Unmarshal(m, b)
}
func (m *Mapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mapping.Marshal(b, m, deterministic)
}
func (dst *Mapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mapping.Merge(dst, src)
}
func (m *Mapping) XXX_Size() int {
	return xxx_messageInfo_Mapping.Size(m)
}
func (m *Mapping) XXX_DiscardUnknown() {
	xxx_messageInfo_Mapping.DiscardUnknown(m)
}

var xxx_messageInfo_Mapping proto.InternalMessageInfo

func (m *Mapping) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Mapping) GetMemoryStart() uint64 {
	if m != nil {
		return m.MemoryStart
	}
	return 0
}

func (m *Mapping) GetMemoryLimit() uint64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *Mapping) GetFileOffset() uint64 {
	if m != nil {
		return m.FileOffset
	}
	return 0
}

func (m *Mapping) GetFilename() int64 {
	if m != nil {
		return m.Filename
	}
	return 0
}

func (m *Mapping) GetBuildId() int64 {
	if m != nil {
		return m.BuildId
	}
	return 0
}

func (m *Mapping) GetHasFunctions() bool {
	if m != nil {
		return m.HasFunctions
	}
	return false
}

func (m *Mapping) GetHasFilenames() bool {
	if m != nil {
		return m.HasFilenames
	}
	return false
}

func (m *Mapping) GetHasLineNumbers() bool {
	if m != nil {
		return m.HasLineNumbers
	}
	return false
}

func (m *Mapping) GetHasInlineFrames() bool {
	if m != nil {
		return m.HasInlineFrames
	}
	return false
}

type Location struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MappingId            uint64   `protobuf:"varint,2,opt,name=mapping_id,json=mappingId,proto3" json:"mapping_id,omitempty"`
	Address              uint64   `protobuf:"varint,3,opt,name=address,proto3" json:"address,omitempty"`
	Line                 []*Line  `protobuf:"bytes,4,rep,name=line,proto3" json:"line,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_7d4530c80c3943ef, []int{5}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Location) GetMappingId() uint64 {
	if m != nil {
		return m.MappingId
	}
	return 0
}

func (m *Location) GetAddress() uint64 {
	if m != nil {
		return m.Address
	}
	return 0
}

func (m *Location) GetLine() []*Line {
	if m != nil {
		return m.Line
	}
	return nil
}

type Line struct {
	FunctionId           uint64   `protobuf:"varint,1,opt,name=function_id,json=functionId,proto3" json:"function_id,omitempty"`
	Line                 int64    `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Line) Reset()         { *m = Line{} }
func (m *Line) String() string { return proto.CompactTextString(m) }
func (*Line) ProtoMessage()    {}
func (*Line) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_7d4530c80c3943ef, []int{6}
}
func (m *Line) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Line.Unmarshal(m, b)
}
func (m *Line) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Line.Marshal(b, m, deterministic)
}
func (dst *Line) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Line.Merge(dst, src)
}
func (m *Line) XXX_Size() int {
	return xxx_messageInfo_Line.Size(m)
}
func (m *Line) XXX_DiscardUnknown() {
	xxx_messageInfo_Line.DiscardUnknown(m)
}

var xxx_messageInfo_Line proto.InternalMessageInfo

func (m *Line) GetFunctionId() uint64 {
	if m != nil {
		return m.FunctionId
	}
	return 0
}

func (m *Line) GetLine() int64 {
	if m != nil {
		return m.Line
	}
	return 0
}

type Function struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 int64    `protobuf:"varint,2,opt,name=name,proto3" json:"name,omitempty"`
	SystemName           int64    `protobuf:"varint,3,opt,name=system_name,json=systemName,proto3" json:"system_name,omitempty"`
	Filename             int64    `protobuf:"varint,4,opt,name=filename,proto3" json:"filename,omitempty"`
	StartLine            int64    `protobuf:"varint,5,opt,name=start_line,json=startLine,proto3" json:"start_line,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Function) Reset()         { *m = Function{} }
func (m *Function) String() string { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()    {}
func (*Function) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_7d4530c80c3943ef, []int{7}
}
func (m *Function) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Function.Unmarshal(m, b)
}
func (m *Function) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Function.Marshal(b, m, deterministic)
}
func (dst *Function) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function.Merge(dst, src)
}
func (m *Function) XXX_Size() int {
	return xxx_messageInfo_Function.Size(m)
}
func (m *Function) XXX_DiscardUnknown() {
	xxx_messageInfo_Function.DiscardUnknown(m)
}

var xxx_messageInfo_Function proto.InternalMessageInfo

func (m *Function) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Function) GetName() int64 {
	if m != nil {
		return m.Name
	}
	return 0
}

func (m *Function) GetSystemName() int64 {
	if m != nil {
		return m.SystemName
	}
	return 0
}

func (m *Function) GetFilename() int64 {
	if m != nil {
		return m.Filename
	}
	return 0
}

func (m *Function) GetStartLine() int64 {
	if m != nil {
		return m.StartLine
	}
	return 0
}

func init() {
	proto.RegisterType((*Profile)(nil), "tensorflow.tfprof.pprof.Profile")
	proto.RegisterType((*ValueType)(nil), "tensorflow.tfprof.pprof.ValueType")
	proto.RegisterType((*Sample)(nil), "tensorflow.tfprof.pprof.Sample")
	proto.RegisterType((*Label)(nil), "tensorflow.tfprof.pprof.Label")
	proto.RegisterType((*Mapping)(nil), "tensorflow.tfprof.pprof.Mapping")
	proto.RegisterType((*Location)(nil), "tensorflow.tfprof.pprof.Location")
	proto.RegisterType((*Line)(nil), "tensorflow.tfprof.pprof.Line")
	proto.RegisterType((*Function)(nil), "tensorflow.tfprof.pprof.Function")
}

func init() {
	proto.RegisterFile("tensorflow/core/profiler/profile.proto", fileDescriptor_profile_7d4530c80c3943ef)
}

var fileDescriptor_profile_7d4530c80c3943ef = []byte{
	// 796 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5d, 0xaf, 0xe3, 0x34,
	0x10, 0x55, 0x9b, 0xf4, 0x6b, 0x7a, 0xef, 0x65, 0xd7, 0x20, 0x08, 0x48, 0xe5, 0x76, 0x83, 0x40,
	0x15, 0x12, 0xad, 0x60, 0x91, 0x90, 0x40, 0xfb, 0xc2, 0x8a, 0x45, 0x95, 0x2e, 0x17, 0x94, 0xbb,
	0xe2, 0x81, 0x97, 0xc8, 0x6d, 0x9c, 0xd6, 0xda, 0xc4, 0x8e, 0x6c, 0x07, 0xd4, 0x3f, 0xc0, 0x03,
	0xaf, 0xfc, 0x38, 0xfe, 0x0e, 0x9a, 0xb1, 0x13, 0x0a, 0xab, 0x8a, 0x7d, 0xa9, 0x66, 0x8e, 0xcf,
	0x71, 0xec, 0x99, 0x33, 0x35, 0x7c, 0xe2, 0x84, 0xb2, 0xda, 0x94, 0x95, 0xfe, 0x6d, 0xb3, 0xd7,
	0x46, 0x6c, 0x1a, 0xa3, 0x4b, 0x59, 0x09, 0xd3, 0x05, 0xeb, 0xc6, 0x68, 0xa7, 0xd9, 0x7b, 0xff,
	0xf0, 0xd6, 0xae, 0xc4, 0xb5, 0x75, 0x83, 0xbf, 0xe9, 0x9f, 0x23, 0x98, 0xfc, 0xe4, 0xa9, 0xec,
	0x39, 0xcc, 0x2d, 0xaf, 0x9b, 0x4a, 0xe4, 0xee, 0xd4, 0x88, 0x64, 0xb0, 0x8c, 0x56, 0xf3, 0x2f,
	0xd2, 0xf5, 0x05, 0xe9, 0xfa, 0x67, 0x5e, 0xb5, 0xe2, 0xe5, 0xa9, 0x11, 0x19, 0x78, 0x19, 0xc6,
	0xec, 0x2b, 0x18, 0xfb, 0x2c, 0x19, 0x92, 0xfe, 0xf6, 0xa2, 0xfe, 0x81, 0x68, 0x59, 0xa0, 0xb3,
	0xaf, 0x61, 0x52, 0xf3, 0xa6, 0x91, 0xea, 0x90, 0x44, 0xa4, 0x5c, 0x5e, 0x54, 0xfe, 0xe0, 0x79,
	0x59, 0x27, 0x60, 0xcf, 0x60, 0x5a, 0xe9, 0x3d, 0x77, 0x52, 0xab, 0x24, 0x26, 0xf1, 0x93, 0x8b,
	0xe2, 0xbb, 0x40, 0xcc, 0x7a, 0x09, 0xca, 0xcb, 0x56, 0xed, 0x49, 0x3e, 0xfa, 0x1f, 0xf9, 0x8b,
	0x40, 0xcc, 0x7a, 0x09, 0x7b, 0x02, 0x57, 0xd6, 0x19, 0xa9, 0x0e, 0xb9, 0xe3, 0xbb, 0x4a, 0x24,
	0xe3, 0x65, 0xb4, 0x9a, 0x65, 0x73, 0x8f, 0xbd, 0x44, 0x88, 0xdd, 0xc2, 0xbc, 0x30, 0xba, 0xc9,
	0x4b, 0xc3, 0x6b, 0x61, 0x93, 0xc9, 0x72, 0xb0, 0x8a, 0x32, 0x40, 0xe8, 0x05, 0x21, 0x48, 0x78,
	0x25, 0x44, 0x4f, 0x98, 0x7a, 0x02, 0x42, 0x81, 0xb0, 0x00, 0x70, 0xb2, 0x16, 0xb9, 0xe2, 0x4a,
	0xdb, 0x64, 0x46, 0xeb, 0x33, 0x44, 0xee, 0x11, 0x60, 0x1f, 0xc3, 0x4d, 0xd1, 0x1a, 0xba, 0x4e,
	0xa0, 0x00, 0x51, 0xae, 0x3b, 0xd4, 0xd3, 0x9e, 0xc3, 0xbc, 0x11, 0x46, 0xea, 0xc2, 0xb7, 0x78,
	0xbe, 0x1c, 0xbc, 0x69, 0x8b, 0xbd, 0x8c, 0x5a, 0xfc, 0x2e, 0x8c, 0x7d, 0x96, 0x5c, 0xd1, 0x37,
	0x42, 0xc6, 0x12, 0x98, 0xec, 0x75, 0x5d, 0x0b, 0xe5, 0x92, 0xeb, 0x65, 0xb4, 0x8a, 0xb2, 0x2e,
	0x65, 0x6b, 0x78, 0xbb, 0x10, 0x25, 0x6f, 0x2b, 0x97, 0x9f, 0x3b, 0xec, 0x86, 0xe4, 0x8f, 0xc3,
	0xd2, 0x43, 0x6f, 0xa2, 0xf4, 0x29, 0xcc, 0xfa, 0x4f, 0x33, 0x06, 0x71, 0xf0, 0x23, 0xb2, 0x29,
	0x46, 0xac, 0x55, 0xd2, 0x25, 0x43, 0x8f, 0x61, 0x9c, 0xb6, 0x30, 0xf6, 0x5b, 0x60, 0x31, 0xbb,
	0xde, 0xe6, 0xb2, 0x20, 0x23, 0xc7, 0x19, 0x74, 0xd0, 0xb6, 0x60, 0xef, 0xc0, 0xe8, 0x57, 0xdc,
	0x9f, 0x3c, 0x1a, 0x65, 0x3e, 0x61, 0x5f, 0xc2, 0xa8, 0xe2, 0x3b, 0x51, 0x05, 0xff, 0x7d, 0x78,
	0xd9, 0x42, 0xc8, 0xca, 0x3c, 0x39, 0x7d, 0x06, 0x23, 0xca, 0xd9, 0x23, 0x88, 0x5e, 0x89, 0x53,
	0x38, 0x26, 0x86, 0x88, 0x58, 0x67, 0xc2, 0x21, 0x31, 0x44, 0x44, 0xb5, 0x75, 0x12, 0x79, 0x44,
	0xb5, 0x75, 0xfa, 0xd7, 0x10, 0x26, 0xc1, 0xcf, 0xec, 0x06, 0x86, 0x74, 0xdc, 0xc1, 0x2a, 0xce,
	0x86, 0xb2, 0x40, 0x63, 0xd5, 0xa2, 0xd6, 0xe6, 0x94, 0x5b, 0xc7, 0x8d, 0xbf, 0x6d, 0x9c, 0xcd,
	0x3d, 0xf6, 0x80, 0xd0, 0x19, 0xa5, 0x92, 0xb5, 0x74, 0xb4, 0x73, 0x4f, 0xb9, 0x43, 0x08, 0xab,
	0x81, 0xe3, 0x9d, 0xeb, 0xb2, 0xb4, 0xc2, 0x25, 0x31, 0x31, 0x00, 0xa1, 0x1f, 0x09, 0x61, 0x1f,
	0xc0, 0x14, 0x33, 0xc5, 0x6b, 0x91, 0x8c, 0xe8, 0x64, 0x7d, 0xce, 0xde, 0x87, 0xe9, 0xae, 0x95,
	0x55, 0x81, 0x75, 0x1c, 0xd3, 0xda, 0x84, 0xf2, 0x6d, 0xc1, 0x3e, 0x82, 0xeb, 0x23, 0xb7, 0x79,
	0x37, 0x06, 0xde, 0xd5, 0xd3, 0xec, 0xea, 0xc8, 0x6d, 0x37, 0x24, 0xb6, 0x27, 0x85, 0xfd, 0xbc,
	0xb3, 0x03, 0xa9, 0xc3, 0xd8, 0x0a, 0x1e, 0x21, 0xa9, 0x92, 0x4a, 0xe4, 0xaa, 0xad, 0x77, 0xc2,
	0x78, 0x87, 0x4f, 0xb3, 0x9b, 0x23, 0xb7, 0x77, 0x52, 0x89, 0x7b, 0x8f, 0xb2, 0x4f, 0xe1, 0x31,
	0x32, 0xa5, 0x22, 0x6e, 0x18, 0x16, 0x20, 0xea, 0x5b, 0x47, 0x6e, 0xb7, 0x84, 0xfb, 0x89, 0x49,
	0x7f, 0x1f, 0xc0, 0xb4, 0x1b, 0xf6, 0xd7, 0x4a, 0xbb, 0x00, 0x08, 0x7f, 0x1e, 0x78, 0x33, 0x5f,
	0xd8, 0x59, 0x40, 0xb6, 0x64, 0x65, 0x5e, 0x14, 0x46, 0x58, 0x1b, 0x2a, 0xda, 0xa5, 0xec, 0x73,
	0x88, 0xf1, 0x1b, 0xe1, 0x6f, 0x66, 0x71, 0xd9, 0x23, 0x52, 0x89, 0x8c, 0xa8, 0xe9, 0x37, 0x10,
	0x63, 0x46, 0x8d, 0x08, 0x85, 0xc9, 0xfb, 0xc3, 0x40, 0x07, 0x6d, 0x0b, 0x74, 0x35, 0xed, 0x1d,
	0x5c, 0x4d, 0xe2, 0x3f, 0x06, 0x30, 0xed, 0xca, 0xf9, 0xda, 0x2d, 0x18, 0xc4, 0xd4, 0xb5, 0x20,
	0xa0, 0x8e, 0xdd, 0xc2, 0xdc, 0x9e, 0xac, 0x13, 0x75, 0x4e, 0x4b, 0xde, 0x6a, 0xe0, 0xa1, 0x7b,
	0x24, 0x9c, 0xb7, 0x3b, 0xfe, 0x4f, 0xbb, 0x17, 0x00, 0x64, 0x35, 0xea, 0x45, 0x30, 0xc3, 0x8c,
	0x10, 0xbc, 0xc1, 0xb7, 0xdf, 0xff, 0xf2, 0xdd, 0x41, 0xba, 0x63, 0xbb, 0x5b, 0xef, 0x75, 0xbd,
	0x51, 0xc2, 0xed, 0x0c, 0x97, 0x6a, 0xe3, 0xca, 0xcf, 0x0e, 0xa6, 0xd9, 0x6f, 0x0e, 0x7a, 0x73,
	0xf6, 0x1e, 0x9d, 0x85, 0x07, 0xfd, 0xef, 0xd7, 0x69, 0x37, 0xa6, 0x67, 0xe9, 0xe9, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x09, 0x0a, 0x19, 0x49, 0xc0, 0x06, 0x00, 0x00,
}
