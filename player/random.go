package player

import "github.com/jwowillo/landgrab/game"

// Random game.Player chooses a random game.Play from all legal game.Plays.
type Random struct{}

// NewRandom game.Player.
func NewRandom() Random {
	return Random{}
}

// Play the turn by returning a random game.Play in the set of legal game.Plays
// from the game.State.
func (p Random) Play(s *game.State) game.Play {
	return random(game.LegalPlays(s))
}
