package test

import (
	"testing"

	"github.com/cuihaoweb/dawn/ctx"
)

var c = &ctx.Ctx{}

func TestJSON(t *testing.T) {
	c.JSON(map[string]interface{}{"name": "李白", "age": 123})
}
