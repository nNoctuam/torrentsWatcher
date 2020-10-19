package impl

import (
	"github.com/jinzhu/gorm"

	"torrentsWatcher/internal/api/models"
)

type TorrentsSqliteStorage struct {
	db *gorm.DB
}

func (t TorrentsSqliteStorage) Save(torrent *models.Torrent) error {
	return t.db.Save(&torrent).Error
}

func (t TorrentsSqliteStorage) Find(torrents *[]models.Torrent, query interface{}, args ...interface{}) error {
	return t.db.Where(query, args).Find(torrents).Error
}

func (t TorrentsSqliteStorage) First(torrent *models.Torrent, query ...interface{}) error {
	return t.db.Where(query).First(torrent).Error
}

func NewTorrentsSqliteStorage(db *gorm.DB) *TorrentsSqliteStorage {
	return &TorrentsSqliteStorage{db: db}
}
