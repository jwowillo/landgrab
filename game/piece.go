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

// pieeMap is an efficient mapping of Pieces to Cells that takes advantage of
// the sequential PieceIDs to hash Pieces into a slie.
type pieceMap struct {
	cells []Cell
}

// newPieceMap where each Player has the given amount of Pieces.
func newPieceMap(pc int) pieceMap {
	return pieceMap{cells: make([]Cell, pc*2)}
}

// Set the Piece to the Cell.
//
// If the Piece's id is NoPieceID, nothing is done.
func (m pieceMap) Set(p Piece, c Cell) {
	pid := p.ID()
	if pid == NoPieceID || int(pid) > len(m.cells) {
		return
	}
	m.cells[pid-1] = c
}

// Get the Cell associated with the Piece from the map.
func (m pieceMap) Get(p Piece) (Cell, bool) {
	pid := p.ID()
	if pid == NoPieceID || int(pid) > len(m.cells) {
		return NoCell, false
	}
	c := m.cells[pid-1]
	return c, c != NoCell
}

// Remove the Cell associated with the Piece and the Piece from the map.
func (m pieceMap) Remove(p Piece) {
	pid := p.ID()
	if pid == NoPieceID || int(pid) > len(m.cells) {
		return
	}
	m.cells[pid-1] = NoCell
}

// clone the pieceMap.
func (m pieceMap) clone() pieceMap {
	return pieceMap{cells: append([]Cell{}, m.cells...)}
}

// pieceIDMap efficiently maps PieceIDs to Pieces.
type pieceIDMap struct {
	pieceCount int
	pieces     []Piece
}

// newPieceIDMap where each Player has the given amount of Pieces.
func newPieceIDMap(pc int) pieceIDMap {
	return pieceIDMap{pieceCount: pc, pieces: make([]Piece, pc*2)}
}

// Set the PieceID to the Piece.
//
// This introduces the ability for the key PieceID to not match the Piece's
// PieceID but is necessary for efficient removal of a Piece from the map. An
// example of this is removing a Piece with PieceID 4. The value at PieceID 4
// will be set to NoPiece, even though NoPiece has a PieceID of 0. This still
// constitutes removing the Piece with PieceID 4 from the map.
func (m pieceIDMap) Set(pid PieceID, p Piece) {
	if pid == NoPieceID {
		return
	}
	m.pieces[pid-1] = p
}

// Get the Piece with the PieceID.
func (m pieceIDMap) Get(pid PieceID) (Piece, bool) {
	if pid == NoPieceID {
		return NoPiece, false
	}
	p := m.pieces[pid-1]
	return p, p != NoPiece
}

// Remove the Piece with the PieceID.
func (m pieceIDMap) Remove(pid PieceID) {
	if pid == NoPieceID {
		return
	}
	m.pieces[pid-1] = NoPiece
}

// Player1Pieces in the map.
func (m pieceIDMap) Player1Pieces() []Piece {
	return m.pieces[:m.pieceCount]
}

// Player2Pieces in the map.
func (m pieceIDMap) Player2Pieces() []Piece {
	return m.pieces[m.pieceCount:]
}

// clone the pieceIDMap.
func (m pieceIDMap) clone() pieceIDMap {
	return pieceIDMap{
		pieceCount: m.pieceCount,
		pieces:     append([]Piece{}, m.pieces...),
	}
}
