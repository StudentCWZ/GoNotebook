/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 14:51
*/

package main

/*
	1. 进行多变量赋值操作时，首先计算出所有右值，然后再依次完成赋值操作。
*/

func main() {
	x, y := 1, 2
	x, y = y+3, x+2 // 先计算出右值 y+3 、x+2，然后再依次完成赋值操作。
	println(x, y)
}
