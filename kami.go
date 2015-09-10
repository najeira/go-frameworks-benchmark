package main

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"

	"github.com/guregu/kami"
)

func init() {
	calcMem("kami", initKami)
}

func initKami() {
	h := kami.New()
	h.Get("/", func(c context.Context, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, responseText)
	})
	registerHandler("kami", h)
}
