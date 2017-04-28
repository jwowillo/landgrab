package game_test

import (
	"testing"

	"github.com/jwowillo/landgrab/game"
)

// BenchmarkPlayerIDString benchmarks the conversion of all game.PlayerIDs to
// their string representations.
func BenchmarkPlayerIDString(b *testing.B) {
	pids := []game.PlayerID{game.NoPlayer, game.Player1, game.Player2}
	for i := 0; i < b.N; i++ {
		for _, pid := range pids {
			pid.String()
		}
	}
}

// TestPlayerIDString tests that game.PlayerID string values are the same as
// defined in the requirements.
func TestPlayerIDString(t *testing.T) {
	t.Parallel()
	if game.NoPlayer.String() != "no player" {
		t.Errorf(
			"game.NoPlayer.String() = %s, want %s",
			game.NoPlayer.String(),
			"no player",
		)
	}
	if game.Player1.String() != "player 1" {
		t.Errorf(
			"game.Player1.String() = %s, want %s",
			game.Player1.String(),
			"player 1",
		)
	}
	if game.Player2.String() != "player 2" {
		t.Errorf(
			"game.Player2.String() = %s, want %s",
			game.Player2.String(),
			"player 2",
		)
	}
}

// TestPlayerIDZeroValue tests that game.PlayerID's zero-value is game.NoPlayer.
func TestPlayerIDZeroValue(t *testing.T) {
	t.Parallel()
	var pid game.PlayerID
	if game.NoPlayer != pid {
		t.Errorf("game.NoPlayer = %v, want %v", game.NoPlayer, pid)
	}
}

// TestPlayerFactory tests that game.PlayerFactory correctly returns
// game.Players and game.SpecialPlayers and allows proper registration.
func TestPlayerFactory(t *testing.T) {
	t.Parallel()
	f := game.NewPlayerFactory()
	if len(f.All()) != 0 {
		t.Errorf("len(f.All()) = %d, want %d", len(f.All()), 0)
	}
	f.Register(func() game.DescribedPlayer { return normal1{} })
	f.Register(func() game.DescribedPlayer { return normal2{} })
	f.RegisterSpecial(
		func() game.DescribedPlayer { return &special1{} },
		func(p game.DescribedPlayer, data map[string]interface{}) {
			p.(*special1).value = data["value"].(string)
		},
	)
	f.RegisterSpecial(
		func() game.DescribedPlayer { return &special2{} },
		func(p game.DescribedPlayer, data map[string]interface{}) {
			p.(*special2).value = data["value"].(string)
		},
	)
	for _, p := range []string{
		"normal1",
		"normal2",
		"special2",
		"special2",
	} {
		doesContain := false
		for _, cp := range f.All() {
			if cp == p {
				doesContain = true
			}
		}
		if !doesContain {
			t.Errorf("f.All() doesn't contain %s", p)
		}
	}
	_, ok := f.Player("normal1").(normal1)
	if !ok {
		t.Errorf("f.Player(\"normal1\") doesn't return normal1")
	}
	_, ok = f.Player("normal2").(normal2)
	if !ok {
		t.Errorf("f.Player(\"normal2\") doesn't return normal2")
	}
	data := map[string]interface{}{"value": "value"}
	p1, ok := f.SpecialPlayer("special1", data).(*special1)
	if !ok {
		t.Errorf("f.Player(\"special1\") doesn't return special1")
	} else {
		if p1.value != "value" {
			t.Errorf("p1.value = %s, want %s", p1.value, "value")
		}
	}
	p2, ok := f.SpecialPlayer("special2", data).(*special2)
	if !ok {
		t.Errorf("f.Player(\"special2\") doesn't return special2")
	} else {
		if p2.value != "value" {
			t.Errorf("p2.value = %s, want %s", p2.value, "value")
		}
	}
}

// normal1 game.DescribedPlayer to test game.PlayerFactory with.
type normal1 struct{}

// Name for normal1.
func (p normal1) Name() string {
	return "normal1"
}

// Name for normal1.
func (p normal1) Description() string {
	return "normal1 description"
}

// Play for normal1.
func (p normal1) Play(s *game.State) game.Play {
	return nil
}

// normal2 game.DescribedPlayer to test game.PlayerFactory with.
type normal2 struct{}

// Name for normal2.
func (p normal2) Name() string {
	return "normal2"
}

// Description for normal2.
func (p normal2) Description() string {
	return "normal2 description"
}

// Play for normal2.
func (p normal2) Play(s *game.State) game.Play {
	return nil
}

// special1 game.DescribedPlayer to test game.PlayerFactory with.
type special1 struct {
	value string
}

// Name for special1.
func (p *special1) Name() string {
	return "special1"
}

// Description for special1.
func (p *special1) Description() string {
	return "special1 description"
}

// Play for special1.
func (p *special1) Play(s *game.State) game.Play {
	return nil
}

// special2 game.DescribedPlayer to test game.PlayerFactory with.
type special2 struct {
	value string
}

// Name for special2.
func (p *special2) Name() string {
	return "special2"
}

// Description for special2.
func (p *special2) Description() string {
	return "special2 description"
}

// Play for special2.
func (p *special2) Play(s *game.State) game.Play {
	return nil
}
