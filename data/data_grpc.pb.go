// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: data.proto

package dataserver

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

// DataServerClient is the client API for DataServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServerClient interface {
	GetData(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Data, error)
}

type dataServerClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServerClient(cc grpc.ClientConnInterface) DataServerClient {
	return &dataServerClient{cc}
}

func (c *dataServerClient) GetData(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/main.DataServer/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServerServer is the server API for DataServer service.
// All implementations must embed UnimplementedDataServerServer
// for forward compatibility
type DataServerServer interface {
	GetData(context.Context, *Request) (*Data, error)
	mustEmbedUnimplementedDataServerServer()
}

// UnimplementedDataServerServer must be embedded to have forward compatible implementations.
type UnimplementedDataServerServer struct {
}

func (UnimplementedDataServerServer) GetData(context.Context, *Request) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedDataServerServer) mustEmbedUnimplementedDataServerServer() {}

// UnsafeDataServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServerServer will
// result in compilation errors.
type UnsafeDataServerServer interface {
	mustEmbedUnimplementedDataServerServer()
}

func RegisterDataServerServer(s grpc.ServiceRegistrar, srv DataServerServer) {
	s.RegisterService(&DataServer_ServiceDesc, srv)
}

func _DataServer_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServerServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.DataServer/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServerServer).GetData(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// DataServer_ServiceDesc is the grpc.ServiceDesc for DataServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.DataServer",
	HandlerType: (*DataServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _DataServer_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}