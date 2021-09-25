package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"torrentsWatcher/internal/core/torrentclient"
)

func GetTransmissionTorrentFiles(torrentClient torrentclient.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			ID int
		}
		var err error
		params.ID, err = strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := torrentClient.GetTorrentFiles([]int{params.ID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, _ := json.Marshal(result[0].Files)
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, string(response))
	}
}
