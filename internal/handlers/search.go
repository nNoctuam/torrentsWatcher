package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	models2 "torrentsWatcher/internal/core/models"
	tracking2 "torrentsWatcher/internal/core/tracking"
	"torrentsWatcher/internal/pb"

	"go.uber.org/zap"

	"google.golang.org/protobuf/proto"
)

func Search(logger *zap.Logger, trackers tracking2.Trackers) func(w http.ResponseWriter, r *http.Request) {
	logger = logger.With(zap.String("method", "Search"))
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

		response, err := proto.Marshal(&pb.TorrentsResponse{
			Torrents: models2.TorrentsToPB(torrents),
		})
		if err != nil {
			logger.Error("failed to marshall torrents", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/protobuf")
		_, _ = fmt.Fprint(w, string(response))
	}
}
