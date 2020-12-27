package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"torrentsWatcher/internal/api/db"
	"torrentsWatcher/internal/pb"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/golang/protobuf/proto"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/parser"
)

func AddTorrent(w http.ResponseWriter, r *http.Request, parsers []*parser.Tracker) {
	var torrent *models.Torrent
	var requestBody struct {
		Url string
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("parsing %s\n", requestBody.Url)

	torrent, err = parser.GetTorrentInfo(requestBody.Url, parsers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.DB.Create(torrent).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := proto.Marshal(&pb.Torrent{
		Id:         uint32(torrent.Id),
		Title:      torrent.Title,
		PageUrl:    torrent.PageUrl,
		FileUrl:    torrent.FileUrl,
		CreatedAt:  &timestamp.Timestamp{Seconds: torrent.CreatedAt.Unix()},
		UpdatedAt:  &timestamp.Timestamp{Seconds: torrent.UpdatedAt.Unix()},
		UploadedAt: &timestamp.Timestamp{Seconds: torrent.UploadedAt.Unix()},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/protobuf")
	fmt.Fprint(w, string(response))
}
