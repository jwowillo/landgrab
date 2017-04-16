// Package api has a constructor to build a trim.Application for the landgrab
// API.
package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jwowillo/landgrab/convert"
	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/pack"
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
		trimming.NewCache(-1, 10000),
	} {
		app.AddTrimming(t)
	}
	for _, c := range []application.DescribedController{
		newController{},
		nextController{},
		playersController{},
		rulesController{},
	} {
		app.AddDescribedController(c)
	}
	app.AddResource("Rules", convert.JSONRules{})
	app.AddResource("Player", convert.JSONPlayer{
		Desc: "?description",
		Arguments: map[string]interface{}{
			"?arguments": "",
		},
	})
	app.AddResource("State", convert.JSONState{Winner: game.Player1})
	app.AddResource("Piece", convert.JSONPiece{})
	return app.Application
}

// descriptionBase is the base folder to find endpoint descriptions in.
const descriptionBase = "description/"

// must wraps a functions output that returns a byte slice and an error and
// panics if the error is not nil and returns the
// application.ControllerDescription otherwise.
func must(bs []byte, err error) *application.ControllerDescription {
	if err != nil {
		log.Println(err)
	}
	d := &application.ControllerDescription{}
	err = json.Unmarshal(bs, d)
	if err != nil {
		log.Println(err)
	}
	return d
}

// raed the file at the path and return its content as a byte slice and an error
// if the file couldn't be read properly.
func read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
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
		pack.AnyMap{"message": fmt.Sprintf("must pass a %s", t)},
		trim.CodeBadRequest,
	)
}
