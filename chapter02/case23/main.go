/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 23:45
*/

package main

import "fmt"

/*
	1. 使用浮点数，须注意小数点位的有效精度。
*/

func main() {
	var a float32 = 1.1234567899
	var b float32 = 1.12345678
	var c float32 = 1.123456781

	println(a, b, c)
	println(a == b, a == c)
	fmt.Printf("%v %v, %v\n", a, b, c)
}
