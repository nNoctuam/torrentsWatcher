package api

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "./torrents.db")
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	DB.Close()
}
