package test

import (
	"fmt"
	"testing"

	"github.com/cuihaoweb/dawn"
	"github.com/valyala/fasthttp"
)

func TestGet(t *testing.T) {
	route := dawn.NewRoute().Get("/hello/book", func(ctx *fasthttp.RequestCtx) {
		ctx.Path()
		fmt.Println("匹配成功")
	}).Get("/hello/id", func(ctx *fasthttp.RequestCtx) {
		ctx.Path()
	}).Post("/hello/:id", func(ctx *fasthttp.RequestCtx) {
		ctx.Path()
	}).Post("/hello/id", func(ctx *fasthttp.RequestCtx) {
		ctx.Path()
	})
	route.GetContainer.Match("/hello/book")(&fasthttp.RequestCtx{})
}
