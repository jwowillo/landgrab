package player

import "github.com/jwowillo/landgrab/game"

// Greedy game.Player chooses the game.Play with the greatest value from all
// legal game.Plays.
type Greedy struct{}

// NewGreedy game.Player.
func NewGreedy() Greedy {
	return Greedy{}
}

// Play the turn by returning a random game.Play in the set of the highest-value
// legal game.Plays from the game.State.
func (p Greedy) Play(s *game.State) game.Play {
	return random(best(s, game.LegalPlays(s)))
}
