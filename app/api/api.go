// Package api ...
package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
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

// New ...
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
	} {
		app.AddDescribedController(c)
	}
	return app.Application
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

func stateToMap(s *game.State, p1, p2 player.Described) pack.AnyMap {
	m := make(pack.AnyMap)
	if s.Winner() != game.NoPlayer {
		m["winner"] = s.Winner()
	}
	m["currentPlayer"] = s.CurrentPlayer()
	m["player1"] = p1.Name()
	m["player2"] = p2.Name()
	m["boardSize"] = s.Rules().BoardSize()
	ps := make(map[int]interface{})
	for _, p := range s.Pieces() {
		pm := make(pack.AnyMap)
		pm["life"] = p.Life()
		pm["damage"] = p.Damage()
		c := s.CellForPiece(p)
		pm["cell"] = []int{c.Row(), c.Column()}
		pm["player"] = s.PlayerForPiece(p)
		ps[int(p.ID())] = pm
	}
	m["pieces"] = ps
	return m
}

type base struct {
	handler trim.Handler
}

func (b *base) Apply(h trim.Handler) {
	b.handler = h
}

func badType(t string) trim.Response {
	return response.NewJSON(
		pack.AnyMap{"message": fmt.Sprintf("must pass a %s", t)},
		trim.CodeBadRequest,
	)
}
