package main

import (
	"fmt"
	"net/http"
	"runtime"
	"testing"
)

var (
	httpHandlers map[string]http.Handler
)

func registerHandler(name string, handler http.Handler) {
	if httpHandlers == nil {
		httpHandlers = make(map[string]http.Handler)
	} else if _, ok := httpHandlers[name]; ok {
		panic("already registered")
	}
	httpHandlers[name] = handler
}

func getHandler(name string) http.Handler {
	if httpHandlers == nil {
		return nil
	}
	handler, _ := httpHandlers[name]
	return handler
}

type mockResponseWriter struct {
}

func (m *mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteHeader(code int) {

}


}

func calcMem(name string, load func()) {
	m := new(runtime.MemStats)

	// before
	runtime.GC()
	runtime.ReadMemStats(m)
	before := m.HeapAlloc

	load()

	// after
	runtime.GC()
	runtime.ReadMemStats(m)
	after := m.HeapAlloc
	println("   "+name+":", after-before, "Bytes")
}

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := mockResponseWriter{}
	u := r.URL
	rq := u.RawQuery
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.RawQuery = rq
		router.ServeHTTP(&w, r)
	}
}

func main() {
	fmt.Println("run: go test -bench=.")
}
