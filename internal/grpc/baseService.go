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
}

func NewRpcServer(
	logger *zap.Logger,
	trackers tracking2.Trackers,
	torrentsStorage storage.Torrents,
) *RpcServer {
	return &RpcServer{logger: logger, trackers: trackers, torrentsStorage: torrentsStorage}
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
	},
	Metadata: "protobuf/baseService.proto",
}
