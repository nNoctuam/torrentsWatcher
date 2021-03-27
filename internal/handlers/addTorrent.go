package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/proto"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/tracking"
	"torrentsWatcher/internal/api/parser"
	"torrentsWatcher/internal/storage"
)

func AddTorrent(w http.ResponseWriter, r *http.Request, trackers tracking.Trackers, torrentsStorage storage.Torrents) {
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

	torrent, err = trackers.GetTorrentInfo(requestBody.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := trackers.DownloadTorrentFile(torrent)
	if err != nil {
		fmt.Printf("Failed to load torrent file '%s': %v", torrent.FileUrl, err)
		return
	}
	torrent.File = file

	err = torrentsStorage.Save(torrent)
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
