package main

import (
	"net/http"
	"testing"
)

func TestOzzoRoutingParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("ozzo-routing"))
}

func BenchmarkOzzoRoutingSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("ozzo-routing"), req)
}

func BenchmarkOzzoRoutingParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher?name=gopher", nil)
	benchRequest(b, getHandler("ozzo-routing"), req)
}
