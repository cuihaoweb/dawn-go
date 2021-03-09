package middleware

import (
	"github.com/cuihaoweb/dawn/types"
)

// MiddleWare 中间件
type MiddleWare struct {
	useGroupHandler map[string]*UseGroupHandler
	useAnyHandler   *UseAnyHandler
}

// AddGroupHandler add into useGroupHandler
func (m *MiddleWare) AddGroupHandler(url string, handler types.UseHandler) {
	if m.useGroupHandler == nil {
		m.useGroupHandler = make(map[string]*UseGroupHandler)
	}

	if m.useGroupHandler[url] == nil {
		m.useGroupHandler[url] = NewUseGroupHandler()
	}

	m.useGroupHandler[url].handlers = append(m.useGroupHandler[url].handlers, handler)

	m.useGroupHandler[url].count++
}

// AddAnyHandler add into useGroupHandler
func (m *MiddleWare) AddAnyHandler(handler types.UseHandler) {
	m.useAnyHandler.handlers = append(m.useAnyHandler.handlers, handler)
	m.useAnyHandler.count++
}

// GetGroupHandler 得到中间件函数
func (m *MiddleWare) GetGroupHandler(url string) []types.UseHandler {
	if m.useGroupHandler[url] == nil {
		return nil
	}

	return m.useGroupHandler[url].handlers
}

// GetAnyHandler 获取中间件函数
func (m *MiddleWare) GetAnyHandler() []types.UseHandler {
	return m.useAnyHandler.handlers
}

// GetAnyCount xx
func (m *MiddleWare) GetAnyCount() int {
	return m.useAnyHandler.count
}

// GetGroupCount xx
func (m *MiddleWare) GetGroupCount(rootURL string) int {
	if m.useGroupHandler[rootURL] == nil {
		return 0
	}

	return m.useGroupHandler[rootURL].count
}

// ReduceGroupUse xx
func (m *MiddleWare) ReduceGroupUse(url string) {
	if m.useGroupHandler[url] == nil {
		return
	}

	m.useGroupHandler[url].count--
}

// ReduceAnyUse xx
func (m *MiddleWare) ReduceAnyUse() {
	m.useAnyHandler.count--
}

// GoAnyUse xx
func (m *MiddleWare) GoAnyUse() {
	m.useAnyHandler.count = len(m.useAnyHandler.handlers)
}

// GoGroupUse xx
func (m *MiddleWare) GoGroupUse(rootURL string) {
	if m.useGroupHandler[rootURL] == nil {
		return
	}

	m.useGroupHandler[rootURL].count = len(m.useGroupHandler[rootURL].handlers)
}

// NewMiddleWare new
func NewMiddleWare() *MiddleWare {
	UseGroupHandler := make(map[string]*UseGroupHandler)

	return &MiddleWare{
		useGroupHandler: UseGroupHandler,
		useAnyHandler:   NewUseAnyHandler(),
	}
}
