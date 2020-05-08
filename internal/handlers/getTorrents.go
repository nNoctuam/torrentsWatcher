package handlers

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"net/http"
	"torrentsWatcher/internal/api/db"
	"torrentsWatcher/internal/pb"

	"torrentsWatcher/internal/api/models"
)

func GetTorrents(w http.ResponseWriter, r *http.Request) {
	var torrents []models.Torrent
	err := db.DB.Find(&torrents).Error
	if err != nil {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//
	w.Header().Add("Content-Type", "application/protobuf")
	fmt.Fprint(w, string(response))
}
