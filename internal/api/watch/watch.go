package watch

import (
	"fmt"
	"log"
	"time"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/notification"
	"torrentsWatcher/internal/api/parser"
	"torrentsWatcher/internal/storage"
)

func Watch(
	period time.Duration,
	parsers []*parser.Tracker,
	notificator notification.Notificator,
	torrentsStorage storage.Torrents,
	cookiesStorage storage.Cookies,
) {
	fmt.Printf("Start checking every %s\n", period)
	for {
		go func() {
			var torrents []models.Torrent
			err := torrentsStorage.Find(&torrents, nil)
			if err != nil {
				log.Print("Couldn't get torrents for check")
			}
			for _, torrent := range torrents {
				checkTorrent(&torrent, parsers, notificator, torrentsStorage)
			}
		}()

		time.Sleep(period)
	}
}

func checkTorrent(
	torrent *models.Torrent,
	parsers []*parser.Tracker,
	notificator notification.Notificator,
	torrentsStorage storage.Torrents,
) {
	updatedTorrent, err := parser.GetTorrentInfo(torrent.PageUrl, parsers)

	if err != nil {
		log.Print("Error parsing torrent: ", err)
		return
	}

	isUpdated := torrent.UploadedAt.Unix() != updatedTorrent.UploadedAt.Unix()

	if isUpdated || torrent.FileUrl != "" && torrent.File == nil {
		file, err := parser.DownloadTorrentFile(torrent, parsers)
		if err != nil {
			fmt.Printf("Failed to load torrent file '%s': %v", torrent.FileUrl, err)
			return
		}
		torrent.File = file
	}

	torrent.UpdateFrom(updatedTorrent)
	err = torrentsStorage.Save(torrent)
	if err != nil {
		log.Printf("Couldn't save torrent: %v", updatedTorrent)
		return
	}

	if isUpdated {
		log.Printf("torrent '%s' (%s) was updated!", torrent.Title, torrent.PageUrl)
		notification.NotifyAbout(torrent, notificator)
	}
}
