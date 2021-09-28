package grpc

import (
	"torrentsWatcher/internal/core/storage"
	tracking2 "torrentsWatcher/internal/core/tracking"
	"torrentsWatcher/internal/pb"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

type RpcServer struct {
	pb.BaseServiceServer
	logger          *zap.Logger
	trackers        tracking2.Trackers
	torrentsStorage storage.Torrents
	downloadFolders map[string]string
}

func NewRpcServer(
	logger *zap.Logger,
	trackers tracking2.Trackers,
	torrentsStorage storage.Torrents,
	downloadFolders map[string]string,
) *RpcServer {
	return &RpcServer{
		logger:          logger,
		trackers:        trackers,
		torrentsStorage: torrentsStorage,
		downloadFolders: downloadFolders,
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
	},
	Metadata: "protobuf/baseService.proto",
}
