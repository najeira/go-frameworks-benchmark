# go-frameworks-benchmark

A micro benchmark for golang http routers and web frameworks.

## Results

Date: 2020-03-11

### Bench

```
BenchmarkBoneSimple-4            2232099               457 ns/op             416 B/op          3 allocs/op
BenchmarkChiSimple-4             1000000              1115 ns/op             848 B/op          6 allocs/op
BenchmarkEchoSimple-4            2693743               738 ns/op             432 B/op          4 allocs/op
BenchmarkFireballSimple-4         422511              3892 ns/op            1200 B/op         15 allocs/op
BenchmarkGinSimple-4             2071533               833 ns/op             464 B/op          4 allocs/op
BenchmarkGocraftSimple-4         1000000              1493 ns/op             696 B/op          8 allocs/op
BenchmarkGojiSimple-4            1437841               709 ns/op             416 B/op          3 allocs/op
BenchmarkGorillaSimple-4          671144              1659 ns/op            1392 B/op         12 allocs/op
BenchmarkHttpSimple-4            1964276               666 ns/op             416 B/op          3 allocs/op
BenchmarkKamiSimple-4            2572772               412 ns/op             416 B/op          3 allocs/op
BenchmarkMartiniSimple-4          229540              5814 ns/op            1309 B/op         16 allocs/op
BenchmarkSiestaSimple-4          1424252               756 ns/op             865 B/op          9 allocs/op
```

```
BenchmarkBoneParam-4              807661              2007 ns/op            1232 B/op         10 allocs/op
BenchmarkChiParam-4              1365415               871 ns/op             864 B/op          7 allocs/op
BenchmarkEchoParam-4             1326015               969 ns/op             464 B/op          6 allocs/op
BenchmarkFireballParam-4          336466              3008 ns/op            1240 B/op         17 allocs/op
BenchmarkGinParam-4              1550900               823 ns/op             480 B/op          5 allocs/op
BenchmarkGocraftParam-4           576402              2051 ns/op            1064 B/op         12 allocs/op
BenchmarkGojiParam-4             1000000              1543 ns/op             768 B/op          6 allocs/op
BenchmarkGorillaParam-4           441978              3080 ns/op            1712 B/op         14 allocs/op
BenchmarkHttpParam-4             1098367              1046 ns/op             848 B/op          7 allocs/op
BenchmarkKamiParam-4             1000000              1333 ns/op            1216 B/op         10 allocs/op
BenchmarkMartiniParam-4           205346              5451 ns/op            1696 B/op         20 allocs/op
BenchmarkSiestaParam-4            656433              2081 ns/op            2017 B/op         22 allocs/op
```

### Memory

```
bone: 504 Bytes
chi: 1184 Bytes
echo: 3704 Bytes
fireball: 1808 Bytes
gin: 984 Bytes
gocraft: 1616 Bytes
goji: 38168 Bytes
gorilla: 5536 Bytes
http: 848 Bytes
kami: 1616 Bytes
martini: 41312 Bytes
siesta: 680 Bytes
```

## Environment

go version go1.13 darwin/amd64

MacBook Pro (13-inch, 2016, Four Thunderbolt 3 Ports)
3.3 GHz Intel Core i7
16 GB 2133 MHz LPDDR3

## How to run

```
go test -bench=. -benchmem
```
