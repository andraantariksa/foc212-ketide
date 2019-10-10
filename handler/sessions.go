package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func IsLoggedIn(c echo.Context) bool {
	sess, err := session.Get("sessions", c)
	if err != nil || sess.Values["username"] == nil {
		return false
	}
	return true
}

func RequiresLogin(handler func(c echo.Context) error) func(echo.Context) error {
	return func(c echo.Context) error {
		if !IsLoggedIn(c) {
			return c.Redirect(http.StatusTemporaryRedirect, "/signin")
		}

		return handler(c)
	}
}
