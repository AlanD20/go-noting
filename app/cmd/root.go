package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-noting",
	Short: "Go-Noting is a simple, easy-to-use, noting web application.",
	Long:  `Learn more at https://github.com/AlanD20/go-noting`,
}

func Execute() {

	rootCmd.AddCommand(NewMigrateCommand())
	rootCmd.AddCommand(NewServeCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
