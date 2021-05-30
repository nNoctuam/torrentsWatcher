package storage

import (
	models2 "torrentsWatcher/internal/core/models"
)

type Torrents interface {
	Save(torrent *models2.Torrent) error
	Find(torrents *[]models2.Torrent, query interface{}, args ...interface{}) error
	First(torrent *models2.Torrent, where ...interface{}) error
	SaveTransmission(torrent *models2.TransmissionTorrent) error
	GetAllTransmission(torrents *[]models2.TransmissionTorrent) error
}
