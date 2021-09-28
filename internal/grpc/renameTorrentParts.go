package grpc

import (
	"context"
	"fmt"
	"torrentsWatcher/internal/pb"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func (s *RpcServer) RenameTorrentParts(ctx context.Context, r *pb.RenameTorrentPartsRequest) (*pb.Empty, error) {
	var err error
	for _, pair := range r.Names {
		err = s.torrentClient.Rename(int(r.Id), pair.OldName, pair.NewName)
		if err != nil {
			s.logger.Error(
				"failed to rename torrent",
				zap.Error(err),
				zap.Int("ID", int(r.Id)),
				zap.String("oldName", pair.OldName),
				zap.String("newName", pair.NewName),
			)
			return nil, fmt.Errorf("rename part: %w", err)
		}
	}
	return &pb.Empty{}, nil
}

func RenameTorrentPartsHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.RenameTorrentPartsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(pb.BaseServiceServer).RenameTorrentParts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/RenameTorrentParts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(pb.BaseServiceServer).RenameTorrentParts(ctx, req.(*pb.RenameTorrentPartsRequest))
	}
	return interceptor(ctx, in, info, handler)
}
