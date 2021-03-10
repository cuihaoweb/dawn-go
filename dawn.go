package dawn

import (
	"errors"

	"github.com/cuihaoweb/dawn/ctx"
	"github.com/cuihaoweb/dawn/route"
	"github.com/cuihaoweb/dawn/types"
	"github.com/cuihaoweb/dawn/utils"
	"github.com/valyala/fasthttp"
)

// Dawn 入口程序
type Dawn struct {
	route *route.Route
}

// Router xx
func (d *Dawn) Router() *route.Route {
	return d.route
}

// match 匹配路由
func (d *Dawn) match(path string, method string, ctx *ctx.Ctx) (types.Handler, map[string]interface{}) {
	// 匹配路由
	var handlerFunc types.Handler
	var query map[string]interface{}
	if handlerFunc = d.route.MatchExactURL(path, method); handlerFunc == nil {
		if handlerFunc, query = d.route.MatchRegURL(path, method); handlerFunc == nil {
			ctx.WriteString("NOT FOUND")
			return nil, nil
		}
	}

	return handlerFunc, query
}

func (d *Dawn) emitAnyUse(ctx *ctx.Ctx) error {
	anyHandlerFunc := d.route.Middleware.GetAnyHandler()
	count := len(anyHandlerFunc)
	for index, val := range anyHandlerFunc {
		val(ctx, func() { count-- })

		if count+index+1 != len(anyHandlerFunc) {
			return errors.New("")
		}
	}

	return nil
}

func (d *Dawn) emitGroupUse(rootURL string, ctx *ctx.Ctx) error {
	var groupHandlerFunc []types.UseHandler

	if groupHandlerFunc = d.route.Middleware.GetGroupHandler(rootURL); groupHandlerFunc == nil {
		return nil
	}

	count := len(groupHandlerFunc)
	for index, val := range groupHandlerFunc {
		val(ctx, func() { count-- })

		if count+index+1 != len(groupHandlerFunc) {
			return errors.New("")
		}
	}

	return nil
}

// New 创建创建路由对象
func New() *Dawn {
	Route := route.NewRoute()
	return &Dawn{route: Route}
}

// Listen 监听服务
func Listen(addr string, app *Dawn) {
	fasthttp.ListenAndServe(addr, func(rw *fasthttp.RequestCtx) {
		c := ctx.NewCtx(rw)
		path := string(c.Path())
		method := string(c.Method())
		rootURL := utils.SplitRootURL(path)

		// 匹配路由
		var handlerFunc types.Handler
		var query map[string]interface{}
		if handlerFunc, query = app.match(path, method, c); handlerFunc == nil {
			c.WriteString("NOT FOUND")
			return
		}
		if query != nil {
			c.SetQuery(query)
		}

		// 执行中间件
		if err := app.emitAnyUse(c); err != nil {
			return
		}
		if err := app.emitGroupUse(rootURL, c); err != nil {
			return
		}

		handlerFunc(c)
	})
}
