package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	logger "github.com/aland20/go-noting/app/loggers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {

	var DB_URL = os.Getenv("DATABASE_URL")

	logger.Info("Connection is open to database")

	db, err := gorm.Open(sqlite.Open(DB_URL), &gorm.Config{})

	if err != nil {
		logger.Panic("failed to connect database")
	}

	return db
}

func StartTemporaryConnection(fn func(*gorm.DB, *sql.DB) error) error {

	var wg sync.WaitGroup
	conn := NewConnection()
	db, err := conn.DB()

	if err != nil {
		logger.Error("Failed to retreive SQL instance!")
		return err
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := fn(conn, db); err != nil {
			fmt.Println(err)
			logger.Panic("Something went wrong during temporary connection!")
		}

	}()

	wg.Wait()

	return nil
}
