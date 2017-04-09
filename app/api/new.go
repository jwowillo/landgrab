package api

import (
	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
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
	p1 := r.Context()[newPlayer1Key].(player.Described)
	p2 := r.Context()[newPlayer2Key].(player.Described)
	s := game.NewState(game.StandardRules, p1, p2)
	return response.NewJSON(stateToMap(s, p1, p2), trim.CodeOK)
}

type validateNew struct {
	*base
}

func newValidateNew() validateNew {
	return validateNew{base: &base{}}
}

func (v validateNew) Handle(r trim.Request) trim.Response {
	p1Args := r.FormArgs()[newPlayer1Key]
	p2Args := r.FormArgs()[newPlayer2Key]
	if len(p1Args) != 1 || len(p2Args) != 1 {
		return errBadPlayer
	}
	players := make(map[string]player.Described)
	for _, p := range player.All() {
		players[p.Name()] = p
	}
	p1 := players[p1Args[0]]
	p2 := players[p2Args[0]]
	if p1 == nil || p2 == nil {
		return errBadPlayer
	}
	r.SetContext(newPlayer1Key, p1)
	r.SetContext(newPlayer2Key, p2)
	return v.handler.Handle(r)
}
