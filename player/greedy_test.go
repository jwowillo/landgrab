package player_test

import (
	"testing"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
)

func BenchmarkGreedy(b *testing.B) {
	p1 := player.Factory.Player("greedy")
	p2 := player.Factory.Player("greedy")
	for i := 0; i < b.N; i++ {
		s := game.NewState(game.StandardRules, p1, p2)
		for s.Winner() == game.NoPlayer {
			s = game.NextState(s)
		}
	}
}
