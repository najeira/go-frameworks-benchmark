package main

import (
	"fmt"

	"github.com/clevergo/clevergo"
)

func init() {
	calcMem("clevergo", initClevergo)
}

func initClevergo() {
	r := clevergo.NewRouter()
	r.Get("/", func(ctx *clevergo.Context) error {
		ctx.Response.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, err := ctx.WriteString("Hello, World")
		return err
	})
	r.Get("/:name", func(ctx *clevergo.Context) error {
		ctx.Response.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, err := ctx.WriteString(fmt.Sprintf("Hello, %s", ctx.Params.String("name")))
		return err
	})
	registerHandler("clevergo", r)
}
