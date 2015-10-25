# go-frameworks-benchmark

Benchmark for golang web frameworks.

## Results

```
BenchmarkGin-4       	 2000000	       600 ns/op	     416 B/op	       3 allocs/op
BenchmarkGocraft-4   	 2000000	       746 ns/op	     280 B/op	       5 allocs/op
BenchmarkGojiSimple-4	 2000000	       721 ns/op	     416 B/op	       3 allocs/op
BenchmarkGojiParam-4 	 1000000	      1253 ns/op	     768 B/op	       6 allocs/op
BenchmarkHttp-4      	 2000000	       817 ns/op	     416 B/op	       3 allocs/op
BenchmarkKami-4      	 1000000	      1153 ns/op	     576 B/op	       9 allocs/op
BenchmarkMartini-4   	  500000	      3550 ns/op	     848 B/op	      13 allocs/op
```

### env

go version: go1.5.1 darwin/amd64
MacBook Pro, 2.8 GHz Intel Core i7, 16 GB 1600 MHz DDR3

## How to run

```
go test -bench=.
```
