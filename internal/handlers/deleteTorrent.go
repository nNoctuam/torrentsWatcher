package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"torrentsWatcher/internal/api/db"
	"torrentsWatcher/internal/api/models"
)

func DeleteTorrent(w http.ResponseWriter, r *http.Request) {
	var torrent models.Torrent
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		_, _ = fmt.Fprintf(w, "invalid torrent id '%s'", chi.URLParam(r, "id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = db.DB.First(&torrent, models.Torrent{Id: uint(id)}).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	now := time.Now()
	torrent.DeletedAt = &now

	if err = db.DB.Save(&torrent).Error; err != nil {
		fmt.Println("error updating torrent", err)
		return
	}
}
