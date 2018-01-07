package main

import (
	"fmt"

	"github.com/zpatrick/fireball"
)

func init() {
	calcMem("fireball", initFireball)
}

func initFireball() {
	routes := []*fireball.Route{
		&fireball.Route{
			Path: "/",
			Handlers: fireball.Handlers{
				"GET": func(c *fireball.Context) (fireball.Response, error) {
					return fireball.NewResponse(200, []byte("Hello, World"), map[string]string{
						"Content-Type": "text/plain; charset=utf-8",
					}), nil
				},
			},
		},
		&fireball.Route{
			Path: "/:name",
			Handlers: fireball.Handlers{
				"GET": func(c *fireball.Context) (fireball.Response, error) {
					name := c.PathVariables["name"]
					body := fmt.Sprintf("Hello, %s", name)
					return fireball.NewResponse(200, []byte(body), map[string]string{
						"Content-Type": "text/plain; charset=utf-8",
					}), nil
				},
			},
		},
	}
	app := fireball.NewApp(routes)
	registerHandler("fireball", app)
}
