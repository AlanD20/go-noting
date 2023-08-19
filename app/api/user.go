package api

import (
	"strings"

	"github.com/aland20/go-noting/app/response"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	App *BaseApp
}

func BindUserApi(app *BaseApp, rg *echo.Group) {

	ctrl := &UserController{App: app}

	rg.GET("/users", ctrl.IndexUser).Name = "user.index"
	rg.GET("/users/:username", ctrl.ShowUser).Name = "user.show"
	rg.POST("/users", ctrl.StoreUser).Name = "user.store"
	rg.PATCH("/users/:username", ctrl.UpdateUser).Name = "user.update"
	rg.DELETE("/users/:username", ctrl.DestroyUser).Name = "user.destroy"
}

func (ctrl *UserController) IndexUser(ctx echo.Context) error {

	users := &[]User{}

	err := ctrl.App.Connection.
		Model(&User{}).
		Select("id", "username", "email", "created_at", "updated_at").
		Where("deleted_at IS NULL").
		Find(&users).
		Error

	if err != nil {
		return response.NotFound(ctx)
	}

	return response.Ok(ctx, users)
}

func (ctrl *UserController) ShowUser(ctx echo.Context) error {

	user := &User{}
	username := ctx.Param("username")

	err := ctrl.App.Connection.Where("username = ?", username).Take(&user).Error

	if err != nil {
		return response.NotFound(ctx)
	}

	return response.Ok(ctx, user)
}

func (ctrl *UserController) StoreUser(ctx echo.Context) error {

	u := &UserSchema{}

	// Bind incoming form to u (UserSchema) type
	if err := ctx.Bind(&u); err != nil {
		return response.BadRequest(ctx)
	}

	// Create user
	user := &User{
		UserSchema: UserSchema{
			Username: strings.ToLower(u.Username),
			Email:    strings.ToLower(u.Email),
			Password: u.Password,
		},
	}

	// Generate UUID
	user.GenerateId()

	if err := ctrl.App.Connection.Create(&user).Error; err != nil {
		return response.BadRequest(ctx)
	}

	return response.Created(ctx, &u)
}

func (ctrl *UserController) UpdateUser(ctx echo.Context) error {

	u := &UserSchema{}
	user := &User{}
	username := ctx.Param("username")

	// Bind incoming form to u (UserSchema) type
	if err := ctx.Bind(&u); err != nil {
		return response.BadRequest(ctx)
	}

	// Retrieve user
	err := ctrl.App.Connection.
		Model(&User{}).
		Select("id", "username", "email", "created_at", "updated_at").
		Where("username = ? AND deleted_at IS NULL", username).
		Take(&user).
		Error

	if err != nil {
		return response.NotFound(ctx)
	}

	// Update user
	user.Username = strings.ToLower(u.Username)
	user.Email = strings.ToLower(u.Email)
	user.Password = u.Password

	// Run update query to database
	if err := ctrl.App.Connection.Updates(&user).Error; err != nil {
		return response.UnprocessableEntity(ctx)
	}

	// Unset sensitive fields
	user.Password = ""

	return response.Update(ctx, user)
}

func (ctrl *UserController) DestroyUser(ctx echo.Context) error {

	user := &User{}
	username := ctx.Param("username")

	errTake := ctrl.App.Connection.Where("username = ?", username).Take(&user).Error

	if errTake != nil {
		return response.NotFound(ctx)
	}

	// Delete user for given username
	errDel := ctrl.App.Connection.Where("username = ?", username).Delete(&user).Error

	if errDel != nil {
		return response.UnprocessableEntity(ctx)
	}

	return response.Delete(ctx)
}
