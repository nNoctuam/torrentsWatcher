package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"torrentsWatcher/internal/core/models"
	"torrentsWatcher/internal/core/storage"
	"torrentsWatcher/internal/core/torrentclient"
)

func GetTransmissionTorrents(
	torrentsStorage storage.Torrents,
	torrentClient torrentclient.Client,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := struct {
			OnlyRegistered bool
		}{OnlyRegistered: true}
		if r.URL.Query().Get("only-registered") != "false" {
			params.OnlyRegistered = false
		}

		var torrents []models.TransmissionTorrent
		err := torrentsStorage.GetAllTransmission(&torrents)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		activeTorrents, err := torrentClient.GetTorrents()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var result []torrentclient.Torrent
		for _, t := range activeTorrents {
			found := false
			if params.OnlyRegistered {
				for _, registeredTorrent := range torrents {
					if registeredTorrent.Hash == t.Hash {
						found = true
						break
					}
				}
			}
			if found || !params.OnlyRegistered {
				result = append(result, t)
			}
		}

		response, _ := json.Marshal(result)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, string(response))
	}
}
