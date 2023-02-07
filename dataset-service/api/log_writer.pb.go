// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: caraml/timber/v1/log_writer.proto

package api

import (
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

// Type of logging application for specified log
type LogWriterType int32

const (
	LogWriterType_LOG_WRITER_TYPE_UNSPECIFIED LogWriterType = 0
	// Fluentd will be used for logging
	LogWriterType_LOG_WRITER_TYPE_FLUENTD LogWriterType = 1
)

// Enum value maps for LogWriterType.
var (
	LogWriterType_name = map[int32]string{
		0: "LOG_WRITER_TYPE_UNSPECIFIED",
		1: "LOG_WRITER_TYPE_FLUENTD",
	}
	LogWriterType_value = map[string]int32{
		"LOG_WRITER_TYPE_UNSPECIFIED": 0,
		"LOG_WRITER_TYPE_FLUENTD":     1,
	}
)

func (x LogWriterType) Enum() *LogWriterType {
	p := new(LogWriterType)
	*p = x
	return p
}

func (x LogWriterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogWriterType) Descriptor() protoreflect.EnumDescriptor {
	return file_caraml_timber_v1_log_writer_proto_enumTypes[0].Descriptor()
}

func (LogWriterType) Type() protoreflect.EnumType {
	return &file_caraml_timber_v1_log_writer_proto_enumTypes[0]
}

func (x LogWriterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogWriterType.Descriptor instead.
func (LogWriterType) EnumDescriptor() ([]byte, []int) {
	return file_caraml_timber_v1_log_writer_proto_rawDescGZIP(), []int{0}
}

// LogWriter describes details of a Log Writer
type LogWriter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          LogWriterType  `protobuf:"varint,1,opt,name=type,proto3,enum=caraml.timber.v1.LogWriterType" json:"type,omitempty"`
	FluentdConfig *FluentdConfig `protobuf:"bytes,2,opt,name=fluentd_config,json=fluentdConfig,proto3" json:"fluentd_config,omitempty"`
}

func (x *LogWriter) Reset() {
	*x = LogWriter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_timber_v1_log_writer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogWriter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogWriter) ProtoMessage() {}

func (x *LogWriter) ProtoReflect() protoreflect.Message {
	mi := &file_caraml_timber_v1_log_writer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogWriter.ProtoReflect.Descriptor instead.
func (*LogWriter) Descriptor() ([]byte, []int) {
	return file_caraml_timber_v1_log_writer_proto_rawDescGZIP(), []int{0}
}

func (x *LogWriter) GetType() LogWriterType {
	if x != nil {
		return x.Type
	}
	return LogWriterType_LOG_WRITER_TYPE_UNSPECIFIED
}

func (x *LogWriter) GetFluentdConfig() *FluentdConfig {
	if x != nil {
		return x.FluentdConfig
	}
	return nil
}

var File_caraml_timber_v1_log_writer_proto protoreflect.FileDescriptor

var file_caraml_timber_v1_log_writer_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f, 0x74, 0x69,
	0x6d, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x88, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12,
	0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63,
	0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x66,
	0x6c, 0x75, 0x65, 0x6e, 0x74, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0x4d, 0x0a, 0x0d,
	0x4c, 0x6f, 0x67, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x1b, 0x4c, 0x4f, 0x47, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b,
	0x0a, 0x17, 0x4c, 0x4f, 0x47, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x46, 0x4c, 0x55, 0x45, 0x4e, 0x54, 0x44, 0x10, 0x01, 0x42, 0xcf, 0x01, 0x0a, 0x14,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x4c, 0x6f, 0x67, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x74, 0x69,
	0x6d, 0x62, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f,
	0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x03,
	0x43, 0x54, 0x58, 0xaa, 0x02, 0x10, 0x43, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x54, 0x69, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x43, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x5c,
	0x54, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x43, 0x61, 0x72, 0x61,
	0x6d, 0x6c, 0x5c, 0x54, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x43, 0x61, 0x72, 0x61, 0x6d,
	0x6c, 0x3a, 0x3a, 0x54, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_caraml_timber_v1_log_writer_proto_rawDescOnce sync.Once
	file_caraml_timber_v1_log_writer_proto_rawDescData = file_caraml_timber_v1_log_writer_proto_rawDesc
)

func file_caraml_timber_v1_log_writer_proto_rawDescGZIP() []byte {
	file_caraml_timber_v1_log_writer_proto_rawDescOnce.Do(func() {
		file_caraml_timber_v1_log_writer_proto_rawDescData = protoimpl.X.CompressGZIP(file_caraml_timber_v1_log_writer_proto_rawDescData)
	})
	return file_caraml_timber_v1_log_writer_proto_rawDescData
}

var file_caraml_timber_v1_log_writer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_caraml_timber_v1_log_writer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_caraml_timber_v1_log_writer_proto_goTypes = []interface{}{
	(LogWriterType)(0),    // 0: caraml.timber.v1.LogWriterType
	(*LogWriter)(nil),     // 1: caraml.timber.v1.LogWriter
	(*FluentdConfig)(nil), // 2: caraml.timber.v1.FluentdConfig
}
var file_caraml_timber_v1_log_writer_proto_depIdxs = []int32{
	0, // 0: caraml.timber.v1.LogWriter.type:type_name -> caraml.timber.v1.LogWriterType
	2, // 1: caraml.timber.v1.LogWriter.fluentd_config:type_name -> caraml.timber.v1.FluentdConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_caraml_timber_v1_log_writer_proto_init() }
func file_caraml_timber_v1_log_writer_proto_init() {
	if File_caraml_timber_v1_log_writer_proto != nil {
		return
	}
	file_caraml_timber_v1_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_caraml_timber_v1_log_writer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogWriter); i {
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
			RawDescriptor: file_caraml_timber_v1_log_writer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_caraml_timber_v1_log_writer_proto_goTypes,
		DependencyIndexes: file_caraml_timber_v1_log_writer_proto_depIdxs,
		EnumInfos:         file_caraml_timber_v1_log_writer_proto_enumTypes,
		MessageInfos:      file_caraml_timber_v1_log_writer_proto_msgTypes,
	}.Build()
	File_caraml_timber_v1_log_writer_proto = out.File
	file_caraml_timber_v1_log_writer_proto_rawDesc = nil
	file_caraml_timber_v1_log_writer_proto_goTypes = nil
	file_caraml_timber_v1_log_writer_proto_depIdxs = nil
}
