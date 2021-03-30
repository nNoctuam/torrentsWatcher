package watcher

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/notification"
	"torrentsWatcher/internal/api/tracking"
	"torrentsWatcher/internal/storage"
)

type Watcher struct {
	ctx             context.Context
	wg              *sync.WaitGroup
	interval        time.Duration
	trackers        tracking.Trackers
	notificator     notification.Notificator
	torrentsStorage storage.Torrents
}

func New(
	ctx context.Context,
	wg *sync.WaitGroup,
	interval time.Duration,
	trackers tracking.Trackers,
	notificator notification.Notificator,
	torrentsStorage storage.Torrents,
) *Watcher {
	return &Watcher{
		ctx:             ctx,
		wg:              wg,
		interval:        interval,
		trackers:        trackers,
		notificator:     notificator,
		torrentsStorage: torrentsStorage,
	}
}

func (w *Watcher) Run() {
	fmt.Printf("Start checking every %s\n", w.interval)
	ticker := time.After(0)
	for {
		select {
		case <-w.ctx.Done():
			w.wg.Done()
			return
		case <-ticker:
			var torrents []models.Torrent
			err := w.torrentsStorage.Find(&torrents, "")
			if err != nil {
				log.Print("Couldn't get torrents for check")
			}
			for _, torrent := range torrents {
				w.checkTorrent(&torrent)
			}
			ticker = time.After(w.interval)
		}
	}
}

func (w *Watcher) checkTorrent(torrent *models.Torrent) {
	updatedTorrent, err := w.trackers.GetTorrentInfo(torrent.PageUrl)

	if err != nil {
		log.Print("Error parsing torrent: ", err)
		return
	}

	isUpdated := torrent.UploadedAt.Unix() != updatedTorrent.UploadedAt.Unix()

	if isUpdated || torrent.FileUrl != "" && torrent.File == nil {
		_, file, err := w.trackers.DownloadTorrentFile(torrent)
		if err != nil {
			fmt.Printf("Failed to load torrent file '%s': %v", torrent.FileUrl, err)
			return
		}
		torrent.File = file
	}

	torrent.UpdateFrom(updatedTorrent)
	err = w.torrentsStorage.Save(torrent)
	if err != nil {
		log.Printf("Couldn't save torrent: %v", updatedTorrent)
		return
	}

	if isUpdated {
		log.Printf("torrent '%s' (%s) was updated!", torrent.Title, torrent.PageUrl)
		notification.NotifyAbout(torrent, w.notificator)
	}
}
