package dbmodels

import (
	"github.com/myOmikron/echotools/utilitymodels"
)

type Settings struct {
	utilitymodels.CommonID
	PageTitle string
	TabTitle  string
}
