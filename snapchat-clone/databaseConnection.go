package snapchat_clone

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func DBConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panicln("Error connecting to database")
	}
	return db
}
