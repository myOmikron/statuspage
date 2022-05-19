package status

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Status(db *gorm.DB) func(ctx echo.Context) error {
	return func(c echo.Context) error {
		return c.Render(200, "status", nil)
	}
}
