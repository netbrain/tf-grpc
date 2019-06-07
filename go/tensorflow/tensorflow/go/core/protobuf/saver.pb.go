// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/saver.proto

package protobuf

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

// A version number that identifies a different on-disk checkpoint format.
// Usually, each subclass of BaseSaverBuilder works with a particular
// version/format.  However, it is possible that the same builder may be
// upgraded to support a newer checkpoint format in the future.
type SaverDef_CheckpointFormatVersion int32

const (
	// Internal legacy format.
	SaverDef_LEGACY SaverDef_CheckpointFormatVersion = 0
	// Deprecated format: tf.Saver() which works with tensorflow::table::Table.
	SaverDef_V1 SaverDef_CheckpointFormatVersion = 1
	// Current format: more efficient.
	SaverDef_V2 SaverDef_CheckpointFormatVersion = 2
)

var SaverDef_CheckpointFormatVersion_name = map[int32]string{
	0: "LEGACY",
	1: "V1",
	2: "V2",
}

var SaverDef_CheckpointFormatVersion_value = map[string]int32{
	"LEGACY": 0,
	"V1":     1,
	"V2":     2,
}

func (x SaverDef_CheckpointFormatVersion) String() string {
	return proto.EnumName(SaverDef_CheckpointFormatVersion_name, int32(x))
}

func (SaverDef_CheckpointFormatVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5551ea1a7581c104, []int{0, 0}
}

// Protocol buffer representing the configuration of a Saver.
type SaverDef struct {
	// The name of the tensor in which to specify the filename when saving or
	// restoring a model checkpoint.
	FilenameTensorName string `protobuf:"bytes,1,opt,name=filename_tensor_name,json=filenameTensorName,proto3" json:"filename_tensor_name,omitempty"`
	// The operation to run when saving a model checkpoint.
	SaveTensorName string `protobuf:"bytes,2,opt,name=save_tensor_name,json=saveTensorName,proto3" json:"save_tensor_name,omitempty"`
	// The operation to run when restoring a model checkpoint.
	RestoreOpName string `protobuf:"bytes,3,opt,name=restore_op_name,json=restoreOpName,proto3" json:"restore_op_name,omitempty"`
	// Maximum number of checkpoints to keep.  If 0, no checkpoints are deleted.
	MaxToKeep int32 `protobuf:"varint,4,opt,name=max_to_keep,json=maxToKeep,proto3" json:"max_to_keep,omitempty"`
	// Shard the save files, one per device that has Variable nodes.
	Sharded bool `protobuf:"varint,5,opt,name=sharded,proto3" json:"sharded,omitempty"`
	// How often to keep an additional checkpoint. If not specified, only the last
	// "max_to_keep" checkpoints are kept; if specified, in addition to keeping
	// the last "max_to_keep" checkpoints, an additional checkpoint will be kept
	// for every n hours of training.
	KeepCheckpointEveryNHours float32                          `protobuf:"fixed32,6,opt,name=keep_checkpoint_every_n_hours,json=keepCheckpointEveryNHours,proto3" json:"keep_checkpoint_every_n_hours,omitempty"`
	Version                   SaverDef_CheckpointFormatVersion `protobuf:"varint,7,opt,name=version,proto3,enum=tensorflow.SaverDef_CheckpointFormatVersion" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                         `json:"-"`
	XXX_unrecognized          []byte                           `json:"-"`
	XXX_sizecache             int32                            `json:"-"`
}

func (m *SaverDef) Reset()         { *m = SaverDef{} }
func (m *SaverDef) String() string { return proto.CompactTextString(m) }
func (*SaverDef) ProtoMessage()    {}
func (*SaverDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_5551ea1a7581c104, []int{0}
}

func (m *SaverDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaverDef.Unmarshal(m, b)
}
func (m *SaverDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaverDef.Marshal(b, m, deterministic)
}
func (m *SaverDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaverDef.Merge(m, src)
}
func (m *SaverDef) XXX_Size() int {
	return xxx_messageInfo_SaverDef.Size(m)
}
func (m *SaverDef) XXX_DiscardUnknown() {
	xxx_messageInfo_SaverDef.DiscardUnknown(m)
}

var xxx_messageInfo_SaverDef proto.InternalMessageInfo

func (m *SaverDef) GetFilenameTensorName() string {
	if m != nil {
		return m.FilenameTensorName
	}
	return ""
}

func (m *SaverDef) GetSaveTensorName() string {
	if m != nil {
		return m.SaveTensorName
	}
	return ""
}

func (m *SaverDef) GetRestoreOpName() string {
	if m != nil {
		return m.RestoreOpName
	}
	return ""
}

func (m *SaverDef) GetMaxToKeep() int32 {
	if m != nil {
		return m.MaxToKeep
	}
	return 0
}

func (m *SaverDef) GetSharded() bool {
	if m != nil {
		return m.Sharded
	}
	return false
}

func (m *SaverDef) GetKeepCheckpointEveryNHours() float32 {
	if m != nil {
		return m.KeepCheckpointEveryNHours
	}
	return 0
}

func (m *SaverDef) GetVersion() SaverDef_CheckpointFormatVersion {
	if m != nil {
		return m.Version
	}
	return SaverDef_LEGACY
}

func init() {
	proto.RegisterEnum("tensorflow.SaverDef_CheckpointFormatVersion", SaverDef_CheckpointFormatVersion_name, SaverDef_CheckpointFormatVersion_value)
	proto.RegisterType((*SaverDef)(nil), "tensorflow.SaverDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/saver.proto", fileDescriptor_5551ea1a7581c104)
}

var fileDescriptor_5551ea1a7581c104 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x71, 0x96, 0x4d, 0x77, 0xbd, 0x62, 0x89, 0x0c, 0x12, 0xe1, 0x00, 0x8a, 0x56, 0x08,
	0xe5, 0x00, 0x09, 0x2c, 0xe2, 0x0e, 0x5b, 0x5a, 0x90, 0x40, 0xa5, 0x0a, 0x55, 0x25, 0xb8, 0x58,
	0x4e, 0x3a, 0xf9, 0xa3, 0x36, 0x9e, 0xc8, 0x71, 0x4a, 0xf9, 0x08, 0x7c, 0x63, 0x8e, 0xc8, 0x2e,
	0xa1, 0xe9, 0x61, 0x4f, 0x99, 0x37, 0xef, 0xf7, 0x32, 0xd2, 0x78, 0xe8, 0x33, 0x0d, 0xb2, 0x45,
	0x95, 0x6f, 0xf0, 0x67, 0x9c, 0xa1, 0x82, 0xb8, 0x51, 0xa8, 0x31, 0xed, 0xf2, 0xb8, 0x15, 0x5b,
	0x50, 0x91, 0x95, 0x8c, 0x1e, 0xa8, 0xab, 0xdf, 0x27, 0xf4, 0xec, 0x9b, 0xf1, 0x3e, 0x40, 0xce,
	0x5e, 0xd1, 0x87, 0x79, 0xb5, 0x01, 0x29, 0x6a, 0xe0, 0x7b, 0x86, 0x9b, 0xda, 0x27, 0x01, 0x09,
	0xcf, 0x13, 0xd6, 0x7b, 0x0b, 0x6b, 0xcd, 0x44, 0x0d, 0x2c, 0xa4, 0x9e, 0xf9, 0xf3, 0x11, 0xed,
	0x58, 0xfa, 0xd2, 0xf4, 0x07, 0xe4, 0x73, 0x7a, 0x5f, 0x41, 0xab, 0x51, 0x01, 0xc7, 0x66, 0x0f,
	0x9e, 0x58, 0xf0, 0xde, 0xbf, 0xf6, 0xd7, 0xc6, 0x72, 0x4f, 0xe9, 0x45, 0x2d, 0x76, 0x5c, 0x23,
	0x5f, 0x03, 0x34, 0xfe, 0xdd, 0x80, 0x84, 0xa7, 0xc9, 0x79, 0x2d, 0x76, 0x0b, 0xfc, 0x0c, 0xd0,
	0x30, 0x9f, 0x8e, 0xda, 0x52, 0xa8, 0x15, 0xac, 0xfc, 0xd3, 0x80, 0x84, 0x67, 0x49, 0x2f, 0xd9,
	0x3b, 0xfa, 0xc4, 0x44, 0x78, 0x56, 0x42, 0xb6, 0x6e, 0xb0, 0x92, 0x9a, 0xc3, 0x16, 0xd4, 0x2f,
	0x2e, 0x79, 0x89, 0x9d, 0x6a, 0x7d, 0x37, 0x20, 0xa1, 0x93, 0x3c, 0x36, 0xd0, 0xf8, 0x3f, 0x33,
	0x31, 0xc8, 0xec, 0x93, 0x01, 0xd8, 0x94, 0x8e, 0xb6, 0xa0, 0xda, 0x0a, 0xa5, 0x3f, 0x0a, 0x48,
	0x78, 0x79, 0xfd, 0x22, 0x3a, 0xac, 0x2a, 0xea, 0xd7, 0x14, 0x1d, 0xc2, 0x53, 0x54, 0xb5, 0xd0,
	0xcb, 0x7d, 0x26, 0xe9, 0xc3, 0x57, 0x6f, 0xe9, 0xa3, 0x5b, 0x18, 0x46, 0xa9, 0xfb, 0x65, 0xf2,
	0xf1, 0xfd, 0xf8, 0xbb, 0x77, 0x87, 0xb9, 0xd4, 0x59, 0xbe, 0xf6, 0x88, 0xfd, 0x5e, 0x7b, 0xce,
	0x8d, 0xa4, 0x0f, 0x50, 0x15, 0xc3, 0x91, 0x9d, 0xae, 0x36, 0x37, 0x17, 0x76, 0xf0, 0xdc, 0x3c,
	0x5d, 0x3b, 0x27, 0x3f, 0x26, 0x45, 0xa5, 0xcb, 0x2e, 0x8d, 0x32, 0xac, 0x63, 0x09, 0x3a, 0x55,
	0xa2, 0x92, 0xb1, 0xce, 0x5f, 0x16, 0xaa, 0xc9, 0xe2, 0x02, 0xe3, 0xc1, 0x09, 0x0c, 0xca, 0x02,
	0x8f, 0x0f, 0xe2, 0x0f, 0x21, 0xa9, 0x6b, 0xc5, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x60,
	0x49, 0x0d, 0x56, 0x36, 0x02, 0x00, 0x00,
}
