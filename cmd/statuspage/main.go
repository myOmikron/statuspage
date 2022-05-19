package main

import (
	"fmt"
	"github.com/hellflame/argparse"
	"github.com/myOmikron/statuspage/server"
	"os"
)

func main() {
	parser := argparse.NewParser("statuspage", "", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	configPath := parser.String("", "config-path", &argparse.Option{
		Default: "/etc/statuspage/config.toml",
		Help:    "Specify an alternative config path. Defaults to: /etc/statuspage/config.toml",
	})

	if err := parser.Parse(nil); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	server.Start(*configPath)
}
