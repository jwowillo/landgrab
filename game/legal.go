package game

// LegalPlays returns all the legal Plays for the State's current Player.
func LegalPlays(s *State) []Play {
	bs := bucketByPiece(LegalMoves(s))
	var noMove Move
	for i, b := range bs {
		bs[i] = append(b, noMove)
	}
	var ps []Play
	for _, p := range combinations(bs) {
		p = removeMove(noMove, p)
		if IsLegalPlay(s, p) {
			ps = append(ps, p)
		}
	}
	return ps
}

// IsLegalPlay returns true iff the Play is legal at the current State.
//
// A Play is legal iff all Moves in the play are legal after performing the
// Moves preceeding them.
func IsLegalPlay(s *State, ms Play) bool {
	s = clone(s)
	for _, m := range ms {
		if !IsLegalMove(s, m) {
			return false
		}
		applyMove(s, m)
	}
	return true
}

// LegalMoves returns all legal Moves for the State's current Player.
func LegalMoves(s *State) []Move {
	var ms []Move
	for _, p := range s.pieces[s.CurrentPlayer()] {
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
//   - the Moves stays within the confines of the Board.
//   - the Move doesn't overlap with any other Board Piece's belonging to the
//   current Player.
func IsLegalMove(s *State, m Move) bool {
	isLegalPiece := false
	for _, p := range s.pieces[s.CurrentPlayer()] {
		isLegalPiece = isLegalPiece || m.Piece() == p
	}
	if !isLegalPiece {
		return false
	}
	cell := nextCell(s.CellForPiece(m.Piece()), m.Direction())
	r := cell.Row()
	c := cell.Column()
	size := 2*s.Rules().PieceCount() + 1
	if r < 0 || r >= size || c < 0 || c >= size {
		return false
	}
	for _, p := range s.pieces[s.CurrentPlayer()] {
		if m.Piece() != p && s.CellForPiece(p) == cell {
			return false
		}
	}
	return true
}

// bucketByPiece buckets the list of Moves by the Piece that made the Move.
func bucketByPiece(moves []Move) [][]Move {
	buckets := make(map[Piece][]Move)
	for _, move := range moves {
		buckets[move.Piece()] = append(buckets[move.Piece()], move)
	}
	bucketed := make([][]Move, len(buckets))
	i := 0
	for _, bucket := range buckets {
		bucketed[i] = bucket
		i++
	}
	return bucketed
}

// combinations of the buckets returns all combinations with one Move from each
// bucket using all buckets.
func combinations(buckets [][]Move) [][]Move {
	sizes := make([]int, len(buckets))
	for i, bucket := range buckets {
		sizes[i] = len(bucket)
	}
	n := combinationCount(sizes)
	combos := make([][]Move, n)
	indices := make([]int, len(buckets))
	for i := 0; i < n; i++ {
		combos[i] = combinationForIndices(buckets, indices)
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
