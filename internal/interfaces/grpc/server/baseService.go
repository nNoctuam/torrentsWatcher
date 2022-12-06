package grpc

import (
	"torrentsWatcher/internal/connectors/torrentclient"
	"torrentsWatcher/internal/pb"
	"torrentsWatcher/internal/services/tracking"
	"torrentsWatcher/internal/storage"

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
	blockViewList   []string
}

func NewRPCServer(
	logger *zap.Logger,
	trackers tracking.Trackers,
	torrentsStorage storage.Torrents,
	downloadFolders map[string]string,
	torrentClient torrentclient.Client,
	blockViewList []string,
) *RPCServer {
	return &RPCServer{
		logger:          logger,
		trackers:        trackers,
		torrentsStorage: torrentsStorage,
		downloadFolders: downloadFolders,
		torrentClient:   torrentClient,
		blockViewList:   blockViewList,
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
