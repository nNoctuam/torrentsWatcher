package grpc

import (
	"context"
	"fmt"
	"sort"
	models2 "torrentsWatcher/internal/core/models"
	"torrentsWatcher/internal/pb"

	"google.golang.org/grpc"
)

func (s *RpcServer) Search(ctx context.Context, r *pb.SearchRequest) (*pb.TorrentsResponse, error) {
	torrents := s.trackers.SearchEverywhere(r.Text)

	sort.Slice(torrents, func(i, j int) bool {
		return torrents[i].Seeders > torrents[j].Seeders
	})

	return &pb.TorrentsResponse{
		Torrents: models2.TorrentsToPB(torrents),
	}, nil
}

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
