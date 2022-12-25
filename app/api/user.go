package api

import (
	"github.com/aland20/go-noting/app/model"
	"github.com/aland20/go-noting/app/response"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	App *BaseApp
}

func BindUserApi(app *BaseApp, rg *echo.Group) {

	api := &UserController{App: app}

	rg.GET("/users", api.IndexUser).Name = "user.index"
	rg.GET("/users/:username", api.ShowUser).Name = "user.show"
	rg.POST("/users", api.StoreUser).Name = "user.store"
	rg.PATCH("/users/:username", api.UpdateUser).Name = "user.update"
	rg.DELETE("/users/:username", api.DestroyUser).Name = "user.destroy"
}

func (ctr *UserController) IndexUser(c echo.Context) error {

	users := &[]model.User{}

	err := ctr.App.Connection.Select("id", "username", "email", "created_at", "updated_at").Find(&users).Error

	if err != nil {
		return response.NotFound(c)
	}

	return response.Ok(c, users)
}

func (ctr *UserController) ShowUser(c echo.Context) error {

	user := &model.User{}
	username := c.Param("username")

	err := ctr.App.Connection.Where("username = ?", username).Take(&user).Error

	if err != nil {
		return response.NotFound(c)
	}

	return response.Ok(c, user)
}

func (ctr *UserController) StoreUser(c echo.Context) error {

	u := &model.UserSchema{}

	// Bind incoming form to u (UserSchema) type
	if err := c.Bind(&u); err != nil {
		return response.BadRequest(c)
	}

	// Construct User type from UserSchema
	user := u.NewUser()

	if err := ctr.App.Connection.Create(&user).Error; err != nil {
		return response.BadRequest(c)
	}

	return response.Created(c, &u)
}

func (ctr *UserController) UpdateUser(c echo.Context) error {

	user := &model.User{}
	u := &model.UserSchema{}
	username := c.Param("username")

	// Bind incoming form to u (UserSchema) type
	if err := c.Bind(&u); err != nil {
		return response.BadRequest(c)
	}

	// Retrieve user
	err := ctr.App.Connection.Where("username = ?", username).Take(&user).Error

	if err != nil {
		return response.NotFound(c)
	}

	// Update user (User) retrieved from DB values with u (User Schema)
	u.UpdateUser(user)

	// Run update query to database
	if err := ctr.App.Connection.Updates(&user).Error; err != nil {
		return response.UnprocessableEntity(c)
	}

	return response.Update(c, user)
}

func (ctr *UserController) DestroyUser(c echo.Context) error {

	user := &model.User{}
	username := c.Param("username")

	errTake := ctr.App.Connection.Where("username = ?", username).Take(&user).Error

	if errTake != nil {
		return response.NotFound(c)
	}

	// Delete user for given username
	errDel := ctr.App.Connection.Where("username = ?", username).Delete(&user).Error

	if errDel != nil {
		return response.UnprocessableEntity(c)
	}

	return response.Delete(c)
}
