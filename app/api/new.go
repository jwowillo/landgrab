package api

import (
	"encoding/json"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
)

var (
	errBadRules  = badType("game.Rules")
	errBadPlayer = badType("game.Player")
)

const (
	// newPath is the newController's path.
	newPath = "/new"
	// newDescription is the path to the newController's description.
	newDescriptionPath = descriptionBase + "new.json"
	// newRulesKey is the key for the game.Rules passed in the trim.Context.
	newRulesKey = "rules"
	// newPlayer1Key is the key for the game.Player for player 1 passed in
	// the trim.Context.
	newPlayer1Key = "player1"
	// newPlayer2Key is the key for the game.Player for player 2 passed in
	// the trim.Context.
	newPlayer2Key = "player2"
)

// newController is a trim.Controller used to create new game.States to play
// landgrab with.
//
// Documentation can be accessed with a GET request to '/schema'.
type newController struct{}

// Path returns newPath.
func (c newController) Path() string {
	return newPath
}

// Description of the newController located at newDescriptionPath.
func (c newController) Description() *application.ControllerDescription {
	return must(read(newDescriptionPath))
}

// Trimmings returns a single trim.Trimming which validates that the
// trim.Request has proper game.Rules and game.Player 1 and 2s passed.
func (c newController) Trimmings() []trim.Trimming {
	return []trim.Trimming{newValidateNew()}
}

// Handle the trim.Request by converting the trim.Context to a new game.State
// and returning a JSON representation of it.
func (c newController) Handle(r trim.Request) trim.Response {
	s := game.NewState(
		r.Context()[newRulesKey].(game.Rules),
		r.Context()[newPlayer1Key].(game.Player),
		r.Context()[newPlayer2Key].(game.Player),
	)
	return response.NewJSON(stateToMap(s), trim.CodeOK)
}

type validateNew struct {
	*base
}

func newValidateNew() validateNew {
	return validateNew{base: &base{}}
}

func (v validateNew) Handle(r trim.Request) trim.Response {
	rArgs := r.FormArgs()[newRulesKey]
	p1Args := r.FormArgs()[newPlayer1Key]
	p2Args := r.FormArgs()[newPlayer2Key]
	if len(rArgs) != 1 {
		return errBadRules
	}
	if len(p1Args) != 1 {
		return errBadPlayer
	}
	if len(p2Args) != 1 {
		return errBadPlayer
	}
	var rules game.Rules
	var p1, p2 game.Player
	if err := json.Unmarshal([]byte(rArgs[0]), rules); err != nil {
		return errBadRules
	}
	if err := json.Unmarshal([]byte(p1Args[0]), p1); err != nil {
		return errBadPlayer
	}
	if err := json.Unmarshal([]byte(p2Args[0]), p2); err != nil {
		return errBadPlayer
	}
	r.SetContext(newRulesKey, rules)
	r.SetContext(newPlayer1Key, p1)
	r.SetContext(newPlayer2Key, p2)
	return v.handler.Handle(r)
}
