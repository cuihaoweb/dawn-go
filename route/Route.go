package route

import (
	"strings"

	"github.com/cuihaoweb/dawn/ctx"
)

// Route 路由模块
type Route struct {
	preURL          string
	getContainer    *DataStructure
	postContainer   *DataStructure
	deleteContainer *DataStructure
	putContainer    *DataStructure
}

// Get get请求
func (r *Route) Get(url string, handler handler) *Route {
	// add url and handler to GetContainer
	r.getContainer.Add(r.preURL+url, handler)
	return r
}

// Post post请求
func (r *Route) Post(url string, handler handler) *Route {
	// add url and handler to PostContainer
	r.postContainer.Add(r.preURL+url, handler)
	return r
}

// Delete delete请求
func (r *Route) Delete(url string, handler handler) *Route {
	// add url and handler to DeleteContainer
	r.deleteContainer.Add(r.preURL+url, handler)
	return r
}

// Put put请求
func (r *Route) Put(url string, handler handler) *Route {
	// add url and handler to PutContainer
	r.putContainer.Add(r.preURL+url, handler)
	return r
}

// RouteGroup 路由组
func (r *Route) RouteGroup(preURL string) *Route {
	r.preURL += preURL
	return r
}

// Match 路由匹配
func (r *Route) Match(url string, ctx *ctx.Ctx) handler {
	method := string(ctx.Method())

	switch strings.ToLower(method) {
	case "get":
		return r.getContainer.Match(url, ctx)
	case "post":
		return r.postContainer.Match(url, ctx)
	case "put":
		return r.putContainer.Match(url, ctx)
	case "delete":
		return r.deleteContainer.Match(url, ctx)
	default:
		return nil
	}
}

// NewRoute 构造函数
func NewRoute() *Route {
	return &Route{
		preURL:          "",
		getContainer:    NewDataStructure(),
		postContainer:   NewDataStructure(),
		deleteContainer: NewDataStructure(),
		putContainer:    NewDataStructure(),
	}
}
