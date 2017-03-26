package game

// Direction to move on a Board.
type Direction int

// Directions of movement possible.
const (
	North Direction = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

// Directions of movement enumerated in a list.
func Directions() []Direction {
	return []Direction{
		North,
		NorthEast,
		East,
		SouthEast,
		South,
		SouthWest,
		West,
		NorthWest,
	}
}

// String representation of the Direction.
func (d Direction) String() string {
	return map[Direction]string{
		North:     "north",
		NorthEast: "north-east",
		East:      "east",
		SouthEast: "south-east",
		South:     "south",
		SouthWest: "south-west",
		West:      "west",
		NorthWest: "north-west",
	}[d]
}
