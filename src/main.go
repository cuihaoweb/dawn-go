package main

import (
	"strconv"

	"github.com/cuihaoweb/dawn"
	"github.com/cuihaoweb/dawn/ctx"
	"github.com/cuihaoweb/dawn/types"
)

func main() {
	app := dawn.New()
	router := app.Router()
	router.Use(func(ctx *ctx.Ctx, next types.Next) {
		ctx.WriteString("hello,word1")
		next()
	}).Use(func(ctx *ctx.Ctx, next types.Next) {
		ctx.WriteString("hello,word2")
		next()
	})

	book := router.Group("/book")
	book.Use(func(ctx *ctx.Ctx, next types.Next) {
		ctx.WriteString("hello,word")
		next()
	})
	book.Get("/hello", func(ctx *ctx.Ctx) {
		a := ctx.Query("name")
		ctx.JSON(map[string]interface{}{"name": a, "age": 123})
	})

	people := router.Group("/people")
	people.Get("/:id/:name", func(ctx *ctx.Ctx) {
		id, _ := strconv.Atoi(ctx.Query("id"))
		name := ctx.Query("name")
		ctx.JSON(map[string]interface{}{"name": name, "age": id})
	})

	dawn.Listen("127.0.0.1:8080", app)
}
