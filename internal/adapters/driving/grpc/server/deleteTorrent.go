package grpc

import (
	"context"

	"torrentsWatcher/internal/ports/adapters/driving/pb"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func (s *RPCServer) DeleteTorrent(ctx context.Context, r *pb.DeleteTorrentRequest) (*pb.Empty, error) {
	err := s.torrents.Delete(uint(r.Id))
	if err != nil {
		s.logger.Error("failed to delete torrent", zap.Uint("id", uint(r.Id)), zap.Error(err))
		return nil, err
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
