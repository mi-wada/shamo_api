package handler

import (
	"api/app/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateRoom(c echo.Context) (err error) {
	room := service.CreateRoom()

	return c.JSON(http.StatusOK, room)
}
