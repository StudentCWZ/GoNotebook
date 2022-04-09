/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 14:32
*/

package main

/*
	1. 简短模式并不总是重新定义变量，也可能是部分退化的赋值操作。
*/

func main() {
	x := 100
	println(&x)

	x, y := 200, "abc"
	println(&x, x)
	println(y)
}
