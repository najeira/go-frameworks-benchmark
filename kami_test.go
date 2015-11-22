package main

import (
	"net/http"
	"testing"
)

func TestKami(t *testing.T) {
	testRequestWithPathParam(t, getHandler("kami"))
}

func BenchmarkKamiSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("kami"), req)
}

func BenchmarkKamiParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("kami"), req)
}
