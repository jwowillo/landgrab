package player

import "github.com/jwowillo/landgrab/game"

// Random game.Player chooses a random game.Play from all legal game.Plays.
type Random struct{}

func newRandom() game.DescribedPlayer {
	return Random{}
}

// Name ...
func (p Random) Name() string {
	return "random"
}

// Description ...
func (p Random) Description() string {
	return "chooses a random play"
}

// Play the turn by returning a random game.Play in the set of legal game.Plays
// from the game.State.
func (p Random) Play(s *game.State) game.Play {
	return random(game.LegalPlays(s))
}
