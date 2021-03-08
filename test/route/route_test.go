package test

import (
	"fmt"
	"testing"

	"github.com/cuihaoweb/dawn"
	"github.com/cuihaoweb/dawn/ctx"
)

func TestGet(t *testing.T) {
	dawn.NewRoute().Get("/hello/book", func(ctx *ctx.Ctx) {
		ctx.Path()
		fmt.Println("匹配成功")
	}).Get("/hello/id", func(ctx *ctx.Ctx) {
		ctx.Path()
	}).Post("/hello/:id", func(ctx *ctx.Ctx) {
		ctx.Path()
	}).Post("/hello/id", func(ctx *ctx.Ctx) {
		ctx.Path()
	})
	// route.GetContainer.Match("/hello/book")(&fasthttp.RequestCtx{})
}
