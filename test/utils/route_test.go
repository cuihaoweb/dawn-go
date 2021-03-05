package utils

import (
	"testing"

	"github.com/cuihaoweb/dawn/src/utils"
)

func TestIsRegularURL(t *testing.T) {
	d1 := utils.IsRegularURL("/hello/*")
	d2 := utils.IsRegularURL("/hello/:id")
	if d1 == false || d2 == false {
		t.Error(d1, d2)
	}
}
func TestSplitURL(t *testing.T) {
	d1, d11 := utils.SplitURL("/")
	d2, d22 := utils.SplitURL("/user/add")

	if d1 != "/" && d11 != "" && d2 != "book" && d22 != "add" {
		t.Error(d1, d11, d2, d22)
	}
}

func TestSubstr(t *testing.T) {
	d1 := utils.Substr("cuihao", 1, -1)

	if d1 != "uihao" {
		t.Error(d1)
	}
}
