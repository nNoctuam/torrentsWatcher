package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB(fileName string) {
	var err error
	DB, err = gorm.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	DB.Close()
}
