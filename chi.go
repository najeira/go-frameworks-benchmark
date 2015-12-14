package main

import (
	"fmt"
	"net/http"

	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

func init() {
	calcMem("chi", initChi)
}

func initChi() {
	h := chi.NewRouter()
	h.Get("/", func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, World")
	})
	h.Get("/:name", func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s", chi.URLParams(ctx)["name"])
	})
	registerHandler("chi", h)
}
