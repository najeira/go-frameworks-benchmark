package main

import (
	"fmt"

	gocraft "github.com/gocraft/web"
)

func init() {
	calcMem("gocraft", initGocraft)
}

func initGocraft() {
	var ctx struct{}
	h := gocraft.New(ctx)
	h.Get("/", func(w gocraft.ResponseWriter, r *gocraft.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, World")
	})
	h.Get("/:name", func(w gocraft.ResponseWriter, r *gocraft.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s", r.PathParams["name"])
	})
	registerHandler("gocraft", h)
}
