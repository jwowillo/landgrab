package api

import (
	"github.com/jwowillo/landgrab/convert"
	"github.com/jwowillo/landgrab/player"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
)

// playersPath is the playersController's path.
const playersPath = "/players"

// playersController is a trim.Controller used to retrieve all implemented
// game.Players.
type playersController struct{}

func (c playersController) Trimmings() []trim.Trimming {
	return []trim.Trimming{newValidateToken()}
}

// Path returns playersPath.
func (c playersController) Path() string {
	return playersPath
}

// Description of the playersController located at playersDescriptionPath.
func (c playersController) Description() *application.ControllerDescription {
	return &application.ControllerDescription{
		Get: &application.MethodDescription{
			Response:       "list of Players which can be used",
			Authentication: "must provide Token",
			Limiting:       "limit of the Token",
		},
	}
}

// Handle the trim.Request by returning all the implemented game.Players.
func (c playersController) Handle(r trim.Request) trim.Response {
	var ps []convert.JSONPlayer
	for _, name := range player.Factory.All() {
		ps = append(
			ps,
			convert.PlayerToJSONPlayer(player.Factory.Player(name)),
		)
	}
	return response.NewJSON(map[string][]convert.JSONPlayer{
		"players": ps,
	}, trim.CodeOK)
}
