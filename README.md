# go-frameworks-benchmark

Benchmark for golang web frameworks.

## Results

Date: 2017-05-08

### Bench

```
BenchmarkBoneParam-4             1000000              1396 ns/op            1104 B/op          9 allocs/op
BenchmarkChiParam-4              2000000               865 ns/op             736 B/op          6 allocs/op
BenchmarkEchoParam-4             2000000               864 ns/op             496 B/op          7 allocs/op
BenchmarkGinParam-4              3000000               575 ns/op             416 B/op          3 allocs/op
BenchmarkGocraftParam-4          1000000              1375 ns/op            1064 B/op         12 allocs/op
BenchmarkGojiParam-4             1000000              1090 ns/op             768 B/op          6 allocs/op
BenchmarkGorillaParam-4           500000              2433 ns/op            1496 B/op         16 allocs/op
BenchmarkHttpParam-4             1000000              1351 ns/op             848 B/op          7 allocs/op
BenchmarkKamiParam-4             1000000              1504 ns/op            1088 B/op          9 allocs/op
BenchmarkMartiniParam-4           200000              6744 ns/op            1680 B/op         20 allocs/op
BenchmarkSiestaParam-4            500000              2326 ns/op            2041 B/op         24 allocs/op
```

### Memory

```
bone: 3584 Bytes
chi: 1136 Bytes
echo: 2944 Bytes
gin: 712 Bytes
gocraft: 1616 Bytes
goji: 40216 Bytes
gorilla: 5464 Bytes
http: 624 Bytes
kami: 1936 Bytes
martini: 82032 Bytes
siesta: 680 Bytes
```

## env

go version go1.8 darwin/amd64

MacBook Pro, 2.8 GHz Intel Core i7, 16 GB 1600 MHz DDR3

## How to run

```
go test -bench=.
```
