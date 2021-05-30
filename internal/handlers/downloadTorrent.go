package handlers

import (
	"fmt"
	"mime"
	"net/http"
	"strconv"
	models2 "torrentsWatcher/internal/core/models"
	storage2 "torrentsWatcher/internal/core/storage"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func DownloadTorrent(logger *zap.Logger, torrentsStorage storage2.Torrents) func(w http.ResponseWriter, r *http.Request) {
	logger = logger.With(zap.String("method", "DownloadTorrent"))
	return func(w http.ResponseWriter, r *http.Request) {
		var torrent models2.Torrent
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			_, _ = fmt.Fprintf(w, "invalid torrent id '%s'", chi.URLParam(r, "id"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = torrentsStorage.First(&torrent, models2.Torrent{Id: uint(id)}); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Add("Content-Type", "application/x-bittorrent")
		w.Header().Add("Content-Length", fmt.Sprintf("%d", len(torrent.File)))
		w.Header().Add("Content-Disposition", mime.FormatMediaType("attachment", map[string]string{"filename": torrent.Title + ".torrent"}))

		_, err = w.Write(torrent.File)
		if err != nil {
			logger.Error("failed to write torrent file", zap.Error(err))
			return
		}
	}
}
