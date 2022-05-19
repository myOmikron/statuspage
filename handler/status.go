package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Wrapper struct {
	DB *gorm.DB
}

func (w *Wrapper) Status(c echo.Context) error {
	return c.Render(200, "status", nil)
}
