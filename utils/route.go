package utils

import (
	"regexp"
	"strings"
)

// IsRegularURL 判断是否是正则路由
func IsRegularURL(str string) bool {
	if strings.Contains(str, "*") {
		return true
	} else if strings.Contains(str, "/:") {
		return true
	}
	return false
}

// SplitURL 判断是否是正则路由
func SplitURL(str string) (string, string) {
	var rootURL []byte
	var endURL string
	var i int

	if str == "" || str == "/" || str == "*" {
		return str, ""
	}

	for i = 0; i < len(str); i++ {
		if i == 0 {
			continue
		}
		if str[i] != '/' {
			rootURL = append(rootURL, str[i])
		} else {
			break
		}
	}
	endURL = SubStr(str, i, -1)

	return string(rootURL), string(endURL)
}

// SubStr 字符串截取
func SubStr(s string, start int, end int) string {
	var str []byte

	if end < 0 {
		for i := start; i <= len(s)+end; i++ {
			str = append(str, s[i])
		}
	} else {
		for i := start; i < end; i++ {
			str = append(str, s[i])
		}
	}
	return string(str)
}

// GetU 匹配正则表达式
func GetU(route string) ([]string, string) {
	const head = "^(/.+?)/:"
	const end = "/:([a-zA-Z][a-zA-Z0-9]*)"
	var headReg = regexp.MustCompile(head)
	var endReg = regexp.MustCompile(end)
	var headRes = headReg.FindString(route)
	var endRes = endReg.FindAllStringSubmatch(route, -1)
	var data []string

	headRes = SubStr(headRes, 0, -3)

	for _, arr := range endRes {
		data = append(data, arr[1])
		headRes += `/(.+)`
	}

	return data, headRes
}

// FindU 匹配正则表达式
func FindU(route string, url string) map[string]interface{} {
	arr, exp := GetU(route)
	if len(arr) <= 0 {
		return nil
	}

	var data = make(map[string]interface{})
	var reg = regexp.MustCompile(exp)
	res := reg.FindStringSubmatch(url)

	for index, val := range arr {
		data[val] = res[index+1]
	}

	return data
}
