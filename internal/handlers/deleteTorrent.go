package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"torrentsWatcher/internal/storage"

	"github.com/go-chi/chi"

	"torrentsWatcher/internal/api/models"
)

func DeleteTorrent(w http.ResponseWriter, r *http.Request, torrentsStorage storage.Torrents) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		_, _ = fmt.Fprintf(w, "invalid torrent id '%s'", chi.URLParam(r, "id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var torrents []models.Torrent
	err = torrentsStorage.Find(&torrents, models.Torrent{
		Id: uint(id),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(torrents) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	torrent := torrents[0]
	now := time.Now()
	torrent.DeletedAt = &now

	if err = torrentsStorage.Save(&torrent); err != nil {
		fmt.Println("error updating torrent", err)
		return
	}
}
