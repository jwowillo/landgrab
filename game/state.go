package game

// State encapsulates all of the game data in an immutable fashion.
type State struct {
	player1PiecesAlive, player2PiecesAlive int
	currentPlayer                          PlayerID
	rules                                  Rules
	players                                []Player
	pieces                                 []Piece
	pieceIDsToCells                        []Cell
	cellsToPieceIDs                        *cellMap
}

// NewState creates an initial game State where the game is being played by
// Player one and Player two with the given Rules.
//
// Player one is set to move first.
func NewState(r Rules, p1, p2 Player) *State {
	pieces := make([]Piece, r.PieceCount()*2)
	cs := make([]Cell, r.PieceCount()*2)
	ps := newCellMap(r.BoardSize())
	for i := 0; i < r.PieceCount(); i++ {
		p1 := newPiece(PieceID(i), r.Life(), r.Damage())
		p2 := newPiece(PieceID(i+r.PieceCount()), r.Life(), r.Damage())
		c1 := NewCell(0, i*2+1)
		c2 := NewCell(r.BoardSize()-1, i*2+1)
		pieces[p1.ID()] = p1
		pieces[p2.ID()] = p2
		cs[p1.ID()] = c1
		cs[p2.ID()] = c2
		ps.Set(c1, p1.ID())
		ps.Set(c2, p2.ID())
	}
	return &State{
		player1PiecesAlive: r.PieceCount(),
		player2PiecesAlive: r.PieceCount(),
		currentPlayer:      Player1,
		rules:              r,
		pieceIDsToCells:    cs,
		cellsToPieceIDs:    ps,
		players:            []Player{Player1: p1, Player2: p2},
		pieces:             pieces,
	}
}

// NextState returns the next State with the Play the current Player chooses.
func NextState(s *State) *State {
	s = NextStateWithPlay(s, s.players[s.CurrentPlayer()].Play(s))
	return s
}

// NextStateWithPlay returns the next State ignoring what the current Player
// would've done and instead uses the moves in the given Play.
func NextStateWithPlay(s *State, p Play) *State {
	s = clone(s)
	if s.Winner() != NoPlayer {
		return s
	}
	set := make([]bool, s.Rules().PieceCount()*2)
	for _, m := range p {
		if ok := set[m.Piece().ID()]; !ok {
			applyMove(s, m)
			set[m.Piece().ID()] = true
		}
	}
	handleDestroyed(s, p)
	s.currentPlayer = s.NextPlayer()
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
	return removePiece(s.currentPlayerPieces(), NoPiece)
}

// NextPlayerPieces returns all the Pieces which belong to the Player who will
// play in the next State.
func (s *State) NextPlayerPieces() []Piece {
	return removePiece(s.nextPlayerPieces(), NoPiece)
}

// Player1Pieces returns all the Pieces which belong to the Player with PlayerID
// Player1.
func (s *State) Player1Pieces() []Piece {
	return removePiece(s.player1Pieces(), NoPiece)
}

// Player2Pieces returns all the Pieces which belong to the Player with PlayerID
// Player2.
func (s *State) Player2Pieces() []Piece {
	return removePiece(s.player2Pieces(), NoPiece)
}

// Pieces ...
func (s *State) Pieces() []Piece {
	return append(s.Player1Pieces(), s.Player2Pieces()...)
}

// CellForPiece returns the Cell the Piece is in or NoCell if the Piece is not
// in a Cell.
func (s *State) CellForPiece(p Piece) Cell {
	if p.ID() == NoPieceID {
		return NoCell
	}
	return s.pieceIDsToCells[p.ID()]
}

// PieceForCell returns the Piece in a Cell of NoPiece if the Cell is empty at
// the current State.
func (s *State) PieceForCell(c Cell) Piece {
	if id := s.cellsToPieceIDs.Get(c); id != NoPieceID {
		return s.pieces[id]
	}
	return NoPiece
}

// PlayerForPiece returns the PlayerID of the Player that owns the Piece.
//
// NoPlayer is returned if no Player owns the Piece.
func (s *State) PlayerForPiece(p Piece) PlayerID {
	if int(p.ID()) < s.Rules().PieceCount() {
		return Player1
	}
	if int(p.ID()) >= s.Rules().PieceCount() {
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
	if s.player1PiecesAlive == 0 {
		return Player2
	}
	if s.player2PiecesAlive == 0 {
		return Player1
	}
	return NoPlayer
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
		s.pieceIDsToCells[m.Piece().ID()] = nextCell(
			s.CellForPiece(m.Piece()),
			m.Direction(),
		)
	}
}

var nextCells = []Cell{
	NewCell(0, -1),
	NewCell(1, -1),
	NewCell(1, 0),
	NewCell(1, 1),
	NewCell(0, 1),
	NewCell(-1, 1),
	NewCell(-1, 0),
	NewCell(-1, -1),
}

// nextCell obtained by moving a cell in the Direction from the original Cell.
func nextCell(c Cell, d Direction) Cell {
	dd := nextCells[d]
	return NewCell(c.Row()+dd.Row(), c.Column()+dd.Column())
}

// clone the mutable parts of a State into a new one.
func clone(s *State) *State {
	return &State{
		player1PiecesAlive: s.player1PiecesAlive,
		player2PiecesAlive: s.player2PiecesAlive,
		players:            s.players,
		currentPlayer:      s.CurrentPlayer(),
		rules:              s.Rules(),
		pieces:             append([]Piece{}, s.pieces...),
		pieceIDsToCells:    append([]Cell{}, s.pieceIDsToCells...),
		cellsToPieceIDs:    s.cellsToPieceIDs.clone(),
	}
}

func (s *State) player1Pieces() []Piece {
	return s.pieces[:s.Rules().PieceCount()]
}

func (s *State) player2Pieces() []Piece {
	return s.pieces[s.Rules().PieceCount():]
}

func (s *State) currentPlayerPieces() []Piece {
	if s.CurrentPlayer() == Player1 {
		return s.player1Pieces()
	}
	if s.CurrentPlayer() == Player2 {
		return s.player2Pieces()
	}
	return nil
}

func (s *State) nextPlayerPieces() []Piece {
	if s.CurrentPlayer() == Player1 {
		return s.player2Pieces()
	}
	if s.CurrentPlayer() == Player2 {
		return s.player1Pieces()
	}
	return nil
}

// handleDestroyed removes all the destroyed Pieces from the State and levels up
// the Pieces that destroyed them according to the State's Rules.
func handleDestroyed(s *State, ms Play) {
	li := s.Rules().LifeIncrease()
	di := s.Rules().DamageIncrease()
	for _, p := range s.nextPlayerPieces() {
		if p != NoPiece && p.Life() <= 0 {
			for _, m := range ms {
				if nextCell(
					s.CellForPiece(m.Piece()),
					m.Direction(),
				) == s.CellForPiece(p) {
					ep := s.pieces[m.Piece().ID()]
					ep.life += li
					ep.damage += di
					s.pieces[m.Piece().ID()] = ep
				}
			}
			if s.NextPlayer() == Player1 {
				s.player1PiecesAlive--
			}
			if s.NextPlayer() == Player2 {
				s.player2PiecesAlive--
			}
			s.pieces[p.ID()] = NoPiece
			s.pieceIDsToCells[p.ID()] = NoCell
		}
	}
}

// applyMove applies the single Move to the State.
func applyMove(s *State, m Move) {
	previous := s.CellForPiece(m.Piece())
	next := nextCell(previous, m.Direction())
	if id := s.cellsToPieceIDs.Get(next); id != NoPieceID {
		if p := s.pieces[id]; p != NoPiece && s.PlayerForPiece(p) == s.CurrentPlayer() {
			return
		}
	}
	for _, p := range s.nextPlayerPieces() {
		if p == NoPiece {
			continue
		}
		if s.CellForPiece(p) == next {
			ep := s.pieces[p.ID()]
			ep.life -= m.Piece().Damage()
			s.pieces[p.ID()] = ep
			return
		}
	}
	s.pieceIDsToCells[m.Piece().ID()] = next
	s.cellsToPieceIDs.Set(next, m.Piece().ID())
	s.cellsToPieceIDs.Remove(previous)
}

func removePiece(ps []Piece, r Piece) []Piece {
	var out []Piece
	for _, p := range ps {
		if p != r {
			out = append(out, p)
		}
	}
	return out
}
