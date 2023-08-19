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

func (ctr *UserController) IndexUser(ctx echo.Context) error {

	users := &[]model.User{}

	err := ctr.App.Connection.Select("id", "username", "email", "created_at", "updated_at").
		Find(&users).
		Error

	if err != nil {
		return response.NotFound(ctx)
	}

	return response.Ok(ctx, users)
}

func (ctr *UserController) ShowUser(ctx echo.Context) error {

	user := &model.User{}
	username := ctx.Param("username")

	err := ctr.App.Connection.Where("username = ?", username).Take(&user).Error

	if err != nil {
		return response.NotFound(ctx)
	}

	return response.Ok(ctx, user)
}

func (ctr *UserController) StoreUser(ctx echo.Context) error {

	u := &model.UserSchema{}

	// Bind incoming form to u (UserSchema) type
	if err := ctx.Bind(&u); err != nil {
		return response.BadRequest(ctx)
	}

	// Construct User type from UserSchema
	user := u.NewUser()

	if err := ctr.App.Connection.Create(&user).Error; err != nil {
		return response.BadRequest(ctx)
	}

	return response.Created(ctx, &u)
}

func (ctr *UserController) UpdateUser(ctx echo.Context) error {

	user := &model.User{}
	u := &model.UserSchema{}
	username := ctx.Param("username")

	// Bind incoming form to u (UserSchema) type
	if err := ctx.Bind(&u); err != nil {
		return response.BadRequest(ctx)
	}

	// Retrieve user
	err := ctr.App.Connection.Where("username = ?", username).Take(&user).Error

	if err != nil {
		return response.NotFound(ctx)
	}

	// Update user (User) retrieved from DB values with u (User Schema)
	u.UpdateUser(user)

	// Run update query to database
	if err := ctr.App.Connection.Updates(&user).Error; err != nil {
		return response.UnprocessableEntity(ctx)
	}

	return response.Update(ctx, user)
}

func (ctr *UserController) DestroyUser(ctx echo.Context) error {

	user := &model.User{}
	username := ctx.Param("username")

	errTake := ctr.App.Connection.Where("username = ?", username).Take(&user).Error

	if errTake != nil {
		return response.NotFound(ctx)
	}

	// Delete user for given username
	errDel := ctr.App.Connection.Where("username = ?", username).Delete(&user).Error

	if errDel != nil {
		return response.UnprocessableEntity(ctx)
	}

	return response.Delete(ctx)
}
