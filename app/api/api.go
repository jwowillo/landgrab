// Package api ...
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
	"github.com/jwowillo/pack"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
	"github.com/jwowillo/trim/trimming"
)

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

// mapToState converts the pack.AnyMap representing a game.State to a
// game.State and the player.Describeds in the pack.AnyMap.
//
// Returns an error if the pack.AnyMap isn't formatted properly.
func mapToState(m pack.AnyMap) (
	*game.State,
	player.Described,
	player.Described,
	error,
) {
	current := game.NoPlayer
	if int(m["currentPlayer"].(float64)) == 1 {
		current = game.Player1
	} else if int(m["currentPlayer"].(float64)) == 2 {
		current = game.Player2
	} else {
		return nil, nil, nil, errors.New("bad \"currentPlayer\"")
	}
	r := game.StandardRules
	p1 := choosePlayer(m["player1"].(string))
	p2 := choosePlayer(m["player2"].(string))
	if p1 == nil || p2 == nil {
		return nil, nil, nil, errors.New(
			"bad \"player1\" or \"player2\"",
		)
	}
	var p1Pieces []game.Piece
	var p2Pieces []game.Piece
	pieces := make(map[game.Cell]game.Piece)
	for rawID, rawPiece := range m["pieces"].(map[string]interface{}) {
		id, err := strconv.Atoi(rawID)
		if err != nil {
			return nil, nil, nil, errors.New("bad \"id\"")
		}
		mapPiece := rawPiece.(map[string]interface{})
		rawCell := mapPiece["cell"].([]interface{})
		row := int(rawCell[0].(float64))
		column := int(rawCell[1].(float64))
		cell := game.NewCell(row, column)
		player := game.NoPlayer
		p := int(mapPiece["player"].(float64))
		if p == 1 {
			player = game.Player1
		}
		if p == 2 {
			player = game.Player2
		}
		life := int(mapPiece["life"].(float64))
		damage := int(mapPiece["damage"].(float64))
		piece := game.NewPiece(game.PieceID(id), life, damage)
		if player == game.Player1 {
			p1Pieces = append(p1Pieces, piece)
		}
		if player == game.Player2 {
			p2Pieces = append(p2Pieces, piece)
		}
		pieces[cell] = piece
	}
	return game.NewStateFromInfo(
		r,
		current,
		p1, p2,
		p1Pieces, p2Pieces,
		pieces,
	), p1, p2, nil
}

// playerToMap converts the player.Described to a pack.AnyMap.
func playerToMap(p player.Described) pack.AnyMap {
	return pack.AnyMap{"name": p.Name(), "description": p.Description()}
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

// choosePlayer with the name from player.Describeds in player.All.
func choosePlayer(name string) player.Described {
	for _, p := range player.All() {
		if p.Name() == name {
			return p
		}
	}
	return nil
}
