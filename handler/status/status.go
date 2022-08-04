package status

import (
	"github.com/labstack/echo/v4"

	"github.com/myOmikron/statuspage/models/dbmodels"
)

type state string

const (
	ok          state = "ok"
	maintenance       = "maintenance"
	warning           = "warning"
	critical          = "critical"
)

type ObjectState struct {
	Name        string
	State       state
	Description string
}

type Data struct {
	Title            string
	PageTitle        string
	OverallState     state
	OverallStateText string
	ObjectStates     []ObjectState
}

func (w *Wrapper) Status(c echo.Context) error {
	var objects []dbmodels.Object

	w.DB.Preload("StateHistory").Find(&objects, "disabled = ?", false)

	data := Data{
		Title:        w.Settings.TabTitle,
		PageTitle:    w.Settings.PageTitle,
		ObjectStates: make([]ObjectState, 0),
		OverallState: ok,
	}

	for _, obj := range objects {
		o := ObjectState{
			Name: obj.Name,
		}
		if len(obj.StateHistory) != 0 {
			o.State = state(obj.StateHistory[len(obj.StateHistory)-1].State)
			o.Description = obj.StateHistory[len(obj.StateHistory)-1].Description

			switch o.State {
			case critical:
				if data.OverallState == maintenance ||
					data.OverallState == warning ||
					data.OverallState == ok {
					data.OverallState = critical
				}
			case warning:
				if data.OverallState == ok || data.OverallState == maintenance {
					data.OverallState = warning
				}
			case maintenance:
				if data.OverallState == ok {
					data.OverallState = maintenance
				}
			}
		}
		data.ObjectStates = append(data.ObjectStates, o)
	}

	switch data.OverallState {
	case critical:
		data.OverallStateText = "Critical error(s) occurred."
	case warning:
		data.OverallStateText = "Warning(s) found."
	case maintenance:
		data.OverallStateText = "At least one component is currently in maintenance."
	case ok:
		data.OverallStateText = "All components are up and running."
	}

	return c.Render(200, "status", &data)
}
