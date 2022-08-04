package protected

import (
	"github.com/labstack/echo/v4"
	"github.com/myOmikron/echotools/auth"
	"github.com/myOmikron/echotools/middleware"
)

func (w *Wrapper) Login(c echo.Context) error {
	return c.Render(200, "login", nil)
}

type LoginRequest struct {
	Username string
	Password string
}

func (w *Wrapper) LoginHandler(c echo.Context) error {
	var form LoginRequest
	echo.FormFieldBinder(c).String("username", &form.Username).String("password", &form.Password)

	if user, err := auth.AuthenticateLocalUser(w.DB, form.Username, form.Password); err != nil {
		return c.Redirect(302, "/login")
	} else {
		if err := middleware.Login(w.DB, user, c); err != nil {
			return c.Redirect(302, "/login")
		} else {
			return c.Redirect(302, "/protected/")
		}
	}
}
