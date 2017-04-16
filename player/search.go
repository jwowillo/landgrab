package player

import (
	"reflect"
	"strings"

	"github.com/jwowillo/landgrab/game"
)

// Search game.Player chooses the game.Play with the greatest value from all
// legal game.Plays.
type Search struct{}

// NewSearch game.Player.
func NewSearch() Search {
	return Search{}
}

// Name ...
func (p Search) Name() string {
	return strings.ToLower(reflect.TypeOf(p).Name())
}

// Description ...
func (p Search) Description() string {
	return "chooses the play leading to the best play within a radius"
}

// Play the turn by returning a random game.Play in the set of the highest-value
// legal game.Plays from the game.State.
func (p Search) Play(s *game.State) game.Play {
	return random(best(s, game.LegalPlays(s)))
}
