package rest

import (
	"api/application"
	"api/infrastructure/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) (err error) {
	users := application.NewUserService(
		repository.NewUserRepository(),
	).GetUserListByRoomId(c.Param("room_id"))

	return c.JSON(http.StatusOK, users)
}
