package cmd

import (
	"github.com/aland20/go-noting/app/helpers"
	"github.com/aland20/go-noting/database/runner"
	"github.com/spf13/cobra"
)

func NewMigrateCommand() *cobra.Command {

	var fresh bool

	migrateCommand := &cobra.Command{
		Use:              "migrate",
		Short:            "Migrate database",
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {

			if fresh {
				helpers.Warn("All the existing data will be erased")
				isDropped := make(chan bool)
				go func() {
					runner.DropTables()
					runner.CreateTables()
					isDropped <- true
				}()

				<-isDropped
			}

			runner.AutoMigrate()
		},
	}

	migrateCommand.Flags().BoolVarP(&fresh, "fresh", "f", false, "Drops all the existing table, then runs migration.")

	return migrateCommand
}
