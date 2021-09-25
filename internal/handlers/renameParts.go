package handlers

import (
	"encoding/json"
	"net/http"
	"torrentsWatcher/internal/core/torrentclient"

	zap "go.uber.org/zap"
)

func RenameParts(
	logger *zap.Logger,
	torrentClient torrentclient.Client,
) func(w http.ResponseWriter, r *http.Request) {
	logger = logger.With(zap.String("method", "Rename"))
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			ID    int        `json:"id"`
			Names [][]string `json:"names"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		logger.Debug("request", zap.Any("request", requestBody))

		for _, n := range requestBody.Names {
			oldName, newName := n[0], n[1]
			err = torrentClient.Rename(requestBody.ID, oldName, newName)
			if err != nil {
				logger.Error(
					"failed to rename torrent",
					zap.Error(err),
					zap.Int("ID", requestBody.ID),
					zap.String("oldName", oldName),
					zap.String("newName", newName),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
