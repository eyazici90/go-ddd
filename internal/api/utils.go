package api

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
)

var ErrInvalidRequest = errors.New("invalid Request params")

func handle(c echo.Context,
	statusCode int,
	fn func(ctx context.Context) error) error {
	if err := fn(c.Request().Context()); err != nil {
		return err
	}
	return c.JSON(statusCode, "")
}

func handleR(c echo.Context,
	statusCode int,
	fn func(ctx context.Context) (interface{}, error)) error {
	result, err := fn(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(statusCode, result)
}
