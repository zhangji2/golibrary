package encode

import "encoding/base64"

// 数组工具函数

// base64 加密函数，a 为加密的字符串，返回加密后的字符串。
func Base64Encode(s string) string {
	encrypted := []byte(s)
	b64 := base64.StdEncoding.EncodeToString(encrypted)
	return b64
}

// base64 解密函数，a 为解密的字符串，返回解密后的字符串。
func Base64Decode(s string) string {
	decrypted, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(decrypted)
}
