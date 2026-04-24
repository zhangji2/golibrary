package encode

import (
	"testing"
)

var str_a_encode = "apple,banana,apple,cherry,banana"
var str_a_key = "1234567890abcdef" // 16, 24, 32 字节的密钥

func TestAesEncryptCBC(t *testing.T) {
	encrypted := AesEncryptCBC(str_a_encode, str_a_key)
	t.Log("encrypted:", encrypted)
}

func TestAesDecryptCBC(t *testing.T) {
	var str_a_decode = "842977723c16d4d4af39ac3bf34070fe8ca2ff7681f929dc4742fd5284ad3b933b6e3e3a2d29b1985c7e387f94faca43"
	decrypted, err := AesDecryptCBC(str_a_decode, str_a_key)
	t.Log("decrypted:", decrypted)
	if err != nil {
		t.Error("err:", err)
		t.Fatal(err)
	}
}

func TestAesEncryptECB(t *testing.T) {
	encrypted := AesEncryptECB(str_a_encode, str_a_key)
	t.Log("encrypted:", encrypted)
}

func TestAesDecryptECB(t *testing.T) {
	var str_a_decode = "961394a2f62ee87c47f1ddc49e107523a982ef256a4e6f49f112aeed36b23e7e02dec5652c0215f93af69f009f0f6d58"
	decrypted, err := AesDecryptECB(str_a_decode, str_a_key)
	t.Log("decrypted:", decrypted)
	if err != nil {
		t.Error("err:", err)
		t.Fatal(err)
	}
}

func TestAesEncryptCFB(t *testing.T) {
	encrypted, err := AesEncryptCFB(str_a_encode, str_a_key)
	t.Log("encrypted:", encrypted)
	if err != nil {
		t.Error("err:", err)
		t.Fatal(err)
	}
}

func TestAesDecryptCFB(t *testing.T) {
	var str_a_decode = "35070fd51c366f13b27dea798c4ff33f9e4529f845d7304a82587afd2bbf1a5aa24a6b3ac4bd2b9c615279c54cd11499"
	decrypted, err := AesDecryptCFB(str_a_decode, str_a_key)
	t.Log("decrypted:", decrypted)
	if err != nil {
		t.Error("err:", err)
		t.Fatal(err)
	}
}

// 在项目中使用：先引入包，再调函数。
// zencode "github.com/zhangji2/golibrary/encode"
// encrypted, err := zencode.AesEncryptCBC(str_a_encode, str_a_key)
// fmt.Printf("AesEncryptCBC: %s %v\n", encrypted, err)
// var str_a_decode = "842977723c16d4d4af39ac3bf34070fe8ca2ff7681f929dc4742fd5284ad3b933b6e3e3a2d29b1985c7e387f94faca43"
// decrypted, err := zencode.AesDecryptCFB(str_a_decode, str_a_key)
// fmt.Printf("AesDecryptCBC: %s %v\n", decrypted, err)
