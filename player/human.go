package player

import (
	"reflect"
	"strings"

	"github.com/jwowillo/landgrab/game"
)

// Human ...
type Human struct {
	play game.Play
}

// NewHuman ...
func NewHuman() *Human {
	return &Human{}
}

// SetPlay ...
func (p *Human) SetPlay(play game.Play) {
	p.play = play
}

// Name ...
func (p *Human) Name() string {
	return strings.ToLower(reflect.TypeOf(p).Name())
}

// Description ...
func (p *Human) Description() string {
	return "makes the play it was told to make"
}

// Play ...
func (p *Human) Play(s *game.State) game.Play {
	return p.play
}
