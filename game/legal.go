package game

import "sync"

const (
	// TODO: Update this as a proportion to the legal moves space.
	bufferSize = 100
	// TODO: Update this to a number propertional to the available logical
	// CPUs.
	readers = 10
)

// LegalPlays returns all the legal Plays for the State's current Player.
func LegalPlays(s *State) []Play {
	bs := bucketByPiece(s)
	for i := range bs {
		bs[i] = append(bs[i], NoMove)
	}
	return legalCombinations(s, bs)
}

// LegalPlaysPipe is a pipe which outputs legal game.Plays.
func LegalPlaysPipe(s *State) chan Play {
	bs := bucketByPiece(s)
	for i := range bs {
		bs[i] = append(bs[i], NoMove)
	}
	ps := make(chan Play, bufferSize)
	go func() {
		var wg sync.WaitGroup
		cs := combinationPipe(s, bs)
		for i := 0; i < readers; i++ {
			wg.Add(1)
			go func() {
				for combo := range cs {
					if IsLegalPlay(s, combo) {
						ps <- removeMove(NoMove, combo)
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(ps)
	}()
	return ps
}

// IsLegalPlay returns true iff the Play is legal at the current State.
//
// A Play is legal iff all Moves in the play are legal after performing the
// Moves preceding them and the same Piece doesnt move more than once.
func IsLegalPlay(s *State, p Play) bool {
	used := make([]bool, 2*s.Rules().PieceCount())
	cm := newCellMap(s.Rules().BoardSize())
	for _, m := range p {
		if !IsLegalMove(s, m) || used[m.Piece().ID()-1] {
			return false
		}
		c := nextCell(s.CellForPiece(m.Piece()), m.Direction())
		pid, ok := cm.Get(c)
		if ok && s.playerForPieceID(pid) == s.CurrentPlayer() {
			return false
		}
		used[m.Piece().ID()-1] = true
		cm.Set(c, m.Piece().ID())
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
	if p := s.PieceForCell(cell); s.PlayerForPiece(p) == s.CurrentPlayer() {
		return false
	}
	return s.PlayerForPiece(m.Piece()) == s.CurrentPlayer()
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

// legalCombinations of the buckets returns all combinations with one Move from
// each bucket using all buckets which represent legal Plays.
func legalCombinations(s *State, buckets [][]Move) []Play {
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

// combinationPipe is a pipe which outputs game.Plays which aren't necessarily
// legal.
func combinationPipe(s *State, buckets [][]Move) chan Play {
	sizes := make([]int, len(buckets))
	for i, bucket := range buckets {
		sizes[i] = len(bucket)
	}
	n := combinationCount(sizes)
	indices := make([]int, len(buckets))
	combos := make(chan Play, bufferSize)
	go func() {
		for i := 0; i < n; i++ {
			combos <- combinationForIndices(buckets, indices)
			increment(indices, sizes)
		}
		close(combos)
	}()
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
