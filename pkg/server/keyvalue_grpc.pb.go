// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pkg/server/keyvalue.proto

package __

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

// KeyValueServiceClient is the client API for KeyValueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyValueServiceClient interface {
	Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*RangeResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	DeleteRange(ctx context.Context, in *DeleteRangeRequest, opts ...grpc.CallOption) (*DeleteRangeResponse, error)
}

type keyValueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyValueServiceClient(cc grpc.ClientConnInterface) KeyValueServiceClient {
	return &keyValueServiceClient{cc}
}

func (c *keyValueServiceClient) Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*RangeResponse, error) {
	out := new(RangeResponse)
	err := c.cc.Invoke(ctx, "/proto.KeyValueService/Range", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/proto.KeyValueService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) DeleteRange(ctx context.Context, in *DeleteRangeRequest, opts ...grpc.CallOption) (*DeleteRangeResponse, error) {
	out := new(DeleteRangeResponse)
	err := c.cc.Invoke(ctx, "/proto.KeyValueService/DeleteRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueServiceServer is the server API for KeyValueService service.
// All implementations should embed UnimplementedKeyValueServiceServer
// for forward compatibility
type KeyValueServiceServer interface {
	Range(context.Context, *RangeRequest) (*RangeResponse, error)
	Put(context.Context, *PutRequest) (*PutResponse, error)
	DeleteRange(context.Context, *DeleteRangeRequest) (*DeleteRangeResponse, error)
}

// UnimplementedKeyValueServiceServer should be embedded to have forward compatible implementations.
type UnimplementedKeyValueServiceServer struct {
}

func (UnimplementedKeyValueServiceServer) Range(context.Context, *RangeRequest) (*RangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Range not implemented")
}
func (UnimplementedKeyValueServiceServer) Put(context.Context, *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedKeyValueServiceServer) DeleteRange(context.Context, *DeleteRangeRequest) (*DeleteRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRange not implemented")
}

// UnsafeKeyValueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyValueServiceServer will
// result in compilation errors.
type UnsafeKeyValueServiceServer interface {
	mustEmbedUnimplementedKeyValueServiceServer()
}

func RegisterKeyValueServiceServer(s grpc.ServiceRegistrar, srv KeyValueServiceServer) {
	s.RegisterService(&KeyValueService_ServiceDesc, srv)
}

func _KeyValueService_Range_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).Range(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KeyValueService/Range",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).Range(ctx, req.(*RangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KeyValueService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_DeleteRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).DeleteRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KeyValueService/DeleteRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).DeleteRange(ctx, req.(*DeleteRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyValueService_ServiceDesc is the grpc.ServiceDesc for KeyValueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyValueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.KeyValueService",
	HandlerType: (*KeyValueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Range",
			Handler:    _KeyValueService_Range_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KeyValueService_Put_Handler,
		},
		{
			MethodName: "DeleteRange",
			Handler:    _KeyValueService_DeleteRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/server/keyvalue.proto",
}

// WatchServiceClient is the client API for WatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchServiceClient interface {
	Watch(ctx context.Context, opts ...grpc.CallOption) (WatchService_WatchClient, error)
}

type watchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchServiceClient(cc grpc.ClientConnInterface) WatchServiceClient {
	return &watchServiceClient{cc}
}

func (c *watchServiceClient) Watch(ctx context.Context, opts ...grpc.CallOption) (WatchService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &WatchService_ServiceDesc.Streams[0], "/proto.WatchService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchServiceWatchClient{stream}
	return x, nil
}

type WatchService_WatchClient interface {
	Send(*WatchRequest) error
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type watchServiceWatchClient struct {
	grpc.ClientStream
}

func (x *watchServiceWatchClient) Send(m *WatchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *watchServiceWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatchServiceServer is the server API for WatchService service.
// All implementations should embed UnimplementedWatchServiceServer
// for forward compatibility
type WatchServiceServer interface {
	Watch(WatchService_WatchServer) error
}

// UnimplementedWatchServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWatchServiceServer struct {
}

func (UnimplementedWatchServiceServer) Watch(WatchService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

// UnsafeWatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchServiceServer will
// result in compilation errors.
type UnsafeWatchServiceServer interface {
	mustEmbedUnimplementedWatchServiceServer()
}

func RegisterWatchServiceServer(s grpc.ServiceRegistrar, srv WatchServiceServer) {
	s.RegisterService(&WatchService_ServiceDesc, srv)
}

func _WatchService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WatchServiceServer).Watch(&watchServiceWatchServer{stream})
}

type WatchService_WatchServer interface {
	Send(*WatchResponse) error
	Recv() (*WatchRequest, error)
	grpc.ServerStream
}

type watchServiceWatchServer struct {
	grpc.ServerStream
}

func (x *watchServiceWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *watchServiceWatchServer) Recv() (*WatchRequest, error) {
	m := new(WatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatchService_ServiceDesc is the grpc.ServiceDesc for WatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WatchService",
	HandlerType: (*WatchServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _WatchService_Watch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/server/keyvalue.proto",
}