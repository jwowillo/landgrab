package api

import (
	"encoding/json"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
)

var errBadState = badType("game.State")

const (
	nextPath            = "/next"
	nextDescriptionPath = descriptionBase + "next.json"
	nextStateKey        = "state"
)

type nextController struct{}

func (c nextController) Path() string {
	return nextPath
}

func (c nextController) Description() *application.ControllerDescription {
	return must(read(nextDescriptionPath))
}

func (c nextController) Trimmings() []trim.Trimming {
	return []trim.Trimming{newValidateNext()}
}

func (c nextController) Handle(r trim.Request) trim.Response {
	s := game.NextState(r.Context()[nextStateKey].(*game.State))
	return response.NewJSON(stateToMap(s), trim.CodeOK)
}

type validateNext struct {
	*base
}

func newValidateNext() validateNext {
	return validateNext{base: &base{}}
}

func (v validateNext) Handle(r trim.Request) trim.Response {
	sArgs := r.FormArgs()[nextStateKey]
	if len(sArgs) != 1 {
		return errBadState
	}
	s := &game.State{}
	if err := json.Unmarshal([]byte(sArgs[0]), s); err != nil {
		return errBadState
	}
	r.SetContext(nextStateKey, s)
	return v.handler.Handle(r)
}
