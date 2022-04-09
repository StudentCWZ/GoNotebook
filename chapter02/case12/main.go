/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 21:16
*/

package main

import "fmt"

/*
	1. 如果显式制定类型，必须确保常量左右值一致，需要时可做显式转换。右值不能超出常量的取值范围，否则会引发溢出错误。
*/

func main() {
	const (
		x, y int  = 99, -999
		b    byte = byte(x) // x被指定为 int 类型，须显式转换为 byte 类型
		//n = unit8(y)	// 报错
	)
	fmt.Println(x, y, b)
}
