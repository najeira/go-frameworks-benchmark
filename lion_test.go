package main

import (
	"net/http"
	"testing"
)

func TestLionParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("lion"))
}

func BenchmarkLionSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("lion"), req)
}

func BenchmarkLionParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher?name=gopher", nil)
	benchRequest(b, getHandler("lion"), req)
}
