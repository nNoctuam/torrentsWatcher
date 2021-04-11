package handlers

import (
	"encoding/json"
	"net/http"
	"torrentsWatcher/internal/api/torrentclient"

	"go.uber.org/zap"
)

func Rename(
	logger *zap.Logger,
	torrentClient *torrentclient.TorrentClient,
) func(w http.ResponseWriter, r *http.Request) {
	logger = logger.With(zap.String("method", "Rename"))
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Hash    string
			NewName string
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		logger.Debug("request", zap.Any("request", requestBody))

		torrents, err := torrentClient.GetTorrents()
		if err != nil {
			logger.Error("failed to get torrents from transmission", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, t := range torrents {
			if t.Hash == requestBody.Hash {
				err = torrentClient.Rename(t.Id, t.Name, requestBody.NewName)
				if err != nil {
					logger.Error("failed to rename torrent", zap.Error(err), zap.String("oldName", t.Name), zap.String("newName", requestBody.NewName))
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
		}

		logger.Error("torrent not found", zap.String("hash", requestBody.Hash), zap.String("name", requestBody.NewName))
		http.Error(w, "Torrent not found", http.StatusInternalServerError)
	}
}
