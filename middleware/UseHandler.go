package middleware

import (
	"github.com/cuihaoweb/dawn/types"
)

// UseGroupHandler 路由组中间件
type UseGroupHandler struct {
	count    int
	handlers []types.UseHandler
}

// NewUseGroupHandler xx
func NewUseGroupHandler() *UseGroupHandler {
	return &UseGroupHandler{}
}

// UseAnyHandler 路由组中间件
type UseAnyHandler struct {
	count    int
	handlers []types.UseHandler
}

// NewUseAnyHandler xx
func NewUseAnyHandler() *UseAnyHandler {
	return &UseAnyHandler{}
}
