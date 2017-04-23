package player

import "github.com/jwowillo/landgrab/game"

// Human makes a set Play provided to it assuming the entity setting the
// game.Play can see the current game.State.
//
// Human is a special game.DescribedPlayer in that it needs its game.Play
// initialized.
type Human struct {
	play game.Play
}

// newHuman constructs an uninitialized Human game.DescribedPlayer.
func newHuman() game.DescribedPlayer {
	return &Human{}
}

// SetPlay that will be made.
func (p *Human) SetPlay(play game.Play) {
	p.play = play
}

// Name returns "human".
func (p *Human) Name() string {
	return "human"
}

// Description of the game.DescribedPlayer.
func (p *Human) Description() string {
	return "makes the play it was told to make"
}

// Play by returning the preset game.Play to make.
func (p *Human) Play(s *game.State) game.Play {
	return p.play
}
