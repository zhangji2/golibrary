package encode

import (
	"log"
	"testing"
)

func TestUtf8ToGbk(t *testing.T) {
	var str_utf8_encode = "您好，世界！"
	str, err := Utf8ToGbk(str_utf8_encode)
	log.Println("Utf8ToGbk:", str)
	if err != nil {
		log.Println("err:", err)
		t.Fatal(err)
	}

	str, err = GbkToUtf8(str)
	log.Println("GbkToUtf8:", str)
	if err != nil {
		log.Println("err:", err)
		t.Fatal(err)
	}
}

func TestUtf8ToGbkIgnoreErr(t *testing.T) {
	var str_utf8_encode = "测试：𪚥𡃁㱿"
	str, err := Utf8ToGbkIgnoreErr(str_utf8_encode)
	log.Println("Utf8ToGbkIgnoreErr:", str)
	if err != nil {
		log.Println("err:", err)
		t.Fatal(err)
	}

	str, err = GbkToUtf8IgnoreErr(str)
	log.Println("GbkToUtf8IgnoreErr:", str)
	if err != nil {
		log.Println("err:", err)
		t.Fatal(err)
	}
}

// 在项目中使用：先引入包，再调函数。
// zencode "github.com/zhangji2/golibrary/encode"
// fmt.Printf("Utf8ToGbk: %v\n", zencode.Utf8ToGbk(str_utf8_encode))
// fmt.Printf("GbkToUtf8: %v\n", zencode.GbkToUtf8(str_utf8_encode))
