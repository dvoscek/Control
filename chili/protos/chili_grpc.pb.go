// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: protos/chili.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChiliClient is the client API for Chili service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChiliClient interface {
	EventStream(ctx context.Context, in *EventStreamRequest, opts ...grpc.CallOption) (Chili_EventStreamClient, error)
}

type chiliClient struct {
	cc grpc.ClientConnInterface
}

func NewChiliClient(cc grpc.ClientConnInterface) ChiliClient {
	return &chiliClient{cc}
}

func (c *chiliClient) EventStream(ctx context.Context, in *EventStreamRequest, opts ...grpc.CallOption) (Chili_EventStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chili_ServiceDesc.Streams[0], "/chili.Chili/EventStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chiliEventStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chili_EventStreamClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type chiliEventStreamClient struct {
	grpc.ClientStream
}

func (x *chiliEventStreamClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChiliServer is the server API for Chili service.
// All implementations should embed UnimplementedChiliServer
// for forward compatibility
type ChiliServer interface {
	EventStream(*EventStreamRequest, Chili_EventStreamServer) error
}

// UnimplementedChiliServer should be embedded to have forward compatible implementations.
type UnimplementedChiliServer struct {
}

func (UnimplementedChiliServer) EventStream(*EventStreamRequest, Chili_EventStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EventStream not implemented")
}

// UnsafeChiliServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChiliServer will
// result in compilation errors.
type UnsafeChiliServer interface {
	mustEmbedUnimplementedChiliServer()
}

func RegisterChiliServer(s grpc.ServiceRegistrar, srv ChiliServer) {
	s.RegisterService(&Chili_ServiceDesc, srv)
}

func _Chili_EventStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChiliServer).EventStream(m, &chiliEventStreamServer{stream})
}

type Chili_EventStreamServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type chiliEventStreamServer struct {
	grpc.ServerStream
}

func (x *chiliEventStreamServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// Chili_ServiceDesc is the grpc.ServiceDesc for Chili service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chili_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chili.Chili",
	HandlerType: (*ChiliServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventStream",
			Handler:       _Chili_EventStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/chili.proto",
}
