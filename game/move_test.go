package game_test

import (
	"testing"

	"github.com/jwowillo/landgrab/game"
)

// TestMove tests that game.Move is initialized properly and the getters return
// the correct values.
func TestMove(t *testing.T) {
	t.Parallel()
	p := game.NewPiece(1, 3, 5)
	d := game.North
	m := game.NewMove(p, d)
	if m.Piece() != p {
		t.Errorf("m.Piece() = %v, want %v", m.Piece(), p)
	}
	if m.Direction() != d {
		t.Errorf("m.Direction() = %v, want %v", m.Direction(), d)
	}
}

// TestNoMoveIsZeroValue tests that game.NoMove is the same as game.Move's
// zero-value.
func TestNoMoveIsZeroValue(t *testing.T) {
	t.Parallel()
	var m game.Move
	if game.NoMove != m {
		t.Errorf("game.NoMove = %v, want %v", game.NoMove, m)
	}
}
