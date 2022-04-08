/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 09:06
*/

package main

/*
	1. 使用 var 定义变量，支持类型推断。
	2. 基础数据类型划分清晰明确，有助于编写跨平台应用。
	3. 编译器确保变量总是被初始化为零值，避免出现意外状况。
	4. 在函数内部，还可省略 var 关键字，使用简单的定义模式。
*/

func main() {
	var x int32
	var s = "hello, world!"

	println(x, s)

	y := 100 // 注意，赋值符号不同
	println(y)
}
