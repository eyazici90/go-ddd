package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var InvalidRequestError = errors.New("Invalid Request params")

func update(c echo.Context, action func(ctx context.Context, identifier string) error) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, InvalidRequestError)
	}
	err := action(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusAccepted, "")
}

func create(c echo.Context, action func(ctx context.Context) error) error {
	if err := action(c.Request().Context()); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

func get(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusOK, result)
}

func getById(c echo.Context, action func(ctx context.Context, id string) interface{}) error {
	id := c.Param("id")
	result := action(c.Request().Context(), id)
	return c.JSON(http.StatusOK, result)
}
