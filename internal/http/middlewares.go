package http

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/eyazici90/go-ddd/internal/domain"
	"github.com/eyazici90/go-ddd/pkg/aggregate"
	"github.com/eyazici90/go-ddd/pkg/httperr"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func (s *Server) useMiddlewares() {
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Second * time.Duration(s.cfg.Server.Timeout),
	}))

	s.useInstrumentation()

	s.useErrorHandler(httperr.NewHandler(
		httperr.DefaultHandler.WithMap(http.StatusBadRequest,
			aggregate.ErrNotFound,
			domain.ErrNotPaid,
			domain.ErrInvalidValue,
		),
		httperr.DefaultHandler.WithMapFunc(
			func(err error) (int, bool) {
				var ve validator.ValidationErrors
				return http.StatusBadRequest, errors.As(err, &ve)
			}),
	))
}

func (s *Server) useErrorHandler(errHandler *httperr.Handler) {
	s.echo.HTTPErrorHandler = errHandler.Handle()
}

func (s *Server) useInstrumentation() {
	skipper := func(c echo.Context) bool {
		return strings.HasPrefix(c.Path(), "/health")
	}
	p := prometheus.NewPrometheus("echo", skipper)
	p.Use(s.echo)
	s.echo.Use(otelecho.Middleware(serverName, otelecho.WithSkipper(skipper)))
}
