package game

// Rules encapsulates the variable parts of the game such as how many Pieces are
// involved, how much damage each Piece can do, how much life they have, and how
// much these increase when they destroy enemy Pieces.
type Rules struct {
	pieceCount, damage, life, damageIncrease, lifeIncrease int
}

// NewRules creates Rules with the given values for the variable parts.
func NewRules(pc, d, l, di, li int) Rules {
	return Rules{
		pieceCount:     pc,
		damage:         d,
		life:           l,
		damageIncrease: di,
		lifeIncrease:   li,
	}
}

// PieceCount specified by the Rules.
func (r Rules) PieceCount() int {
	return r.pieceCount
}

// BoardSize is 2 times the piece count + 1.
func (r Rules) BoardSize() int {
	return r.PieceCount()*2 + 1
}

// Damage specified by the rules.
func (r Rules) Damage() int {
	return r.damage
}

// Life specified by the rules.
func (r Rules) Life() int {
	return r.life
}

// DamageIncrease specified by the rules.
func (r Rules) DamageIncrease() int {
	return r.damageIncrease
}

// LifeIncrease specified by the rules.
func (r Rules) LifeIncrease() int {
	return r.lifeIncrease
}

// StandardRules ...
var StandardRules = NewRules(5, 1, 3, 1, 1)
