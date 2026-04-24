package utils

import (
	"testing"
)

func TestRecurisonListPath(t *testing.T) {
	slice := make([]string, 0)
	RecurisonListPath("/path/to/test", &slice)
	t.Log("RecurisonListPath:", slice)
}

func TestGetCurrPath(t *testing.T) {
	t.Log("GetCurrPath:", GetCurrPath())
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// fmt.Printf("GetCurrPath: %v\n", zutils.GetCurrPath())
