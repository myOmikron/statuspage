package dbmodels

import (
	"github.com/myOmikron/echotools/utilitymodels"
)

type State struct {
	utilitymodels.Common
	State       string
	Description string
}
