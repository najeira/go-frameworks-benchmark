# go-frameworks-benchmark

A micro benchmark for golang http routers and web frameworks.

## Results

Date: 2020-03-11

### Bench

```
BenchmarkBoneParam-4                      289956              4812 ns/op            1232 B/op         10 allocs/op
BenchmarkChiParam-4                      1395741              1251 ns/op             864 B/op          7 allocs/op
BenchmarkClevergoParam-4                 2012955               579 ns/op             464 B/op          6 allocs/op
BenchmarkEchoParam-4                     1511348               663 ns/op             464 B/op          6 allocs/op
BenchmarkFireballParam-4                  400944              2854 ns/op            1240 B/op         17 allocs/op
BenchmarkGinParam-4                      1910553               696 ns/op             480 B/op          5 allocs/op
BenchmarkGocraftParam-4                  1013068              1243 ns/op            1064 B/op         12 allocs/op
BenchmarkGojiParam-4                     1478815               843 ns/op             768 B/op          6 allocs/op
BenchmarkGorillaParam-4                   616470              2005 ns/op            1712 B/op         14 allocs/op
BenchmarkGorouterParam-4                  191655              5906 ns/op            3970 B/op         57 allocs/op
BenchmarkHttpParam-4                     1227997               963 ns/op             848 B/op          7 allocs/op
BenchmarkHttptreemuxParam-4              1000000              1474 ns/op             784 B/op          7 allocs/op
BenchmarkJulienschmidtParam-4            2395546               484 ns/op             464 B/op          5 allocs/op
BenchmarkKamiParam-4                     1117251              1303 ns/op            1216 B/op         10 allocs/op
BenchmarkLionParam-4                     1378180              1029 ns/op             864 B/op          7 allocs/op
BenchmarkMartiniParam-4                   237154              5157 ns/op            1697 B/op         20 allocs/op
BenchmarkOzzoRoutingParam-4              1883518               604 ns/op             480 B/op          7 allocs/op
BenchmarkSiestaParam-4                    602625              1864 ns/op            2017 B/op         22 allocs/op
```

### Initial Memory

```
bone: 488 Bytes
chi: 1184 Bytes
clevergo: 38512 Bytes
echo: 3784 Bytes
fireball: 1904 Bytes
gin: 984 Bytes
gocraft: 1616 Bytes
goji: 38168 Bytes
gorilla: 5824 Bytes
gorouter: 912 Bytes
http: 560 Bytes
httptreemux: 992 Bytes
julienschmidt: 824 Bytes
kami: 1904 Bytes
lion: 1816 Bytes
martini: 41312 Bytes
ozzo-routing: 7608 Bytes
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

## Conclusion

Faster than net/http:
- [httprouter](https://github.com/julienschmidt/httprouter)
- [clevergo](https://github.com/clevergo/clevergo)
- [ozzo-routing](https://github.com/go-ozzo/ozzo-routing): regular expression matching
- [echo](https://github.com/labstack/echo)
- [gin](https://github.com/gin-gonic/gin)
- [goji](https://github.com/zenazn/goji)

Fast same as net/http:
- [lion](https://github.com/celrenheit/lion)
- [gocraft](https://github.com/gocraft/web)
- [chi](https://github.com/pressly/chi): regular expression matching
- [kami](https://github.com/guregu/kami)
- [httptreemux](https://github.com/dimfeld/httptreemux)

2x slower net/http:
- [siesta](https://github.com/VividCortex/siesta)
- [gorilla](https://github.com/gorilla/mux)
- [fireball](https://github.com/zpatrick/fireball)
- [bone](https://github.com/go-zoo/bone)
- [martini](https://github.com/go-martini/martini): no longer maintained
- [gorouter](https://github.com/xujiajun/gorouter): regular expression matching
