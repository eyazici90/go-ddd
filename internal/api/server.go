package api

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	Server struct {
		Port    string `json:"port"`
		Timeout int    `json:"timeout"`
	} `json:"server"`
	MongoDb struct {
		URL      string `json:"url"`
		Database string `json:"database"`
	} `json:"mongoDb"`
	Context struct {
		Timeout int `json:"timeout"`
	} `json:"context"`
}

type Server struct {
	cfg                    Config
	echo                   *echo.Echo
	orderCommandController *OrderCommandController
	orderQueryController   *OrderQueryController
}

func NewServer(cfg Config,
	echo *echo.Echo,
	cmdController *OrderCommandController,
	querycontroller *OrderQueryController) *Server {
	server := &Server{
		cfg:                    cfg,
		echo:                   echo,
		orderCommandController: cmdController,
		orderQueryController:   querycontroller,
	}

	server.health()
	server.setRoutes()
	server.setMiddlewares()

	return server
}

func (s *Server) setMiddlewares() {
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.RequestID())
	s.setTimeout()
}

func (s *Server) setTimeout() {
	s.echo.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Second * time.Duration(s.cfg.Server.Timeout),
	}))
}

func (s *Server) Start() error {
	port := s.cfg.Server.Port
	return s.echo.Start(port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *Server) Fatal(err error) { s.echo.Logger.Fatal(err) }
