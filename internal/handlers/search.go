package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"torrentsWatcher/internal/pb"

	"github.com/golang/protobuf/proto"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/parser"
)

func Search(w http.ResponseWriter, r *http.Request, parsers []*parser.Tracker) {
	var requestBody struct {
		Text string
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var torrents []*models.Torrent

	for _, p := range parsers {
		found, _ := p.Search(requestBody.Text)
		torrents = append(torrents, found...)
	}

	response, err := proto.Marshal(&pb.Torrents{
		Torrents: models.TorrentsToPB(torrents),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/protobuf")
	fmt.Fprint(w, string(response))
}
