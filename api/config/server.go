package config

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SettingForServer(server *echo.Echo) {
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlRequestHeaders,
			echo.HeaderAccessControlRequestMethod,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAccessControlAllowOrigin,
		},
	}))
	server.Use(middleware.Recover())
	server.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Fprintf(os.Stderr, "Request: %v\n", string(reqBody))
		fmt.Fprintf(os.Stderr, "Headers: %v\n", c.Request().Header)
		fmt.Fprintf(os.Stderr, "Method: %v\n", c.Request().Method)
	}))
}
