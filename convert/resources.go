package convert

import "github.com/jwowillo/landgrab/game"

// JSONPlay ...
type JSONPlay struct {
	Moves []JSONMove `json:"moves"`
}

// Description ...
func (p JSONPlay) Description() string {
	return "play of Moves"
}

// JSONMove ...
type JSONMove struct {
	Direction game.Direction `json:"direction"`
	Piece     JSONPiece      `json:"piece"`
}

// Description ...
func (p JSONMove) Description() string {
	return "move made by a Piece"
}

// JSONPlayer ...
type JSONPlayer struct {
	Name      string                 `json:"name"`
	Desc      string                 `json:"description,omitempty"`
	Arguments map[string]interface{} `json:"arguments,omitempty"`
}

// Description ...
func (p JSONPlayer) Description() string {
	return "makes moves"
}

// JSONPiece ...
type JSONPiece struct {
	ID     game.PieceID  `json:"id"`
	Player game.PlayerID `json:"player"`
	Life   int           `json:"life"`
	Damage int           `json:"damage"`
	Cell   [2]int        `json:"cell"`
}

// Description ...
func (s JSONPiece) Description() string {
	return "piece on the board"
}

// JSONRules ...
type JSONRules struct {
	TimerDuration  int `json:"timerDuration"`
	PieceCount     int `json:"pieceCount"`
	BoardSize      int `json:"boardSize"`
	Life           int `json:"life"`
	Damage         int `json:"damage"`
	LifeIncrease   int `json:"lifeIncrease"`
	DamageIncrease int `json:"damageIncrease"`
}

// Description ...
func (r JSONRules) Description() string {
	return "defines variable parts of a game"
}

// JSONState ...
type JSONState struct {
	CurrentPlayer game.PlayerID `json:"currentPlayer"`
	Winner        game.PlayerID `json:"winner,omitempty"`
	Rules         JSONRules     `json:"rules"`
	Player1       JSONPlayer    `json:"player1"`
	Player2       JSONPlayer    `json:"player2"`
	Pieces        []JSONPiece   `json:"pieces"`
}

// Description ...
func (s JSONState) Description() string {
	return "state of the game"
}
