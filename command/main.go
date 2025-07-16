package main

import (
	"task-manager-golang/config"
	"task-manager-golang/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()

	e := echo.New()
	routes.Routes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
