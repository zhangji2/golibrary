package utils

import (
	"bytes"
	"math/rand"
	"strings"
	"time"
)

// 随机相关函数

// 创建一个全局独立的随机数生成器，只初始化一次
var ZRandomInt = rand.New(rand.NewSource(time.Now().UnixNano()))
var ZRandomInt64 = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomInt 随机生成区间数值
func RandomInt(min, max int) int {
	return ZRandomInt.Intn(max-min) + min
}

// RandomInt64 生成一个区间的随机数
func RandomInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return ZRandomInt64.Int63n(max-min) + min
}

// RandomString 随机生成字符串
// RandomString(8, "A")
// RandomString(8, "a0")
// RandomString(20, "Aa0")
func RandomString(randLength int, randType string) (result string) {
	var num = "0123456789"
	var lower = "abcdefghijklmnopqrstuvwxyz"
	var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := bytes.Buffer{}
	if strings.Contains(randType, "0") {
		b.WriteString(num)
	}
	if strings.Contains(randType, "a") {
		b.WriteString(lower)
	}
	if strings.Contains(randType, "A") {
		b.WriteString(upper)
	}
	var str = b.String()
	var strLen = len(str)
	if strLen == 0 {
		result = ""
		return
	}

	rand.Seed(time.Now().UnixNano())
	b = bytes.Buffer{}
	for i := 0; i < randLength; i++ {
		b.WriteByte(str[rand.Intn(strLen)])
	}
	result = b.String()
	return
}
