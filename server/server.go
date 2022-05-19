package server

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/myOmikron/echotools/color"
	"github.com/myOmikron/echotools/execution"
	mw "github.com/myOmikron/echotools/middleware"
	"github.com/myOmikron/statuspage/conf"
	"github.com/pelletier/go-toml"
	"html/template"
	"io/fs"
	"io/ioutil"
	"os"
	"time"
)

func Start(configPath string) {
	config := &conf.Config{}

	if configBytes, err := ioutil.ReadFile(configPath); errors.Is(err, fs.ErrNotExist) {
		color.Printf(color.RED, "Config was not found at %s\n", configPath)
		b, _ := toml.Marshal(config)
		fmt.Print(string(b))
		os.Exit(1)
	} else {
		if err := toml.Unmarshal(configBytes, config); err != nil {
			panic(err)
		}
	}

	// Check for valid config values
	if err := config.Check(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	db := initializeDatabase(config)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Template rendering
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.gohtml")),
	}
	e.Renderer = renderer

	e.Use(emw.Logger())
	e.Use(emw.Recover())
	allowedHosts := make([]mw.AllowedHost, 0)
	for _, host := range config.Server.AllowedHosts {
		allowedHosts = append(allowedHosts, mw.AllowedHost{
			Host:  host.Host,
			Https: host.Https,
		})
	}
	secConfig := &mw.SecurityConfig{
		AllowedHosts:            allowedHosts,
		UseForwardedProtoHeader: config.Server.UseForwardedProtoHeader,
	}
	e.Use(mw.Security(secConfig))
	cookieAge := time.Hour * 4
	e.Use(mw.Session(db, &mw.SessionConfig{
		CookieName: "sessionid",
		CookieAge:  &cookieAge,
		CookiePath: "/",
	}))

	// Define routes
	defineRoutes(e, db, config)

	fmt.Println("http://127.0.0.1:8080")
	execution.SignalStart(e, "127.0.0.1:8080", &execution.Config{
		ReloadFunc: func() {
			Start(configPath)
		},
		StopFunc: func() {

		},
		TerminateFunc: func() {

		},
	})
}
