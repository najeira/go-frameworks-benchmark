package main

import (
	"net/http"
	"testing"
)

func TestGocraftParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("gocraft"))
}

func BenchmarkGocraftSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("gocraft"), req)
}

func BenchmarkGocraftParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("gocraft"), req)
}
