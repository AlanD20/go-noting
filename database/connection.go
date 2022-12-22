package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	DB_URL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(sqlite.Open(DB_URL), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
