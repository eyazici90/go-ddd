package http

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"

	"github.com/eyazici90/go-ddd/internal/domain"
	"github.com/eyazici90/go-ddd/pkg/aggregate"
	"github.com/eyazici90/go-ddd/pkg/httperr"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) useMiddlewares() {
	s.useLogger()
	s.useRecover()
	s.useRequestID()
	s.useTimeout()
	s.useMetrics()

	s.useErrorHandler(httperr.NewHandler(
		httperr.DefaultHandler.WithMap(http.StatusBadRequest,
			aggregate.ErrNotFound,
			domain.ErrNotPaid,
			domain.ErrInvalidValue,
		),
		httperr.DefaultHandler.WithMapFunc(
			func(err error) (int, bool) {
				_, ok := err.(validator.ValidationErrors)
				return http.StatusBadRequest, ok
			}),
	))
}

func (s *Server) useTimeout() {
	s.echo.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Second * time.Duration(s.cfg.Server.Timeout),
	}))
}

func (s *Server) useLogger() {
	s.echo.Use(middleware.Logger())
}

func (s *Server) useRecover() {
	s.echo.Use(middleware.Recover())
}

func (s *Server) useRequestID() {
	s.echo.Use(middleware.RequestID())
}

func (s *Server) useErrorHandler(httpErrHandler *httperr.Handler) {
	s.echo.HTTPErrorHandler = httpErrHandler.Handle()
}

func (s *Server) useMetrics() {
	p := prometheus.NewPrometheus("echo", urlSkipper)
	p.Use(s.echo)
}

func urlSkipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/health")
}
