package game

import "testing"

// cellMapSize is the size of cellMap to use.
const cellMapSize = 11

// BenchmarkCellMapOperations benchmarks the efficiency of setting, getting, and
// removing from a cellMap.
func BenchmarkCellMapOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := newCellMap(cellMapSize)
		for j := 0; j < cellMapSize; j++ {
			for k := 0; k < cellMapSize; k++ {
				c := NewCell(j, k)
				m.Set(c, NewPiece(1, 1, 1))
				m.Get(c)
				m.Remove(c)
			}
		}
	}
}

// BenchmarkCellMapClone benchmarks the efficiency of cloning a cellMap.
func BenchmarkCellMapClone(b *testing.B) {
	m := newCellMap(cellMapSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.clone()
	}
}

// TestCellMap tests that cellMap values can properly be set, gotten, and
// removed.
func TestCellMap(t *testing.T) {
	t.Parallel()
	m := newCellMap(cellMapSize).clone()
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			m.Set(NewCell(i, j), NewPiece(1, 1, 1))
		}
	}
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if p, ok := m.Get(NewCell(i, j)); p.ID() != 1 || !ok {
				t.Errorf(
					"p=%v, ok=%v := "+
						"m.Get(NewCell(%d, %d)), "+
						"want pid=%v, ok=%v",
					p, ok,
					i, j,
					1, true,
				)
			}
		}
	}
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			m.Remove(NewCell(i, j))
		}
	}
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if _, ok := m.Get(NewCell(i, j)); ok {
				t.Errorf(
					"_, ok=%v := m.Get(NewCell(%d, %d)), "+
						"want ok=%v",
					ok,
					i, j,
					true,
				)
			}
		}
	}
	m.Get(NoCell)
	m.Set(NoCell, NoPiece)
	m.Remove(NoCell)
}
