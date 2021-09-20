package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	models2 "torrentsWatcher/internal/core/models"
	storage2 "torrentsWatcher/internal/core/storage"

	"go.uber.org/zap"

	"github.com/go-chi/chi"
)

func DeleteTorrent(logger *zap.Logger, torrentsStorage storage2.Torrents) func(w http.ResponseWriter, r *http.Request) {
	logger = logger.With(zap.String("method", "DeleteTorrent"))
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			_, _ = fmt.Fprintf(w, "invalid torrent id '%s'", chi.URLParam(r, "id"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var torrents []models2.Torrent
		err = torrentsStorage.Find(&torrents, models2.Torrent{
			ID: uint(id),
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
			logger.Error("failed to update torrent", zap.Error(err))
			return
		}
	}
}
