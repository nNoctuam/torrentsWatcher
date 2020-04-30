package models

import "github.com/jinzhu/gorm"

type AuthCookie struct {
	gorm.Model
	Domain string `gorm:"varchar(50)"`
	Name   string `gorm:"varchar(50)"`
	Value  string `gorm:"varchar(255)"`
}
