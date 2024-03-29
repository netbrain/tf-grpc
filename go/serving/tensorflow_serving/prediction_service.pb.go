// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/prediction_service.proto

package tensorflow_serving // import "github.com/netbrain/tf-grpc/go/serving/tensorflow_serving"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Classify.
	Classify(ctx context.Context, in *ClassificationRequest, opts ...grpc.CallOption) (*ClassificationResponse, error)
	// Regress.
	Regress(ctx context.Context, in *RegressionRequest, opts ...grpc.CallOption) (*RegressionResponse, error)
	// Predict -- provides access to loaded TensorFlow model.
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
	// MultiInference API for multi-headed models.
	MultiInference(ctx context.Context, in *MultiInferenceRequest, opts ...grpc.CallOption) (*MultiInferenceResponse, error)
	// GetModelMetadata - provides access to metadata for loaded models.
	GetModelMetadata(ctx context.Context, in *GetModelMetadataRequest, opts ...grpc.CallOption) (*GetModelMetadataResponse, error)
}

type predictionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPredictionServiceClient(cc *grpc.ClientConn) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Classify(ctx context.Context, in *ClassificationRequest, opts ...grpc.CallOption) (*ClassificationResponse, error) {
	out := new(ClassificationResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Classify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) Regress(ctx context.Context, in *RegressionRequest, opts ...grpc.CallOption) (*RegressionResponse, error) {
	out := new(RegressionResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Regress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) MultiInference(ctx context.Context, in *MultiInferenceRequest, opts ...grpc.CallOption) (*MultiInferenceResponse, error) {
	out := new(MultiInferenceResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/MultiInference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) GetModelMetadata(ctx context.Context, in *GetModelMetadataRequest, opts ...grpc.CallOption) (*GetModelMetadataResponse, error) {
	out := new(GetModelMetadataResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/GetModelMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Classify.
	Classify(context.Context, *ClassificationRequest) (*ClassificationResponse, error)
	// Regress.
	Regress(context.Context, *RegressionRequest) (*RegressionResponse, error)
	// Predict -- provides access to loaded TensorFlow model.
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
	// MultiInference API for multi-headed models.
	MultiInference(context.Context, *MultiInferenceRequest) (*MultiInferenceResponse, error)
	// GetModelMetadata - provides access to metadata for loaded models.
	GetModelMetadata(context.Context, *GetModelMetadataRequest) (*GetModelMetadataResponse, error)
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Classify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Classify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Classify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Classify(ctx, req.(*ClassificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_Regress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Regress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Regress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Regress(ctx, req.(*RegressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_MultiInference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiInferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).MultiInference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/MultiInference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).MultiInference(ctx, req.(*MultiInferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_GetModelMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).GetModelMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/GetModelMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).GetModelMetadata(ctx, req.(*GetModelMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.serving.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Classify",
			Handler:    _PredictionService_Classify_Handler,
		},
		{
			MethodName: "Regress",
			Handler:    _PredictionService_Regress_Handler,
		},
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
		{
			MethodName: "MultiInference",
			Handler:    _PredictionService_MultiInference_Handler,
		},
		{
			MethodName: "GetModelMetadata",
			Handler:    _PredictionService_GetModelMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow_serving/apis/prediction_service.proto",
}

func init() {
	proto.RegisterFile("tensorflow_serving/apis/prediction_service.proto", fileDescriptor_prediction_service_74d07554922802e5)
}

var fileDescriptor_prediction_service_74d07554922802e5 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0x55, 0x7d, 0xa5, 0x6f, 0x91, 0x07, 0x04, 0x1e, 0x33, 0x82, 0xca, 0xcf, 0x12, 0x23,
	0x98, 0x58, 0x18, 0xe8, 0x80, 0x18, 0x2a, 0xa1, 0xb0, 0x20, 0x96, 0xca, 0x75, 0x2f, 0xc6, 0x52,
	0x62, 0x07, 0xfb, 0x02, 0xe2, 0x3f, 0x67, 0x60, 0x40, 0xc5, 0x76, 0x50, 0xdb, 0xa4, 0xed, 0x9a,
	0x7c, 0xde, 0xbb, 0xbb, 0xe7, 0x47, 0x2e, 0x11, 0xb4, 0x33, 0x36, 0x2f, 0xcc, 0xc7, 0xc4, 0x81,
	0x7d, 0x57, 0x5a, 0x32, 0x5e, 0x29, 0xc7, 0x2a, 0x0b, 0x33, 0x25, 0x50, 0x19, 0xed, 0xbf, 0x0b,
	0x48, 0x2b, 0x6b, 0xd0, 0x50, 0xfa, 0xa7, 0x48, 0x83, 0x22, 0x19, 0x76, 0xb9, 0x88, 0x82, 0x3b,
	0xa7, 0x72, 0x25, 0xf8, 0xdc, 0xc9, 0x3b, 0x24, 0x9d, 0x33, 0x25, 0xe0, 0xa4, 0x34, 0x33, 0x28,
	0x26, 0x25, 0x20, 0x9f, 0x71, 0xe4, 0x41, 0x71, 0xdc, 0xa5, 0x50, 0x3a, 0x07, 0x0b, 0x3a, 0x2e,
	0x97, 0x0c, 0x36, 0x9c, 0x13, 0xb0, 0x93, 0x2e, 0xcc, 0x82, 0xb4, 0xe0, 0x5c, 0xb3, 0xeb, 0xd5,
	0xf7, 0x3f, 0xb2, 0xff, 0xd8, 0x44, 0xf1, 0xe4, 0x93, 0xa0, 0x9c, 0xec, 0x8c, 0xfc, 0x65, 0x9f,
	0xf4, 0x34, 0x5d, 0x0d, 0x24, 0x1d, 0x2d, 0xdc, 0x9d, 0xc1, 0x5b, 0x0d, 0x0e, 0x93, 0xb3, 0x6d,
	0x50, 0x57, 0x19, 0xed, 0x80, 0x3e, 0x93, 0x7e, 0xe6, 0x97, 0xa1, 0x83, 0x36, 0x59, 0xd6, 0x6c,
	0x1a, 0xdd, 0x8f, 0x36, 0x61, 0xc1, 0x39, 0x23, 0xfd, 0x70, 0x11, 0x3d, 0x68, 0x93, 0x84, 0x9f,
	0xd1, 0xf6, 0x70, 0x2d, 0x13, 0x3c, 0x25, 0xd9, 0x1d, 0xd7, 0x05, 0xaa, 0x87, 0xf8, 0x1e, 0xed,
	0xb1, 0x2c, 0x32, 0x6b, 0x63, 0x59, 0x46, 0xc3, 0xa0, 0x92, 0xec, 0xdd, 0x03, 0x8e, 0xe7, 0x25,
	0x19, 0x87, 0x8e, 0xd0, 0xf3, 0x36, 0xfd, 0x32, 0x15, 0x87, 0x0d, 0xb7, 0x83, 0xfd, 0xb8, 0xbb,
	0xdb, 0x97, 0x1b, 0xa9, 0xf0, 0xb5, 0x9e, 0xa6, 0xc2, 0x94, 0x4c, 0x03, 0x4e, 0x2d, 0x57, 0x9a,
	0x61, 0x7e, 0x21, 0x6d, 0x25, 0x98, 0x34, 0x2c, 0xd6, 0x67, 0xb5, 0x51, 0x5f, 0xbd, 0xde, 0xf4,
	0xff, 0x6f, 0x8b, 0xae, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x31, 0x72, 0x0b, 0xad, 0x67, 0x03,
	0x00, 0x00,
}
