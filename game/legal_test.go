package game_test

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/jwowillo/landgrab/game"
)

// TODO: TestLegalMoves.
// TODO: TestIsLegalPlay.
//   Make sure to test immutability.
// TODO: TestLegalPlay.

// BenchmarkLegalPlays benchmarks the performance of finding the legal
// game.Plays from a game.State.
func BenchmarkLegalPlays(b *testing.B) {
	s := game.NewState(game.StandardRules, normal1{}, normal2{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.LegalPlays(s)
	}
}

// BenchmarkIsLegalPlay benchmarks the cost of determining if a game.Play is
// legal.
func BenchmarkIsLegalPlay(b *testing.B) {
	s := game.NewState(game.StandardRules, normal1{}, normal2{})
	ps := game.LegalPlays(s)
	p := ps[rand.Intn(len(ps))]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.IsLegalPlay(s, p)
	}
}

// BenchmarkLegalMoves benchmarks the performance of finding all the legal
// game.Moves at a game.State.
func BenchmarkLegalMoves(b *testing.B) {
	s := game.NewState(game.StandardRules, normal1{}, normal2{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.LegalMoves(s)
	}
}

// BenchmarkIsLegalMove benchmarks the cost of determining if a game.Move is
// legal or not.
func BenchmarkIsLegalMove(b *testing.B) {
	s := game.NewState(game.StandardRules, normal1{}, normal2{})
	ms := game.LegalMoves(s)
	m := ms[rand.Intn(len(ms))]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.IsLegalMove(s, m)
	}
}

func TestLegalMoves(t *testing.T) {
	t.Parallel()
	rules := game.NewRules(30*time.Second, 2, 1, 1, 1, 1)
	p1 := game.NewPiece(game.PieceID(1), 1, 1)
	p2 := game.NewPiece(game.PieceID(2), 1, 1)
	p3 := game.NewPiece(game.PieceID(rules.PieceCount()+1), 1, 1)
	p4 := game.NewPiece(game.PieceID(rules.PieceCount()+2), 1, 1)
	cases := []struct {
		State *game.State
		Moves []game.Move
	}{
		{
			State: game.NewStateFromInfo(
				rules,
				game.Player2,
				normal1{}, normal2{},
				nil,
			),
			Moves: nil,
		},
		{
			State: game.NewStateFromInfo(
				rules,
				game.Player2,
				normal1{}, normal2{},
				map[game.Cell]game.Piece{
					game.NewCell(0, 0): p1,
					game.NewCell(0, 4): p2,
					game.NewCell(4, 0): p3,
					game.NewCell(4, 4): p4,
				},
			),
			Moves: []game.Move{
				game.NewMove(p3, game.North),
				game.NewMove(p3, game.NorthEast),
				game.NewMove(p3, game.East),
				game.NewMove(p4, game.North),
				game.NewMove(p4, game.NorthWest),
				game.NewMove(p4, game.West),
			},
		},
		{
			State: game.NewState(rules, normal1{}, normal2{}),
			Moves: []game.Move{
				game.NewMove(p1, game.West),
				game.NewMove(p1, game.SouthWest),
				game.NewMove(p1, game.South),
				game.NewMove(p1, game.SouthEast),
				game.NewMove(p1, game.East),
				game.NewMove(p2, game.West),
				game.NewMove(p2, game.SouthWest),
				game.NewMove(p2, game.South),
				game.NewMove(p2, game.SouthEast),
				game.NewMove(p2, game.East),
			},
		},
		{

			State: game.NewStateFromInfo(
				rules,
				game.Player2,
				normal1{}, normal2{},
				map[game.Cell]game.Piece{
					game.NewCell(2, 2): p1,
					game.NewCell(2, 3): p2,
					game.NewCell(3, 2): p3,
					game.NewCell(3, 3): p4,
				},
			),
			Moves: []game.Move{
				game.NewMove(p3, game.North),
				game.NewMove(p3, game.NorthWest),
				game.NewMove(p3, game.NorthEast),
				game.NewMove(p3, game.West),
				game.NewMove(p3, game.SouthWest),
				game.NewMove(p3, game.SouthEast),
				game.NewMove(p3, game.South),
				game.NewMove(p4, game.North),
				game.NewMove(p4, game.NorthWest),
				game.NewMove(p4, game.NorthEast),
				game.NewMove(p4, game.East),
				game.NewMove(p4, game.SouthWest),
				game.NewMove(p4, game.SouthEast),
				game.NewMove(p4, game.South),
			},
		},
		{
			State: game.NewStateFromInfo(
				rules,
				game.Player2,
				normal1{}, normal2{},
				map[game.Cell]game.Piece{
					game.NewCell(1, 1): p1,
					game.NewCell(1, 3): p2,
					game.NewCell(3, 1): p3,
					game.NewCell(3, 3): p4,
				},
			),
			Moves: []game.Move{
				game.NewMove(p3, game.North),
				game.NewMove(p3, game.NorthWest),
				game.NewMove(p3, game.NorthEast),
				game.NewMove(p3, game.West),
				game.NewMove(p3, game.East),
				game.NewMove(p3, game.SouthWest),
				game.NewMove(p3, game.SouthEast),
				game.NewMove(p3, game.South),
				game.NewMove(p4, game.North),
				game.NewMove(p4, game.NorthWest),
				game.NewMove(p4, game.NorthEast),
				game.NewMove(p4, game.West),
				game.NewMove(p4, game.East),
				game.NewMove(p4, game.SouthWest),
				game.NewMove(p4, game.SouthEast),
				game.NewMove(p4, game.South),
			},
		},
		{
			State: game.NewStateFromInfo(
				rules,
				game.Player2,
				normal1{}, normal2{},
				map[game.Cell]game.Piece{
					game.NewCell(0, 1): p1,
					game.NewCell(0, 3): p2,
				},
			),
			Moves: nil,
		},
		{
			State: game.NewStateFromInfo(
				rules,
				game.Player2,
				normal1{}, normal2{},
				map[game.Cell]game.Piece{
					game.NewCell(4, 1): p3,
					game.NewCell(4, 3): p4,
				},
			),
			Moves: []game.Move{
				game.NewMove(p3, game.West),
				game.NewMove(p3, game.North),
				game.NewMove(p3, game.NorthEast),
				game.NewMove(p3, game.NorthWest),
				game.NewMove(p3, game.East),
				game.NewMove(p4, game.West),
				game.NewMove(p4, game.North),
				game.NewMove(p4, game.NorthEast),
				game.NewMove(p4, game.NorthWest),
				game.NewMove(p4, game.East),
			},
		},
	}
	for _, test := range cases {
		ms := game.LegalMoves(test.State)
		actual := make(map[game.Move]struct{})
		expected := make(map[game.Move]struct{})
		for _, m := range ms {
			actual[m] = struct{}{}
			if !game.IsLegalMove(test.State, m) {
				t.Errorf(
					"game.LegalMoves(%v) contains %v",
					test.State, m,
				)
			}
		}
		for _, m := range test.Moves {
			expected[m] = struct{}{}
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected moves %v, got %v", expected, actual)
		}
	}
}

// TestIsLegalMove runs a table of test-cases against game.IsLegalMove to check
// against expected results.
func TestIsLegalMove(t *testing.T) {
	t.Parallel()
	rules := game.NewRules(30*time.Second, 2, 1, 1, 1, 1)
	p1 := game.NewPiece(game.PieceID(1), 1, 1)
	p2 := game.NewPiece(game.PieceID(2), 1, 1)
	p3 := game.NewPiece(game.PieceID(rules.PieceCount()+1), 1, 1)
	p4 := game.NewPiece(game.PieceID(rules.PieceCount()+2), 1, 1)
	p5 := game.NewPiece(game.PieceID(rules.PieceCount()+3), 1, 1)
	s := game.NewStateFromInfo(
		rules,
		game.Player1,
		normal1{}, normal2{},
		map[game.Cell]game.Piece{
			game.NewCell(0, 1): p1,
			game.NewCell(3, 3): p2,
			game.NewCell(3, 1): p3,
			game.NewCell(4, 3): p4,
		},
	)
	s2 := game.NewStateFromInfo(
		rules,
		game.Player1,
		normal1{}, normal2{},
		map[game.Cell]game.Piece{
			game.NewCell(1, 1): p1,
			game.NewCell(0, 4): p2,
			game.NewCell(4, 1): p3,
			game.NewCell(4, 3): p4,
		},
	)
	s3 := game.NewStateFromInfo(
		rules,
		game.Player1,
		normal1{}, normal2{},
		map[game.Cell]game.Piece{
			game.NewCell(0, 0): p1,
			game.NewCell(0, 1): p2,
			game.NewCell(4, 0): p3,
			game.NewCell(4, 1): p4,
		},
	)
	cases := []struct {
		State   *game.State
		Move    game.Move
		IsLegal bool
	}{
		{
			State:   s,
			Move:    game.NewMove(p1, game.North),
			IsLegal: false,
		},
		{
			State:   s,
			Move:    game.NewMove(p5, game.North),
			IsLegal: false,
		},
		{
			State:   s,
			Move:    game.NewMove(p1, game.South),
			IsLegal: true,
		},
		{
			State:   s,
			Move:    game.NewMove(game.NoPiece, game.South),
			IsLegal: false,
		},
		{
			State:   s,
			Move:    game.NewMove(p4, game.North),
			IsLegal: false,
		},
		{
			State:   s,
			Move:    game.NewMove(p4, game.South),
			IsLegal: false,
		},
		{
			State:   s,
			Move:    game.NewMove(p2, game.South),
			IsLegal: true,
		},
		{
			State:   s,
			Move:    game.NewMove(p4, game.North),
			IsLegal: false,
		},
		{
			State:   s2,
			Move:    game.NewMove(p1, game.NoDirection),
			IsLegal: false,
		},
		{
			State:   s2,
			Move:    game.NewMove(p1, game.North),
			IsLegal: true,
		},
		{
			State:   s2,
			Move:    game.NewMove(p1, game.NorthEast),
			IsLegal: true,
		},
		{
			State:   s2,
			Move:    game.NewMove(p1, game.East),
			IsLegal: true,
		},
		{
			State:   s2,
			Move:    game.NewMove(p1, game.SouthEast),
			IsLegal: true,
		},
		{
			State:   s2,
			Move:    game.NewMove(p1, game.South),
			IsLegal: true,
		},
		{
			State:   s2,
			Move:    game.NewMove(p1, game.SouthWest),
			IsLegal: true,
		},
		{
			State:   s2,
			Move:    game.NewMove(p1, game.West),
			IsLegal: true,
		},
		{
			State:   s2,
			Move:    game.NewMove(p1, game.NorthWest),
			IsLegal: true,
		},
		{
			State:   s3,
			Move:    game.NewMove(p1, game.East),
			IsLegal: false,
		},
		{
			State:   s3,
			Move:    game.NewMove(p3, game.East),
			IsLegal: false,
		},
	}
	for _, test := range cases {
		if game.IsLegalMove(test.State, test.Move) != test.IsLegal {
			t.Errorf(
				"game.IsLegalMove(%v, %v) = %v, want %v",
				test.State, test.Move,
				!test.IsLegal, test.IsLegal,
			)
		}
	}
}
