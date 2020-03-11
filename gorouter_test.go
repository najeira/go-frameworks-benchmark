package main

import (
	"net/http"
	"testing"
)

func TestGorouterParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("gorouter"))
}

func BenchmarkGorouterSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("gorouter"), req)
}

func BenchmarkGorouterParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher?name=gopher", nil)
	benchRequest(b, getHandler("gorouter"), req)
}
