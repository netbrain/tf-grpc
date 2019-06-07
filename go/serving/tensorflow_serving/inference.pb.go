// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/inference.proto

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

// Inference request such as classification, regression, etc...
type InferenceTask struct {
	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	// All ModelSpecs in a MultiInferenceRequest must access the same model name.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Signature's method_name. Should be one of the method names defined in
	// third_party/tensorflow/python/saved_model/signature_constants.py.
	// e.g. "tensorflow/serving/classify".
	MethodName           string   `protobuf:"bytes,2,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InferenceTask) Reset()         { *m = InferenceTask{} }
func (m *InferenceTask) String() string { return proto.CompactTextString(m) }
func (*InferenceTask) ProtoMessage()    {}
func (*InferenceTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_4842740eef0fb1a6, []int{0}
}

func (m *InferenceTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferenceTask.Unmarshal(m, b)
}
func (m *InferenceTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferenceTask.Marshal(b, m, deterministic)
}
func (m *InferenceTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferenceTask.Merge(m, src)
}
func (m *InferenceTask) XXX_Size() int {
	return xxx_messageInfo_InferenceTask.Size(m)
}
func (m *InferenceTask) XXX_DiscardUnknown() {
	xxx_messageInfo_InferenceTask.DiscardUnknown(m)
}

var xxx_messageInfo_InferenceTask proto.InternalMessageInfo

func (m *InferenceTask) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *InferenceTask) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

// Inference result, matches the type of request or is an error.
type InferenceResult struct {
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Types that are valid to be assigned to Result:
	//	*InferenceResult_ClassificationResult
	//	*InferenceResult_RegressionResult
	Result               isInferenceResult_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *InferenceResult) Reset()         { *m = InferenceResult{} }
func (m *InferenceResult) String() string { return proto.CompactTextString(m) }
func (*InferenceResult) ProtoMessage()    {}
func (*InferenceResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4842740eef0fb1a6, []int{1}
}

func (m *InferenceResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferenceResult.Unmarshal(m, b)
}
func (m *InferenceResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferenceResult.Marshal(b, m, deterministic)
}
func (m *InferenceResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferenceResult.Merge(m, src)
}
func (m *InferenceResult) XXX_Size() int {
	return xxx_messageInfo_InferenceResult.Size(m)
}
func (m *InferenceResult) XXX_DiscardUnknown() {
	xxx_messageInfo_InferenceResult.DiscardUnknown(m)
}

var xxx_messageInfo_InferenceResult proto.InternalMessageInfo

func (m *InferenceResult) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

type isInferenceResult_Result interface {
	isInferenceResult_Result()
}

type InferenceResult_ClassificationResult struct {
	ClassificationResult *ClassificationResult `protobuf:"bytes,2,opt,name=classification_result,json=classificationResult,proto3,oneof"`
}

type InferenceResult_RegressionResult struct {
	RegressionResult *RegressionResult `protobuf:"bytes,3,opt,name=regression_result,json=regressionResult,proto3,oneof"`
}

func (*InferenceResult_ClassificationResult) isInferenceResult_Result() {}

func (*InferenceResult_RegressionResult) isInferenceResult_Result() {}

func (m *InferenceResult) GetResult() isInferenceResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *InferenceResult) GetClassificationResult() *ClassificationResult {
	if x, ok := m.GetResult().(*InferenceResult_ClassificationResult); ok {
		return x.ClassificationResult
	}
	return nil
}

func (m *InferenceResult) GetRegressionResult() *RegressionResult {
	if x, ok := m.GetResult().(*InferenceResult_RegressionResult); ok {
		return x.RegressionResult
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*InferenceResult) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*InferenceResult_ClassificationResult)(nil),
		(*InferenceResult_RegressionResult)(nil),
	}
}

// Inference request containing one or more requests.
type MultiInferenceRequest struct {
	// Inference tasks.
	Tasks []*InferenceTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	// Input data.
	Input                *Input   `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiInferenceRequest) Reset()         { *m = MultiInferenceRequest{} }
func (m *MultiInferenceRequest) String() string { return proto.CompactTextString(m) }
func (*MultiInferenceRequest) ProtoMessage()    {}
func (*MultiInferenceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4842740eef0fb1a6, []int{2}
}

func (m *MultiInferenceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiInferenceRequest.Unmarshal(m, b)
}
func (m *MultiInferenceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiInferenceRequest.Marshal(b, m, deterministic)
}
func (m *MultiInferenceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiInferenceRequest.Merge(m, src)
}
func (m *MultiInferenceRequest) XXX_Size() int {
	return xxx_messageInfo_MultiInferenceRequest.Size(m)
}
func (m *MultiInferenceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiInferenceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiInferenceRequest proto.InternalMessageInfo

func (m *MultiInferenceRequest) GetTasks() []*InferenceTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *MultiInferenceRequest) GetInput() *Input {
	if m != nil {
		return m.Input
	}
	return nil
}

// Inference request containing one or more responses.
type MultiInferenceResponse struct {
	// List of results; one for each InferenceTask in the request, returned in the
	// same order as the request.
	Results              []*InferenceResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MultiInferenceResponse) Reset()         { *m = MultiInferenceResponse{} }
func (m *MultiInferenceResponse) String() string { return proto.CompactTextString(m) }
func (*MultiInferenceResponse) ProtoMessage()    {}
func (*MultiInferenceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4842740eef0fb1a6, []int{3}
}

func (m *MultiInferenceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiInferenceResponse.Unmarshal(m, b)
}
func (m *MultiInferenceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiInferenceResponse.Marshal(b, m, deterministic)
}
func (m *MultiInferenceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiInferenceResponse.Merge(m, src)
}
func (m *MultiInferenceResponse) XXX_Size() int {
	return xxx_messageInfo_MultiInferenceResponse.Size(m)
}
func (m *MultiInferenceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiInferenceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MultiInferenceResponse proto.InternalMessageInfo

func (m *MultiInferenceResponse) GetResults() []*InferenceResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*InferenceTask)(nil), "tensorflow.serving.InferenceTask")
	proto.RegisterType((*InferenceResult)(nil), "tensorflow.serving.InferenceResult")
	proto.RegisterType((*MultiInferenceRequest)(nil), "tensorflow.serving.MultiInferenceRequest")
	proto.RegisterType((*MultiInferenceResponse)(nil), "tensorflow.serving.MultiInferenceResponse")
}

func init() {
	proto.RegisterFile("tensorflow_serving/apis/inference.proto", fileDescriptor_4842740eef0fb1a6)
}

var fileDescriptor_4842740eef0fb1a6 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x8f, 0xd3, 0x30,
	0x10, 0x85, 0x49, 0xab, 0x16, 0xea, 0x08, 0x01, 0x16, 0x45, 0xa5, 0x12, 0xa2, 0xa4, 0x48, 0xe4,
	0x00, 0x89, 0x54, 0x0e, 0x08, 0x09, 0x38, 0x94, 0x0b, 0x1c, 0xca, 0xc1, 0x45, 0x42, 0xe2, 0x12,
	0xb9, 0xee, 0x24, 0xb5, 0x9a, 0xd8, 0xc6, 0x76, 0xe0, 0xcc, 0x2f, 0xe0, 0xef, 0x72, 0x5c, 0x6d,
	0xdc, 0xa4, 0xdb, 0x36, 0xdd, 0x3d, 0xec, 0xcd, 0x1a, 0xbd, 0xf9, 0xde, 0x9b, 0xf1, 0xa0, 0x57,
	0x16, 0x84, 0x91, 0x3a, 0xcd, 0xe5, 0x9f, 0xc4, 0x80, 0xfe, 0xcd, 0x45, 0x16, 0x53, 0xc5, 0x4d,
	0xcc, 0x45, 0x0a, 0x1a, 0x04, 0x83, 0x48, 0x69, 0x69, 0x25, 0xc6, 0x7b, 0x61, 0xb4, 0x13, 0x8e,
	0x5f, 0x9f, 0x6b, 0x66, 0x39, 0x35, 0x86, 0xa7, 0x9c, 0x51, 0xcb, 0xa5, 0x70, 0x84, 0xf1, 0xf4,
	0xbc, 0x95, 0x2a, 0xed, 0x4d, 0xa2, 0x42, 0xae, 0x21, 0xdf, 0x89, 0xc2, 0x73, 0x22, 0x0d, 0x99,
	0x06, 0x63, 0x1a, 0xcf, 0x40, 0xa0, 0xfb, 0x5f, 0xeb, 0x41, 0xbe, 0x53, 0xb3, 0xc5, 0x1f, 0x10,
	0xaa, 0x48, 0x89, 0x51, 0xc0, 0x46, 0xde, 0xc4, 0x0b, 0xfd, 0xd9, 0xb3, 0xe8, 0x74, 0xb6, 0x68,
	0x71, 0xa9, 0x5a, 0x2a, 0x60, 0x64, 0x50, 0xd4, 0x4f, 0xfc, 0x1c, 0xf9, 0x05, 0xd8, 0x8d, 0x5c,
	0x27, 0x82, 0x16, 0x30, 0xea, 0x4c, 0xbc, 0x70, 0x40, 0x90, 0x2b, 0x7d, 0xa3, 0x05, 0x04, 0xff,
	0x3a, 0xe8, 0x41, 0x63, 0x48, 0xc0, 0x94, 0xb9, 0xbd, 0xa5, 0x65, 0x82, 0x86, 0x87, 0xdb, 0x4c,
	0x74, 0x85, 0xad, 0xcc, 0xfd, 0x59, 0xd8, 0x06, 0xfa, 0x7c, 0xd0, 0xe0, 0x62, 0x7c, 0xb9, 0x43,
	0x1e, 0xb3, 0x96, 0x3a, 0x5e, 0xa2, 0x47, 0xfb, 0xb5, 0xd5, 0xf0, 0x6e, 0x05, 0x7f, 0xd9, 0x06,
	0x27, 0x8d, 0xb8, 0x01, 0x3f, 0xd4, 0x47, 0xb5, 0xf9, 0x3d, 0xd4, 0x77, 0xa4, 0xe0, 0xaf, 0x87,
	0x86, 0x8b, 0x32, 0xb7, 0xfc, 0xca, 0x5a, 0x7e, 0x95, 0x60, 0x2c, 0x7e, 0x87, 0x7a, 0x96, 0x9a,
	0xad, 0x19, 0x79, 0x93, 0x6e, 0xe8, 0xcf, 0x5e, 0xb4, 0x99, 0x1d, 0x7c, 0x1e, 0x71, 0x7a, 0x1c,
	0xa3, 0x5e, 0x75, 0x32, 0xbb, 0x15, 0x3c, 0x6d, 0x6f, 0x54, 0xa5, 0x25, 0x4e, 0x17, 0xfc, 0x40,
	0x4f, 0x8e, 0x23, 0x18, 0x25, 0x85, 0x01, 0xfc, 0x11, 0xdd, 0x75, 0x39, 0xeb, 0x14, 0xd3, 0x6b,
	0x53, 0xb8, 0xe9, 0x48, 0xdd, 0x33, 0xff, 0xf4, 0xf3, 0x7d, 0xc6, 0xed, 0xa6, 0x5c, 0x45, 0x4c,
	0x16, 0xb1, 0x00, 0xbb, 0xd2, 0x94, 0x8b, 0xd8, 0xa6, 0x6f, 0x32, 0xad, 0x58, 0x9c, 0xc9, 0xb8,
	0x3e, 0xcf, 0xd3, 0x8b, 0xfd, 0xef, 0x79, 0xab, 0x7e, 0x75, 0xa5, 0x6f, 0x2f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xe9, 0xb0, 0x05, 0x18, 0x86, 0x03, 0x00, 0x00,
}