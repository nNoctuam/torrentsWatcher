package grpc

import (
	"context"
	"fmt"

	"torrentsWatcher/internal/ports/pb"

	"google.golang.org/grpc"
)

func (s *RPCServer) GetActiveTorrentParts(ctx context.Context, r *pb.GetActiveTorrentPartsRequest) (*pb.ActiveTorrentPartsResponse, error) {
	files, err := s.torrentClient.GetTorrentFiles([]int{int(r.Id)})
	if err != nil {
		return nil, fmt.Errorf("get files: %w", err)
	}

	result := make([]*pb.ActiveTorrentPart, len(files))
	for i, file := range files {
		result[i] = &pb.ActiveTorrentPart{
			Name: file.Name,
		}
	}

	return &pb.ActiveTorrentPartsResponse{Parts: result}, nil
}

// nolint: revive
func GetActiveTorrentPartsHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.GetActiveTorrentPartsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(pb.BaseServiceServer).GetActiveTorrentParts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/GetActiveTorrentsParts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(pb.BaseServiceServer).GetActiveTorrentParts(ctx, req.(*pb.GetActiveTorrentPartsRequest))
	}
	return interceptor(ctx, in, info, handler)
}
