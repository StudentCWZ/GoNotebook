/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 14:20
*/

package main

import "fmt"

/*
	1. 除 var 关键字外，还可使用更加简短的变量定义和初始化语法。
	2. 只是要注意，简短模式有些限制：
       - 定义变量，同时显式初始化
       - 不能提供数据类型
       - 只能用在函数内部
*/

func main() {
	x := 100
	a, s := 1, "abc"

	fmt.Println(x, a, s)
}
