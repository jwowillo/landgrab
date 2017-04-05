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
var NoCell = Cell{-1, -1}

// TODO: Make follow standard map conventions and make more tolerant.

type cellMap struct {
	size  int
	cells []PieceID
}

func newCellMap(s int) *cellMap {
	cells := make([]PieceID, s*s)
	for i := 0; i < s*s; i++ {
		cells[i] = NoPieceID
	}
	return &cellMap{size: s, cells: cells}
}

func (l *cellMap) Set(c Cell, id PieceID) {
	l.cells[l.index(c)] = id
}

func (l *cellMap) Get(c Cell) PieceID {
	return l.cells[l.index(c)]
}

func (l *cellMap) Remove(c Cell) {
	l.cells[l.index(c)] = NoPieceID
}

func (l *cellMap) index(c Cell) int {
	return l.size*c.Row() + c.Column()
}

func (l *cellMap) clone() *cellMap {
	return &cellMap{size: l.size, cells: append([]PieceID{}, l.cells...)}
}
