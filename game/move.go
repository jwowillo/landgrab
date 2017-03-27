package game

// Move of a Piece in a Direction.
type Move struct {
	piece     Piece
	direction Direction
}

// NewMove with a valid Piece and Direction.
//
// NoPiece should not be passed as the Piece and -1 shouldn't be passed as the
// Direction as they're reserved for special uses.
func NewMove(p Piece, d Direction) Move {
	return Move{piece: p, direction: d}
}

// Piece making the Move.
func (m Move) Piece() Piece {
	return m.piece
}

// Direction of the Move.
func (m Move) Direction() Direction {
	return m.direction
}

// Play is a turn in the game represented by a list of Moves the Player is
// making.
type Play []Move

// NoMove is the absence of a Move.
var NoMove = NewMove(NoPiece, Direction(-1))
