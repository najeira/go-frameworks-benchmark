package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func init() {
	calcMem("echo", initEcho)
}

func initEcho() {
	h := echo.New()
	h.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World")
	})
	h.GET("/:name", func(c echo.Context) error {
		return c.String(200, fmt.Sprintf("Hello, %s", c.Param("name")))
	})
	registerHandler("echo", h)
}
