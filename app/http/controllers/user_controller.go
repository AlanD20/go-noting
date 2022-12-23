package controllers

import (
	"github.com/aland20/go-noting/app/http/responses"
	"github.com/aland20/go-noting/app/models"
	"github.com/labstack/echo/v4"
)

func AllUser(c echo.Context) error {

	return c.String(200, "")
}

func ShowUser(c echo.Context) error {

	return c.String(200, "")
}

func StoreUser(c echo.Context) error {

	u := new(models.UserSchema)

	if err := c.Bind(u); err != nil {
		return responses.BadRequest(c)
	}

	if err := u.Create(); err != nil {
		return responses.BadRequest(c)
	}

	return responses.Created(c, u)
}

func UpdateUser(c echo.Context) error {

	return c.String(200, "")
}

func DestroyUser(c echo.Context) error {

	return c.String(200, "")
}
