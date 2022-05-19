package protected

import (
	"github.com/labstack/echo/v4"
	"github.com/myOmikron/echotools/auth"
	"github.com/myOmikron/echotools/middleware"
	"gorm.io/gorm"
)

func Login(c echo.Context) error {
	return c.Render(200, "login", nil)
}

type LoginRequest struct {
	Username string
	Password string
}

func LoginHandler(db *gorm.DB) func(ctx echo.Context) error {
	return func(c echo.Context) error {
		var form LoginRequest
		echo.FormFieldBinder(c).String("username", &form.Username).String("password", &form.Password)

		if user, err := auth.Authenticate(db, form.Username, form.Password); err != nil {
			return c.String(400, err.Error())
		} else {
			if err := middleware.Login(db, user, c); err != nil {
				return c.String(400, err.Error())
			} else {
				return c.Redirect(302, "/protected/")
			}
		}
	}
}
