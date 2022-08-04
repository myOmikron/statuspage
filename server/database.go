package server

import (
	"fmt"
	"net"
	"net/url"
	"strconv"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/myOmikron/echotools/database"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/myOmikron/statuspage/models/conf"
	"github.com/myOmikron/statuspage/models/dbmodels"
)

func initializeDatabase(config *conf.Config) (db *gorm.DB) {
	var driver gorm.Dialector
	switch config.Database.Driver {
	case "sqlite":
		driver = sqlite.Open(config.Database.Name)
	case "mysql":
		mysqlConf := mysqlDriver.NewConfig()
		mysqlConf.Net = fmt.Sprintf("tcp(%s)", net.JoinHostPort(config.Database.Host, strconv.Itoa(int(config.Database.Port))))
		mysqlConf.DBName = config.Database.Name
		mysqlConf.User = config.Database.User
		mysqlConf.Passwd = config.Database.Password
		mysqlConf.ParseTime = true
		mysqlConf.Params = map[string]string{
			"charset": "utf8mb4",
		}
		driver = mysql.Open(mysqlConf.FormatDSN())
	case "postgresql":
		dsn := url.URL{
			Scheme: "postgres",
			User:   url.UserPassword(config.Database.User, config.Database.Password),
			Host:   net.JoinHostPort(config.Database.Host, strconv.Itoa(int(config.Database.Port))),
			Path:   config.Database.Name,
		}
		driver = postgres.Open(dsn.String())
	}

	db = database.Initialize(
		driver,

		dbmodels.Settings{},
		dbmodels.Object{},
		dbmodels.State{},
	)

	var settingCount int64
	db.Find(&dbmodels.Settings{}).Count(&settingCount)
	if settingCount == 0 {
		db.Create(&dbmodels.Settings{
			TabTitle:  "demo page - statuspage",
			PageTitle: "demo page",
		})
	}

	return
}
