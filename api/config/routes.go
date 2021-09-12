package config

import (
	"api/app/handler"

	"github.com/labstack/echo/v4"
)

func SetRoutes(server *echo.Echo) {
	server.POST("/rooms", handler.CreateRoom)

	server.GET("/rooms/:roomId/users", handler.GetUsers)
	server.GET("/rooms/:roomId/payments", handler.GetPayments)

	server.POST("/rooms/:roomId/payments", handler.CreatePayment)
	server.DELETE("/rooms/:roomId/payments", handler.DeletePayment)
}
