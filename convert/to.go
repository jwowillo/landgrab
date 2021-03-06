package convert

import (
	"encoding/json"
	"time"

	"github.com/jwowillo/landgrab/game"
)

// PlayToJSONPlay ...
func PlayToJSONPlay(p game.Play, s *game.State) JSONPlay {
	ms := make([]JSONMove, len(p))
	for i, m := range p {
		ms[i] = MoveToJSONMove(m, s)
	}
	return JSONPlay{Moves: ms}
}

// JSONToJSONPlay ...
func JSONToJSONPlay(bs []byte) (JSONPlay, error) {
	p := JSONPlay{}
	err := json.Unmarshal(bs, &p)
	return p, err
}

// JSONPlayToPlay ...
func JSONPlayToPlay(p JSONPlay) game.Play {
	play := make(game.Play, len(p.Moves))
	for i, move := range p.Moves {
		play[i] = JSONMoveToMove(move)
	}
	return play
}

// JSONToPlay ...
func JSONToPlay(bs []byte) (game.Play, error) {
	play, err := JSONToJSONPlay(bs)
	return JSONPlayToPlay(play), err
}

// MoveToJSONMove ...
func MoveToJSONMove(m game.Move, s *game.State) JSONMove {
	return JSONMove{
		Direction: m.Direction().String(),
		Piece:     PieceToJSONPiece(s, m.Piece()),
	}
}

// JSONToJSONMove ...
func JSONToJSONMove(bs []byte) (JSONMove, error) {
	m := JSONMove{}
	err := json.Unmarshal(bs, &m)
	return m, err
}

// JSONMoveToMove ...
func JSONMoveToMove(m JSONMove) game.Move {
	var d game.Direction
	switch m.Direction {
	case "north":
		d = game.North
	case "north-east":
		d = game.NorthEast
	case "east":
		d = game.East
	case "south-east":
		d = game.SouthEast
	case "south":
		d = game.South
	case "south-west":
		d = game.SouthWest
	case "west":
		d = game.West
	case "north-west":
		d = game.NorthWest
	}
	return game.NewMove(JSONPieceToPiece(m.Piece), d)
}

// JSONToMove ...
func JSONToMove(bs []byte) (game.Move, error) {
	move, err := JSONToJSONMove(bs)
	return JSONMoveToMove(move), err
}

// PieceToJSONPiece ...
func PieceToJSONPiece(s *game.State, p game.Piece) JSONPiece {
	raw := JSONPiece{}
	raw.ID = p.ID()
	raw.Life = p.Life()
	raw.Damage = p.Damage()
	c := s.CellForPiece(p)
	raw.Cell = [2]int{c.Row(), c.Column()}
	raw.Player = s.PlayerForPiece(p).String()
	return raw
}

// JSONToJSONPiece ...
func JSONToJSONPiece(bs []byte) (JSONPiece, error) {
	p := JSONPiece{}
	err := json.Unmarshal(bs, &p)
	return p, err
}

// JSONPieceToPiece ...
func JSONPieceToPiece(p JSONPiece) game.Piece {
	return game.NewPiece(p.ID, p.Life, p.Damage)
}

// JSONToPiece ...
func JSONToPiece(bs []byte) (game.Piece, error) {
	p, err := JSONToJSONPiece(bs)
	return JSONPieceToPiece(p), err
}

func stringToPlayerID(x string) game.PlayerID {
	switch x {
	case "no player":
		return game.NoPlayer
	case "player 1":
		return game.Player1
	case "player 2":
		return game.Player2
	}
	return game.NoPlayer
}

// StateToJSONState ...
func StateToJSONState(s *game.State) JSONState {
	raw := JSONState{}
	if s.Winner() != game.NoPlayer {
		raw.Winner = s.Winner().String()
	}
	raw.CurrentPlayer = s.CurrentPlayer().String()
	p1, p1Ok := s.Player1().(game.DescribedPlayer)
	p2, p2Ok := s.Player2().(game.DescribedPlayer)
	if !p1Ok || !p2Ok {
		p1 = nil
		p2 = nil
	}
	raw.Player1 = PlayerToJSONPlayer(p1)
	raw.Player2 = PlayerToJSONPlayer(p2)
	raw.Rules = RulesToJSONRules(s.Rules())
	for _, p := range s.Pieces() {
		raw.Pieces = append(raw.Pieces, PieceToJSONPiece(s, p))
	}
	return raw
}

// JSONToJSONState ...
func JSONToJSONState(bs []byte) (JSONState, error) {
	s := JSONState{}
	err := json.Unmarshal(bs, &s)
	return s, err
}

// JSONStateToState ...
func JSONStateToState(s JSONState, factory *game.PlayerFactory) *game.State {
	p1 := JSONPlayerToPlayer(s.Player1, factory)
	p2 := JSONPlayerToPlayer(s.Player2, factory)
	Pieces := make(map[game.Cell]game.Piece)
	for _, rawPiece := range s.Pieces {
		Piece := JSONPieceToPiece(rawPiece)
		Pieces[game.NewCell(rawPiece.Cell[0], rawPiece.Cell[1])] = Piece
	}
	return game.NewStateFromInfo(
		JSONRulesToRules(s.Rules),
		stringToPlayerID(s.CurrentPlayer),
		p1, p2,
		Pieces,
	)
}

// JSONToState ...
func JSONToState(bs []byte, factory *game.PlayerFactory) (*game.State, error) {
	rs, err := JSONToJSONState(bs)
	s := JSONStateToState(rs, factory)
	return s, err
}

// PlayerToJSONPlayer ...
func PlayerToJSONPlayer(p game.DescribedPlayer) JSONPlayer {
	return JSONPlayer{Name: p.Name(), Desc: p.Description()}
}

// JSONToJSONPlayer ...
func JSONToJSONPlayer(bs []byte) (JSONPlayer, error) {
	p := JSONPlayer{}
	err := json.Unmarshal(bs, &p)
	return p, err
}

// JSONPlayerToPlayer ...
func JSONPlayerToPlayer(
	raw JSONPlayer,
	factory *game.PlayerFactory,
) game.DescribedPlayer {
	for _, name := range factory.All() {
		if name == raw.Name {
			return factory.SpecialPlayer(name, raw.Arguments)
		}
	}
	return nil
}

// JSONToPlayer ...
func JSONToPlayer(
	bs []byte,
	factory *game.PlayerFactory,
) (game.DescribedPlayer, error) {
	p, err := JSONToJSONPlayer(bs)
	return JSONPlayerToPlayer(p, factory), err
}

// RulesToJSONRules ...
func RulesToJSONRules(r game.Rules) JSONRules {
	return JSONRules{
		TimerDuration:  int(r.TimerDuration() / time.Second),
		PieceCount:     r.PieceCount(),
		BoardSize:      r.BoardSize(),
		Life:           r.Life(),
		Damage:         r.Damage(),
		LifeIncrease:   r.LifeIncrease(),
		DamageIncrease: r.DamageIncrease(),
	}
}

// JSONToJSONRules ...
func JSONToJSONRules(bs []byte) (JSONRules, error) {
	r := JSONRules{}
	err := json.Unmarshal(bs, &r)
	return r, err
}

// JSONRulesToRules ...
func JSONRulesToRules(r JSONRules) game.Rules {
	return game.NewRules(
		time.Duration(r.TimerDuration)*time.Second,
		r.PieceCount,
		r.Life,
		r.Damage,
		r.LifeIncrease,
		r.DamageIncrease,
	)
}

// JSONToRules ...
func JSONToRules(bs []byte) (game.Rules, error) {
	r, err := JSONToJSONRules(bs)
	return JSONRulesToRules(r), err
}
