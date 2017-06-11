package game

// State encapsulates all of the game data in an immutable fashion.
type State struct {
	player1PiecesAlive, player2PiecesAlive int
	currentPlayer                          PlayerID
	rules                                  Rules
	players                                []Player
	pieces                                 pieceIDMap
	piecesToCells                          pieceMap
	cellsToPieceIDs                        cellMap
}

// NewState creates an initial game State where the game is being played by
// Player one and Player two with the given Rules.
//
// Player one is set to move first.
func NewState(r Rules, p1, p2 Player) *State {
	pieces := newPieceIDMap(r.PieceCount())
	cs := newPieceMap(r.PieceCount())
	ps := newCellMap(r.BoardSize())
	for i := 0; i < r.PieceCount(); i++ {
		p1id := PieceID(i + 1)
		p2id := PieceID(i + r.PieceCount() + 1)
		p1 := NewPiece(p1id, r.Life(), r.Damage())
		p2 := NewPiece(p2id, r.Life(), r.Damage())
		c1 := NewCell(0, i*2+1)
		c2 := NewCell(r.BoardSize()-1, i*2+1)
		pieces.Set(p1id, p1)
		pieces.Set(p2id, p2)
		cs.Set(p1, c1)
		cs.Set(p2, c2)
		ps.Set(c1, p1id)
		ps.Set(c2, p2id)
	}
	return &State{
		player1PiecesAlive: r.PieceCount(),
		player2PiecesAlive: r.PieceCount(),
		currentPlayer:      Player1,
		rules:              r,
		piecesToCells:      cs,
		cellsToPieceIDs:    ps,
		players:            []Player{Player1: p1, Player2: p2},
		pieces:             pieces,
	}
}

// NewStateFromInfo creates a State using info from a game already in progress.
func NewStateFromInfo(
	rules Rules,
	currentPlayer PlayerID,
	p1 Player, p2 Player,
	pieces map[Cell]Piece,
) *State {
	ps := newPieceIDMap(rules.PieceCount())
	cs := newPieceMap(rules.PieceCount())
	cm := newCellMap(rules.BoardSize())
	var p1Alive, p2Alive int
	for c, p := range pieces {
		if p.ID() != NoPieceID {
			pid := int(p.ID())
			if pid <= rules.PieceCount() {
				p1Alive++
			}
			if pid > rules.PieceCount() {
				p2Alive++
			}
		}
		ps.Set(p.ID(), p)
		cs.Set(p, c)
		cm.Set(c, p.ID())
	}
	return &State{
		player1PiecesAlive: p1Alive,
		player2PiecesAlive: p2Alive,
		currentPlayer:      currentPlayer,
		rules:              rules,
		players:            []Player{Player1: p1, Player2: p2},
		pieces:             ps,
		piecesToCells:      cs,
		cellsToPieceIDs:    cm,
	}
}

// NextState returns the next State with the Play the current Player chooses.
func NextState(s *State) *State {
	return NextStateWithPlay(s, s.players[s.CurrentPlayer()].Play(s))
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
		if ok := set[m.Piece().ID()-1]; !ok {
			applyMove(s, m)
			set[m.Piece().ID()-1] = true
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
	switch s.CurrentPlayer() {
	case Player1:
		return Player2
	case Player2:
		return Player1
	default:
		return NoPlayer
	}
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

// Pieces returns the Pieces for Player one and Player two.
func (s *State) Pieces() []Piece {
	return append(s.Player1Pieces(), s.Player2Pieces()...)
}

// CellForPiece returns the Cell the Piece is in or NoCell if the Piece is not
// in a Cell.
func (s *State) CellForPiece(p Piece) Cell {
	if c, ok := s.piecesToCells.Get(p); ok {
		return c
	}
	return NoCell
}

// PieceForCell returns the Piece in a Cell of NoPiece if the Cell is empty at
// the current State.
func (s *State) PieceForCell(c Cell) Piece {
	if pid, ok := s.cellsToPieceIDs.Get(c); ok {
		if p, ok := s.pieces.Get(pid); ok {
			return p
		}
		return NoPiece
	}
	return NoPiece
}

// PlayerForPiece returns the PlayerID of the Player that owns the Piece.
//
// NoPlayer is returned if no Player owns the Piece.
func (s *State) PlayerForPiece(p Piece) PlayerID {
	return s.playerForPieceID(p.ID())
}

func (s *State) playerForPieceID(pid PieceID) PlayerID {
	id := int(pid)
	pc := s.Rules().PieceCount()
	if id > 0 && id <= pc {
		return Player1
	}
	if id > pc && id <= 2*pc {
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
		s.piecesToCells.Set(m.Piece(), nextCell(
			s.CellForPiece(m.Piece()),
			m.Direction(),
		))
	}
}

// nextCells is a hardcoded list of direction Cells for use in determining
// possible next Cells.
var nextCells = []Cell{
	NewCell(0, 0),
	NewCell(-1, 0),
	NewCell(-1, 1),
	NewCell(0, 1),
	NewCell(1, 1),
	NewCell(1, 0),
	NewCell(1, -1),
	NewCell(0, -1),
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
		pieces:             s.pieces.clone(),
		piecesToCells:      s.piecesToCells.clone(),
		cellsToPieceIDs:    s.cellsToPieceIDs.clone(),
	}
}

// player1Pieces returns a non-copied list of Player one's Pieces.
func (s *State) player1Pieces() []Piece {
	return s.pieces.Player1Pieces()
}

// player2Pieces returns a non-copied list of Player two's Pieces.
func (s *State) player2Pieces() []Piece {
	return s.pieces.Player2Pieces()
}

// currentPlayerPieces returns a non-copied list of the current Player's Pieces.
func (s *State) currentPlayerPieces() []Piece {
	switch s.CurrentPlayer() {
	case Player1:
		return s.player1Pieces()
	case Player2:
		return s.player2Pieces()
	default:
		return nil
	}
}

// nextPlayerPieces returns a non-copied list of the next Player's Pieces.
func (s *State) nextPlayerPieces() []Piece {
	switch s.CurrentPlayer() {
	case Player1:
		return s.player2Pieces()
	case Player2:
		return s.player1Pieces()
	default:
		return nil
	}
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
				) != s.CellForPiece(p) {
					continue
				}
				s.pieces.Set(m.Piece().ID(), NewPiece(
					m.Piece().ID(),
					m.Piece().Life()+li,
					m.Piece().Damage()+di,
				))
			}
			switch s.NextPlayer() {
			case Player1:
				s.player1PiecesAlive--
			case Player2:
				s.player2PiecesAlive--
			}
			s.pieces.Set(p.ID(), NoPiece)
			s.piecesToCells.Set(p, NoCell)
		}
	}
}

// applyMove applies the single Move to the State.
func applyMove(s *State, m Move) {
	previous := s.CellForPiece(m.Piece())
	next := nextCell(previous, m.Direction())
	if pid, ok := s.cellsToPieceIDs.Get(next); ok {
		if s.playerForPieceID(pid) == s.CurrentPlayer() {
			return
		}
	}
	if p := s.PieceForCell(next); s.PlayerForPiece(p) == s.NextPlayer() {
		s.pieces.Set(p.ID(), NewPiece(
			p.ID(),
			p.Life()-m.Piece().Damage(),
			p.Damage(),
		))
		return
	}
	s.piecesToCells.Set(m.Piece(), next)
	s.cellsToPieceIDs.Set(next, m.Piece().ID())
	s.cellsToPieceIDs.Remove(previous)
}

// removePiece occurances in list of Pieces.
func removePiece(ps []Piece, r Piece) []Piece {
	var out []Piece
	for _, p := range ps {
		if p != r {
			out = append(out, p)
		}
	}
	return out
}
