package api

import (
	"time"

	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) useMiddlewares() {
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.RequestID())
	s.useTimeout()
}

func (s *Server) useTimeout() {
	s.echo.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Second * time.Duration(s.cfg.Server.Timeout),
	}))
}
