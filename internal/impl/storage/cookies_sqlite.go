package storage

import (
	"torrentsWatcher/internal/core/models"

	"github.com/jinzhu/gorm"
)

type cookiesSqliteStorage struct {
	db *gorm.DB
}

func (t cookiesSqliteStorage) Save(cookie *models.AuthCookie) error {
	return t.db.Save(&cookie).Error
}

func (t cookiesSqliteStorage) Find(cookies *[]models.AuthCookie, query interface{}, args ...interface{}) error {
	return t.db.Where(query, args).Find(cookies).Error
}

func (t cookiesSqliteStorage) First(cookie *models.AuthCookie, query ...interface{}) error {
	return t.db.Where(query).First(cookie).Error
}

func (t cookiesSqliteStorage) Delete(query ...interface{}) error {
	return t.db.Where(query).Delete(&models.AuthCookie{}).Error
}

func (t cookiesSqliteStorage) DeleteByDomain(domain string) error {
	return t.db.Where("domain = ?", domain).Delete(&models.AuthCookie{}).Error
}

func NewCookiesSqliteStorage(db *gorm.DB) *cookiesSqliteStorage {
	return &cookiesSqliteStorage{db: db}
}
