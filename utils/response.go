package utils

import (
	"errors"

	"github.com/IbadT/go_server/types"
	"github.com/labstack/echo/v4"
)

func SuccessResponse(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, data)
}

func ErrorResponse(c echo.Context, status int, message string) error {
	return c.JSON(status, types.Response{Status: status, Message: message})
}

func NewError(message string) error {
	return errors.New(message)
}
