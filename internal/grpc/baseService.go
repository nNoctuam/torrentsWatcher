package grpc

import (
	"torrentsWatcher/internal/core/storage"
	"torrentsWatcher/internal/core/torrentclient"
	"torrentsWatcher/internal/services/tracking"
	"torrentsWatcher/internal/pb"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

type RPCServer struct {
	pb.BaseServiceServer
	logger          *zap.Logger
	trackers        tracking.Trackers
	torrentsStorage storage.Torrents
	downloadFolders map[string]string
	torrentClient   torrentclient.Client
}

func NewRPCServer(
	logger *zap.Logger,
	trackers tracking.Trackers,
	torrentsStorage storage.Torrents,
	downloadFolders map[string]string,
	torrentClient torrentclient.Client,
) *RPCServer {
	return &RPCServer{
		logger:          logger,
		trackers:        trackers,
		torrentsStorage: torrentsStorage,
		downloadFolders: downloadFolders,
		torrentClient:   torrentClient,
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
