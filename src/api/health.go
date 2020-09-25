package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Health() func(echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy")
	}
}
