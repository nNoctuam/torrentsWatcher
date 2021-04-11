package handlers

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"torrentsWatcher/internal/pb"
	"torrentsWatcher/internal/storage"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/proto"

	"torrentsWatcher/internal/api/models"
)

func GetTorrents(logger *zap.Logger, torrentsStorage storage.Torrents) func(w http.ResponseWriter, r *http.Request) {
	logger = logger.With(zap.String("method", "GetTorrents"))
	return func(w http.ResponseWriter, r *http.Request) {
		var torrents []models.Torrent
		err := torrentsStorage.Find(&torrents, "")
		if err != nil {
			logger.Error("failed to get torrents", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		transformed := &pb.Torrents{}
		for _, torrent := range torrents {
			transformed.Torrents = append(transformed.Torrents, &pb.Torrent{
				Id:         uint32(torrent.Id),
				Title:      torrent.Title,
				PageUrl:    torrent.PageUrl,
				FileUrl:    torrent.FileUrl,
				CreatedAt:  &timestamp.Timestamp{Seconds: torrent.CreatedAt.Unix()},
				UpdatedAt:  &timestamp.Timestamp{Seconds: torrent.UpdatedAt.Unix()},
				UploadedAt: &timestamp.Timestamp{Seconds: torrent.UploadedAt.Unix()},
			})
		}

		response, err := proto.Marshal(transformed)
		if err != nil {
			logger.Error("failed to marshall torrents", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//
		w.Header().Add("Content-Type", "application/protobuf")
		_, _ = fmt.Fprint(w, string(response))
	}
}
