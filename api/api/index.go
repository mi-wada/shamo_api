package handler

import (
	"api/config"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := echo.New()

	config.SettingForServer(server)
	config.SetRoutes(server)

	fmt.Fprintf(w, r.URL.Path)
}
