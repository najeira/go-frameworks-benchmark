# go-frameworks-benchmark

A micro benchmark for golang http routers and web frameworks.

## Results

Date: 2022-12-03

### Bench

```
BenchmarkBoneParam-8            	 2217018	       534.8 ns/op	    1104 B/op	       9 allocs/op
BenchmarkChiParam-8             	 3191725	       377.7 ns/op	     736 B/op	       6 allocs/op
BenchmarkClevergoParam-8        	 4589001	       266.2 ns/op	     464 B/op	       6 allocs/op
BenchmarkEchoParam-8            	 4291570	       293.3 ns/op	     464 B/op	       6 allocs/op
BenchmarkFireballParam-8        	  959980	      1255 ns/op	    1240 B/op	      17 allocs/op
BenchmarkGinParam-8             	 4614988	       259.5 ns/op	     480 B/op	       5 allocs/op
BenchmarkGocraftParam-8         	 2232187	       530.7 ns/op	    1056 B/op	      12 allocs/op
BenchmarkGojiParam-8            	 3029848	       396.2 ns/op	     768 B/op	       6 allocs/op
BenchmarkGorillaParam-8         	 1405196	       856.7 ns/op	    1457 B/op	      12 allocs/op
BenchmarkGorouterParam-8        	  439477	      2729 ns/op	    3957 B/op	      56 allocs/op
BenchmarkHttpParam-8            	 2983761	       400.8 ns/op	     848 B/op	       7 allocs/op
BenchmarkHttptreemuxParam-8     	 3493615	       343.5 ns/op	     784 B/op	       7 allocs/op
BenchmarkJulienschmidtParam-8   	 5174398	       237.3 ns/op	     464 B/op	       5 allocs/op
BenchmarkKamiParam-8            	 2554264	       488.9 ns/op	    1088 B/op	       9 allocs/op
BenchmarkLionParam-8            	 3220908	       363.8 ns/op	     736 B/op	       6 allocs/op
BenchmarkMartiniParam-8         	  481544	      2461 ns/op	    1717 B/op	      23 allocs/op
BenchmarkOzzoRoutingParam-8     	 4360780	       285.6 ns/op	     480 B/op	       7 allocs/op
BenchmarkSiestaParam-8          	 1359814	       865.5 ns/op	    1986 B/op	      21 allocs/op
```

### Initial Memory

```
bone: 2440 Bytes
chi: 3424 Bytes
clevergo: 38984 Bytes
echo: 4608 Bytes
fireball: 2200 Bytes
gin: 1048 Bytes
gocraft: 1616 Bytes
goji: 38680 Bytes
gorilla: 6320 Bytes
gorouter: 912 Bytes
http: 856 Bytes
httptreemux: 1008 Bytes
julienschmidt: 536 Bytes
kami: 1992 Bytes
lion: 1832 Bytes
martini: 42488 Bytes
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
