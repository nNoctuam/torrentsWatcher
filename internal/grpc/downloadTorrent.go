package grpc

import (
	"context"
	"fmt"
	"torrentsWatcher/internal/core/models"
	"torrentsWatcher/internal/pb"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func (s *RpcServer) DownloadTorrent(ctx context.Context, r *pb.DownloadTorrentRequest) (*pb.DownloadTorrentResponse, error) {
	folder, ok := s.downloadFolders[r.Folder]
	s.logger.Debug(
		"folders matching",
		zap.String("folderName", r.Folder),
		zap.String("path", folder),
		zap.Bool("found", ok),
	)

	torrent, err := s.trackers.GetTorrentInfo(r.Url)
	if err != nil || torrent.FileURL == "" {
		s.logger.Error("failed to get link to .torrent file", zap.Error(err))
		return nil, fmt.Errorf("cannot get link to .torrent file: %w", err)
	}

	_, content, err := s.trackers.DownloadTorrentFile(torrent)
	if err != nil {
		s.logger.Error("failed to download .torrent file", zap.Error(err))
		return nil, fmt.Errorf("cannot download .torrent file: %w", err)
	}

	addedTorrent, err := s.torrentClient.AddTorrent(content, folder, false)
	if err != nil {
		s.logger.Error("failed to add .torrent to client", zap.Error(err), zap.String("name", addedTorrent.Name))
		return nil, fmt.Errorf("cannot add torrent: %w", err)
	}
	transmissionTorrent := &models.TransmissionTorrent{
		Hash: addedTorrent.Hash,
	}
	err = s.torrentsStorage.SaveTransmission(transmissionTorrent)
	if err != nil {
		s.logger.Error("failed to save torrent to storage", zap.Error(err))
		return nil, fmt.Errorf("cannot save transmissionTorrent: %w", err)
	}

	return &pb.DownloadTorrentResponse{
		ID:   int32(addedTorrent.ID),
		Name: addedTorrent.Name,
		Hash: addedTorrent.Hash,
	}, nil
}

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
