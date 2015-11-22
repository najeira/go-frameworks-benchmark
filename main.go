package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
	"strings"
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

type simpleResponseWriter struct {
	code   int
	body   bytes.Buffer
	header http.Header
}

func (m *simpleResponseWriter) Header() http.Header {
	return m.header
}

func (m *simpleResponseWriter) Write(p []byte) (int, error) {
	return m.body.Write(p)
}

func (m *simpleResponseWriter) WriteHeader(code int) {
	m.code = code
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
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(&w, r)

		// clear caches
		r.Form = nil
		r.PostForm = nil
		r.MultipartForm = nil
	}
}

func sendRequest(router http.Handler, r *http.Request) (int, []byte, http.Header) {
	w := simpleResponseWriter{header: http.Header{}}
	router.ServeHTTP(&w, r)
	return w.code, w.body.Bytes(), w.header
}

func testRequestWithPathParam(t *testing.T, handler http.Handler) {
	req, _ := http.NewRequest("GET", "/gopher?name=gopher", nil)
	c, b, h := sendRequest(handler, req)
	if c != 0 && c != 200 {
		t.Errorf("invalid status code: %d", c)
	}
	if !strings.Contains(string(b), "gopher") {
		t.Errorf("invalid body: %s", string(b))
	}
	if h == nil {
		t.Errorf("invalid header")
	} else {
		ct := h["Content-Type"]
		if len(ct) <= 0 {
			t.Errorf("invalid header")
		} else {
			if !strings.Contains(ct[0], "text/plain") {
				t.Errorf("invalid header")
			}
		}
	}
}

func main() {
	fmt.Println("run: go test -bench=.")
}
