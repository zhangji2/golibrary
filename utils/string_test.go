package utils

import (
	"testing"
)

var str_a = "apple,banana,apple,cherry,banana"

func TestHasPrefix(t *testing.T) {
	t.Log("HasPrefix:", HasPrefix(str_a, "apple,"))
	t.Error("HasPrefix:", HasPrefix(str_a, "banana,"))
}

func TestInString(t *testing.T) {
	t.Log("InString:", InString(str_a, "apple"))
	t.Error("InString:", InString(str_a, "orange"))
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// fmt.Printf("InString: %v\n", zutils.InString(str_a, "apple"))
// fmt.Printf("HasPrefix: %v\n", zutils.HasPrefix(str_a, "apple,"))
