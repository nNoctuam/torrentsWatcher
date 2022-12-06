package storage

import (
	"torrentsWatcher/internal/domain/models"
)

type Cookies interface {
	Save(authCookie *models.AuthCookie) error
	Find(authCookies *[]models.AuthCookie, query interface{}, args ...interface{}) error
	First(authCookie *models.AuthCookie, where ...interface{}) error
	Delete(where ...interface{}) error
	DeleteByDomain(domain string) error
}
