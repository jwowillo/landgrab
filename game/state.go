package game

// State encapsulates all of the game data in an immutable fashion.
type State struct {
	players       map[PlayerID]Player
	pieces        map[PlayerID]map[PieceID]Piece
	currentPlayer PlayerID
	winner        PlayerID
	rules         Rules
	cells         map[PieceID]Cell
	cellPieces    map[Cell]PieceID
}

// NewState creates an initial game State where the game is being played by
// Player one and Player two with the given Rules.
//
// Player one is set to move first.
func NewState(r Rules, p1, p2 Player) *State {
	p1s := make(map[PieceID]Piece)
	p2s := make(map[PieceID]Piece)
	ls := make(map[PieceID]Cell)
	cs := make(map[Cell]PieceID)
	for i := 0; i < r.PieceCount(); i++ {
		p1 := NewPiece(PieceID(i*2+1), r.Life(), r.Damage())
		p2 := NewPiece(PieceID(i*2), r.Life(), r.Damage())
		c1 := NewCell(0, i*2+1)
		c2 := NewCell(r.BoardSize()-1, i*2+1)
		p1s[p1.ID()] = p1
		p2s[p2.ID()] = p2
		ls[p1.ID()] = c1
		ls[p2.ID()] = c2
		cs[c1] = p1.ID()
		cs[c2] = p2.ID()
	}
	return &State{
		currentPlayer: Player1,
		winner:        NoPlayer,
		rules:         r,
		cells:         ls,
		cellPieces:    cs,
		players:       map[PlayerID]Player{Player1: p1, Player2: p2},
		pieces: map[PlayerID]map[PieceID]Piece{
			Player1: p1s,
			Player2: p2s,
		},
	}
}

// NextState returns the next State with the Play the current Player chooses.
func NextState(s *State) *State {
	s = NextStateWithPlay(s, s.players[s.CurrentPlayer()].Play(s))
	return s
}

// NextStateWithPlay returns the next State ignoring what the current Player
// would've done and instead uses the moves in the given Play.
func NextStateWithPlay(s *State, ms Play) *State {
	s = clone(s)
	if s.IsGameOver() {
		return s
	}
	set := make(map[Piece]struct{})
	for _, m := range ms {
		if _, ok := set[m.Piece()]; !ok {
			applyMove(s, m)
			set[m.Piece()] = struct{}{}
		}
	}
	handleDestroyed(s, ms)
	s.currentPlayer = s.NextPlayer()
	s.winner = winner(s)
	return s
}

// CurrentPlayer returns the PlayerID of the Player who is playing in this
// State.
func (s *State) CurrentPlayer() PlayerID {
	return s.currentPlayer
}

// NextPlayer returns the PlayerID of the Player who will play in the next
// State.
func (s *State) NextPlayer() PlayerID {
	var p PlayerID
	if s.CurrentPlayer() == Player1 {
		p = Player2
	}
	if s.CurrentPlayer() == Player2 {
		p = Player1
	}
	return p
}

// CurrentPlayerPieces returns all the Pieces which belong to the Player who is
// playing in this State.
func (s *State) CurrentPlayerPieces() []Piece {
	return s.playerPieces(s.CurrentPlayer())
}

// NextPlayerPieces returns all the Pieces which belong to the Player who will
// play in the next State.
func (s *State) NextPlayerPieces() []Piece {
	return s.playerPieces(s.NextPlayer())
}

// Player1Pieces returns all the Pieces which belong to the Player with PlayerID
// Player1.
func (s *State) Player1Pieces() []Piece {
	return s.playerPieces(Player1)
}

// Player2Pieces returns all the Pieces which belong to the Player with PlayerID
// Player2.
func (s *State) Player2Pieces() []Piece {
	return s.playerPieces(Player2)
}

// Pieces ...
func (s *State) Pieces() []Piece {
	return append(s.Player1Pieces(), s.Player2Pieces()...)
}

// CellForPiece returns the Cell the Piece is in or NoCell if the Piece is not
// in a Cell.
func (s *State) CellForPiece(p Piece) Cell {
	if c, ok := s.cells[p.ID()]; ok {
		return c
	}
	return NoCell
}

// PieceForCell returns the Piece in a Cell of NoPiece if the Cell is empty at
// the current State.
func (s *State) PieceForCell(c Cell) Piece {
	if id, ok := s.cellPieces[c]; ok {
		if p, ok := s.pieces[Player1][id]; ok {
			return p
		}
		if p, ok := s.pieces[Player2][id]; ok {
			return p
		}
	}
	return NoPiece
}

// PlayerForPiece returns the PlayerID of the Player that owns the Piece.
//
// NoPlayer is returned if no Player owns the Piece.
func (s *State) PlayerForPiece(p Piece) PlayerID {
	if _, ok := s.pieces[Player1][p.ID()]; ok {
		return Player1
	}
	if _, ok := s.pieces[Player2][p.ID()]; ok {
		return Player2
	}
	return NoPlayer
}

// Player1 of the game.
func (s *State) Player1() Player {
	return s.players[Player1]
}

// Player2 of the game.
func (s *State) Player2() Player {
	return s.players[Player2]
}

// Rules which control the game.
func (s *State) Rules() Rules {
	return s.rules
}

// Winner of the game at the State if there is one.
//
// NoPlayer is returned if there is no winner.
func (s *State) Winner() PlayerID {
	return s.winner
}

// IsGameOver returns true iff the game is over.
//
// The game is over if there is a winner or if both Players have lost.
func (s *State) IsGameOver() bool {
	return len(s.pieces[Player1]) == 0 || len(s.pieces[Player2]) == 0
}

// handleMoves applies all the Moves in the Play to the State.
//
// Destroyed Pieces aren't removed from the State and levels aren't given to the
// Piecs that destroyed those Pieces.
func handleMoves(s *State, ms Play) {
	for _, m := range ms {
		if !IsLegalMove(s, m) {
			continue
		}
		s.cells[m.Piece().ID()] = nextCell(
			s.CellForPiece(m.Piece()),
			m.Direction(),
		)
	}
}

// winner returns the PlayerID of the winner, if it exists, at the State.
//
// There is no winner if both Players have no Pieces left or both Players still
// have Pieces. If both Player's have no Pieces, the game is over and both lose.
// Otherwise, the winner is the Player that still has Pieces.
func winner(s *State) PlayerID {
	w := NoPlayer
	if len(s.pieces[Player1]) == 0 && len(s.pieces[Player2]) == 0 {
		return w
	}
	if len(s.pieces[Player1]) == 0 {
		w = Player2
	}
	if len(s.pieces[Player2]) == 0 {
		w = Player1
	}
	return w
}

// nextCell obtained by moving a cell in the Direction from the original Cell.
func nextCell(c Cell, d Direction) Cell {
	dd := map[Direction]Cell{
		North:     NewCell(0, -1),
		NorthEast: NewCell(1, -1),
		East:      NewCell(1, 0),
		SouthEast: NewCell(1, 1),
		South:     NewCell(0, 1),
		SouthWest: NewCell(-1, 1),
		West:      NewCell(-1, 0),
		NorthWest: NewCell(-1, -1),
	}[d]
	return NewCell(c.Row()+dd.Row(), c.Column()+dd.Column())
}

// clone the mutable parts of a State into a new one.
func clone(s *State) *State {
	cs := make(map[PieceID]Cell)
	ps := make(map[Cell]PieceID)
	for p, c := range s.cells {
		cs[p] = c
		ps[c] = p
	}
	p1s := make(map[PieceID]Piece)
	p2s := make(map[PieceID]Piece)
	for _, p := range s.pieces[Player1] {
		p1s[p.ID()] = p
	}
	for _, p := range s.pieces[Player2] {
		p2s[p.ID()] = p
	}
	return &State{
		currentPlayer: s.CurrentPlayer(),
		winner:        s.Winner(),
		players: map[PlayerID]Player{
			Player1: s.Player1(),
			Player2: s.Player2(),
		},
		pieces: map[PlayerID]map[PieceID]Piece{
			Player1: p1s,
			Player2: p2s,
		},
		rules:      s.Rules(),
		cells:      cs,
		cellPieces: ps,
	}
}

// playerPieces returns a copy of the Pieces for the Player with the given
// PlayerID.
func (s *State) playerPieces(pid PlayerID) []Piece {
	var ps []Piece
	for _, p := range s.pieces[pid] {
		ps = append(ps, p)
	}
	return ps
}

// handleDestroyed removes all the destroyed Pieces from the State and levels up
// the Pieces that destroyed them according to the State's Rules.
func handleDestroyed(s *State, ms Play) {
	li := s.Rules().LifeIncrease()
	di := s.Rules().DamageIncrease()
	for _, p := range s.pieces[s.NextPlayer()] {
		if p.Life() <= 0 {
			ps := s.pieces[s.CurrentPlayer()]
			for _, m := range ms {
				if nextCell(
					s.CellForPiece(m.Piece()),
					m.Direction(),
				) == s.CellForPiece(p) {
					ep := ps[m.Piece().ID()]
					ep.life += li
					ep.damage += di
					ps[m.Piece().ID()] = ep
				}
			}
			delete(s.pieces[s.NextPlayer()], p.ID())
			delete(s.cells, p.ID())
		}
	}
}

// applyMove applies the single Move to the State.
func applyMove(s *State, m Move) {
	previous := s.CellForPiece(m.Piece())
	next := nextCell(previous, m.Direction())
	if p, ok := s.cellPieces[next]; ok {
		if _, ok := s.pieces[s.CurrentPlayer()][p]; ok {
			return
		}
	}
	for _, p := range s.pieces[s.NextPlayer()] {
		if s.CellForPiece(p) == next {
			ps := s.pieces[s.NextPlayer()]
			ep := ps[p.ID()]
			ep.life -= m.Piece().Damage()
			ps[p.ID()] = ep
			return
		}
	}
	s.cells[m.Piece().ID()] = next
	s.cellPieces[next] = m.Piece().ID()
	delete(s.cellPieces, previous)
}
