package utils

// 国际短信工具函数

// IsGSM7Encode 判断编码类型（ASCII/GSM-7 vs Unicode）
func IsGSM7Encode(content string) (bool, int) {
	if content == "" {
		return true, 0
	}
	// GSM-7 扩展字符集‌在编码时需要使用转义机制‌，因此‌实际占用空间为 2 个 7 位字符的位置‌，即“按 2 个计算”。
	// 使用map提高查找效率
	gsm7ExtendedMap := map[rune]bool{
		'\f': true, '{': true, '}': true, '\\': true, '[': true, '~': true, ']': true, '|': true, '€': true,
	}
	// type gsm7ExtendedInfo struct {
	// 	Dec    int
	// 	Hex    int
	// 	Name   string
	// 	NameCn string
	// }
	// var gsm7ExtendedInfoMap = map[rune]gsm7ExtendedInfo{
	// 	'\f': {Dec: 12, Hex: 0x0C, Name: "Form Feed", NameCn: "换页符"},
	// 	'{':  {Dec: 123, Hex: 0x7B, Name: "Left Curly Bracket", NameCn: "左花括号"},
	// 	'}':  {Dec: 125, Hex: 0x7D, Name: "Right Curly Bracket", NameCn: "右花括号"},
	// 	'\\': {Dec: 92, Hex: 0x5C, Name: "Backslash", NameCn: "反斜杠"},
	// 	'[':  {Dec: 91, Hex: 0x5B, Name: "Left Square Bracket", NameCn: "左方括号"},
	// 	'~':  {Dec: 126, Hex: 0x7E, Name: "Tilde", NameCn: "波浪号"},
	// 	']':  {Dec: 93, Hex: 0x5D, Name: "Right Square Bracket", NameCn: "右方括号"},
	// 	'|':  {Dec: 124, Hex: 0x7C, Name: "Pipe", NameCn: "管道符"},
	// 	'€':  {Dec: 8364, Hex: 0x20AC, Name: "Euro Sign", NameCn: "欧元符号"},
	// }
	// fmt.Printf("GSM-7 Extended Characters Info:\n")
	// for char, info := range gsm7ExtendedInfoMap {
	// 	fmt.Printf("Character: '%c', Dec: %d, Hex: 0x%X, Name: %s, NameCn: %s\n",
	// 		char, info.Dec, info.Hex, info.Name, info.NameCn)
	// }
	isGSM7 := true
	gsm7ExtendedCount := 0
	// 遍历每个字符判断编码类型
	for _, char := range content {
		// 如果字符ASCII码 > 127 且不是€，则不是GSM-7编码
		// fmt.Printf("char: %c\n", char)
		if char > 127 && char != '€' {
			isGSM7 = false
			break
		}
		// 如果是扩展字符，增加扩展长度（使用map查找，比strings.ContainsRune更快）
		if gsm7ExtendedMap[char] {
			gsm7ExtendedCount++
		}
	}
	return isGSM7, gsm7ExtendedCount
}
