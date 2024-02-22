// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: executor.proto

package executor

import (
	dbadapter "/dbadapter"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutorClient interface {
	Execute(ctx context.Context, in *dbadapter.Task, opts ...grpc.CallOption) (*dbadapter.Task, error)
	ExecuteBatch(ctx context.Context, in *dbadapter.Tasks, opts ...grpc.CallOption) (Executor_ExecuteBatchClient, error)
	ExecuteAsync(ctx context.Context, in *dbadapter.Task, opts ...grpc.CallOption) (*dbadapter.Task, error)
	ExecuteBatchAsync(ctx context.Context, in *dbadapter.Tasks, opts ...grpc.CallOption) (*dbadapter.Task, error)
	LoadReport(ctx context.Context, in *dbadapter.Filter, opts ...grpc.CallOption) (*dbadapter.Tasks, error)
}

type executorClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutorClient(cc grpc.ClientConnInterface) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) Execute(ctx context.Context, in *dbadapter.Task, opts ...grpc.CallOption) (*dbadapter.Task, error) {
	out := new(dbadapter.Task)
	err := c.cc.Invoke(ctx, "/Executor/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) ExecuteBatch(ctx context.Context, in *dbadapter.Tasks, opts ...grpc.CallOption) (Executor_ExecuteBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Executor_ServiceDesc.Streams[0], "/Executor/ExecuteBatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorExecuteBatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Executor_ExecuteBatchClient interface {
	Recv() (*dbadapter.Task, error)
	grpc.ClientStream
}

type executorExecuteBatchClient struct {
	grpc.ClientStream
}

func (x *executorExecuteBatchClient) Recv() (*dbadapter.Task, error) {
	m := new(dbadapter.Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *executorClient) ExecuteAsync(ctx context.Context, in *dbadapter.Task, opts ...grpc.CallOption) (*dbadapter.Task, error) {
	out := new(dbadapter.Task)
	err := c.cc.Invoke(ctx, "/Executor/ExecuteAsync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) ExecuteBatchAsync(ctx context.Context, in *dbadapter.Tasks, opts ...grpc.CallOption) (*dbadapter.Task, error) {
	out := new(dbadapter.Task)
	err := c.cc.Invoke(ctx, "/Executor/ExecuteBatchAsync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) LoadReport(ctx context.Context, in *dbadapter.Filter, opts ...grpc.CallOption) (*dbadapter.Tasks, error) {
	out := new(dbadapter.Tasks)
	err := c.cc.Invoke(ctx, "/Executor/LoadReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
// All implementations must embed UnimplementedExecutorServer
// for forward compatibility
type ExecutorServer interface {
	Execute(context.Context, *dbadapter.Task) (*dbadapter.Task, error)
	ExecuteBatch(*dbadapter.Tasks, Executor_ExecuteBatchServer) error
	ExecuteAsync(context.Context, *dbadapter.Task) (*dbadapter.Task, error)
	ExecuteBatchAsync(context.Context, *dbadapter.Tasks) (*dbadapter.Task, error)
	LoadReport(context.Context, *dbadapter.Filter) (*dbadapter.Tasks, error)
	mustEmbedUnimplementedExecutorServer()
}

// UnimplementedExecutorServer must be embedded to have forward compatible implementations.
type UnimplementedExecutorServer struct {
}

func (UnimplementedExecutorServer) Execute(context.Context, *dbadapter.Task) (*dbadapter.Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedExecutorServer) ExecuteBatch(*dbadapter.Tasks, Executor_ExecuteBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteBatch not implemented")
}
func (UnimplementedExecutorServer) ExecuteAsync(context.Context, *dbadapter.Task) (*dbadapter.Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteAsync not implemented")
}
func (UnimplementedExecutorServer) ExecuteBatchAsync(context.Context, *dbadapter.Tasks) (*dbadapter.Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteBatchAsync not implemented")
}
func (UnimplementedExecutorServer) LoadReport(context.Context, *dbadapter.Filter) (*dbadapter.Tasks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadReport not implemented")
}
func (UnimplementedExecutorServer) mustEmbedUnimplementedExecutorServer() {}

// UnsafeExecutorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutorServer will
// result in compilation errors.
type UnsafeExecutorServer interface {
	mustEmbedUnimplementedExecutorServer()
}

func RegisterExecutorServer(s grpc.ServiceRegistrar, srv ExecutorServer) {
	s.RegisterService(&Executor_ServiceDesc, srv)
}

func _Executor_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dbadapter.Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Executor/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Execute(ctx, req.(*dbadapter.Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_ExecuteBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(dbadapter.Tasks)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutorServer).ExecuteBatch(m, &executorExecuteBatchServer{stream})
}

type Executor_ExecuteBatchServer interface {
	Send(*dbadapter.Task) error
	grpc.ServerStream
}

type executorExecuteBatchServer struct {
	grpc.ServerStream
}

func (x *executorExecuteBatchServer) Send(m *dbadapter.Task) error {
	return x.ServerStream.SendMsg(m)
}

func _Executor_ExecuteAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dbadapter.Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).ExecuteAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Executor/ExecuteAsync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).ExecuteAsync(ctx, req.(*dbadapter.Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_ExecuteBatchAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dbadapter.Tasks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).ExecuteBatchAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Executor/ExecuteBatchAsync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).ExecuteBatchAsync(ctx, req.(*dbadapter.Tasks))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_LoadReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dbadapter.Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).LoadReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Executor/LoadReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).LoadReport(ctx, req.(*dbadapter.Filter))
	}
	return interceptor(ctx, in, info, handler)
}

// Executor_ServiceDesc is the grpc.ServiceDesc for Executor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Executor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Executor_Execute_Handler,
		},
		{
			MethodName: "ExecuteAsync",
			Handler:    _Executor_ExecuteAsync_Handler,
		},
		{
			MethodName: "ExecuteBatchAsync",
			Handler:    _Executor_ExecuteBatchAsync_Handler,
		},
		{
			MethodName: "LoadReport",
			Handler:    _Executor_LoadReport_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteBatch",
			Handler:       _Executor_ExecuteBatch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "executor.proto",
}
