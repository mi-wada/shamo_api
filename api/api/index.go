package handler

import (
	"api/config"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := echo.New()

	config.SettingForServer(server)
	config.SetRoutes(server)

	os.Setenv("AWS_ACCESS_KEY_ID", os.Getenv("MY_AWS_ACCESS_KEY_ID"))
	os.Setenv("AWS_SECRET_ACCESS_KEY", os.Getenv("MY_AWS_SECRET_ACCESS_KEY"))
	os.Setenv("AWS_DEFAULT_REGION", os.Getenv("MY_AWS_DEFAULT_REGION"))

	server.ServeHTTP(w, r)
}
