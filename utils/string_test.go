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

func TestIsGSM7Encode(t *testing.T) {
	isGSM7, gsm7ExtendedCount := IsGSM7Encode(str_a + "€{]中")
	t.Log("isGSM7:", isGSM7)
	t.Error("gsm7ExtendedCount:", gsm7ExtendedCount)
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// fmt.Printf("InString: %v\n", zutils.InString(str_a, "apple"))
// fmt.Printf("HasPrefix: %v\n", zutils.HasPrefix(str_a, "apple,"))
// isGSM7, gsm7ExtendedCount := zutils.IsGSM7Encode(str_a + "€{]中")
// fmt.Printf("IsGSM7Encode: %v, Extended Count: %d\n", isGSM7, gsm7ExtendedCount)
