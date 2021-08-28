// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package entityproto

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

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error)
	Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*Entity, error)
	Create(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*CommandResponse, error)
	Patch(ctx context.Context, in *KeyEntity, opts ...grpc.CallOption) (*CommandResponse, error)
	Put(ctx context.Context, in *KeyEntity, opts ...grpc.CallOption) (*CommandResponse, error)
	Delete(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*CommandResponse, error)
	Import(ctx context.Context, opts ...grpc.CallOption) (RPC_ImportClient, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := c.cc.Invoke(ctx, "/entityproto.RPC/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/entityproto.RPC/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Create(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/entityproto.RPC/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Patch(ctx context.Context, in *KeyEntity, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/entityproto.RPC/Patch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Put(ctx context.Context, in *KeyEntity, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/entityproto.RPC/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Delete(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/entityproto.RPC/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Import(ctx context.Context, opts ...grpc.CallOption) (RPC_ImportClient, error) {
	stream, err := c.cc.NewStream(ctx, &RPC_ServiceDesc.Streams[0], "/entityproto.RPC/Import", opts...)
	if err != nil {
		return nil, err
	}
	x := &rPCImportClient{stream}
	return x, nil
}

type RPC_ImportClient interface {
	Send(*Entity) error
	CloseAndRecv() (*ImportResponse, error)
	grpc.ClientStream
}

type rPCImportClient struct {
	grpc.ClientStream
}

func (x *rPCImportClient) Send(m *Entity) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rPCImportClient) CloseAndRecv() (*ImportResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ImportResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	Index(context.Context, *IndexRequest) (*IndexResponse, error)
	Get(context.Context, *KeyRequest) (*Entity, error)
	Create(context.Context, *Entity) (*CommandResponse, error)
	Patch(context.Context, *KeyEntity) (*CommandResponse, error)
	Put(context.Context, *KeyEntity) (*CommandResponse, error)
	Delete(context.Context, *KeyRequest) (*CommandResponse, error)
	Import(RPC_ImportServer) error
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) Index(context.Context, *IndexRequest) (*IndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedRPCServer) Get(context.Context, *KeyRequest) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRPCServer) Create(context.Context, *Entity) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRPCServer) Patch(context.Context, *KeyEntity) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Patch not implemented")
}
func (UnimplementedRPCServer) Put(context.Context, *KeyEntity) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedRPCServer) Delete(context.Context, *KeyRequest) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRPCServer) Import(RPC_ImportServer) error {
	return status.Errorf(codes.Unimplemented, "method Import not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entityproto.RPC/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entityproto.RPC/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Get(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entityproto.RPC/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Create(ctx, req.(*Entity))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entityproto.RPC/Patch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Patch(ctx, req.(*KeyEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entityproto.RPC/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Put(ctx, req.(*KeyEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entityproto.RPC/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Delete(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Import_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RPCServer).Import(&rPCImportServer{stream})
}

type RPC_ImportServer interface {
	SendAndClose(*ImportResponse) error
	Recv() (*Entity, error)
	grpc.ServerStream
}

type rPCImportServer struct {
	grpc.ServerStream
}

func (x *rPCImportServer) SendAndClose(m *ImportResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rPCImportServer) Recv() (*Entity, error) {
	m := new(Entity)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entityproto.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _RPC_Index_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RPC_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RPC_Create_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _RPC_Patch_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _RPC_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RPC_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Import",
			Handler:       _RPC_Import_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "entity.proto",
}