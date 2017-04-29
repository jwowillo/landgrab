package game_test

import (
	"testing"

	"github.com/jwowillo/landgrab/game"
)

// BenchmarkDirectionString benchmarks the conversion of all game.Directions to
// their string forms.
func BenchmarkDirectionString(b *testing.B) {
	ds := game.Directions()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, d := range ds {
			d.String()
		}
	}
}

// TestDirections tests that the game.Directions are the same as their
// prescripted values in the requirements.
func TestDirections(t *testing.T) {
	t.Parallel()
	if game.North.String() != "north" {
		t.Errorf(
			"game.North.String() = %s, want %s",
			game.North.String(),
			"north",
		)
	}
	if game.NorthEast.String() != "north-east" {
		t.Errorf(
			"game.NorthEast.String() = %s, want %s",
			game.NorthEast.String(),
			"north-east",
		)
	}
	if game.East.String() != "east" {
		t.Errorf(
			"game.East.String() = %s, want %s",
			game.East.String(),
			"east",
		)
	}
	if game.SouthEast.String() != "south-east" {
		t.Errorf(
			"game.SouthEast.String() = %s, want %s",
			game.SouthEast.String(),
			"south-east",
		)
	}
	if game.South.String() != "south" {
		t.Errorf(
			"game.South.String() = %s, want %s",
			game.South.String(),
			"south",
		)
	}
	if game.SouthWest.String() != "south-west" {
		t.Errorf(
			"game.SouthWest.String() = %s, want %s",
			game.SouthWest.String(),
			"south-west",
		)
	}
	if game.West.String() != "west" {
		t.Errorf(
			"game.West.String() = %s, want %s",
			game.West.String(),
			"west",
		)
	}
	if game.NorthWest.String() != "north-west" {
		t.Errorf(
			"game.NorthWest.String() = %s, want %s",
			game.NorthWest.String(),
			"north-west",
		)
	}
	if game.NoDirection.String() != "" {
		t.Errorf(
			"game.NoDirection.String() = %s, want %s",
			game.NoDirection.String(),
			"",
		)
	}
}

// TestDIrectionZeroValue tests that game.NoDirection is the same as
// game.Direction's zero-value.
func TestDirectionZeroValue(t *testing.T) {
	t.Parallel()
	var d game.Direction
	if game.NoDirection != d {
		t.Errorf("game.NoDirection = %v, want %v", game.NoDirection, d)
	}
}
