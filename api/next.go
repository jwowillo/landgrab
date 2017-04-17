package api

import (
	"encoding/json"
	"net/url"

	"github.com/jwowillo/landgrab/convert"
	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
)

var (
	// errBadState is an error trim.Response returned when a bad value is provided
	// for the game.State type.
	errBadState = badType("game.State")
	errBadPlay  = badType("game.Play")
)

const (
	// nextPath is the nextController's path.
	nextPath = "/next"
	// nextStateKey is the key for the game.State passed in the
	// trim.Context.
	nextStateKey     = "state"
	nextJSONStateKey = "json-state"
	nextPlayKey      = "play"
)

// nextController is a trim.Controller used to get the next game.State from a
// game.State.
type nextController struct{}

// Path returns nextPath.
func (c nextController) Path() string {
	return nextPath
}

// Description of the nextController located at nextDescriptionPath.
func (c nextController) Description() *application.ControllerDescription {
	return &application.ControllerDescription{
		Get: &application.MethodDescription{
			FormArguments: map[string]string{
				nextStateKey:      "State to find the next State of",
				"?" + nextPlayKey: "optional Play to use for the next State",
			},
			Response: "next State",
		},
	}
}

// Trimmings returns a single trim.Trimming which validates that the
// trim.Request has a valid game.State and game.Player 1 and 2s passed.
func (c nextController) Trimmings() []trim.Trimming {
	return []trim.Trimming{newValidateNext()}
}

// Handle the trim.Request by returning the next game.State to the game.State
// passed in the trim.Requests context.
func (c nextController) Handle(r trim.Request) trim.Response {
	s := r.Context()[nextStateKey].(*game.State)
	ojs := r.Context()[nextJSONStateKey].(convert.JSONState)
	p, ok := r.Context()[nextPlayKey]
	if ok {
		s = game.NextStateWithPlay(s, p.(game.Play))
	} else {
		s = game.NextState(s)
	}
	js := convert.StateToJSONState(s)
	js.Player1 = ojs.Player1
	js.Player2 = ojs.Player2
	return response.NewJSON(js, trim.CodeOK)
}

// validateNext is a validating trim.Trimming that validates input to the
// nextController.
type validateNext struct {
	*base
}

// newValidateNext creates a validateNext.
func newValidateNext() validateNext {
	return validateNext{base: &base{}}
}

// Handle the trim.Request by parsing the query arguments into their real types
// and returning the nextController's trim.Response.
//
// If the query arguments aren't a game.State, errBadState is returned.
func (v validateNext) Handle(r trim.Request) trim.Response {
	if err := parseState(r, nextStateKey, nextJSONStateKey); err != nil {
		return err
	}
	pArgs, ok := r.FormArgs()[nextPlayKey]
	if ok {
		if len(pArgs) != 1 {
			return errBadPlay
		}
		unquoted, err := url.QueryUnescape(pArgs[0])
		if err != nil {
			return errBadPlay
		}
		p, err := convert.JSONToPlay([]byte(unquoted))
		if err != nil {
			return errBadPlay
		}
		r.SetContext(nextPlayKey, p)
	}
	return v.handler.Handle(r)
}

func parseState(r trim.Request, skey, jskey string) trim.Response {
	sArgs := r.FormArgs()[skey]
	if len(sArgs) != 1 {
		return errBadState
	}
	unquoted, err := url.QueryUnescape(sArgs[0])
	if err != nil {
		return errBadState
	}
	js, err := convert.JSONToJSONState([]byte(unquoted))
	if err != nil {
		return errBadState
	}
	s := convert.JSONStateToState(js, player.All())
	if s.Rules() != game.StandardRules {
		return errBadState
	}
	handleSpecial(s.Player1().(game.DescribedPlayer), js.Player1)
	handleSpecial(s.Player2().(game.DescribedPlayer), js.Player2)
	r.SetContext(jskey, js)
	r.SetContext(skey, s)
	return nil
}

func handleSpecial(p game.DescribedPlayer, raw convert.JSONPlayer) {
	switch real := p.(type) {
	case *player.Human:
		val, ok := raw.Arguments["play"]
		if !ok {
			return
		}
		bs, err := json.Marshal(val)
		if err != nil {
			return
		}
		play, err := convert.JSONToPlay(bs)
		if err != nil {
			return
		}
		real.SetPlay(play)
	case *player.API:
		val, ok := raw.Arguments["url"]
		if !ok {
			return
		}
		url, ok := val.(string)
		if !ok {
			return
		}
		real.SetURL(url)
	}
}
