/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/12 09:07
*/

package main

/*
	1. 位移右操作数必须是无符号整数，或可以转换的无显式类型常量
*/

func main() {
	b := 23
	x := 1 << b
	println(x)
}
