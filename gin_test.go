package main

import (
	"net/http"
	"testing"
)

func TestGinParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("gin"))
}

func BenchmarkGinSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("gin"), req)
}

func BenchmarkGinParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("gin"), req)
}
