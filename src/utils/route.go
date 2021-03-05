package utils

import (
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
	endURL = Substr(str, i, -1)

	return string(rootURL), string(endURL)
}

// Substr 字符串截取
func Substr(s string, start int, end int) string {
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
