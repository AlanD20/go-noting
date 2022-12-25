package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json:"message"`
	Code    uint16      `json:"code"`
	Data    interface{} `json:"data"`
}

// Mostly used responses

func Ok(c echo.Context, dt interface{}) error {

	return c.JSON(http.StatusOK, Response{
		Message: "Data retrieved",
		Code:    http.StatusOK,
		Data:    dt,
	})
}

func Update(c echo.Context, dt interface{}) error {

	return c.JSON(http.StatusOK, Response{
		Message: "Record Updated successfully",
		Code:    http.StatusOK,
		Data:    dt,
	})
}

func Delete(c echo.Context) error {

	return c.JSON(http.StatusOK, Response{
		Message: "Record deleted successfully",
		Code:    http.StatusOK,
	})
}

func Created(c echo.Context, dt interface{}) error {

	return c.JSON(http.StatusCreated, Response{
		Message: "Created Successfully",
		Code:    http.StatusCreated,
		Data:    dt,
	})
}

func BadRequest(c echo.Context) error {

	return c.JSON(http.StatusBadRequest, Response{
		Message: "Bad request",
		Code:    http.StatusBadRequest,
	})
}

func NotFound(c echo.Context) error {

	return c.JSON(http.StatusNotFound, Response{
		Message: "Not Found",
		Code:    http.StatusNotFound,
	})
}

func Unauthorized(c echo.Context) error {

	return c.JSON(http.StatusUnauthorized, Response{
		Message: "Unauthorized",
		Code:    http.StatusUnauthorized,
	})
}

func UnprocessableEntity(c echo.Context) error {

	return c.JSON(http.StatusUnprocessableEntity, Response{
		Message: "Unprocessable Entity",
		Code:    http.StatusUnprocessableEntity,
	})
}
