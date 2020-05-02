package handlers

import (
	"fmt"
	"mime"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"torrentsWatcher/internal/api/db"
	"torrentsWatcher/internal/api/models"
)

func DownloadTorrent(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Add("Content-Type", "application/x-bittorrent")
	w.Header().Add("Content-Length", fmt.Sprintf("%d", len(torrent.File)))
	w.Header().Add("Content-Disposition", mime.FormatMediaType("attachment", map[string]string{"filename": torrent.Title + ".torrent"}))

	_, err = w.Write(torrent.File)
	if err != nil {
		fmt.Println("error writing torrent file")
		return
	}
}
