// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0--rc3
// source: playlist_service.proto

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
	PlaylistService_CreatePlaylist_FullMethodName         = "/mpdb.PlaylistService/CreatePlaylist"
	PlaylistService_GetPlaylistByID_FullMethodName        = "/mpdb.PlaylistService/GetPlaylistByID"
	PlaylistService_GetPlaylistByTitle_FullMethodName     = "/mpdb.PlaylistService/GetPlaylistByTitle"
	PlaylistService_GetPlaylistByCreatorID_FullMethodName = "/mpdb.PlaylistService/GetPlaylistByCreatorID"
	PlaylistService_UpdatePlaylistTitle_FullMethodName    = "/mpdb.PlaylistService/UpdatePlaylistTitle"
	PlaylistService_AddSongToPlaylist_FullMethodName      = "/mpdb.PlaylistService/AddSongToPlaylist"
	PlaylistService_DeletePlaylist_FullMethodName         = "/mpdb.PlaylistService/DeletePlaylist"
)

// PlaylistServiceClient is the client API for PlaylistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaylistServiceClient interface {
	// CREATE REQUEST
	CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*CreatePlaylistResponse, error)
	// READ REQUEST
	GetPlaylistByID(ctx context.Context, in *GetPlaylistByIDRequest, opts ...grpc.CallOption) (*GetPlaylistByIDResponse, error)
	GetPlaylistByTitle(ctx context.Context, in *GetPlaylistByTitleRequest, opts ...grpc.CallOption) (*GetPlaylistByTitleResponse, error)
	GetPlaylistByCreatorID(ctx context.Context, in *GetPlaylistByCreatorIDRequest, opts ...grpc.CallOption) (*GetPlaylistByCreatorIDResponse, error)
	// UPDATE REQUEST
	UpdatePlaylistTitle(ctx context.Context, in *UpdatePlaylistTitleRequest, opts ...grpc.CallOption) (*UpdatePlaylistTitleResponse, error)
	AddSongToPlaylist(ctx context.Context, in *AddSongToPlaylistRequest, opts ...grpc.CallOption) (*AddSongToPlaylistResponse, error)
	// DELETE REQUEST
	DeletePlaylist(ctx context.Context, in *DeletePlaylistRequest, opts ...grpc.CallOption) (*DeletePlaylistResponse, error)
}

type playlistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaylistServiceClient(cc grpc.ClientConnInterface) PlaylistServiceClient {
	return &playlistServiceClient{cc}
}

func (c *playlistServiceClient) CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*CreatePlaylistResponse, error) {
	out := new(CreatePlaylistResponse)
	err := c.cc.Invoke(ctx, PlaylistService_CreatePlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) GetPlaylistByID(ctx context.Context, in *GetPlaylistByIDRequest, opts ...grpc.CallOption) (*GetPlaylistByIDResponse, error) {
	out := new(GetPlaylistByIDResponse)
	err := c.cc.Invoke(ctx, PlaylistService_GetPlaylistByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) GetPlaylistByTitle(ctx context.Context, in *GetPlaylistByTitleRequest, opts ...grpc.CallOption) (*GetPlaylistByTitleResponse, error) {
	out := new(GetPlaylistByTitleResponse)
	err := c.cc.Invoke(ctx, PlaylistService_GetPlaylistByTitle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) GetPlaylistByCreatorID(ctx context.Context, in *GetPlaylistByCreatorIDRequest, opts ...grpc.CallOption) (*GetPlaylistByCreatorIDResponse, error) {
	out := new(GetPlaylistByCreatorIDResponse)
	err := c.cc.Invoke(ctx, PlaylistService_GetPlaylistByCreatorID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) UpdatePlaylistTitle(ctx context.Context, in *UpdatePlaylistTitleRequest, opts ...grpc.CallOption) (*UpdatePlaylistTitleResponse, error) {
	out := new(UpdatePlaylistTitleResponse)
	err := c.cc.Invoke(ctx, PlaylistService_UpdatePlaylistTitle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) AddSongToPlaylist(ctx context.Context, in *AddSongToPlaylistRequest, opts ...grpc.CallOption) (*AddSongToPlaylistResponse, error) {
	out := new(AddSongToPlaylistResponse)
	err := c.cc.Invoke(ctx, PlaylistService_AddSongToPlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) DeletePlaylist(ctx context.Context, in *DeletePlaylistRequest, opts ...grpc.CallOption) (*DeletePlaylistResponse, error) {
	out := new(DeletePlaylistResponse)
	err := c.cc.Invoke(ctx, PlaylistService_DeletePlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaylistServiceServer is the server API for PlaylistService service.
// All implementations must embed UnimplementedPlaylistServiceServer
// for forward compatibility
type PlaylistServiceServer interface {
	// CREATE REQUEST
	CreatePlaylist(context.Context, *CreatePlaylistRequest) (*CreatePlaylistResponse, error)
	// READ REQUEST
	GetPlaylistByID(context.Context, *GetPlaylistByIDRequest) (*GetPlaylistByIDResponse, error)
	GetPlaylistByTitle(context.Context, *GetPlaylistByTitleRequest) (*GetPlaylistByTitleResponse, error)
	GetPlaylistByCreatorID(context.Context, *GetPlaylistByCreatorIDRequest) (*GetPlaylistByCreatorIDResponse, error)
	// UPDATE REQUEST
	UpdatePlaylistTitle(context.Context, *UpdatePlaylistTitleRequest) (*UpdatePlaylistTitleResponse, error)
	AddSongToPlaylist(context.Context, *AddSongToPlaylistRequest) (*AddSongToPlaylistResponse, error)
	// DELETE REQUEST
	DeletePlaylist(context.Context, *DeletePlaylistRequest) (*DeletePlaylistResponse, error)
	mustEmbedUnimplementedPlaylistServiceServer()
}

// UnimplementedPlaylistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlaylistServiceServer struct {
}

func (UnimplementedPlaylistServiceServer) CreatePlaylist(context.Context, *CreatePlaylistRequest) (*CreatePlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlaylist not implemented")
}
func (UnimplementedPlaylistServiceServer) GetPlaylistByID(context.Context, *GetPlaylistByIDRequest) (*GetPlaylistByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylistByID not implemented")
}
func (UnimplementedPlaylistServiceServer) GetPlaylistByTitle(context.Context, *GetPlaylistByTitleRequest) (*GetPlaylistByTitleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylistByTitle not implemented")
}
func (UnimplementedPlaylistServiceServer) GetPlaylistByCreatorID(context.Context, *GetPlaylistByCreatorIDRequest) (*GetPlaylistByCreatorIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylistByCreatorID not implemented")
}
func (UnimplementedPlaylistServiceServer) UpdatePlaylistTitle(context.Context, *UpdatePlaylistTitleRequest) (*UpdatePlaylistTitleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlaylistTitle not implemented")
}
func (UnimplementedPlaylistServiceServer) AddSongToPlaylist(context.Context, *AddSongToPlaylistRequest) (*AddSongToPlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSongToPlaylist not implemented")
}
func (UnimplementedPlaylistServiceServer) DeletePlaylist(context.Context, *DeletePlaylistRequest) (*DeletePlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlaylist not implemented")
}
func (UnimplementedPlaylistServiceServer) mustEmbedUnimplementedPlaylistServiceServer() {}

// UnsafePlaylistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaylistServiceServer will
// result in compilation errors.
type UnsafePlaylistServiceServer interface {
	mustEmbedUnimplementedPlaylistServiceServer()
}

func RegisterPlaylistServiceServer(s grpc.ServiceRegistrar, srv PlaylistServiceServer) {
	s.RegisterService(&PlaylistService_ServiceDesc, srv)
}

func _PlaylistService_CreatePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).CreatePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlaylistService_CreatePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).CreatePlaylist(ctx, req.(*CreatePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_GetPlaylistByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaylistByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).GetPlaylistByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlaylistService_GetPlaylistByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).GetPlaylistByID(ctx, req.(*GetPlaylistByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_GetPlaylistByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaylistByTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).GetPlaylistByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlaylistService_GetPlaylistByTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).GetPlaylistByTitle(ctx, req.(*GetPlaylistByTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_GetPlaylistByCreatorID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaylistByCreatorIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).GetPlaylistByCreatorID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlaylistService_GetPlaylistByCreatorID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).GetPlaylistByCreatorID(ctx, req.(*GetPlaylistByCreatorIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_UpdatePlaylistTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlaylistTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).UpdatePlaylistTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlaylistService_UpdatePlaylistTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).UpdatePlaylistTitle(ctx, req.(*UpdatePlaylistTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_AddSongToPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSongToPlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).AddSongToPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlaylistService_AddSongToPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).AddSongToPlaylist(ctx, req.(*AddSongToPlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_DeletePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).DeletePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlaylistService_DeletePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).DeletePlaylist(ctx, req.(*DeletePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlaylistService_ServiceDesc is the grpc.ServiceDesc for PlaylistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlaylistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mpdb.PlaylistService",
	HandlerType: (*PlaylistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlaylist",
			Handler:    _PlaylistService_CreatePlaylist_Handler,
		},
		{
			MethodName: "GetPlaylistByID",
			Handler:    _PlaylistService_GetPlaylistByID_Handler,
		},
		{
			MethodName: "GetPlaylistByTitle",
			Handler:    _PlaylistService_GetPlaylistByTitle_Handler,
		},
		{
			MethodName: "GetPlaylistByCreatorID",
			Handler:    _PlaylistService_GetPlaylistByCreatorID_Handler,
		},
		{
			MethodName: "UpdatePlaylistTitle",
			Handler:    _PlaylistService_UpdatePlaylistTitle_Handler,
		},
		{
			MethodName: "AddSongToPlaylist",
			Handler:    _PlaylistService_AddSongToPlaylist_Handler,
		},
		{
			MethodName: "DeletePlaylist",
			Handler:    _PlaylistService_DeletePlaylist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playlist_service.proto",
}
