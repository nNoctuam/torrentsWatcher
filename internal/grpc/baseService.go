package grpc

import (
	tracking2 "torrentsWatcher/internal/core/tracking"
	"torrentsWatcher/internal/pb"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

type RpcServer struct {
	pb.BaseServiceServer
	logger   *zap.Logger
	trackers tracking2.Trackers
}

func NewRpcServer(
	logger *zap.Logger,
	trackers tracking2.Trackers,
) *RpcServer {
	return &RpcServer{logger: logger, trackers: trackers}
}

var BaseServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.BaseService",
	HandlerType: (*pb.BaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    SearchHandler,
		},
	},
	Metadata: "protobuf/baseService.proto",
}
