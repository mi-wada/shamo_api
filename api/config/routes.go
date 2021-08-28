package config

import (
	"api/app/handler"

	"github.com/labstack/echo/v4"
)

func SetRoutes(server *echo.Echo) {
	server.GET("/rooms/:roomId/users", handler.GetUsers)

	server.GET("/rooms/:roomId/payments", handler.GetPayments)
	server.POST("/payments", handler.CreatePayment)
	server.DELETE("/payments/:id", handler.DeletePayment)
}
