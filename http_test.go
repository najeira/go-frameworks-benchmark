package main

import (
	"net/http"
	"testing"
)

func TestHttpParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("http"))
}

func BenchmarkHttpSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("http"), req)
}

func BenchmarkHttpParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher?name=gopher", nil)
	benchRequest(b, getHandler("http"), req)
}
