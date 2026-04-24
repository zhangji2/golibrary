package encode

import (
	"testing"
)

func TestUtf8ToGbk(t *testing.T) {
	var str_utf8_encode = "您好，世界！"
	str, err := Utf8ToGbk(str_utf8_encode)
	t.Log("Utf8ToGbk:", str)
	if err != nil {
		t.Error("err:", err)
		t.Fatal(err)
	}

	str, err = GbkToUtf8(str)
	t.Log("GbkToUtf8:", str)
	if err != nil {
		t.Error("err:", err)
		t.Fatal(err)
	}
}

func TestUtf8ToGbkIgnoreErr(t *testing.T) {
	var str_utf8_encode = "测试：𪚥𡃁㱿" //在gbk中不存在的字会丢失，替换为问号。
	str, err := Utf8ToGbkIgnoreErr(str_utf8_encode)
	t.Log("Utf8ToGbkIgnoreErr:", str)
	if err != nil {
		t.Error("err:", err)
		t.Fatal(err)
	}

	str, err = GbkToUtf8(str)
	t.Log("GbkToUtf8IgnoreErr:", str)
	if err != nil {
		t.Error("err:", err)
		t.Fatal(err)
	}
}

// 在项目中使用：先引入包，再调函数。
// zencode "github.com/zhangji2/golibrary/encode"
// fmt.Printf("Utf8ToGbk: %v\n", zencode.Utf8ToGbk(str_utf8_encode))
// fmt.Printf("GbkToUtf8: %v\n", zencode.GbkToUtf8(str_utf8_encode))
