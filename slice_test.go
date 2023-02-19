package utils

import (
	"fmt"
	"testing"
)

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

// As fast as if written by hand.
func BenchmarkFilter(b *testing.B) {
	ints := makeInts(b.N)
	f := func(n int) bool {
		return n%2 == 0
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ints = Filter(ints, f)
	}
}

// To compare with benchmark above.
func BenchmarkInlinedFilter(b *testing.B) {
	ints := makeInts(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var n int
		for _, v := range ints {
			if v%2 == 0 {
				ints[n] = v
				n++
			}
		}
		ints = ints[:n]
	}
}

func TestMap(t *testing.T) {
	ints := makeInts(8)
	double := func(i int) int {
		return i * 2
	}
	doubles := Map(makeInts(8), double)
	for i := range doubles {
		if double(ints[i]) != doubles[i] {
			t.Fatalf("expected %v got %v",
				double(ints[i]), doubles[i])
		}
	}
}

func TestMapTo(t *testing.T) {
	ints := makeInts(8)
	itoa := func(i int) string {
		return fmt.Sprintf("%d", i)
	}
	strs := MapTo(ints, itoa)
	for i := range strs {
		if itoa(ints[i]) != strs[i] {
			t.Fatalf("expected %v got %v", itoa(ints[i]), strs[i])
		}
	}
}

func makeInts(n int) []int {
	ints := make([]int, n)
	for i := range ints {
		ints[i] = i
	}
	return ints
}
