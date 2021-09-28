package grpc

import (
	"context"
	"torrentsWatcher/internal/core/models"
	"torrentsWatcher/internal/pb"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func (s *RpcServer) AddTorrent(ctx context.Context, r *pb.AddTorrentRequest) (*pb.TorrentResponse, error) {
	var torrent *models.Torrent
	s.logger.Info("parsing ", zap.String("url", r.Url))

	torrent, err := s.trackers.GetTorrentInfo(r.Url)
	if err != nil {
		return nil, err
	}

	_, file, err := s.trackers.DownloadTorrentFile(torrent)
	if err != nil {
		s.logger.Error("Failed to load torrent file", zap.Error(err), zap.String("url", torrent.FileURL))
		return nil, err
	}
	torrent.File = file

	err = s.torrentsStorage.Save(torrent)
	if err != nil {
		s.logger.Error("Failed to save torrent to storage", zap.Error(err))
		return nil, err
	}

	return &pb.TorrentResponse{Torrent: torrent.ToPB()}, nil
}

func AddTorrentHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.AddTorrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(pb.BaseServiceServer).AddTorrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/AddTorrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(pb.BaseServiceServer).AddTorrent(ctx, req.(*pb.AddTorrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}
