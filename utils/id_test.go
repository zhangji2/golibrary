package utils

import (
	"testing"
)

func TestGetGuid(t *testing.T) {
	t.Log("GetGuid:", GetGuid())
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// fmt.Printf("GetGuid: %v\n", zutils.GetGuid())
