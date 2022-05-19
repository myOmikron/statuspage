package server

import (
	"github.com/labstack/echo/v4"
	"github.com/myOmikron/statuspage/conf"
	"github.com/myOmikron/statuspage/handler/protected"
	"github.com/myOmikron/statuspage/handler/status"
	"gorm.io/gorm"
)

func defineRoutes(e *echo.Echo, db *gorm.DB, config *conf.Config) {
	e.GET("/", status.Status(db))

	e.GET("/login", protected.Login)
	e.POST("/frontend/login", protected.LoginHandler(db))

	e.Static("/static/", config.Server.StaticPath)
}
