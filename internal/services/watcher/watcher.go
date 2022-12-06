package watcher

import (
	"context"
	"fmt"
	"sync"
	"time"

	"torrentsWatcher/internal/models"
	"torrentsWatcher/internal/ports"
	"torrentsWatcher/internal/services/tracking"
	"torrentsWatcher/internal/storage"

	"go.uber.org/zap"
)

type Watcher struct {
	ctx             context.Context
	wg              *sync.WaitGroup
	logger          *zap.Logger
	interval        time.Duration
	trackers        tracking.Trackers
	torrentClient   ports.TorrentClient
	torrentsStorage storage.Torrents
}

func New(
	ctx context.Context,
	wg *sync.WaitGroup,
	logger *zap.Logger,
	interval time.Duration,
	trackers tracking.Trackers,
	torrentClient ports.TorrentClient,
	torrentsStorage storage.Torrents,
) *Watcher {
	return &Watcher{
		ctx:             ctx,
		wg:              wg,
		logger:          logger,
		interval:        interval,
		trackers:        trackers,
		torrentClient:   torrentClient,
		torrentsStorage: torrentsStorage,
	}
}

func (w *Watcher) Run() {
	w.logger.Info(fmt.Sprintf("Start checking every %s\n", w.interval))
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
			for i := range torrents {
				w.checkTorrent(&torrents[i])
			}
			ticker = time.After(w.interval)
		}
	}
}

func (w *Watcher) checkTorrent(torrent *models.Torrent) {
	updatedTorrent, err := w.trackers.GetTorrentInfo(torrent.PageURL)
	if err != nil {
		w.logger.Error("Failed to parse torrent", zap.Error(err))
		return
	}

	isUpdated := torrent.UploadedAt.Unix() != updatedTorrent.UploadedAt.Unix()

	if isUpdated || torrent.FileURL != "" && torrent.File == nil {
		_, file, err := w.trackers.DownloadTorrentFile(torrent)
		if err != nil {
			w.logger.Error("Failed to load torrent file", zap.Error(err), zap.String("url", torrent.FileURL))
			return
		}
		torrent.File = file
	}

	torrent.UpdateFrom(updatedTorrent)

	if isUpdated {
		w.logger.Info("torrent was updated", zap.String("title", torrent.Title), zap.String("url", torrent.PageURL))
		err = w.torrentClient.UpdateTorrent(torrent.PageURL, torrent.File)
		if err != nil {
			w.logger.Error(
				"torrent replace",
				zap.String("title", torrent.Title),
				zap.String("url", torrent.PageURL),
				zap.Error(err),
			)
		}
	}

	err = w.torrentsStorage.Save(torrent)
	if err != nil {
		w.logger.Error("Failed to save torrent to storage", zap.Error(err), zap.Any("torrent", updatedTorrent))
		return
	}
}
