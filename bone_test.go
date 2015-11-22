package main

import (
	"net/http"
	"testing"
)

func TestBoneParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("bone"))
}

func BenchmarkBoneSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("bone"), req)
}

func BenchmarkBoneParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("bone"), req)
}
