package main

import (
	"healthCheck/app"
	"healthCheck/config"
	"healthCheck/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	e := echo.New()
	routes.Register(e)
	app.StartServer(e)
}
