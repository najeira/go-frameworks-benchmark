package main

import (
	"github.com/gin-gonic/gin"
)

func init() {
	calcMem("gin", initGin)
}

func initGin() {
	gin.SetMode(gin.ReleaseMode)
	h := gin.New()
	h.GET("/", func(c *gin.Context) {
		// gin writes Content-Type internally
		c.String(200, "Hello, World")
	})
	h.GET("/:name", func(c *gin.Context) {
		// gin writes Content-Type internally
		name := c.Param("name")
		c.String(200, "Hello, %s", name)
	})
	registerHandler("gin", h)
}
