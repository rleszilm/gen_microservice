// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: cmd/protoc-gen-go-genms/annotations/annotations.proto

package annotations

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MicroServiceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rest    bool `protobuf:"varint,1,opt,name=rest,proto3" json:"rest,omitempty"`
	Graphql bool `protobuf:"varint,2,opt,name=graphql,proto3" json:"graphql,omitempty"`
}

func (x *MicroServiceOptions) Reset() {
	*x = MicroServiceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_genms_annotations_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MicroServiceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MicroServiceOptions) ProtoMessage() {}

func (x *MicroServiceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_genms_annotations_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MicroServiceOptions.ProtoReflect.Descriptor instead.
func (*MicroServiceOptions) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *MicroServiceOptions) GetRest() bool {
	if x != nil {
		return x.Rest
	}
	return false
}

func (x *MicroServiceOptions) GetGraphql() bool {
	if x != nil {
		return x.Graphql
	}
	return false
}

var file_cmd_protoc_gen_go_genms_annotations_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*MicroServiceOptions)(nil),
		Field:         50000,
		Name:          "annotations.genms_service",
		Tag:           "bytes,50000,opt,name=genms_service",
		Filename:      "cmd/protoc-gen-go-genms/annotations/annotations.proto",
	},
}

// Extension fields to descriptor.ServiceOptions.
var (
	// optional annotations.MicroServiceOptions genms_service = 50000;
	E_GenmsService = &file_cmd_protoc_gen_go_genms_annotations_annotations_proto_extTypes[0]
)

var File_cmd_protoc_gen_go_genms_annotations_annotations_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x13, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x72, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x3a, 0x6b, 0x0a, 0x0d, 0x67,
	0x65, 0x6e, 0x6d, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6c, 0x65, 0x73, 0x7a, 0x69, 0x6c, 0x6d, 0x2f,
	0x67, 0x65, 0x6e, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDescData = file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDesc
)

func file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDescData
}

var file_cmd_protoc_gen_go_genms_annotations_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_genms_annotations_annotations_proto_goTypes = []interface{}{
	(*MicroServiceOptions)(nil),       // 0: annotations.MicroServiceOptions
	(*descriptor.ServiceOptions)(nil), // 1: google.protobuf.ServiceOptions
}
var file_cmd_protoc_gen_go_genms_annotations_annotations_proto_depIdxs = []int32{
	1, // 0: annotations.genms_service:extendee -> google.protobuf.ServiceOptions
	0, // 1: annotations.genms_service:type_name -> annotations.MicroServiceOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_genms_annotations_annotations_proto_init() }
func file_cmd_protoc_gen_go_genms_annotations_annotations_proto_init() {
	if File_cmd_protoc_gen_go_genms_annotations_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_genms_annotations_annotations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MicroServiceOptions); i {
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
			RawDescriptor: file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_genms_annotations_annotations_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_genms_annotations_annotations_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_genms_annotations_annotations_proto_msgTypes,
		ExtensionInfos:    file_cmd_protoc_gen_go_genms_annotations_annotations_proto_extTypes,
	}.Build()
	File_cmd_protoc_gen_go_genms_annotations_annotations_proto = out.File
	file_cmd_protoc_gen_go_genms_annotations_annotations_proto_rawDesc = nil
	file_cmd_protoc_gen_go_genms_annotations_annotations_proto_goTypes = nil
	file_cmd_protoc_gen_go_genms_annotations_annotations_proto_depIdxs = nil
}
