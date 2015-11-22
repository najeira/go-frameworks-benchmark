package main

import (
	"net/http"
	"testing"
)

func TestSiestaParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("siesta"))
}

func BenchmarkSiestaSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("siesta"), req)
}

func BenchmarkSiestaParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("siesta"), req)
}
