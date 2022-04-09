/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 14:26
*/

package main

/*
	1. 对于粗心的新手，使用简短模式的时候，可能会造成意外错误。比如原本打算修改全局变量，结果变成重新定义同名局部变量。
*/

var x = 100

func main() {
	println(&x, x) // 全局变量
	x := "abc"
	println(&x, x) // 重新定义和初始化同名局部变量
}
