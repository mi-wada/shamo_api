package rest

import (
	"api/application"
	"api/infrastructure/cookie"
	"api/infrastructure/jwt"
	"api/infrastructure/repository"
	"api/infrastructure/session_store"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) (err error) {
	claims, _ := jwt.ParseIdToken(c.QueryParam("id_token"))
	userId := claims.UserId

	user_service := application.NewUserService(repository.NewUserRepository())

	if !user_service.Exist(userId) {
		user_service.CreateUser(
			userId,
			claims.Name,
			"",
		)
	}

	sessionId := session_store.NewSessionStore().Set(userId)
	cookie.WriteCookie(c, uuid.NewString(), sessionId)

	// TODO: ログイン済みかチェックする

	return c.Redirect(http.StatusFound, "/rooms")
}