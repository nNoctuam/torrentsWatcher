package watch

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"torrentsWatcher/internal/api/db"
	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/notification"
	"torrentsWatcher/internal/api/tracking"
)

func Run(ctx context.Context, wg *sync.WaitGroup, period time.Duration, trackers tracking.Trackers, notificator notification.Notificator) {
	fmt.Printf("Start checking every %s\n", period)
	ticker := time.After(0)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-ticker:
			var torrents []models.Torrent
			err := db.DB.Find(&torrents).Error
			if err != nil {
				log.Print("Couldn't get torrents for check")
			}
			for _, torrent := range torrents {
				checkTorrent(&torrent, trackers, notificator)
			}
			ticker = time.After(period)
		}
	}
}

func checkTorrent(torrent *models.Torrent, trackers tracking.Trackers, notificator notification.Notificator) {
	updatedTorrent, err := trackers.GetTorrentInfo(torrent.PageUrl)

	if err != nil {
		log.Print("Error parsing torrent: ", err)
		return
	}

	isUpdated := torrent.UploadedAt.Unix() != updatedTorrent.UploadedAt.Unix()

	if isUpdated || torrent.FileUrl != "" && torrent.File == nil {
		file, err := trackers.DownloadTorrentFile(torrent)
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
		notification.NotifyAbout(torrent, notificator)
	}
}
