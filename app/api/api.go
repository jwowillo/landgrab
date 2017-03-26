// Package api ...
package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/pack"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
	"github.com/jwowillo/trim/trimming"
)

// TODO: Configure the trimming.Package on API correctly.
// TODO: Add JSON descriptions.
// TODO: Document.
// TODO: Properly build game.Rules, game.Players, and game.States when that is
// implemented.

// Configure the application.API to serve the landgrab api.
func Configure(api *application.API) {
	for _, t := range []trim.Trimming{
		trimming.NewAllow(trim.MethodGet),
		trimming.NewCache(-1, 10000),
	} {
		api.AddTrimming(t)
	}
	for _, c := range []application.DescribedController{
		newController{},
		nextController{},
	} {
		api.AddDescribedController(c)
	}
}

const descriptionBase = "description/"

func must(bs []byte, err error) *application.ControllerDescription {
	if err != nil {
		panic(err)
	}
	d := &application.ControllerDescription{}
	err = json.Unmarshal(bs, d)
	if err != nil {
		panic(err)
	}
	return d
}

func read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func stateToMap(s *game.State) pack.AnyMap {
	return pack.AnyMap{}
}

type base struct {
	handler trim.Handler
}

func (b *base) Apply(h trim.Handler) {
	b.handler = h
}

func badType(t string) trim.Response {
	return response.NewJSON(
		pack.AnyMap{"error": fmt.Sprintf("must pass a %s", t)},
		trim.CodeBadRequest,
	)
}
