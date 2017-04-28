package game

import "time"

// Rules encapsulates the variable parts of games such as how many Pieces are
// involved, how much life and damage each Piece has, and how much these
// increase when they destroy enemy Pieces.
type Rules struct {
	timerDuration                                          time.Duration
	pieceCount, damage, life, damageIncrease, lifeIncrease int
}

// NewRules creates Rules with the given values for the variable parts.
func NewRules(td time.Duration, pc, l, d, li, di int) Rules {
	return Rules{
		timerDuration:  td,
		pieceCount:     pc,
		damage:         d,
		life:           l,
		damageIncrease: di,
		lifeIncrease:   li,
	}
}

// TimerDuration is the duration for the timer that limits the duration of each
// turn.
func (r Rules) TimerDuration() time.Duration {
	return r.timerDuration
}

// PieceCount is the number of Pieces held initially by each Player.
func (r Rules) PieceCount() int {
	return r.pieceCount
}

// BoardSize is 2 times the piece count plus 1 representing the length of 1 side
// of the board.
func (r Rules) BoardSize() int {
	return r.PieceCount()*2 + 1
}

// Life each Piece initially has which defines how much damage it can take
// before being destroyed.
func (r Rules) Life() int {
	return r.life
}

// Damage each Piece does when colliding with another Piece.
func (r Rules) Damage() int {
	return r.damage
}

// LifeIncrease is the amount a Piece's life increases every time it levels up.
func (r Rules) LifeIncrease() int {
	return r.lifeIncrease
}

// DamageIncrease is the amount a Piece's damage increases every time it levels
// up.
func (r Rules) DamageIncrease() int {
	return r.damageIncrease
}

// StandardRules a game is meant to be played by.
var StandardRules = NewRules(30*time.Second, 5, 3, 1, 1, 1)
