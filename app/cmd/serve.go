package cmd

import (
	"fmt"

	"github.com/aland20/go-noting/routes"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the web application",
	Run: func(cmd *cobra.Command, args []string) {

		e := echo.New()

		groups := e.Group("/api")

		routes.BindUserApi(groups)

		fmt.Println("🚀 Serving on http://127.0.0.1:8000")

		e.Logger.Fatal(e.Start(":8000"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
