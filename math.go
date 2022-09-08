package utils

import "math"

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}

	return n
}

func Magnitude[T Number](x, y T) T {
	_x, _y := float64(x), float64(y)
	return T(math.Sqrt(_x*_x + _y*_y))
}

func Normalize[T Float](x, y T) (T, T) {
	if x == 0 || y == 0 {
		return x, y
	}

	m := Magnitude(x, y)
	return x / m, y / m
}

func Distance[T Number](x1, y1, x2, y2 T) T {
	return Abs(x2-x1) + Abs(y2-y1)
}

func Clamp[T Number](n, low, high T) T {
	return Max(low, Min(n, high))
}

// Minimum absolute value or zero.
func MinOrZero[T Number](v, min T) T {
	if Abs(v) < min {
		return 0
	}

	return v
}

func Lerp[T Number](start, end, t T) T {
	return (start + (end-start)*t)
}
