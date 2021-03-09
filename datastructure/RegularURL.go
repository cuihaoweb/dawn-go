package datastructure

import "github.com/cuihaoweb/dawn/types"

// RegularURL 正则路由
type RegularURL struct {
	url     string
	handler types.Handler
}

// NewRegularURL xx
func NewRegularURL(url string, handler types.Handler) *RegularURL {
	return &RegularURL{
		url:     url,
		handler: handler,
	}
}
