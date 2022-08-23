# UTILS

A library for common helper functions.

Most implementations are generic and no alloc.

## BENCHMARKS
```
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz

BenchmarkRemove-8        1000000000     1.137 ns/op    0 B/op    0 allocs/op
BenchmarkRemoveMany-8     872314512     1.554 ns/op    0 B/op    0 allocs/op
BenchmarkFilter-8            122851    102905 ns/op    0 B/op    0 allocs/op
```