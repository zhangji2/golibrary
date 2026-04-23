package utils

import (
	"testing"
)

var arr_a = []string{"apple", "banana", "apple", "cherry", "banana"}

func TestInArray(t *testing.T) {
	t.Log("InArray:", InArray(arr_a, "apple"))
	t.Error("InArray:", InArray(arr_a, "orange"))
}

func TestUniqueArray(t *testing.T) {
	t.Log("UniqueArray:", UniqueArray(arr_a))
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// fmt.Printf("InArray: %v\n", zutils.InArray(arr_a, "apple"))
// fmt.Printf("UniqueArray: %v\n", zutils.UniqueArray(arr_a))
