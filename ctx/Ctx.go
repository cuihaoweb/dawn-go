package ctx

import (
	"encoding/json"
	"reflect"

	"github.com/valyala/fasthttp"
)

// Ctx 上下文
type Ctx struct {
	*fasthttp.RequestCtx
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
	return &Ctx{ctx}
}
