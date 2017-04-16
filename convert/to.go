package convert

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/landgrab/player"
)

// PieceToJSONPiece ...
func PieceToJSONPiece(s *game.State, p game.Piece) JSONPiece {
	raw := JSONPiece{}
	raw.ID = p.ID()
	raw.Life = p.Life()
	raw.Damage = p.Damage()
	c := s.CellForPiece(p)
	raw.Cell = [2]int{c.Row(), c.Column()}
	raw.Player = s.PlayerForPiece(p)
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

// StateToJSONState ...
func StateToJSONState(s *game.State, p1, p2 player.Described) JSONState {
	raw := JSONState{}
	if s.Winner() != game.NoPlayer {
		raw.Winner = s.Winner()
	}
	raw.CurrentPlayer = s.CurrentPlayer()
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
func JSONStateToState(s JSONState) (
	*game.State,
	player.Described,
	player.Described,
) {
	p1 := JSONPlayerToPlayer(s.Player1)
	p2 := JSONPlayerToPlayer(s.Player2)
	var p1Pieces []game.Piece
	var p2Pieces []game.Piece
	Pieces := make(map[game.Cell]game.Piece)
	for _, rawPiece := range s.Pieces {
		Piece := JSONPieceToPiece(rawPiece)
		if rawPiece.Player == game.Player1 {
			p1Pieces = append(p1Pieces, Piece)
		}
		if rawPiece.Player == game.Player2 {
			p2Pieces = append(p2Pieces, Piece)
		}
		Pieces[game.NewCell(rawPiece.Cell[0], rawPiece.Cell[1])] = Piece
	}
	return game.NewStateFromInfo(
		JSONRulesToRules(s.Rules),
		s.CurrentPlayer,
		p1, p2,
		p1Pieces, p2Pieces,
		Pieces,
	), p1, p2
}

// JSONToState ...
func JSONToState(bs []byte) (
	*game.State,
	player.Described,
	player.Described,
	error,
) {
	rs, err := JSONToJSONState(bs)
	s, p1, p2 := JSONStateToState(rs)
	if p1 == nil || p2 == nil {
		err = errors.New("bad \"game.Player\"")
	}
	return s, p1, p2, err
}

// PlayerToJSONPlayer ...
func PlayerToJSONPlayer(p player.Described) JSONPlayer {
	return JSONPlayer{Name: p.Name(), Desc: p.Description()}
}

// JSONToJSONPlayer ...
func JSONToJSONPlayer(bs []byte) (JSONPlayer, error) {
	p := JSONPlayer{}
	err := json.Unmarshal(bs, &p)
	return p, err
}

// JSONPlayerToPlayer ...
func JSONPlayerToPlayer(raw JSONPlayer) player.Described {
	for _, p := range player.All() {
		if p.Name() == raw.Name {
			return p
		}
	}
	return nil
}

// JSONToPlayer ...
func JSONToPlayer(bs []byte) (player.Described, error) {
	p, err := JSONToJSONPlayer(bs)
	return JSONPlayerToPlayer(p), err
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
