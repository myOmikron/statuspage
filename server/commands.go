package server

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/myOmikron/echotools/color"
	"github.com/myOmikron/echotools/database"
	"github.com/myOmikron/echotools/utilitymodels"
	"github.com/myOmikron/statuspage/conf"
	"github.com/pelletier/go-toml"
	"golang.org/x/term"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
)

func CreateAdminUser(configPath string) {
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

	// Ask things
	reader := bufio.NewReader(os.Stdin)

	color.Print(color.PURPLE, "Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		color.Println(color.RED, err.Error())
		os.Exit(1)
	}
	username = strings.TrimSuffix(username, "\n")

	color.Print(color.PURPLE, "Enter Password: ")
	bytePassword, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		color.Println(color.RED, err.Error())
		os.Exit(1)
	}
	fmt.Println()

	password := string(bytePassword)

	var count int64
	db.Find(&utilitymodels.LocalUser{}, "username = ?", username).Count(&count)
	if count != 0 {
		color.Println(color.RED, "There is already a user with that username registered")
		os.Exit(1)
	}

	if _, err := database.CreateLocalUser(db, username, password, nil); err != nil {
		color.Println(color.RED, err.Error())
		os.Exit(1)
	}

	color.Println(color.GREEN, "User created")
}
