package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"torrentsWatcher/internal/api/torrentclient"

	"torrentsWatcher/internal/api/tracking"
)

func DownloadWithClient(
	trackers tracking.Trackers,
	torrentClient *torrentclient.TorrentClient,
	folders map[string]string,
) func(w http.ResponseWriter, r *http.Request) {
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
		log.Println(folder, ok)

		torrent, err := trackers.GetTorrentInfo(requestBody.Url)
		if err != nil || torrent.FileUrl == "" {
			http.Error(w, "cannot get link to .torrent file", http.StatusUnprocessableEntity)
			return
		}

		name, content, err := trackers.DownloadTorrentFile(torrent)
		if err != nil {
			log.Println("cannot download .torrent file", err)
			http.Error(w, "cannot download .torrent file", http.StatusUnprocessableEntity)
			return
		}
		if name == "" {
			name = torrent.Title + ".torrent" // todo sanitize
		}

		if err := torrentClient.AddTorrent(content, folder); err != nil {
			log.Printf("cannot save .torrent file [%s]: %s", name, err)
			http.Error(w, "cannot add torrent: "+err.Error(), http.StatusUnprocessableEntity)
			return
		}
	}
}
