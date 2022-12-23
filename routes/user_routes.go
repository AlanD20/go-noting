package routes

import (
	"github.com/aland20/go-noting/app/http/controllers"
	"github.com/labstack/echo/v4"
)

func BindUserApi(rg *echo.Group) {

	rg.GET("/users", controllers.AllUser).Name = "user.index"
	rg.GET("/users/:username", controllers.ShowUser).Name = "user.show"
	rg.POST("/users", controllers.StoreUser).Name = "user.store"
	rg.PATCH("/users/:username", controllers.UpdateUser).Name = "user.update"
	rg.DELETE("/users/:username", controllers.DestroyUser).Name = "user.destroy"
}
