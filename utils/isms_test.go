package utils

import (
	"testing"
)

func TestIsGSM7Encode(t *testing.T) {
	isGSM7, gsm7ExtendedCount := IsGSM7Encode("€{]")
	t.Log("isGSM7:", isGSM7)
	t.Log("gsm7ExtendedCount:", gsm7ExtendedCount)
	if !isGSM7 {
		t.Error("IsGSM7Encode: Expected true, got false")
	}
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// isGSM7, gsm7ExtendedCount := zutils.IsGSM7Encode("€{]中")
// fmt.Printf("IsGSM7Encode: %v, Extended Count: %d\n", isGSM7, gsm7ExtendedCount)
