package ctx

import (
	"encoding/json"
	"reflect"

	"github.com/valyala/fasthttp"
)

// Ctx 上下文
type Ctx struct {
	*fasthttp.RequestCtx
	query map[string]interface{}
}

// SetQuery xx
func (c *Ctx) SetQuery(query map[string]interface{}) {
	c.query = query
}

// SetQueryVal xx
func (c *Ctx) SetQueryVal(key string, val interface{}) {
	if c.query == nil {
		c.query = make(map[string]interface{})
	}

	c.query[key] = val
}

// Query xx
func (c *Ctx) Query(key string) string {
	if c.query == nil || c.query[key] == nil {
		return ""
	}

	return c.query[key].(string)
}

// JSON xx
func (c *Ctx) JSON(data interface{}) {
	tp := reflect.TypeOf(data)
	kind := tp.Kind().String()
	if !(kind == "map" || kind == "struct") {
		panic("argument must be map || struct")
	}

	byte, _ := json.Marshal(data)

	c.Write(byte)
}

// NewCtx 构造函数
func NewCtx(ctx *fasthttp.RequestCtx) *Ctx {
	Query := make(map[string]interface{})

	return &Ctx{ctx, Query}
}
