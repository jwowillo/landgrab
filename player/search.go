package player

import "github.com/jwowillo/landgrab/game"

// Search game.Player chooses the game.Play which leads to the greatest value
// game.State within a search radius.
type Search struct{}

// newSearch game.Player.
func newSearch() game.DescribedPlayer {
	return Search{}
}

// Name returns "search".
func (p Search) Name() string {
	return "search"
}

// Description ...
func (p Search) Description() string {
	return "chooses the play leading to the best play within a radius"
}

// Play by searching for the highest value game.State within a set search radius
// and returning the game.Play that leads to it.
func (p Search) Play(s *game.State) game.Play {
	return random(best(s, game.LegalPlays(s)))
}
