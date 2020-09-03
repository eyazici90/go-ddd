package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var InvalidRequestError = errors.New("Invalid Request params")

func handle(c echo.Context, action func(identifier string)) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, InvalidRequestError)
	}

	action(id)

	return c.JSON(http.StatusAccepted, "")
}
