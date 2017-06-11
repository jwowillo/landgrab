package player

import "github.com/jwowillo/landgrab/game"

// Greedy game.Player chooses the game.Play with the greatest value from all
// legal game.Plays.
type Greedy struct{}

// newGreedy game.DescribedPlayer.
func newGreedy() game.DescribedPlayer {
	return Greedy{}
}

// Name returns "greedy".
func (p Greedy) Name() string {
	return "greedy"
}

// Description of the game.DescribedPlayer.
func (p Greedy) Description() string {
	return "chooses the best play directly available"
}

// Play the turn by returning a random game.Play in the set of the highest-value
// legal game.Plays from the game.State.
func (p Greedy) Play(s *game.State) game.Play {
	return random(best(s))
}
