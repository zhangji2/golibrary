package utils

// 数组工具函数

// InArray 判断数组 a 是否包含字符串 str。
func InArray(a []string, str string) bool {
	if len(a) == 0 || str == "" {
		return false
	}
	for _, s := range a {
		if s == str {
			return true
		}
	}
	return false
}

// UniqueArray 数组去重 返回一个新的数组，包含 a 中的唯一元素。
func UniqueArray(a []string) []string {
	if len(a) == 0 {
		return []string{}
	}
	uniqueMap := make(map[string]struct{})
	for _, s := range a {
		uniqueMap[s] = struct{}{}
	}
	uniqueArray := make([]string, 0, len(uniqueMap))
	for s := range uniqueMap {
		uniqueArray = append(uniqueArray, s)
	}
	return uniqueArray
}
