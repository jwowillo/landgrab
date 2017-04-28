package game_test

import (
	"testing"

	"github.com/jwowillo/landgrab/game"
)

// TestPiece tests that game.Piece is initialized properly and the ID, life, and
// damage are correctly returned.
func TestPiece(t *testing.T) {
	t.Parallel()
	p := game.NewPiece(1, 3, 5)
	if p.ID() != 1 {
		t.Errorf("p.ID() = %d, want %d", p.ID(), 1)
	}
	if p.Life() != 3 {
		t.Errorf("p.Life() = %d, want %d", p.Life(), 3)
	}
	if p.Damage() != 5 {
		t.Errorf("p.Damage() = %d, want %d", p.Damage(), 5)
	}
}

// TestNoPieceIDIsZeroValue tests that the zero-value of game.PieceID is the
// same as game.NoPieceID.
func TestNoPieceIDIsZeroValue(t *testing.T) {
	t.Parallel()
	var pid game.PieceID
	if game.NoPieceID != pid {
		t.Errorf("game.NoPieceID = %d, want %d", game.NoPieceID, pid)
	}
}

// TestNoPieceIsZeroValue tests that the zero-value of game.Piece is the same as
// game.NoPiece.
func TestNoPieceIsZeroValue(t *testing.T) {
	t.Parallel()
	var p game.Piece
	if game.NoPiece != p {
		t.Errorf("game.NoPiece = %v, want %v", game.NoPiece, p)
	}
}
