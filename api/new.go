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

// errBadPlayer is an error trim.Response returned when a bad value is provided
// for the game.Player type.
var errBadPlayer = badType("game.Player")

const (
	// newPath is the newController's path.
	newPath = "/new"
	// newPlayer1Key is the key for the game.Player 1 passed in the
	// trim.Context.
	newPlayer1Key = "player1"
	// newPlayer2Key is the key for the game.Player 2 passed in the
	// trim.Context.
	newPlayer2Key = "player2"
)

// newController is a trim.Controller used to create new game.States to play
// landgrab with.
type newController struct{}

// Path returns newPath.
func (c newController) Path() string {
	return newPath
}

// Description of the newController located at newDescriptionPath.
func (c newController) Description() *application.ControllerDescription {
	return &application.ControllerDescription{
		Get: &application.MethodDescription{
			FormArguments: map[string]string{
				newPlayer1Key: "Player for player 1",
				newPlayer2Key: "Player for player 2",
			},
			Response: "initial State",
		},
	}
}

// Trimmings returns a single trim.Trimming which validates that the
// trim.Request has valid game.Player 1 and 2s passed.
func (c newController) Trimmings() []trim.Trimming {
	return []trim.Trimming{newValidateNew()}
}

// Handle the trim.Request by converting the trim.Request's context to a new
// game.State and returning a JSON representation of it.
func (c newController) Handle(r trim.Request) trim.Response {
	p1 := r.Context()[newPlayer1Key].(game.DescribedPlayer)
	p2 := r.Context()[newPlayer2Key].(game.DescribedPlayer)
	s := game.NewState(game.StandardRules, p1, p2)
	return response.NewJSON(
		convert.StateToJSONState(s),
		trim.CodeOK,
	)
}

// validateNew is a validating trim.Trimming that validates input to the
// newController.
type validateNew struct {
	*base
}

// newValidateNew creates a validateNew.
func newValidateNew() validateNew {
	return validateNew{base: &base{}}
}

// Handle the trim.Request by parsing the query arguments into their real types
// and returning the newController's trim.Response.
//
// If the query arguments aren't game.Players, errBadPlayer is returned.
func (v validateNew) Handle(r trim.Request) trim.Response {
	p1Args := r.FormArgs()[newPlayer1Key]
	p2Args := r.FormArgs()[newPlayer2Key]
	if len(p1Args) != 1 || len(p2Args) != 1 {
		return errBadPlayer
	}
	up1, err := url.QueryUnescape(p1Args[0])
	up2, err := url.QueryUnescape(p2Args[0])
	if err != nil {
		return errBadPlayer
	}
	p1, err := convert.JSONToPlayer([]byte(up1), player.All())
	p2, err := convert.JSONToPlayer([]byte(up2), player.All())
	if p1 == nil || p2 == nil || err != nil {
		return errBadPlayer
	}
	r.SetContext(newPlayer1Key, p1)
	r.SetContext(newPlayer2Key, p2)
	return v.handler.Handle(r)
}
