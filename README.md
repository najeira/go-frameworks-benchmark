# go-frameworks-benchmark

Benchmark for golang web frameworks.

## Results

```
BenchmarkGocraft-4	 2000000	       967 ns/op	     280 B/op	       5 allocs/op
BenchmarkGoji-4   	 2000000	       933 ns/op	     416 B/op	       3 allocs/op
BenchmarkKami-4   	 2000000	       786 ns/op	     416 B/op	       3 allocs/op
BenchmarkMartini-4	  200000	      5672 ns/op	     848 B/op	      13 allocs/op
```

### env
go version: go1.5.1 darwin/amd64
MacBook Pro, 2.5 GHz Intel Core i5, 8 GB 1600 MHz DDR3

## How to run

```
go test -bench=.
```
