package encode

import (
	"crypto/md5"
	"encoding/hex"
)

// 数组工具函数

// md5 加密函数，a 为加密的字符串，len 为加密后的长度 16/32，返回加密后的字符串。
func Md5Encode(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
