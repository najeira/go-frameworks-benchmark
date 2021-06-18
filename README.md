# go-frameworks-benchmark

A micro benchmark for golang http routers and web frameworks.

## Results

Date: 2021-06-18

### Bench

```
BenchmarkBoneParam-8            	 1374938	       871.2 ns/op	    1248 B/op	      10 allocs/op
BenchmarkChiParam-8             	 1671458	       710.7 ns/op	     880 B/op	       7 allocs/op
BenchmarkClevergoParam-8        	 2632844	       457.8 ns/op	     464 B/op	       6 allocs/op
BenchmarkEchoParam-8            	 2471078	       480.3 ns/op	     464 B/op	       6 allocs/op
BenchmarkFireballParam-8        	  690672	      1766 ns/op	    1240 B/op	      17 allocs/op
BenchmarkGinParam-8             	 2649759	       450.3 ns/op	     480 B/op	       5 allocs/op
BenchmarkGocraftParam-8         	 1444513	       825.6 ns/op	    1056 B/op	      12 allocs/op
BenchmarkGojiParam-8            	 1892113	       625.3 ns/op	     768 B/op	       6 allocs/op
BenchmarkGorillaParam-8         	  706734	      1536 ns/op	    1745 B/op	      14 allocs/op
BenchmarkGorouterParam-8        	  283389	      4106 ns/op	    4002 B/op	      57 allocs/op
BenchmarkHttpParam-8            	 1831071	       652.7 ns/op	     848 B/op	       7 allocs/op
BenchmarkHttptreemuxParam-8     	 2313579	       517.1 ns/op	     784 B/op	       7 allocs/op
BenchmarkJulienschmidtParam-8   	 3045826	       391.5 ns/op	     464 B/op	       5 allocs/op
BenchmarkKamiParam-8            	 1584115	       750.8 ns/op	    1232 B/op	      10 allocs/op
BenchmarkLionParam-8            	 1856485	       643.7 ns/op	     880 B/op	       7 allocs/op
BenchmarkMartiniParam-8         	  268746	      4296 ns/op	    1689 B/op	      20 allocs/op
BenchmarkOzzoRoutingParam-8     	 2480985	       474.8 ns/op	     480 B/op	       7 allocs/op
BenchmarkSiestaParam-8          	  788655	      1348 ns/op	    2009 B/op	      22 allocs/op
```

### Initial Memory

```
bone: 1104 Bytes
chi: 1184 Bytes
clevergo: 39000 Bytes
echo: 4408 Bytes
fireball: 2112 Bytes
gin: 1032 Bytes
gocraft: 1616 Bytes
goji: 38680 Bytes
gorilla: 6336 Bytes
gorouter: 1008 Bytes
http: 560 Bytes
httptreemux: 1008 Bytes
julienschmidt: 632 Bytes
kami: 2184 Bytes
lion: 1816 Bytes
martini: 42296 Bytes
ozzo-routing: 7592 Bytes
siesta: 648 Bytes
```

## Environment

go version go1.16 darwin/amd64

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
- [httptreemux](https://github.com/dimfeld/httptreemux)

Fast same as net/http:
- [goji](https://github.com/zenazn/goji)
- [lion](https://github.com/celrenheit/lion)
- [chi](https://github.com/pressly/chi): regular expression matching
- [kami](https://github.com/guregu/kami)
- [gocraft](https://github.com/gocraft/web)
- [bone](https://github.com/go-zoo/bone)

2x slower net/http:
- [siesta](https://github.com/VividCortex/siesta)
- [gorilla](https://github.com/gorilla/mux)
- [fireball](https://github.com/zpatrick/fireball)

6x slower net/http:
- [gorouter](https://github.com/xujiajun/gorouter): regular expression matching
- [martini](https://github.com/go-martini/martini): no longer maintained
