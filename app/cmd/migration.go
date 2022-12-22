package cmd

import (
	"github.com/aland20/go-noting/database"
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		database.MigrateUp()
	},
}

func init() {

	rootCmd.AddCommand(migrationCmd)
}
