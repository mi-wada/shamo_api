package handler

import (
	"api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	roomId := c.Param("roomId")

	users := service.GetUsersByRoomId(roomId)

	return c.JSON(http.StatusOK, users)
}
