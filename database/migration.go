package database

import (
	"github.com/aland20/go-noting/app/helpers"
	"github.com/aland20/go-noting/database/migrations"
)

func MigrateUp() {

	db := Connect()

	hasMigrated := false

	helpers.Info("Migrating database...")

	if !db.Migrator().HasTable(&migrations.User{}) {
		if err := db.Migrator().CreateTable(&migrations.User{}); err != nil {
			panic("Failed to migrate User schema")
		}
		hasMigrated = true
	}

	if !db.Migrator().HasTable(&migrations.Note{}) {
		if err := db.Migrator().CreateTable(&migrations.Note{}); err != nil {
			panic("Failed to migrate Note schema")
		}
		hasMigrated = true
	}

	if !db.Migrator().HasConstraint(&migrations.User{}, "fk_users_notes") {
		if err := db.Migrator().CreateConstraint(&migrations.User{}, "fk_users_notes"); err != nil {
			panic("Failed to create foreign key on Note Schema for User table")
		}
		hasMigrated = true
	}

	if hasMigrated {

		helpers.Success("Database migrated successfully")
	} else {

		helpers.Warn("There is nothing to migrate")
	}
}
