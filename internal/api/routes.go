package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const orderBaseURL string = "/orders"
const version string = "v1"

func (s *Server) setRoutes() {

	v1 := s.echo.Group("/api/" + version)
	{
		v1.GET(orderBaseURL, s.orderQueryController.getOrders)
		v1.GET(orderBaseURL+"/:id", s.orderQueryController.getOrder)

		v1.POST(orderBaseURL, s.orderCommandController.create)

		v1.PUT(orderBaseURL+"/pay"+"/:id", s.orderCommandController.pay)
		v1.PUT(orderBaseURL+"/ship"+"/:id", s.orderCommandController.ship)
	}
}

func (s *Server) health() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy")
	})
}
