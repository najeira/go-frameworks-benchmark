package main

import (
	"fmt"
	"github.com/VividCortex/siesta"
"net/http"
)

func init() {
	calcMem("siesta", initSiesta)
}

func initSiesta() {
	h := siesta.NewService("/")
	h.Route("GET", "/", "",
		func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, World")
	})
	h.Route("GET", "/:name", "",
		func(w http.ResponseWriter, r *http.Request) {
		var params siesta.Params
		name := params.String("name", "", "")
		params.Parse(r.Form)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s", *name)
	})
	registerHandler("siesta", h)
}
