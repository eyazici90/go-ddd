package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var InvalidRequestError = errors.New("Invalid Request params")

func update(c echo.Context, action func(identifier string)) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, InvalidRequestError)
	}

	action(id)

	return c.JSON(http.StatusAccepted, "")
}

func updateErr(c echo.Context, action func(identifier string) error) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, InvalidRequestError)
	}
	result := action(id)
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusAccepted, "")
}

func create(c echo.Context, action func()) error {
	action()

	return c.JSON(http.StatusCreated, "")
}
