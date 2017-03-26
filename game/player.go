package game

// Player in the game.
//
// The Player will be asked to return their choice of Play for each State where
// it is their turn in a game.
type Player interface {
	Play(*State) Play
}

// PlayerID identifies a Player in a game.
type PlayerID int

// PlayerIDs which can be given to Players.
const (
	NoPlayer PlayerID = iota
	Player1
	Player2
)

// String representation of the PlayerID.
func (id PlayerID) String() string {
	return map[PlayerID]string{
		NoPlayer: "no player",
		Player1:  "player 1",
		Player2:  "player 2",
	}[id]
}
