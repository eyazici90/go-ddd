package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	orderBaseURL string = "/orders"
	version      string = "v1"
)

func (s *Server) useRoutes() {
	v1 := s.echo.Group("/api/" + version)

	v1.GET(orderBaseURL, s.queryCtrl.getOrders)
	v1.GET(orderBaseURL+"/:id", s.queryCtrl.getOrder)

	v1.POST(orderBaseURL, s.cmdCtrl.create)

	v1.PUT(orderBaseURL+"/pay"+"/:id", s.cmdCtrl.pay)
	v1.PUT(orderBaseURL+"/ship"+"/:id", s.cmdCtrl.ship)
	v1.PUT(orderBaseURL+"/cancel"+"/:id", s.cmdCtrl.cancel)
}

func (s *Server) useHealth() {
	s.echo.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy")
	})
}
