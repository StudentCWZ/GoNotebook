/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/8 09:36
*/

package main

/*
	1. 函数是第一类型，可作为参数或返回值
*/

func test(x int) func() { // 返回函数类型
	return func() { // 匿名函数
		println(x) // 闭包
	}
}

func main() {
	x := 100
	f := test(x)
	f()
}
