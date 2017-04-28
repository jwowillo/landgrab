package game

// PieceID uniquely identifies a Piece in a game.
//
// PieceID 0 is left for special purposes and shouldn't be assigned to Pieces
// outside of the package.
type PieceID int

// Piece in a game.
//
// Pieces are uniquely identified within a game by a PieceID. They also have
// life which indicates a piece is destroyed if the life is zero. Finally,
// Pieces have damage which is how much they deduct from other enemy Pieces in
// collisions.
//
// The zero-value Piece represents the absence of a Piece and shouldn't be used
// outside of the package.
type Piece struct {
	id           PieceID
	life, damage int
}

// NewPiece identified by the PieceID with the given life and damage.
func NewPiece(id PieceID, l, d int) Piece {
	return Piece{id: id, life: l, damage: d}
}

// ID uniquely identifying the Piece within a game.
func (p Piece) ID() PieceID {
	return p.id
}

// Life remaining on the Piece.
func (p Piece) Life() int {
	return p.life
}

// Damage the Piece does to other Pieces.
func (p Piece) Damage() int {
	return p.damage
}

// NoPieceID is the ID of no Piece.
//
// Note that this is the same as the zero-value for PieceID.
const NoPieceID = 0

// NoPiece is the absence of a Piece on a Cell.
//
// Note that this is the same as the zero-value for Piece.
var NoPiece = NewPiece(NoPieceID, 0, 0)
