package main

import (
	"github.com/cuihaoweb/dawn"
	"github.com/cuihaoweb/dawn/ctx"
	"github.com/cuihaoweb/dawn/types"
)

func main() {
	app := dawn.New()
	router := app.Router()
	router.Use(func(ctx *ctx.Ctx, next types.Next) {
		println("use1中间件")
		ctx.WriteString("hello,word1")
		next()
	}).Use(func(ctx *ctx.Ctx, next types.Next) {
		println("use2中间件")
		ctx.WriteString("hello,word2")
		next()
	})

	book := router.Group("/book")
	book.Use(func(ctx *ctx.Ctx, next types.Next) {
		ctx.WriteString("hello,word")
		next()
	})
	book.Get("/hello", func(ctx *ctx.Ctx) {
		ctx.JSON(map[string]interface{}{"name": "李白", "age": 123})
	})

	people := router.Group("/people")
	people.Get("/:id", func(ctx *ctx.Ctx) {
		ctx.JSON(map[string]interface{}{"name": "杜甫", "age": 123})
	})

	dawn.Listen("127.0.0.1:8080", app)
}
