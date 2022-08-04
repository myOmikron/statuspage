package dbmodels

import (
	"github.com/myOmikron/echotools/utilitymodels"
)

type Object struct {
	utilitymodels.CommonID
	Name     string
	Disabled bool

	// Method to check the state of the object
	// 0: Manual - state changes are only allowed via API or logged-in user
	// 1: Ping
	// 2: HTTP Return Code
	// 3: TCP Connection
	CheckMethod uint

	StateHistory []State `gorm:"many2many:object__state;"`
}
