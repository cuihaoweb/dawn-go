package route

import (
	"github.com/cuihaoweb/dawn/types"
)

// Group 路由组
type Group struct {
	route  *Route
	preURL string
}

// Get get请求
func (g *Group) Get(url string, handler types.Handler) *Group {
	// add url and handler to GetContainer
	g.route.GetContainer.Add(g.preURL+url, handler)
	return g
}

// Post post请求
func (g *Group) Post(url string, handler types.Handler) *Group {
	// add url and handler to PostContainer
	g.route.PostContainer.Add(g.preURL+url, handler)
	return g
}

// Delete delete请求
func (g *Group) Delete(url string, handler types.Handler) *Group {
	// add url and handler to DeleteContainer
	g.route.DeleteContainer.Add(g.preURL+url, handler)
	return g
}

// Put put请求
func (g *Group) Put(url string, handler types.Handler) *Group {
	// add url and handler to PutContainer
	g.route.PutContainer.Add(g.preURL+url, handler)
	return g
}

// Use 中间件
func (g *Group) Use(handler types.UseHandler) *Group {
	g.route.Middleware.AddGroupHandler(g.preURL, handler)
	return g
}

// NewGroup xx
func NewGroup(preURL string, route *Route) *Group {
	return &Group{route: route, preURL: preURL}
}
