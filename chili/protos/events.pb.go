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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: protos/events.proto

package pb

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

type ClientType int32

const (
	ClientType_NULL ClientType = 0
	ClientType_CLI  ClientType = 1
	ClientType_GUI  ClientType = 2
)

// Enum value maps for ClientType.
var (
	ClientType_name = map[int32]string{
		0: "NULL",
		1: "CLI",
		2: "GUI",
	}
	ClientType_value = map[string]int32{
		"NULL": 0,
		"CLI":  1,
		"GUI":  2,
	}
)

func (x ClientType) Enum() *ClientType {
	p := new(ClientType)
	*p = x
	return p
}

func (x ClientType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_events_proto_enumTypes[0].Descriptor()
}

func (ClientType) Type() protoreflect.EnumType {
	return &file_protos_events_proto_enumTypes[0]
}

func (x ClientType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientType.Descriptor instead.
func (ClientType) EnumDescriptor() ([]byte, []int) {
	return file_protos_events_proto_rawDescGZIP(), []int{0}
}

type Event_MesosHeartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Event_MesosHeartbeat) Reset() {
	*x = Event_MesosHeartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_MesosHeartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_MesosHeartbeat) ProtoMessage() {}

func (x *Event_MesosHeartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_protos_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_MesosHeartbeat.ProtoReflect.Descriptor instead.
func (*Event_MesosHeartbeat) Descriptor() ([]byte, []int) {
	return file_protos_events_proto_rawDescGZIP(), []int{0}
}

type Ev_MetaEvent_Subscribed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *Ev_MetaEvent_Subscribed) Reset() {
	*x = Ev_MetaEvent_Subscribed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ev_MetaEvent_Subscribed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ev_MetaEvent_Subscribed) ProtoMessage() {}

func (x *Ev_MetaEvent_Subscribed) ProtoReflect() protoreflect.Message {
	mi := &file_protos_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ev_MetaEvent_Subscribed.ProtoReflect.Descriptor instead.
func (*Ev_MetaEvent_Subscribed) Descriptor() ([]byte, []int) {
	return file_protos_events_proto_rawDescGZIP(), []int{1}
}

func (x *Ev_MetaEvent_Subscribed) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type Ev_EnvironmentEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvironmentId    string `protobuf:"bytes,1,opt,name=environmentId,proto3" json:"environmentId,omitempty"`
	State            string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	CurrentRunNumber uint32 `protobuf:"varint,3,opt,name=currentRunNumber,proto3" json:"currentRunNumber,omitempty"`
	Error            string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Message          string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ev_EnvironmentEvent) Reset() {
	*x = Ev_EnvironmentEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ev_EnvironmentEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ev_EnvironmentEvent) ProtoMessage() {}

func (x *Ev_EnvironmentEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protos_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ev_EnvironmentEvent.ProtoReflect.Descriptor instead.
func (*Ev_EnvironmentEvent) Descriptor() ([]byte, []int) {
	return file_protos_events_proto_rawDescGZIP(), []int{2}
}

func (x *Ev_EnvironmentEvent) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *Ev_EnvironmentEvent) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Ev_EnvironmentEvent) GetCurrentRunNumber() uint32 {
	if x != nil {
		return x.CurrentRunNumber
	}
	return 0
}

func (x *Ev_EnvironmentEvent) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Ev_EnvironmentEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Ev_TaskEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Taskid    string `protobuf:"bytes,2,opt,name=taskid,proto3" json:"taskid,omitempty"`
	State     string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Hostname  string `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"`
	ClassName string `protobuf:"bytes,6,opt,name=className,proto3" json:"className,omitempty"`
}

func (x *Ev_TaskEvent) Reset() {
	*x = Ev_TaskEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ev_TaskEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ev_TaskEvent) ProtoMessage() {}

func (x *Ev_TaskEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protos_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ev_TaskEvent.ProtoReflect.Descriptor instead.
func (*Ev_TaskEvent) Descriptor() ([]byte, []int) {
	return file_protos_events_proto_rawDescGZIP(), []int{3}
}

func (x *Ev_TaskEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ev_TaskEvent) GetTaskid() string {
	if x != nil {
		return x.Taskid
	}
	return ""
}

func (x *Ev_TaskEvent) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Ev_TaskEvent) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Ev_TaskEvent) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Ev_TaskEvent) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

type Ev_RoleEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	State    string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	RolePath string `protobuf:"bytes,4,opt,name=rolePath,proto3" json:"rolePath,omitempty"`
}

func (x *Ev_RoleEvent) Reset() {
	*x = Ev_RoleEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ev_RoleEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ev_RoleEvent) ProtoMessage() {}

func (x *Ev_RoleEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protos_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ev_RoleEvent.ProtoReflect.Descriptor instead.
func (*Ev_RoleEvent) Descriptor() ([]byte, []int) {
	return file_protos_events_proto_rawDescGZIP(), []int{4}
}

func (x *Ev_RoleEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ev_RoleEvent) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Ev_RoleEvent) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Ev_RoleEvent) GetRolePath() string {
	if x != nil {
		return x.RolePath
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*Event_EnvironmentEvent
	//	*Event_TaskEvent
	//	*Event_RoleEvent
	//	*Event_MetaEvent
	Payload isEvent_Payload `protobuf_oneof:"Payload"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_protos_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_protos_events_proto_rawDescGZIP(), []int{5}
}

func (x *Event) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (m *Event) GetPayload() isEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Event) GetEnvironmentEvent() *Ev_EnvironmentEvent {
	if x, ok := x.GetPayload().(*Event_EnvironmentEvent); ok {
		return x.EnvironmentEvent
	}
	return nil
}

func (x *Event) GetTaskEvent() *Ev_TaskEvent {
	if x, ok := x.GetPayload().(*Event_TaskEvent); ok {
		return x.TaskEvent
	}
	return nil
}

func (x *Event) GetRoleEvent() *Ev_RoleEvent {
	if x, ok := x.GetPayload().(*Event_RoleEvent); ok {
		return x.RoleEvent
	}
	return nil
}

func (x *Event) GetMetaEvent() *Ev_MetaEvent_Subscribed {
	if x, ok := x.GetPayload().(*Event_MetaEvent); ok {
		return x.MetaEvent
	}
	return nil
}

type isEvent_Payload interface {
	isEvent_Payload()
}

type Event_EnvironmentEvent struct {
	EnvironmentEvent *Ev_EnvironmentEvent `protobuf:"bytes,2,opt,name=environmentEvent,proto3,oneof"`
}

type Event_TaskEvent struct {
	TaskEvent *Ev_TaskEvent `protobuf:"bytes,3,opt,name=taskEvent,proto3,oneof"`
}

type Event_RoleEvent struct {
	RoleEvent *Ev_RoleEvent `protobuf:"bytes,4,opt,name=roleEvent,proto3,oneof"`
}

type Event_MetaEvent struct {
	MetaEvent *Ev_MetaEvent_Subscribed `protobuf:"bytes,5,opt,name=metaEvent,proto3,oneof"`
}

func (*Event_EnvironmentEvent) isEvent_Payload() {}

func (*Event_TaskEvent) isEvent_Payload() {}

func (*Event_RoleEvent) isEvent_Payload() {}

func (*Event_MetaEvent) isEvent_Payload() {}

type EventStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId      string     `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	ClientType    ClientType `protobuf:"varint,2,opt,name=clientType,proto3,enum=events.ClientType" json:"clientType,omitempty"`
	ClientName    string     `protobuf:"bytes,3,opt,name=clientName,proto3" json:"clientName,omitempty"`
	ClientVersion string     `protobuf:"bytes,4,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
}

func (x *EventStreamRequest) Reset() {
	*x = EventStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStreamRequest) ProtoMessage() {}

func (x *EventStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_events_proto_msgTypes[6]
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
	return file_protos_events_proto_rawDescGZIP(), []int{6}
}

func (x *EventStreamRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *EventStreamRequest) GetClientType() ClientType {
	if x != nil {
		return x.ClientType
	}
	return ClientType_NULL
}

func (x *EventStreamRequest) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *EventStreamRequest) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

var File_protos_events_proto protoreflect.FileDescriptor

var file_protos_events_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x16, 0x0a,
	0x14, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x4d, 0x65, 0x73, 0x6f, 0x73, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x22, 0x35, 0x0a, 0x17, 0x45, 0x76, 0x5f, 0x4d, 0x65, 0x74, 0x61,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xad, 0x01, 0x0a,
	0x13, 0x45, 0x76, 0x5f, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x75, 0x6e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x52, 0x75, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa2, 0x01, 0x0a,
	0x0c, 0x45, 0x76, 0x5f, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x6c, 0x0a, 0x0c, 0x45, 0x76, 0x5f, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22,
	0xa8, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x49, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x5f, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45,
	0x76, 0x5f, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x09, 0x74,
	0x61, 0x73, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x5f, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x48, 0x00, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3f,
	0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x5f, 0x4d, 0x65,
	0x74, 0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x64, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0x09, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x12, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x28, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x43, 0x4c, 0x49, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x55, 0x49, 0x10,
	0x02, 0x42, 0x53, 0x0a, 0x1f, 0x63, 0x68, 0x2e, 0x63, 0x65, 0x72, 0x6e, 0x2e, 0x61, 0x6c, 0x69,
	0x63, 0x65, 0x2e, 0x6f, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x6c, 0x69, 0x63, 0x65, 0x4f, 0x32, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_events_proto_rawDescOnce sync.Once
	file_protos_events_proto_rawDescData = file_protos_events_proto_rawDesc
)

func file_protos_events_proto_rawDescGZIP() []byte {
	file_protos_events_proto_rawDescOnce.Do(func() {
		file_protos_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_events_proto_rawDescData)
	})
	return file_protos_events_proto_rawDescData
}

var file_protos_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_events_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_events_proto_goTypes = []interface{}{
	(ClientType)(0),                 // 0: events.ClientType
	(*Event_MesosHeartbeat)(nil),    // 1: events.Event_MesosHeartbeat
	(*Ev_MetaEvent_Subscribed)(nil), // 2: events.Ev_MetaEvent_Subscribed
	(*Ev_EnvironmentEvent)(nil),     // 3: events.Ev_EnvironmentEvent
	(*Ev_TaskEvent)(nil),            // 4: events.Ev_TaskEvent
	(*Ev_RoleEvent)(nil),            // 5: events.Ev_RoleEvent
	(*Event)(nil),                   // 6: events.Event
	(*EventStreamRequest)(nil),      // 7: events.EventStreamRequest
}
var file_protos_events_proto_depIdxs = []int32{
	3, // 0: events.Event.environmentEvent:type_name -> events.Ev_EnvironmentEvent
	4, // 1: events.Event.taskEvent:type_name -> events.Ev_TaskEvent
	5, // 2: events.Event.roleEvent:type_name -> events.Ev_RoleEvent
	2, // 3: events.Event.metaEvent:type_name -> events.Ev_MetaEvent_Subscribed
	0, // 4: events.EventStreamRequest.clientType:type_name -> events.ClientType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_events_proto_init() }
func file_protos_events_proto_init() {
	if File_protos_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_MesosHeartbeat); i {
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
		file_protos_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ev_MetaEvent_Subscribed); i {
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
		file_protos_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ev_EnvironmentEvent); i {
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
		file_protos_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ev_TaskEvent); i {
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
		file_protos_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ev_RoleEvent); i {
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
		file_protos_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_protos_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_protos_events_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Event_EnvironmentEvent)(nil),
		(*Event_TaskEvent)(nil),
		(*Event_RoleEvent)(nil),
		(*Event_MetaEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_events_proto_goTypes,
		DependencyIndexes: file_protos_events_proto_depIdxs,
		EnumInfos:         file_protos_events_proto_enumTypes,
		MessageInfos:      file_protos_events_proto_msgTypes,
	}.Build()
	File_protos_events_proto = out.File
	file_protos_events_proto_rawDesc = nil
	file_protos_events_proto_goTypes = nil
	file_protos_events_proto_depIdxs = nil
}
