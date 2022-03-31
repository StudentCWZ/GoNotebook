package main

import "fmt"

/*
	1. 支持 != 、==、<、>、+、+= 操作符
	2. 允许以索引号访问字节数组（非字符），但不能获取元素地址。
*/

func StringOperator() {
	s := "ab" +
		"cd"
	println(s == "abcd")
	println(s > "abc")
}

func StringIndex() {
	s := "abc"
	println(s[1])
	fmt.Printf("%c\n", s[1])
	// print(&s[1])	// 错误：不能获取字符串中元素地址
}

func main() {
	StringOperator()
	StringIndex()

}
