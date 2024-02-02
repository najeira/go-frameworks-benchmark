# go-frameworks-benchmark

A micro benchmark for golang http routers and web frameworks.

## Results

Date: 2024-02-02

### Bench

```
BenchmarkBoneParam-8            	 2365204	       496.2 ns/op	    1104 B/op	       9 allocs/op
BenchmarkChiParam-8             	 3401697	       355.3 ns/op	     736 B/op	       6 allocs/op
BenchmarkClevergoParam-8        	 4677952	       258.9 ns/op	     464 B/op	       6 allocs/op
BenchmarkEchoParam-8            	 4262938	       294.9 ns/op	     464 B/op	       6 allocs/op
BenchmarkFireballParam-8        	  973047	      1234 ns/op	    1240 B/op	      17 allocs/op
BenchmarkGinParam-8             	 4233492	       278.7 ns/op	     480 B/op	       5 allocs/op
BenchmarkGocraftParam-8         	 2217714	       527.6 ns/op	    1056 B/op	      12 allocs/op
BenchmarkGojiParam-8            	 3061669	       414.6 ns/op	     768 B/op	       6 allocs/op
BenchmarkGorillaParam-8         	 1464541	       847.1 ns/op	    1457 B/op	      12 allocs/op
BenchmarkGorouterParam-8        	  461496	      2661 ns/op	    3954 B/op	      56 allocs/op
BenchmarkHttpParam-8            	 2778646	       403.4 ns/op	     848 B/op	       7 allocs/op
BenchmarkHttptreemuxParam-8     	 3266467	       378.3 ns/op	     784 B/op	       7 allocs/op
BenchmarkJulienschmidtParam-8   	 4764955	       236.5 ns/op	     464 B/op	       5 allocs/op
BenchmarkKamiParam-8            	 2673991	       475.4 ns/op	    1088 B/op	       9 allocs/op
BenchmarkLionParam-8            	 3400897	       357.7 ns/op	     736 B/op	       6 allocs/op
BenchmarkMartiniParam-8         	  484983	      2376 ns/op	    1721 B/op	      23 allocs/op
BenchmarkOzzoRoutingParam-8     	 3598252	       283.1 ns/op	     480 B/op	       7 allocs/op
BenchmarkSiestaParam-8          	 1348377	       887.0 ns/op	    2002 B/op	      21 allocs/op
```

### Initial Memory

```
bone: 2552 Bytes
chi: 3424 Bytes
clevergo: 38984 Bytes
echo: 4584 Bytes
fireball: 2080 Bytes
gin: 1048 Bytes
gocraft: 1616 Bytes
goji: 38680 Bytes
gorilla: 6320 Bytes
gorouter: 1008 Bytes
http: 560 Bytes
httptreemux: 1008 Bytes
julienschmidt: 1208 Bytes
kami: 1608 Bytes
lion: 1832 Bytes
martini: 44616 Bytes
ozzo-routing: 7592 Bytes
siesta: 648 Bytes
```

## Environment

go version go1.21 darwin/amd64

Mac mini (M1, 2020)
Apple M1 8 cores
16 GB

## How to run

```
go test -bench=. -benchmem
```

## Conclusion

Faster than net/http:
- [httprouter](https://github.com/julienschmidt/httprouter)
- [gin](https://github.com/gin-gonic/gin)
- [clevergo](https://github.com/clevergo/clevergo)
- [ozzo-routing](https://github.com/go-ozzo/ozzo-routing): regular expression matching
- [echo](https://github.com/labstack/echo)

Fast same as net/http:
- [httptreemux](https://github.com/dimfeld/httptreemux)
- [lion](https://github.com/celrenheit/lion)
- [chi](https://github.com/pressly/chi): regular expression matching
- [goji](https://github.com/zenazn/goji)
- [kami](https://github.com/guregu/kami)
- [gocraft](https://github.com/gocraft/web)
- [bone](https://github.com/go-zoo/bone)

2x-3x slower net/http:
- [gorilla](https://github.com/gorilla/mux)
- [siesta](https://github.com/VividCortex/siesta)
- [fireball](https://github.com/zpatrick/fireball)

6x slower net/http:
- [martini](https://github.com/go-martini/martini): no longer maintained
- [gorouter](https://github.com/xujiajun/gorouter): regular expression matching
