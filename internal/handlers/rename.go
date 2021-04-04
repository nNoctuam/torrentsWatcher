package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"torrentsWatcher/internal/api/torrentclient"
)

func Rename(
	torrentClient *torrentclient.TorrentClient,
) func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Printf("%+v", requestBody)

		torrents, err := torrentClient.GetTorrents()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, t := range torrents {
			if t.Hash == requestBody.Hash {
				err = torrentClient.Rename(t.Id, t.Name, requestBody.NewName)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
		}

		http.Error(w, "Torrent not found", http.StatusInternalServerError)
	}
}
