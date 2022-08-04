package status

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/myOmikron/statuspage/models/dbmodels"
)

type SystemState struct {
	Name        string
	State       string
	Description string
}

type State struct {
	Title            string //
	PageTitle        string //
	OverallState     string //
	OverallStateText string //
	SystemStates     []SystemState
}

func Status(db *gorm.DB) func(ctx echo.Context) error {
	return func(c echo.Context) error {
		var settings dbmodels.Settings
		db.Find(&settings)

		state := State{
			Title:            settings.TabTitle,
			PageTitle:        settings.PageTitle,
			OverallState:     "critical",
			OverallStateText: "Critical error. There is no more Gulasch.",
			SystemStates:     make([]SystemState, 0),
		}
		state.SystemStates = append(state.SystemStates, SystemState{
			Name:        "Waffeln",
			State:       "ok",
			Description: "Normal operation",
		})
		state.SystemStates = append(state.SystemStates, SystemState{
			Name:        "Mail",
			State:       "maintenance",
			Description: "Planned maintenance due to bad sysadmins",
		})
		state.SystemStates = append(state.SystemStates, SystemState{
			Name:        "Gulasch",
			State:       "critical",
			Description: "There is no more Gulasch.",
		})
		state.SystemStates = append(state.SystemStates, SystemState{
			Name:        "Mate",
			State:       "warning",
			Description: "There is only a limited amount of Mate left",
		})
		return c.Render(200, "status", &state)
	}
}
