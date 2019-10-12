package codes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"gitlab.com/cikadev/ketide/handler"
	"gitlab.com/cikadev/ketide/repository/codes"
)

func Get(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		return c.String(http.StatusBadRequest, "Err")
	}

	code := codes.Codes{
		ID: id,
	}

	foundCode, err := code.FindByID()

	if foundCode == nil || err != nil {
		return c.Render(http.StatusNotFound, "notfound.html", map[string]interface{}{})
	}

	data := map[string]interface{}{
		"code":  foundCode,
		"owner": "guest",
		"title": id,
	}

	return c.Render(http.StatusOK, "single-code.html", handler.AppendSessionData(c, data))
}
