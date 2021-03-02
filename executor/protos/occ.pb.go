//
// === This file is part of ALICE O² ===
//
// Copyright 2018 CERN and copyright holders of ALICE O².
// Author: Teo Mrnjavac <teo.mrnjavac@cern.ch>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// In applying this license CERN does not waive the privileges and
// immunities granted to it by virtue of its status as an
// Intergovernmental Organization or submit itself to any jurisdiction.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: protos/occ.proto

package pb

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

type StateChangeTrigger int32

const (
	StateChangeTrigger_EXECUTOR           StateChangeTrigger = 0
	StateChangeTrigger_DEVICE_INTENTIONAL StateChangeTrigger = 1
	StateChangeTrigger_DEVICE_ERROR       StateChangeTrigger = 2
)

// Enum value maps for StateChangeTrigger.
var (
	StateChangeTrigger_name = map[int32]string{
		0: "EXECUTOR",
		1: "DEVICE_INTENTIONAL",
		2: "DEVICE_ERROR",
	}
	StateChangeTrigger_value = map[string]int32{
		"EXECUTOR":           0,
		"DEVICE_INTENTIONAL": 1,
		"DEVICE_ERROR":       2,
	}
)

func (x StateChangeTrigger) Enum() *StateChangeTrigger {
	p := new(StateChangeTrigger)
	*p = x
	return p
}

func (x StateChangeTrigger) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateChangeTrigger) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_occ_proto_enumTypes[0].Descriptor()
}

func (StateChangeTrigger) Type() protoreflect.EnumType {
	return &file_protos_occ_proto_enumTypes[0]
}

func (x StateChangeTrigger) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateChangeTrigger.Descriptor instead.
func (StateChangeTrigger) EnumDescriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{0}
}

type StateType int32

const (
	StateType_STATE_STABLE       StateType = 0
	StateType_STATE_INTERMEDIATE StateType = 1
)

// Enum value maps for StateType.
var (
	StateType_name = map[int32]string{
		0: "STATE_STABLE",
		1: "STATE_INTERMEDIATE",
	}
	StateType_value = map[string]int32{
		"STATE_STABLE":       0,
		"STATE_INTERMEDIATE": 1,
	}
)

func (x StateType) Enum() *StateType {
	p := new(StateType)
	*p = x
	return p
}

func (x StateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_occ_proto_enumTypes[1].Descriptor()
}

func (StateType) Type() protoreflect.EnumType {
	return &file_protos_occ_proto_enumTypes[1]
}

func (x StateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateType.Descriptor instead.
func (StateType) EnumDescriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{1}
}

type DeviceEventType int32

const (
	DeviceEventType_NULL_DEVICE_EVENT     DeviceEventType = 0
	DeviceEventType_END_OF_STREAM         DeviceEventType = 1
	DeviceEventType_BASIC_TASK_TERMINATED DeviceEventType = 2
)

// Enum value maps for DeviceEventType.
var (
	DeviceEventType_name = map[int32]string{
		0: "NULL_DEVICE_EVENT",
		1: "END_OF_STREAM",
		2: "BASIC_TASK_TERMINATED",
	}
	DeviceEventType_value = map[string]int32{
		"NULL_DEVICE_EVENT":     0,
		"END_OF_STREAM":         1,
		"BASIC_TASK_TERMINATED": 2,
	}
)

func (x DeviceEventType) Enum() *DeviceEventType {
	p := new(DeviceEventType)
	*p = x
	return p
}

func (x DeviceEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_occ_proto_enumTypes[2].Descriptor()
}

func (DeviceEventType) Type() protoreflect.EnumType {
	return &file_protos_occ_proto_enumTypes[2]
}

func (x DeviceEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceEventType.Descriptor instead.
func (DeviceEventType) EnumDescriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{2}
}

type StateStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StateStreamRequest) Reset() {
	*x = StateStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_occ_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateStreamRequest) ProtoMessage() {}

func (x *StateStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_occ_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateStreamRequest.ProtoReflect.Descriptor instead.
func (*StateStreamRequest) Descriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{0}
}

type StateStreamReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  StateType `protobuf:"varint,1,opt,name=type,proto3,enum=occ_pb.StateType" json:"type,omitempty"`
	State string    `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *StateStreamReply) Reset() {
	*x = StateStreamReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_occ_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateStreamReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateStreamReply) ProtoMessage() {}

func (x *StateStreamReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_occ_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateStreamReply.ProtoReflect.Descriptor instead.
func (*StateStreamReply) Descriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{1}
}

func (x *StateStreamReply) GetType() StateType {
	if x != nil {
		return x.Type
	}
	return StateType_STATE_STABLE
}

func (x *StateStreamReply) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type EventStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventStreamRequest) Reset() {
	*x = EventStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_occ_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStreamRequest) ProtoMessage() {}

func (x *EventStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_occ_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStreamRequest.ProtoReflect.Descriptor instead.
func (*EventStreamRequest) Descriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{2}
}

type DeviceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type DeviceEventType `protobuf:"varint,1,opt,name=type,proto3,enum=occ_pb.DeviceEventType" json:"type,omitempty"`
}

func (x *DeviceEvent) Reset() {
	*x = DeviceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_occ_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceEvent) ProtoMessage() {}

func (x *DeviceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protos_occ_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceEvent.ProtoReflect.Descriptor instead.
func (*DeviceEvent) Descriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceEvent) GetType() DeviceEventType {
	if x != nil {
		return x.Type
	}
	return DeviceEventType_NULL_DEVICE_EVENT
}

type EventStreamReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *DeviceEvent `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventStreamReply) Reset() {
	*x = EventStreamReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_occ_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStreamReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStreamReply) ProtoMessage() {}

func (x *EventStreamReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_occ_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStreamReply.ProtoReflect.Descriptor instead.
func (*EventStreamReply) Descriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{4}
}

func (x *EventStreamReply) GetEvent() *DeviceEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type GetStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStateRequest) Reset() {
	*x = GetStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_occ_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateRequest) ProtoMessage() {}

func (x *GetStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_occ_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateRequest.ProtoReflect.Descriptor instead.
func (*GetStateRequest) Descriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{5}
}

type GetStateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Pid   int32  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *GetStateReply) Reset() {
	*x = GetStateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_occ_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateReply) ProtoMessage() {}

func (x *GetStateReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_occ_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateReply.ProtoReflect.Descriptor instead.
func (*GetStateReply) Descriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{6}
}

func (x *GetStateReply) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GetStateReply) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type ConfigEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ConfigEntry) Reset() {
	*x = ConfigEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_occ_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigEntry) ProtoMessage() {}

func (x *ConfigEntry) ProtoReflect() protoreflect.Message {
	mi := &file_protos_occ_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigEntry.ProtoReflect.Descriptor instead.
func (*ConfigEntry) Descriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{7}
}

func (x *ConfigEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ConfigEntry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TransitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcState        string         `protobuf:"bytes,1,opt,name=srcState,proto3" json:"srcState,omitempty"`
	TransitionEvent string         `protobuf:"bytes,2,opt,name=transitionEvent,proto3" json:"transitionEvent,omitempty"`
	Arguments       []*ConfigEntry `protobuf:"bytes,3,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *TransitionRequest) Reset() {
	*x = TransitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_occ_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitionRequest) ProtoMessage() {}

func (x *TransitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_occ_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitionRequest.ProtoReflect.Descriptor instead.
func (*TransitionRequest) Descriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{8}
}

func (x *TransitionRequest) GetSrcState() string {
	if x != nil {
		return x.SrcState
	}
	return ""
}

func (x *TransitionRequest) GetTransitionEvent() string {
	if x != nil {
		return x.TransitionEvent
	}
	return ""
}

func (x *TransitionRequest) GetArguments() []*ConfigEntry {
	if x != nil {
		return x.Arguments
	}
	return nil
}

type TransitionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trigger         StateChangeTrigger `protobuf:"varint,1,opt,name=trigger,proto3,enum=occ_pb.StateChangeTrigger" json:"trigger,omitempty"`
	State           string             `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	TransitionEvent string             `protobuf:"bytes,3,opt,name=transitionEvent,proto3" json:"transitionEvent,omitempty"`
	Ok              bool               `protobuf:"varint,4,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *TransitionReply) Reset() {
	*x = TransitionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_occ_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransitionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitionReply) ProtoMessage() {}

func (x *TransitionReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_occ_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitionReply.ProtoReflect.Descriptor instead.
func (*TransitionReply) Descriptor() ([]byte, []int) {
	return file_protos_occ_proto_rawDescGZIP(), []int{9}
}

func (x *TransitionReply) GetTrigger() StateChangeTrigger {
	if x != nil {
		return x.Trigger
	}
	return StateChangeTrigger_EXECUTOR
}

func (x *TransitionReply) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *TransitionReply) GetTransitionEvent() string {
	if x != nil {
		return x.TransitionEvent
	}
	return ""
}

func (x *TransitionReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_protos_occ_proto protoreflect.FileDescriptor

var file_protos_occ_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x63, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x6f, 0x63, 0x63, 0x5f, 0x70, 0x62, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x4f, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6f, 0x63, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x14, 0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6f, 0x63, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x3d, 0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x63, 0x63, 0x5f, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x35,
	0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x72, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x72, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x31, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x63, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6f, 0x63, 0x63, 0x5f,
	0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x2a, 0x4c,
	0x0a, 0x12, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x4f, 0x52,
	0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x54,
	0x45, 0x4e, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x45,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x2a, 0x35, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54,
	0x45, 0x10, 0x01, 0x2a, 0x56, 0x0a, 0x0f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x55, 0x4c, 0x4c, 0x5f, 0x44,
	0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x45, 0x4e, 0x44, 0x5f, 0x4f, 0x46, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x10, 0x01,
	0x12, 0x19, 0x0a, 0x15, 0x42, 0x41, 0x53, 0x49, 0x43, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54,
	0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0x99, 0x02, 0x0a, 0x03,
	0x4f, 0x63, 0x63, 0x12, 0x47, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x1a, 0x2e, 0x6f, 0x63, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6f, 0x63, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x0b,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1a, 0x2e, 0x6f, 0x63,
	0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6f, 0x63, 0x63, 0x5f, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x2e, 0x6f, 0x63, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6f, 0x63, 0x63,
	0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x2e, 0x6f, 0x63, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6f,
	0x63, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x29, 0x0a, 0x1c, 0x63, 0x68, 0x2e, 0x63, 0x65,
	0x72, 0x6e, 0x2e, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x6f, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x6f, 0x63, 0x63, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_occ_proto_rawDescOnce sync.Once
	file_protos_occ_proto_rawDescData = file_protos_occ_proto_rawDesc
)

func file_protos_occ_proto_rawDescGZIP() []byte {
	file_protos_occ_proto_rawDescOnce.Do(func() {
		file_protos_occ_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_occ_proto_rawDescData)
	})
	return file_protos_occ_proto_rawDescData
}

var file_protos_occ_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protos_occ_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_occ_proto_goTypes = []interface{}{
	(StateChangeTrigger)(0),    // 0: occ_pb.StateChangeTrigger
	(StateType)(0),             // 1: occ_pb.StateType
	(DeviceEventType)(0),       // 2: occ_pb.DeviceEventType
	(*StateStreamRequest)(nil), // 3: occ_pb.StateStreamRequest
	(*StateStreamReply)(nil),   // 4: occ_pb.StateStreamReply
	(*EventStreamRequest)(nil), // 5: occ_pb.EventStreamRequest
	(*DeviceEvent)(nil),        // 6: occ_pb.DeviceEvent
	(*EventStreamReply)(nil),   // 7: occ_pb.EventStreamReply
	(*GetStateRequest)(nil),    // 8: occ_pb.GetStateRequest
	(*GetStateReply)(nil),      // 9: occ_pb.GetStateReply
	(*ConfigEntry)(nil),        // 10: occ_pb.ConfigEntry
	(*TransitionRequest)(nil),  // 11: occ_pb.TransitionRequest
	(*TransitionReply)(nil),    // 12: occ_pb.TransitionReply
}
var file_protos_occ_proto_depIdxs = []int32{
	1,  // 0: occ_pb.StateStreamReply.type:type_name -> occ_pb.StateType
	2,  // 1: occ_pb.DeviceEvent.type:type_name -> occ_pb.DeviceEventType
	6,  // 2: occ_pb.EventStreamReply.event:type_name -> occ_pb.DeviceEvent
	10, // 3: occ_pb.TransitionRequest.arguments:type_name -> occ_pb.ConfigEntry
	0,  // 4: occ_pb.TransitionReply.trigger:type_name -> occ_pb.StateChangeTrigger
	5,  // 5: occ_pb.Occ.EventStream:input_type -> occ_pb.EventStreamRequest
	3,  // 6: occ_pb.Occ.StateStream:input_type -> occ_pb.StateStreamRequest
	8,  // 7: occ_pb.Occ.GetState:input_type -> occ_pb.GetStateRequest
	11, // 8: occ_pb.Occ.Transition:input_type -> occ_pb.TransitionRequest
	7,  // 9: occ_pb.Occ.EventStream:output_type -> occ_pb.EventStreamReply
	4,  // 10: occ_pb.Occ.StateStream:output_type -> occ_pb.StateStreamReply
	9,  // 11: occ_pb.Occ.GetState:output_type -> occ_pb.GetStateReply
	12, // 12: occ_pb.Occ.Transition:output_type -> occ_pb.TransitionReply
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_protos_occ_proto_init() }
func file_protos_occ_proto_init() {
	if File_protos_occ_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_occ_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateStreamRequest); i {
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
		file_protos_occ_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateStreamReply); i {
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
		file_protos_occ_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventStreamRequest); i {
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
		file_protos_occ_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceEvent); i {
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
		file_protos_occ_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventStreamReply); i {
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
		file_protos_occ_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStateRequest); i {
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
		file_protos_occ_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStateReply); i {
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
		file_protos_occ_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigEntry); i {
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
		file_protos_occ_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransitionRequest); i {
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
		file_protos_occ_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransitionReply); i {
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
			RawDescriptor: file_protos_occ_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_occ_proto_goTypes,
		DependencyIndexes: file_protos_occ_proto_depIdxs,
		EnumInfos:         file_protos_occ_proto_enumTypes,
		MessageInfos:      file_protos_occ_proto_msgTypes,
	}.Build()
	File_protos_occ_proto = out.File
	file_protos_occ_proto_rawDesc = nil
	file_protos_occ_proto_goTypes = nil
	file_protos_occ_proto_depIdxs = nil
}
