package grpc

import (
	"torrentsWatcher/internal/domain/services/torrents"
	"torrentsWatcher/internal/ports/adapters/driving/pb"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

type RPCServer struct {
	pb.BaseServiceServer
	logger   *zap.Logger
	torrents *torrents.Torrents
}

func NewRPCServer(
	logger *zap.Logger,
	torrents *torrents.Torrents,
) *RPCServer {
	return &RPCServer{
		logger:   logger,
		torrents: torrents,
	}
}

var BaseServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.BaseService",
	HandlerType: (*pb.BaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    SearchHandler,
		},
		{
			MethodName: "GetMonitoredTorrents",
			Handler:    GetMonitoredTorrentsHandler,
		},
		{
			MethodName: "GetDownloadFolders",
			Handler:    GetDownloadFoldersHandler,
		},
		{
			MethodName: "AddTorrent",
			Handler:    AddTorrentHandler,
		},
		{
			MethodName: "DeleteTorrent",
			Handler:    DeleteTorrentHandler,
		},
		{
			MethodName: "DownloadTorrent",
			Handler:    DownloadTorrentHandler,
		},
		{
			MethodName: "RenameTorrentParts",
			Handler:    RenameTorrentPartsHandler,
		},
		{
			MethodName: "GetActiveTorrents",
			Handler:    GetActiveTorrentsHandler,
		},
		{
			MethodName: "GetActiveTorrentParts",
			Handler:    GetActiveTorrentPartsHandler,
		},
	},
	Metadata: "protobuf/baseService.proto",
}
