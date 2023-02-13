// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: caraml/timber/v1/log_writer.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogWriterSourceType int32

const (
	// log type is not specified
	LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_UNSPECIFIED LogWriterSourceType = 0
	// log type for consuming prediction log
	LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_PREDICTION_LOG LogWriterSourceType = 1
	// log type for consuming router log
	LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG LogWriterSourceType = 2
)

// Enum value maps for LogWriterSourceType.
var (
	LogWriterSourceType_name = map[int32]string{
		0: "LOG_WRITER_SOURCE_TYPE_UNSPECIFIED",
		1: "LOG_WRITER_SOURCE_TYPE_PREDICTION_LOG",
		2: "LOG_WRITER_SOURCE_TYPE_ROUTER_LOG",
	}
	LogWriterSourceType_value = map[string]int32{
		"LOG_WRITER_SOURCE_TYPE_UNSPECIFIED":    0,
		"LOG_WRITER_SOURCE_TYPE_PREDICTION_LOG": 1,
		"LOG_WRITER_SOURCE_TYPE_ROUTER_LOG":     2,
	}
)

func (x LogWriterSourceType) Enum() *LogWriterSourceType {
	p := new(LogWriterSourceType)
	*p = x
	return p
}

func (x LogWriterSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogWriterSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_caraml_timber_v1_log_writer_proto_enumTypes[0].Descriptor()
}

func (LogWriterSourceType) Type() protoreflect.EnumType {
	return &file_caraml_timber_v1_log_writer_proto_enumTypes[0]
}

func (x LogWriterSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogWriterSourceType.Descriptor instead.
func (LogWriterSourceType) EnumDescriptor() ([]byte, []int) {
	return file_caraml_timber_v1_log_writer_proto_rawDescGZIP(), []int{0}
}

// Details of the log writer data source
type LogWriterSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Log type. It determines whether prediction_log_source or router_log_source is populated.
	// If the value is LOG_TYPE_PREDICTION_LOG, then prediction_log_source should be valid.
	// Whereas, if the value is LOG_TYPE_ROUTER_LOG, then router_log_source should be valid.
	LogType LogWriterSourceType `protobuf:"varint,1,opt,name=log_type,json=logType,proto3,enum=caraml.timber.v1.LogWriterSourceType" json:"log_type,omitempty"`
	// Prediction log source details
	PredictionLogSource *PredictionLogSource `protobuf:"bytes,2,opt,name=prediction_log_source,json=predictionLogSource,proto3" json:"prediction_log_source,omitempty"`
	// Router log source details
	RouterLogSource *RouterLogSource `protobuf:"bytes,3,opt,name=router_log_source,json=routerLogSource,proto3" json:"router_log_source,omitempty"`
}

func (x *LogWriterSource) Reset() {
	*x = LogWriterSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_timber_v1_log_writer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogWriterSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogWriterSource) ProtoMessage() {}

func (x *LogWriterSource) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use LogWriterSource.ProtoReflect.Descriptor instead.
func (*LogWriterSource) Descriptor() ([]byte, []int) {
	return file_caraml_timber_v1_log_writer_proto_rawDescGZIP(), []int{0}
}

func (x *LogWriterSource) GetLogType() LogWriterSourceType {
	if x != nil {
		return x.LogType
	}
	return LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_UNSPECIFIED
}

func (x *LogWriterSource) GetPredictionLogSource() *PredictionLogSource {
	if x != nil {
		return x.PredictionLogSource
	}
	return nil
}

func (x *LogWriterSource) GetRouterLogSource() *RouterLogSource {
	if x != nil {
		return x.RouterLogSource
	}
	return nil
}

// Prediction log source details
type PredictionLogSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the model producing the prediction log
	ModelId int64 `protobuf:"varint,1,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	// name of the model producing the prediction log
	ModelName string `protobuf:"bytes,2,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	// kafka source configuration where the prediction logs are located
	Kafka *KafkaConfig `protobuf:"bytes,3,opt,name=kafka,proto3" json:"kafka,omitempty"`
}

func (x *PredictionLogSource) Reset() {
	*x = PredictionLogSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_timber_v1_log_writer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictionLogSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictionLogSource) ProtoMessage() {}

func (x *PredictionLogSource) ProtoReflect() protoreflect.Message {
	mi := &file_caraml_timber_v1_log_writer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictionLogSource.ProtoReflect.Descriptor instead.
func (*PredictionLogSource) Descriptor() ([]byte, []int) {
	return file_caraml_timber_v1_log_writer_proto_rawDescGZIP(), []int{1}
}

func (x *PredictionLogSource) GetModelId() int64 {
	if x != nil {
		return x.ModelId
	}
	return 0
}

func (x *PredictionLogSource) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *PredictionLogSource) GetKafka() *KafkaConfig {
	if x != nil {
		return x.Kafka
	}
	return nil
}

type RouterLogSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the router producing the router logs
	RouterId int64 `protobuf:"varint,1,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	// name of the router producing the router logs
	RouterName string `protobuf:"bytes,2,opt,name=router_name,json=routerName,proto3" json:"router_name,omitempty"`
	// kafka source configuration where the router logs are located
	Kafka *KafkaConfig `protobuf:"bytes,3,opt,name=kafka,proto3" json:"kafka,omitempty"`
}

func (x *RouterLogSource) Reset() {
	*x = RouterLogSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_timber_v1_log_writer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterLogSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterLogSource) ProtoMessage() {}

func (x *RouterLogSource) ProtoReflect() protoreflect.Message {
	mi := &file_caraml_timber_v1_log_writer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterLogSource.ProtoReflect.Descriptor instead.
func (*RouterLogSource) Descriptor() ([]byte, []int) {
	return file_caraml_timber_v1_log_writer_proto_rawDescGZIP(), []int{2}
}

func (x *RouterLogSource) GetRouterId() int64 {
	if x != nil {
		return x.RouterId
	}
	return 0
}

func (x *RouterLogSource) GetRouterName() string {
	if x != nil {
		return x.RouterName
	}
	return ""
}

func (x *RouterLogSource) GetKafka() *KafkaConfig {
	if x != nil {
		return x.Kafka
	}
	return nil
}

// LogWriter describes details of a Log Writer
type LogWriter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Project id that owns the log writer
	ProjectId int64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Log writer's ID
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the log writer
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Log writer source
	Source *LogWriterSource `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	// TODO: Add details of where the log is stored at
	// Status of the observation service
	Status Status `protobuf:"varint,10,opt,name=status,proto3,enum=caraml.timber.v1.Status" json:"status,omitempty"`
	// Creation timestamp
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Last update timestamp
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *LogWriter) Reset() {
	*x = LogWriter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caraml_timber_v1_log_writer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogWriter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogWriter) ProtoMessage() {}

func (x *LogWriter) ProtoReflect() protoreflect.Message {
	mi := &file_caraml_timber_v1_log_writer_proto_msgTypes[3]
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
	return file_caraml_timber_v1_log_writer_proto_rawDescGZIP(), []int{3}
}

func (x *LogWriter) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *LogWriter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LogWriter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LogWriter) GetSource() *LogWriterSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *LogWriter) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *LogWriter) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *LogWriter) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_caraml_timber_v1_log_writer_proto protoreflect.FileDescriptor

var file_caraml_timber_v1_log_writer_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f, 0x74, 0x69,
	0x6d, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2f, 0x74, 0x69, 0x6d, 0x62,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x63, 0x61, 0x72, 0x61,
	0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x6c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x59, 0x0a, 0x15, 0x70, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d,
	0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x13, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x6c,
	0x6f, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x05,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x61,
	0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b,
	0x61, 0x66, 0x6b, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x22, 0xb5, 0x02, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x63, 0x61, 0x72, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x69, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x8f, 0x01, 0x0a, 0x13, 0x4c, 0x6f,
	0x67, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x26, 0x0a, 0x22, 0x4c, 0x4f, 0x47, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52, 0x5f,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x25, 0x4c, 0x4f, 0x47,
	0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x44, 0x49, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c,
	0x4f, 0x47, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x4c, 0x4f, 0x47, 0x5f, 0x57, 0x52, 0x49, 0x54,
	0x45, 0x52, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x4f, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x10, 0x02, 0x42, 0xcf, 0x01, 0x0a, 0x14,
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
var file_caraml_timber_v1_log_writer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_caraml_timber_v1_log_writer_proto_goTypes = []interface{}{
	(LogWriterSourceType)(0),      // 0: caraml.timber.v1.LogWriterSourceType
	(*LogWriterSource)(nil),       // 1: caraml.timber.v1.LogWriterSource
	(*PredictionLogSource)(nil),   // 2: caraml.timber.v1.PredictionLogSource
	(*RouterLogSource)(nil),       // 3: caraml.timber.v1.RouterLogSource
	(*LogWriter)(nil),             // 4: caraml.timber.v1.LogWriter
	(*KafkaConfig)(nil),           // 5: caraml.timber.v1.KafkaConfig
	(Status)(0),                   // 6: caraml.timber.v1.Status
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_caraml_timber_v1_log_writer_proto_depIdxs = []int32{
	0, // 0: caraml.timber.v1.LogWriterSource.log_type:type_name -> caraml.timber.v1.LogWriterSourceType
	2, // 1: caraml.timber.v1.LogWriterSource.prediction_log_source:type_name -> caraml.timber.v1.PredictionLogSource
	3, // 2: caraml.timber.v1.LogWriterSource.router_log_source:type_name -> caraml.timber.v1.RouterLogSource
	5, // 3: caraml.timber.v1.PredictionLogSource.kafka:type_name -> caraml.timber.v1.KafkaConfig
	5, // 4: caraml.timber.v1.RouterLogSource.kafka:type_name -> caraml.timber.v1.KafkaConfig
	1, // 5: caraml.timber.v1.LogWriter.source:type_name -> caraml.timber.v1.LogWriterSource
	6, // 6: caraml.timber.v1.LogWriter.status:type_name -> caraml.timber.v1.Status
	7, // 7: caraml.timber.v1.LogWriter.create_time:type_name -> google.protobuf.Timestamp
	7, // 8: caraml.timber.v1.LogWriter.update_time:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_caraml_timber_v1_log_writer_proto_init() }
func file_caraml_timber_v1_log_writer_proto_init() {
	if File_caraml_timber_v1_log_writer_proto != nil {
		return
	}
	file_caraml_timber_v1_kafka_proto_init()
	file_caraml_timber_v1_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_caraml_timber_v1_log_writer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogWriterSource); i {
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
		file_caraml_timber_v1_log_writer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictionLogSource); i {
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
		file_caraml_timber_v1_log_writer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterLogSource); i {
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
		file_caraml_timber_v1_log_writer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   4,
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
