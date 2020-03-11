package main

import (
	"net/http"
	"testing"
)

func TestClevergoParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("clevergo"))
}

func BenchmarkClevergoSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("clevergo"), req)
}

func BenchmarkClevergoParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher?name=gopher", nil)
	benchRequest(b, getHandler("clevergo"), req)
}
