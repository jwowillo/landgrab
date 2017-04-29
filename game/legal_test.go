package game_test

import (
	"math/rand"
	"testing"

	"github.com/jwowillo/landgrab/game"
)

func BenchmarkLegalPlays(b *testing.B) {
	s := game.NewState(game.StandardRules, normal1{}, normal2{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.LegalPlays(s)
	}
}

func BenchmarkIsLegalPlay(b *testing.B) {
	s := game.NewState(game.StandardRules, normal1{}, normal2{})
	ps := game.LegalPlays(s)
	p := ps[rand.Intn(len(ps))]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.IsLegalPlay(s, p)
	}
}

func BenchmarkLegalMoves(b *testing.B) {
	s := game.NewState(game.StandardRules, normal1{}, normal2{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.LegalMoves(s)
	}
}

func BenchmarkIsLegalMove(b *testing.B) {
	s := game.NewState(game.StandardRules, normal1{}, normal2{})
	ms := game.LegalMoves(s)
	m := ms[rand.Intn(len(ms))]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.IsLegalMove(s, m)
	}
}
