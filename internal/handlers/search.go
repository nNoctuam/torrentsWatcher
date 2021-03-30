package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"torrentsWatcher/internal/pb"

	"google.golang.org/protobuf/proto"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/tracking"
)

func Search(trackers tracking.Trackers) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Text string
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		torrents := trackers.SearchEverywhere(requestBody.Text)

		sort.Slice(torrents, func(i, j int) bool {
			return torrents[i].Seeders > torrents[j].Seeders
		})

		response, err := proto.Marshal(&pb.Torrents{
			Torrents: models.TorrentsToPB(torrents),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/protobuf")
		fmt.Fprint(w, string(response))
	}
}
