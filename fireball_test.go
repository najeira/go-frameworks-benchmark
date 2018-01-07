package main

import (
	"net/http"
	"testing"
)

func TestFireballParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("fireball"))
}

func BenchmarkFireballSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("fireball"), req)
}

func BenchmarkFireballParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("fireball"), req)
}
