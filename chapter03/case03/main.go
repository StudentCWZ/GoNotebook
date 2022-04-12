/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/12 09:11
*/

package main

import "fmt"

/*
	1. 如果是非常量位移表达式，那么会优先将无显式类型的常量左操作转型。
*/

func main() {
	a := 1.0 << 3                // 常量表达式
	fmt.Printf("%T, %v\n", a, a) // int, 8

	var s uint = 3
	// b := 1.0 << s	// 报错
	fmt.Println(s)

	var c int32 = 1.0 << s       // 自动将 1.0 转换为 int 32 类型
	fmt.Printf("%T, %v\n", c, c) // int32, 8
}
