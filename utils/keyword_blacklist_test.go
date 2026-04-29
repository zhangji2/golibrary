package utils

import (
	"testing"
)

func TestCheckKeywordBlacklist(t *testing.T) {
	var kl = "apple,banana\napple,cherry,banana"
	c := "this is an apple"
	t.Log("content:", CheckKeywordBlacklist(kl, c))
}

func TestFindKeywordBlacklist(t *testing.T) {
	var kl = "apple,banana\napple,cherry,banana"
	c := "this is an apple"
	t.Log("content:", FindKeywordBlacklist(kl, c))
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// fmt.Printf("CheckKeywordBlacklist: %v\n", zutils.CheckKeywordBlacklist("apple,banana\napple,cherry,banana", "this is an apple"))
