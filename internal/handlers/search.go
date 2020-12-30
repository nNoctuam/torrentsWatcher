package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"torrentsWatcher/internal/pb"

	"google.golang.org/protobuf/proto"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/parser"
)

func Search(w http.ResponseWriter, r *http.Request, parsers []*parser.Tracker) {
	var requestBody struct {
		Text string
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var torrents []*models.Torrent

	wg := sync.WaitGroup{}
	tChan := make(chan []*models.Torrent)
	for _, p := range parsers {
		wg.Add(1)
		go func(p *parser.Tracker) {
			found, _ := p.Search(requestBody.Text)
			tChan <- found
		}(p)
	}

	q := make(chan interface{})
	go func() {
		for {
			select {
			case t := <-tChan:
				torrents = append(torrents, t...)
				wg.Done()
			case <-q:
				return
			}
		}
	}()

	wg.Wait()
	q <- nil

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
