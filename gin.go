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
		c.String(200, "Hello, World")
	})
	registerHandler("gin", h)
}
