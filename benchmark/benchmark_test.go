package benchmark_test

import (
	"testing"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
)

var standardRules = game.NewRules(5, 1, 3, 1, 1)

func BenchmarkNextState(b *testing.B) {
	p1 := player.NewGreedy()
	p2 := player.NewGreedy()
	for i := 0; i < b.N; i++ {
		game.NextState(game.NewState(standardRules, p1, p2))
	}
}

func BenchmarkGame(b *testing.B) {
	p1 := player.NewGreedy()
	p2 := player.NewGreedy()
	for i := 0; i < b.N; i++ {
		s := game.NewState(standardRules, p1, p2)
		for s.Winner() == game.NoPlayer {
			s = game.NextState(s)
		}
	}
}
