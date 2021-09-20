package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	models2 "torrentsWatcher/internal/core/models"
	storage2 "torrentsWatcher/internal/core/storage"
	tracking2 "torrentsWatcher/internal/core/tracking"

	"go.uber.org/zap"

	"google.golang.org/protobuf/proto"
)

func AddTorrent(
	logger *zap.Logger,
	trackers tracking2.Trackers,
	torrentsStorage storage2.Torrents,
) func(w http.ResponseWriter, r *http.Request) {
	logger = logger.With(zap.String("method", "AddTorrent"))
	return func(w http.ResponseWriter, r *http.Request) {
		var torrent *models2.Torrent
		var requestBody struct {
			URL string
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Info("parsing ", zap.String("url", requestBody.URL))

		torrent, err = trackers.GetTorrentInfo(requestBody.URL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, file, err := trackers.DownloadTorrentFile(torrent)
		if err != nil {
			logger.Error("Failed to load torrent file", zap.Error(err), zap.String("url", torrent.FileURL))
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
