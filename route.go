package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"gitlab.com/cikadev/ketide/handler"
	"gitlab.com/cikadev/ketide/handler/users"
)

func route(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
	})

	e.GET("/recent", handler.RequiresLogin(users.RecentCodesHandler))
}
