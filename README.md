# go-frameworks-benchmark

Benchmark for golang web frameworks.

## Results

### Bench

```
BenchmarkBoneParam-4   	 1000000	      1463 ns/op	     800 B/op	       7 allocs/op
BenchmarkEchoParam-4   	 2000000	       879 ns/op	     544 B/op	       5 allocs/op
BenchmarkGinParam-4    	 2000000	       734 ns/op	     416 B/op	       3 allocs/op
BenchmarkGocraftParam-4	 1000000	      1915 ns/op	    1064 B/op	      12 allocs/op
BenchmarkGojiParam-4   	 1000000	      1299 ns/op	     768 B/op	       6 allocs/op
BenchmarkGorillaParam-4	  500000	      3944 ns/op	    1264 B/op	      15 allocs/op
BenchmarkHttpParam-4   	 1000000	      1854 ns/op	     848 B/op	       7 allocs/op
BenchmarkKamiParam-4   	 1000000	      1237 ns/op	     576 B/op	       9 allocs/op
BenchmarkMartiniParam-4	  200000	      5823 ns/op	    1729 B/op	      22 allocs/op
BenchmarkSiestaParam-4 	  500000	      3215 ns/op	    2056 B/op	      24 allocs/op
```

### Memory

```
bone: 3360 Bytes
echo: 800 Bytes
gin: 712 Bytes
gocraft: 1600 Bytes
goji: 40232 Bytes
gorilla: 5512 Bytes
http: 624 Bytes
kami: 1160 Bytes
martini: 82096 Bytes
siesta: 680 Bytes
```

## env

go version: go1.5.1 darwin/amd64

MacBook Pro, 2.8 GHz Intel Core i7, 16 GB 1600 MHz DDR3

## How to run

```
go test -bench=.
```
