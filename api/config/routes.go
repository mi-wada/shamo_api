package config

import (
	"api/app/handler"

	"github.com/labstack/echo/v4"
)

func SetRoutes(server *echo.Echo) {
	server.GET("/room/:roomId/users", handler.GetUsers)

	server.GET("/room/:roomId/payments", handler.GetPayments)
	server.POST("/payment", handler.CreatePayment)
	server.DELETE("/payment/:id", handler.DeletePayment)
}
