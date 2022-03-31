package main

import "fmt"

/*
	1. 字符串是不可变字节（byte）序列，其本身是一个复合结构。
	2. 头部指针指向字节数组，但没有 NULL 结尾。
	3. 默认以 UTF-8 编码存储 Unicode 字符，字面量里允许使用十六进制、八进制和 UTF 编码格式。
	4. 注意：内置函数 len 返回字节数组长度，cap 不接受字符串类型参数。
*/

// type stringStruct struct {
// 	str unsafe.Pointer
// 	len int
// }

func main() {
	s := "雨痕\x61\142\u0041"
	fmt.Printf("%s\n", s)
	fmt.Printf("% x, len: %d\n", s, len(s))
}
