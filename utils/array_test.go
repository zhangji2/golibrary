package utils

import (
	"testing"
)

var test_arr_a = []string{"apple", "banana", "apple", "cherry", "banana", "orange"}

func TestInArrayStr(t *testing.T) {
	t.Log("InArray:", InArrayStr(test_arr_a, "apple"))
	if !InArrayStr(test_arr_a, "orange") {
		t.Error("InArray:", "orange should not be in the array")
	}
}

func TestUniqueArray(t *testing.T) {
	t.Log("UniqueArray:", UniqueArray(test_arr_a))
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// fmt.Printf("InArrayStr: %v\n", zutils.InArrayStr(test_arr_a, "apple"))
// fmt.Printf("UniqueArray: %v\n", zutils.UniqueArray(test_arr_a))
