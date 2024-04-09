package http

import (
	"context"
	"errors"

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
	cfg       Config
	echo      *echo.Echo
	cmdCtrl   CommandController
	queryCtrl OrderQueryController

	shutdowns []func(ctx context.Context) error
}

func NewServer(cfg Config,
	e *echo.Echo,
	cmdCtrl CommandController,
	queryCtrl OrderQueryController,
) *Server {
	server := Server{
		cfg:       cfg,
		echo:      e,
		cmdCtrl:   cmdCtrl,
		queryCtrl: queryCtrl,
	}

	server.useProbes()
	server.useRoutes()
	server.useMiddlewares()

	return &server
}

func (s *Server) Start() error {
	s.shutdowns = append(s.shutdowns, s.echo.Shutdown)
	port := s.cfg.Server.Port
	return s.echo.Start(port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	var err error
	for _, shutdown := range s.shutdowns {
		err = errors.Join(err, shutdown(ctx))
	}
	return err
}

func (s *Server) Fatal(err error) { s.echo.Logger.Fatal(err) }

func (s *Server) Config() Config { return s.cfg }
