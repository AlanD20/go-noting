package cmd

import (
	"github.com/aland20/go-noting/app"
	"github.com/spf13/cobra"
)

func NewServeCommand() *cobra.Command {

	serveCommand := &cobra.Command{
		Use:   "serve",
		Short: "Serve the web application",
		Run: func(cmd *cobra.Command, args []string) {

			err := app.AppInit()

			if err != nil {
				panic("Failed to start server!")
			}
		},
	}

	return serveCommand
}
