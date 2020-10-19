package handlers

import (
	"fmt"
	"mime"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/storage"
)

func DownloadTorrent(w http.ResponseWriter, r *http.Request, torrentsStorage storage.Torrents) {
	var torrent models.Torrent
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		_, _ = fmt.Fprintf(w, "invalid torrent id '%s'", chi.URLParam(r, "id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = torrentsStorage.First(&torrent, models.Torrent{Id: uint(id)}); err != nil {
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
