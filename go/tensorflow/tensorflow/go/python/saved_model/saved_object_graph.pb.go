// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/python/saved_model/saved_object_graph.proto

package saved_model // import "github.com/netbrain/tf-grpc/go/tensorflow/tensorflow/go/python/saved_model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protobuf "github.com/netbrain/tf-grpc/go/tensorflow/tensorflow/go/core/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SavedObjectGraph struct {
	// List of objects in the SavedModel.
	//
	// The position of the object in this list indicates its id.
	// Nodes[0] is considered the root node.
	Nodes                []*SavedObject `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SavedObjectGraph) Reset()         { *m = SavedObjectGraph{} }
func (m *SavedObjectGraph) String() string { return proto.CompactTextString(m) }
func (*SavedObjectGraph) ProtoMessage()    {}
func (*SavedObjectGraph) Descriptor() ([]byte, []int) {
	return fileDescriptor_saved_object_graph_f5c2092bc03810fa, []int{0}
}
func (m *SavedObjectGraph) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedObjectGraph.Unmarshal(m, b)
}
func (m *SavedObjectGraph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedObjectGraph.Marshal(b, m, deterministic)
}
func (dst *SavedObjectGraph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedObjectGraph.Merge(dst, src)
}
func (m *SavedObjectGraph) XXX_Size() int {
	return xxx_messageInfo_SavedObjectGraph.Size(m)
}
func (m *SavedObjectGraph) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedObjectGraph.DiscardUnknown(m)
}

var xxx_messageInfo_SavedObjectGraph proto.InternalMessageInfo

func (m *SavedObjectGraph) GetNodes() []*SavedObject {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type SavedObject struct {
	// Objects which this object depends on: named edges in the dependency
	// graph.
	//
	// Note: only valid if kind == "object".
	Children []*protobuf.CheckpointableObjectGraph_CheckpointableObject_ObjectReference `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty"`
	// Slot variables owned by this object. This describes the three-way
	// (optimizer, variable, slot variable) relationship; none of the three
	// depend on the others directly.
	//
	// Note: only valid if kind == "object".
	SlotVariables []*protobuf.CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference `protobuf:"bytes,3,rep,name=slot_variables,json=slotVariables,proto3" json:"slot_variables,omitempty"`
	// Types that are valid to be assigned to Kind:
	//	*SavedObject_UserObject
	//	*SavedObject_Asset
	//	*SavedObject_Function
	Kind                 isSavedObject_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SavedObject) Reset()         { *m = SavedObject{} }
func (m *SavedObject) String() string { return proto.CompactTextString(m) }
func (*SavedObject) ProtoMessage()    {}
func (*SavedObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_saved_object_graph_f5c2092bc03810fa, []int{1}
}
func (m *SavedObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedObject.Unmarshal(m, b)
}
func (m *SavedObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedObject.Marshal(b, m, deterministic)
}
func (dst *SavedObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedObject.Merge(dst, src)
}
func (m *SavedObject) XXX_Size() int {
	return xxx_messageInfo_SavedObject.Size(m)
}
func (m *SavedObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedObject.DiscardUnknown(m)
}

var xxx_messageInfo_SavedObject proto.InternalMessageInfo

func (m *SavedObject) GetChildren() []*protobuf.CheckpointableObjectGraph_CheckpointableObject_ObjectReference {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *SavedObject) GetSlotVariables() []*protobuf.CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference {
	if m != nil {
		return m.SlotVariables
	}
	return nil
}

type isSavedObject_Kind interface {
	isSavedObject_Kind()
}

type SavedObject_UserObject struct {
	UserObject *SavedUserObject `protobuf:"bytes,4,opt,name=user_object,json=userObject,proto3,oneof"`
}

type SavedObject_Asset struct {
	Asset *SavedAsset `protobuf:"bytes,5,opt,name=asset,proto3,oneof"`
}

type SavedObject_Function struct {
	Function *SavedPolymorphicFunction `protobuf:"bytes,6,opt,name=function,proto3,oneof"`
}

func (*SavedObject_UserObject) isSavedObject_Kind() {}

func (*SavedObject_Asset) isSavedObject_Kind() {}

func (*SavedObject_Function) isSavedObject_Kind() {}

func (m *SavedObject) GetKind() isSavedObject_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *SavedObject) GetUserObject() *SavedUserObject {
	if x, ok := m.GetKind().(*SavedObject_UserObject); ok {
		return x.UserObject
	}
	return nil
}

func (m *SavedObject) GetAsset() *SavedAsset {
	if x, ok := m.GetKind().(*SavedObject_Asset); ok {
		return x.Asset
	}
	return nil
}

func (m *SavedObject) GetFunction() *SavedPolymorphicFunction {
	if x, ok := m.GetKind().(*SavedObject_Function); ok {
		return x.Function
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SavedObject) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SavedObject_OneofMarshaler, _SavedObject_OneofUnmarshaler, _SavedObject_OneofSizer, []interface{}{
		(*SavedObject_UserObject)(nil),
		(*SavedObject_Asset)(nil),
		(*SavedObject_Function)(nil),
	}
}

func _SavedObject_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SavedObject)
	// kind
	switch x := m.Kind.(type) {
	case *SavedObject_UserObject:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserObject); err != nil {
			return err
		}
	case *SavedObject_Asset:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Asset); err != nil {
			return err
		}
	case *SavedObject_Function:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Function); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SavedObject.Kind has unexpected type %T", x)
	}
	return nil
}

func _SavedObject_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SavedObject)
	switch tag {
	case 4: // kind.user_object
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SavedUserObject)
		err := b.DecodeMessage(msg)
		m.Kind = &SavedObject_UserObject{msg}
		return true, err
	case 5: // kind.asset
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SavedAsset)
		err := b.DecodeMessage(msg)
		m.Kind = &SavedObject_Asset{msg}
		return true, err
	case 6: // kind.function
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SavedPolymorphicFunction)
		err := b.DecodeMessage(msg)
		m.Kind = &SavedObject_Function{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SavedObject_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SavedObject)
	// kind
	switch x := m.Kind.(type) {
	case *SavedObject_UserObject:
		s := proto.Size(x.UserObject)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SavedObject_Asset:
		s := proto.Size(x.Asset)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SavedObject_Function:
		s := proto.Size(x.Function)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A SavedUserObject is an object (in the object-oriented language of the
// TensorFlow program) of some user- or framework-defined class other than
// those handled specifically by the other kinds of SavedObjects.
//
// This object cannot be evaluated as a tensor, and therefore cannot be bound
// to an input of a function.
type SavedUserObject struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SavedUserObject) Reset()         { *m = SavedUserObject{} }
func (m *SavedUserObject) String() string { return proto.CompactTextString(m) }
func (*SavedUserObject) ProtoMessage()    {}
func (*SavedUserObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_saved_object_graph_f5c2092bc03810fa, []int{2}
}
func (m *SavedUserObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedUserObject.Unmarshal(m, b)
}
func (m *SavedUserObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedUserObject.Marshal(b, m, deterministic)
}
func (dst *SavedUserObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedUserObject.Merge(dst, src)
}
func (m *SavedUserObject) XXX_Size() int {
	return xxx_messageInfo_SavedUserObject.Size(m)
}
func (m *SavedUserObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedUserObject.DiscardUnknown(m)
}

var xxx_messageInfo_SavedUserObject proto.InternalMessageInfo

// A SavedAsset represents a file in a SavedModel.
//
// When bound to a function this object evaluates to a Variable from which the
// absolute filename can be read. Users should not expect the filename to be
// maintained.
type SavedAsset struct {
	// Index into `MetaGraphDef.asset_file_def[]` that describes the Asset.
	//
	// Only the field `AssetFileDef.filename` is used. Other fields, such as
	// `AssetFileDef.tensor_info`, MUST be ignored.
	AssetFileDefIndex    uint32   `protobuf:"varint,1,opt,name=asset_file_def_index,json=assetFileDefIndex,proto3" json:"asset_file_def_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SavedAsset) Reset()         { *m = SavedAsset{} }
func (m *SavedAsset) String() string { return proto.CompactTextString(m) }
func (*SavedAsset) ProtoMessage()    {}
func (*SavedAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_saved_object_graph_f5c2092bc03810fa, []int{3}
}
func (m *SavedAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedAsset.Unmarshal(m, b)
}
func (m *SavedAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedAsset.Marshal(b, m, deterministic)
}
func (dst *SavedAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedAsset.Merge(dst, src)
}
func (m *SavedAsset) XXX_Size() int {
	return xxx_messageInfo_SavedAsset.Size(m)
}
func (m *SavedAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedAsset.DiscardUnknown(m)
}

var xxx_messageInfo_SavedAsset proto.InternalMessageInfo

func (m *SavedAsset) GetAssetFileDefIndex() uint32 {
	if m != nil {
		return m.AssetFileDefIndex
	}
	return 0
}

// A function with multiple signatures, possibly with non-Tensor arguments.
type SavedPolymorphicFunction struct {
	MonomorphicFunction  []*SavedMonomorphicFunction `protobuf:"bytes,1,rep,name=monomorphic_function,json=monomorphicFunction,proto3" json:"monomorphic_function,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SavedPolymorphicFunction) Reset()         { *m = SavedPolymorphicFunction{} }
func (m *SavedPolymorphicFunction) String() string { return proto.CompactTextString(m) }
func (*SavedPolymorphicFunction) ProtoMessage()    {}
func (*SavedPolymorphicFunction) Descriptor() ([]byte, []int) {
	return fileDescriptor_saved_object_graph_f5c2092bc03810fa, []int{4}
}
func (m *SavedPolymorphicFunction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedPolymorphicFunction.Unmarshal(m, b)
}
func (m *SavedPolymorphicFunction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedPolymorphicFunction.Marshal(b, m, deterministic)
}
func (dst *SavedPolymorphicFunction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedPolymorphicFunction.Merge(dst, src)
}
func (m *SavedPolymorphicFunction) XXX_Size() int {
	return xxx_messageInfo_SavedPolymorphicFunction.Size(m)
}
func (m *SavedPolymorphicFunction) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedPolymorphicFunction.DiscardUnknown(m)
}

var xxx_messageInfo_SavedPolymorphicFunction proto.InternalMessageInfo

func (m *SavedPolymorphicFunction) GetMonomorphicFunction() []*SavedMonomorphicFunction {
	if m != nil {
		return m.MonomorphicFunction
	}
	return nil
}

type SavedMonomorphicFunction struct {
	// A reference to a TensorFlow function in the MetaGraph's FunctionDefLibrary
	ConcreteFunction     string   `protobuf:"bytes,1,opt,name=concrete_function,json=concreteFunction,proto3" json:"concrete_function,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SavedMonomorphicFunction) Reset()         { *m = SavedMonomorphicFunction{} }
func (m *SavedMonomorphicFunction) String() string { return proto.CompactTextString(m) }
func (*SavedMonomorphicFunction) ProtoMessage()    {}
func (*SavedMonomorphicFunction) Descriptor() ([]byte, []int) {
	return fileDescriptor_saved_object_graph_f5c2092bc03810fa, []int{5}
}
func (m *SavedMonomorphicFunction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedMonomorphicFunction.Unmarshal(m, b)
}
func (m *SavedMonomorphicFunction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedMonomorphicFunction.Marshal(b, m, deterministic)
}
func (dst *SavedMonomorphicFunction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedMonomorphicFunction.Merge(dst, src)
}
func (m *SavedMonomorphicFunction) XXX_Size() int {
	return xxx_messageInfo_SavedMonomorphicFunction.Size(m)
}
func (m *SavedMonomorphicFunction) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedMonomorphicFunction.DiscardUnknown(m)
}

var xxx_messageInfo_SavedMonomorphicFunction proto.InternalMessageInfo

func (m *SavedMonomorphicFunction) GetConcreteFunction() string {
	if m != nil {
		return m.ConcreteFunction
	}
	return ""
}

func init() {
	proto.RegisterType((*SavedObjectGraph)(nil), "tensorflow.SavedObjectGraph")
	proto.RegisterType((*SavedObject)(nil), "tensorflow.SavedObject")
	proto.RegisterType((*SavedUserObject)(nil), "tensorflow.SavedUserObject")
	proto.RegisterType((*SavedAsset)(nil), "tensorflow.SavedAsset")
	proto.RegisterType((*SavedPolymorphicFunction)(nil), "tensorflow.SavedPolymorphicFunction")
	proto.RegisterType((*SavedMonomorphicFunction)(nil), "tensorflow.SavedMonomorphicFunction")
}

func init() {
	proto.RegisterFile("tensorflow/python/saved_model/saved_object_graph.proto", fileDescriptor_saved_object_graph_f5c2092bc03810fa)
}

var fileDescriptor_saved_object_graph_f5c2092bc03810fa = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcf, 0x8b, 0xd3, 0x4e,
	0x14, 0x6f, 0xbe, 0xfd, 0x41, 0xbf, 0xaf, 0xac, 0xb6, 0xe3, 0xa2, 0x41, 0x2f, 0x25, 0x78, 0x28,
	0xc8, 0x26, 0xb0, 0x82, 0x07, 0x41, 0x61, 0xab, 0xec, 0xd6, 0x82, 0xec, 0x92, 0x45, 0x05, 0x2f,
	0x21, 0x99, 0xbc, 0x34, 0xe3, 0xa6, 0xf3, 0xc2, 0xcc, 0x64, 0xd7, 0xfd, 0xdf, 0xfc, 0xc3, 0x3c,
	0x4a, 0xd3, 0xb4, 0x89, 0xa6, 0x9e, 0xbc, 0x4d, 0xde, 0xe7, 0xe7, 0x0c, 0x79, 0xf0, 0xca, 0xa0,
	0xd4, 0xa4, 0x92, 0x8c, 0xee, 0xbc, 0xfc, 0xde, 0xa4, 0x24, 0x3d, 0x1d, 0xde, 0x62, 0x1c, 0xac,
	0x29, 0xc6, 0xac, 0x3a, 0x53, 0xf4, 0x0d, 0xb9, 0x09, 0x56, 0x2a, 0xcc, 0x53, 0x37, 0x57, 0x64,
	0x88, 0x41, 0xad, 0x7b, 0xfa, 0xba, 0xe1, 0xc1, 0x49, 0xa1, 0x57, 0x52, 0xa2, 0x22, 0xf1, 0x78,
	0x8a, 0xfc, 0x26, 0x27, 0x21, 0x4d, 0x18, 0x65, 0x78, 0xc0, 0xc7, 0x39, 0x83, 0xf1, 0xf5, 0x26,
	0xe3, 0xb2, 0x84, 0x2e, 0x36, 0x08, 0x3b, 0x81, 0xbe, 0xa4, 0x18, 0xb5, 0x6d, 0x4d, 0xbb, 0xb3,
	0xd1, 0xe9, 0x13, 0xb7, 0xf6, 0x77, 0x1b, 0x64, 0x7f, 0xcb, 0x72, 0x7e, 0x74, 0x61, 0xd4, 0x18,
	0xb3, 0x04, 0x86, 0x3c, 0x15, 0x59, 0xac, 0x50, 0x56, 0x0e, 0xcb, 0xa6, 0xc3, 0xbb, 0xdf, 0x3a,
	0x35, 0x72, 0x0f, 0x22, 0x6e, 0x95, 0x85, 0x09, 0x2a, 0x94, 0x1c, 0xfd, 0xbd, 0x37, 0xbb, 0x83,
	0x07, 0x3a, 0x23, 0x13, 0xdc, 0x86, 0x4a, 0x6c, 0x04, 0xda, 0xee, 0x96, 0x69, 0x57, 0xff, 0x90,
	0x76, 0x9d, 0x91, 0xf9, 0x5c, 0xf9, 0xd5, 0x99, 0x47, 0xba, 0x31, 0xd6, 0xec, 0x2d, 0x8c, 0x0a,
	0x8d, 0xaa, 0x7a, 0x4e, 0xbb, 0x37, 0xb5, 0x66, 0xa3, 0xd3, 0x67, 0xad, 0x57, 0xfa, 0xa4, 0x51,
	0x6d, 0x6d, 0x17, 0x1d, 0x1f, 0x8a, 0xfd, 0x17, 0x73, 0xa1, 0x1f, 0x6a, 0x8d, 0xc6, 0xee, 0x97,
	0xca, 0xc7, 0x2d, 0xe5, 0xd9, 0x06, 0x5d, 0x74, 0xfc, 0x2d, 0x8d, 0xcd, 0x61, 0x98, 0x14, 0x92,
	0x1b, 0x41, 0xd2, 0x1e, 0x94, 0x92, 0xe7, 0x2d, 0xc9, 0x15, 0x65, 0xf7, 0x6b, 0x52, 0x79, 0x2a,
	0xf8, 0x79, 0xc5, 0x5d, 0x74, 0xfc, 0xbd, 0x6e, 0x3e, 0x80, 0xde, 0x8d, 0x90, 0xf1, 0xb2, 0x37,
	0xfc, 0x6f, 0xdc, 0xf5, 0x21, 0x34, 0x46, 0x89, 0xa8, 0x30, 0xa8, 0x9d, 0x09, 0x3c, 0xfc, 0xa3,
	0xae, 0xf3, 0x06, 0xa0, 0xee, 0xc1, 0x3c, 0x38, 0x2e, 0x7b, 0x04, 0x89, 0xc8, 0x30, 0x88, 0x31,
	0x09, 0x84, 0x8c, 0xf1, 0xbb, 0x6d, 0x4d, 0xad, 0xd9, 0x91, 0x3f, 0x29, 0xb1, 0x73, 0x91, 0xe1,
	0x7b, 0x4c, 0x3e, 0x6c, 0x00, 0x47, 0x83, 0xfd, 0xb7, 0x4e, 0xec, 0x0b, 0x1c, 0xaf, 0x49, 0x52,
	0x35, 0x0e, 0xf6, 0xf7, 0xda, 0xfe, 0x28, 0xed, 0x7b, 0x7d, 0xac, 0xc9, 0x3b, 0x0f, 0xff, 0xd1,
	0xba, 0x3d, 0x74, 0x2e, 0xaa, 0xd0, 0x03, 0x02, 0xf6, 0x02, 0x26, 0x9c, 0x24, 0x57, 0x68, 0xb0,
	0x99, 0x68, 0xcd, 0xfe, 0xf7, 0xc7, 0x3b, 0x60, 0x47, 0x9e, 0x5f, 0x7e, 0x5d, 0xae, 0x84, 0x49,
	0x8b, 0xc8, 0xe5, 0xb4, 0xf6, 0x24, 0x9a, 0x48, 0x85, 0x42, 0x7a, 0x26, 0x39, 0x59, 0xa9, 0x9c,
	0x7b, 0x2b, 0xf2, 0x1a, 0xeb, 0xd6, 0x38, 0xae, 0xe8, 0xc0, 0x02, 0xff, 0xb4, 0xac, 0x68, 0x50,
	0x6e, 0xda, 0xcb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x52, 0xe2, 0x3f, 0x30, 0xeb, 0x03, 0x00,
	0x00,
}
