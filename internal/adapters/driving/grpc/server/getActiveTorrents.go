package grpc

import (
	"context"
	"fmt"

	"torrentsWatcher/internal/ports/adapters/driving/pb"

	"github.com/golang/protobuf/ptypes/timestamp"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func (s *RPCServer) GetActiveTorrents(ctx context.Context, r *pb.GetActiveTorrentsRequest) (*pb.ActiveTorrentsResponse, error) {
	torrents, err := s.torrents.GetActive(r.OnlyRegistered)
	if err != nil {
		s.logger.Warn("cannot get active torrents", zap.Error(err))
		return nil, fmt.Errorf("cannot get active torrents: %w", err)
	}

	result := make([]*pb.ActiveTorrent, len(torrents))
	for _, t := range torrents {
		result = append(result, &pb.ActiveTorrent{
			ID:          int32(t.ID),
			Name:        t.Name,
			Hash:        t.Hash,
			Comment:     t.Comment,
			DownloadDir: t.DownloadDir,
			DateCreated: &timestamp.Timestamp{Seconds: t.DateCreated.Unix()},
		})
	}

	return &pb.ActiveTorrentsResponse{Torrents: result}, nil
}

// nolint: revive
func GetActiveTorrentsHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.GetActiveTorrentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(pb.BaseServiceServer).GetActiveTorrents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/GetActiveTorrents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(pb.BaseServiceServer).GetActiveTorrents(ctx, req.(*pb.GetActiveTorrentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}
