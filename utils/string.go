package utils

import "strings"

// 字符串工具函数

// HasPrefix 判断字符串 s 是否以前缀 prefix 开头。
func HasPrefix(s string, prefix string) bool {
	if s == "" || prefix == "" {
		return false
	}
	return strings.HasPrefix(s, prefix)
}

// HasSuffix 判断字符串 s 是否以后缀 suffix 结尾。
func HasSuffix(s string, suffix string) bool {
	if s == "" || suffix == "" {
		return false
	}
	return strings.HasSuffix(s, suffix)
}

// InString 判断字符串 s 是否包含子字符串 substr。
func InString(s string, substr string) bool {
	if s == "" || substr == "" {
		return false
	}
	return strings.Contains(s, substr)
}
