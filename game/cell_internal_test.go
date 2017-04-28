package game

import "testing"

// mapSize is the size of cellMap to use.
const mapSize = 11

// BenchmarkCellMapOperations benchmarks the efficiency of setting, getting, and
// removing from a cellMap.
func BenchmarkCellMapOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := newCellMap(mapSize)
		for j := 0; j < 11; j++ {
			for k := 0; k < mapSize; k++ {
				c := NewCell(j, k)
				m.Set(c, 1)
				m.Get(c)
				m.Remove(c)
			}
		}
	}
}

// BenchmarkCellMapClone benchmarks the efficiency of cloning a cellMap.
func BenchmarkCellMapClone(b *testing.B) {
	m := newCellMap(mapSize)
	for i := 0; i < b.N; i++ {
		m.clone()
	}
}

// TestCellMap tests that cellMap values can properly be set, gotten, and
// removed.
func TestCellMap(t *testing.T) {
	t.Parallel()
	m := newCellMap(mapSize)
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			m.Set(NewCell(i, j), 1)
		}
	}
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			pid, ok := m.Get(NewCell(i, j))
			if pid != 1 || !ok {
				t.Errorf(
					"pid=%d, ok=%v := "+
						"m.Get(NewCell(%d, %d)), "+
						"want pid=%d, ok=%v",
					pid, ok,
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
			_, ok := m.Get(NewCell(i, j))
			if ok {
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
}
