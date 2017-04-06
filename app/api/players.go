package api

import (
	"github.com/jwowillo/landgrab/player"
	"github.com/jwowillo/pack"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
)

type playersController struct{}

func (c playersController) Path() string {
	return "/players"
}

func (c playersController) Description() *application.ControllerDescription {
	return must(read(descriptionBase + "players.json"))
}

func (c playersController) Trimmings() []trim.Trimming {
	return nil
}

func (c playersController) Handle(r trim.Request) trim.Response {
	var ps []pack.AnyMap
	for _, p := range player.All() {
		ps = append(ps, playerToMap(p))
	}
	return response.NewJSON(pack.AnyMap{"players": ps}, trim.CodeOK)
}

func playerToMap(p player.Described) pack.AnyMap {
	return pack.AnyMap{"name": p.Name(), "description": p.Description()}
}
