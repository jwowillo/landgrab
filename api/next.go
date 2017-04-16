package api

import (
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
	nextStateKey = "state"
	nextPlayKey  = "play"
	// nextPlayer1Key is the key for the game.Player 1 passed in the
	// trim.Context.
	nextPlayer1Key = "player1"
	// nextPlayer2Key is the key for the game.Player 2 passed in the
	// trim.Context.
	nextPlayer2Key = "player2"
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
	p, ok := r.Context()[nextPlayKey]
	if ok {
		s = game.NextStateWithPlay(s, p.(game.Play))
	} else {
		s = game.NextState(s)
	}
	p1 := r.Context()[nextPlayer1Key].(game.DescribedPlayer)
	p2 := r.Context()[nextPlayer2Key].(game.DescribedPlayer)
	return response.NewJSON(
		convert.StateToJSONState(s, p1, p2),
		trim.CodeOK,
	)
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
	sArgs := r.FormArgs()[nextStateKey]
	if len(sArgs) != 1 {
		return errBadState
	}
	unquoted, err := url.QueryUnescape(sArgs[0])
	if err != nil {
		return errBadState
	}
	s, p1, p2, err := convert.JSONToState([]byte(unquoted), player.All())
	if s.Rules() != game.StandardRules || err != nil {
		return errBadState
	}
	r.SetContext(nextStateKey, s)
	r.SetContext(nextPlayer1Key, p1)
	r.SetContext(nextPlayer2Key, p2)
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
