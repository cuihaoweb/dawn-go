package types

import "github.com/cuihaoweb/dawn/ctx"

// UseHandler 中间件函数
type UseHandler func(*ctx.Ctx, Next)
