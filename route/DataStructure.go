package route

import (
	"fmt"

	"github.com/cuihaoweb/dawn/ctx"
	"github.com/cuihaoweb/dawn/utils"
)

// DataStructure 数据结构
type DataStructure struct {
	ExactURL   map[string]handler
	RegularURL []*regularURL
}

// Add xx
func (d *DataStructure) Add(url string, handle handler) {
	if utils.IsRegularURL(url) {
		// 添加到正则路由
		d.RegularURL = append(d.RegularURL, NewRegularURL(url, handle))
	} else {
		// 添加到精准路由
		if d.ExactURL == nil {
			d.ExactURL = make(map[string]handler)
		}
		d.ExactURL[url] = handle
	}
}

// Match xx
func (d *DataStructure) Match(url string, ctx *ctx.Ctx) handler {
	if utils.IsRegularURL(url) {
		// 到正则路由中匹配
		for _, val := range d.RegularURL {
			params := utils.FindU(val.URL, url)
			if params != nil {
				fmt.Println(params)
				return val.Handler
			}
		}
		return nil
	} else {
		// 到精准路由中匹配
		handle, ok := d.ExactURL[url]
		if !ok {
			return nil
		}
		return handle
	}
}

// NewDataStructure 创建数据结构
func NewDataStructure() *DataStructure {
	var regularURLs []*regularURL
	var ExactURLs = make(map[string]handler)
	return &DataStructure{
		ExactURL:   ExactURLs,
		RegularURL: regularURLs,
	}
}

type regularURL struct {
	URL     string
	Handler handler
}

// NewRegularURL xx
func NewRegularURL(url string, handler handler) *regularURL {
	return &regularURL{
		URL:     url,
		Handler: handler,
	}
}
