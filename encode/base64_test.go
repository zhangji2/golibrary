package encode

import (
	"testing"
)

var str_base64_encode = "apple,banana,apple,cherry,banana"
var str_base64_decode = "YXBwbGUsYmFuYW5hLGFwcGxlLGNoZXJyeSxiYW5hbmE="

func TestBase64Encode(t *testing.T) {
	t.Log("Base64Encode:", Base64Encode(str_base64_encode))
}

func TestBase64Decode(t *testing.T) {
	t.Log("Base64Decode:", Base64Decode(str_base64_decode))
}

// 在项目中使用：先引入包，再调函数。
// zencode "github.com/zhangji2/golibrary/encode"
// fmt.Printf("Base64Encode: %v\n", zencode.Base64Encode(str_base64_encode))
// fmt.Printf("Base64Decode: %v\n", zencode.Base64Decode(str_base64_decode))
