package codes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"gitlab.com/cikadev/ketide/handler"
	"gitlab.com/cikadev/ketide/repository/codes"
)

type codeForm struct {
	Code     string
	Language string `validate:"required"`
}

func Exec(c echo.Context) error {
	cf := new(codeForm)
	if err := c.Bind(cf); err != nil {
		return c.String(http.StatusOK, "err")
	}

	if err := c.Validate(cf); err != nil {
		return c.String(http.StatusOK, "not validated")
	}

	sessUserId := handler.GetSessionData(c)["id"]
	var userId uint64 = 0
	if sessUserId != nil {
		userId = sessUserId.(uint64)
	}

	code := codes.Codes{
		Code:     cf.Code,
		Language: cf.Language,
		Owner:    userId,
	}

	if err := code.Create(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/%d", code.ID))
}
