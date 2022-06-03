package grpc

import (
	"context"
	"fmt"

	"torrentsWatcher/internal/domain/services/torrents"
	"torrentsWatcher/internal/ports/adapters/driving/pb"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func (s *RPCServer) RenameTorrentParts(ctx context.Context, r *pb.RenameTorrentPartsRequest) (*pb.Empty, error) {
	var parts []*torrents.PartToRename
	for _, n := range r.Names {
		parts = append(parts, &torrents.PartToRename{OldName: n.OldName, NewName: n.NewName})
	}

	if err := s.torrents.RenameParts(uint(r.Id), parts); err != nil {
		s.logger.Warn("cannot rename parts", zap.Error(err))
		return nil, fmt.Errorf("cannot rename parts: %w", err)
	}

	return &pb.Empty{}, nil
}

// nolint: revive
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
