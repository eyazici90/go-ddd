package create

import (
	"context"
	"net/http"

	"github.com/eyazici90/go-ddd/internal/app/feature"
	"github.com/eyazici90/go-mediator/mediator"
	"github.com/labstack/echo/v4"
)

func CommandController(med *mediator.Mediator) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")

		return feature.Handle(c, http.StatusCreated, func(ctx context.Context) error {
			return med.Send(ctx, OrderCommand{OrderID: id})
		})
	}
}
