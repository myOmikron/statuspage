package server

import (
	"github.com/labstack/echo/v4"
	"github.com/myOmikron/statuspage/conf"
	"github.com/myOmikron/statuspage/handler"
	"gorm.io/gorm"
)

func defineRoutes(e *echo.Echo, db *gorm.DB, config *conf.Config) {
	statusWrapper := handler.Wrapper{DB: db}

	e.GET("/", statusWrapper.Status)

	e.Static("/static/", config.Server.StaticPath)
}
