package grpc

import (
	"context"

	"torrentsWatcher/internal/ports/pb"

	"google.golang.org/grpc"
)

func (s *RPCServer) GetDownloadFolders(ctx context.Context, in *pb.Empty) (*pb.DownloadFoldersResponse, error) {
	folders := make([]string, len(s.downloadFolders))
	i := 0
	for folder := range s.downloadFolders {
		folders[i] = folder
		i++
	}

	return &pb.DownloadFoldersResponse{Folders: folders}, nil
}

// nolint: revive
func GetDownloadFoldersHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(pb.BaseServiceServer).GetDownloadFolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BaseService/GetDownloadFolders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(pb.BaseServiceServer).GetDownloadFolders(ctx, req.(*pb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}
