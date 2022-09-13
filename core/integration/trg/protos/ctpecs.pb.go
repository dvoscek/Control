// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: protos/ctpecs.proto

package ctpecs

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

// -config, e.g. "trig\nPH_bc 100 inf" or
// -key in case the config is available in Consul[key]
// -"" use default configuration
//
// Optional parameters follows, overriding parameters in config string
// for the begining let's agree on three parameters :
type RunStartRequest_Mode int32

const (
	RunStartRequest_NULL RunStartRequest_Mode = 0 // use default (from config)
	RunStartRequest_TRIG RunStartRequest_Mode = 1
	RunStartRequest_CONT RunStartRequest_Mode = 2
)

// Enum value maps for RunStartRequest_Mode.
var (
	RunStartRequest_Mode_name = map[int32]string{
		0: "NULL",
		1: "TRIG",
		2: "CONT",
	}
	RunStartRequest_Mode_value = map[string]int32{
		"NULL": 0,
		"TRIG": 1,
		"CONT": 2,
	}
)

func (x RunStartRequest_Mode) Enum() *RunStartRequest_Mode {
	p := new(RunStartRequest_Mode)
	*p = x
	return p
}

func (x RunStartRequest_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunStartRequest_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_ctpecs_proto_enumTypes[0].Descriptor()
}

func (RunStartRequest_Mode) Type() protoreflect.EnumType {
	return &file_protos_ctpecs_proto_enumTypes[0]
}

func (x RunStartRequest_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunStartRequest_Mode.Descriptor instead.
func (RunStartRequest_Mode) EnumDescriptor() ([]byte, []int) {
	return file_protos_ctpecs_proto_rawDescGZIP(), []int{1, 0}
}

//import "ctpecs_m.proto";
// global runs only:
type RunLoadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runn   uint32 `protobuf:"varint,1,opt,name=runn,proto3" json:"runn,omitempty"` // run number
	Config string `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// -config, i.e. the content of partname.pd
	// -partname -the content is in ctp config dir partname.pd
	// -key in case the config is available in Consul[key]
	Detectors string `protobuf:"bytes,3,opt,name=detectors,proto3" json:"detectors,omitempty"` // available detectors, "" -all required must be ready
}

func (x *RunLoadRequest) Reset() {
	*x = RunLoadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ctpecs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunLoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunLoadRequest) ProtoMessage() {}

func (x *RunLoadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ctpecs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunLoadRequest.ProtoReflect.Descriptor instead.
func (*RunLoadRequest) Descriptor() ([]byte, []int) {
	return file_protos_ctpecs_proto_rawDescGZIP(), []int{0}
}

func (x *RunLoadRequest) GetRunn() uint32 {
	if x != nil {
		return x.Runn
	}
	return 0
}

func (x *RunLoadRequest) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *RunLoadRequest) GetDetectors() string {
	if x != nil {
		return x.Detectors
	}
	return ""
}

// global + stdalone runs (common for ctpd/ltud):
type RunStartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runn   uint32               `protobuf:"varint,1,opt,name=runn,proto3" json:"runn,omitempty"` // run number
	Config string               `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Mode   RunStartRequest_Mode `protobuf:"varint,3,opt,name=mode,proto3,enum=ctpd.RunStartRequest_Mode" json:"mode,omitempty"`
	PhBc   string               `protobuf:"bytes,4,opt,name=ph_bc,json=phBc,proto3" json:"ph_bc,omitempty"`    // BC-downscaled rate of physics triggers
	PhRnd  string               `protobuf:"bytes,5,opt,name=ph_rnd,json=phRnd,proto3" json:"ph_rnd,omitempty"` // rate of physics triggers generated randomly
	// ph_ ...: "": use default (from config)
	//          N : number of BCs between 2 triggers
	//          2.3khz  rate of triggers: 2.3 khz
	Detector string `protobuf:"bytes,6,opt,name=detector,proto3" json:"detector,omitempty"` // detector has to be given for stdalone runs, "" for global run
}

func (x *RunStartRequest) Reset() {
	*x = RunStartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ctpecs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunStartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunStartRequest) ProtoMessage() {}

func (x *RunStartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ctpecs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunStartRequest.ProtoReflect.Descriptor instead.
func (*RunStartRequest) Descriptor() ([]byte, []int) {
	return file_protos_ctpecs_proto_rawDescGZIP(), []int{1}
}

func (x *RunStartRequest) GetRunn() uint32 {
	if x != nil {
		return x.Runn
	}
	return 0
}

func (x *RunStartRequest) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *RunStartRequest) GetMode() RunStartRequest_Mode {
	if x != nil {
		return x.Mode
	}
	return RunStartRequest_NULL
}

func (x *RunStartRequest) GetPhBc() string {
	if x != nil {
		return x.PhBc
	}
	return ""
}

func (x *RunStartRequest) GetPhRnd() string {
	if x != nil {
		return x.PhRnd
	}
	return ""
}

func (x *RunStartRequest) GetDetector() string {
	if x != nil {
		return x.Detector
	}
	return ""
}

type RunStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runn uint32 `protobuf:"varint,1,opt,name=runn,proto3" json:"runn,omitempty"`
}

func (x *RunStatusRequest) Reset() {
	*x = RunStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ctpecs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunStatusRequest) ProtoMessage() {}

func (x *RunStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ctpecs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunStatusRequest.ProtoReflect.Descriptor instead.
func (*RunStatusRequest) Descriptor() ([]byte, []int) {
	return file_protos_ctpecs_proto_rawDescGZIP(), []int{2}
}

func (x *RunStatusRequest) GetRunn() uint32 {
	if x != nil {
		return x.Runn
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ctpecs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ctpecs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_ctpecs_proto_rawDescGZIP(), []int{3}
}

type RunStopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runn     uint32 `protobuf:"varint,1,opt,name=runn,proto3" json:"runn,omitempty"`
	Detector string `protobuf:"bytes,2,opt,name=detector,proto3" json:"detector,omitempty"` // only for stdalone runs (forced stop)
}

func (x *RunStopRequest) Reset() {
	*x = RunStopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ctpecs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunStopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunStopRequest) ProtoMessage() {}

func (x *RunStopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ctpecs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunStopRequest.ProtoReflect.Descriptor instead.
func (*RunStopRequest) Descriptor() ([]byte, []int) {
	return file_protos_ctpecs_proto_rawDescGZIP(), []int{4}
}

func (x *RunStopRequest) GetRunn() uint32 {
	if x != nil {
		return x.Runn
	}
	return 0
}

func (x *RunStopRequest) GetDetector() string {
	if x != nil {
		return x.Detector
	}
	return ""
}

type RunReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rc int32 `protobuf:"varint,1,opt,name=rc,proto3" json:"rc,omitempty"` // 0: ok
	// RunStatus rc: 0:active/running 1:paused 2:loaded 3:does not exist
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` // response (usually "" when rc is 0)
}

func (x *RunReply) Reset() {
	*x = RunReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ctpecs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunReply) ProtoMessage() {}

func (x *RunReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ctpecs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunReply.ProtoReflect.Descriptor instead.
func (*RunReply) Descriptor() ([]byte, []int) {
	return file_protos_ctpecs_proto_rawDescGZIP(), []int{5}
}

func (x *RunReply) GetRc() int32 {
	if x != nil {
		return x.Rc
	}
	return 0
}

func (x *RunReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_protos_ctpecs_proto protoreflect.FileDescriptor

var file_protos_ctpecs_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x74, 0x70, 0x65, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x74, 0x70, 0x64, 0x22, 0x5a, 0x0a, 0x0e, 0x52,
	0x75, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x75, 0x6e, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x75, 0x6e,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x75, 0x6e, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x74, 0x70, 0x64, 0x2e, 0x52, 0x75, 0x6e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x68, 0x5f, 0x62, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x68, 0x42, 0x63, 0x12, 0x15, 0x0a, 0x06,
	0x70, 0x68, 0x5f, 0x72, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x52, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22,
	0x24, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x52, 0x49, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x43,
	0x4f, 0x4e, 0x54, 0x10, 0x02, 0x22, 0x26, 0x0a, 0x10, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6e,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x6e, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x40, 0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x2c, 0x0a, 0x08, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x72, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xec, 0x02, 0x0a, 0x04, 0x43, 0x54, 0x50, 0x64, 0x12,
	0x31, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x2e, 0x63, 0x74, 0x70,
	0x64, 0x2e, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x63, 0x74, 0x70, 0x64, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x14, 0x2e, 0x63, 0x74, 0x70, 0x64, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x63, 0x74, 0x70, 0x64, 0x2e, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x52, 0x75, 0x6e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x63, 0x74, 0x70, 0x64, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x63, 0x74, 0x70,
	0x64, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09,
	0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x63, 0x74, 0x70, 0x64,
	0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x63, 0x74, 0x70, 0x64, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0b,
	0x2e, 0x63, 0x74, 0x70, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x63, 0x74,
	0x70, 0x64, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x07, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x2e, 0x63, 0x74, 0x70, 0x64, 0x2e,
	0x52, 0x75, 0x6e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x63, 0x74, 0x70, 0x64, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x33, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x2e,
	0x63, 0x74, 0x70, 0x64, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x63, 0x74, 0x70, 0x64, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x69, 0x63, 0x65, 0x4f, 0x32, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x63, 0x74, 0x70, 0x65, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_ctpecs_proto_rawDescOnce sync.Once
	file_protos_ctpecs_proto_rawDescData = file_protos_ctpecs_proto_rawDesc
)

func file_protos_ctpecs_proto_rawDescGZIP() []byte {
	file_protos_ctpecs_proto_rawDescOnce.Do(func() {
		file_protos_ctpecs_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_ctpecs_proto_rawDescData)
	})
	return file_protos_ctpecs_proto_rawDescData
}

var file_protos_ctpecs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_ctpecs_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_ctpecs_proto_goTypes = []interface{}{
	(RunStartRequest_Mode)(0), // 0: ctpd.RunStartRequest.Mode
	(*RunLoadRequest)(nil),    // 1: ctpd.RunLoadRequest
	(*RunStartRequest)(nil),   // 2: ctpd.RunStartRequest
	(*RunStatusRequest)(nil),  // 3: ctpd.RunStatusRequest
	(*Empty)(nil),             // 4: ctpd.Empty
	(*RunStopRequest)(nil),    // 5: ctpd.RunStopRequest
	(*RunReply)(nil),          // 6: ctpd.RunReply
}
var file_protos_ctpecs_proto_depIdxs = []int32{
	0, // 0: ctpd.RunStartRequest.mode:type_name -> ctpd.RunStartRequest.Mode
	1, // 1: ctpd.CTPd.RunLoad:input_type -> ctpd.RunLoadRequest
	5, // 2: ctpd.CTPd.RunUnload:input_type -> ctpd.RunStopRequest
	2, // 3: ctpd.CTPd.RunStart:input_type -> ctpd.RunStartRequest
	3, // 4: ctpd.CTPd.RunStatus:input_type -> ctpd.RunStatusRequest
	4, // 5: ctpd.CTPd.RunList:input_type -> ctpd.Empty
	5, // 6: ctpd.CTPd.RunStop:input_type -> ctpd.RunStopRequest
	5, // 7: ctpd.CTPd.RunConfig:input_type -> ctpd.RunStopRequest
	6, // 8: ctpd.CTPd.RunLoad:output_type -> ctpd.RunReply
	6, // 9: ctpd.CTPd.RunUnload:output_type -> ctpd.RunReply
	6, // 10: ctpd.CTPd.RunStart:output_type -> ctpd.RunReply
	6, // 11: ctpd.CTPd.RunStatus:output_type -> ctpd.RunReply
	6, // 12: ctpd.CTPd.RunList:output_type -> ctpd.RunReply
	6, // 13: ctpd.CTPd.RunStop:output_type -> ctpd.RunReply
	6, // 14: ctpd.CTPd.RunConfig:output_type -> ctpd.RunReply
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_ctpecs_proto_init() }
func file_protos_ctpecs_proto_init() {
	if File_protos_ctpecs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_ctpecs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunLoadRequest); i {
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
		file_protos_ctpecs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunStartRequest); i {
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
		file_protos_ctpecs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunStatusRequest); i {
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
		file_protos_ctpecs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protos_ctpecs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunStopRequest); i {
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
		file_protos_ctpecs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunReply); i {
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
			RawDescriptor: file_protos_ctpecs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_ctpecs_proto_goTypes,
		DependencyIndexes: file_protos_ctpecs_proto_depIdxs,
		EnumInfos:         file_protos_ctpecs_proto_enumTypes,
		MessageInfos:      file_protos_ctpecs_proto_msgTypes,
	}.Build()
	File_protos_ctpecs_proto = out.File
	file_protos_ctpecs_proto_rawDesc = nil
	file_protos_ctpecs_proto_goTypes = nil
	file_protos_ctpecs_proto_depIdxs = nil
}
