package route

import "github.com/valyala/fasthttp"

// handler 处理函数
type handler func(*fasthttp.RequestCtx)
