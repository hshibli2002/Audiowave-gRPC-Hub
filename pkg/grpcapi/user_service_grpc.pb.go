// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0--rc3
// source: user_service.proto

package grpcapi

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
	UserService_CreateUser_FullMethodName              = "/mpdb.UserService/CreateUser"
	UserService_GetUserById_FullMethodName             = "/mpdb.UserService/GetUserById"
	UserService_GetUserByUsername_FullMethodName       = "/mpdb.UserService/GetUserByUsername"
	UserService_GetUserByEmail_FullMethodName          = "/mpdb.UserService/GetUserByEmail"
	UserService_GetAllUsers_FullMethodName             = "/mpdb.UserService/GetAllUsers"
	UserService_UpdateUsername_FullMethodName          = "/mpdb.UserService/UpdateUsername"
	UserService_UpdateEmail_FullMethodName             = "/mpdb.UserService/UpdateEmail"
	UserService_DeleteUser_FullMethodName              = "/mpdb.UserService/DeleteUser"
	UserService_IncrementLikes_FullMethodName          = "/mpdb.UserService/IncrementLikes"
	UserService_IncrementFollowingCount_FullMethodName = "/mpdb.UserService/IncrementFollowingCount"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// CREATE REQUESTS
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// READ REQUESTS
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
	GetUserByUsername(ctx context.Context, in *GetUsernameRequest, opts ...grpc.CallOption) (*GetUsernameResponse, error)
	GetUserByEmail(ctx context.Context, in *GetEmailRequest, opts ...grpc.CallOption) (*GetEmailResponse, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	// UPDATE REQUESTS
	UpdateUsername(ctx context.Context, in *UpdateUsernameRequest, opts ...grpc.CallOption) (*UpdateUsernameResponse, error)
	UpdateEmail(ctx context.Context, in *UpdateUserEmailRequest, opts ...grpc.CallOption) (*UpdateUserEmailResponse, error)
	// DELETE REQUESTS
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	// Additional Requests
	IncrementLikes(ctx context.Context, in *IncrementLikesRequest, opts ...grpc.CallOption) (*IncrementLikesResponse, error)
	IncrementFollowingCount(ctx context.Context, in *IncrementFollowingCountRequest, opts ...grpc.CallOption) (*IncrementFollowingCountResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByUsername(ctx context.Context, in *GetUsernameRequest, opts ...grpc.CallOption) (*GetUsernameResponse, error) {
	out := new(GetUsernameResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserByUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByEmail(ctx context.Context, in *GetEmailRequest, opts ...grpc.CallOption) (*GetEmailResponse, error) {
	out := new(GetEmailResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, UserService_GetAllUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUsername(ctx context.Context, in *UpdateUsernameRequest, opts ...grpc.CallOption) (*UpdateUsernameResponse, error) {
	out := new(UpdateUsernameResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateEmail(ctx context.Context, in *UpdateUserEmailRequest, opts ...grpc.CallOption) (*UpdateUserEmailResponse, error) {
	out := new(UpdateUserEmailResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IncrementLikes(ctx context.Context, in *IncrementLikesRequest, opts ...grpc.CallOption) (*IncrementLikesResponse, error) {
	out := new(IncrementLikesResponse)
	err := c.cc.Invoke(ctx, UserService_IncrementLikes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IncrementFollowingCount(ctx context.Context, in *IncrementFollowingCountRequest, opts ...grpc.CallOption) (*IncrementFollowingCountResponse, error) {
	out := new(IncrementFollowingCountResponse)
	err := c.cc.Invoke(ctx, UserService_IncrementFollowingCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// CREATE REQUESTS
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// READ REQUESTS
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	GetUserByUsername(context.Context, *GetUsernameRequest) (*GetUsernameResponse, error)
	GetUserByEmail(context.Context, *GetEmailRequest) (*GetEmailResponse, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
	// UPDATE REQUESTS
	UpdateUsername(context.Context, *UpdateUsernameRequest) (*UpdateUsernameResponse, error)
	UpdateEmail(context.Context, *UpdateUserEmailRequest) (*UpdateUserEmailResponse, error)
	// DELETE REQUESTS
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	// Additional Requests
	IncrementLikes(context.Context, *IncrementLikesRequest) (*IncrementLikesResponse, error)
	IncrementFollowingCount(context.Context, *IncrementFollowingCountRequest) (*IncrementFollowingCountResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) GetUserByUsername(context.Context, *GetUsernameRequest) (*GetUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}
func (UnimplementedUserServiceServer) GetUserByEmail(context.Context, *GetEmailRequest) (*GetEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserServiceServer) UpdateUsername(context.Context, *UpdateUsernameRequest) (*UpdateUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUsername not implemented")
}
func (UnimplementedUserServiceServer) UpdateEmail(context.Context, *UpdateUserEmailRequest) (*UpdateUserEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmail not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) IncrementLikes(context.Context, *IncrementLikesRequest) (*IncrementLikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementLikes not implemented")
}
func (UnimplementedUserServiceServer) IncrementFollowingCount(context.Context, *IncrementFollowingCountRequest) (*IncrementFollowingCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementFollowingCount not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByUsername(ctx, req.(*GetUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByEmail(ctx, req.(*GetEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUsername(ctx, req.(*UpdateUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateEmail(ctx, req.(*UpdateUserEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IncrementLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementLikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IncrementLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_IncrementLikes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IncrementLikes(ctx, req.(*IncrementLikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IncrementFollowingCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementFollowingCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IncrementFollowingCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_IncrementFollowingCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IncrementFollowingCount(ctx, req.(*IncrementFollowingCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mpdb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
		{
			MethodName: "GetUserByUsername",
			Handler:    _UserService_GetUserByUsername_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _UserService_GetUserByEmail_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
		{
			MethodName: "UpdateUsername",
			Handler:    _UserService_UpdateUsername_Handler,
		},
		{
			MethodName: "UpdateEmail",
			Handler:    _UserService_UpdateEmail_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "IncrementLikes",
			Handler:    _UserService_IncrementLikes_Handler,
		},
		{
			MethodName: "IncrementFollowingCount",
			Handler:    _UserService_IncrementFollowingCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
