package routes

import (
	"github.com/aland20/go-noting/app/http/controllers"
	"github.com/labstack/echo/v4"
)

func BindUserApi(rg *echo.Group) {

	gr := rg.Group("/users")

	// gr.GET("/:username", controllers.UserCreate)
	gr.POST("/", controllers.UserCreate).Name = "user.register"
	// gr.PUT("/:username", controllers.UserCreate)
	// gr.DELETE("/:username", controllers.UserCreate)
}
