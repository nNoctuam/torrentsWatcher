package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"torrentsWatcher/internal/api/db"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/parsing"
)

func AddTorrent(w http.ResponseWriter, r *http.Request) {
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

	torrent, err = parsing.GetTorrentInfo(requestBody.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.DB.Create(torrent).Error
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
