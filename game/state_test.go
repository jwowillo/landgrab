package game_test

import (
	"math/rand"
	"testing"

	"github.com/jwowillo/landgrab/game"
)

func BenchmarkNewState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		game.NewState(game.StandardRules, normal1{}, normal2{})
	}
}

func BenchmarkNewStateFromInfo(b *testing.B) {
	r := game.StandardRules
	p1 := normal1{}
	p2 := normal2{}
	s := game.NewState(r, p1, p2)
	cp := s.CurrentPlayer()
	ps := make(map[game.Cell]game.Piece)
	for i := 0; i < r.BoardSize(); i++ {
		for j := 0; j < r.BoardSize(); j++ {
			c := game.NewCell(i, j)
			p := s.PieceForCell(c)
			if c == game.NoCell {
				continue
			}
			ps[c] = p
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.NewStateFromInfo(r, cp, p1, p2, ps)
	}
}

func BenchmarkNextState(b *testing.B) {
	s := game.NewState(game.StandardRules, normal1{}, normal2{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.NextState(s)
	}
}

func BenchmarkNextStateWithPlay(b *testing.B) {
	s := game.NewState(game.StandardRules, normal1{}, normal2{})
	ps := game.LegalPlays(s)
	p := ps[rand.Intn(len(ps))]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.NextStateWithPlay(s, p)
	}
}
