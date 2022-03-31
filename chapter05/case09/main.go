package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	1. 类型 rune 专门来存储 Unicode 码点（code point），它是 int32 的别名，相当于 UCS-4/UTF-32 编码格式。
	2. 使用单引号的字面量，其默认类型就是 rune
	3. 除 []rune 外，还可直接在 rune、byte、string 间进行转换
*/

func printType() {
	r := '我'
	fmt.Printf("%T\n", r)
}

func unicodeConvert() {
	r := '我'
	s := string(r) // rune to string
	b := byte(r)   // rune to byte

	s2 := string(b) // byte to string
	r2 := rune(b)   // byte to rune
	fmt.Println(s, b, s2, r2)
	fmt.Println(string(r2))
}

func isTrueString() {
	s := "雨痕"
	s = string(s[0:1] + s[3:4]) // 截取拼接一个不合法的字符串
	fmt.Println(s, utf8.ValidString(s))
}

func unicodeLen() {
	s := "雨.痕"
	println(len(s), utf8.RuneCountInString(s))
}

func main() {
	printType()
	unicodeConvert()
	isTrueString()
	unicodeLen()
}
