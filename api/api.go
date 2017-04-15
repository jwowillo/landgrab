// Package api has a constructor to build a trim.Application for the landgrab
// API.
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
	"github.com/jwowillo/pack"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
	"github.com/jwowillo/trim/trimming"
)

// TODO: Put descriptions in code.

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
	app.AddResource("Rules", jsonRules{})
	app.AddResource("Player", jsonPlayer{})
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

// stateToMap converts a game.State and two player.Describeds into a
// pack.AnyMap.
func stateToMap(s *game.State, p1, p2 player.Described) pack.AnyMap {
	m := make(pack.AnyMap)
	if s.Winner() != game.NoPlayer {
		m["winner"] = s.Winner()
	}
	m["currentPlayer"] = s.CurrentPlayer()
	m["player1"] = playerToJSON(p1)
	m["player2"] = playerToJSON(p2)
	m["rules"] = rulesToJSON(s.Rules())
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
	bs, err := json.Marshal(m["rules"])
	if err != nil {
		return nil, nil, nil, err
	}
	r, err := jsonToRules(bs)
	if err != nil {
		return nil, nil, nil, err
	}
	bs, err = json.Marshal(m["player1"])
	if err != nil {
		return nil, nil, nil, err
	}
	p1, err := jsonToPlayer(bs)
	if err != nil {
		return nil, nil, nil, err
	}
	bs, err = json.Marshal(m["player2"])
	if err != nil {
		return nil, nil, nil, err
	}
	p2, err := jsonToPlayer(bs)
	if err != nil {
		return nil, nil, nil, err
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

type jsonPlayer struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Arguments   map[string]string `json:"arguments,omitempty"`
}

func playerToJSON(p player.Described) jsonPlayer {
	return jsonPlayer{Name: p.Name(), Description: p.Description()}
}

func jsonToPlayer(bs []byte) (player.Described, error) {
	raw := jsonPlayer{}
	err := json.Unmarshal(bs, &raw)
	for _, p := range player.All() {
		if p.Name() == raw.Name {
			return p, err
		}
	}
	return nil, errors.New("bad \"game.Player\"")
}

type jsonRules struct {
	TimerDuration  int `json:"timerDuration"`
	PieceCount     int `json:"pieceCount"`
	BoardSize      int `json:"boardSize"`
	Life           int `json:"life"`
	Damage         int `json:"damage"`
	LifeIncrease   int `json:"lifeIncrease"`
	DamageIncrease int `json:"damageIncrease"`
}

func rulesToJSON(r game.Rules) jsonRules {
	return jsonRules{
		TimerDuration:  int(r.TimerDuration() / time.Second),
		PieceCount:     r.PieceCount(),
		BoardSize:      r.BoardSize(),
		Life:           r.Life(),
		Damage:         r.Damage(),
		LifeIncrease:   r.LifeIncrease(),
		DamageIncrease: r.DamageIncrease(),
	}
}

// mapToRules converts the pack.AnyMap to a game.Rules.
func jsonToRules(bs []byte) (game.Rules, error) {
	r := jsonRules{}
	err := json.Unmarshal(bs, &r)
	return game.NewRules(
		time.Duration(r.TimerDuration)*time.Second,
		r.PieceCount,
		r.Life,
		r.Damage,
		r.LifeIncrease,
		r.DamageIncrease,
	), err
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
