package rest

import (
	"api/application"
	"api/infra/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateRoom(c echo.Context) (err error) {
	room_service := application.NewRoomService(repository.NewRoomRepository())
	room := room_service.CreateRoom()
	return c.JSON(http.StatusOK, room)
}
