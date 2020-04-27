package models

import (
	"github.com/jinzhu/gorm"
)

type Torrent struct {
	gorm.Model
	Title   string
	PageUrl string `gorm:"unique"`
	FileUrl string
}
