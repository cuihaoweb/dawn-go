package main

import (
	"github.com/cuihaoweb/dawn"
	"github.com/cuihaoweb/dawn/ctx"
)

func main() {
	route := dawn.NewRoute()

	route.RouteGroup("/book")
	route.Get("/hello", func(ctx *ctx.Ctx) {
		ctx.JSON(map[string]interface{}{"name": "李白", "age": 123})
	})

	dawn.Listen("127.0.0.1:8080", route)
}
