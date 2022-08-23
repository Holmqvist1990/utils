package utils

import "testing"

func BenchmarkRemove(b *testing.B) {
	ints := makeInts(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ints = Remove(ints, len(ints)/2)
	}
}

func BenchmarkRemoveMany(b *testing.B) {
	ints := makeInts(Max(4, b.N*2))
	b.ResetTimer()
	for i := 0; i < b.N/4; i++ {
		l := Max(0, (len(ints)/2)-4)
		ints = RemoveMany(ints, l, l+1, l+2, l+3)
	}
}

func makeInts(n int) []int {
	ints := make([]int, n)
	for i := range ints {
		ints[i] = i
	}
	return ints
}
