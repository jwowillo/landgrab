package benchmark_test

import (
	"testing"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
)

func BenchmarkNextState(b *testing.B) {
	p1 := player.NewGreedy()
	p2 := player.NewGreedy()
	for i := 0; i < b.N; i++ {
		s := game.NewState(game.StandardRules, p1, p2)
		s = game.NextState(s)
		s = game.NextState(s)
	}
}

func BenchmarkGame(b *testing.B) {
	p1 := player.NewGreedy()
	p2 := player.NewGreedy()
	for i := 0; i < b.N; i++ {
		s := game.NewState(game.StandardRules, p1, p2)
		for s.Winner() == game.NoPlayer {
			s = game.NextState(s)
		}
	}
}
