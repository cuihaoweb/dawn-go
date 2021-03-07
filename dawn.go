package dawn

import "github.com/cuihaoweb/dawn/route"

// NewRoute 创建创建路由对象
func NewRoute() *route.Route {
	return &route.Route{
		GetContainer:    route.NewDataStructure(),
		PostContainer:   route.NewDataStructure(),
		DeleteContainer: route.NewDataStructure(),
		PutContainer:    route.NewDataStructure(),
	}
}
