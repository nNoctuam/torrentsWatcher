package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"torrentsWatcher/internal/api/db"

	"google.golang.org/protobuf/proto"

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

	file, err := parser.DownloadTorrentFile(torrent, parsers)
	if err != nil {
		fmt.Printf("Failed to load torrent file '%s': %v", torrent.FileUrl, err)
		return
	}
	torrent.File = file

	err = db.DB.Create(torrent).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := proto.Marshal(torrent.ToPB())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/protobuf")
	fmt.Fprint(w, string(response))
}
