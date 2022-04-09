/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 22:04
*/

package main

/*
	1. 在实际编码中，建议用自定义类型实现用途明确的枚举类型。但这并不能将取值范围限制在预定义的枚举值内。
*/

type color byte // 自定义类型

const (
	black color = iota // 指定常量类型
	red
	blue
)

func test(c color) {
	println(c)
}

func main() {
	test(red)
	test(blue)
	test(100) // 100 并未超出 color/byte 类型取值范围

	x := 2
	//test(x)	// 报错：类型错误
	println(x)
}
