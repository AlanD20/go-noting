package runner

import (
	"database/sql"

	logger "github.com/aland20/go-noting/app/loggers"
	"github.com/aland20/go-noting/app/model"
	"github.com/aland20/go-noting/database"
	"gorm.io/gorm"
)

func AutoMigrate() {

	database.StartTemporaryConnection(func(conn *gorm.DB, _ *sql.DB) error {

		logger.Info("Migrating database...")

		err := conn.AutoMigrate(&model.User{}, &model.Note{})

		if err != nil {
			return err
		}

		logger.Success("Database migrated successfully")

		return nil
	})
}

func CreateTables() {

	database.StartTemporaryConnection(func(conn *gorm.DB, _ *sql.DB) error {

		hasMigrated := false

		logger.Info("Creating tables...")

		if !conn.Migrator().HasTable(&model.User{}) {
			if err := conn.Migrator().CreateTable(&model.User{}); err != nil {
				logger.Panic("Failed to create `users` table")
			}
			hasMigrated = true
		}

		if !conn.Migrator().HasTable(&model.Note{}) {
			if err := conn.Migrator().CreateTable(&model.Note{}); err != nil {
				logger.Panic("Failed to create `notes` table")
			}
			hasMigrated = true
		}

		if !conn.Migrator().HasConstraint(&model.User{}, "fk_users_notes") {
			if err := conn.Migrator().CreateConstraint(&model.User{}, "fk_users_notes"); err != nil {
				logger.Panic("Failed to create foreign key on `notes` for `users` table")
			}
			hasMigrated = true
		}

		if hasMigrated {

			logger.Success("Tables created successfully")
		} else {

			logger.Warn("Tables are already exist")
		}

		return nil
	})
}

func DropTables() {

	database.StartTemporaryConnection(func(conn *gorm.DB, _ *sql.DB) error {

		logger.Info("Dropping tables...")

		if err := conn.Migrator().DropTable(&model.User{}); err != nil {
			logger.Panic("Failed to drop `user` table")
		}

		if err := conn.Migrator().DropTable(&model.User{}); err != nil {
			logger.Panic("Failed to drop `note` table")
		}

		logger.Success("Tables dropped successfully")

		return nil
	})
}
