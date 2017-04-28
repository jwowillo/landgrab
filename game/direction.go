package game

// Direction to move on a Board.
type Direction int

// Directions of movement possible.
const (
	NoDirection Direction = iota // Direction zero-value.
	North
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
	switch d {
	case North:
		return "north"
	case NorthEast:
		return "north-east"
	case East:
		return "east"
	case SouthEast:
		return "south-east"
	case South:
		return "south"
	case SouthWest:
		return "south-west"
	case West:
		return "west"
	case NorthWest:
		return "north-west"
	default:
		return ""
	}
}
