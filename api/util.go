package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func handle(c echo.Context, action func(identifier string)) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, InvalidRequestError)
	}

	action(id)

	return c.JSON(http.StatusAccepted, "")
}
