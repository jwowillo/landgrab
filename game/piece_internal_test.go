package game

import "testing"

// pieceMapSize is the number of game.Pieces in a pieceMap.
const pieceMapSize = 5

// BenchmarkPieceMapOperations benchmarks the efficiency of setting, getting,
// and removing from a pieceMap.
func BenchmarkPieceMapOperations(b *testing.B) {
	p := NewPiece(1, 1, 1)
	for i := 0; i < b.N; i++ {
		m := newPieceMap(pieceMapSize)
		for j := 0; j < pieceMapSize; j++ {
			for k := 0; k < pieceMapSize*2; k++ {
				c := NewCell(j, k)
				m.Set(p, c)
				m.Get(p)
				m.Remove(p)
			}
		}
	}
}

// BenchmarkPieceMapClone benchmarks the efficiency of cloning a pieceMap.
func BenchmarkPieceMapClone(b *testing.B) {
	m := newPieceMap(pieceMapSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.clone()
	}
}

// TestPieceMap tests that pieceMap values can properly be set, gotten, and
// removed.
func TestPieceMap(t *testing.T) {
	t.Parallel()
	m := newPieceMap(pieceMapSize).clone()
	c := NewCell(3, 5)
	for i := 1; i <= pieceMapSize*2; i++ {
		m.Set(NewPiece(PieceID(i), 1, 1), c)
	}
	if _, ok := m.Get(NoPiece); ok {
		t.Errorf("_, ok=%v := m.Get(NoPiece), want ok=%v", ok, false)
	}
	m.Set(NoPiece, NoCell)
	m.Remove(NoPiece)
	for i := 1; i <= 10; i++ {
		if cc, ok := m.Get(NewPiece(PieceID(i), 1, 1)); cc != c || !ok {
			t.Errorf(
				"c=%v, ok=%v := "+
					"m.Get(NewPiece(%d, %d, %d)), "+
					"want c=%d, ok=%v",
				cc, ok,
				i, 1, 1,
				c, true,
			)
		}
	}
	for i := 1; i <= 10; i++ {
		m.Remove(NewPiece(PieceID(i), 1, 1))
	}
	for i := 1; i <= 10; i++ {
		if _, ok := m.Get(NewPiece(PieceID(i), 1, 1)); ok {
			t.Errorf(
				"_, ok=%v := m.Get(NewPiece(%d, %d, %d)), "+
					"want ok=%v",
				ok,
				i, 1, 1,
				true,
			)
		}
	}
}

// BenchmarkPieceMapIDOperations benchmarks the efficiency of setting, getting,
// and removing from a pieceIDMap.
func BenchmarkPieceIDMapOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := newPieceIDMap(pieceMapSize)
		for j := 1; j <= pieceMapSize*2; j++ {
			pid := PieceID(j)
			m.Set(pid, NewPiece(pid, 1, 1))
			m.Get(pid)
			m.Remove(pid)
		}
	}
}

// BenchmarkPieceIDMapClone benchmarks the efficiency of cloning a pieceIDMap.
func BenchmarkPieceIDMapClone(b *testing.B) {
	m := newPieceIDMap(pieceMapSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.clone()
	}
}

// TestPieceIDMap tests that pieceIDMap values can properly be set, gotten, and
// removed.
func TestPieceIDMap(t *testing.T) {
	t.Parallel()
	m := newPieceIDMap(pieceMapSize).clone()
	for i := 1; i <= pieceMapSize*2; i++ {
		pid := PieceID(i)
		m.Set(pid, NewPiece(pid, 1, 1))
	}
	for i := 1; i <= 10; i++ {
		pid := PieceID(i)
		p := NewPiece(pid, 1, 1)
		if pp, ok := m.Get(pid); p != pp || !ok {
			t.Errorf(
				"p=%v, ok=%v := m.Get(%d), want p=%d, ok=%v",
				pp, ok,
				pid,
				p, true,
			)
		}
	}
	if _, ok := m.Get(NoPieceID); ok {
		t.Errorf("_, ok=%v := m.Get(NoPieceID), want ok=%v", ok, false)
	}
	m.Set(NoPieceID, NoPiece)
	m.Remove(NoPieceID)
	for i, p := range m.Player1Pieces() {
		pid := PieceID(i + 1)
		if pid != p.ID() {
			t.Errorf("pid=%d, want %d", p.ID(), pid)
		}
	}
	for i, p := range m.Player2Pieces() {
		pid := PieceID(i + pieceMapSize + 1)
		if pid != p.ID() {
			t.Errorf("pid=%d, want %d", p.ID(), pid)
		}
	}
	for i := 1; i <= 10; i++ {
		m.Remove(PieceID(i))
	}
	for i := 1; i <= 10; i++ {
		pid := PieceID(i)
		if _, ok := m.Get(pid); ok {
			t.Errorf(
				"_, ok=%v := m.Get(%d), want ok=%v",
				ok, pid, true,
			)
		}
	}
}
