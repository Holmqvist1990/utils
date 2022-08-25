package utils

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

func Distance[T Number](x1, y1, x2, y2 T) T {
	return Abs(x2-x1) + Abs(y2-y1)
}

func Clamp[T Number](n, low, high T) T {
	return Max(low, Min(n, high))
}

func Lerp[T Number](start, end, t T) T {
	return (start + (end-start)*t)
}
