package grpc

import (
	"context"

	"torrentsWatcher/internal/models"
	"torrentsWatcher/internal/pb"

	"github.com/golang/protobuf/ptypes/timestamp"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func (s *RPCServer) GetMonitoredTorrents(ctx context.Context, r *pb.Empty) (*pb.TorrentsResponse, error) {
	var torrents []models.Torrent
	err := s.torrentsStorage.Find(&torrents, "")
	if err != nil {
		s.logger.Error("failed to get torrents", zap.Error(err))
		return nil, err
	}

	transformed := &pb.TorrentsResponse{}
	for _, torrent := range torrents {
		transformed.Torrents = append(transformed.Torrents, &pb.Torrent{
			Id:         uint32(torrent.ID),
			Title:      torrent.Title,
			PageUrl:    torrent.PageURL,
			FileUrl:    torrent.FileURL,
			CreatedAt:  &timestamp.Timestamp{Seconds: torrent.CreatedAt.Unix()},
			UpdatedAt:  &timestamp.Timestamp{Seconds: torrent.UpdatedAt.Unix()},
			UploadedAt: &timestamp.Timestamp{Seconds: torrent.UploadedAt.Unix()},
		})
	}

	return transformed, nil
}

// nolint: revive
func GetMonitoredTorrentsHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(pb.BaseServiceServer).GetMonitoredTorrents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/GetMonitoredTorrents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(pb.BaseServiceServer).GetMonitoredTorrents(ctx, req.(*pb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}
