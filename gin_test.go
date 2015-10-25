package main

import (
	"net/http"
	"testing"
)

func BenchmarkGin(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("gin"), req)
}
