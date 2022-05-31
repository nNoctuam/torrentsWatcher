package grpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"torrentsWatcher/internal/models"
	"torrentsWatcher/internal/pb"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func (s *RPCServer) DeleteTorrent(ctx context.Context, r *pb.DeleteTorrentRequest) (*pb.Empty, error) {
	var torrents []models.Torrent
	err := s.torrentsStorage.Find(&torrents, models.Torrent{
		ID: uint(r.Id),
	})
	if err != nil {
		return nil, err
	}
	if len(torrents) == 0 {
		return nil, errors.New("torrent not found")
	}

	torrent := torrents[0]
	now := time.Now()
	torrent.DeletedAt = &now

	if err = s.torrentsStorage.Save(&torrent); err != nil {
		s.logger.Error("failed to update torrent", zap.Error(err))
		return nil, fmt.Errorf("failed to update torrent: %w", err)
	}

	return &pb.Empty{}, nil
}

// nolint: revive
func DeleteTorrentHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.DeleteTorrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(pb.BaseServiceServer).DeleteTorrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/DeleteTorrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(pb.BaseServiceServer).DeleteTorrent(ctx, req.(*pb.DeleteTorrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}
