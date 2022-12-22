package app

import (
	"fmt"

	"github.com/aland20/go-noting/routes"
	"github.com/labstack/echo/v4"
)

func AppInit() error {

	e := echo.New()

	groups := e.Group("/api")

	routes.BindUserApi(groups)

	fmt.Println("🚀 Serving on http://127.0.0.1:8000")

	return e.Start(":8000")
}
