package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RecentCodesHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "code.html", map[string]interface{}{})
}
