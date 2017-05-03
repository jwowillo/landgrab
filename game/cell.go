package game

// Cell located at a row and column.
type Cell struct {
	row, column int
}

// NewCell located at the row and column.
//
// Negative rows and columns shouldn't be passed as they are reserved for use
// inside the package.
func NewCell(r, c int) Cell {
	return Cell{row: r, column: c}
}

// Row of the cell.
func (c Cell) Row() int {
	return c.row
}

// Column of the Cell.
func (c Cell) Column() int {
	return c.column
}

// NoCell represents no Cell exists.
//
// This Cell should not be used to indicate anything other than the absence of a
// Cell.
//
// Unlike other variables in the package which represent similar things, this is
// not equivalent to the zero-value of a Cell. The reason for this is that the
// zero-value is Cell{0, 0}, which would mean counting over the game-grid would
// have to start at 1 everywhere. That breaks too many conventions.
var NoCell = Cell{-1, -1}

// cellMap is a mapping which efficiently relates Cells to Pieces by taking
// advantage of the Cell and grid structure to perfectly hash the Cells into a
// slice.
type cellMap struct {
	size  int
	cells []Piece
}

// newCellMap where the grid has sides of the given size.
func newCellMap(s int) cellMap {
	return cellMap{size: s, cells: make([]Piece, s*s)}
}

// Set the key Cell to the PieceID.
func (l cellMap) Set(c Cell, p Piece) {
	i := l.index(c)
	if i < 0 || i >= len(l.cells) {
		return
	}
	l.cells[i] = p
}

// Get the PieceID at the key Cell from the mapping.
//
// Also return a bool that is true iff the mapping contained the Cell as a key.
func (l cellMap) Get(c Cell) (Piece, bool) {
	i := l.index(c)
	if i < 0 || i >= len(l.cells) {
		return NoPiece, false
	}
	p := l.cells[i]
	return p, p != NoPiece
}

// Remove the Cell and its mapped PieceID from the mapping.
func (l cellMap) Remove(c Cell) {
	i := l.index(c)
	if i < 0 || i >= len(l.cells) {
		return
	}
	l.cells[i] = NoPiece
}

// index of a Cell in the mapping.
func (l cellMap) index(c Cell) int {
	return l.size*c.Row() + c.Column()
}

// clone the cellMap into a new one.
func (l cellMap) clone() cellMap {
	return cellMap{size: l.size, cells: append([]Piece{}, l.cells...)}
}
