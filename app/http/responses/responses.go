package responses

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string
	Code    uint16
	Data    interface{}
}

func Created(c echo.Context, dt interface{}) error {

	return c.JSON(http.StatusCreated, Response{
		Message: "Created Successfully",
		Code:    http.StatusCreated,
		Data:    dt,
	})
}

func BadRequest(c echo.Context) error {

	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": "Bad request",
		"code":    fmt.Sprint(http.StatusBadRequest),
	})
}
