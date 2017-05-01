package game

// LegalPlays returns all the legal Plays for the State's current Player.
func LegalPlays(s *State) []Play {
	bs := bucketByPiece(s)
	for i := range bs {
		bs[i] = append(bs[i], NoMove)
	}
	return combinations(s, bs)
}

// IsLegalPlay returns true iff the Play is legal at the current State.
//
// A Play is legal iff all Moves in the play are legal after performing the
// Moves preceding them.
func IsLegalPlay(s *State, p Play) bool {
	s = clone(s)
	for _, m := range p {
		if m != NoMove && !IsLegalMove(s, m) {
			return false
		}
		applyMove(s, m)
	}
	return true
}

// LegalMoves returns all legal Moves for the State's current Player.
func LegalMoves(s *State) []Move {
	var ms []Move
	for _, p := range s.currentPlayerPieces() {
		if p == NoPiece {
			continue
		}
		for _, d := range Directions() {
			m := NewMove(p, d)
			if IsLegalMove(s, m) {
				ms = append(ms, m)
			}
		}
	}
	return ms
}

// IsLegalMove returns true iff the Move is legal at the current State.
//
// A Move is legal iff:
//   - the Move's Piece belongs to the current Player.
//   - the Move's stays within the confines of the Board.
//   - the Move doesn't overlap with any other Board Piece's belonging to the
//   current Player.
func IsLegalMove(s *State, m Move) bool {
	cell := nextCell(s.CellForPiece(m.Piece()), m.Direction())
	r := cell.Row()
	c := cell.Column()
	size := s.Rules().BoardSize()
	if r < 0 || r >= size || c < 0 || c >= size {
		return false
	}
	if p := s.PieceForCell(cell); playerForPiece(s, p) == s.CurrentPlayer() {
		return false
	}
	return playerForPiece(s, m.Piece()) == s.CurrentPlayer()
}

// bucketByPiece buckets the list of Moves by the Piece that made the Move.
func bucketByPiece(s *State) [][]Move {
	buckets := make([][]Move, s.Rules().PieceCount())
	pc := s.Rules().PieceCount()
	for _, move := range LegalMoves(s) {
		id := (int(move.Piece().ID()) - 1) % pc
		buckets[id] = append(buckets[id], move)
	}
	var bucketed [][]Move
	for _, bucket := range buckets {
		if len(bucket) != 0 {
			bucketed = append(bucketed, bucket)
		}
	}
	return bucketed
}

// combinations of the buckets returns all combinations with one Move from each
// bucket using all buckets.
func combinations(s *State, buckets [][]Move) []Play {
	sizes := make([]int, len(buckets))
	for i, bucket := range buckets {
		sizes[i] = len(bucket)
	}
	n := combinationCount(sizes)
	combos := make([]Play, 0, n)
	indices := make([]int, len(buckets))
	for i := 0; i < n; i++ {
		p := combinationForIndices(buckets, indices)
		if IsLegalPlay(s, p) {
			combos = append(combos, removeMove(NoMove, p))
		}
		increment(indices, sizes)
	}
	return combos
}

// combinationCount returns the number of combinations of sets with sizes in the
// list of sizes.
func combinationCount(sizes []int) int {
	n := 1
	for _, size := range sizes {
		n *= size
	}
	return n
}

// combinationForIndices returns the combination corresponding to the indices
// into each bucket.
func combinationForIndices(buckets [][]Move, indices []int) []Move {
	combo := make([]Move, len(buckets))
	for i, index := range indices {
		combo[i] = buckets[i][index]
	}
	return combo
}

// increment the indices like a number was having 1 added to it with carry and
// each digit has a base equal to the corresponding element in sizes.
func increment(indices []int, sizes []int) {
	i := 0
	indices[i]++
	for i < len(indices)-1 && indices[i] >= sizes[i] {
		indices[i] = 0
		i++
		indices[i]++
	}
}

// remove the Move from the list of Moves.
func removeMove(m Move, ms []Move) []Move {
	var out []Move
	for _, cm := range ms {
		if cm != m {
			out = append(out, cm)
		}
	}
	return out
}
