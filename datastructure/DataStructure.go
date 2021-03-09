package datastructure

import (
	"github.com/cuihaoweb/dawn/types"
	"github.com/cuihaoweb/dawn/utils"
)

// DataStructure 数据结构
type DataStructure struct {
	ExactURL   map[string]types.Handler
	RegularURL []*RegularURL
}

// Add xx
func (d *DataStructure) Add(url string, handle types.Handler) {
	if utils.IsRegularURL(url) {
		// 添加到正则路由
		d.RegularURL = append(d.RegularURL, NewRegularURL(url, handle))
	} else {
		// 添加到精准路由
		if d.ExactURL == nil {
			d.ExactURL = make(map[string]types.Handler)
		}
		d.ExactURL[url] = handle
	}
}

// MatchExactURL xx
func (d *DataStructure) MatchExactURL(url string) types.Handler {
	// 到精准路由中匹配
	if handle, ok := d.ExactURL[url]; ok {
		return handle
	}

	return nil
}

// MatchRegURL xx
func (d *DataStructure) MatchRegURL(url string) (types.Handler, map[string]interface{}) {
	// 到正则路由中匹配
	for _, val := range d.RegularURL {
		params := utils.FindU(val.url, url)
		if params != nil {
			return val.handler, params
		}
	}

	return nil, nil
}

// NewDataStructure 创建数据结构
func NewDataStructure() *DataStructure {
	var ExactURLs = make(map[string]types.Handler)
	return &DataStructure{
		ExactURL: ExactURLs,
	}
}
