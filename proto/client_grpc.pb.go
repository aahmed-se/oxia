// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// OxiaClientClient is the client API for OxiaClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OxiaClientClient interface {
	//*
	// Gets all shard-to-server assignments as a stream. Each set of assignments
	// in the response stream will contain all the assignments to bring the client
	// up to date. For example, if a shard is split, the stream will return a
	// single response containing all the new shard assignments as opposed to
	// multiple stream responses, each containing a single shard assignment.
	//
	// Clients should connect to a single random server which will stream the
	// assignments for all shards on all servers.
	ShardAssignments(ctx context.Context, in *ShardAssignmentsRequest, opts ...grpc.CallOption) (OxiaClient_ShardAssignmentsClient, error)
	//*
	// Batches put, delete and delete_range requests.
	//
	// Clients should send this request to the shard leader. In the future,
	// this may be handled server-side in a proxy layer.
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	//*
	// Batches get and get_range requests.
	//
	// Clients should send this request to the shard leader. In the future,
	// this may be handled server-side in a proxy layer.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	GetNotifications(ctx context.Context, in *NotificationsRequest, opts ...grpc.CallOption) (OxiaClient_GetNotificationsClient, error)
	//
	// Creates a new client session. Sessions are kept alive by regularly sending heartbeats via the KeepAlive rpc.
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	//
	// Sends heartbeats to prevent the session from timing out.
	KeepAlive(ctx context.Context, opts ...grpc.CallOption) (OxiaClient_KeepAliveClient, error)
	//
	// Closes a session and removes all ephemeral values associated with it.
	CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*CloseSessionResponse, error)
}

type oxiaClientClient struct {
	cc grpc.ClientConnInterface
}

func NewOxiaClientClient(cc grpc.ClientConnInterface) OxiaClientClient {
	return &oxiaClientClient{cc}
}

func (c *oxiaClientClient) ShardAssignments(ctx context.Context, in *ShardAssignmentsRequest, opts ...grpc.CallOption) (OxiaClient_ShardAssignmentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &OxiaClient_ServiceDesc.Streams[0], "/io.streamnative.oxia.proto.OxiaClient/ShardAssignments", opts...)
	if err != nil {
		return nil, err
	}
	x := &oxiaClientShardAssignmentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OxiaClient_ShardAssignmentsClient interface {
	Recv() (*ShardAssignmentsResponse, error)
	grpc.ClientStream
}

type oxiaClientShardAssignmentsClient struct {
	grpc.ClientStream
}

func (x *oxiaClientShardAssignmentsClient) Recv() (*ShardAssignmentsResponse, error) {
	m := new(ShardAssignmentsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oxiaClientClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/io.streamnative.oxia.proto.OxiaClient/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oxiaClientClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/io.streamnative.oxia.proto.OxiaClient/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oxiaClientClient) GetNotifications(ctx context.Context, in *NotificationsRequest, opts ...grpc.CallOption) (OxiaClient_GetNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &OxiaClient_ServiceDesc.Streams[1], "/io.streamnative.oxia.proto.OxiaClient/GetNotifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &oxiaClientGetNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OxiaClient_GetNotificationsClient interface {
	Recv() (*NotificationBatch, error)
	grpc.ClientStream
}

type oxiaClientGetNotificationsClient struct {
	grpc.ClientStream
}

func (x *oxiaClientGetNotificationsClient) Recv() (*NotificationBatch, error) {
	m := new(NotificationBatch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oxiaClientClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, "/io.streamnative.oxia.proto.OxiaClient/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oxiaClientClient) KeepAlive(ctx context.Context, opts ...grpc.CallOption) (OxiaClient_KeepAliveClient, error) {
	stream, err := c.cc.NewStream(ctx, &OxiaClient_ServiceDesc.Streams[2], "/io.streamnative.oxia.proto.OxiaClient/KeepAlive", opts...)
	if err != nil {
		return nil, err
	}
	x := &oxiaClientKeepAliveClient{stream}
	return x, nil
}

type OxiaClient_KeepAliveClient interface {
	Send(*SessionHeartbeat) error
	CloseAndRecv() (*KeepAliveResponse, error)
	grpc.ClientStream
}

type oxiaClientKeepAliveClient struct {
	grpc.ClientStream
}

func (x *oxiaClientKeepAliveClient) Send(m *SessionHeartbeat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *oxiaClientKeepAliveClient) CloseAndRecv() (*KeepAliveResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(KeepAliveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oxiaClientClient) CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*CloseSessionResponse, error) {
	out := new(CloseSessionResponse)
	err := c.cc.Invoke(ctx, "/io.streamnative.oxia.proto.OxiaClient/CloseSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OxiaClientServer is the server API for OxiaClient service.
// All implementations must embed UnimplementedOxiaClientServer
// for forward compatibility
type OxiaClientServer interface {
	//*
	// Gets all shard-to-server assignments as a stream. Each set of assignments
	// in the response stream will contain all the assignments to bring the client
	// up to date. For example, if a shard is split, the stream will return a
	// single response containing all the new shard assignments as opposed to
	// multiple stream responses, each containing a single shard assignment.
	//
	// Clients should connect to a single random server which will stream the
	// assignments for all shards on all servers.
	ShardAssignments(*ShardAssignmentsRequest, OxiaClient_ShardAssignmentsServer) error
	//*
	// Batches put, delete and delete_range requests.
	//
	// Clients should send this request to the shard leader. In the future,
	// this may be handled server-side in a proxy layer.
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	//*
	// Batches get and get_range requests.
	//
	// Clients should send this request to the shard leader. In the future,
	// this may be handled server-side in a proxy layer.
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	GetNotifications(*NotificationsRequest, OxiaClient_GetNotificationsServer) error
	//
	// Creates a new client session. Sessions are kept alive by regularly sending heartbeats via the KeepAlive rpc.
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	//
	// Sends heartbeats to prevent the session from timing out.
	KeepAlive(OxiaClient_KeepAliveServer) error
	//
	// Closes a session and removes all ephemeral values associated with it.
	CloseSession(context.Context, *CloseSessionRequest) (*CloseSessionResponse, error)
	mustEmbedUnimplementedOxiaClientServer()
}

// UnimplementedOxiaClientServer must be embedded to have forward compatible implementations.
type UnimplementedOxiaClientServer struct {
}

func (UnimplementedOxiaClientServer) ShardAssignments(*ShardAssignmentsRequest, OxiaClient_ShardAssignmentsServer) error {
	return status.Errorf(codes.Unimplemented, "method ShardAssignments not implemented")
}
func (UnimplementedOxiaClientServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedOxiaClientServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedOxiaClientServer) GetNotifications(*NotificationsRequest, OxiaClient_GetNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}
func (UnimplementedOxiaClientServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedOxiaClientServer) KeepAlive(OxiaClient_KeepAliveServer) error {
	return status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (UnimplementedOxiaClientServer) CloseSession(context.Context, *CloseSessionRequest) (*CloseSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseSession not implemented")
}
func (UnimplementedOxiaClientServer) mustEmbedUnimplementedOxiaClientServer() {}

// UnsafeOxiaClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OxiaClientServer will
// result in compilation errors.
type UnsafeOxiaClientServer interface {
	mustEmbedUnimplementedOxiaClientServer()
}

func RegisterOxiaClientServer(s grpc.ServiceRegistrar, srv OxiaClientServer) {
	s.RegisterService(&OxiaClient_ServiceDesc, srv)
}

func _OxiaClient_ShardAssignments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ShardAssignmentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OxiaClientServer).ShardAssignments(m, &oxiaClientShardAssignmentsServer{stream})
}

type OxiaClient_ShardAssignmentsServer interface {
	Send(*ShardAssignmentsResponse) error
	grpc.ServerStream
}

type oxiaClientShardAssignmentsServer struct {
	grpc.ServerStream
}

func (x *oxiaClientShardAssignmentsServer) Send(m *ShardAssignmentsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OxiaClient_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OxiaClientServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.streamnative.oxia.proto.OxiaClient/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OxiaClientServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OxiaClient_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OxiaClientServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.streamnative.oxia.proto.OxiaClient/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OxiaClientServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OxiaClient_GetNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotificationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OxiaClientServer).GetNotifications(m, &oxiaClientGetNotificationsServer{stream})
}

type OxiaClient_GetNotificationsServer interface {
	Send(*NotificationBatch) error
	grpc.ServerStream
}

type oxiaClientGetNotificationsServer struct {
	grpc.ServerStream
}

func (x *oxiaClientGetNotificationsServer) Send(m *NotificationBatch) error {
	return x.ServerStream.SendMsg(m)
}

func _OxiaClient_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OxiaClientServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.streamnative.oxia.proto.OxiaClient/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OxiaClientServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OxiaClient_KeepAlive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OxiaClientServer).KeepAlive(&oxiaClientKeepAliveServer{stream})
}

type OxiaClient_KeepAliveServer interface {
	SendAndClose(*KeepAliveResponse) error
	Recv() (*SessionHeartbeat, error)
	grpc.ServerStream
}

type oxiaClientKeepAliveServer struct {
	grpc.ServerStream
}

func (x *oxiaClientKeepAliveServer) SendAndClose(m *KeepAliveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *oxiaClientKeepAliveServer) Recv() (*SessionHeartbeat, error) {
	m := new(SessionHeartbeat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OxiaClient_CloseSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OxiaClientServer).CloseSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.streamnative.oxia.proto.OxiaClient/CloseSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OxiaClientServer).CloseSession(ctx, req.(*CloseSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OxiaClient_ServiceDesc is the grpc.ServiceDesc for OxiaClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OxiaClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.streamnative.oxia.proto.OxiaClient",
	HandlerType: (*OxiaClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _OxiaClient_Write_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _OxiaClient_Read_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _OxiaClient_CreateSession_Handler,
		},
		{
			MethodName: "CloseSession",
			Handler:    _OxiaClient_CloseSession_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ShardAssignments",
			Handler:       _OxiaClient_ShardAssignments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetNotifications",
			Handler:       _OxiaClient_GetNotifications_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "KeepAlive",
			Handler:       _OxiaClient_KeepAlive_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/client.proto",
}
