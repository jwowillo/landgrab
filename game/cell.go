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
