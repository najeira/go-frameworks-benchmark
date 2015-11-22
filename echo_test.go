package main

import (
	"net/http"
	"testing"
)

func TestEchoParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("echo"))
}

func BenchmarkEchoSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("echo"), req)
}

func BenchmarkEchoParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("echo"), req)
}
