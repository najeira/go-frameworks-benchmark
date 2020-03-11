package main

import (
	"fmt"

	"github.com/go-ozzo/ozzo-routing/v2"
)

func init() {
	calcMem("ozzo-routing", initOzzoRouting)
}

func initOzzoRouting() {
	r := routing.New()
	r.Get("/", func(c *routing.Context) error {
		c.Response.Header().Set("Content-Type", "text/plain; charset=utf-8")
		return c.Write("Hello, World")
	})
	r.Get("/<name>", func(c *routing.Context) error {
		c.Response.Header().Set("Content-Type", "text/plain; charset=utf-8")
		return c.Write(fmt.Sprintf("Hello, %s", c.Param("name")))
	})
	registerHandler("ozzo-routing", r)
}
