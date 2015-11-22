package main

import (
	"net/http"
	"testing"
)

func TestMartiniParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("martini"))
}

func BenchmarkMartiniSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("martini"), req)
}

func BenchmarkMartiniParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("martini"), req)
}
