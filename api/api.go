// Package api has a constructor to build a trim.Application for the landgrab
// API.
package api

import (
	"fmt"

	"github.com/jwowillo/landgrab/convert"
	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
	"github.com/jwowillo/trim/trimming"
)

// New landgrab API.
func New() *application.Application {
	app := application.NewAPI()
	for _, t := range []trim.Trimming{
		trimming.NewAllow(trim.MethodGet),
		trimming.NewPreflight(),
	} {
		app.AddTrimming(t)
	}
	for _, c := range []application.DescribedController{
		newController{},
		nextController{},
		playersController{},
		rulesController{},
		movesController{},
		newTokenController(),
	} {
		app.AddDescribedController(c)
	}
	app.AddResource("Token", &token{})
	app.AddResource("Rules", convert.JSONRules{})
	app.AddResource("Player", convert.JSONPlayer{
		Desc: "?description",
		Arguments: map[string]interface{}{
			"?arguments": "",
		},
	})
	app.AddResource("State", convert.JSONState{Winner: game.Player1.String()})
	app.AddResource("Piece", convert.JSONPiece{})
	app.AddResource("Play", convert.JSONPlay{})
	app.AddResource("Move", convert.JSONMove{})
	return app.Application
}

// base trim.Trimming.
type base struct {
	handler trim.Handler
}

// Apply the trim.Handler to the trim.Trimming.
func (b *base) Apply(h trim.Handler) {
	b.handler = h
}

// badType returns an error.response when a type passed in a query string is
// bad.
func badType(t string) trim.Response {
	return response.NewJSON(
		map[string]string{"message": fmt.Sprintf("must pass a %s", t)},
		trim.CodeBadRequest,
	)
}
