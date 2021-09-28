// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// BaseServiceClient is the client API for BaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaseServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*TorrentsResponse, error)
	GetMonitoredTorrents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TorrentsResponse, error)
	GetDownloadFolders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DownloadFoldersResponse, error)
	AddTorrent(ctx context.Context, in *AddTorrentRequest, opts ...grpc.CallOption) (*TorrentResponse, error)
	DeleteTorrent(ctx context.Context, in *DeleteTorrentRequest, opts ...grpc.CallOption) (*Empty, error)
}

type baseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseServiceClient(cc grpc.ClientConnInterface) BaseServiceClient {
	return &baseServiceClient{cc}
}

func (c *baseServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*TorrentsResponse, error) {
	out := new(TorrentsResponse)
	err := c.cc.Invoke(ctx, "/protobuf.BaseService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) GetMonitoredTorrents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TorrentsResponse, error) {
	out := new(TorrentsResponse)
	err := c.cc.Invoke(ctx, "/protobuf.BaseService/GetMonitoredTorrents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) GetDownloadFolders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DownloadFoldersResponse, error) {
	out := new(DownloadFoldersResponse)
	err := c.cc.Invoke(ctx, "/protobuf.BaseService/GetDownloadFolders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) AddTorrent(ctx context.Context, in *AddTorrentRequest, opts ...grpc.CallOption) (*TorrentResponse, error) {
	out := new(TorrentResponse)
	err := c.cc.Invoke(ctx, "/protobuf.BaseService/AddTorrent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) DeleteTorrent(ctx context.Context, in *DeleteTorrentRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protobuf.BaseService/DeleteTorrent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseServiceServer is the server API for BaseService service.
// All implementations must embed UnimplementedBaseServiceServer
// for forward compatibility
type BaseServiceServer interface {
	Search(context.Context, *SearchRequest) (*TorrentsResponse, error)
	GetMonitoredTorrents(context.Context, *Empty) (*TorrentsResponse, error)
	GetDownloadFolders(context.Context, *Empty) (*DownloadFoldersResponse, error)
	AddTorrent(context.Context, *AddTorrentRequest) (*TorrentResponse, error)
	DeleteTorrent(context.Context, *DeleteTorrentRequest) (*Empty, error)
	mustEmbedUnimplementedBaseServiceServer()
}

// UnimplementedBaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBaseServiceServer struct {
}

func (UnimplementedBaseServiceServer) Search(context.Context, *SearchRequest) (*TorrentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedBaseServiceServer) GetMonitoredTorrents(context.Context, *Empty) (*TorrentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonitoredTorrents not implemented")
}
func (UnimplementedBaseServiceServer) GetDownloadFolders(context.Context, *Empty) (*DownloadFoldersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadFolders not implemented")
}
func (UnimplementedBaseServiceServer) AddTorrent(context.Context, *AddTorrentRequest) (*TorrentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTorrent not implemented")
}
func (UnimplementedBaseServiceServer) DeleteTorrent(context.Context, *DeleteTorrentRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTorrent not implemented")
}
func (UnimplementedBaseServiceServer) mustEmbedUnimplementedBaseServiceServer() {}

// UnsafeBaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseServiceServer will
// result in compilation errors.
type UnsafeBaseServiceServer interface {
	mustEmbedUnimplementedBaseServiceServer()
}

func RegisterBaseServiceServer(s grpc.ServiceRegistrar, srv BaseServiceServer) {
	s.RegisterService(&BaseService_ServiceDesc, srv)
}

func _BaseService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_GetMonitoredTorrents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).GetMonitoredTorrents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/GetMonitoredTorrents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).GetMonitoredTorrents(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_GetDownloadFolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).GetDownloadFolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/GetDownloadFolders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).GetDownloadFolders(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_AddTorrent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTorrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).AddTorrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/AddTorrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).AddTorrent(ctx, req.(*AddTorrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_DeleteTorrent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTorrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).DeleteTorrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/DeleteTorrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).DeleteTorrent(ctx, req.(*DeleteTorrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BaseService_ServiceDesc is the grpc.ServiceDesc for BaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.BaseService",
	HandlerType: (*BaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _BaseService_Search_Handler,
		},
		{
			MethodName: "GetMonitoredTorrents",
			Handler:    _BaseService_GetMonitoredTorrents_Handler,
		},
		{
			MethodName: "GetDownloadFolders",
			Handler:    _BaseService_GetDownloadFolders_Handler,
		},
		{
			MethodName: "AddTorrent",
			Handler:    _BaseService_AddTorrent_Handler,
		},
		{
			MethodName: "DeleteTorrent",
			Handler:    _BaseService_DeleteTorrent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "baseService.proto",
}
