// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: protos/kafka.proto

package kafka

import (
	proto "github.com/golang/protobuf/proto"
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

type NewStateNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvInfo   *EnvInfo `protobuf:"bytes,1,opt,name=envInfo,proto3" json:"envInfo,omitempty"`
	Timestamp uint64   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // ms since epoch
}

func (x *NewStateNotification) Reset() {
	*x = NewStateNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_kafka_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewStateNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewStateNotification) ProtoMessage() {}

func (x *NewStateNotification) ProtoReflect() protoreflect.Message {
	mi := &file_protos_kafka_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewStateNotification.ProtoReflect.Descriptor instead.
func (*NewStateNotification) Descriptor() ([]byte, []int) {
	return file_protos_kafka_proto_rawDescGZIP(), []int{0}
}

func (x *NewStateNotification) GetEnvInfo() *EnvInfo {
	if x != nil {
		return x.EnvInfo
	}
	return nil
}

func (x *NewStateNotification) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ActiveRunsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActiveRuns []*EnvInfo `protobuf:"bytes,1,rep,name=activeRuns,proto3" json:"activeRuns,omitempty"`
	Timestamp  uint64     `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // ms since epoch
}

func (x *ActiveRunsList) Reset() {
	*x = ActiveRunsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_kafka_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveRunsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRunsList) ProtoMessage() {}

func (x *ActiveRunsList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_kafka_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRunsList.ProtoReflect.Descriptor instead.
func (*ActiveRunsList) Descriptor() ([]byte, []int) {
	return file_protos_kafka_proto_rawDescGZIP(), []int{1}
}

func (x *ActiveRunsList) GetActiveRuns() []*EnvInfo {
	if x != nil {
		return x.ActiveRuns
	}
	return nil
}

func (x *ActiveRunsList) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type EnvInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvironmentId       string   `protobuf:"bytes,1,opt,name=environmentId,proto3" json:"environmentId,omitempty"`
	RunNumber           *uint32  `protobuf:"varint,2,opt,name=runNumber,proto3,oneof" json:"runNumber,omitempty"`
	RunType             *string  `protobuf:"bytes,3,opt,name=runType,proto3,oneof" json:"runType,omitempty"`
	State               string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Detectors           []string `protobuf:"bytes,5,rep,name=detectors,proto3" json:"detectors,omitempty"`
	EnterStateTimestamp uint64   `protobuf:"varint,6,opt,name=enterStateTimestamp,proto3" json:"enterStateTimestamp,omitempty"` // ms since epoch.
}

func (x *EnvInfo) Reset() {
	*x = EnvInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_kafka_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvInfo) ProtoMessage() {}

func (x *EnvInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_kafka_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvInfo.ProtoReflect.Descriptor instead.
func (*EnvInfo) Descriptor() ([]byte, []int) {
	return file_protos_kafka_proto_rawDescGZIP(), []int{2}
}

func (x *EnvInfo) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *EnvInfo) GetRunNumber() uint32 {
	if x != nil && x.RunNumber != nil {
		return *x.RunNumber
	}
	return 0
}

func (x *EnvInfo) GetRunType() string {
	if x != nil && x.RunType != nil {
		return *x.RunType
	}
	return ""
}

func (x *EnvInfo) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *EnvInfo) GetDetectors() []string {
	if x != nil {
		return x.Detectors
	}
	return nil
}

func (x *EnvInfo) GetEnterStateTimestamp() uint64 {
	if x != nil {
		return x.EnterStateTimestamp
	}
	return 0
}

var File_protos_kafka_proto protoreflect.FileDescriptor

var file_protos_kafka_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x22, 0x5e, 0x0a, 0x14, 0x4e,
	0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x07, 0x65, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x45, 0x6e, 0x76,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x5e, 0x0a, 0x0e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x45, 0x6e, 0x76, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6e, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xf1, 0x01, 0x0a, 0x07,
	0x45, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x09, 0x72, 0x75, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x00, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x13, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x75, 0x6e, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_kafka_proto_rawDescOnce sync.Once
	file_protos_kafka_proto_rawDescData = file_protos_kafka_proto_rawDesc
)

func file_protos_kafka_proto_rawDescGZIP() []byte {
	file_protos_kafka_proto_rawDescOnce.Do(func() {
		file_protos_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_kafka_proto_rawDescData)
	})
	return file_protos_kafka_proto_rawDescData
}

var file_protos_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_kafka_proto_goTypes = []interface{}{
	(*NewStateNotification)(nil), // 0: kafka.NewStateNotification
	(*ActiveRunsList)(nil),       // 1: kafka.ActiveRunsList
	(*EnvInfo)(nil),              // 2: kafka.EnvInfo
}
var file_protos_kafka_proto_depIdxs = []int32{
	2, // 0: kafka.NewStateNotification.envInfo:type_name -> kafka.EnvInfo
	2, // 1: kafka.ActiveRunsList.activeRuns:type_name -> kafka.EnvInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_kafka_proto_init() }
func file_protos_kafka_proto_init() {
	if File_protos_kafka_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_kafka_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewStateNotification); i {
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
		file_protos_kafka_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveRunsList); i {
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
		file_protos_kafka_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvInfo); i {
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
	file_protos_kafka_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_kafka_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_kafka_proto_goTypes,
		DependencyIndexes: file_protos_kafka_proto_depIdxs,
		MessageInfos:      file_protos_kafka_proto_msgTypes,
	}.Build()
	File_protos_kafka_proto = out.File
	file_protos_kafka_proto_rawDesc = nil
	file_protos_kafka_proto_goTypes = nil
	file_protos_kafka_proto_depIdxs = nil
}
