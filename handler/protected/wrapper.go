package protected

import (
	"gorm.io/gorm"

	"github.com/myOmikron/statuspage/models/conf"
	"github.com/myOmikron/statuspage/models/dbmodels"
)

type Wrapper struct {
	Config             *conf.Config
	DB                 *gorm.DB
	Settings           *dbmodels.Settings
	SettingsReloadFunc func()
}
