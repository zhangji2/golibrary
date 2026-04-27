package utils

import (
	"encoding/json"
	"strings"
	"unicode"
)

// 模板工具函数

func TemplateContentReplaceVariables(content string, varMapStr string) (string, error) {
	if len(content) == 0 {
		return content, nil
	}
	if len(varMapStr) == 0 {
		return content, nil
	}
	content = TemplateContentPreprocess(content)
	varMap := make(map[string]string)
	err := json.Unmarshal([]byte(varMapStr), &varMap)
	if err != nil {
		return "", err
	}
	for k, v := range varMap {
		content = strings.Replace(content, k, v, -1)
	}
	return content, nil
}

// TemplateContentPreprocess 内容预处理：统一换行符（\r\n→\n）、过滤违规特殊字符
func TemplateContentPreprocess(content string) string {
	// 统一换行符
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")

	// 过滤违规特殊字符（移除控制字符，保留可打印字符和换行符）
	var builder strings.Builder
	for _, r := range content {
		// 保留可打印字符、换行符、制表符
		if unicode.IsPrint(r) || r == '\n' || r == '\t' {
			builder.WriteRune(r)
		}
		// 其他控制字符（如 \x00-\x1F 除了 \n, \t）会被过滤掉
	}
	return builder.String()
}
