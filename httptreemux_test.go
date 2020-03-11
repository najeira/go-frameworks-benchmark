package main

import (
	"net/http"
	"testing"
)

func TestHttptreemuxParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("httptreemux"))
}

func BenchmarkHttptreemuxSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("httptreemux"), req)
}

func BenchmarkHttptreemuxParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher?name=gopher", nil)
	benchRequest(b, getHandler("httptreemux"), req)
}
