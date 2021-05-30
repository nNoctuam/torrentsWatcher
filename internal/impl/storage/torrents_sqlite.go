package storage

import (
	"torrentsWatcher/internal/core/models"

	"github.com/jinzhu/gorm"
)

type torrentsSqliteStorage struct {
	db *gorm.DB
}

func (t *torrentsSqliteStorage) Save(torrent *models.Torrent) error {
	return t.db.Save(&torrent).Error
}

func (t *torrentsSqliteStorage) Find(torrents *[]models.Torrent, query interface{}, args ...interface{}) error {
	return t.db.Where(query, args).Find(torrents).Error
}

func (t *torrentsSqliteStorage) First(torrent *models.Torrent, query ...interface{}) error {
	return t.db.Where(query).First(torrent).Error
}

func (t *torrentsSqliteStorage) SaveTransmission(torrent *models.TransmissionTorrent) error {
	existing := &models.TransmissionTorrent{}
	_ = t.FirstTransmission(existing, "hash = ?", torrent.Hash)
	if existing.Id != 0 {
		torrent.Id = existing.Id
	}

	return t.db.Save(&torrent).Error
}

func (t *torrentsSqliteStorage) GetAllTransmission(torrents *[]models.TransmissionTorrent) error {
	return t.db.Debug().Where("").Find(torrents).Error
}

func (t *torrentsSqliteStorage) FirstTransmission(torrent *models.TransmissionTorrent, query ...interface{}) error {
	return t.db.Where(query).First(torrent).Error
}

func NewTorrentsSqliteStorage(db *gorm.DB) *torrentsSqliteStorage {
	return &torrentsSqliteStorage{db: db}
}
