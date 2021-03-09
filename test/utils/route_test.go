package utils

import (
	"fmt"
	"testing"

	"github.com/cuihaoweb/dawn/utils"
)

func TestIsRegularURL(t *testing.T) {
	d1 := utils.IsRegularURL("/hello/*")
	d2 := utils.IsRegularURL("/hello/:id")
	if d1 == false || d2 == false {
		t.Error(d1, d2)
	}
}
func TestSplitRootURL(t *testing.T) {
	d1 := utils.SplitRootURL("/")
	d2 := utils.SplitRootURL("/user/add")

	if d1 != "/" || d2 != "/user" {
		t.Error(d1, d2)
	}
}
func TestSplitEndURL(t *testing.T) {
	d1 := utils.SplitEndURL("/")
	d2 := utils.SplitEndURL("/user/add")

	if d1 != "/" || d2 != "/add" {
		t.Error(d1, d2)
	}
}

func TestSubstr(t *testing.T) {
	d1 := utils.SubStr("cuihao", 1, -1)

	if d1 != "uihao" {
		t.Error(d1)
	}
}
func TestGetU(t *testing.T) {
	res, exp := utils.GetU("/book/:id/:name")
	fmt.Println(res, exp)
}

func TestFindU(t *testing.T) {
	res := utils.FindU("/hello/:id/:name", "/hello/1/李白")
	fmt.Println(res)
}
