package utils

import (
	"testing"
)

func TestTemplateContentReplaceVariables(t *testing.T) {
	contnet, err := TemplateContentReplaceVariables("你好，${name}，订单号是${orderon}！", `{"${name}":"张三", "${orderon}":"123456"}`)
	t.Log("content:", contnet)
	if err != nil {
		t.Error("TemplateContentReplaceVariables: Expected true, got false")
	}
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// contnet, err := zutils.TemplateContentReplaceVariables("你好，${name}，订单号是${orderon}！", `{"${name}":"张三", "${orderon}":"123456"}`)
// fmt.Printf("TemplateContentReplaceVariables: %s, Extended Count: %v\n", contnet, err)
