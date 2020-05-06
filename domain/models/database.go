package models

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitDatabase function initialize database
func InitDatabase() (db *gorm.DB, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	var databasePath string = wd + "/messenger.db"
	db, err = gorm.Open("sqlite3", databasePath)

	// Apply Database migrations
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Message{})
	db.AutoMigrate(&Accesstoken{})

	db.LogMode(true)

	return
}
