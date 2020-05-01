package watch

import (
	"fmt"
	"log"
	"time"

	"torrentsWatcher/internal/api/db"
	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/notification"
	"torrentsWatcher/internal/api/parsing"
)

func Watch(period time.Duration) {
	fmt.Printf("Start checking every %s\n", period)
	for {
		go func() {
			var torrents []models.Torrent
			err := db.DB.Find(&torrents).Error
			if err != nil {
				log.Print("Couldn't get torrents for check")
			}
			for _, torrent := range torrents {
				checkTorrent(&torrent)
			}
		}()

		time.Sleep(period)
	}
}

func checkTorrent(torrent *models.Torrent) {
	updatedTorrent, err := parsing.GetTorrentInfo(torrent.PageUrl)

	if err != nil {
		log.Print("Error parsing torrent: ", err)
		return
	}

	isUpdated := torrent.UploadedAt.Unix() != updatedTorrent.UploadedAt.Unix()

	if isUpdated || torrent.FileUrl != "" && torrent.File == nil {
		file, err := parsing.DownloadTorrentFile(torrent)
		if err != nil {
			fmt.Printf("Failed to load torrent file '%s': %v", torrent.FileUrl, err)
			return
		}
		torrent.File = file
	}

	err = torrent.UpdateFrom(updatedTorrent)
	if err != nil {
		log.Printf("Couldn't save torrent: %v", updatedTorrent)
		return
	}

	if isUpdated {
		log.Printf("torrent '%s' (%s) was updated!", torrent.Title, torrent.PageUrl)
		notification.NotifyAbout(torrent)
	}
}
