// Package player has game.Player implementations which take different
// approaches to playing landgrab.
package player

import (
	"math/rand"
	"time"

	"github.com/jwowillo/landgrab/game"
)

// gen random values.
var gen = rand.New(rand.NewSource(time.Now().Unix()))

const (
	// max int.
	max = int(^uint(0) >> 1)
	// min int.
	min = -max - 1
)

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
	var bestPlays []game.Play
	for _, p := range ps {
		v := value(game.NextStateWithPlay(s, p))
		if v == best {
			bestPlays = append(bestPlays, p)
		}
		if v < best {
			best = v
			bestPlays = []game.Play{p}
		}
	}
	return bestPlays
}
