package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"torrentsWatcher/internal/api/torrentclient"

	"torrentsWatcher/internal/api/tracking"
)

func DownloadWithClient(
	w http.ResponseWriter,
	r *http.Request,
	trackers tracking.Trackers,
	torrentClient *torrentclient.TorrentClient,
) {
	var requestBody struct {
		Url string
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	torrent, err := trackers.GetTorrentInfo(requestBody.Url)
	if err != nil || torrent.FileUrl == "" {
		http.Error(w, "cannot get link to .torrent file", http.StatusUnprocessableEntity)
		return
	}

	name, content, err := trackers.DownloadTorrentFile(torrent)
	if err != nil {
		http.Error(w, "cannot download .torrent file", http.StatusUnprocessableEntity)
		return
	}
	if name == "" {
		name = torrent.Title + ".torrent" // todo sanitize
	}

	if err := torrentClient.StartDownload(name, content); err != nil {
		log.Printf("cannot save .torrent file [%s]: %s", name, err)
		http.Error(w, "cannot save .torrent file to auto-download dir", http.StatusUnprocessableEntity)
		return
	}

}
