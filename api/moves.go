package api

import (
	"github.com/jwowillo/landgrab/convert"
	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
)

const (
	movesPath     = "/moves"
	movesStateKey = "state"
)

type movesController struct{}

func (c movesController) Path() string {
	return movesPath
}

func (c movesController) Description() *application.ControllerDescription {
	return &application.ControllerDescription{
		Get: &application.MethodDescription{
			FormArguments: map[string]string{
				movesStateKey: "State to find moves for",
			},
			Response:       "Moves for each Piece",
			Authentication: "must provide Token",
			Limiting:       "limit of the Token",
		},
	}
}

func (c movesController) Trimmings() []trim.Trimming {
	return []trim.Trimming{
		newValidateToken(),
		newValidateMoves(),
	}
}

func (c movesController) Handle(r trim.Request) trim.Response {
	s := r.Context()[movesStateKey].(*game.State)
	setMoves := make(map[game.PieceID]map[game.Move]struct{})
	for _, m := range game.LegalMoves(s) {
		if s.PlayerForPiece(m.Piece()) != s.CurrentPlayer() {
			continue
		}
		id := m.Piece().ID()
		if _, ok := setMoves[id]; !ok {
			setMoves[id] = make(map[game.Move]struct{})
		}
		setMoves[id][m] = struct{}{}
	}
	moves := make(map[game.PieceID][]convert.JSONMove)
	for id, ms := range setMoves {
		for m := range ms {
			moves[id] = append(
				moves[id],
				convert.MoveToJSONMove(m, s),
			)
		}
	}
	return response.NewJSON(moves, trim.CodeOK)
}

type validateMoves struct {
	*base
}

func newValidateMoves() validateMoves {
	return validateMoves{base: &base{}}
}

func (v validateMoves) Handle(r trim.Request) trim.Response {
	if err := parseState(r, movesStateKey, ""); err != nil {
		return err
	}
	return v.handler.Handle(r)
}
