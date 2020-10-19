package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/parser"
	"torrentsWatcher/internal/storage"
)

func AddTorrent(w http.ResponseWriter, r *http.Request, parsers []*parser.Tracker, torrentsStorage storage.Torrents) {
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

	err = torrentsStorage.Save(torrent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(torrent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(response))
}
