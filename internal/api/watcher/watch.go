package watcher

import (
	"context"
	"fmt"
	"sync"
	"time"
	"torrentsWatcher/internal/api/notification"
	"torrentsWatcher/internal/api/torrentclient"

	"go.uber.org/zap"

	"torrentsWatcher/internal/api/models"
	"torrentsWatcher/internal/api/tracking"
	"torrentsWatcher/internal/storage"
)

type Watcher struct {
	ctx             context.Context
	wg              *sync.WaitGroup
	logger          *zap.Logger
	interval        time.Duration
	trackers        tracking.Trackers
	notificator     notification.Notificator
	torrentClient   *torrentclient.TorrentClient
	torrentsStorage storage.Torrents
}

func New(
	ctx context.Context,
	wg *sync.WaitGroup,
	logger *zap.Logger,
	interval time.Duration,
	trackers tracking.Trackers,
	notificator notification.Notificator,
	torrentClient *torrentclient.TorrentClient,
	torrentsStorage storage.Torrents,
) *Watcher {
	return &Watcher{
		ctx:             ctx,
		wg:              wg,
		logger:          logger,
		interval:        interval,
		trackers:        trackers,
		notificator:     notificator,
		torrentClient:   torrentClient,
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
				w.logger.Error("Couldn't get torrents for check", zap.Error(err))
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
		w.logger.Error("Failed to parse torrent", zap.Error(err))
		return
	}

	isUpdated := torrent.UploadedAt.Unix() != updatedTorrent.UploadedAt.Unix()

	if isUpdated || torrent.FileUrl != "" && torrent.File == nil {
		_, file, err := w.trackers.DownloadTorrentFile(torrent)
		if err != nil {
			w.logger.Error("Failed to load torrent file", zap.Error(err), zap.String("url", torrent.FileUrl))
			return
		}
		torrent.File = file
	}

	torrent.UpdateFrom(updatedTorrent)
	err = w.torrentsStorage.Save(torrent)
	if err != nil {
		w.logger.Error("Failed to save torrent to storage", zap.Error(err), zap.Any("torrent", updatedTorrent))
		return
	}

	if isUpdated {
		w.logger.Info("torrent was updated", zap.String("title", torrent.Title), zap.String("url", torrent.PageUrl))
		err = w.torrentClient.UpdateTorrent(torrent.PageUrl, torrent.File)
		if err != nil {
			w.logger.Error("torrent replace", zap.String("title", torrent.Title), zap.String("url", torrent.PageUrl), zap.Error(err))
		}
	}
}
