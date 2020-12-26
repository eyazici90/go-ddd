package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrInvalidRequest = errors.New("invalid Request params")

func update(c echo.Context, fn func(ctx context.Context, identifier string) error) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, ErrInvalidRequest)
	}
	err := fn(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusAccepted, "")
}

func create(c echo.Context, fn func(ctx context.Context) error) error {
	if err := fn(c.Request().Context()); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

func get(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusOK, result)
}

func getByID(c echo.Context, fn func(ctx context.Context, id string) interface{}) error {
	id := c.Param("id")
	result := fn(c.Request().Context(), id)
	return c.JSON(http.StatusOK, result)
}
