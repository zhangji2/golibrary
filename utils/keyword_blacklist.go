package utils

import "strings"

// 关键字工具函数

// DFA 节点
type keywordBlacklistNode map[rune]keywordBlacklistNode

// 初始化字典树
var keywordBlacklistRoot = make(keywordBlacklistNode)

// AddKeywordBlacklist 添加敏感词到字典树
func AddKeywordBlacklist(wordStr string) {
	wordStr = strings.Replace(wordStr, "\r\n", ",", -1)
	wordStr = strings.Replace(wordStr, "\n", ",", -1)
	words := strings.Split(wordStr, ",")
	for _, word := range words {
		p := keywordBlacklistRoot
		for _, r := range word {
			if _, ok := p[r]; !ok {
				p[r] = make(keywordBlacklistNode)
			}
			p = p[r]
		}
		// 结束标记
		p['#'] = nil
	}
}

// CheckKeywordBlacklist 检测是否含敏感词
func CheckKeywordBlacklist(keyword string, text string) bool {
	if len(keyword) == 0 || len(text) == 0 {
		return false
	}
	AddKeywordBlacklist(keyword)
	runes := []rune(text)
	length := len(runes)
	for i := 0; i < length; i++ {
		p := keywordBlacklistRoot
		for j := i; j < length; j++ {
			r := runes[j]
			if _, ok := p[r]; !ok {
				break
			}
			p = p[r]
			// 找到结束符，匹配到敏感词
			if _, ok := p['#']; ok {
				return true
			}
		}
	}
	return false
}

// FindKeywordBlacklist 找到第一个命中的敏感词并返回
func FindKeywordBlacklist(keyword string, text string) string {
	if len(keyword) == 0 || len(text) == 0 {
		return ""
	}
	AddKeywordBlacklist(keyword)
	runes := []rune(text)
	n := len(runes)
	for i := 0; i < n; i++ {
		p := keywordBlacklistRoot
		var word []rune
		for j := i; j < n; j++ {
			r := runes[j]
			if _, ok := p[r]; !ok {
				break
			}
			word = append(word, r)
			p = p[r]
			// 匹配到完整敏感词
			if _, ok := p['#']; ok {
				return string(word)
			}
		}
	}
	return ""
}
