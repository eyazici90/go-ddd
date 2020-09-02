package main

import (
	"orderContext/api"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	api.RegisterHandlers(e)

	e.Logger.Fatal(e.Start(":8080"))
}
