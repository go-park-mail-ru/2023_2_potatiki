// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: api.proto

package protomodels

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

const (
	AuthServiceTest_SayHello_FullMethodName = "/AuthServiceTest/SayHello"
)

// AuthServiceTestClient is the client API for AuthServiceTest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceTestClient interface {
	SayHello(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error)
}

type authServiceTestClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceTestClient(cc grpc.ClientConnInterface) AuthServiceTestClient {
	return &authServiceTestClient{cc}
}

func (c *authServiceTestClient) SayHello(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error) {
	out := new(Hello)
	err := c.cc.Invoke(ctx, AuthServiceTest_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceTestServer is the server API for AuthServiceTest service.
// All implementations must embed UnimplementedAuthServiceTestServer
// for forward compatibility
type AuthServiceTestServer interface {
	SayHello(context.Context, *Hello) (*Hello, error)
	mustEmbedUnimplementedAuthServiceTestServer()
}

// UnimplementedAuthServiceTestServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceTestServer struct {
}

func (UnimplementedAuthServiceTestServer) SayHello(context.Context, *Hello) (*Hello, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedAuthServiceTestServer) mustEmbedUnimplementedAuthServiceTestServer() {}

// UnsafeAuthServiceTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceTestServer will
// result in compilation errors.
type UnsafeAuthServiceTestServer interface {
	mustEmbedUnimplementedAuthServiceTestServer()
}

func RegisterAuthServiceTestServer(s grpc.ServiceRegistrar, srv AuthServiceTestServer) {
	s.RegisterService(&AuthServiceTest_ServiceDesc, srv)
}

func _AuthServiceTest_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceTestServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthServiceTest_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceTestServer).SayHello(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthServiceTest_ServiceDesc is the grpc.ServiceDesc for AuthServiceTest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthServiceTest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthServiceTest",
	HandlerType: (*AuthServiceTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _AuthServiceTest_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
