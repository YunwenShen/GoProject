package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 使用字节切片构造字符串
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
	// 字符串长度 和 字符长度
	fmt.Println("字符串长度", utf8.RuneCountInString(str))
	fmt.Println("字符长度", len(str))
}
