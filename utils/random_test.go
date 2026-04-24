package utils

import (
	"testing"
)

func TestRandomInt(t *testing.T) {
	t.Log("RandomInt:", RandomInt(0, 100))
}

func TestRandomInt64(t *testing.T) {
	t.Log("RandomInt64:", RandomInt64(20, 100))
}

func TestRandomString(t *testing.T) {
	t.Log("RandomString1:", RandomString(8, "A"))
	t.Log("RandomString2:", RandomString(16, "a0"))
	t.Log("RandomString3:", RandomString(32, "Aa0"))
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// fmt.Printf("RandomInt: %d\n", zutil.RandomInt(0, 100))
// fmt.Printf("RandomInt64: %d\n", zutil.RandomInt64(0, 100))
// fmt.Printf("RandomString: %s\n", zutil.RandomString(8, "A"))
// fmt.Printf("RandomString: %s\n", zutil.RandomString(16, "a0"))
// fmt.Printf("RandomString: %s\n", zutil.RandomString(32, "Aa0"))
