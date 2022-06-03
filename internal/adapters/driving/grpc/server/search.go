package grpc

import (
	"context"
	"fmt"

	"torrentsWatcher/internal/ports/adapters/driving/pb"

	"google.golang.org/grpc"
)

func (s *RPCServer) Search(ctx context.Context, r *pb.SearchRequest) (*pb.TorrentsResponse, error) {
	torrents := s.torrents.Search(r.Text)
	return &pb.TorrentsResponse{
		Torrents: pb.TorrentsToPB(torrents),
	}, nil
}

// nolint: revive
func SearchHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.SearchRequest)
	fmt.Println("got search request")
	fmt.Printf("%+v\n", in)
	if err := dec(in); err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}
	if interceptor == nil {
		return srv.(pb.BaseServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(pb.BaseServiceServer).Search(ctx, req.(*pb.SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}
