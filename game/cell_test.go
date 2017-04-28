package game_test

import (
	"testing"

	"github.com/jwowillo/landgrab/game"
)

// TestCell tests that game.Cell row and column are properly initialized and
// returned by getters.
func TestCell(t *testing.T) {
	t.Parallel()
	c := game.NewCell(3, 5)
	if c.Row() != 3 {
		t.Errorf("c.Row() = %d, want %d", c.Row(), 3)
	}
	if c.Column() != 5 {
		t.Errorf("c.Column() = %d, want %d", c.Column(), 5)
	}
}

// TestNoCellIsNegative tests that game.NoCell has a negative row and column.
func TestNoCellIsNegative(t *testing.T) {
	t.Parallel()
	if game.NoCell.Row() >= 0 || game.NoCell.Column() >= 0 {
		t.Errorf(
			"game.NoCell.Row() or game.NoCell.Column() is " +
				"positive, want negative",
		)
	}
}
