package main

import (
	"fmt"
	"net/http"

	"github.com/go-zoo/bone"
)

func init() {
	calcMem("bone", initBone)
}

func initBone() {
	h := bone.New()
	h.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, World")
	}))
	h.Get("/:name", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s", bone.GetValue(r, "name"))
	}))
	registerHandler("bone", h)
}
