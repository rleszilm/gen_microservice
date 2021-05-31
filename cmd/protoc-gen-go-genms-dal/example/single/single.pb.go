// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: single.proto

package single

import (
	_ "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	_ "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations/types"
	_ "google.golang.org/genproto/googleapis/type/latlng"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Single_Enum int32

const (
	Single_EnumOne Single_Enum = 0
	Single_EnumTwo Single_Enum = 1
)

// Enum value maps for Single_Enum.
var (
	Single_Enum_name = map[int32]string{
		0: "EnumOne",
		1: "EnumTwo",
	}
	Single_Enum_value = map[string]int32{
		"EnumOne": 0,
		"EnumTwo": 1,
	}
)

func (x Single_Enum) Enum() *Single_Enum {
	p := new(Single_Enum)
	*p = x
	return p
}

func (x Single_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Single_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_single_proto_enumTypes[0].Descriptor()
}

func (Single_Enum) Type() protoreflect.EnumType {
	return &file_single_proto_enumTypes[0]
}

func (x Single_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Single_Enum.Descriptor instead.
func (Single_Enum) EnumDescriptor() ([]byte, []int) {
	return file_single_proto_rawDescGZIP(), []int{0, 0}
}

type Single struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScalarInt32     int32           `protobuf:"varint,1,opt,name=scalar_int32,json=scalarInt32,proto3" json:"scalar_int32,omitempty"`
	ScalarInt64     int64           `protobuf:"varint,2,opt,name=scalar_int64,json=scalarInt64,proto3" json:"scalar_int64,omitempty"`
	ScalarFloat32   float32         `protobuf:"fixed32,3,opt,name=scalar_float32,json=scalarFloat32,proto3" json:"scalar_float32,omitempty"`
	ScalarFloat64   float64         `protobuf:"fixed64,4,opt,name=scalar_float64,json=scalarFloat64,proto3" json:"scalar_float64,omitempty"`
	ScalarString    string          `protobuf:"bytes,5,opt,name=scalar_string,json=scalarString,proto3" json:"scalar_string,omitempty"`
	ScalarBool      bool            `protobuf:"varint,6,opt,name=scalar_bool,json=scalarBool,proto3" json:"scalar_bool,omitempty"`
	ScalarEnum      Single_Enum     `protobuf:"varint,7,opt,name=scalar_enum,json=scalarEnum,proto3,enum=greeter.Single_Enum" json:"scalar_enum,omitempty"`
	ObjMessage      *Single_Message `protobuf:"bytes,8,opt,name=obj_message,json=objMessage,proto3" json:"obj_message,omitempty"`
	Ignored         string          `protobuf:"bytes,9,opt,name=ignored,proto3" json:"ignored,omitempty"`
	Renamed         string          `protobuf:"bytes,10,opt,name=renamed,proto3" json:"renamed,omitempty"`
	IgnoredPostgres string          `protobuf:"bytes,11,opt,name=ignored_postgres,json=ignoredPostgres,proto3" json:"ignored_postgres,omitempty"`
	RenamedPostgres string          `protobuf:"bytes,12,opt,name=renamed_postgres,json=renamedPostgres,proto3" json:"renamed_postgres,omitempty"`
	IgnoredRest     string          `protobuf:"bytes,13,opt,name=ignored_rest,json=ignoredRest,proto3" json:"ignored_rest,omitempty"`
	RenamedRest     string          `protobuf:"bytes,14,opt,name=renamed_rest,json=renamedRest,proto3" json:"renamed_rest,omitempty"`
	IgnoredMongo    string          `protobuf:"bytes,15,opt,name=ignored_mongo,json=ignoredMongo,proto3" json:"ignored_mongo,omitempty"`
	RenamedMongo    string          `protobuf:"bytes,16,opt,name=renamed_mongo,json=renamedMongo,proto3" json:"renamed_mongo,omitempty"`
}

func (x *Single) Reset() {
	*x = Single{}
	if protoimpl.UnsafeEnabled {
		mi := &file_single_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Single) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Single) ProtoMessage() {}

func (x *Single) ProtoReflect() protoreflect.Message {
	mi := &file_single_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Single.ProtoReflect.Descriptor instead.
func (*Single) Descriptor() ([]byte, []int) {
	return file_single_proto_rawDescGZIP(), []int{0}
}

func (x *Single) GetScalarInt32() int32 {
	if x != nil {
		return x.ScalarInt32
	}
	return 0
}

func (x *Single) GetScalarInt64() int64 {
	if x != nil {
		return x.ScalarInt64
	}
	return 0
}

func (x *Single) GetScalarFloat32() float32 {
	if x != nil {
		return x.ScalarFloat32
	}
	return 0
}

func (x *Single) GetScalarFloat64() float64 {
	if x != nil {
		return x.ScalarFloat64
	}
	return 0
}

func (x *Single) GetScalarString() string {
	if x != nil {
		return x.ScalarString
	}
	return ""
}

func (x *Single) GetScalarBool() bool {
	if x != nil {
		return x.ScalarBool
	}
	return false
}

func (x *Single) GetScalarEnum() Single_Enum {
	if x != nil {
		return x.ScalarEnum
	}
	return Single_EnumOne
}

func (x *Single) GetObjMessage() *Single_Message {
	if x != nil {
		return x.ObjMessage
	}
	return nil
}

func (x *Single) GetIgnored() string {
	if x != nil {
		return x.Ignored
	}
	return ""
}

func (x *Single) GetRenamed() string {
	if x != nil {
		return x.Renamed
	}
	return ""
}

func (x *Single) GetIgnoredPostgres() string {
	if x != nil {
		return x.IgnoredPostgres
	}
	return ""
}

func (x *Single) GetRenamedPostgres() string {
	if x != nil {
		return x.RenamedPostgres
	}
	return ""
}

func (x *Single) GetIgnoredRest() string {
	if x != nil {
		return x.IgnoredRest
	}
	return ""
}

func (x *Single) GetRenamedRest() string {
	if x != nil {
		return x.RenamedRest
	}
	return ""
}

func (x *Single) GetIgnoredMongo() string {
	if x != nil {
		return x.IgnoredMongo
	}
	return ""
}

func (x *Single) GetRenamedMongo() string {
	if x != nil {
		return x.RenamedMongo
	}
	return ""
}

type Single_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Single_Message) Reset() {
	*x = Single_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_single_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Single_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Single_Message) ProtoMessage() {}

func (x *Single_Message) ProtoReflect() protoreflect.Message {
	mi := &file_single_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Single_Message.ProtoReflect.Descriptor instead.
func (*Single_Message) Descriptor() ([]byte, []int) {
	return file_single_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Single_Message) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_single_proto protoreflect.FileDescriptor

var file_single_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x35, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2d, 0x64, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2d, 0x64, 0x61,
	0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x86, 0x09, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x63,
	0x61, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72,
	0x46, 0x6c, 0x6f, 0x61, 0x74, 0x33, 0x32, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6c, 0x61,
	0x72, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x62, 0x6f,
	0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72,
	0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52,
	0x0a, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x38, 0x0a, 0x0b, 0x6f,
	0x62, 0x6a, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0xf3, 0x18, 0x02, 0x08, 0x01, 0x52, 0x07,
	0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xc2, 0xf3, 0x18, 0x09, 0x12, 0x07,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x64, 0x52, 0x07, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64,
	0x12, 0x33, 0x0a, 0x10, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xc2, 0xf3, 0x18, 0x04,
	0x1a, 0x02, 0x08, 0x01, 0x52, 0x0f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x10, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64,
	0x5f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x18, 0xc2, 0xf3, 0x18, 0x14, 0x1a, 0x12, 0x12, 0x10, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x64,
	0x5f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x52, 0x0f, 0x72, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0c, 0x69, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xc2, 0xf3, 0x18, 0x04, 0x22, 0x02, 0x08, 0x01, 0x52, 0x0b, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x72, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xc2,
	0xf3, 0x18, 0x10, 0x22, 0x0e, 0x12, 0x0c, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x74, 0x52, 0x0b, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x0d, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xc2, 0xf3, 0x18, 0x04, 0x2a, 0x02, 0x08,
	0x01, 0x52, 0x0c, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x12,
	0x3a, 0x0a, 0x0d, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xc2, 0xf3, 0x18, 0x11, 0x2a, 0x0f, 0x12, 0x0d,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x52, 0x0c, 0x72,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x1a, 0x19, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b,
	0x0a, 0x07, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45,
	0x6e, 0x75, 0x6d, 0x54, 0x77, 0x6f, 0x10, 0x01, 0x3a, 0xdd, 0x02, 0xc2, 0xf3, 0x18, 0xd8, 0x02,
	0x0a, 0x1b, 0x0a, 0x09, 0x6f, 0x6e, 0x65, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a,
	0x0c, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x0a, 0x42, 0x0a,
	0x0e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x0e, 0x0a, 0x0c, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x0e, 0x0a, 0x0c, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12,
	0x10, 0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x33,
	0x32, 0x0a, 0x1e, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x12, 0x0d, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x0a, 0x23, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x0c, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x18, 0x02, 0x0a, 0x7a, 0x0a, 0x09, 0x77, 0x69, 0x74, 0x68, 0x20, 0x72,
	0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0c, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x2a, 0x19, 0x0a, 0x17, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73,
	0x74, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x12,
	0x0a, 0x0c, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x2a, 0x02,
	0x10, 0x01, 0x12, 0x14, 0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x33, 0x32, 0x2a, 0x02, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6c,
	0x61, 0x72, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x2a, 0x02, 0x10, 0x03, 0x32, 0x02,
	0x08, 0x01, 0x0a, 0x16, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x73,
	0x74, 0x75, 0x62, 0x20, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x01, 0x0a, 0x17, 0x0a, 0x13, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x20, 0x73, 0x74, 0x75, 0x62, 0x20, 0x6f, 0x6e, 0x6c,
	0x79, 0x18, 0x02, 0x12, 0x03, 0x01, 0x02, 0x03, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6c, 0x65, 0x73, 0x7a, 0x69, 0x6c, 0x6d, 0x2f,
	0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2d, 0x64, 0x61,
	0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_single_proto_rawDescOnce sync.Once
	file_single_proto_rawDescData = file_single_proto_rawDesc
)

func file_single_proto_rawDescGZIP() []byte {
	file_single_proto_rawDescOnce.Do(func() {
		file_single_proto_rawDescData = protoimpl.X.CompressGZIP(file_single_proto_rawDescData)
	})
	return file_single_proto_rawDescData
}

var file_single_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_single_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_single_proto_goTypes = []interface{}{
	(Single_Enum)(0),       // 0: greeter.Single.Enum
	(*Single)(nil),         // 1: greeter.Single
	(*Single_Message)(nil), // 2: greeter.Single.Message
}
var file_single_proto_depIdxs = []int32{
	0, // 0: greeter.Single.scalar_enum:type_name -> greeter.Single.Enum
	2, // 1: greeter.Single.obj_message:type_name -> greeter.Single.Message
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_single_proto_init() }
func file_single_proto_init() {
	if File_single_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_single_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Single); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_single_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Single_Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_single_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_single_proto_goTypes,
		DependencyIndexes: file_single_proto_depIdxs,
		EnumInfos:         file_single_proto_enumTypes,
		MessageInfos:      file_single_proto_msgTypes,
	}.Build()
	File_single_proto = out.File
	file_single_proto_rawDesc = nil
	file_single_proto_goTypes = nil
	file_single_proto_depIdxs = nil
}
