package utils

// Mutates in place, no allocations.
// Does not keep order.
func Remove[T any](s []T, i int) []T {
	j := len(s) - 1
	s[i], s[j] = s[j], s[i]
	return s[:j]
}

// Mutates in place, no allocations.
// Does not keep order.
func RemoveMany[T any](s []T, ii ...int) []T {
	for i := range ii {
		s = Remove(s, ii[i])
	}
	return s
}

// Mutates in place, no allocations.
func Filter[T any](s []T, f func(T) bool) []T {
	var n int
	for _, v := range s {
		if f(v) {
			s[n] = v
			n++
		}
	}
	return s[:n]
}

// Mutates in place, no allocations.
func Map[T any](s []T, f func(T) T) []T {
	for i, v := range s {
		s[i] = f(v)
	}
	return s
}
