package player

import "github.com/jwowillo/landgrab/game"

// Human ...
type Human struct {
	play game.Play
}

func newHuman() game.DescribedPlayer {
	return &Human{}
}

// SetPlay ...
func (p *Human) SetPlay(play game.Play) {
	p.play = play
}

// Name ...
func (p *Human) Name() string {
	return "human"
}

// Description ...
func (p *Human) Description() string {
	return "makes the play it was told to make"
}

// Play ...
func (p *Human) Play(s *game.State) game.Play {
	return p.play
}
