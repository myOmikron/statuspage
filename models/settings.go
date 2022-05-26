package models

import "github.com/myOmikron/echotools/utilitymodels"

type Settings struct {
	utilitymodels.Common
	PageTitle string
	TabTitle  string
}
