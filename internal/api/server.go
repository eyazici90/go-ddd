package api

import (
	"github.com/labstack/echo/v4"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
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
	server.routes()

	return server
}

func (s *Server) Start() error {
	port := s.cfg.Server.Port
	return s.echo.Start(port)
}
