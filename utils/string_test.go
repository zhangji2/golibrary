package utils

import (
	"testing"
)

var test_str_a = "apple,banana,apple,cherry,banana"

func TestHasPrefix(t *testing.T) {
	t.Log("HasPrefix:", HasPrefix(test_str_a, "apple,"))
	if !HasPrefix(test_str_a, "apple,") {
		t.Error("HasPrefix:", "apple, should be a prefix")
	}
}

func TestInString(t *testing.T) {
	t.Log("InString:", InString(test_str_a, "apple"))
	if !InString(test_str_a, "apple") {
		t.Error("InString:", "apple should be in the string")
	}
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// fmt.Printf("InString: %v\n", zutils.InString(test_str_a, "apple"))
// fmt.Printf("HasPrefix: %v\n", zutils.HasPrefix(test_str_a, "apple,"))
