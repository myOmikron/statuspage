package server

import (
	"github.com/labstack/echo/v4"
	"github.com/myOmikron/statuspage/conf"
	"gorm.io/gorm"
)

func defineRoutes(e *echo.Echo, db *gorm.DB, config *conf.Config) {

	e.Static("/static/", config.Server.StaticPath)
}
