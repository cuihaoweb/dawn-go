package dawn

import (
	"github.com/cuihaoweb/dawn/ctx"
	"github.com/cuihaoweb/dawn/route"
	"github.com/valyala/fasthttp"
)

// NewRoute 创建创建路由对象
func NewRoute() *route.Route {
	return route.NewRoute()
}

// Listen 监听服务
func Listen(addr string, router *route.Route) {
	fasthttp.ListenAndServe(addr, func(rw *fasthttp.RequestCtx) {
		c := ctx.NewCtx(rw)
		handlerFunc := router.Match(string(c.Path()), c)
		if handlerFunc == nil {
			c.WriteString("NOT FOUND")
			return
		}
		handlerFunc(c)
	})
}
