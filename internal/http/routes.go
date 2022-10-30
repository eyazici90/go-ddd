package http

import (
	"net/http"

	"github.com/eyazici90/go-ddd/internal/app/cancel"
	"github.com/eyazici90/go-ddd/internal/app/create"
	"github.com/eyazici90/go-ddd/internal/app/pay"
	"github.com/eyazici90/go-ddd/internal/app/ship"
	"github.com/labstack/echo/v4"
)

const (
	orderBaseURL string = "/orders"
	version      string = "v1"
)

func (s *Server) useRoutes() {
	v1 := s.echo.Group("/api/" + version)

	v1.GET(orderBaseURL, s.qry.GetOrders)
	v1.GET(orderBaseURL+"/:id", s.qry.GetOrder)

	v1.POST(orderBaseURL, create.CommandController(s.mediator))

	v1.PUT(orderBaseURL+"/pay"+"/:id", pay.CommandController(s.mediator))
	v1.PUT(orderBaseURL+"/ship"+"/:id", ship.CommandController(s.mediator))
	v1.PUT(orderBaseURL+"/cancel"+"/:id", cancel.CommandController(s.mediator))
}

func (s *Server) useHealth() {
	s.echo.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy")
	})
}
