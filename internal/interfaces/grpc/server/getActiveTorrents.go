package grpc

import (
	"context"
	"fmt"
	"strings"

	"torrentsWatcher/internal/core/models"
	"torrentsWatcher/internal/pb"

	"github.com/golang/protobuf/ptypes/timestamp"

	"google.golang.org/grpc"
)

func (s *RPCServer) GetActiveTorrents(ctx context.Context, r *pb.GetActiveTorrentsRequest) (*pb.ActiveTorrentsResponse, error) {
	var torrents []models.TransmissionTorrent
	err := s.torrentsStorage.GetAllTransmission(&torrents)
	if err != nil {
		return nil, fmt.Errorf("get registered: %w", err)
	}

	activeTorrents, err := s.torrentClient.GetTorrents()
	if err != nil {
		return nil, fmt.Errorf("get from client: %w", err)
	}

	var result []*pb.ActiveTorrent
	for _, t := range activeTorrents {
		blocked := false
		for _, path := range s.blockViewList {
			if strings.Contains(t.DownloadDir+t.Name, path) {
				blocked = true
			}
		}
		if blocked {
			continue
		}

		found := false
		if r.OnlyRegistered {
			for _, registeredTorrent := range torrents {
				if registeredTorrent.Hash == t.Hash {
					found = true
					break
				}
			}
		}
		if found || !r.OnlyRegistered {
			result = append(result, &pb.ActiveTorrent{
				ID:          int32(t.ID),
				Name:        t.Name,
				Hash:        t.Hash,
				Comment:     t.Comment,
				DownloadDir: t.DownloadDir,
				DateCreated: &timestamp.Timestamp{Seconds: t.DateCreated.Unix()},
			})
		}
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
