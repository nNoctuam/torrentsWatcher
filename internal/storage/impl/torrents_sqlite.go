package impl

import (
	"github.com/jinzhu/gorm"

	"torrentsWatcher/internal/api/models"
)

type torrentsSqliteStorage struct {
	db *gorm.DB
}

func (t torrentsSqliteStorage) Save(torrent *models.Torrent) error {
	return t.db.Save(&torrent).Error
}

func (t torrentsSqliteStorage) Find(torrents *[]models.Torrent, query interface{}, args ...interface{}) error {
	return t.db.Where(query, args).Find(torrents).Error
}

func (t torrentsSqliteStorage) First(torrent *models.Torrent, query ...interface{}) error {
	return t.db.Where(query).First(torrent).Error
}

func NewTorrentsSqliteStorage(db *gorm.DB) *torrentsSqliteStorage {
	return &torrentsSqliteStorage{db: db}
}
