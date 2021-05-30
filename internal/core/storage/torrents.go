package storage

import (
	"torrentsWatcher/internal/core/models"
)

type Torrents interface {
	Save(torrent *models.Torrent) error
	Find(torrents *[]models.Torrent, query interface{}, args ...interface{}) error
	First(torrent *models.Torrent, where ...interface{}) error
	SaveTransmission(torrent *models.TransmissionTorrent) error
	GetAllTransmission(torrents *[]models.TransmissionTorrent) error
}
