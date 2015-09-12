package main

import (
	"fmt"
	"net/http"

	goji "github.com/zenazn/goji/web"
)

func init() {
	calcMem("goji", initGoji)
}

func initGoji() {
	h := goji.New()
	h.Get("/", func(c goji.C, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, World")
	})
	registerHandler("goji", h)
}
