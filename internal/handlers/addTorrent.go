package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"google.golang.org/protobuf/proto"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/tracking"
	"torrentsWatcher/internal/storage"
)

func AddTorrent(logger *zap.Logger, trackers tracking.Trackers, torrentsStorage storage.Torrents) func(w http.ResponseWriter, r *http.Request) {
	logger = logger.With(zap.String("method", "AddTorrent"))
	return func(w http.ResponseWriter, r *http.Request) {
		var torrent *models.Torrent
		var requestBody struct {
			Url string
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Info("parsing ", zap.String("url", requestBody.Url))

		torrent, err = trackers.GetTorrentInfo(requestBody.Url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, file, err := trackers.DownloadTorrentFile(torrent)
		if err != nil {
			logger.Error("Failed to load torrent file", zap.Error(err), zap.String("url", torrent.FileUrl))
			return
		}
		torrent.File = file

		err = torrentsStorage.Save(torrent)
		if err != nil {
			logger.Error("Failed to save torrent to storage", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := proto.Marshal(torrent.ToPB())
		if err != nil {
			logger.Error("Failed to marshal response", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/protobuf")
		_, _ = fmt.Fprint(w, string(response))
	}
}
