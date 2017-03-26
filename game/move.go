package game

// Move of a Piece in a Direction.
type Move struct {
	piece     Piece
	direction Direction
}

// NewMove with a valid Piece and Direction.
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
