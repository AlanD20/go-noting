package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-noting",
	Short: "Go-Noting is a simple, easy-to-use, noting web application.",
	Long:  `Learn more at https://github.com/AlanD20/go-noting`,
}

func Execute() error {

	rootCmd.AddCommand(NewMigrateCommand())
	rootCmd.AddCommand(NewServeCommand())
	rootCmd.AddCommand(NewKeyCommand())

	return rootCmd.Execute()
}
