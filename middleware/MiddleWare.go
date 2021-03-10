package middleware

import (
	"github.com/cuihaoweb/dawn/types"
)

// MiddleWare 中间件
type MiddleWare struct {
	useGroupHandler map[string][]types.UseHandler
	useAnyHandler   []types.UseHandler
}

// AddGroupHandler add into useGroupHandler
func (m *MiddleWare) AddGroupHandler(url string, handler types.UseHandler) {
	if m.useGroupHandler == nil {
		m.useGroupHandler = make(map[string][]types.UseHandler)
	}

	m.useGroupHandler[url] = append(m.useGroupHandler[url], handler)
}

// AddAnyHandler add into useGroupHandler
func (m *MiddleWare) AddAnyHandler(handler types.UseHandler) {
	m.useAnyHandler = append(m.useAnyHandler, handler)
}

// GetGroupHandler 得到中间件函数
func (m *MiddleWare) GetGroupHandler(url string) []types.UseHandler {
	if m.useGroupHandler[url] == nil {
		return nil
	}

	return m.useGroupHandler[url]
}

// GetAnyHandler 获取中间件函数
func (m *MiddleWare) GetAnyHandler() []types.UseHandler {
	return m.useAnyHandler
}

// NewMiddleWare new
func NewMiddleWare() *MiddleWare {
	UseGroupHandler := make(map[string][]types.UseHandler)

	return &MiddleWare{useGroupHandler: UseGroupHandler}
}
