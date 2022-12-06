package grpc

import (
	"context"
	"fmt"

	"torrentsWatcher/internal/ports/adapters/driving/pb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func (s *RPCServer) DownloadTorrent(ctx context.Context, r *pb.DownloadTorrentRequest) (*pb.DownloadTorrentResponse, error) {
	addedTorrent, err := s.torrents.Download(r.Url, r.Folder)
	if err != nil {
		s.logger.Warn("cannot download torrent", zap.Error(err))
		return nil, fmt.Errorf("cannot download torrent: %w", err)
	}

	return &pb.DownloadTorrentResponse{
		ID:   int32(addedTorrent.ID),
		Name: addedTorrent.Name,
		Hash: addedTorrent.Hash,
	}, nil
}

// nolint: revive
func DownloadTorrentHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.DownloadTorrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(pb.BaseServiceServer).DownloadTorrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/DownloadTorrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(pb.BaseServiceServer).DownloadTorrent(ctx, req.(*pb.DownloadTorrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}
