package handler

import (
	"api/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	config.SettingForServer(server)
	config.SetRoutes(server)

	server.Logger.Fatal(server.Start(":8080"))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server := echo.New()

	config.SettingForServer(server)
	config.SetRoutes(server)

	server.ServeHTTP(w, r)
}
