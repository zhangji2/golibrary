package encode

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// Utf-8编码转换函数

// Utf8ToGbk utf8转换gbk
func Utf8ToGbk(str string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "nil", e
	}
	return string(d), nil
}

// GbkToUtf8 gbk转换utf8
func GbkToUtf8(str string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "ioutil.ReadAll", e
	}
	return string(d), nil
}

// Utf8ToGbkIgnoreErr utf8转换gbk
// 一个一个字符转换忽视转换错误
// 错误字符不转换
func Utf8ToGbkIgnoreErr(str string) (string, error) {
	runeChar := []rune(str)
	var buffer bytes.Buffer
	for _, row := range runeChar {
		new, err := Utf8ToGbk(string(row))
		if err != nil {
			// fmt.Println("change character encoding error, char:", string(row), "error:", err)
			// return "change character encoding error", err
			buffer.WriteString("?")
		} else {
			buffer.WriteString(string(new))
		}
	}
	return buffer.String(), nil
}

// GbkToUtf8IgnoreErr utf8转换gbk
// 一个一个字符转换忽视转换错误
// 错误字符不转换
func GbkToUtf8IgnoreErr(str string) (string, error) {
	runeChar := []rune(str)
	var buffer bytes.Buffer
	for _, row := range runeChar {
		new, err := GbkToUtf8(string(row))
		if err != nil {
			// fmt.Println("change character encoding error, char:", string(row), "error:", err)
			// return "change character encoding error", err
			buffer.WriteString("?")
		} else {
			buffer.WriteString(string(new))
		}
	}
	return buffer.String(), nil
}
