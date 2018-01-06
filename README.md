# go-frameworks-benchmark

A simple benchmark for golang web frameworks.

## Results

Date: 2018-01-06

### Bench

```
BenchmarkBoneParam-4      	 1000000	      1321 ns/op	    1232 B/op	      10 allocs/op
BenchmarkChiParam-4       	 2000000	       952 ns/op	     864 B/op	       7 allocs/op
BenchmarkEchoParam-4      	 2000000	       959 ns/op	     592 B/op	       9 allocs/op
BenchmarkGinParam-4       	 2000000	       640 ns/op	     480 B/op	       5 allocs/op
BenchmarkGocraftParam-4   	 1000000	      1342 ns/op	    1064 B/op	      12 allocs/op
BenchmarkGojiParam-4      	 2000000	       968 ns/op	     768 B/op	       6 allocs/op
BenchmarkGorillaParam-4   	  500000	      2401 ns/op	    1712 B/op	      14 allocs/op
BenchmarkHttpParam-4      	 1000000	      1104 ns/op	     848 B/op	       7 allocs/op
BenchmarkKamiParam-4      	 1000000	      1223 ns/op	    1216 B/op	      10 allocs/op
BenchmarkMartiniParam-4   	  200000	      6145 ns/op	    1680 B/op	      20 allocs/op
BenchmarkSiestaParam-4    	 1000000	      2209 ns/op	    2025 B/op	      23 allocs/op
```

### Memory

```
bone: 1456 Bytes
chi: 1168 Bytes
echo: 2736 Bytes
gin: 824 Bytes
gocraft: 1616 Bytes
goji: 40216 Bytes
gorilla: 4808 Bytes
http: 624 Bytes
kami: 1968 Bytes
martini: 82032 Bytes
siesta: 680 Bytes
```

## Environment

go version go1.9.1 darwin/amd64

MacBook Pro, 2.8 GHz Intel Core i7, 16 GB 1600 MHz DDR3

## How to run

```
go test -bench=. -benchmem
```
