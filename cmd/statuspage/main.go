package main

import (
	"fmt"
	"github.com/hellflame/argparse"
	"github.com/myOmikron/statuspage/server"
	"os"
)

func main() {
	parser := argparse.NewParser("statuspage", "", &argparse.ParserConfig{})

	configPath := parser.String("", "config-path", &argparse.Option{
		Inheritable: true,
		Default:     "/etc/statuspage/config.toml",
		Help:        "Specify an alternative config path. Defaults to: /etc/statuspage/config.toml",
	})

	startServer := parser.AddCommand("start-server", "Starts the server", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	createAdminUser := parser.AddCommand(
		"create-admin-user",
		"Creates an administration user",
		&argparse.ParserConfig{
			DisableDefaultShowHelp: true,
		},
	)

	if err := parser.Parse(nil); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	switch {
	case startServer.Invoked:
		server.Start(*configPath)
	case createAdminUser.Invoked:
		server.CreateAdminUser(*configPath)
	}
}
