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

type App struct {
	cfg                    Config
	echo                   *echo.Echo
	orderCommandController *OrderCommandController
	orderQueryController   *OrderQueryController
}

func NewApp(cfg Config,
	echo *echo.Echo,
	cmdController *OrderCommandController,
	querycontroller *OrderQueryController) *App {
	app := &App{
		cfg:                    cfg,
		echo:                   echo,
		orderCommandController: cmdController,
		orderQueryController:   querycontroller,
	}
	app.routes()

	return app
}

func (a *App) Start() error {
	port := a.cfg.Server.Port
	return a.echo.Start(port)
}
