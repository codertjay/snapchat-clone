package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func DBConnection() *gorm.DB {
	connection, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panicln("Error connecting to database")
	}
	return connection
}

func CloseDB() {
	db := DBConnection()
	close, err := db.DB()
	if err != nil {
		log.Panicln("Error closing database connection ", err)
	}
	close.Close()
}
