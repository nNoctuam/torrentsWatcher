package driven

import (
	"torrentsWatcher/internal/domain/models"
)

type TrackersAdapter interface {
	GetTorrentInfo(torrentURL string) (*models.Torrent, error)
	DownloadTorrentFile(torrent *models.Torrent) (string, []byte, error)
	SearchEverywhere(text string) (torrents []*models.Torrent)
}
