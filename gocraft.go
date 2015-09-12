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
		fmt.Fprintf(w, "Hello, World")
	})
	registerHandler("gocraft", h)
}
