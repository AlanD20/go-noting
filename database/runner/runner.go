package runner

import (
	logger "github.com/aland20/go-noting/app/helpers"
	"github.com/aland20/go-noting/app/models"
	"github.com/aland20/go-noting/database"
)

func AutoMigrate() {

	db := database.Connect()

	logger.Info("Migrating database...")

	db.AutoMigrate(&models.User{}, &models.Note{})

	logger.Success("Database migrated successfully")

}

func CreateTables() {

	db := database.Connect()

	hasMigrated := false

	logger.Info("Creating tables...")

	if !db.Migrator().HasTable(&models.User{}) {
		if err := db.Migrator().CreateTable(&models.User{}); err != nil {
			panic("Failed to create `users` table")
		}
		hasMigrated = true
	}

	if !db.Migrator().HasTable(&models.Note{}) {
		if err := db.Migrator().CreateTable(&models.Note{}); err != nil {
			panic("Failed to create `notes` table")
		}
		hasMigrated = true
	}

	if !db.Migrator().HasConstraint(&models.User{}, "fk_users_notes") {
		if err := db.Migrator().CreateConstraint(&models.User{}, "fk_users_notes"); err != nil {
			panic("Failed to create foreign key on `notes` for `users` table")
		}
		hasMigrated = true
	}

	if hasMigrated {

		logger.Success("Tables created successfully")
	} else {

		logger.Warn("Tables are already exist")
	}
}

func DropTables() {

	db := database.Connect()

	logger.Info("Dropping tables...")

	if err := db.Migrator().DropTable(&models.User{}); err != nil {
		panic("Failed to drop `user` table")
	}

	if err := db.Migrator().DropTable(&models.User{}); err != nil {
		panic("Failed to drop `note` table")
	}

	logger.Success("Tables dropped successfully")
}
