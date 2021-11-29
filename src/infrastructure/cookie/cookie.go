package cookie

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const DEFAULT_COOKIE_KEY = "shamo_session"

func WriteCookie(c echo.Context, key, value string) {
	if key == "" {
		key = DEFAULT_COOKIE_KEY
	}

	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}

func ReadCookie(c echo.Context, key string) (string, error) {
	if key == "" {
		key = DEFAULT_COOKIE_KEY
	}

	cookie, err := c.Cookie(key)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}
