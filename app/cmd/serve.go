package cmd

import (
	"github.com/aland20/go-noting/app/api"
	logger "github.com/aland20/go-noting/app/loggers"
	"github.com/spf13/cobra"
)

func NewServeCommand() *cobra.Command {

	serveCommand := &cobra.Command{
		Use:   "serve",
		Short: "Serve the web application",
		Run: func(cmd *cobra.Command, args []string) {

			if err := api.NewBaseApp(); err != nil {
				logger.Panic("Failed to start the application!")
			}

		},
	}

	return serveCommand
}
