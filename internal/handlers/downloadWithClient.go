package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/torrentclient"
	"torrentsWatcher/internal/storage"

	"torrentsWatcher/internal/api/tracking"

	"go.uber.org/zap"
)

func DownloadWithClient(
	logger *zap.Logger,
	trackers tracking.Trackers,
	torrentClient *torrentclient.TorrentClient,
	torrentsStorage storage.Torrents,
	folders map[string]string,
) func(w http.ResponseWriter, r *http.Request) {
	logger = logger.With(zap.String("method", "DownloadWithClient"))
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Url    string
			Folder string
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		folder, ok := folders[requestBody.Folder]
		logger.Debug("folders matching", zap.String("folderName", requestBody.Folder), zap.String("path", folder), zap.Bool("found", ok))

		torrent, err := trackers.GetTorrentInfo(requestBody.Url)
		if err != nil || torrent.FileUrl == "" {
			logger.Error("failed to get link to .torrent file", zap.Error(err))
			http.Error(w, "cannot get link to .torrent file", http.StatusUnprocessableEntity)
			return
		}

		_, content, err := trackers.DownloadTorrentFile(torrent)
		if err != nil {
			logger.Error("failed to download .torrent file", zap.Error(err))
			http.Error(w, "cannot download .torrent file", http.StatusUnprocessableEntity)
			return
		}

		hash, name, err := torrentClient.AddTorrent(content, folder)
		if err != nil {
			logger.Error("failed to add .torrent to client", zap.Error(err), zap.String("name", name))
			http.Error(w, "cannot add torrent: "+err.Error(), http.StatusUnprocessableEntity)
			return
		}
		transmissionTorrent := &models.TransmissionTorrent{
			Hash: hash,
		}
		err = torrentsStorage.SaveTransmission(transmissionTorrent)
		if err != nil {
			logger.Error("failed to save torrent to storage", zap.Error(err))
			http.Error(w, "cannot save transmissionTorrent: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response, _ := json.Marshal(map[string]string{
			"hash": hash,
			"name": name,
		})

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprintln(w, string(response))
	}
}
