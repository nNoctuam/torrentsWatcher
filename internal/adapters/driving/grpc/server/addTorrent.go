package grpc

import (
	"context"

	"torrentsWatcher/internal/ports/adapters/driving/pb"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func (s *RPCServer) AddTorrent(ctx context.Context, r *pb.AddTorrentRequest) (*pb.TorrentResponse, error) {
	torrent, err := s.torrents.Add(r.Url)
	if err != nil {
		s.logger.Warn("cannot add torrent", zap.String("url", r.Url), zap.Error(err))
		return nil, err
	}

	return &pb.TorrentResponse{Torrent: pb.TorrentToPB(torrent)}, nil
}

// nolint: revive
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
