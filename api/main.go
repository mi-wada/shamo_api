package main

import (
	"api/config"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	config.SettingForServer(server)
	config.SetRoutes(server)

	server.Logger.Fatal(server.Start(":8080"))
}
