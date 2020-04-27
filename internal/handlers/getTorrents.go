package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"torrentsWatcher/internal/api/db"

	"torrentsWatcher/internal/api/models"
)

func GetTorrents(w http.ResponseWriter, r *http.Request) {
	var torrents []models.Torrent
	err := db.DB.Find(&torrents).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(torrents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(response))
}
