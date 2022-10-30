package http

import (
	"context"

	"github.com/eyazici90/go-ddd/internal/app/query"
	"github.com/eyazici90/go-mediator/mediator"
	"github.com/labstack/echo/v4"
)

type Config struct {
	Server struct {
		Port    string `json:"port"`
		Timeout int    `json:"timeout"`
	} `json:"server"`
	MongoDB struct {
		URL      string `json:"url"`
		Database string `json:"database"`
	} `json:"mongoDb"`
	Context struct {
		Timeout int `json:"timeout"`
	} `json:"context"`
}

type Server struct {
	cfg      Config
	echo     *echo.Echo
	qry      query.OrderController
	mediator *mediator.Mediator
}

func NewServer(cfg Config,
	e *echo.Echo,
	qry query.OrderController,
	med *mediator.Mediator,
) *Server {
	server := Server{
		cfg:      cfg,
		echo:     e,
		mediator: med,
	}

	server.useHealth()
	server.useRoutes()
	server.useMiddlewares()

	return &server
}

func (s *Server) Start() error {
	port := s.cfg.Server.Port
	return s.echo.Start(port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *Server) Fatal(err error) { s.echo.Logger.Fatal(err) }

func (s *Server) Config() Config { return s.cfg }
