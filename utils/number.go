package utils

import (
	"strconv"
	"strings"
)

func AutoConvert(s string) interface{} {
	s = strings.TrimSpace(s)
	lowerS := strings.ToLower(s)

	// 判断是否是布尔值
	switch lowerS {
	case "true":
		return true
	case "false":
		return false
	}

	// 判断是否是整数
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}

	// 判断是否是浮点数
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	// 否则返回字符串
	return s
}
