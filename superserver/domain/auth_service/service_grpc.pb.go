// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: auth/service.proto

package auth_service

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
	AuthService_ListUser_FullMethodName                = "/auth_service.AuthService/ListUser"
	AuthService_OperateUser_FullMethodName             = "/auth_service.AuthService/OperateUser"
	AuthService_BatchOperateUser_FullMethodName        = "/auth_service.AuthService/BatchOperateUser"
	AuthService_ListRole_FullMethodName                = "/auth_service.AuthService/ListRole"
	AuthService_OperateRole_FullMethodName             = "/auth_service.AuthService/OperateRole"
	AuthService_BatchOperateRole_FullMethodName        = "/auth_service.AuthService/BatchOperateRole"
	AuthService_ListUserRoleRef_FullMethodName         = "/auth_service.AuthService/ListUserRoleRef"
	AuthService_OperateUserRoleRef_FullMethodName      = "/auth_service.AuthService/OperateUserRoleRef"
	AuthService_BatchOperateUserRoleRef_FullMethodName = "/auth_service.AuthService/BatchOperateUserRoleRef"
	AuthService_ListAccess_FullMethodName              = "/auth_service.AuthService/ListAccess"
	AuthService_OperateAccess_FullMethodName           = "/auth_service.AuthService/OperateAccess"
	AuthService_BatchOperateAccess_FullMethodName      = "/auth_service.AuthService/BatchOperateAccess"
	AuthService_ListPermission_FullMethodName          = "/auth_service.AuthService/ListPermission"
	AuthService_OperatePermission_FullMethodName       = "/auth_service.AuthService/OperatePermission"
	AuthService_BatchOperatePermission_FullMethodName  = "/auth_service.AuthService/BatchOperatePermission"
	AuthService_ListAuthToken_FullMethodName           = "/auth_service.AuthService/ListAuthToken"
	AuthService_OperateAuthToken_FullMethodName        = "/auth_service.AuthService/OperateAuthToken"
	AuthService_BatchOperateAuthToken_FullMethodName   = "/auth_service.AuthService/BatchOperateAuthToken"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	// 用户
	ListUser(ctx context.Context, in *ListUserParams, opts ...grpc.CallOption) (*ListUserReply, error)
	OperateUser(ctx context.Context, in *OperateUserParams, opts ...grpc.CallOption) (*OperateUserReply, error)
	BatchOperateUser(ctx context.Context, in *BatchOperateUserParams, opts ...grpc.CallOption) (*BatchOperateUserReply, error)
	// 角色
	ListRole(ctx context.Context, in *ListRoleParams, opts ...grpc.CallOption) (*ListRoleReply, error)
	OperateRole(ctx context.Context, in *OperateRoleParams, opts ...grpc.CallOption) (*OperateRoleReply, error)
	BatchOperateRole(ctx context.Context, in *BatchOperateRoleParams, opts ...grpc.CallOption) (*BatchOperateRoleReply, error)
	// 用户&角色
	ListUserRoleRef(ctx context.Context, in *ListUserRoleRefParams, opts ...grpc.CallOption) (*ListUserRoleRefReply, error)
	OperateUserRoleRef(ctx context.Context, in *OperateUserRoleRefParams, opts ...grpc.CallOption) (*OperateUserRoleRefReply, error)
	BatchOperateUserRoleRef(ctx context.Context, in *BatchOperateUserRoleRefParams, opts ...grpc.CallOption) (*BatchOperateUserRoleRefReply, error)
	// api接口
	ListAccess(ctx context.Context, in *ListAccessParams, opts ...grpc.CallOption) (*ListAccessReply, error)
	OperateAccess(ctx context.Context, in *OperateAccessParams, opts ...grpc.CallOption) (*OperateAccessReply, error)
	BatchOperateAccess(ctx context.Context, in *BatchOperateAccessParams, opts ...grpc.CallOption) (*BatchOperateAccessReply, error)
	// api接口权限
	ListPermission(ctx context.Context, in *ListPermissionParams, opts ...grpc.CallOption) (*ListPermissionReply, error)
	OperatePermission(ctx context.Context, in *OperatePermissionParams, opts ...grpc.CallOption) (*OperatePermissionReply, error)
	BatchOperatePermission(ctx context.Context, in *BatchOperatePermissionParams, opts ...grpc.CallOption) (*BatchOperatePermissionReply, error)
	// 认证令牌
	ListAuthToken(ctx context.Context, in *ListAuthTokenParams, opts ...grpc.CallOption) (*ListAuthTokenReply, error)
	OperateAuthToken(ctx context.Context, in *OperateAuthTokenParams, opts ...grpc.CallOption) (*OperateAuthTokenReply, error)
	BatchOperateAuthToken(ctx context.Context, in *BatchOperateAuthTokenParams, opts ...grpc.CallOption) (*BatchOperateAuthTokenReply, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) ListUser(ctx context.Context, in *ListUserParams, opts ...grpc.CallOption) (*ListUserReply, error) {
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, AuthService_ListUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) OperateUser(ctx context.Context, in *OperateUserParams, opts ...grpc.CallOption) (*OperateUserReply, error) {
	out := new(OperateUserReply)
	err := c.cc.Invoke(ctx, AuthService_OperateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) BatchOperateUser(ctx context.Context, in *BatchOperateUserParams, opts ...grpc.CallOption) (*BatchOperateUserReply, error) {
	out := new(BatchOperateUserReply)
	err := c.cc.Invoke(ctx, AuthService_BatchOperateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListRole(ctx context.Context, in *ListRoleParams, opts ...grpc.CallOption) (*ListRoleReply, error) {
	out := new(ListRoleReply)
	err := c.cc.Invoke(ctx, AuthService_ListRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) OperateRole(ctx context.Context, in *OperateRoleParams, opts ...grpc.CallOption) (*OperateRoleReply, error) {
	out := new(OperateRoleReply)
	err := c.cc.Invoke(ctx, AuthService_OperateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) BatchOperateRole(ctx context.Context, in *BatchOperateRoleParams, opts ...grpc.CallOption) (*BatchOperateRoleReply, error) {
	out := new(BatchOperateRoleReply)
	err := c.cc.Invoke(ctx, AuthService_BatchOperateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListUserRoleRef(ctx context.Context, in *ListUserRoleRefParams, opts ...grpc.CallOption) (*ListUserRoleRefReply, error) {
	out := new(ListUserRoleRefReply)
	err := c.cc.Invoke(ctx, AuthService_ListUserRoleRef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) OperateUserRoleRef(ctx context.Context, in *OperateUserRoleRefParams, opts ...grpc.CallOption) (*OperateUserRoleRefReply, error) {
	out := new(OperateUserRoleRefReply)
	err := c.cc.Invoke(ctx, AuthService_OperateUserRoleRef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) BatchOperateUserRoleRef(ctx context.Context, in *BatchOperateUserRoleRefParams, opts ...grpc.CallOption) (*BatchOperateUserRoleRefReply, error) {
	out := new(BatchOperateUserRoleRefReply)
	err := c.cc.Invoke(ctx, AuthService_BatchOperateUserRoleRef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListAccess(ctx context.Context, in *ListAccessParams, opts ...grpc.CallOption) (*ListAccessReply, error) {
	out := new(ListAccessReply)
	err := c.cc.Invoke(ctx, AuthService_ListAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) OperateAccess(ctx context.Context, in *OperateAccessParams, opts ...grpc.CallOption) (*OperateAccessReply, error) {
	out := new(OperateAccessReply)
	err := c.cc.Invoke(ctx, AuthService_OperateAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) BatchOperateAccess(ctx context.Context, in *BatchOperateAccessParams, opts ...grpc.CallOption) (*BatchOperateAccessReply, error) {
	out := new(BatchOperateAccessReply)
	err := c.cc.Invoke(ctx, AuthService_BatchOperateAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListPermission(ctx context.Context, in *ListPermissionParams, opts ...grpc.CallOption) (*ListPermissionReply, error) {
	out := new(ListPermissionReply)
	err := c.cc.Invoke(ctx, AuthService_ListPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) OperatePermission(ctx context.Context, in *OperatePermissionParams, opts ...grpc.CallOption) (*OperatePermissionReply, error) {
	out := new(OperatePermissionReply)
	err := c.cc.Invoke(ctx, AuthService_OperatePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) BatchOperatePermission(ctx context.Context, in *BatchOperatePermissionParams, opts ...grpc.CallOption) (*BatchOperatePermissionReply, error) {
	out := new(BatchOperatePermissionReply)
	err := c.cc.Invoke(ctx, AuthService_BatchOperatePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListAuthToken(ctx context.Context, in *ListAuthTokenParams, opts ...grpc.CallOption) (*ListAuthTokenReply, error) {
	out := new(ListAuthTokenReply)
	err := c.cc.Invoke(ctx, AuthService_ListAuthToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) OperateAuthToken(ctx context.Context, in *OperateAuthTokenParams, opts ...grpc.CallOption) (*OperateAuthTokenReply, error) {
	out := new(OperateAuthTokenReply)
	err := c.cc.Invoke(ctx, AuthService_OperateAuthToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) BatchOperateAuthToken(ctx context.Context, in *BatchOperateAuthTokenParams, opts ...grpc.CallOption) (*BatchOperateAuthTokenReply, error) {
	out := new(BatchOperateAuthTokenReply)
	err := c.cc.Invoke(ctx, AuthService_BatchOperateAuthToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	// 用户
	ListUser(context.Context, *ListUserParams) (*ListUserReply, error)
	OperateUser(context.Context, *OperateUserParams) (*OperateUserReply, error)
	BatchOperateUser(context.Context, *BatchOperateUserParams) (*BatchOperateUserReply, error)
	// 角色
	ListRole(context.Context, *ListRoleParams) (*ListRoleReply, error)
	OperateRole(context.Context, *OperateRoleParams) (*OperateRoleReply, error)
	BatchOperateRole(context.Context, *BatchOperateRoleParams) (*BatchOperateRoleReply, error)
	// 用户&角色
	ListUserRoleRef(context.Context, *ListUserRoleRefParams) (*ListUserRoleRefReply, error)
	OperateUserRoleRef(context.Context, *OperateUserRoleRefParams) (*OperateUserRoleRefReply, error)
	BatchOperateUserRoleRef(context.Context, *BatchOperateUserRoleRefParams) (*BatchOperateUserRoleRefReply, error)
	// api接口
	ListAccess(context.Context, *ListAccessParams) (*ListAccessReply, error)
	OperateAccess(context.Context, *OperateAccessParams) (*OperateAccessReply, error)
	BatchOperateAccess(context.Context, *BatchOperateAccessParams) (*BatchOperateAccessReply, error)
	// api接口权限
	ListPermission(context.Context, *ListPermissionParams) (*ListPermissionReply, error)
	OperatePermission(context.Context, *OperatePermissionParams) (*OperatePermissionReply, error)
	BatchOperatePermission(context.Context, *BatchOperatePermissionParams) (*BatchOperatePermissionReply, error)
	// 认证令牌
	ListAuthToken(context.Context, *ListAuthTokenParams) (*ListAuthTokenReply, error)
	OperateAuthToken(context.Context, *OperateAuthTokenParams) (*OperateAuthTokenReply, error)
	BatchOperateAuthToken(context.Context, *BatchOperateAuthTokenParams) (*BatchOperateAuthTokenReply, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) ListUser(context.Context, *ListUserParams) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedAuthServiceServer) OperateUser(context.Context, *OperateUserParams) (*OperateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperateUser not implemented")
}
func (UnimplementedAuthServiceServer) BatchOperateUser(context.Context, *BatchOperateUserParams) (*BatchOperateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchOperateUser not implemented")
}
func (UnimplementedAuthServiceServer) ListRole(context.Context, *ListRoleParams) (*ListRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedAuthServiceServer) OperateRole(context.Context, *OperateRoleParams) (*OperateRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperateRole not implemented")
}
func (UnimplementedAuthServiceServer) BatchOperateRole(context.Context, *BatchOperateRoleParams) (*BatchOperateRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchOperateRole not implemented")
}
func (UnimplementedAuthServiceServer) ListUserRoleRef(context.Context, *ListUserRoleRefParams) (*ListUserRoleRefReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserRoleRef not implemented")
}
func (UnimplementedAuthServiceServer) OperateUserRoleRef(context.Context, *OperateUserRoleRefParams) (*OperateUserRoleRefReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperateUserRoleRef not implemented")
}
func (UnimplementedAuthServiceServer) BatchOperateUserRoleRef(context.Context, *BatchOperateUserRoleRefParams) (*BatchOperateUserRoleRefReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchOperateUserRoleRef not implemented")
}
func (UnimplementedAuthServiceServer) ListAccess(context.Context, *ListAccessParams) (*ListAccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccess not implemented")
}
func (UnimplementedAuthServiceServer) OperateAccess(context.Context, *OperateAccessParams) (*OperateAccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperateAccess not implemented")
}
func (UnimplementedAuthServiceServer) BatchOperateAccess(context.Context, *BatchOperateAccessParams) (*BatchOperateAccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchOperateAccess not implemented")
}
func (UnimplementedAuthServiceServer) ListPermission(context.Context, *ListPermissionParams) (*ListPermissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPermission not implemented")
}
func (UnimplementedAuthServiceServer) OperatePermission(context.Context, *OperatePermissionParams) (*OperatePermissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperatePermission not implemented")
}
func (UnimplementedAuthServiceServer) BatchOperatePermission(context.Context, *BatchOperatePermissionParams) (*BatchOperatePermissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchOperatePermission not implemented")
}
func (UnimplementedAuthServiceServer) ListAuthToken(context.Context, *ListAuthTokenParams) (*ListAuthTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthToken not implemented")
}
func (UnimplementedAuthServiceServer) OperateAuthToken(context.Context, *OperateAuthTokenParams) (*OperateAuthTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperateAuthToken not implemented")
}
func (UnimplementedAuthServiceServer) BatchOperateAuthToken(context.Context, *BatchOperateAuthTokenParams) (*BatchOperateAuthTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchOperateAuthToken not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListUser(ctx, req.(*ListUserParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_OperateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateUserParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).OperateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_OperateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).OperateUser(ctx, req.(*OperateUserParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_BatchOperateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperateUserParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).BatchOperateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_BatchOperateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).BatchOperateUser(ctx, req.(*BatchOperateUserParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListRole(ctx, req.(*ListRoleParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_OperateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateRoleParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).OperateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_OperateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).OperateRole(ctx, req.(*OperateRoleParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_BatchOperateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperateRoleParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).BatchOperateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_BatchOperateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).BatchOperateRole(ctx, req.(*BatchOperateRoleParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListUserRoleRef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRoleRefParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListUserRoleRef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListUserRoleRef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListUserRoleRef(ctx, req.(*ListUserRoleRefParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_OperateUserRoleRef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateUserRoleRefParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).OperateUserRoleRef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_OperateUserRoleRef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).OperateUserRoleRef(ctx, req.(*OperateUserRoleRefParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_BatchOperateUserRoleRef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperateUserRoleRefParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).BatchOperateUserRoleRef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_BatchOperateUserRoleRef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).BatchOperateUserRoleRef(ctx, req.(*BatchOperateUserRoleRefParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccessParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListAccess(ctx, req.(*ListAccessParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_OperateAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateAccessParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).OperateAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_OperateAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).OperateAccess(ctx, req.(*OperateAccessParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_BatchOperateAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperateAccessParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).BatchOperateAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_BatchOperateAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).BatchOperateAccess(ctx, req.(*BatchOperateAccessParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPermissionParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListPermission(ctx, req.(*ListPermissionParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_OperatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperatePermissionParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).OperatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_OperatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).OperatePermission(ctx, req.(*OperatePermissionParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_BatchOperatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperatePermissionParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).BatchOperatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_BatchOperatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).BatchOperatePermission(ctx, req.(*BatchOperatePermissionParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthTokenParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListAuthToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListAuthToken(ctx, req.(*ListAuthTokenParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_OperateAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateAuthTokenParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).OperateAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_OperateAuthToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).OperateAuthToken(ctx, req.(*OperateAuthTokenParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_BatchOperateAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperateAuthTokenParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).BatchOperateAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_BatchOperateAuthToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).BatchOperateAuthToken(ctx, req.(*BatchOperateAuthTokenParams))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_service.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _AuthService_ListUser_Handler,
		},
		{
			MethodName: "OperateUser",
			Handler:    _AuthService_OperateUser_Handler,
		},
		{
			MethodName: "BatchOperateUser",
			Handler:    _AuthService_BatchOperateUser_Handler,
		},
		{
			MethodName: "ListRole",
			Handler:    _AuthService_ListRole_Handler,
		},
		{
			MethodName: "OperateRole",
			Handler:    _AuthService_OperateRole_Handler,
		},
		{
			MethodName: "BatchOperateRole",
			Handler:    _AuthService_BatchOperateRole_Handler,
		},
		{
			MethodName: "ListUserRoleRef",
			Handler:    _AuthService_ListUserRoleRef_Handler,
		},
		{
			MethodName: "OperateUserRoleRef",
			Handler:    _AuthService_OperateUserRoleRef_Handler,
		},
		{
			MethodName: "BatchOperateUserRoleRef",
			Handler:    _AuthService_BatchOperateUserRoleRef_Handler,
		},
		{
			MethodName: "ListAccess",
			Handler:    _AuthService_ListAccess_Handler,
		},
		{
			MethodName: "OperateAccess",
			Handler:    _AuthService_OperateAccess_Handler,
		},
		{
			MethodName: "BatchOperateAccess",
			Handler:    _AuthService_BatchOperateAccess_Handler,
		},
		{
			MethodName: "ListPermission",
			Handler:    _AuthService_ListPermission_Handler,
		},
		{
			MethodName: "OperatePermission",
			Handler:    _AuthService_OperatePermission_Handler,
		},
		{
			MethodName: "BatchOperatePermission",
			Handler:    _AuthService_BatchOperatePermission_Handler,
		},
		{
			MethodName: "ListAuthToken",
			Handler:    _AuthService_ListAuthToken_Handler,
		},
		{
			MethodName: "OperateAuthToken",
			Handler:    _AuthService_OperateAuthToken_Handler,
		},
		{
			MethodName: "BatchOperateAuthToken",
			Handler:    _AuthService_BatchOperateAuthToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/service.proto",
}