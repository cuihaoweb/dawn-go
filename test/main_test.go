package test

import (
	"fmt"
	"testing"

	"github.com/cuihaoweb/dawn"
	"github.com/cuihaoweb/dawn/ctx"
	"github.com/cuihaoweb/dawn/types"
	"github.com/valyala/fasthttp"
)

func TestMain(t *testing.T) {
	rw := &fasthttp.RequestCtx{}
	app := dawn.New()
	router := app.Router()
	router.Use(func(ctx *ctx.Ctx, next types.Next) {
		println("use中间件")
		next()
	})

	book := router.Group("/book")
	book.Get("/hello", func(ctx *ctx.Ctx) {
		ctx.JSON(map[string]interface{}{"name": "李白", "age": 123})
	})

	people := router.Group("/people")
	people.Get("/:id", func(ctx *ctx.Ctx) {
		ctx.JSON(map[string]interface{}{"name": "杜甫", "age": 123})
	})
	fmt.Println(app.Router().Middleware.GetAnyHandler())

	c := ctx.NewCtx(rw)
	path := "/book/1"
	method := "get"
	// rootURL := utils.SplitRootURL(path)

	// 匹配路由
	var handlerFunc types.Handler
	var query map[string]interface{}
	if handlerFunc = router.MatchExactURL(path, method); handlerFunc == nil {
		if handlerFunc, query = router.MatchRegURL(path, method); handlerFunc == nil {
			c.WriteString("NOT FOUND")
			return
		}
	}
	println(query)
	// 执行中间件
	if anyHandlerFunc := router.Middleware.GetAnyHandler(); anyHandlerFunc != nil {
		for index, val := range anyHandlerFunc {
			val(c, func() {
				router.Middleware.ReduceAnyUse()
				println("执行了next")
			})

			println(router.Middleware.GetAnyCount(), router.Middleware.GetAnyCount()+index)
			if router.Middleware.GetAnyCount()+index != len(anyHandlerFunc) {
				c.WriteString("")
				return
			}

			if router.Middleware.GetAnyCount() <= 0 {
				router.Middleware.GoAnyUse()
			}
		}
	}

	// if groupHandlerFunc := app.route.Middleware.GetGroupHandler(rootURL); groupHandlerFunc != nil {
	// 	for index, val := range groupHandlerFunc {
	// 		val(c, func() { app.route.Middleware.GetGroupHandler(rootURL) })
	// 		app.route.Middleware.ReduceGroupUse(rootURL)

	// 		if app.route.Middleware.GetGroupCount(rootURL)+index != len(groupHandlerFunc) {
	// 			return
	// 		}
	// 	}
	// }

	// router.EmitAnyUse(c)
	// router.EmitGroupUse(path, c)
	handlerFunc(c)
}
