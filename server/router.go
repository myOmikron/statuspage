package server

import (
	"github.com/labstack/echo/v4"
	"github.com/myOmikron/statuspage/models/dbmodels"
	"gorm.io/gorm"

	"github.com/myOmikron/statuspage/handler/protected"
	"github.com/myOmikron/statuspage/handler/status"
	"github.com/myOmikron/statuspage/models/conf"
)

func defineRoutes(
	e *echo.Echo,
	db *gorm.DB,
	config *conf.Config,
	settings *dbmodels.Settings,
	settingsReloadFunc func(),
) {
	protectedWrapper := protected.Wrapper{
		DB:                 db,
		Config:             config,
		Settings:           settings,
		SettingsReloadFunc: settingsReloadFunc,
	}

	statusWrapper := status.Wrapper{
		Config:             config,
		DB:                 db,
		Settings:           settings,
		SettingsReloadFunc: settingsReloadFunc,
	}

	e.GET("/", statusWrapper.Status)

	e.GET("/login", protectedWrapper.Login)
	e.POST("/frontend/login", protectedWrapper.LoginHandler)

	e.Static("/static/", config.Server.StaticPath)
}
