package utils

import (
	"testing"
)

func TestIsEqual(t *testing.T) {
	t.Log("IsEqual:", IsEqual(1, 11))
}

func TestIsContains(t *testing.T) {
	t.Log("IsContains:", IsContains(1, 1))
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// fmt.Printf("IsEqual: %v\n", zutils.IsEqual(1, 11))
// fmt.Printf("IsContains: %v\n", zutils.IsContains(1, 1))
