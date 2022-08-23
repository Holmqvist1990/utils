# 2D-UTILS

A no-alloc, low time complexity, generics based utilities library for 2D games.

* Math (vectors, clamping, lerping...)
* Slices (remove, filter in place...)
* Tile (1D, 2D, safe indexing...)

## BENCHMARKS for slice operations
```
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz

BenchmarkRemove-8        1000000000     1.137 ns/op    0 B/op    0 allocs/op
BenchmarkRemoveMany-8     872314512     1.554 ns/op    0 B/op    0 allocs/op
BenchmarkFilter-8            114692    101617 ns/op    0 B/op    0 allocs/op
BenchmarkInlinedFilter-8     334354    107903 ns/op    0 B/op    0 allocs/op
```