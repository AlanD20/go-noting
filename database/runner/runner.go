package runner

import (
	"github.com/aland20/go-noting/app/helpers"
	"github.com/aland20/go-noting/database"
	"github.com/aland20/go-noting/database/migrations"
)

func AutoMigrate() {

	db := database.Connect()

	helpers.Info("Migrating database...")

	db.AutoMigrate(&migrations.User{}, &migrations.Note{})

	helpers.Success("Database migrated successfully")

}

func CreateTables() {

	db := database.Connect()

	hasMigrated := false

	helpers.Info("Creating tables...")

	if !db.Migrator().HasTable(&migrations.User{}) {
		if err := db.Migrator().CreateTable(&migrations.User{}); err != nil {
			panic("Failed to create `users` table")
		}
		hasMigrated = true
	}

	if !db.Migrator().HasTable(&migrations.Note{}) {
		if err := db.Migrator().CreateTable(&migrations.Note{}); err != nil {
			panic("Failed to create `notes` table")
		}
		hasMigrated = true
	}

	if !db.Migrator().HasConstraint(&migrations.User{}, "fk_users_notes") {
		if err := db.Migrator().CreateConstraint(&migrations.User{}, "fk_users_notes"); err != nil {
			panic("Failed to create foreign key on `notes` for `users` table")
		}
		hasMigrated = true
	}

	if hasMigrated {

		helpers.Success("Tables created successfully")
	} else {

		helpers.Warn("Tables are already exist")
	}
}

func DropTables() {

	db := database.Connect()

	helpers.Info("Dropping tables...")

	if err := db.Migrator().DropTable(&migrations.User{}); err != nil {
		panic("Failed to drop `user` table")
	}

	if err := db.Migrator().DropTable(&migrations.User{}); err != nil {
		panic("Failed to drop `note` table")
	}

	helpers.Success("Tables dropped successfully")
}
