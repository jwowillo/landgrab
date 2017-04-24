// Package player has game.DescribedPlayer implementations which take different
// approaches to playing landgrab and an exported game.PlayerFactory instance
// used to construct the game.DescribedPlayers correctly.
//
// Anywhere random choices can be made, a time-seeded pseudo-random
// number-generator is used.
//
// Values of game.States are defined as the sum of the current
// game.DescribedPlayer's game.Piece's life and damage with ties broken by the
// lower manhattan-distance between all pieces. This has the effect of the
// game.DescribedPlayers tending to move their game.Pieces closer together but
// only when advantageous.
package player

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/jwowillo/landgrab/convert"
	"github.com/jwowillo/landgrab/game"
)

// Factory is a game.PlayerFactory instance which has all of the
// game.DescribedPlayers in this package registered properly.
//
// Possible names for game.DescribedPlayers along with any required data are
//   - "api" ({"url": <String URL for API>})
//   - "human" ({"play": <Play to execute in convert.JSONPlay form>})
//   - "random"
//   - "greedy"
//   - "search"
var Factory = game.NewPlayerFactory()

// init registers the implemented game.DescribedPlayers to the Factory instance.
func init() {
	Factory.Register(newGreedy)
	Factory.Register(newRandom)
	Factory.Register(newSearch)
	Factory.RegisterSpecial(
		newHuman,
		func(x game.DescribedPlayer, data map[string]interface{}) {
			p, ok := x.(*Human)
			if !ok {
				return
			}
			bs, err := json.Marshal(data)
			if err != nil {
				return
			}
			play, err := convert.JSONToPlay(bs)
			if err != nil {
				return
			}
			p.SetPlay(play)
		},
	)
	Factory.RegisterSpecial(
		newAPI,
		func(x game.DescribedPlayer, data map[string]interface{}) {
			p, ok := x.(*API)
			if !ok {
				return
			}
			val, ok := data["url"]
			if !ok {
				return
			}
			url, ok := val.(string)
			if !ok {
				return
			}
			p.SetURL(url)
		},
	)
}

// gen random values.
var gen = rand.New(rand.NewSource(time.Now().Unix()))

// Commonly used numeric constants.
const (
	// max int.
	max = int(^uint(0) >> 1)
	// min int.
	min = -max - 1
)

// random game.Play from the list.
func random(ps []game.Play) game.Play {
	return ps[gen.Intn(len(ps))]
}

// best game.Plays from the given game.State.
//
// Returns a list of game.Plays that all had the highest found value from the
// given game.State. To find this, the value of the next game.State from the
// current one is minimized, since the next game.State is for the enemy.
func best(s *game.State, ps []game.Play) []game.Play {
	best := max
	bestDistance := max
	var bestPlays []game.Play
	for _, p := range ps {
		n := game.NextStateWithPlay(s, p)
		v := value(n)
		if v == best {
			d := totalDistance(n)
			if totalDistance(n) < bestDistance {
				bestDistance = d
				bestPlays = []game.Play{p}
			} else {
				bestPlays = append(bestPlays, p)
			}
		}
		if v < best {
			best = v
			bestPlays = []game.Play{p}
		}
	}
	return bestPlays
}

// value of the game.States is the sum of the current game.Player's lifes and
// damages minus the sum o fthe next game.Player's lifes and damages.
//
// Two special cases are max is returned if a winning game.State is passed and
// min is returned otherwise.
func value(s *game.State) int {
	x := 0
	if s.Winner() == s.NextPlayer() {
		return min
	}
	if s.Winner() == s.CurrentPlayer() {
		return max
	}
	for _, p := range s.NextPlayerPieces() {
		x -= p.Life() + p.Damage()
	}
	for _, p := range s.CurrentPlayerPieces() {
		x += p.Life() + p.Damage()
	}
	return x
}

// totalDistance using the manhattan metric between all game.Pieces.
func totalDistance(s *game.State) int {
	total := 0
	for _, pa := range s.CurrentPlayerPieces() {
		for _, pb := range s.NextPlayerPieces() {
			a := s.CellForPiece(pa)
			b := s.CellForPiece(pb)
			total += manhattanDistance(a, b)
		}
	}
	return total
}

// manhattanDistance between two game.Cells.
func manhattanDistance(a, b game.Cell) int {
	dr := a.Row() - b.Row()
	if dr < 0 {
		dr = -dr
	}
	dc := a.Column() - b.Column()
	if dc < 0 {
		dc = -dc
	}
	return dr + dc
}
