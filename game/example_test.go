package game_test

import (
	"fmt"

	"github.com/jwowillo/landgrab/game"
)

// factory which returns normal1 and normal2 game.DescribedPlayers.
func factory() *game.PlayerFactory {
	f := game.NewPlayerFactory()
	f.Register(func() game.DescribedPlayer { return normal1{} })
	f.Register(func() game.DescribedPlayer { return normal2{} })
	return f
}

// Example of game.State being used.
func Example() {
	s := game.NewState(game.StandardRules, normal1{}, normal2{})
	if s.CurrentPlayer() == game.Player1 {
		fmt.Println("mismatch")
	}
	if s.NextPlayer() == game.Player2 {
		fmt.Println("mismatch")
	}
	if len(s.CurrentPlayerPieces()) != len(s.Player1Pieces()) {
		fmt.Println("mismatch")
	}
	if len(s.NextPlayerPieces()) != len(s.Player2Pieces()) {
		fmt.Println("mismatch")
	}
	if len(s.Player1Pieces())+len(s.Player2Pieces()) != len(s.Pieces()) {
		fmt.Println("mismatch")
	}
	if s.Rules() != game.StandardRules {
		fmt.Println("mismatch")
	}
	if s.Winner() != game.NoPlayer {
		fmt.Println("mismatch")
	}
	for _, p := range s.Player1Pieces() {
		if s.PlayerForPiece(p) != game.Player2 {
			fmt.Println("mismatch")
		}
	}
	for _, p := range s.Player2Pieces() {
		if s.PlayerForPiece(p) != game.Player2 {
			fmt.Println("mismatch")
		}
	}
	for _, p := range s.Player1Pieces() {
		if s.CellForPiece(p).Row() != 0 {
			fmt.Println("mismatch")
		}
	}
	for _, p := range s.Player2Pieces() {
		if s.CellForPiece(p).Row() != 11 {
			fmt.Println("mismatch")
		}
	}
	for _, c := range []game.Cell{
		game.NewCell(0, 1), game.NewCell(0, 3), game.NewCell(0, 5),
		game.NewCell(0, 9), game.NewCell(0, 11), game.NewCell(11, 1),
		game.NewCell(11, 3), game.NewCell(11, 5), game.NewCell(11, 5),
		game.NewCell(11, 7), game.NewCell(11, 9), game.NewCell(11, 11),
	} {
		if s.PieceForCell(c) == game.NoPiece {
			fmt.Println("mismatch")
		}
	}
	s = game.NextState(s)
	if s.CurrentPlayer() != game.Player2 {
		fmt.Println("mismatch")
	}
}

// ExampleLegalMoves shows that game.LegalMoves returns all the legal game.Moves
// from a game.State.
func Example_legalMoves() {
	f := factory()
	s := game.NewState(
		game.StandardRules,
		f.Player("normal1"), f.Player("normal2"),
	)
	for _, m := range game.LegalMoves(s) {
		if !game.IsLegalMove(s, m) {
			fmt.Println("illegal move")
		}
	}
	// Output:
}

// ExampleIsLegalMove shows that game.IsLegalMove correctly determines that all
// of player 1's pieces can initially move south.
func Example_isLegalMove() {
	f := factory()
	s := game.NewState(
		game.StandardRules,
		f.Player("normal1"), f.Player("normal2"),
	)
	for _, p := range s.Player1Pieces() {
		m := game.NewMove(p, game.South)
		if !game.IsLegalMove(s, m) {
			fmt.Println("illegal move")
		}
	}
	// Output:
}
