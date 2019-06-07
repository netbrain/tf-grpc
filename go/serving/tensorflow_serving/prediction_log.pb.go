// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/prediction_log.proto

package tensorflow_serving

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

type ClassifyLog struct {
	Request              *ClassificationRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *ClassificationResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClassifyLog) Reset()         { *m = ClassifyLog{} }
func (m *ClassifyLog) String() string { return proto.CompactTextString(m) }
func (*ClassifyLog) ProtoMessage()    {}
func (*ClassifyLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_db62cb1da263d301, []int{0}
}

func (m *ClassifyLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassifyLog.Unmarshal(m, b)
}
func (m *ClassifyLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassifyLog.Marshal(b, m, deterministic)
}
func (m *ClassifyLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassifyLog.Merge(m, src)
}
func (m *ClassifyLog) XXX_Size() int {
	return xxx_messageInfo_ClassifyLog.Size(m)
}
func (m *ClassifyLog) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassifyLog.DiscardUnknown(m)
}

var xxx_messageInfo_ClassifyLog proto.InternalMessageInfo

func (m *ClassifyLog) GetRequest() *ClassificationRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ClassifyLog) GetResponse() *ClassificationResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type RegressLog struct {
	Request              *RegressionRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *RegressionResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RegressLog) Reset()         { *m = RegressLog{} }
func (m *RegressLog) String() string { return proto.CompactTextString(m) }
func (*RegressLog) ProtoMessage()    {}
func (*RegressLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_db62cb1da263d301, []int{1}
}

func (m *RegressLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegressLog.Unmarshal(m, b)
}
func (m *RegressLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegressLog.Marshal(b, m, deterministic)
}
func (m *RegressLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegressLog.Merge(m, src)
}
func (m *RegressLog) XXX_Size() int {
	return xxx_messageInfo_RegressLog.Size(m)
}
func (m *RegressLog) XXX_DiscardUnknown() {
	xxx_messageInfo_RegressLog.DiscardUnknown(m)
}

var xxx_messageInfo_RegressLog proto.InternalMessageInfo

func (m *RegressLog) GetRequest() *RegressionRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *RegressLog) GetResponse() *RegressionResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type PredictLog struct {
	Request              *PredictRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *PredictResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PredictLog) Reset()         { *m = PredictLog{} }
func (m *PredictLog) String() string { return proto.CompactTextString(m) }
func (*PredictLog) ProtoMessage()    {}
func (*PredictLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_db62cb1da263d301, []int{2}
}

func (m *PredictLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictLog.Unmarshal(m, b)
}
func (m *PredictLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictLog.Marshal(b, m, deterministic)
}
func (m *PredictLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictLog.Merge(m, src)
}
func (m *PredictLog) XXX_Size() int {
	return xxx_messageInfo_PredictLog.Size(m)
}
func (m *PredictLog) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictLog.DiscardUnknown(m)
}

var xxx_messageInfo_PredictLog proto.InternalMessageInfo

func (m *PredictLog) GetRequest() *PredictRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *PredictLog) GetResponse() *PredictResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type MultiInferenceLog struct {
	Request              *MultiInferenceRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *MultiInferenceResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MultiInferenceLog) Reset()         { *m = MultiInferenceLog{} }
func (m *MultiInferenceLog) String() string { return proto.CompactTextString(m) }
func (*MultiInferenceLog) ProtoMessage()    {}
func (*MultiInferenceLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_db62cb1da263d301, []int{3}
}

func (m *MultiInferenceLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiInferenceLog.Unmarshal(m, b)
}
func (m *MultiInferenceLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiInferenceLog.Marshal(b, m, deterministic)
}
func (m *MultiInferenceLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiInferenceLog.Merge(m, src)
}
func (m *MultiInferenceLog) XXX_Size() int {
	return xxx_messageInfo_MultiInferenceLog.Size(m)
}
func (m *MultiInferenceLog) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiInferenceLog.DiscardUnknown(m)
}

var xxx_messageInfo_MultiInferenceLog proto.InternalMessageInfo

func (m *MultiInferenceLog) GetRequest() *MultiInferenceRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *MultiInferenceLog) GetResponse() *MultiInferenceResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type SessionRunLog struct {
	Request              *SessionRunRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *SessionRunResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SessionRunLog) Reset()         { *m = SessionRunLog{} }
func (m *SessionRunLog) String() string { return proto.CompactTextString(m) }
func (*SessionRunLog) ProtoMessage()    {}
func (*SessionRunLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_db62cb1da263d301, []int{4}
}

func (m *SessionRunLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRunLog.Unmarshal(m, b)
}
func (m *SessionRunLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRunLog.Marshal(b, m, deterministic)
}
func (m *SessionRunLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRunLog.Merge(m, src)
}
func (m *SessionRunLog) XXX_Size() int {
	return xxx_messageInfo_SessionRunLog.Size(m)
}
func (m *SessionRunLog) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRunLog.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRunLog proto.InternalMessageInfo

func (m *SessionRunLog) GetRequest() *SessionRunRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *SessionRunLog) GetResponse() *SessionRunResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

// Logged model inference request.
type PredictionLog struct {
	LogMetadata *LogMetadata `protobuf:"bytes,1,opt,name=log_metadata,json=logMetadata,proto3" json:"log_metadata,omitempty"`
	// Types that are valid to be assigned to LogType:
	//	*PredictionLog_ClassifyLog
	//	*PredictionLog_RegressLog
	//	*PredictionLog_PredictLog
	//	*PredictionLog_MultiInferenceLog
	//	*PredictionLog_SessionRunLog
	LogType              isPredictionLog_LogType `protobuf_oneof:"log_type"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PredictionLog) Reset()         { *m = PredictionLog{} }
func (m *PredictionLog) String() string { return proto.CompactTextString(m) }
func (*PredictionLog) ProtoMessage()    {}
func (*PredictionLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_db62cb1da263d301, []int{5}
}

func (m *PredictionLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionLog.Unmarshal(m, b)
}
func (m *PredictionLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionLog.Marshal(b, m, deterministic)
}
func (m *PredictionLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionLog.Merge(m, src)
}
func (m *PredictionLog) XXX_Size() int {
	return xxx_messageInfo_PredictionLog.Size(m)
}
func (m *PredictionLog) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionLog.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionLog proto.InternalMessageInfo

func (m *PredictionLog) GetLogMetadata() *LogMetadata {
	if m != nil {
		return m.LogMetadata
	}
	return nil
}

type isPredictionLog_LogType interface {
	isPredictionLog_LogType()
}

type PredictionLog_ClassifyLog struct {
	ClassifyLog *ClassifyLog `protobuf:"bytes,2,opt,name=classify_log,json=classifyLog,proto3,oneof"`
}

type PredictionLog_RegressLog struct {
	RegressLog *RegressLog `protobuf:"bytes,3,opt,name=regress_log,json=regressLog,proto3,oneof"`
}

type PredictionLog_PredictLog struct {
	PredictLog *PredictLog `protobuf:"bytes,6,opt,name=predict_log,json=predictLog,proto3,oneof"`
}

type PredictionLog_MultiInferenceLog struct {
	MultiInferenceLog *MultiInferenceLog `protobuf:"bytes,4,opt,name=multi_inference_log,json=multiInferenceLog,proto3,oneof"`
}

type PredictionLog_SessionRunLog struct {
	SessionRunLog *SessionRunLog `protobuf:"bytes,5,opt,name=session_run_log,json=sessionRunLog,proto3,oneof"`
}

func (*PredictionLog_ClassifyLog) isPredictionLog_LogType() {}

func (*PredictionLog_RegressLog) isPredictionLog_LogType() {}

func (*PredictionLog_PredictLog) isPredictionLog_LogType() {}

func (*PredictionLog_MultiInferenceLog) isPredictionLog_LogType() {}

func (*PredictionLog_SessionRunLog) isPredictionLog_LogType() {}

func (m *PredictionLog) GetLogType() isPredictionLog_LogType {
	if m != nil {
		return m.LogType
	}
	return nil
}

func (m *PredictionLog) GetClassifyLog() *ClassifyLog {
	if x, ok := m.GetLogType().(*PredictionLog_ClassifyLog); ok {
		return x.ClassifyLog
	}
	return nil
}

func (m *PredictionLog) GetRegressLog() *RegressLog {
	if x, ok := m.GetLogType().(*PredictionLog_RegressLog); ok {
		return x.RegressLog
	}
	return nil
}

func (m *PredictionLog) GetPredictLog() *PredictLog {
	if x, ok := m.GetLogType().(*PredictionLog_PredictLog); ok {
		return x.PredictLog
	}
	return nil
}

func (m *PredictionLog) GetMultiInferenceLog() *MultiInferenceLog {
	if x, ok := m.GetLogType().(*PredictionLog_MultiInferenceLog); ok {
		return x.MultiInferenceLog
	}
	return nil
}

func (m *PredictionLog) GetSessionRunLog() *SessionRunLog {
	if x, ok := m.GetLogType().(*PredictionLog_SessionRunLog); ok {
		return x.SessionRunLog
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PredictionLog) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PredictionLog_ClassifyLog)(nil),
		(*PredictionLog_RegressLog)(nil),
		(*PredictionLog_PredictLog)(nil),
		(*PredictionLog_MultiInferenceLog)(nil),
		(*PredictionLog_SessionRunLog)(nil),
	}
}

func init() {
	proto.RegisterType((*ClassifyLog)(nil), "tensorflow.serving.ClassifyLog")
	proto.RegisterType((*RegressLog)(nil), "tensorflow.serving.RegressLog")
	proto.RegisterType((*PredictLog)(nil), "tensorflow.serving.PredictLog")
	proto.RegisterType((*MultiInferenceLog)(nil), "tensorflow.serving.MultiInferenceLog")
	proto.RegisterType((*SessionRunLog)(nil), "tensorflow.serving.SessionRunLog")
	proto.RegisterType((*PredictionLog)(nil), "tensorflow.serving.PredictionLog")
}

func init() {
	proto.RegisterFile("tensorflow_serving/apis/prediction_log.proto", fileDescriptor_db62cb1da263d301)
}

var fileDescriptor_db62cb1da263d301 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbb, 0x14, 0x4a, 0x35, 0x6e, 0x84, 0x6a, 0x2e, 0x51, 0x0f, 0x14, 0x82, 0x0a, 0x05,
	0x51, 0x5b, 0x82, 0x13, 0x12, 0xa2, 0x6a, 0x82, 0x50, 0x11, 0xad, 0x54, 0x99, 0x03, 0x12, 0x97,
	0xc8, 0x71, 0x37, 0xcb, 0x4a, 0xce, 0xae, 0xd9, 0x5d, 0x83, 0xf2, 0x0c, 0x9c, 0x90, 0x38, 0x71,
	0xe2, 0x95, 0x78, 0x1b, 0x8e, 0xd5, 0xfe, 0x89, 0x9d, 0x66, 0x63, 0x37, 0xca, 0x31, 0xd6, 0xcc,
	0x6f, 0xbf, 0xcc, 0x37, 0xf3, 0xc1, 0x0b, 0x85, 0x99, 0xe4, 0x62, 0x9c, 0xf3, 0x1f, 0x43, 0x89,
	0xc5, 0x77, 0xca, 0x48, 0x9c, 0x16, 0x54, 0xc6, 0x85, 0xc0, 0x97, 0x34, 0x53, 0x94, 0xb3, 0x61,
	0xce, 0x49, 0x54, 0x08, 0xae, 0x78, 0x18, 0xd6, 0xd5, 0x91, 0xab, 0xde, 0x6b, 0x24, 0x64, 0x79,
	0x2a, 0x25, 0x1d, 0xd3, 0x2c, 0xd5, 0x14, 0x4b, 0xd8, 0x7b, 0xda, 0x54, 0x4d, 0xd9, 0x18, 0x0b,
	0xcc, 0x32, 0xec, 0x0a, 0x0f, 0x6e, 0x10, 0xe6, 0xca, 0x0e, 0x9b, 0xca, 0x04, 0x26, 0x02, 0x4b,
	0x59, 0xbf, 0x7c, 0xd4, 0x54, 0x29, 0x6d, 0x99, 0xfd, 0xd8, 0xfa, 0x7e, 0xc6, 0x05, 0x8e, 0x73,
	0x4e, 0x08, 0x65, 0x6e, 0x22, 0xbd, 0x3f, 0x08, 0x82, 0x81, 0xfd, 0xa3, 0xd3, 0x33, 0x4e, 0xc2,
	0x01, 0xdc, 0x15, 0xf8, 0x5b, 0x89, 0xa5, 0xea, 0xa2, 0x87, 0xe8, 0x30, 0x78, 0xf9, 0x2c, 0xf2,
	0x67, 0x16, 0x0d, 0xae, 0x8d, 0x26, 0xb1, 0x0d, 0xc9, 0xac, 0x33, 0x7c, 0x0f, 0xdb, 0x02, 0xcb,
	0x82, 0x33, 0x89, 0xbb, 0xb7, 0x0c, 0xe5, 0xf9, 0x2a, 0x14, 0xdb, 0x91, 0x54, 0xbd, 0xbd, 0x5f,
	0x08, 0x20, 0xb1, 0x73, 0xd0, 0xda, 0x8e, 0x17, 0xb5, 0x1d, 0x2c, 0xa3, 0x26, 0xd5, 0xe0, 0x3c,
	0x5d, 0x7d, 0x4f, 0xd7, 0x93, 0x9b, 0x08, 0x9e, 0xa6, 0x9f, 0x08, 0xe0, 0xc2, 0x5a, 0xa8, 0x35,
	0xbd, 0x59, 0xd4, 0xd4, 0x5b, 0x46, 0x74, 0x0d, 0x9e, 0xa0, 0x63, 0x4f, 0xd0, 0xe3, 0xd6, 0x76,
	0x4f, 0xcd, 0x5f, 0x04, 0xbb, 0xe7, 0x65, 0xae, 0xe8, 0x87, 0xd9, 0xfa, 0xad, 0x6e, 0xe2, 0xf5,
	0xbe, 0x75, 0x4d, 0x5c, 0xa4, 0x78, 0x12, 0x7f, 0x23, 0xe8, 0x7c, 0x72, 0xe3, 0x2c, 0xd9, 0xea,
	0x3e, 0xd6, 0x3d, 0xeb, 0xfa, 0x38, 0x4f, 0xf0, 0x64, 0xfd, 0xdb, 0x84, 0xce, 0x45, 0x95, 0x11,
	0x5a, 0x56, 0x1f, 0x76, 0x72, 0x4e, 0x86, 0x13, 0xac, 0xd2, 0xcb, 0x54, 0xa5, 0x4e, 0xdb, 0xfe,
	0x32, 0xf2, 0x19, 0x27, 0xe7, 0xae, 0x2c, 0x09, 0xf2, 0xfa, 0x47, 0xf8, 0x0e, 0x76, 0x5c, 0x6c,
	0x4c, 0x75, 0xec, 0x38, 0x75, 0xfb, 0x2d, 0xdb, 0xaf, 0xaf, 0xee, 0x74, 0x23, 0x09, 0xb2, 0xb9,
	0x23, 0x3c, 0x81, 0xc0, 0x9d, 0xbf, 0x81, 0x6c, 0x1a, 0xc8, 0x83, 0x96, 0x55, 0xb5, 0x0c, 0x10,
	0xf5, 0xad, 0x9c, 0x40, 0xe0, 0x82, 0xc6, 0x20, 0xb6, 0x9a, 0x11, 0xf5, 0x32, 0x6b, 0x44, 0x51,
	0xaf, 0xf6, 0x67, 0xb8, 0x3f, 0xd1, 0xe6, 0x0e, 0xab, 0x68, 0x33, 0xa8, 0xdb, 0xcd, 0x96, 0x79,
	0x9b, 0x78, 0xba, 0x91, 0xec, 0x4e, 0xbc, 0xf5, 0xfc, 0x08, 0xf7, 0x66, 0x99, 0x25, 0x4a, 0x13,
	0xcf, 0xdd, 0x3b, 0x06, 0xfa, 0xa8, 0xdd, 0x45, 0x0b, 0xec, 0xc8, 0xf9, 0x0f, 0x7d, 0x80, 0x6d,
	0xed, 0x9a, 0x9a, 0x16, 0xb8, 0xff, 0xf6, 0xcb, 0x6b, 0x42, 0xd5, 0xd7, 0x72, 0x14, 0x65, 0x7c,
	0x12, 0x33, 0xac, 0x46, 0x22, 0xa5, 0x2c, 0x56, 0xe3, 0x23, 0x22, 0x8a, 0x2c, 0x26, 0x3c, 0x9e,
	0x25, 0xa1, 0x1f, 0x8e, 0xff, 0x11, 0x1a, 0x6d, 0x99, 0x4c, 0x7c, 0x75, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x45, 0x84, 0x89, 0x53, 0x55, 0x06, 0x00, 0x00,
}
