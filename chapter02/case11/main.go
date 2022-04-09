/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 20:46
*/

package main

/*
	1. 可在函数代码块中定义常量，不曾使用的常量不会引发编译错误。
*/

func main() {
	const x = 123

	println(x)

	const y = 1.23 // 未使用，不会引发编译错误。

	{
		const x = "abc" // 在不同作用域定义同名常量
		println(x)
	}
}
