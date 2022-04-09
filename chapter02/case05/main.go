/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 14:36
*/

package main

/*
	1. 退化赋值的前提条件是：最少有一个新变量被定义，且必须是同一作用域。
*/

func main() {
	x := 100
	println(&x)

	//x := 200 // 报错
	x = 200 // 正确
	println(&x, x)
}
