// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/summary.proto

package framework // import "github.com/netbrain/tf-grpc/go/tensorflow/tensorflow/go/core/framework"

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

// Metadata associated with a series of Summary data
type SummaryDescription struct {
	// Hint on how plugins should process the data in this series.
	// Supported values include "scalar", "histogram", "image", "audio"
	TypeHint             string   `protobuf:"bytes,1,opt,name=type_hint,json=typeHint,proto3" json:"type_hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SummaryDescription) Reset()         { *m = SummaryDescription{} }
func (m *SummaryDescription) String() string { return proto.CompactTextString(m) }
func (*SummaryDescription) ProtoMessage()    {}
func (*SummaryDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_summary_5fe47b9d4a6e0071, []int{0}
}
func (m *SummaryDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummaryDescription.Unmarshal(m, b)
}
func (m *SummaryDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummaryDescription.Marshal(b, m, deterministic)
}
func (dst *SummaryDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummaryDescription.Merge(dst, src)
}
func (m *SummaryDescription) XXX_Size() int {
	return xxx_messageInfo_SummaryDescription.Size(m)
}
func (m *SummaryDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_SummaryDescription.DiscardUnknown(m)
}

var xxx_messageInfo_SummaryDescription proto.InternalMessageInfo

func (m *SummaryDescription) GetTypeHint() string {
	if m != nil {
		return m.TypeHint
	}
	return ""
}

// Serialization format for histogram module in
// core/lib/histogram/histogram.h
type HistogramProto struct {
	Min        float64 `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	Max        float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	Num        float64 `protobuf:"fixed64,3,opt,name=num,proto3" json:"num,omitempty"`
	Sum        float64 `protobuf:"fixed64,4,opt,name=sum,proto3" json:"sum,omitempty"`
	SumSquares float64 `protobuf:"fixed64,5,opt,name=sum_squares,json=sumSquares,proto3" json:"sum_squares,omitempty"`
	// Parallel arrays encoding the bucket boundaries and the bucket values.
	// bucket(i) is the count for the bucket i.  The range for
	// a bucket is:
	//   i == 0:  -DBL_MAX .. bucket_limit(0)
	//   i != 0:  bucket_limit(i-1) .. bucket_limit(i)
	BucketLimit          []float64 `protobuf:"fixed64,6,rep,packed,name=bucket_limit,json=bucketLimit,proto3" json:"bucket_limit,omitempty"`
	Bucket               []float64 `protobuf:"fixed64,7,rep,packed,name=bucket,proto3" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *HistogramProto) Reset()         { *m = HistogramProto{} }
func (m *HistogramProto) String() string { return proto.CompactTextString(m) }
func (*HistogramProto) ProtoMessage()    {}
func (*HistogramProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_summary_5fe47b9d4a6e0071, []int{1}
}
func (m *HistogramProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistogramProto.Unmarshal(m, b)
}
func (m *HistogramProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistogramProto.Marshal(b, m, deterministic)
}
func (dst *HistogramProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistogramProto.Merge(dst, src)
}
func (m *HistogramProto) XXX_Size() int {
	return xxx_messageInfo_HistogramProto.Size(m)
}
func (m *HistogramProto) XXX_DiscardUnknown() {
	xxx_messageInfo_HistogramProto.DiscardUnknown(m)
}

var xxx_messageInfo_HistogramProto proto.InternalMessageInfo

func (m *HistogramProto) GetMin() float64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *HistogramProto) GetMax() float64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *HistogramProto) GetNum() float64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *HistogramProto) GetSum() float64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *HistogramProto) GetSumSquares() float64 {
	if m != nil {
		return m.SumSquares
	}
	return 0
}

func (m *HistogramProto) GetBucketLimit() []float64 {
	if m != nil {
		return m.BucketLimit
	}
	return nil
}

func (m *HistogramProto) GetBucket() []float64 {
	if m != nil {
		return m.Bucket
	}
	return nil
}

// A SummaryMetadata encapsulates information on which plugins are able to make
// use of a certain summary value.
type SummaryMetadata struct {
	// Data that associates a summary with a certain plugin.
	PluginData *SummaryMetadata_PluginData `protobuf:"bytes,1,opt,name=plugin_data,json=pluginData,proto3" json:"plugin_data,omitempty"`
	// Display name for viewing in TensorBoard.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Longform readable description of the summary sequence. Markdown supported.
	SummaryDescription   string   `protobuf:"bytes,3,opt,name=summary_description,json=summaryDescription,proto3" json:"summary_description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SummaryMetadata) Reset()         { *m = SummaryMetadata{} }
func (m *SummaryMetadata) String() string { return proto.CompactTextString(m) }
func (*SummaryMetadata) ProtoMessage()    {}
func (*SummaryMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_summary_5fe47b9d4a6e0071, []int{2}
}
func (m *SummaryMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummaryMetadata.Unmarshal(m, b)
}
func (m *SummaryMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummaryMetadata.Marshal(b, m, deterministic)
}
func (dst *SummaryMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummaryMetadata.Merge(dst, src)
}
func (m *SummaryMetadata) XXX_Size() int {
	return xxx_messageInfo_SummaryMetadata.Size(m)
}
func (m *SummaryMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_SummaryMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_SummaryMetadata proto.InternalMessageInfo

func (m *SummaryMetadata) GetPluginData() *SummaryMetadata_PluginData {
	if m != nil {
		return m.PluginData
	}
	return nil
}

func (m *SummaryMetadata) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *SummaryMetadata) GetSummaryDescription() string {
	if m != nil {
		return m.SummaryDescription
	}
	return ""
}

type SummaryMetadata_PluginData struct {
	// The name of the plugin this data pertains to.
	PluginName string `protobuf:"bytes,1,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
	// The content to store for the plugin. The best practice is for this to be
	// a binary serialized protocol buffer.
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SummaryMetadata_PluginData) Reset()         { *m = SummaryMetadata_PluginData{} }
func (m *SummaryMetadata_PluginData) String() string { return proto.CompactTextString(m) }
func (*SummaryMetadata_PluginData) ProtoMessage()    {}
func (*SummaryMetadata_PluginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_summary_5fe47b9d4a6e0071, []int{2, 0}
}
func (m *SummaryMetadata_PluginData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummaryMetadata_PluginData.Unmarshal(m, b)
}
func (m *SummaryMetadata_PluginData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummaryMetadata_PluginData.Marshal(b, m, deterministic)
}
func (dst *SummaryMetadata_PluginData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummaryMetadata_PluginData.Merge(dst, src)
}
func (m *SummaryMetadata_PluginData) XXX_Size() int {
	return xxx_messageInfo_SummaryMetadata_PluginData.Size(m)
}
func (m *SummaryMetadata_PluginData) XXX_DiscardUnknown() {
	xxx_messageInfo_SummaryMetadata_PluginData.DiscardUnknown(m)
}

var xxx_messageInfo_SummaryMetadata_PluginData proto.InternalMessageInfo

func (m *SummaryMetadata_PluginData) GetPluginName() string {
	if m != nil {
		return m.PluginName
	}
	return ""
}

func (m *SummaryMetadata_PluginData) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// A Summary is a set of named values to be displayed by the
// visualizer.
//
// Summaries are produced regularly during training, as controlled by
// the "summary_interval_secs" attribute of the training operation.
// Summaries are also produced at the end of an evaluation.
type Summary struct {
	// Set of values for the summary.
	Value                []*Summary_Value `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Summary) Reset()         { *m = Summary{} }
func (m *Summary) String() string { return proto.CompactTextString(m) }
func (*Summary) ProtoMessage()    {}
func (*Summary) Descriptor() ([]byte, []int) {
	return fileDescriptor_summary_5fe47b9d4a6e0071, []int{3}
}
func (m *Summary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary.Unmarshal(m, b)
}
func (m *Summary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary.Marshal(b, m, deterministic)
}
func (dst *Summary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary.Merge(dst, src)
}
func (m *Summary) XXX_Size() int {
	return xxx_messageInfo_Summary.Size(m)
}
func (m *Summary) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary.DiscardUnknown(m)
}

var xxx_messageInfo_Summary proto.InternalMessageInfo

func (m *Summary) GetValue() []*Summary_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type Summary_Image struct {
	// Dimensions of the image.
	Height int32 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Width  int32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	// Valid colorspace values are
	//   1 - grayscale
	//   2 - grayscale + alpha
	//   3 - RGB
	//   4 - RGBA
	//   5 - DIGITAL_YUV
	//   6 - BGRA
	Colorspace int32 `protobuf:"varint,3,opt,name=colorspace,proto3" json:"colorspace,omitempty"`
	// Image data in encoded format.  All image formats supported by
	// image_codec::CoderUtil can be stored here.
	EncodedImageString   []byte   `protobuf:"bytes,4,opt,name=encoded_image_string,json=encodedImageString,proto3" json:"encoded_image_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Summary_Image) Reset()         { *m = Summary_Image{} }
func (m *Summary_Image) String() string { return proto.CompactTextString(m) }
func (*Summary_Image) ProtoMessage()    {}
func (*Summary_Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_summary_5fe47b9d4a6e0071, []int{3, 0}
}
func (m *Summary_Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary_Image.Unmarshal(m, b)
}
func (m *Summary_Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary_Image.Marshal(b, m, deterministic)
}
func (dst *Summary_Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary_Image.Merge(dst, src)
}
func (m *Summary_Image) XXX_Size() int {
	return xxx_messageInfo_Summary_Image.Size(m)
}
func (m *Summary_Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Summary_Image proto.InternalMessageInfo

func (m *Summary_Image) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Summary_Image) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Summary_Image) GetColorspace() int32 {
	if m != nil {
		return m.Colorspace
	}
	return 0
}

func (m *Summary_Image) GetEncodedImageString() []byte {
	if m != nil {
		return m.EncodedImageString
	}
	return nil
}

type Summary_Audio struct {
	// Sample rate of the audio in Hz.
	SampleRate float32 `protobuf:"fixed32,1,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	// Number of channels of audio.
	NumChannels int64 `protobuf:"varint,2,opt,name=num_channels,json=numChannels,proto3" json:"num_channels,omitempty"`
	// Length of the audio in frames (samples per channel).
	LengthFrames int64 `protobuf:"varint,3,opt,name=length_frames,json=lengthFrames,proto3" json:"length_frames,omitempty"`
	// Encoded audio data and its associated RFC 2045 content type (e.g.
	// "audio/wav").
	EncodedAudioString   []byte   `protobuf:"bytes,4,opt,name=encoded_audio_string,json=encodedAudioString,proto3" json:"encoded_audio_string,omitempty"`
	ContentType          string   `protobuf:"bytes,5,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Summary_Audio) Reset()         { *m = Summary_Audio{} }
func (m *Summary_Audio) String() string { return proto.CompactTextString(m) }
func (*Summary_Audio) ProtoMessage()    {}
func (*Summary_Audio) Descriptor() ([]byte, []int) {
	return fileDescriptor_summary_5fe47b9d4a6e0071, []int{3, 1}
}
func (m *Summary_Audio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary_Audio.Unmarshal(m, b)
}
func (m *Summary_Audio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary_Audio.Marshal(b, m, deterministic)
}
func (dst *Summary_Audio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary_Audio.Merge(dst, src)
}
func (m *Summary_Audio) XXX_Size() int {
	return xxx_messageInfo_Summary_Audio.Size(m)
}
func (m *Summary_Audio) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary_Audio.DiscardUnknown(m)
}

var xxx_messageInfo_Summary_Audio proto.InternalMessageInfo

func (m *Summary_Audio) GetSampleRate() float32 {
	if m != nil {
		return m.SampleRate
	}
	return 0
}

func (m *Summary_Audio) GetNumChannels() int64 {
	if m != nil {
		return m.NumChannels
	}
	return 0
}

func (m *Summary_Audio) GetLengthFrames() int64 {
	if m != nil {
		return m.LengthFrames
	}
	return 0
}

func (m *Summary_Audio) GetEncodedAudioString() []byte {
	if m != nil {
		return m.EncodedAudioString
	}
	return nil
}

func (m *Summary_Audio) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

type Summary_Value struct {
	// This field is deprecated and will not be set.
	NodeName string `protobuf:"bytes,7,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Tag name for the data. Used by TensorBoard plugins to organize data. Tags
	// are often organized by scope (which contains slashes to convey
	// hierarchy). For example: foo/bar/0
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Contains metadata on the summary value such as which plugins may use it.
	// Take note that many summary values may lack a metadata field. This is
	// because the FileWriter only keeps a metadata object on the first summary
	// value with a certain tag for each tag. TensorBoard then remembers which
	// tags are associated with which plugins. This saves space.
	Metadata *SummaryMetadata `protobuf:"bytes,9,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Value associated with the tag.
	//
	// Types that are valid to be assigned to Value:
	//	*Summary_Value_SimpleValue
	//	*Summary_Value_ObsoleteOldStyleHistogram
	//	*Summary_Value_Image
	//	*Summary_Value_Histo
	//	*Summary_Value_Audio
	//	*Summary_Value_Tensor
	Value                isSummary_Value_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Summary_Value) Reset()         { *m = Summary_Value{} }
func (m *Summary_Value) String() string { return proto.CompactTextString(m) }
func (*Summary_Value) ProtoMessage()    {}
func (*Summary_Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_summary_5fe47b9d4a6e0071, []int{3, 2}
}
func (m *Summary_Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary_Value.Unmarshal(m, b)
}
func (m *Summary_Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary_Value.Marshal(b, m, deterministic)
}
func (dst *Summary_Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary_Value.Merge(dst, src)
}
func (m *Summary_Value) XXX_Size() int {
	return xxx_messageInfo_Summary_Value.Size(m)
}
func (m *Summary_Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Summary_Value proto.InternalMessageInfo

func (m *Summary_Value) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Summary_Value) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Summary_Value) GetMetadata() *SummaryMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type isSummary_Value_Value interface {
	isSummary_Value_Value()
}

type Summary_Value_SimpleValue struct {
	SimpleValue float32 `protobuf:"fixed32,2,opt,name=simple_value,json=simpleValue,proto3,oneof"`
}

type Summary_Value_ObsoleteOldStyleHistogram struct {
	ObsoleteOldStyleHistogram []byte `protobuf:"bytes,3,opt,name=obsolete_old_style_histogram,json=obsoleteOldStyleHistogram,proto3,oneof"`
}

type Summary_Value_Image struct {
	Image *Summary_Image `protobuf:"bytes,4,opt,name=image,proto3,oneof"`
}

type Summary_Value_Histo struct {
	Histo *HistogramProto `protobuf:"bytes,5,opt,name=histo,proto3,oneof"`
}

type Summary_Value_Audio struct {
	Audio *Summary_Audio `protobuf:"bytes,6,opt,name=audio,proto3,oneof"`
}

type Summary_Value_Tensor struct {
	Tensor *TensorProto `protobuf:"bytes,8,opt,name=tensor,proto3,oneof"`
}

func (*Summary_Value_SimpleValue) isSummary_Value_Value() {}

func (*Summary_Value_ObsoleteOldStyleHistogram) isSummary_Value_Value() {}

func (*Summary_Value_Image) isSummary_Value_Value() {}

func (*Summary_Value_Histo) isSummary_Value_Value() {}

func (*Summary_Value_Audio) isSummary_Value_Value() {}

func (*Summary_Value_Tensor) isSummary_Value_Value() {}

func (m *Summary_Value) GetValue() isSummary_Value_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Summary_Value) GetSimpleValue() float32 {
	if x, ok := m.GetValue().(*Summary_Value_SimpleValue); ok {
		return x.SimpleValue
	}
	return 0
}

func (m *Summary_Value) GetObsoleteOldStyleHistogram() []byte {
	if x, ok := m.GetValue().(*Summary_Value_ObsoleteOldStyleHistogram); ok {
		return x.ObsoleteOldStyleHistogram
	}
	return nil
}

func (m *Summary_Value) GetImage() *Summary_Image {
	if x, ok := m.GetValue().(*Summary_Value_Image); ok {
		return x.Image
	}
	return nil
}

func (m *Summary_Value) GetHisto() *HistogramProto {
	if x, ok := m.GetValue().(*Summary_Value_Histo); ok {
		return x.Histo
	}
	return nil
}

func (m *Summary_Value) GetAudio() *Summary_Audio {
	if x, ok := m.GetValue().(*Summary_Value_Audio); ok {
		return x.Audio
	}
	return nil
}

func (m *Summary_Value) GetTensor() *TensorProto {
	if x, ok := m.GetValue().(*Summary_Value_Tensor); ok {
		return x.Tensor
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Summary_Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Summary_Value_OneofMarshaler, _Summary_Value_OneofUnmarshaler, _Summary_Value_OneofSizer, []interface{}{
		(*Summary_Value_SimpleValue)(nil),
		(*Summary_Value_ObsoleteOldStyleHistogram)(nil),
		(*Summary_Value_Image)(nil),
		(*Summary_Value_Histo)(nil),
		(*Summary_Value_Audio)(nil),
		(*Summary_Value_Tensor)(nil),
	}
}

func _Summary_Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Summary_Value)
	// value
	switch x := m.Value.(type) {
	case *Summary_Value_SimpleValue:
		b.EncodeVarint(2<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.SimpleValue)))
	case *Summary_Value_ObsoleteOldStyleHistogram:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.ObsoleteOldStyleHistogram)
	case *Summary_Value_Image:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Image); err != nil {
			return err
		}
	case *Summary_Value_Histo:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Histo); err != nil {
			return err
		}
	case *Summary_Value_Audio:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Audio); err != nil {
			return err
		}
	case *Summary_Value_Tensor:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tensor); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Summary_Value.Value has unexpected type %T", x)
	}
	return nil
}

func _Summary_Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Summary_Value)
	switch tag {
	case 2: // value.simple_value
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &Summary_Value_SimpleValue{math.Float32frombits(uint32(x))}
		return true, err
	case 3: // value.obsolete_old_style_histogram
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &Summary_Value_ObsoleteOldStyleHistogram{x}
		return true, err
	case 4: // value.image
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Summary_Image)
		err := b.DecodeMessage(msg)
		m.Value = &Summary_Value_Image{msg}
		return true, err
	case 5: // value.histo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HistogramProto)
		err := b.DecodeMessage(msg)
		m.Value = &Summary_Value_Histo{msg}
		return true, err
	case 6: // value.audio
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Summary_Audio)
		err := b.DecodeMessage(msg)
		m.Value = &Summary_Value_Audio{msg}
		return true, err
	case 8: // value.tensor
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TensorProto)
		err := b.DecodeMessage(msg)
		m.Value = &Summary_Value_Tensor{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Summary_Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Summary_Value)
	// value
	switch x := m.Value.(type) {
	case *Summary_Value_SimpleValue:
		n += 1 // tag and wire
		n += 4
	case *Summary_Value_ObsoleteOldStyleHistogram:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ObsoleteOldStyleHistogram)))
		n += len(x.ObsoleteOldStyleHistogram)
	case *Summary_Value_Image:
		s := proto.Size(x.Image)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Summary_Value_Histo:
		s := proto.Size(x.Histo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Summary_Value_Audio:
		s := proto.Size(x.Audio)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Summary_Value_Tensor:
		s := proto.Size(x.Tensor)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*SummaryDescription)(nil), "tensorflow.SummaryDescription")
	proto.RegisterType((*HistogramProto)(nil), "tensorflow.HistogramProto")
	proto.RegisterType((*SummaryMetadata)(nil), "tensorflow.SummaryMetadata")
	proto.RegisterType((*SummaryMetadata_PluginData)(nil), "tensorflow.SummaryMetadata.PluginData")
	proto.RegisterType((*Summary)(nil), "tensorflow.Summary")
	proto.RegisterType((*Summary_Image)(nil), "tensorflow.Summary.Image")
	proto.RegisterType((*Summary_Audio)(nil), "tensorflow.Summary.Audio")
	proto.RegisterType((*Summary_Value)(nil), "tensorflow.Summary.Value")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/summary.proto", fileDescriptor_summary_5fe47b9d4a6e0071)
}

var fileDescriptor_summary_5fe47b9d4a6e0071 = []byte{
	// 785 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x92, 0x1b, 0x35,
	0x10, 0xde, 0xf1, 0x30, 0xf6, 0xba, 0xc7, 0x81, 0x2d, 0x91, 0x82, 0x89, 0x43, 0x41, 0x70, 0x8a,
	0xb0, 0x17, 0x3c, 0xac, 0x39, 0x70, 0x8e, 0x49, 0x25, 0xa6, 0x8a, 0x9f, 0x2d, 0x39, 0xc5, 0x81,
	0xcb, 0x94, 0x3c, 0xa3, 0x1d, 0xab, 0x32, 0x92, 0x86, 0x91, 0x26, 0x1b, 0x3f, 0x01, 0x4f, 0xc3,
	0x1b, 0xe4, 0xca, 0x33, 0xc1, 0x91, 0x52, 0x4b, 0x6b, 0x9b, 0xb0, 0xcb, 0x4d, 0xfa, 0xf4, 0xb5,
	0xbe, 0x56, 0x7f, 0xad, 0x86, 0x2f, 0x2d, 0x57, 0x46, 0x77, 0x57, 0x8d, 0xbe, 0xce, 0x4b, 0xdd,
	0xf1, 0xfc, 0xaa, 0x63, 0x92, 0x5f, 0xeb, 0xee, 0x55, 0x6e, 0x7a, 0x29, 0x59, 0xb7, 0x9b, 0xb7,
	0x9d, 0xb6, 0x9a, 0xc0, 0x81, 0x38, 0x7d, 0x72, 0x77, 0x90, 0x3f, 0xf1, 0x31, 0xb3, 0x0b, 0x20,
	0x6b, 0x7f, 0xc9, 0x33, 0x6e, 0xca, 0x4e, 0xb4, 0x56, 0x68, 0x45, 0x1e, 0xc2, 0xd8, 0xee, 0x5a,
	0x5e, 0x6c, 0x85, 0xb2, 0x59, 0xf4, 0x28, 0x3a, 0x1f, 0xd3, 0x53, 0x07, 0xac, 0x84, 0xb2, 0xb3,
	0xb7, 0x11, 0xbc, 0xbf, 0x12, 0xc6, 0xea, 0xba, 0x63, 0xf2, 0x12, 0x95, 0xcf, 0x20, 0x96, 0x42,
	0x21, 0x33, 0xa2, 0x6e, 0x89, 0x08, 0x7b, 0x93, 0x0d, 0x02, 0xc2, 0xde, 0x38, 0x44, 0xf5, 0x32,
	0x8b, 0x3d, 0xa2, 0x7a, 0xe9, 0x10, 0xd3, 0xcb, 0xec, 0x3d, 0x8f, 0x98, 0x5e, 0x92, 0xcf, 0x20,
	0x35, 0xbd, 0x2c, 0xcc, 0x6f, 0x3d, 0xeb, 0xb8, 0xc9, 0x12, 0x3c, 0x01, 0xd3, 0xcb, 0xb5, 0x47,
	0xc8, 0x17, 0x30, 0xd9, 0xf4, 0xe5, 0x2b, 0x6e, 0x8b, 0x46, 0x48, 0x61, 0xb3, 0xe1, 0xa3, 0xf8,
	0x3c, 0x5a, 0x0e, 0xce, 0x22, 0x9a, 0x7a, 0xfc, 0x07, 0x07, 0x93, 0x29, 0x0c, 0xfd, 0x36, 0x1b,
	0xed, 0x09, 0x01, 0x99, 0xfd, 0x15, 0xc1, 0x07, 0xe1, 0xc9, 0x3f, 0x72, 0xcb, 0x2a, 0x66, 0x19,
	0x79, 0x01, 0x69, 0xdb, 0xf4, 0xb5, 0x50, 0x85, 0xdb, 0xe2, 0x3b, 0xd2, 0xc5, 0x93, 0xf9, 0xa1,
	0x86, 0xf3, 0x77, 0x22, 0xe6, 0x97, 0x48, 0x7f, 0xc6, 0x2c, 0xa3, 0xd0, 0xee, 0xd7, 0xe4, 0x73,
	0x98, 0x54, 0xc2, 0xb4, 0x0d, 0xdb, 0x15, 0x8a, 0x49, 0x8e, 0xef, 0x1f, 0xd3, 0x34, 0x60, 0x3f,
	0x31, 0xc9, 0x49, 0x0e, 0x1f, 0x06, 0xdb, 0x8a, 0xea, 0x50, 0x72, 0xac, 0xcb, 0x98, 0x12, 0xf3,
	0x1f, 0x33, 0xa6, 0x2f, 0x00, 0x0e, 0x6a, 0xae, 0x44, 0x21, 0x55, 0x14, 0xf0, 0xe6, 0x84, 0x14,
	0xf0, 0xfe, 0x0c, 0x46, 0xa5, 0x56, 0x96, 0x2b, 0x8b, 0xea, 0x13, 0x7a, 0xb3, 0x9d, 0xbd, 0x1d,
	0xc2, 0x28, 0xbc, 0x83, 0xe4, 0x90, 0xbc, 0x66, 0x4d, 0xef, 0x2e, 0x88, 0xcf, 0xd3, 0xc5, 0x83,
	0x5b, 0xde, 0x3a, 0xff, 0xc5, 0x11, 0xa8, 0xe7, 0x4d, 0x7f, 0x8f, 0x20, 0xf9, 0x5e, 0xb2, 0x9a,
	0x93, 0x8f, 0x60, 0xb8, 0xe5, 0xa2, 0xde, 0xfa, 0xce, 0x48, 0x68, 0xd8, 0x91, 0xfb, 0x90, 0x5c,
	0x8b, 0xca, 0x6e, 0x51, 0x36, 0xa1, 0x7e, 0x43, 0x3e, 0x05, 0x28, 0x75, 0xa3, 0x3b, 0xd3, 0xb2,
	0x92, 0xe3, 0x2b, 0x13, 0x7a, 0x84, 0x90, 0xaf, 0xe1, 0x3e, 0x57, 0xa5, 0xae, 0x78, 0x55, 0x08,
	0x77, 0x7d, 0x61, 0x6c, 0x27, 0x54, 0x8d, 0x5d, 0x31, 0xa1, 0x24, 0x9c, 0xa1, 0xf2, 0x1a, 0x4f,
	0xa6, 0x7f, 0x46, 0x90, 0x3c, 0xed, 0x2b, 0xa1, 0xb1, 0x5d, 0x98, 0x6c, 0x1b, 0x5e, 0x74, 0xcc,
	0xfa, 0x5a, 0x0c, 0x28, 0x78, 0x88, 0x32, 0xcb, 0x9d, 0x1d, 0xaa, 0x97, 0x45, 0xb9, 0x65, 0x4a,
	0xf1, 0xc6, 0x60, 0x66, 0x31, 0x4d, 0x55, 0x2f, 0xbf, 0x0b, 0x10, 0x79, 0x0c, 0xf7, 0x1a, 0xae,
	0x6a, 0xbb, 0x2d, 0xf0, 0x87, 0x18, 0x4c, 0x31, 0xa6, 0x13, 0x0f, 0x3e, 0x47, 0xec, 0x38, 0x49,
	0xe6, 0x94, 0x6f, 0x4f, 0x12, 0x93, 0xf2, 0x49, 0x3a, 0xe5, 0x50, 0xf6, 0xc2, 0x7d, 0x1c, 0x6c,
	0xe5, 0x31, 0x4d, 0x03, 0xf6, 0x72, 0xd7, 0xf2, 0xe9, 0x1f, 0x31, 0x24, 0x58, 0x62, 0xf7, 0xdd,
	0x94, 0xae, 0xb8, 0x77, 0x74, 0xe4, 0xbf, 0x9b, 0x03, 0xd0, 0xcf, 0x33, 0x88, 0x2d, 0xab, 0x83,
	0xd1, 0x6e, 0x49, 0xbe, 0x85, 0x53, 0x19, 0xfa, 0x30, 0x1b, 0x63, 0xab, 0x3e, 0xfc, 0x9f, 0x56,
	0xa5, 0x7b, 0x32, 0x79, 0x0c, 0x13, 0x23, 0xb0, 0x5e, 0xde, 0x7b, 0x57, 0x8e, 0xc1, 0xea, 0x84,
	0xa6, 0x1e, 0xf5, 0xc9, 0x3c, 0x85, 0x4f, 0xf4, 0xc6, 0xe8, 0x86, 0x5b, 0x5e, 0xe8, 0xa6, 0x2a,
	0x8c, 0xdd, 0x35, 0x6e, 0x12, 0x84, 0x0f, 0x8f, 0xf5, 0x99, 0xac, 0x4e, 0xe8, 0x83, 0x1b, 0xd6,
	0xcf, 0x4d, 0xb5, 0x76, 0x9c, 0xfd, 0x4c, 0x20, 0x17, 0x90, 0xa0, 0x97, 0x58, 0x9f, 0x3b, 0x9a,
	0x0b, 0x1d, 0x5d, 0x9d, 0x50, 0xcf, 0x24, 0x0b, 0x48, 0x50, 0x02, 0x0b, 0x95, 0x2e, 0xa6, 0xc7,
	0x21, 0xff, 0x1e, 0x36, 0x2e, 0x06, 0xa9, 0x4e, 0x06, 0xdd, 0xc8, 0x86, 0x77, 0xcb, 0xa0, 0x27,
	0x2e, 0x04, 0x99, 0xe4, 0x02, 0x86, 0x9e, 0x94, 0x9d, 0x62, 0xcc, 0xc7, 0xc7, 0x31, 0x2f, 0x71,
	0x79, 0x23, 0x12, 0x88, 0xcb, 0x51, 0xf8, 0x29, 0xcb, 0xd7, 0x90, 0xe9, 0xae, 0x3e, 0x0e, 0xd8,
	0xcf, 0xd4, 0xe5, 0xbd, 0xa0, 0x87, 0xc1, 0xe6, 0x32, 0xfa, 0xf5, 0x79, 0x2d, 0xec, 0xb6, 0xdf,
	0xcc, 0x4b, 0x2d, 0x73, 0xc5, 0xed, 0xa6, 0x63, 0x42, 0xe5, 0xf6, 0xea, 0xab, 0xba, 0x6b, 0xcb,
	0xbc, 0xd6, 0xf9, 0xd1, 0x78, 0x3e, 0x5a, 0xd6, 0xfa, 0x9d, 0x61, 0xfd, 0x77, 0x14, 0x6d, 0x86,
	0x38, 0xa9, 0xbf, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x26, 0x9c, 0x13, 0x08, 0x06, 0x00,
	0x00,
}
