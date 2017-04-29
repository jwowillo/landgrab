package game

import "testing"

func BenchmarkNextCell(b *testing.B) {
	c := NewCell(0, 0)
	ds := Directions()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, d := range ds {
			nextCell(c, d)
		}
	}
}
