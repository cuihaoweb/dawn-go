package route

import (
	"strings"

	"github.com/cuihaoweb/dawn/datastructure"
	"github.com/cuihaoweb/dawn/middleware"
	"github.com/cuihaoweb/dawn/types"
)

// Route 路由模块
type Route struct {
	Middleware      *middleware.MiddleWare
	GetContainer    *datastructure.DataStructure
	PostContainer   *datastructure.DataStructure
	DeleteContainer *datastructure.DataStructure
	PutContainer    *datastructure.DataStructure
}

// Get get请求
func (r *Route) Get(url string, handler types.Handler) *Route {
	// add url and handler to GetContainer
	r.GetContainer.Add(url, handler)
	return r
}

// Post post请求
func (r *Route) Post(url string, handler types.Handler) *Route {
	// add url and handler to PostContainer
	r.PostContainer.Add(url, handler)
	return r
}

// Delete delete请求
func (r *Route) Delete(url string, handler types.Handler) *Route {
	// add url and handler to DeleteContainer
	r.DeleteContainer.Add(url, handler)
	return r
}

// Put put请求
func (r *Route) Put(url string, handler types.Handler) *Route {
	// add url and handler to PutContainer
	r.PutContainer.Add(url, handler)
	return r
}

// MatchExactURL 路由匹配
func (r *Route) MatchExactURL(url string, method string) types.Handler {
	switch strings.ToLower(method) {
	case "get":
		return r.GetContainer.MatchExactURL(url)
	case "post":
		return r.PostContainer.MatchExactURL(url)
	case "put":
		return r.PutContainer.MatchExactURL(url)
	case "delete":
		return r.DeleteContainer.MatchExactURL(url)
	default:
		return nil
	}
}

// Use 中间件
func (r *Route) Use(handler types.UseHandler) *Route {
	r.Middleware.AddAnyHandler(handler)
	return r
}

// MatchRegURL 路由匹配
func (r *Route) MatchRegURL(url string, method string) (types.Handler, map[string]interface{}) {
	switch strings.ToLower(method) {
	case "get":
		return r.GetContainer.MatchRegURL(url)
	case "post":
		return r.PostContainer.MatchRegURL(url)
	case "put":
		return r.PutContainer.MatchRegURL(url)
	case "delete":
		return r.DeleteContainer.MatchRegURL(url)
	default:
		return nil, nil
	}
}

// Group xx
func (r *Route) Group(rootURL string) *Group {
	return NewGroup(rootURL, r)
}

// NewRoute 构造函数
func NewRoute() *Route {
	return &Route{
		Middleware:      middleware.NewMiddleWare(),
		GetContainer:    datastructure.NewDataStructure(),
		PostContainer:   datastructure.NewDataStructure(),
		DeleteContainer: datastructure.NewDataStructure(),
		PutContainer:    datastructure.NewDataStructure(),
	}
}
