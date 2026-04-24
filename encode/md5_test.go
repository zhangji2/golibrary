package encode

import (
	"testing"
)

var str_a = "apple,banana,apple,cherry,banana"

func TestMd5Encode(t *testing.T) {
	t.Log("Md5Encode:", Md5Encode(str_a))
}

// 在项目中使用：先引入包，再调函数。
// zencode "github.com/zhangji2/golibrary/encode"
// fmt.Printf("Md5Encode: %v\n", zencode.Md5Encode(str_a))
