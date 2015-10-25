# go-frameworks-benchmark

Benchmark for golang web frameworks.

## Results

```
BenchmarkGinSimple-4 	 2000000	       596 ns/op	     416 B/op	       3 allocs/op
BenchmarkGinParam-4  	 2000000	       719 ns/op	     416 B/op	       3 allocs/op
BenchmarkGocraft-4   	 2000000	       752 ns/op	     280 B/op	       5 allocs/op
BenchmarkGojiSimple-4	 2000000	       723 ns/op	     416 B/op	       3 allocs/op
BenchmarkGojiParam-4 	 1000000	      1279 ns/op	     768 B/op	       6 allocs/op
BenchmarkHttp-4      	 2000000	       825 ns/op	     416 B/op	       3 allocs/op
BenchmarkKamiSimple-4	 3000000	       580 ns/op	     416 B/op	       3 allocs/op
BenchmarkKamiParam-4 	 1000000	      1077 ns/op	     576 B/op	       9 allocs/op
BenchmarkMartini-4   	  500000	      3499 ns/op	     848 B/op	      13 allocs/op
```

### env

go version: go1.5.1 darwin/amd64
MacBook Pro, 2.8 GHz Intel Core i7, 16 GB 1600 MHz DDR3

## How to run

```
go test -bench=.
```
