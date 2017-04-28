package game_test

import (
	"testing"
	"time"

	"github.com/jwowillo/landgrab/game"
)

// TestRules tests that game.Rules is properly initialized and has getters that
// return the proper values.
func TestRules(t *testing.T) {
	t.Parallel()
	r := game.NewRules(time.Second, 3, 5, 7, 9, 11)
	testRules(t, r, time.Second, 3, 5, 7, 9, 11)
}

// TestStandardRules test that game.StandardRules has the values defined by the
// requirements.
func TestStandardRules(t *testing.T) {
	t.Parallel()
	testRules(t, game.StandardRules, 30*time.Second, 5, 3, 1, 1, 1)
}

// testRules tests that the game.Rules getters return the given values.
func testRules(
	t *testing.T,
	r game.Rules,
	td time.Duration, pc, l, d, li, di int,
) {
	if r.TimerDuration() != td {
		t.Errorf(
			"r.TimerDuration() = %v, want %v",
			r.TimerDuration(),
			td,
		)
	}
	if r.PieceCount() != pc {
		t.Errorf("r.PieceCount() = %d, want %d", r.PieceCount(), pc)
	}
	if r.BoardSize() != 2*pc+1 {
		t.Errorf("r.BoardSize() = %d, want %d", r.BoardSize(), 2*pc+1)
	}
	if r.Life() != l {
		t.Errorf("r.Life() = %d, want %d", r.Life(), l)
	}
	if r.Damage() != d {
		t.Errorf("r.Damage() = %d, want %d", r.Damage(), d)
	}
	if r.LifeIncrease() != li {
		t.Errorf("r.LifeIncrease() = %d, want %d", r.LifeIncrease(), li)
	}
	if r.DamageIncrease() != di {
		t.Errorf(
			"r.DamageIncrease() = %d, want %d",
			r.DamageIncrease(),
			li,
		)
	}
}
