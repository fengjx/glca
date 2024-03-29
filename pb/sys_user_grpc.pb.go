// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: sys_user.proto

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

// UserLoginClient is the client API for UserLogin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserLoginClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
}

type userLoginClient struct {
	cc grpc.ClientConnInterface
}

func NewUserLoginClient(cc grpc.ClientConnInterface) UserLoginClient {
	return &userLoginClient{cc}
}

func (c *userLoginClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/pb.UserLogin/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLoginServer is the server API for UserLogin service.
// All implementations must embed UnimplementedUserLoginServer
// for forward compatibility
type UserLoginServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	mustEmbedUnimplementedUserLoginServer()
}

// UnimplementedUserLoginServer must be embedded to have forward compatible implementations.
type UnimplementedUserLoginServer struct {
}

func (UnimplementedUserLoginServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserLoginServer) mustEmbedUnimplementedUserLoginServer() {}

// UnsafeUserLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserLoginServer will
// result in compilation errors.
type UnsafeUserLoginServer interface {
	mustEmbedUnimplementedUserLoginServer()
}

func RegisterUserLoginServer(s grpc.ServiceRegistrar, srv UserLoginServer) {
	s.RegisterService(&UserLogin_ServiceDesc, srv)
}

func _UserLogin_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserLogin/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserLogin_ServiceDesc is the grpc.ServiceDesc for UserLogin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserLogin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserLogin",
	HandlerType: (*UserLoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserLogin_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys_user.proto",
}
