package storage

import (
	models2 "torrentsWatcher/internal/core/models"
)

type Cookies interface {
	Save(authCookie *models2.AuthCookie) error
	Find(authCookies *[]models2.AuthCookie, query interface{}, args ...interface{}) error
	First(authCookie *models2.AuthCookie, where ...interface{}) error
	Delete(where ...interface{}) error
	DeleteByDomain(domain string) error
}
