package api

import (
	"github.com/jwowillo/landgrab/player"
	"github.com/jwowillo/pack"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/controller"
	"github.com/jwowillo/trim/response"
)

const (
	// playersPath is the playersController's path.
	playersPath = "/players"
	// playersDescriptionPath is the path to the playersController's
	// description.
	playersDescriptionPath = descriptionBase + "players.json"
)

// playersController is a trim.Controller used to retrieve all implemented
// game.Players.
type playersController struct {
	controller.Bare
}

// Path returns playersPath.
func (c playersController) Path() string {
	return playersPath
}

// Description of the playersController located at playersDescriptionPath.
func (c playersController) Description() *application.ControllerDescription {
	return must(read(playersDescriptionPath))
}

// Handle the trim.Request by returning all the implemented game.Players.
func (c playersController) Handle(r trim.Request) trim.Response {
	var ps []pack.AnyMap
	for _, p := range player.All() {
		ps = append(ps, playerToMap(p))
	}
	return response.NewJSON(pack.AnyMap{"players": ps}, trim.CodeOK)
}
