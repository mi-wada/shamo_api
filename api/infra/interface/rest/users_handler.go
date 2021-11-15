package rest

import (
	"api/application"
	"api/infra/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) (err error) {
	user_service := application.NewUserService(repository.NewUserRepository())

	users := user_service.GetUserListByRoomId(c.Param("room_id"))

	return c.JSON(http.StatusOK, users)
}
